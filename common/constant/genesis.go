// Constants to generate/validate genesis block
// Note: this file has been auto-generated by 'genesis generate' command
package constant

const (
	MainchainGenesisBlockID int64 = -3748523504887037308
)

type (
	GenesisConfigEntry struct {
		AccountAddress     string
		AccountBalance     int64
		NodePublicKey      []byte
		NodeAddress        string
		LockedBalance      int64
		ParticipationScore int64
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
	MainchainGenesisAccountAddress = "ZBC_AQTEH7K4_L45WJPLL_HCEC65ZH_7XC5N3XD_YNKPHK45_POH7PQME_AFAFBDWM"
	MainchainGenesisBlockSeed      = make([]byte, 64)
	MainchainGenesisNodePublicKey  = make([]byte, 32)
	GenesisConfig                  = []GenesisConfigEntry{
		{
			AccountAddress: "ZBC_ZFGNFDUB_UDNZUCNI_GPSQRQCU_2QD54YIH_ZXGCDFDJ_WLNHZVRB_INVJQTTR",
			AccountBalance: 0,
			NodePublicKey: []byte{164, 79, 0, 107, 231, 65, 39, 56, 34, 137, 87, 71, 46, 25, 202, 88, 201, 14, 28,
				197, 192, 218, 72, 64, 51, 14, 210, 98, 194, 226, 123, 250},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_DGNSM3EU_RFML2XKY_OUCE4AW6_PT5HGNW3_WW5JQBIS_XN7O3DHS_OOJNLRAY",
			AccountBalance: 0,
			NodePublicKey: []byte{76, 102, 195, 215, 187, 57, 47, 6, 81, 103, 16, 8, 41, 99, 86, 206, 202, 214, 166,
				252, 120, 3, 187, 239, 180, 45, 21, 96, 255, 36, 40, 159},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_VXJTE6PG_57ZDSJ6C_3AWYR3EI_MHIJQVNF_KO735JRF_27SERTWD_57H3RDEP",
			AccountBalance: 0,
			NodePublicKey: []byte{151, 252, 138, 168, 227, 149, 246, 45, 202, 202, 115, 72, 139, 153, 246, 157, 140, 222, 125,
				82, 4, 143, 46, 251, 161, 62, 149, 64, 152, 243, 110, 142},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_KAQWCVGR_5R566OBJ_DLSZQMKX_MBS6S2FL_27UB5HUL_DKXBMIGH_KO7XQFV2",
			AccountBalance: 0,
			NodePublicKey: []byte{171, 228, 199, 163, 160, 41, 14, 209, 185, 227, 115, 224, 234, 49, 184, 7, 238, 135, 78,
				160, 178, 234, 203, 109, 129, 132, 241, 184, 46, 132, 255, 33},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_TI7THWVS_24TEHF6H_CIE7NTBK_MLUDLC72_4CWRWWD2_PEW6SDSR_GZ7U4FKQ",
			AccountBalance: 0,
			NodePublicKey: []byte{191, 99, 18, 180, 37, 131, 87, 231, 113, 220, 163, 148, 191, 188, 193, 209, 177, 63, 241,
				188, 228, 194, 70, 164, 148, 35, 186, 178, 91, 71, 22, 76},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_KCAT4QDI_5H64UNT5_BWG2FXRY_QKQFU7DR_JUALSOED_3OQHZFYJ_LEMVF7HN",
			AccountBalance: 0,
			NodePublicKey: []byte{148, 127, 120, 253, 67, 170, 188, 180, 163, 203, 157, 17, 213, 217, 191, 188, 68, 3, 57,
				94, 27, 58, 64, 215, 7, 132, 151, 215, 107, 74, 154, 192},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_OLIVVD5V_BPEI2VRA_AERUCLYP_DBTOZGVA_QUPH7IUO_67ZNXAV4_ZPPDPHPB",
			AccountBalance: 0,
			NodePublicKey: []byte{130, 73, 155, 54, 41, 216, 63, 54, 82, 218, 171, 156, 110, 80, 40, 237, 35, 97, 18,
				138, 93, 30, 12, 141, 8, 193, 202, 41, 26, 194, 33, 225},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_TB36S5FH_4LGCVEF3_VNBHOENR_FRPIVTGZ_XQ2UKSTG_XUHLMV4A_5NC4KJZD",
			AccountBalance: 0,
			NodePublicKey: []byte{238, 230, 196, 164, 215, 93, 38, 98, 237, 221, 217, 166, 14, 18, 48, 224, 190, 131, 155,
				249, 74, 253, 83, 172, 150, 33, 42, 58, 61, 29, 253, 191},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_7BK5GIT3_2OI4GD67_ENFRTE3J_NTNFX2ZG_T7I2YP6D_VCRRLXS5_MCUBTNXF",
			AccountBalance: 0,
			NodePublicKey: []byte{38, 157, 20, 84, 33, 120, 124, 214, 139, 82, 54, 3, 80, 26, 62, 39, 226, 73, 152,
				153, 240, 13, 161, 223, 187, 3, 206, 251, 80, 75, 52, 195},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_5P2EAARC_BRR4GCSC_D5K2YLMG_ZVQMYNR6_JOOWBAOY_UG3257Y5_YCZZO4S4",
			AccountBalance: 0,
			NodePublicKey: []byte{205, 83, 100, 40, 67, 104, 224, 172, 208, 119, 210, 129, 36, 227, 159, 73, 1, 118, 185,
				138, 204, 203, 25, 12, 116, 236, 233, 155, 105, 251, 31, 56},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_YFTGNMPF_54ZQZASV_7QKDSYT7_XFYUNAZT_V7B2G4NN_YGCMEFHK_PB3GBUPR",
			AccountBalance: 0,
			NodePublicKey: []byte{234, 148, 121, 227, 151, 211, 72, 164, 252, 242, 246, 247, 48, 253, 165, 197, 33, 35, 47,
				33, 195, 51, 13, 250, 126, 223, 133, 200, 3, 113, 102, 54},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_HC7SBIGT_CWUVGNF3_KUBMYSQY_ZUENSSZS_FOCVV6BY_3CRTBNUA_JECG7LMQ",
			AccountBalance: 0,
			NodePublicKey: []byte{160, 193, 81, 111, 167, 145, 119, 240, 189, 184, 81, 232, 141, 169, 190, 15, 235, 138, 19,
				44, 160, 176, 71, 9, 76, 166, 23, 142, 193, 225, 115, 146},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_453A3GVP_W4MKTJYT_DCIKXZ3X_7HMZB6VN_FQZNFIPO_EFHFVTVQ_VJI64B4T",
			AccountBalance: 0,
			NodePublicKey: []byte{207, 184, 3, 118, 111, 181, 145, 30, 234, 39, 93, 199, 196, 69, 29, 239, 146, 11, 186,
				59, 124, 98, 33, 255, 248, 134, 63, 248, 25, 77, 15, 21},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_MVBI3TVQ_H2Y3XVST_OSLSEQJC_RPM3UDYP_2MAPSVPM_VM2UWPQF_52QMI2RR",
			AccountBalance: 0,
			NodePublicKey: []byte{77, 249, 254, 159, 84, 63, 69, 164, 193, 120, 88, 97, 154, 227, 110, 148, 35, 100, 8,
				136, 73, 77, 169, 42, 50, 212, 166, 17, 211, 80, 133, 189},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_JLKMA3NN_PTA7OXVC_YD6D4SEF_Z5W5NBVA_5WR7YI5W_4Q3TKSHG_J3Y62Z5B",
			AccountBalance: 0,
			NodePublicKey: []byte{143, 212, 71, 169, 73, 57, 251, 216, 204, 209, 196, 81, 43, 125, 215, 65, 157, 89, 11,
				111, 241, 74, 126, 207, 185, 139, 148, 31, 6, 7, 235, 121},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_3R63S7MU_ZCJJLVSW_QB5IT3S4_CTEYHTRC_YB5KVU4N_KSOFHKNM_HQUM5W4Z",
			AccountBalance: 0,
			NodePublicKey: []byte{189, 77, 29, 149, 61, 131, 203, 69, 219, 181, 107, 93, 38, 62, 164, 10, 205, 85, 198,
				22, 209, 165, 79, 139, 142, 209, 48, 160, 135, 217, 19, 208},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_WLMA4N6S_WA77IZFP_WSEPVYXK_I6KUEEPF_ZD54ODSO_EQ6Z3E6J_FZWKPAVG",
			AccountBalance: 0,
			NodePublicKey: []byte{97, 191, 150, 141, 207, 121, 250, 78, 62, 211, 97, 98, 184, 51, 17, 52, 101, 47, 246,
				199, 200, 208, 96, 30, 29, 222, 235, 30, 62, 221, 152, 106},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_6ZQVCLHK_IFUH7TZX_VN5OIJ5Z_V3WPBNHZ_MZHEX245_VUOTNYZW_6ABJNV7L",
			AccountBalance: 0,
			NodePublicKey: []byte{211, 155, 221, 141, 111, 58, 224, 168, 88, 12, 178, 125, 135, 74, 125, 34, 87, 162, 90,
				1, 121, 189, 65, 59, 20, 186, 22, 31, 53, 46, 13, 220},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_3CMJOD65_GHCPEAX3_7EBZW6JM_QRCQCIVK_P2FGSJ6E_SNXDI3O7_JVW54KSY",
			AccountBalance: 0,
			NodePublicKey: []byte{137, 109, 86, 100, 109, 40, 167, 183, 102, 31, 231, 72, 172, 184, 10, 166, 61, 30, 58,
				106, 216, 62, 136, 105, 157, 85, 109, 205, 214, 121, 84, 36},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_L7QXVYTY_HT5IPP5R_5RS2VYQS_PSM7G7AV_OCWBC5FB_4YR7DDZJ_7NVFA5VF",
			AccountBalance: 0,
			NodePublicKey: []byte{204, 67, 134, 72, 36, 104, 122, 60, 172, 214, 85, 28, 127, 30, 91, 121, 184, 230, 236,
				45, 17, 219, 103, 150, 206, 145, 85, 232, 87, 117, 116, 32},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_OTKVCM3J_6AKG2DC6_DAC6VISW_EUTR3GPZ_SEJFZXFX_GEPRUOFL_SNUS57HY",
			AccountBalance: 0,
			NodePublicKey: []byte{224, 242, 175, 80, 17, 165, 213, 239, 114, 161, 235, 207, 225, 99, 242, 94, 83, 179, 163,
				39, 81, 49, 141, 181, 165, 30, 195, 222, 233, 181, 149, 127},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_WY6GZJHI_QLIJOFH3_FUKZ253K_WDPSP6L6_BB2GXXZD_OABMPPTD_T4HJ5PPZ",
			AccountBalance: 0,
			NodePublicKey: []byte{44, 80, 222, 83, 90, 64, 144, 203, 87, 2, 245, 238, 241, 13, 210, 149, 199, 227, 251,
				220, 216, 73, 74, 98, 17, 207, 99, 254, 155, 137, 241, 77},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_OZ33G6D3_VAZGPQX3_CDILP7BL_E6KAG3IW_AY4PUUWE_E5SEZNJO_HUDSOLXP",
			AccountBalance: 0,
			NodePublicKey: []byte{195, 163, 244, 69, 88, 225, 117, 80, 74, 254, 30, 6, 238, 119, 179, 92, 177, 200, 154,
				189, 31, 57, 70, 180, 175, 157, 105, 199, 221, 223, 123, 19},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_M2MMNRDR_2M4OSRMA_Y7PEQNEC_JJLGAM37_UVSL3X7F_BANWBZH6_AZJOAA7C",
			AccountBalance: 0,
			NodePublicKey: []byte{76, 44, 237, 244, 17, 203, 77, 73, 189, 62, 189, 180, 53, 8, 205, 190, 235, 43, 31,
				184, 1, 208, 238, 85, 191, 217, 57, 58, 62, 198, 243, 115},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_RYZRLHH6_3JFUYH4H_QKLZANVU_Q4E7RR6H_ZUKTQSPQ_Z2OHZ6QP_TYRSQEZA",
			AccountBalance: 0,
			NodePublicKey: []byte{59, 173, 139, 184, 132, 171, 0, 223, 146, 140, 250, 205, 116, 0, 79, 106, 120, 10, 5,
				184, 248, 8, 210, 127, 9, 17, 9, 45, 121, 191, 16, 214},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_3RGCVJT4_4PBAGATS_6K4JVTHQ_XJTOM54O_LL2ARWQT_XMKFVGVS_EP327Y43",
			AccountBalance: 0,
			NodePublicKey: []byte{117, 184, 211, 177, 141, 122, 14, 210, 249, 124, 174, 223, 144, 107, 182, 219, 93, 127, 145,
				183, 71, 108, 153, 240, 135, 210, 244, 104, 216, 14, 71, 111},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_Z5MY6H7S_YASIVWQK_XFAXKRE4_35IZ6S5Y_OND5TZ36_D2VBVDZK_PLXB3QNF",
			AccountBalance: 0,
			NodePublicKey: []byte{89, 247, 55, 149, 56, 121, 236, 186, 182, 74, 96, 146, 11, 131, 9, 174, 8, 50, 213,
				31, 153, 165, 159, 226, 14, 240, 151, 171, 56, 11, 81, 181},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_AXSTVUQO_KYJI3YV2_W6VRBQCO_ZC2FKLW3_E5DDKDUJ_QPRO6Y5Y_RX47J4PR",
			AccountBalance: 0,
			NodePublicKey: []byte{11, 151, 242, 190, 153, 77, 200, 226, 170, 246, 195, 182, 233, 159, 85, 45, 4, 89, 144,
				105, 251, 92, 4, 156, 187, 210, 215, 181, 158, 30, 79, 77},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_PGBDE77L_M6POPSFK_6QRX7AEA_S4MCNT4M_OHZ7V72D_2DCR47ZP_2NMT4YHT",
			AccountBalance: 0,
			NodePublicKey: []byte{11, 77, 192, 191, 253, 22, 159, 146, 249, 52, 135, 145, 22, 32, 86, 166, 53, 229, 232,
				220, 210, 185, 38, 153, 63, 86, 2, 108, 16, 51, 112, 183},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_PH7Y2VPN_OC3OYKPC_GYAK25OE_ZNS2LN55_VZAC63G6_DVMDIGGV_NKOGLXFM",
			AccountBalance: 0,
			NodePublicKey: []byte{241, 200, 74, 67, 99, 152, 224, 64, 181, 158, 254, 185, 133, 3, 186, 230, 87, 170, 87,
				147, 226, 144, 23, 200, 94, 125, 253, 219, 108, 189, 47, 176},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_3BS6RH4E_CNOS77CO_ZXLNPR7Z_VUBJHEKA_QY64FLHW_KOA7T6AN_2FBSZ6OS",
			AccountBalance: 0,
			NodePublicKey: []byte{81, 121, 232, 229, 129, 0, 81, 133, 150, 83, 116, 87, 200, 135, 14, 11, 191, 99, 199,
				34, 229, 65, 86, 162, 76, 156, 44, 161, 98, 255, 6, 175},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_W7II4UKX_GOX3MJSV_252IXITE_G4RHC4ZL_TMRXUJSY_S5TOC434_GEVUTCW7",
			AccountBalance: 0,
			NodePublicKey: []byte{56, 117, 229, 171, 173, 147, 183, 79, 189, 179, 82, 4, 28, 181, 7, 30, 1, 99, 220,
				8, 169, 200, 174, 197, 174, 119, 121, 144, 234, 120, 134, 210},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_6JFFQ4BA_TRATZW3H_XJEOSVWI_ZFSXPPP4_LWCNKSU6_53HCV7QR_LN3DIX2H",
			AccountBalance: 0,
			NodePublicKey: []byte{43, 127, 237, 166, 195, 50, 111, 235, 159, 165, 201, 147, 165, 218, 44, 31, 192, 230, 33,
				164, 122, 124, 249, 77, 249, 176, 250, 19, 8, 119, 81, 253},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_X3LDC747_PWIL525Y_A6IIRZ6O_WIODY72H_XTNNWRVQ_43EST6FS_VJYDRVAA",
			AccountBalance: 0,
			NodePublicKey: []byte{108, 248, 190, 81, 87, 235, 102, 221, 161, 171, 187, 124, 60, 231, 9, 2, 208, 161, 202,
				190, 207, 175, 132, 46, 159, 101, 122, 81, 213, 134, 235, 57},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_XBIOMPVX_GWVFWTP3_3DUK3JP2_XRXFJZYJ_JFIXXCNA_SH7FNMCC_LGMJA56X",
			AccountBalance: 0,
			NodePublicKey: []byte{180, 93, 255, 224, 66, 103, 224, 118, 175, 234, 126, 112, 241, 40, 114, 235, 58, 222, 79,
				150, 40, 100, 251, 61, 22, 11, 92, 226, 118, 234, 24, 164},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_K5I2S3LX_6NY6QDAW_PL2QCSAR_ZIUBW3NU_T2ZFHE3L_EMMA5WVY_QJ2JGY7Y",
			AccountBalance: 0,
			NodePublicKey: []byte{147, 206, 165, 1, 201, 79, 138, 97, 52, 8, 79, 61, 58, 41, 224, 80, 31, 156, 2,
				217, 185, 213, 102, 184, 180, 253, 70, 216, 178, 1, 123, 204},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_FQATUUTQ_YZZFAOPD_RLIDN73M_Y7A5DNQE_X66JFHQM_C75YVY6P_DEV6DOWN",
			AccountBalance: 0,
			NodePublicKey: []byte{16, 171, 64, 117, 112, 146, 217, 194, 247, 113, 246, 94, 22, 0, 137, 120, 107, 35, 45,
				68, 237, 138, 61, 7, 146, 232, 54, 133, 20, 217, 171, 97},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_U52PUAXY_DR7JDG7N_IQ6MGBJD_CSHOJJKI_EWWDXOV7_MCYS32ET_S6A2BR6B",
			AccountBalance: 0,
			NodePublicKey: []byte{240, 248, 57, 19, 216, 206, 204, 98, 187, 107, 189, 229, 234, 60, 213, 39, 112, 44, 16,
				174, 253, 97, 11, 135, 160, 95, 26, 137, 88, 176, 11, 253},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_PTNC232Y_NAHSWT27_3SY5WXQ2_JXW32OZK_JYXTVHUA_SQ37AFFO_3MUPSBQJ",
			AccountBalance: 0,
			NodePublicKey: []byte{145, 218, 29, 132, 81, 96, 109, 113, 36, 253, 37, 42, 223, 32, 5, 237, 236, 96, 239,
				186, 194, 227, 60, 70, 191, 176, 73, 105, 239, 90, 111, 66},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_THZTR3IX_TUWCL34K_HGFIHWKS_6SXQSA2B_XDCDABO2_3NMWRXDN_BGRG73JJ",
			AccountBalance: 0,
			NodePublicKey: []byte{157, 0, 47, 98, 154, 34, 251, 136, 138, 29, 172, 13, 210, 41, 80, 63, 250, 175, 11,
				57, 55, 60, 99, 49, 60, 27, 15, 134, 120, 116, 242, 251},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_UGTSDPUX_5BN3VSJK_72MPAK22_5ORGTFSQ_CZIZMZRR_VGY3J7VN_N5DATDNW",
			AccountBalance: 0,
			NodePublicKey: []byte{240, 241, 45, 92, 100, 247, 118, 128, 15, 177, 188, 106, 241, 11, 109, 49, 244, 63, 173,
				207, 204, 115, 75, 162, 49, 185, 176, 49, 136, 157, 229, 117},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_LYQ2M6PL_WEG7TAXY_W4I6XGNP_OHNCJTOA_5XT4HCGK_WZVYN5K7_4QKTLV4G",
			AccountBalance: 0,
			NodePublicKey: []byte{87, 175, 101, 41, 130, 148, 235, 102, 84, 134, 204, 41, 214, 53, 107, 173, 212, 22, 167,
				64, 156, 86, 7, 195, 111, 242, 114, 29, 137, 162, 236, 227},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_XLHFQ3BU_NOLYZY2H_FC4VX7NA_UMJMV7J2_CUZZ4QMK_IIHPKJ2E_C5GAQAM3",
			AccountBalance: 0,
			NodePublicKey: []byte{137, 72, 237, 134, 177, 153, 235, 177, 15, 52, 17, 135, 153, 178, 237, 190, 81, 101, 191,
				115, 197, 162, 97, 8, 229, 103, 109, 24, 212, 113, 130, 229},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_MNYBIFMG_QQXGHMDL_PACDZZNR_I3LGTVV2_3FECXPET_BOUOS2S5_6VPGPWEE",
			AccountBalance: 0,
			NodePublicKey: []byte{90, 139, 117, 111, 179, 160, 202, 13, 211, 226, 228, 106, 93, 133, 178, 207, 217, 161, 90,
				215, 200, 163, 52, 77, 94, 113, 87, 75, 85, 161, 156, 76},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_ETXGYW2L_B4PAERRX_VKUU4TQ7_Y5AJRQMC_CM72BXSF_RD7AR4N5_SAQPZHNY",
			AccountBalance: 0,
			NodePublicKey: []byte{159, 166, 43, 2, 13, 176, 222, 18, 227, 136, 1, 146, 128, 48, 211, 177, 186, 44, 40,
				68, 140, 81, 179, 176, 186, 216, 20, 246, 254, 9, 97, 100},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_RERG3XD7_GAKOZZKY_VMZP2SQE_LBP45DC6_VDFGDTFK_3BZFBQGK_JMWELLO7",
			AccountBalance: 0,
			NodePublicKey: []byte{170, 183, 88, 172, 12, 251, 227, 191, 181, 213, 94, 106, 210, 175, 131, 150, 72, 155, 131,
				43, 219, 57, 214, 209, 195, 93, 1, 128, 232, 110, 4, 191},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_N4MLJ6QD_HVZRDRR5_QMEWR54I_NNQLHJXK_COHH2MFE_D3MU67SZ_4DWJPNGK",
			AccountBalance: 0,
			NodePublicKey: []byte{219, 207, 59, 152, 154, 166, 59, 4, 46, 133, 84, 68, 20, 38, 84, 138, 161, 21, 56,
				109, 17, 176, 163, 250, 18, 121, 118, 184, 200, 205, 155, 160},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_LGHLZ77R_23KXAXYT_MIIWRUUX_EBXBUDEJ_MJWOPFWD_SVMKGBJL_4N6LRLUF",
			AccountBalance: 0,
			NodePublicKey: []byte{165, 24, 43, 170, 30, 156, 56, 83, 18, 161, 208, 39, 129, 201, 141, 18, 102, 227, 40,
				176, 127, 94, 213, 77, 2, 86, 60, 45, 207, 252, 79, 172},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_Q7P6XRIC_QFLNPOEP_EI5ZOJRQ_KF5AWSN5_XMFCHDDZ_I3DZW2D2_I3NEY5P4",
			AccountBalance: 0,
			NodePublicKey: []byte{203, 4, 64, 91, 11, 51, 161, 68, 144, 30, 24, 15, 20, 237, 155, 13, 39, 104, 55,
				76, 129, 227, 135, 8, 57, 200, 25, 172, 238, 97, 217, 196},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_N4OFFA4E_ZZ4O5KU3_K6KGFN5T_IYATNM5Q_7WMQDLHM_LDQD2HUK_W2JZ6RVT",
			AccountBalance: 0,
			NodePublicKey: []byte{24, 37, 201, 143, 69, 165, 17, 190, 156, 188, 93, 141, 135, 12, 210, 6, 10, 68, 53,
				134, 144, 206, 65, 197, 245, 215, 8, 81, 108, 125, 176, 238},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_3LNOIUOP_52VRV4E7_AVGZXJH5_NXNOYTDE_VOSNAM2O_2UWKHRJN_GVSSYKIF",
			AccountBalance: 0,
			NodePublicKey: []byte{213, 154, 78, 44, 52, 255, 23, 247, 175, 64, 103, 38, 243, 239, 24, 150, 107, 35, 104,
				5, 204, 181, 15, 151, 38, 41, 100, 246, 234, 100, 93, 199},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_QMZXQEDK_JHQW7TZ2_JZSEHGMQ_6OKIMQHP_IKH3GZV2_ZUSJCXI5_LVQ5NNUC",
			AccountBalance: 0,
			NodePublicKey: []byte{69, 32, 114, 232, 70, 138, 30, 217, 45, 103, 170, 178, 45, 5, 226, 123, 209, 19, 199,
				59, 204, 202, 90, 19, 220, 252, 13, 173, 9, 129, 126, 144},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_B3GEXDCX_SROWXC3I_PCTEVVI4_MIQ2ELEZ_7NMQG7PJ_XT5J2TDV_FKXEFCOU",
			AccountBalance: 0,
			NodePublicKey: []byte{47, 82, 66, 168, 216, 3, 237, 10, 211, 119, 92, 179, 54, 146, 124, 158, 168, 9, 123,
				194, 174, 44, 41, 155, 119, 95, 49, 104, 202, 146, 99, 251},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_CDPUKUG5_45VAQ4BT_2KETZQYA_7QUSYQEV_GXSM7FXS_Q67SQKI5_D57F53H2",
			AccountBalance: 0,
			NodePublicKey: []byte{12, 79, 130, 200, 184, 241, 13, 136, 82, 70, 195, 242, 120, 217, 154, 209, 5, 76, 73,
				210, 45, 17, 55, 162, 178, 14, 196, 33, 42, 183, 55, 163},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_ARHIVIHU_F4MKHYBV_GIKYSSY2_ZOTGTLJQ_2XTNJEUV_NA3RTVQJ_VSKY67LL",
			AccountBalance: 0,
			NodePublicKey: []byte{85, 230, 157, 112, 124, 153, 122, 229, 92, 200, 151, 201, 18, 111, 102, 5, 136, 239, 236,
				152, 183, 158, 20, 32, 72, 72, 225, 197, 200, 82, 227, 139},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_VTTWO3DM_QIA4AGKN_GEOOTHMC_UAVK36XF_UVYDDLEJ_AXM2AQN6_EES52PES",
			AccountBalance: 0,
			NodePublicKey: []byte{8, 104, 214, 113, 133, 103, 168, 215, 253, 144, 120, 17, 27, 129, 110, 44, 228, 155, 62,
				209, 20, 26, 180, 68, 126, 205, 202, 21, 119, 25, 164, 24},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_S6T5KN6X_IB57IV47_BE56CSDR_L2UWEE4M_K6MBGM3Z_RU5MKN73_VTU62DKV",
			AccountBalance: 0,
			NodePublicKey: []byte{201, 54, 198, 82, 189, 28, 139, 25, 230, 74, 24, 23, 216, 125, 179, 108, 106, 151, 50,
				181, 118, 202, 16, 240, 127, 210, 211, 104, 239, 240, 82, 27},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
	}
)
