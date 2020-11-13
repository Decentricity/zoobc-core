// Constants to generate/validate genesis block
// Note: this file has been auto-generated by 'genesis generate' command
package constant

const (
	MainchainGenesisBlockID int64 = -8052505283446226463
)

type (
	GenesisConfigEntry struct {
		AccountAddressType int32
		AccountAddress     string
		AccountBalance     int64
		NodePublicKey      []byte
		NodeAddress        string
		LockedBalance      int64
		ParticipationScore int64
		Message            string
	}
)

var (
	ApplicationCodeName            = "ZBC_main"
	ApplicationVersion             = "1.0.0"
	MainchainGenesisBlocksmithID   = []byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	MainchainGenesisBlockSignature = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	MainchainGenesisTransactionSignature = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	MainchainGenesisBlockTimestamp = int64(1596708000)
	// MainchainGenesisAccountAddress encoded "ZBC_AQTEH7K4_L45WJPLL_HCEC65ZH_7XC5N3XD_YNKPHK45_POH7PQME_AFAFBDWM"
	// MainchainGenesisAccountAddress hex "042643fd5c5f3b64bd6b38882f7727fdc5d6eee3c354f3ab9d7b8ff7c1840140"
	MainchainGenesisAccountAddress = []byte{0, 0, 0, 0, 4, 38, 67, 253, 92, 95, 59, 100, 189, 107, 56, 136, 47, 119, 39, 253, 197, 214,
		238, 227, 195, 84, 243, 171, 157, 123, 143, 247, 193, 132, 1, 64}
	MainchainGenesisBlockSeed     = make([]byte, 64)
	MainchainGenesisNodePublicKey = make([]byte, 32)
	GenesisConfig                 = []GenesisConfigEntry{
		{
			AccountAddressType: 0,
			AccountAddress:     "ZBC_F5YUYDXD_WFDJSAV5_K3Y72RCM_GLQP32XI_QDVXOGGD_J7CGSSSK_5VKR7YML",
			AccountBalance:     0,
			NodePublicKey: []byte{153, 58, 50, 200, 7, 61, 108, 229, 204, 48, 199, 145, 21, 99, 125, 75, 49, 45, 118,
				97, 219, 80, 242, 244, 100, 134, 144, 246, 37, 144, 213, 135},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
			Message:            "Fifer",
		},
		{
			AccountAddressType: 0,
			AccountAddress:     "ZBC_EEX2EIIS_UFQI5SRU_4UGYMSXA_RDSKDNAF_P2A7M5PC_CXWCMZZX_SNPIQ2SW",
			AccountBalance:     0,
			NodePublicKey: []byte{0, 14, 6, 218, 170, 54, 60, 50, 2, 66, 130, 119, 226, 235, 126, 203, 5, 12, 152,
				194, 170, 146, 43, 63, 224, 101, 127, 241, 62, 152, 187, 255},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
			Message:            "Fiddler",
		},
		{
			AccountAddressType: 0,
			AccountAddress:     "ZBC_3WWDF4S2_IZVG2HHD_VOPSCNGN_COLYZ2OZ_M4QJZ4OL_44YHTKVC_2TPZBZAU",
			AccountBalance:     0,
			NodePublicKey: []byte{91, 36, 228, 70, 101, 94, 186, 246, 186, 4, 78, 142, 173, 162, 187, 173, 202, 81, 243,
				92, 141, 120, 148, 220, 41, 160, 208, 94, 174, 166, 62, 207},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
			Message:            "Practical",
		},
	}
)
