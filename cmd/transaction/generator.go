package transaction

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/zoobc/zoobc-core/common/accounttype"
	"github.com/zoobc/zoobc-core/common/signaturetype"
	"log"
	"strings"
	"time"

	"github.com/zoobc/zoobc-core/cmd/admin"
	"github.com/zoobc/zoobc-core/common/constant"
	"github.com/zoobc/zoobc-core/common/model"
	"github.com/zoobc/zoobc-core/common/query"
	rpcService "github.com/zoobc/zoobc-core/common/service"
	"github.com/zoobc/zoobc-core/common/transaction"
	"github.com/zoobc/zoobc-core/common/util"
	"golang.org/x/crypto/sha3"
	"google.golang.org/grpc"
)

// GenerateTxSendMoney return send money transaction based on provided basic transaction & ammunt
func GenerateTxSendMoney(tx *model.Transaction, sendAmount int64) *model.Transaction {
	tx.TransactionType = util.ConvertBytesToUint32(txTypeMap["sendMoney"])
	tx.TransactionBody = &model.Transaction_SendMoneyTransactionBody{
		SendMoneyTransactionBody: &model.SendMoneyTransactionBody{
			Amount: sendAmount,
		},
	}
	tx.TransactionBodyBytes = util.ConvertUint64ToBytes(uint64(sendAmount))
	tx.TransactionBodyLength = 8
	return tx
}

/*
GenerateTxRegisterNode return register node transaction based on provided basic transaction &
others specific field for generate register node transaction
*/
func GenerateTxRegisterNode(
	tx *model.Transaction,
	lockedBalance int64,
	nodePubKey []byte,
	proofOfOwnerShip *model.ProofOfOwnership,
) *model.Transaction {

	txBody := &model.NodeRegistrationTransactionBody{
		AccountAddress: tx.SenderAccountAddress,
		NodePublicKey:  nodePubKey,
		LockedBalance:  lockedBalance,
		Poown:          proofOfOwnerShip,
	}
	txBodyBytes, _ := (&transaction.NodeRegistration{
		Body:                  txBody,
		NodeRegistrationQuery: query.NewNodeRegistrationQuery(),
	}).GetBodyBytes()

	tx.TransactionType = util.ConvertBytesToUint32(txTypeMap["registerNode"])
	tx.TransactionBody = &model.Transaction_NodeRegistrationTransactionBody{
		NodeRegistrationTransactionBody: txBody,
	}
	tx.TransactionBodyBytes = txBodyBytes
	tx.TransactionBodyLength = uint32(len(txBodyBytes))

	return tx
}

/*
GenerateTxUpdateNode return update node transaction based on provided basic transaction &
others specific field for update register node transaction
*/
func GenerateTxUpdateNode(
	tx *model.Transaction,
	lockedBalance int64,
	nodePubKey []byte,
	proofOfOwnerShip *model.ProofOfOwnership,
) *model.Transaction {
	txBody := &model.UpdateNodeRegistrationTransactionBody{
		NodePublicKey: nodePubKey,
		LockedBalance: lockedBalance,
		Poown:         proofOfOwnerShip,
	}
	txBodyBytes, _ := (&transaction.UpdateNodeRegistration{
		Body:                  txBody,
		NodeRegistrationQuery: query.NewNodeRegistrationQuery(),
	}).GetBodyBytes()

	tx.TransactionBodyLength = uint32(len(txBodyBytes))
	tx.TransactionType = util.ConvertBytesToUint32(txTypeMap["updateNodeRegistration"])
	tx.TransactionBody = &model.Transaction_UpdateNodeRegistrationTransactionBody{
		UpdateNodeRegistrationTransactionBody: txBody,
	}
	tx.TransactionBodyBytes = txBodyBytes
	return tx
}

/*
GenerateTxRemoveNode return remove node transaction based on provided basic transaction &
others specific field for remove node transaction
*/
func GenerateTxRemoveNode(tx *model.Transaction, nodePubKey []byte) *model.Transaction {
	txBody := &model.RemoveNodeRegistrationTransactionBody{
		NodePublicKey: nodePubKey,
	}
	txBodyBytes, _ := (&transaction.RemoveNodeRegistration{
		Body:                  txBody,
		NodeRegistrationQuery: query.NewNodeRegistrationQuery(),
	}).GetBodyBytes()

	tx.TransactionType = util.ConvertBytesToUint32(txTypeMap["removeNodeRegistration"])
	tx.TransactionBody = &model.Transaction_RemoveNodeRegistrationTransactionBody{
		RemoveNodeRegistrationTransactionBody: txBody,
	}
	tx.TransactionBodyBytes = txBodyBytes
	tx.TransactionBodyLength = uint32(len(txBodyBytes))
	return tx
}

/*
GenerateTxClaimNode return claim node transaction based on provided basic transaction &
others specific field for claim node transaction
*/
func GenerateTxClaimNode(
	tx *model.Transaction,
	nodePubKey []byte,
	proofOfOwnerShip *model.ProofOfOwnership,
) *model.Transaction {
	var (
		txBody = &model.ClaimNodeRegistrationTransactionBody{
			NodePublicKey: nodePubKey,
			Poown:         proofOfOwnerShip,
		}
		txBodyBytes, _ = (&transaction.ClaimNodeRegistration{
			Body: txBody,
		}).GetBodyBytes()
	)
	tx.TransactionType = util.ConvertBytesToUint32(txTypeMap["claimNodeRegistration"])
	tx.TransactionBody = &model.Transaction_ClaimNodeRegistrationTransactionBody{
		ClaimNodeRegistrationTransactionBody: txBody,
	}
	tx.TransactionBodyBytes = txBodyBytes
	tx.TransactionBodyLength = uint32(len(txBodyBytes))
	return tx
}

// GenerateProofOfOwnership generate proof of owner ship for transaction related with node registry
func GenerateProofOfOwnership(
	dbPath, dbname, nodeOwnerAccountAddress, nodeSeed, proofOfOwnershipHex string,
) *model.ProofOfOwnership {
	if proofOfOwnershipHex != "" {
		powBytes, err := hex.DecodeString(proofOfOwnershipHex)
		if err != nil {
			panic(fmt.Sprintln("failed decode proofOfOwnershipHex, ", err.Error()))
		}
		pow, err := util.ParseProofOfOwnershipBytes(powBytes)
		if err != nil {
			panic(fmt.Sprintln("failed parse proofOfOwnership, ", err.Error()))
		}
		return pow
	}
	return admin.GetProofOfOwnerShip(dbPath, dbname, nodeOwnerAccountAddress, nodeSeed)
}

/*
GenerateTxSetupAccountDataset return setup account dataset transaction based on provided basic transaction &
others specific field for setup account dataset transaction
*/
func GenerateTxSetupAccountDataset(
	tx *model.Transaction,
	property, value string,
) *model.Transaction {
	txBody := &model.SetupAccountDatasetTransactionBody{
		Property: property,
		Value:    value,
	}
	txBodyBytes, _ := (&transaction.SetupAccountDataset{
		Body: txBody,
	}).GetBodyBytes()

	tx.TransactionType = util.ConvertBytesToUint32(txTypeMap["setupAccountDataset"])
	tx.TransactionBody = &model.Transaction_SetupAccountDatasetTransactionBody{
		SetupAccountDatasetTransactionBody: txBody,
	}
	tx.TransactionBodyBytes = txBodyBytes
	tx.TransactionBodyLength = uint32(len(txBodyBytes))
	return tx
}

/*
GenerateTxRemoveAccountDataset return remove account dataset transaction based on provided basic transaction &
others specific field for remove account dataset transaction
*/
func GenerateTxRemoveAccountDataset(
	tx *model.Transaction,
	property, value string,
) *model.Transaction {
	txBody := &model.RemoveAccountDatasetTransactionBody{
		Property: property,
		Value:    value,
	}
	txBodyBytes, _ := (&transaction.RemoveAccountDataset{
		Body: txBody,
	}).GetBodyBytes()

	tx.TransactionType = util.ConvertBytesToUint32(txTypeMap["removeAccountDataset"])
	tx.TransactionBody = &model.Transaction_RemoveAccountDatasetTransactionBody{
		RemoveAccountDatasetTransactionBody: txBody,
	}
	tx.TransactionBodyBytes = txBodyBytes
	tx.TransactionBodyLength = uint32(len(txBodyBytes))
	return tx
}

func getAccountTypeFromAccountHex(senderAccountAddressHex string) accounttype.AccountTypeInterface {
	accountAddress, err := hex.DecodeString(senderAccountAddressHex)
	if err != nil {
		panic(fmt.Sprintln(
			"GenerateBasicTransaction-Failed DecodeHexAddress",
			err.Error(),
		))
	}
	accountType, err := accounttype.NewAccountTypeFromAccount(accountAddress)
	if err != nil {
		panic(fmt.Sprintln(
			"GenerateBasicTransaction-Failed DecodeAccountTypeFromAddress",
			err.Error(),
		))
	}
	return accountType
}

// GenerateBasicTransaction return  basic transaction based on common transaction field
func GenerateBasicTransaction(
	senderAccountAddressHex, senderSeed string,
	version uint32,
	timestamp, fee int64,
	recipientAccountAddressHex,
	message string,
) *model.Transaction {
	if senderAccountAddressHex == "" && senderSeed != "" {
		accountType := getAccountTypeFromAccountHex(senderAccountAddressHex)
		// TODO: move this into AccountType interface
		switch accountType.GetSignatureType() {
		case model.SignatureType_DefaultSignature:
			b, err := signaturetype.NewEd25519Signature().GetPrivateKeyFromSeedUseSlip10(senderSeed)
			if err != nil {
				panic(err.Error())
			}
			bb, err := signaturetype.NewEd25519Signature().GetPublicKeyFromPrivateKeyUseSlip10(b)
			if err != nil {
				panic(err.Error())
			}
			senderAccountAddressHex, err = signaturetype.NewEd25519Signature().GetAddressFromPublicKey(constant.PrefixZoobcDefaultAccount, bb)
			if err != nil {
				panic(err.Error())
			}
		case model.SignatureType_BitcoinSignature:
			var (
				bitcoinSig  = signaturetype.NewBitcoinSignature(signaturetype.DefaultBitcoinNetworkParams(), signaturetype.DefaultBitcoinCurve())
				pubKey, err = bitcoinSig.GetPublicKeyFromSeed(
					senderSeed,
					signaturetype.DefaultBitcoinPublicKeyFormat(),
					signaturetype.DefaultBitcoinPrivateKeyLength(),
				)
			)
			if err != nil {
				panic(fmt.Sprintln(
					"GenerateBasicTransaction-BitcoinSignature-Failed GetPublicKey",
					err.Error(),
				))
			}
			senderAccountAddressHex, err = bitcoinSig.GetAddressFromPublicKey(pubKey)
			if err != nil {
				panic(fmt.Sprintln(
					"GenerateBasicTransaction-BitcoinSignature-Failed GetPublicKey",
					err.Error(),
				))
			}
		default:
			panic("GenerateBasicTransaction-Invalid Signature Type")
		}
	}

	if timestamp <= 0 {
		timestamp = time.Now().Unix()
	}
	decodedSenderAddress, err := hex.DecodeString(senderAccountAddressHex)
	if err != nil {
		panic(err)
	}
	decodedRecipientAddress, err := hex.DecodeString(recipientAccountAddressHex)
	if err != nil {
		panic(err)
	}
	return &model.Transaction{
		Version:                 version,
		Timestamp:               timestamp,
		SenderAccountAddress:    decodedSenderAddress,
		RecipientAccountAddress: decodedRecipientAddress,
		Fee:                     fee,
		Escrow: &model.Escrow{
			ApproverAddress: nil,
			Commission:      0,
			Timeout:         0,
		},
		Message: message,
	}
}

// PrintTx will print out the signed transaction based on provided format
func PrintTx(signedTxBytes []byte, outputType string) {
	var resultStr string
	switch outputType {
	case "hex":
		resultStr = hex.EncodeToString(signedTxBytes)
	default:
		var byteStrArr []string
		for _, bt := range signedTxBytes {
			byteStrArr = append(byteStrArr, fmt.Sprintf("%v", bt))
		}
		resultStr = strings.Join(byteStrArr, ", ")
	}
	if post {
		conn, err := grpc.Dial(postHost, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		c := rpcService.NewTransactionServiceClient(conn)

		response, err := c.PostTransaction(context.Background(), &model.PostTransactionRequest{
			TransactionBytes: signedTxBytes,
		})
		if err != nil {
			fmt.Printf("post failed: %v\n", err)
		} else {
			fmt.Printf("\n\nresult: %v\n", response)
		}
	} else {
		fmt.Println("")
		fmt.Printf("Length: %d\n", len(signedTxBytes))
		fmt.Printf("bytes: %s\n", resultStr)
	}
}

// GenerateSignedTxBytes return signed transaction bytes
func GenerateSignedTxBytes(
	tx *model.Transaction,
	senderSeed string,
	accountTypeInt int32,
	optionalSignParams ...interface{},
) []byte {
	var (
		transactionUtil = &transaction.Util{}
		txType          transaction.TypeAction
		err             error
	)
	txType, err = (&transaction.TypeSwitcher{}).GetTransactionType(tx)
	if err != nil {
		log.Fatalf("fail get transaction type: %s", err)
	}
	minimumFee, err := txType.GetMinimumFee()
	if err != nil {
		log.Fatalf("fail get minimum fee: %s", err)
	}
	tx.Fee += minimumFee

	unsignedTxBytes, _ := transactionUtil.GetTransactionBytes(tx, false)
	if senderSeed == "" {
		return unsignedTxBytes
	}
	txBytesHash := sha3.Sum256(unsignedTxBytes)
	tx.Signature, err = signature.Sign(
		txBytesHash[:],
		model.AccountType(accountTypeInt),
		senderSeed,
		optionalSignParams...,
	)
	if err != nil {
		log.Fatalf("fail get sign tx: %s", err)
	}
	signedTxBytes, err := transactionUtil.GetTransactionBytes(tx, true)
	if err != nil {
		log.Fatalf("fail get get signed transactionBytes: %s", err)
	}
	return signedTxBytes
}

// GenerateEscrowApprovalTransaction set escrow approval body
func GenerateEscrowApprovalTransaction(tx *model.Transaction) *model.Transaction {

	var chosen model.EscrowApproval
	switch approval {
	case true:
		chosen = model.EscrowApproval_Approve
	default:
		chosen = model.EscrowApproval_Reject
	}

	txBody := &model.ApprovalEscrowTransactionBody{
		Approval:      chosen,
		TransactionID: transactionID,
	}
	txBodyBytes, _ := (&transaction.ApprovalEscrowTransaction{
		Body: txBody,
	}).GetBodyBytes()

	tx.TransactionBody = txBody
	tx.TransactionBodyBytes = txBodyBytes
	tx.TransactionBodyLength = constant.EscrowApprovalBytesLength
	tx.TransactionType = util.ConvertBytesToUint32(txTypeMap["approvalEscrow"])

	return tx
}

/*
GenerateEscrowedTransaction inject escrow. Need:
		1. esApproverAddressHex
		2. Commission
		3. Timeout
Invalid escrow validation when those fields has not set
*/
func GenerateEscrowedTransaction(
	tx *model.Transaction,
) *model.Transaction {
	decodedApproverAddress, err := hex.DecodeString(esApproverAddressHex)
	if err != nil {
		panic(err)
	}
	tx.Escrow = &model.Escrow{
		ApproverAddress: decodedApproverAddress,
		Commission:      esCommission,
		Timeout:         esTimeout,
		Instruction:     esInstruction,
	}
	return tx
}

/*
GeneratedMultiSignatureTransaction inject escrow. Need:
		1. unsignedTxHex
		2. signatures
		3. multisigInfo:
			- minSignature
			- nonce
			- addressesHex
Invalid escrow validation when those fields has not set
*/
func GeneratedMultiSignatureTransaction(
	tx *model.Transaction,
	minSignature uint32,
	nonce int64,
	unsignedTxHex, txHash string,
	addressSignatures map[string]string, addressesHex []string,
) *model.Transaction {
	var (
		signatures    = make(map[string][]byte)
		signatureInfo *model.SignatureInfo
		unsignedTx    []byte
		multiSigInfo  *model.MultiSignatureInfo
		err           error
		fullAddresses [][]byte
	)
	for _, addrHex := range addressesHex {
		decodedAddr, err := hex.DecodeString(addrHex)
		if err != nil {
			panic(err)
		}
		fullAddresses = append(fullAddresses, decodedAddr)
	}
	if minSignature > 0 && len(addressesHex) > 0 {
		multiSigInfo = &model.MultiSignatureInfo{
			MinimumSignatures: minSignature,
			Nonce:             nonce,
			Addresses:         fullAddresses,
		}
	}
	if unsignedTxHex != "" {
		unsignedTx, err = hex.DecodeString(unsignedTxHex)
		if err != nil {
			return nil
		}
	}
	if txHash != "" {
		transactionHash, err := hex.DecodeString(txHash)
		if err != nil {
			return nil
		}
		for k, v := range addressSignatures {
			if v == "" {
				sigType := util.ConvertUint32ToBytes(2)
				signatures[k] = sigType
			} else {
				signature, err := hex.DecodeString(v)
				if err != nil {
					return nil
				}
				signatures[k] = signature
			}
		}
		fmt.Printf("signatures: %v\n\n\n", signatures)
		signatureInfo = &model.SignatureInfo{
			TransactionHash: transactionHash,
			Signatures:      signatures,
		}
	}
	tx.TransactionType = util.ConvertBytesToUint32(txTypeMap["multiSignature"])
	txBody := &model.MultiSignatureTransactionBody{
		MultiSignatureInfo:       multiSigInfo,
		UnsignedTransactionBytes: unsignedTx,
		SignatureInfo:            signatureInfo,
	}
	if tx.TransactionBodyBytes, err = (&transaction.MultiSignatureTransaction{
		Body: txBody,
	}).GetBodyBytes(); err != nil {
		panic(err)
	}
	fmt.Printf("length: %v\n", len(tx.TransactionBodyBytes))
	tx.TransactionBodyLength = uint32(len(tx.TransactionBodyBytes))
	return tx
}

func GenerateTxRemoveNodeHDwallet(tx *model.Transaction, nodePubKey []byte) *model.Transaction {
	txBody := &model.RemoveNodeRegistrationTransactionBody{
		NodePublicKey: nodePubKey,
	}
	txBodyBytes, _ := (&transaction.RemoveNodeRegistration{
		Body:                  txBody,
		NodeRegistrationQuery: query.NewNodeRegistrationQuery(),
	}).GetBodyBytes()

	tx.TransactionType = util.ConvertBytesToUint32(txTypeMap["removeNodeRegistration"])
	tx.TransactionBody = &model.Transaction_RemoveNodeRegistrationTransactionBody{
		RemoveNodeRegistrationTransactionBody: txBody,
	}
	tx.TransactionBodyBytes = txBodyBytes
	tx.TransactionBodyLength = uint32(len(txBodyBytes))
	return tx
}

/*
GenerateTxFeeVoteCommitment return fee vote commit vote transaction based on provided basic transaction &
others specific field for fee vote commit vote transaction
*/
func GenerateTxFeeVoteCommitment(
	tx *model.Transaction,
	voteHash []byte,
) *model.Transaction {
	var (
		txBody = &model.FeeVoteCommitTransactionBody{
			VoteHash: voteHash,
		}
		txBodyBytes, _ = (&transaction.FeeVoteCommitTransaction{Body: txBody}).GetBodyBytes()
	)
	tx.TransactionType = util.ConvertBytesToUint32(txTypeMap["feeVoteCommit"])
	tx.TransactionBody = &model.Transaction_FeeVoteCommitTransactionBody{
		FeeVoteCommitTransactionBody: txBody,
	}
	tx.TransactionBodyBytes = txBodyBytes
	tx.TransactionBodyLength = uint32(len(txBodyBytes))
	return tx
}

func GenerateTxFeeVoteRevealPhase(tx *model.Transaction, voteInfo *model.FeeVoteInfo, voteInfoSigned []byte) *model.Transaction {

	var (
		txBody = &model.FeeVoteRevealTransactionBody{
			FeeVoteInfo:    voteInfo,
			VoterSignature: voteInfoSigned,
		}
		txBodyBytes, _ = (&transaction.FeeVoteRevealTransaction{
			Body: txBody,
		}).GetBodyBytes()
	)
	tx.TransactionType = util.ConvertBytesToUint32(txTypeMap["feeVoteReveal"])
	tx.TransactionBody = &model.Transaction_FeeVoteRevealTransactionBody{
		FeeVoteRevealTransactionBody: txBody,
	}
	tx.TransactionBodyBytes = txBodyBytes
	tx.TransactionBodyLength = uint32(len(txBodyBytes))
	return tx
}

// GenerateTxLiquidPayment return liquid payment transaction based on provided basic transaction & ammunt
func GenerateTxLiquidPayment(tx *model.Transaction, sendAmount int64, completeMinutes uint64) *model.Transaction {
	txBody := &model.LiquidPaymentTransactionBody{
		Amount:          sendAmount,
		CompleteMinutes: completeMinutes,
	}
	tx.TransactionType = util.ConvertBytesToUint32(txTypeMap["liquidPayment"])
	tx.TransactionBody = &model.Transaction_LiquidPaymentTransactionBody{
		LiquidPaymentTransactionBody: txBody,
	}
	txBodyBytes, _ := (&transaction.LiquidPaymentTransaction{
		Body: txBody,
	}).GetBodyBytes()
	tx.TransactionBodyBytes = txBodyBytes
	tx.TransactionBodyLength = uint32(len(txBodyBytes))
	return tx
}

// GenerateTxLiquidPaymentStop return liquid payment stop transaction based on provided basic transaction & amount
func GenerateTxLiquidPaymentStop(tx *model.Transaction, transactionID int64) *model.Transaction {
	txBody := &model.LiquidPaymentStopTransactionBody{
		TransactionID: transactionID,
	}
	tx.TransactionType = util.ConvertBytesToUint32(txTypeMap["liquidPaymentStop"])
	tx.TransactionBody = &model.Transaction_LiquidPaymentStopTransactionBody{
		LiquidPaymentStopTransactionBody: txBody,
	}
	txBodyBytes, _ := (&transaction.LiquidPaymentStopTransaction{
		Body: txBody,
	}).GetBodyBytes()
	tx.TransactionBodyBytes = txBodyBytes
	tx.TransactionBodyLength = uint32(len(txBodyBytes))
	return tx
}

// GenerateTxAtomic return atomic transaction based on provided basic transaction
// inner transaction is send money with sender address is the tx.SenderAddress
func GenerateTxAtomic(tx *model.Transaction, innerCount uint32) *model.Transaction {
	var (
		atomicInnerTX = make(map[string]*model.AtomicInnerTransaction)
	)
	tx.TransactionType = util.ConvertBytesToUint32(txTypeMap["atomic"])
	for i := 0; i < int(innerCount); i++ {
		innerTx := GenerateBasicTransaction(
			senderAddressHex,
			"",
			senderSignatureType,
			version,
			timestamp,
			fee,
			recipientAccountAddressHex,
		)
		innerTx = GenerateTxSendMoney(tx, 10)
		unsignedTxBytes, err := (&transaction.Util{}).GetTransactionBytes(innerTx, false)
		if err != nil {
			return nil
		}
		atomicInnerTX[senderAddressHex] = &model.AtomicInnerTransaction{
			AtomicInnerItem: [][]byte{
				unsignedTxBytes,
			},
		}

	}
	return tx
}
