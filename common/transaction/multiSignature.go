package transaction

import (
	"bytes"
	"database/sql"
	"encoding/hex"
	"fmt"

	"github.com/zoobc/zoobc-core/common/accounttype"
	"github.com/zoobc/zoobc-core/common/blocker"
	"github.com/zoobc/zoobc-core/common/constant"
	"github.com/zoobc/zoobc-core/common/crypto"
	"github.com/zoobc/zoobc-core/common/fee"
	"github.com/zoobc/zoobc-core/common/model"
	"github.com/zoobc/zoobc-core/common/query"
	"github.com/zoobc/zoobc-core/common/util"
	"golang.org/x/crypto/sha3"
)

type (
	// MultiSignatureTransaction represent wrapper transaction type that require multiple signer to approve the transaction
	// wrapped
	MultiSignatureTransaction struct {
		ID              int64
		SenderAddress   []byte
		Fee             int64
		Body            *model.MultiSignatureTransactionBody
		NormalFee       fee.FeeModelInterface
		EscrowFee       fee.FeeModelInterface
		TransactionUtil UtilInterface
		TypeSwitcher    TypeActionSwitcher
		Signature       crypto.SignatureInterface
		Height          uint32
		BlockID         int64
		Escrow          *model.Escrow
		EscrowQuery     query.EscrowTransactionQueryInterface
		// multisig helpers
		MultisigUtil             MultisigTransactionUtilInterface
		SignatureInfoHelper      SignatureInfoHelperInterface
		MultisignatureInfoHelper MultisignatureInfoHelperInterface
		PendingTransactionHelper PendingTransactionHelperInterface
		// general helpers
		AccountBalanceHelper AccountBalanceHelperInterface
		TransactionHelper    TransactionHelperInterface
		QueryExecutor        query.ExecutorInterface
	}
	// SignatureInfoHelperInterface multisignature helpers
	SignatureInfoHelperInterface interface {
		InsertPendingSignature(
			pendingSignature *model.PendingSignature,
		) error
		InsertPendingSignatures(
			pendingSignatures []*model.PendingSignature,
		) (err error)
		GetPendingSignatureByTransactionHash(
			transactionHash []byte, txHeight uint32,
		) ([]*model.PendingSignature, error)
	}

	MultisignatureInfoHelperInterface interface {
		GetMultisigInfoByAddress(
			multisigInfo *model.MultiSignatureInfo,
			multisigAddress []byte,
			blockHeight uint32,
		) error
		InsertMultisignatureInfo(
			multisigInfo *model.MultiSignatureInfo,
		) error
		InsertMultiSignaturesInfo(multiSignaturesInfo []*model.MultiSignatureInfo) (err error)
	}

	PendingTransactionHelperInterface interface {
		InsertPendingTransaction(
			pendingTransaction *model.PendingTransaction,
		) error
		GetPendingTransactionByHash(
			pendingTransaction *model.PendingTransaction,
			pendingTransactionHash []byte,
			pendingTransactionStatuses []model.PendingTransactionStatus,
			blockHeight uint32,
			dbTx bool,
		) error
		GetPendingTransactionBySenderAddress(
			senderAddress []byte, txHeight uint32,
		) ([]*model.PendingTransaction, error)
		ApplyUnconfirmedPendingTransaction(pendingTransactionBytes []byte) error
		UndoApplyUnconfirmedPendingTransaction(pendingTransactionBytes []byte) error
		ApplyConfirmedPendingTransaction(
			pendingTransaction []byte, txHeight uint32, blockTimestamp int64,
		) (*model.Transaction, error)
	}

	SignatureInfoHelper struct {
		PendingSignatureQuery   query.PendingSignatureQueryInterface
		PendingTransactionQuery query.PendingTransactionQueryInterface
		QueryExecutor           query.ExecutorInterface
		Signature               crypto.SignatureInterface
	}

	MultisignatureInfoHelper struct {
		MultisignatureInfoQuery        query.MultisignatureInfoQueryInterface
		MultiSignatureParticipantQuery query.MultiSignatureParticipantQueryInterface
		QueryExecutor                  query.ExecutorInterface
	}

	PendingTransactionHelper struct {
		MultisignatureInfoQuery query.MultisignatureInfoQueryInterface
		PendingTransactionQuery query.PendingTransactionQueryInterface
		TransactionUtil         UtilInterface
		TypeSwitcher            TypeActionSwitcher
		QueryExecutor           query.ExecutorInterface
	}
)

func (pth *PendingTransactionHelper) GetPendingTransactionByHash(
	pendingTx *model.PendingTransaction,
	pendingTransactionHash []byte,
	pendingTransactionStatuses []model.PendingTransactionStatus,
	blockHeight uint32,
	dbTx bool,
) error {
	var (
		err error
	)
	q, args := pth.PendingTransactionQuery.GetPendingTransactionByHash(
		pendingTransactionHash, pendingTransactionStatuses,
		blockHeight, constant.MinRollbackBlocks,
	)
	row, _ := pth.QueryExecutor.ExecuteSelectRow(q, dbTx, args...)
	err = pth.PendingTransactionQuery.Scan(pendingTx, row)
	if err != nil {
		return err
	}
	return nil
}

func (pth *PendingTransactionHelper) GetPendingTransactionBySenderAddress(
	senderAddress []byte, txHeight uint32,
) ([]*model.PendingTransaction, error) {
	var pendingTxs []*model.PendingTransaction
	q, args := pth.PendingTransactionQuery.GetPendingTransactionsBySenderAddress(
		senderAddress, model.PendingTransactionStatus_PendingTransactionPending,
		txHeight, constant.MinRollbackBlocks,
	)
	pendingTxRows, err := pth.QueryExecutor.ExecuteSelect(q, false, args...)
	if err != nil {
		return nil, err
	}
	defer pendingTxRows.Close()
	_, err = pth.PendingTransactionQuery.BuildModel(pendingTxs, pendingTxRows)
	if err != nil {
		return nil, err
	}
	return pendingTxs, nil
}

func (pth *PendingTransactionHelper) InsertPendingTransaction(
	pendingTransaction *model.PendingTransaction,
) error {
	insertPendingTxQry := pth.PendingTransactionQuery.InsertPendingTransaction(pendingTransaction)
	err := pth.QueryExecutor.ExecuteTransactions(insertPendingTxQry)
	if err != nil {
		return blocker.NewBlocker(blocker.DBErr, err.Error())
	}
	return nil
}

func (pth *PendingTransactionHelper) ApplyUnconfirmedPendingTransaction(
	pendingTransactionBytes []byte,
) error {
	// parse and apply unconfirmed
	innerTx, err := pth.TransactionUtil.ParseTransactionBytes(pendingTransactionBytes, false)
	if err != nil {
		return blocker.NewBlocker(
			blocker.ValidationErr,
			"FailToParseTransactionBytes",
		)
	}
	innerTa, err := pth.TypeSwitcher.GetTransactionType(innerTx)
	if err != nil {
		return blocker.NewBlocker(
			blocker.ValidationErr,
			"FailToCastInnerTransaction",
		)
	}
	err = innerTa.ApplyUnconfirmed()
	if err != nil {
		return blocker.NewBlocker(
			blocker.ValidationErr,
			"FailToApplyUnconfirmedInnerTx",
		)
	}
	return nil
}

func (pth *PendingTransactionHelper) UndoApplyUnconfirmedPendingTransaction(pendingTransactionBytes []byte) error {
	// parse and apply unconfirmed
	innerTx, err := pth.TransactionUtil.ParseTransactionBytes(pendingTransactionBytes, false)
	if err != nil {
		return blocker.NewBlocker(
			blocker.ValidationErr,
			"FailToParseTransactionBytes",
		)
	}
	innerTa, err := pth.TypeSwitcher.GetTransactionType(innerTx)
	if err != nil {
		return blocker.NewBlocker(
			blocker.ValidationErr,
			"FailToCastInnerTransaction",
		)
	}
	err = innerTa.UndoApplyUnconfirmed()
	if err != nil {
		return blocker.NewBlocker(
			blocker.ValidationErr,
			"FailToUndoApplyUnconfirmedInnerTx",
		)
	}
	return nil
}

func (pth *PendingTransactionHelper) ApplyConfirmedPendingTransaction(
	pendingTransactionBytes []byte, txHeight uint32, blockTimestamp int64,
) (*model.Transaction, error) {
	utx, err := pth.TransactionUtil.ParseTransactionBytes(pendingTransactionBytes, false)
	if err != nil {
		return utx, err
	}
	utx.Height = txHeight
	utxAct, err := pth.TypeSwitcher.GetTransactionType(utx)
	if err != nil {
		return utx, err
	}
	err = utxAct.UndoApplyUnconfirmed()
	if err != nil {
		return utx, blocker.NewBlocker(
			blocker.ValidationErr,
			"FailToApplyUndoUnconfirmedInnerTx",
		)
	}
	// call ApplyConfirmed() to inner transaction
	err = utxAct.ApplyConfirmed(blockTimestamp)
	if err != nil {
		return utx, err
	}
	return utx, nil
}

func (sih *SignatureInfoHelper) InsertPendingSignature(
	pendingSignature *model.PendingSignature,
) error {
	insertPendingSigQ := sih.PendingSignatureQuery.InsertPendingSignature(pendingSignature)
	err := sih.QueryExecutor.ExecuteTransactions(insertPendingSigQ)
	if err != nil {
		return blocker.NewBlocker(blocker.DBErr, err.Error())
	}
	return nil
}
func (sih *SignatureInfoHelper) InsertPendingSignatures(
	pendingSignatures []*model.PendingSignature,
) (err error) {
	bulkQ, bulkArgs := sih.PendingSignatureQuery.InsertPendingSignatures(pendingSignatures)
	err = sih.QueryExecutor.ExecuteTransaction(bulkQ, bulkArgs...)
	if err != nil {
		return blocker.NewBlocker(blocker.DBErr, err.Error())
	}
	return nil
}

func (sih *SignatureInfoHelper) GetPendingSignatureByTransactionHash(transactionHash []byte, txHeight uint32) ([]*model.PendingSignature, error) {
	var pendingSigs []*model.PendingSignature
	q, args := sih.PendingSignatureQuery.GetPendingSignatureByHash(
		transactionHash,
		txHeight, constant.MinRollbackBlocks,
	)
	pendingSigRows, err := sih.QueryExecutor.ExecuteSelect(q, false, args...)
	if err != nil {
		return nil, err
	}
	defer pendingSigRows.Close()
	pendingSigs, err = sih.PendingSignatureQuery.BuildModel(pendingSigs, pendingSigRows)
	if err != nil {
		return nil, err
	}
	return pendingSigs, nil
}

func (msi *MultisignatureInfoHelper) GetMultisigInfoByAddress(
	multisigInfo *model.MultiSignatureInfo,
	multisigAddress []byte,
	blockHeight uint32,
) error {
	var (
		err              error
		multisigInfos    []*model.MultiSignatureInfo
		multisigAccounts [][]byte
	)
	q, args := msi.MultisignatureInfoQuery.GetMultisignatureInfoByAddressWithParticipants(
		multisigAddress, blockHeight, constant.MinRollbackBlocks,
	)
	rows, err := msi.QueryExecutor.ExecuteSelect(q, false, args...)
	if err != nil {
		return err
	}
	defer rows.Close()
	multisigInfos, err = msi.MultisignatureInfoQuery.BuildModelWithParticipant(multisigInfos, rows)
	if err != nil {
		return err
	}

	if len(multisigInfos) == 0 {
		return sql.ErrNoRows
	}
	// make sure we have all data from db when returning
	multisigInfo.MultisigAddress = multisigInfos[0].GetMultisigAddress()
	multisigInfo.Latest = multisigInfos[0].Latest
	multisigInfo.BlockHeight = multisigInfos[0].BlockHeight
	multisigInfo.MinimumSignatures = multisigInfos[0].MinimumSignatures
	multisigInfo.Nonce = multisigInfos[0].Nonce
	for _, msInfo := range multisigInfos {
		if len(msInfo.Addresses[0]) > 0 {
			multisigAccounts = append(multisigAccounts, msInfo.Addresses[0])
		}
	}
	multisigInfo.Addresses = multisigAccounts
	return nil
}

func (msi *MultisignatureInfoHelper) InsertMultisignatureInfo(multisigInfo *model.MultiSignatureInfo) error {
	var (
		queries              = msi.MultisignatureInfoQuery.InsertMultisignatureInfo(multisigInfo)
		participantAddresses []*model.MultiSignatureParticipant
	)
	for k, participant := range multisigInfo.GetAddresses() {
		participantAddresses = append(participantAddresses, &model.MultiSignatureParticipant{
			MultiSignatureAddress: multisigInfo.GetMultisigAddress(),
			AccountAddressIndex:   uint32(k),
			AccountAddress:        participant,
			BlockHeight:           multisigInfo.GetBlockHeight(),
			Latest:                true,
		})
	}
	participantQ := msi.MultiSignatureParticipantQuery.InsertMultisignatureParticipants(participantAddresses)
	queries = append(queries, participantQ...)
	return msi.QueryExecutor.ExecuteTransactions(queries)
}

func (msi *MultisignatureInfoHelper) InsertMultiSignaturesInfo(multiSignaturesInfo []*model.MultiSignatureInfo) (err error) {
	bulkQ := msi.MultisignatureInfoQuery.InsertMultiSignatureInfos(multiSignaturesInfo)
	err = msi.QueryExecutor.ExecuteTransactions(bulkQ)
	if err != nil {
		return blocker.NewBlocker(blocker.DBErr, err.Error())
	}
	return nil
}
func (tx *MultiSignatureTransaction) ApplyConfirmed(blockTimestamp int64) (err error) {

	var (
		address             []byte
		pendingTransaction  model.PendingTransaction
		multiSignaturesInfo []*model.MultiSignatureInfo
		pendingSignatures   []*model.PendingSignature
	)

	// if have multisig info, MultisigInfoService.AddMultisigInfo() -> noop duplicate
	if tx.Body.MultiSignatureInfo != nil {
		_, address, err = tx.TransactionUtil.GenerateMultiSigAddress(tx.Body.MultiSignatureInfo)
		if err != nil {
			return err
		}
		tx.Body.MultiSignatureInfo.MultisigAddress = address
		tx.Body.MultiSignatureInfo.BlockHeight = tx.Height
		tx.Body.MultiSignatureInfo.Latest = true
		err = tx.MultisignatureInfoHelper.InsertMultisignatureInfo(tx.Body.MultiSignatureInfo)
		if err != nil {
			return err
		}
	}
	// if have transaction bytes, PendingTransactionService.AddPendingTransaction() -> noop duplicate
	if len(tx.Body.UnsignedTransactionBytes) > 0 {
		var innerTx *model.Transaction
		innerTx, err = tx.TransactionUtil.ParseTransactionBytes(tx.Body.UnsignedTransactionBytes, false)
		if err != nil {
			return blocker.NewBlocker(
				blocker.ValidationErr,
				"FailToParseTransactionBytes",
			)
		}
		txHash := sha3.Sum256(tx.Body.UnsignedTransactionBytes)
		err = tx.PendingTransactionHelper.GetPendingTransactionByHash(&pendingTransaction, txHash[:], []model.PendingTransactionStatus{
			model.PendingTransactionStatus_PendingTransactionPending,
		}, tx.Height, true)
		if err == sql.ErrNoRows {
			// apply-unconfirmed on pending transaction
			err = tx.PendingTransactionHelper.ApplyUnconfirmedPendingTransaction(tx.Body.UnsignedTransactionBytes)
			if err != nil {
				return err
			}
			// save the pending transaction
			pendingTransaction = model.PendingTransaction{
				SenderAddress:    innerTx.SenderAccountAddress,
				TransactionHash:  txHash[:],
				TransactionBytes: tx.Body.UnsignedTransactionBytes,
				Status:           model.PendingTransactionStatus_PendingTransactionPending,
				BlockHeight:      tx.Height,
				Latest:           true,
			}
			err = tx.PendingTransactionHelper.InsertPendingTransaction(&pendingTransaction)
			if err != nil {
				return err
			}
		} else {
			return blocker.NewBlocker(
				blocker.ValidationErr,
				"PendingTransactionDuplicateOrAlreadyExecuted",
			)
		}
	}
	// if have signature, PendingSignature.AddPendingSignature -> noop duplicate
	if tx.Body.SignatureInfo != nil {
		var musigAddress []byte
		if tx.Body.GetMultiSignatureInfo() != nil {
			_, musigAddress, err = tx.TransactionUtil.GenerateMultiSigAddress(tx.Body.GetMultiSignatureInfo())
			if err != nil {
				return err
			}
		}
		for addrHex, sig := range tx.Body.SignatureInfo.Signatures {
			var (
				addr []byte
			)
			addr, err = hex.DecodeString(addrHex)
			if err != nil {
				return err
			}
			if util.ConvertBytesToUint32(addr[:constant.AccountAddressTypeLength]) == uint32(model.AccountType_MultiSignatureAccountType) {
				var (
					mis []*model.MultiSignatureInfo
					ps  []*model.PendingSignature
				)
				err = tx.MultisigUtil.ParseSignatureInfoBytesAsCandidates(
					tx.Body.GetSignatureInfo().GetTransactionHash(),
					musigAddress,
					addr,
					sig,
					tx.Height,
					&mis,
					&ps,
				)
				if err != nil {
					return err
				}
				err = tx.MultisignatureInfoHelper.InsertMultiSignaturesInfo(mis)
				if err != nil {
					return err
				}
				err = tx.SignatureInfoHelper.InsertPendingSignatures(ps)
				if err != nil {
					return err
				}
				multiSignaturesInfo = append(multiSignaturesInfo, mis...)
				pendingSignatures = append(pendingSignatures, ps...)
			} else {
				pendingSig := &model.PendingSignature{
					TransactionHash:       tx.Body.SignatureInfo.TransactionHash,
					AccountAddress:        addr,
					Signature:             sig,
					BlockHeight:           tx.Height,
					Latest:                true,
					MultiSignatureAddress: musigAddress,
				}

				err = tx.SignatureInfoHelper.InsertPendingSignature(pendingSig)
				if err != nil {
					return blocker.NewBlocker(blocker.DBErr, err.Error())
				}
				pendingSignatures = append(pendingSignatures, pendingSig)
			}
		}
	}
	// checks for completion, if multisigInfo && txBytes && signatureInfo exist, check if signature info complete
	err = tx.MultisigUtil.CheckMultisigComplete(
		tx.TransactionUtil,
		tx.MultisignatureInfoHelper,
		tx.SignatureInfoHelper,
		tx.PendingTransactionHelper,
		tx.Body,
		tx.Height,
		&multiSignaturesInfo,
		&pendingSignatures,
		&pendingTransaction,
	)
	if err != nil {
		if errType, ok := err.(blocker.Blocker); !ok || errType.Type != blocker.MultiSignatureNotComplete {
			return err
		}
	}
	// every element in txs will have all three optional field filled, to avoid infinite recursive calls.
	for _, pendingSignature := range pendingSignatures {
		// parse the UnsignedTransactionBytes
		var utx *model.Transaction
		utx, err = tx.PendingTransactionHelper.ApplyConfirmedPendingTransaction(
			pendingTransaction.GetTransactionBytes(),
			tx.Height,
			blockTimestamp,
		)
		if err != nil {
			return err
		}
		// update pending transaction status
		pendingTx := &model.PendingTransaction{
			SenderAddress:    pendingSignature.GetMultiSignatureAddress(),
			TransactionHash:  pendingSignature.GetTransactionHash(),
			TransactionBytes: pendingTransaction.GetTransactionBytes(),
			Status:           model.PendingTransactionStatus_PendingTransactionExecuted,
			BlockHeight:      tx.Height,
			Latest:           true,
		}
		// update pendingTx
		err = tx.PendingTransactionHelper.InsertPendingTransaction(pendingTx)
		if err != nil {
			return err
		}
		// save multisig_child transaction
		utx.MultisigChild = true
		utx.BlockID = tx.BlockID
		err = tx.TransactionHelper.InsertTransaction(utx)
		if err != nil {
			return err
		}
	}
	// deduct fee from sender
	err = tx.AccountBalanceHelper.AddAccountBalance(
		tx.SenderAddress,
		-tx.Fee,
		model.EventType_EventMultiSignatureTransaction,
		tx.Height,
		tx.ID,
		uint64(blockTimestamp),
	)
	if err != nil {
		return err
	}
	return nil
}

func (tx *MultiSignatureTransaction) ApplyUnconfirmed() error {
	var (
		err error
	)
	// reduce fee from sender
	err = tx.AccountBalanceHelper.AddAccountSpendableBalance(tx.SenderAddress, -tx.Fee)
	if err != nil {
		return err
	}
	// Run ApplyUnconfirmed of inner transaction
	if len(tx.Body.UnsignedTransactionBytes) > 0 {
		// parse and apply unconfirmed
		err = tx.PendingTransactionHelper.ApplyUnconfirmedPendingTransaction(tx.Body.UnsignedTransactionBytes)
		if err != nil {
			return err
		}
	}
	return nil
}

func (tx *MultiSignatureTransaction) UndoApplyUnconfirmed() error {
	// recover fee
	err := tx.AccountBalanceHelper.AddAccountSpendableBalance(tx.SenderAddress, tx.Fee)
	if err != nil {
		return err
	}
	if len(tx.Body.UnsignedTransactionBytes) > 0 {
		err = tx.PendingTransactionHelper.UndoApplyUnconfirmedPendingTransaction(tx.Body.UnsignedTransactionBytes)
		if err != nil {
			return err
		}
	}
	return nil
}

// Validate dbTx specify whether validation should read from transaction state or db state
func (tx *MultiSignatureTransaction) Validate(dbTx bool) error {
	var (
		multisigInfoAddresses = make(map[string]bool)
		enough                bool
		body                  = tx.Body
		err                   error
	)

	if body.MultiSignatureInfo == nil && body.SignatureInfo == nil && body.UnsignedTransactionBytes == nil {
		return blocker.NewBlocker(
			blocker.ValidationErr,
			"AtLeastTxBytesSignatureInfoOrMultisignatureInfoMustBeProvided",
		)
	}

	// check existing & balance account sender
	enough, err = tx.AccountBalanceHelper.HasEnoughSpendableBalance(dbTx, tx.SenderAddress, tx.Fee)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
		return blocker.NewBlocker(blocker.ValidationErr, "AccountAddressNotFound")
	}
	if !enough {
		return blocker.NewBlocker(blocker.ValidationErr, "AccountBalanceNotEnough")
	}

	if body.GetMultiSignatureInfo() != nil {
		err = tx.MultisigUtil.ValidateMultisignatureInfo(body.GetMultiSignatureInfo())
		if err != nil {
			return err
		}
		for _, address := range body.GetMultiSignatureInfo().GetAddresses() {
			multisigInfoAddresses[hex.EncodeToString(address)] = true
		}
	}
	if len(body.GetUnsignedTransactionBytes()) > 0 {
		// if ok will refill body.MultiSignatureInfo and validate signature false
		err = tx.MultisigUtil.ValidatePendingTransactionBytes(
			tx.TransactionUtil,
			tx.TypeSwitcher,
			tx.MultisignatureInfoHelper,
			tx.PendingTransactionHelper,
			body.GetMultiSignatureInfo(),
			tx.SenderAddress,
			body.GetUnsignedTransactionBytes(),
			tx.Height,
			dbTx,
		)
		if err != nil {
			return err
		}
		// multi signature info would not nil anymore
		for _, address := range body.GetMultiSignatureInfo().GetAddresses() {
			multisigInfoAddresses[hex.EncodeToString(address)] = true
		}
	}
	if body.GetSignatureInfo() != nil {

		var (
			pendingTX model.PendingTransaction
			txHash    [32]byte
		)
		if len(body.GetUnsignedTransactionBytes()) != 0 {
			txHash = sha3.Sum256(body.GetUnsignedTransactionBytes())
			err = tx.PendingTransactionHelper.GetPendingTransactionByHash(
				&pendingTX,
				txHash[:],
				[]model.PendingTransactionStatus{
					model.PendingTransactionStatus_PendingTransactionExecuted,
					model.PendingTransactionStatus_PendingTransactionPending,
				},
				tx.Height,
				dbTx,
			)
			if err != nil && err != sql.ErrNoRows {
				return err
			}
		}
		if body.GetMultiSignatureInfo() != nil {
			err = tx.MultisignatureInfoHelper.GetMultisigInfoByAddress(body.MultiSignatureInfo, pendingTX.GetSenderAddress(), tx.Height)
			if err != nil && err != sql.ErrNoRows {
				return err
			}
		}
		if body.GetSignatureInfo().GetTransactionHash() == nil { // transaction hash has to come with at least one signature
			return blocker.NewBlocker(
				blocker.ValidationErr,
				"TransactionHashRequiredInSignatureInfo",
			)
		}
		if len(body.GetSignatureInfo().GetSignatures()) < 1 {
			return blocker.NewBlocker(
				blocker.ValidationErr,
				"MinimumOneSignatureRequiredInSignatureInfo",
			)
		}

		if !bytes.Equal(txHash[:], body.GetSignatureInfo().GetTransactionHash()) {
			return blocker.NewBlocker(blocker.ValidationErr, "TransactionHashNotMatch")
		}

		for addressHex, sig := range body.GetSignatureInfo().GetSignatures() {
			var (
				decodedAcc          []byte
				multiSignaturesInfo []*model.MultiSignatureInfo
				pendingSignatures   []*model.PendingSignature
			)
			decodedAcc, err = hex.DecodeString(addressHex)
			if err != nil {
				return err
			}
			err = tx.Signature.VerifySignature(body.GetSignatureInfo().GetTransactionHash(), sig, decodedAcc)
			if err != nil {
				if util.ConvertBytesToUint32(decodedAcc[:constant.AccountAddressTypeLength]) == uint32(model.AccountType_MultiSignatureAccountType) {
					var musigAddress []byte
					if tx.Body.GetMultiSignatureInfo() != nil {
						_, musigAddress, err = tx.TransactionUtil.GenerateMultiSigAddress(tx.Body.GetMultiSignatureInfo())
						if err != nil {
							return err
						}
					}
					err = tx.MultisigUtil.ParseSignatureInfoBytesAsCandidates(
						body.GetSignatureInfo().GetTransactionHash(),
						musigAddress,
						decodedAcc,
						sig,
						tx.Height,
						&multiSignaturesInfo,
						&pendingSignatures,
					)
					if err != nil {
						return err
					}

					for k, innerMultiSignatureInfo := range multiSignaturesInfo {
						var innerMultiSignatureAddress []byte
						err = tx.MultisigUtil.ValidateMultisignatureInfo(innerMultiSignatureInfo)
						if err != nil {
							return err
						}
						// Get Participant address, index should the last
						if k == len(multiSignaturesInfo)-1 {
							_, innerMultiSignatureAddress, err = tx.TransactionUtil.GenerateMultiSigAddress(innerMultiSignatureInfo)
							if err != nil {
								return err
							}
							if _, ok := multisigInfoAddresses[hex.EncodeToString(innerMultiSignatureAddress)]; !ok {
								return blocker.NewBlocker(
									blocker.ValidationErr,
									"SignerNotInParticipantList",
								)
							}
						}
					}
					for _, pendingSignature := range pendingSignatures {
						err = tx.Signature.VerifySignature(
							body.GetSignatureInfo().GetTransactionHash(),
							pendingSignature.GetSignature(),
							pendingSignature.GetAccountAddress(),
						)
						if err != nil {
							return err
						}
					}
				} else {
					return err
				}
			} else {
				if _, ok := multisigInfoAddresses[addressHex]; !ok {
					return blocker.NewBlocker(
						blocker.ValidationErr,
						"SignerNotInParticipantList",
					)
				}
			}
		}
	}
	return nil
}

func (tx *MultiSignatureTransaction) GetMinimumFee() (int64, error) {
	if tx.Escrow != nil && tx.Escrow.GetApproverAddress() != nil && !bytes.Equal(tx.Escrow.GetApproverAddress(), []byte{}) {
		return tx.EscrowFee.CalculateTxMinimumFee(tx.Body, tx.Escrow)
	}
	return tx.NormalFee.CalculateTxMinimumFee(tx.Body, tx.Escrow)
}

func (*MultiSignatureTransaction) GetAmount() int64 {
	return 0
}

func (tx *MultiSignatureTransaction) GetSize() (uint32, error) {
	var (
		txByteSize, signaturesSize, multisigInfoSize uint32
	)
	// MultisigInfo
	multisigInfo := tx.Body.GetMultiSignatureInfo()
	multisigInfoSize += constant.MultisigFieldLength
	if multisigInfo != nil {
		multisigInfoSize += constant.MultiSigInfoMinSignature
		multisigInfoSize += constant.MultiSigInfoNonce
		multisigInfoSize += constant.MultiSigNumberOfAddress
		for _, v := range multisigInfo.GetAddresses() {
			multisigInfoSize += constant.MultiSigAddressLength
			multisigInfoSize += uint32(len(v))
		}
	}
	// TransactionBytes
	txByteSize = constant.MultiSigUnsignedTxBytesLength + uint32(len(tx.Body.GetUnsignedTransactionBytes()))
	// SignatureInfo
	signaturesSize += constant.MultisigFieldLength
	if tx.Body.GetSignatureInfo() != nil {
		signaturesSize += constant.MultiSigTransactionHash
		signaturesSize += constant.MultiSigNumberOfSignatures
		for address, sig := range tx.Body.SignatureInfo.Signatures {
			signaturesSize += constant.MultiSigSignatureAddressLength
			signaturesSize += uint32(len([]byte(address)))
			signaturesSize += constant.MultiSigSignatureLength
			signaturesSize += uint32(len(sig))
		}
	}

	return txByteSize + signaturesSize + multisigInfoSize, nil
}

func (tx *MultiSignatureTransaction) ParseBodyBytes(txBodyBytes []byte) (model.TransactionBodyInterface, error) {
	var (
		addresses     [][]byte
		signatures    = make(map[string][]byte)
		multisigInfo  *model.MultiSignatureInfo
		signatureInfo *model.SignatureInfo
		err           error
	)
	bufferBytes := bytes.NewBuffer(txBodyBytes)
	// MultisigInfo
	multisigInfoPresent := util.ConvertBytesToUint32(bufferBytes.Next(int(constant.MultisigFieldLength)))
	if multisigInfoPresent == constant.MultiSigFieldPresent {
		minSignatures := util.ConvertBytesToUint32(bufferBytes.Next(int(constant.MultiSigInfoMinSignature)))
		nonce := util.ConvertBytesToUint64(bufferBytes.Next(int(constant.MultiSigInfoNonce)))
		addressesLength := util.ConvertBytesToUint32(bufferBytes.Next(int(constant.MultiSigNumberOfAddress)))
		for i := 0; i < int(addressesLength); i++ {
			var (
				accType accounttype.AccountTypeInterface
				address []byte
			)

			accType, err = accounttype.ParseBytesToAccountType(bufferBytes)
			if err != nil {
				return nil, err
			}
			address, err = accType.GetAccountAddress()
			if err != nil {
				return nil, err
			}
			addresses = append(addresses, address)
		}
		multisigInfo = &model.MultiSignatureInfo{
			MinimumSignatures: minSignatures,
			Nonce:             int64(nonce),
			Addresses:         addresses,
		}
	}
	// TransactionBytes
	unsignedTxLength := util.ConvertBytesToUint32(bufferBytes.Next(int(constant.MultiSigUnsignedTxBytesLength)))
	unsignedTx, err := util.ReadTransactionBytes(bufferBytes, int(unsignedTxLength))
	if err != nil {
		return nil, err
	}
	// SignatureInfo
	signatureInfoPresent := util.ConvertBytesToUint32(bufferBytes.Next(int(constant.MultisigFieldLength)))
	if signatureInfoPresent == constant.MultiSigFieldPresent {
		var transactionHash []byte
		transactionHash, err = util.ReadTransactionBytes(bufferBytes, int(constant.MultiSigTransactionHash))
		if err != nil {
			return nil, err
		}
		numberOfSignatures := util.ConvertBytesToUint32(bufferBytes.Next(int(constant.MultiSigNumberOfSignatures)))
		for i := 0; i < int(numberOfSignatures); i++ {
			var (
				accTypeBytes, address, signature, multiSignatureInfoBytes []byte
				accType                                                   accounttype.AccountTypeInterface
			)

			accTypeBytes = bufferBytes.Next(int(constant.AccountAddressTypeLength))
			if model.AccountType(util.ConvertBytesToUint32(accTypeBytes)) == model.AccountType_MultiSignatureAccountType {
				lenUint := util.ConvertBytesToUint32(bufferBytes.Next(int(constant.MultiSigInfoSize)))
				multiSignatureInfoBytes = bufferBytes.Next(int(lenUint))
				if len(multiSignatureInfoBytes) == 0 {
					return nil, fmt.Errorf("invalid signature info format")
				}
				rebuildAddress := bytes.NewBuffer(accTypeBytes)
				rebuildAddress.Write(util.ConvertUint32ToBytes(lenUint))
				rebuildAddress.Write(multiSignatureInfoBytes)
				address = rebuildAddress.Bytes()
			} else {
				accType, err = accounttype.NewAccountType(int32(util.ConvertBytesToUint32(accTypeBytes)), []byte{})
				if err != nil {
					return nil, err
				}
				address = append(accTypeBytes, bufferBytes.Next(int(accType.GetAccountPublicKeyLength()))...)
			}

			signatureLength := util.ConvertBytesToUint32(bufferBytes.Next(int(constant.MultiSigSignatureLength)))
			signature, err = util.ReadTransactionBytes(bufferBytes, int(signatureLength))
			if err != nil {
				return nil, err
			}
			// encode address to hex string to be able to build the map (signature order must be preserved, so we can't use slices)
			signatures[hex.EncodeToString(address)] = signature
		}
		signatureInfo = &model.SignatureInfo{
			TransactionHash: transactionHash,
			Signatures:      signatures,
		}
	}

	return &model.MultiSignatureTransactionBody{
		MultiSignatureInfo:       multisigInfo,
		UnsignedTransactionBytes: unsignedTx,
		SignatureInfo:            signatureInfo,
	}, nil
}

func (tx *MultiSignatureTransaction) GetBodyBytes() ([]byte, error) {
	var (
		buffer = bytes.NewBuffer([]byte{})
	)
	// Multisig Info
	if tx.Body.GetMultiSignatureInfo() != nil {
		buffer.Write(util.ConvertUint32ToBytes(constant.MultiSigFieldPresent))
		buffer.Write(util.ConvertUint32ToBytes(tx.Body.GetMultiSignatureInfo().GetMinimumSignatures()))
		buffer.Write(util.ConvertUint64ToBytes(uint64(tx.Body.GetMultiSignatureInfo().GetNonce())))
		buffer.Write(util.ConvertUint32ToBytes(uint32(len(tx.Body.GetMultiSignatureInfo().GetAddresses()))))
		for _, v := range tx.Body.GetMultiSignatureInfo().GetAddresses() {
			buffer.Write(v)
		}
	} else {
		buffer.Write(util.ConvertUint32ToBytes(constant.MultiSigFieldMissing))
	}
	// Transaction Bytes
	buffer.Write(util.ConvertUint32ToBytes(uint32(len(tx.Body.GetUnsignedTransactionBytes()))))
	buffer.Write(tx.Body.GetUnsignedTransactionBytes())
	// SignatureInfo
	if tx.Body.GetSignatureInfo() != nil {
		buffer.Write(util.ConvertUint32ToBytes(constant.MultiSigFieldPresent))
		buffer.Write(tx.Body.GetSignatureInfo().GetTransactionHash())
		buffer.Write(util.ConvertUint32ToBytes(uint32(len(tx.Body.GetSignatureInfo().GetSignatures()))))
		for addressHex, sig := range tx.Body.GetSignatureInfo().GetSignatures() {
			var (
				accountAddress []byte
				err            error
			)

			accountAddress, err = hex.DecodeString(addressHex)
			if err != nil {
				return nil, err
			}

			if _, ok := model.AccountType_name[int32(util.ConvertBytesToUint32(accountAddress[:constant.AccountAddressTypeLength]))]; !ok {
				return []byte{}, fmt.Errorf("invalid participant format")
			}
			buffer.Write(accountAddress)
			buffer.Write(util.ConvertUint32ToBytes(uint32(len(sig))))
			buffer.Write(sig)
		}
	} else {
		buffer.Write(util.ConvertUint32ToBytes(constant.MultiSigFieldMissing))
	}

	return buffer.Bytes(), nil
}

func (tx *MultiSignatureTransaction) GetTransactionBody(transaction *model.Transaction) {
	transaction.TransactionBody = &model.Transaction_MultiSignatureTransactionBody{
		MultiSignatureTransactionBody: tx.Body,
	}
}

func (*MultiSignatureTransaction) SkipMempoolTransaction([]*model.Transaction, int64, uint32) (bool, error) {
	return false, nil
}

func (tx *MultiSignatureTransaction) Escrowable() (EscrowTypeAction, bool) {
	if tx.Escrow.GetApproverAddress() != nil && !bytes.Equal(tx.Escrow.GetApproverAddress(), []byte{}) {
		tx.Escrow = &model.Escrow{
			ID:              tx.ID,
			SenderAddress:   tx.SenderAddress,
			ApproverAddress: tx.Escrow.GetApproverAddress(),
			Commission:      tx.Escrow.GetCommission(),
			Timeout:         tx.Escrow.GetTimeout(),
			Status:          tx.Escrow.GetStatus(),
			BlockHeight:     tx.Height,
			Latest:          true,
			Instruction:     tx.Escrow.GetInstruction(),
		}
		return EscrowTypeAction(tx), true
	}
	return nil, false
}

func (tx *MultiSignatureTransaction) EscrowApplyConfirmed(blockTimestamp int64) error {
	return tx.AccountBalanceHelper.AddAccountBalance(
		tx.SenderAddress,
		-(tx.Fee + tx.Escrow.GetCommission()),
		model.EventType_EventEscrowedTransaction,
		tx.Height,
		tx.ID,
		uint64(blockTimestamp),
	)
}

func (tx *MultiSignatureTransaction) EscrowApplyUnconfirmed() error {
	return tx.AccountBalanceHelper.AddAccountSpendableBalance(tx.SenderAddress, -tx.Fee)
}

func (tx *MultiSignatureTransaction) EscrowUndoApplyUnconfirmed() error {
	return tx.AccountBalanceHelper.AddAccountSpendableBalance(tx.SenderAddress, tx.Fee)
}

func (tx *MultiSignatureTransaction) EscrowValidate(dbTx bool) error {

	var (
		err    error
		enough bool
	)
	if tx.Escrow.GetApproverAddress() == nil || bytes.Equal(tx.Escrow.GetApproverAddress(), []byte{}) {
		return blocker.NewBlocker(blocker.RequestParameterErr, "ApproverAddressRequired")
	}
	if tx.Escrow.GetCommission() <= 0 {
		return blocker.NewBlocker(blocker.RequestParameterErr, "CommissionRequired")
	}
	if tx.Escrow.GetTimeout() > uint64(constant.MinRollbackBlocks) {
		return blocker.NewBlocker(blocker.ValidationErr, "TimeoutRequired")
	}
	err = tx.Validate(dbTx)
	if err != nil {
		return err
	}

	enough, err = tx.AccountBalanceHelper.HasEnoughSpendableBalance(dbTx, tx.SenderAddress, tx.Fee+tx.Escrow.GetCommission())
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
		return blocker.NewBlocker(blocker.ValidationErr, "AccountBalanceNotFound")
	}
	if !enough {
		return blocker.NewBlocker(blocker.ValidationErr, "AccountBalanceNotEnough")
	}
	return nil
}

func (tx *MultiSignatureTransaction) EscrowApproval(blockTimestamp int64, txBody *model.ApprovalEscrowTransactionBody) (err error) {
	switch txBody.GetApproval() {
	case model.EscrowApproval_Approve:
		tx.Escrow.Status = model.EscrowStatus_Approved
		err = tx.AccountBalanceHelper.AddAccountBalance(
			tx.SenderAddress,
			tx.Fee,
			model.EventType_EventEscrowedTransaction,
			tx.Height,
			tx.ID,
			uint64(blockTimestamp),
		)
		if err != nil {
			return err
		}
		err = tx.ApplyConfirmed(blockTimestamp)
		if err != nil {
			return err
		}
		err = tx.AccountBalanceHelper.AddAccountBalance(
			tx.Escrow.GetApproverAddress(),
			tx.Escrow.GetCommission(),
			model.EventType_EventApprovalEscrowTransaction,
			tx.Height,
			tx.ID,
			uint64(blockTimestamp),
		)
		if err != nil {
			return err
		}
	case model.EscrowApproval_Reject:
		tx.Escrow.Status = model.EscrowStatus_Rejected
		err = tx.AccountBalanceHelper.AddAccountBalance(
			tx.Escrow.GetApproverAddress(),
			tx.Escrow.GetCommission(),
			model.EventType_EventApprovalEscrowTransaction,
			tx.Height,
			tx.ID,
			uint64(blockTimestamp),
		)
		if err != nil {
			return err
		}
	default:
		err = tx.AccountBalanceHelper.AddAccountBalance(
			tx.SenderAddress,
			tx.Escrow.GetCommission(),
			model.EventType_EventApprovalEscrowTransaction,
			tx.Height,
			tx.ID,
			uint64(blockTimestamp),
		)
		if err != nil {
			return err
		}
	}
	escrowQ := tx.EscrowQuery.InsertEscrowTransaction(tx.Escrow)
	err = tx.QueryExecutor.ExecuteTransactions(escrowQ)
	if err != nil {
		return err
	}
	return nil
}
