package cmbbs

import "github.com/Ptt-official-app/go-pttbbs/ptttype"

var (
	testUserecRaw1 = &ptttype.UserecRaw{
		Version:    ptttype.PASSWD_VERSION,
		UserID:     ptttype.UserID_t{83, 89, 83, 79, 80},                         // SYSOP
		RealName:   ptttype.RealName_t{67, 111, 100, 105, 110, 103, 77, 97, 110}, // CodingMan
		Nickname:   ptttype.Nickname_t{175, 171},                                 // 神
		PasswdHash: ptttype.Passwd_t{98, 104, 119, 118, 79, 74, 116, 102, 84, 49, 84, 65, 73, 0},

		UFlag:        ptttype.UF_CURSOR_ASCII | ptttype.UF_DBCS_DROP_REPEAT | ptttype.UF_DBCS_AWARE | ptttype.UF_ADBANNER | ptttype.UF_BRDSORT,
		UserLevel:    ptttype.PERM_DEFAULT | ptttype.PERM_SYSSUBOP | ptttype.PERM_BM,
		NumLoginDays: 2,
		NumPosts:     0,
		FirstLogin:   1600681288,
		LastLogin:    1600756094,
		LastHost:     ptttype.IPv4_t{53, 57, 46, 49, 50, 52, 46, 49, 54, 55, 46, 50, 50, 54},
		Address:      ptttype.Address_t{183, 115, 166, 203, 191, 164, 164, 108, 181, 234, 182, 109, 175, 81, 166, 179, 167, 248, 53, 52, 51, 184, 185},
		Over18:       true,
		Pager:        ptttype.PAGER_ON,
		Career:       ptttype.Career_t{165, 254, 180, 186, 179, 110, 197, 233}, // 全景軟體
		LastSeen:     1600681288,
	}

	testUserecRaw1Load = &ptttype.UserecRaw{
		Version:    ptttype.PASSWD_VERSION,
		UserID:     ptttype.UserID_t{83, 89, 83, 79, 80},                         // SYSOP
		RealName:   ptttype.RealName_t{67, 111, 100, 105, 110, 103, 77, 97, 110}, // CodingMan
		Nickname:   ptttype.Nickname_t{175, 171},                                 // 神
		PasswdHash: ptttype.Passwd_t{98, 104, 119, 118, 79, 74, 116, 102, 84, 49, 84, 65, 73, 0},

		UFlag:        ptttype.UF_CURSOR_ASCII | ptttype.UF_DBCS_DROP_REPEAT | ptttype.UF_DBCS_AWARE | ptttype.UF_ADBANNER | ptttype.UF_BRDSORT,
		UserLevel:    ptttype.PERM_DEFAULT | ptttype.PERM_SYSSUBOP | ptttype.PERM_BM,
		NumLoginDays: 2,
		NumPosts:     100,
		FirstLogin:   1600681288,
		LastLogin:    1600756094,
		LastHost:     ptttype.IPv4_t{53, 57, 46, 49, 50, 52, 46, 49, 54, 55, 46, 50, 50, 54},
		Address:      ptttype.Address_t{183, 115, 166, 203, 191, 164, 164, 108, 181, 234, 182, 109, 175, 81, 166, 179, 167, 248, 53, 52, 51, 184, 185},
		Over18:       true,
		Pager:        ptttype.PAGER_ON,
		Career:       ptttype.Career_t{165, 254, 180, 186, 179, 110, 197, 233}, // 全景軟體
		LastSeen:     1600681288,
	}

	testUserecRaw1Updated = &ptttype.UserecRaw{
		Version:    ptttype.PASSWD_VERSION,
		UserID:     ptttype.UserID_t{83, 89, 83, 79, 80},                         // SYSOP
		RealName:   ptttype.RealName_t{67, 111, 100, 105, 110, 103, 77, 97, 110}, // CodingMan
		Nickname:   ptttype.Nickname_t{175, 171},                                 // 神
		PasswdHash: ptttype.Passwd_t{98, 104, 119, 118, 79, 74, 116, 102, 84, 49, 84, 65, 73, 0},

		UFlag:        ptttype.UF_CURSOR_ASCII | ptttype.UF_DBCS_DROP_REPEAT | ptttype.UF_DBCS_AWARE | ptttype.UF_ADBANNER | ptttype.UF_BRDSORT,
		UserLevel:    ptttype.PERM_DEFAULT | ptttype.PERM_SYSSUBOP | ptttype.PERM_BM,
		NumLoginDays: 2,
		NumPosts:     100,
		FirstLogin:   1600681288,
		LastLogin:    1600756094,
		LastHost:     ptttype.IPv4_t{53, 57, 46, 49, 50, 52, 46, 49, 54, 55, 46, 50, 50, 54},
		Address:      ptttype.Address_t{183, 115, 166, 203, 191, 164, 164, 108, 181, 234, 182, 109, 175, 81, 166, 179, 167, 248, 53, 52, 51, 184, 185},
		Over18:       true,
		Pager:        ptttype.PAGER_ON,
		Career:       ptttype.Career_t{165, 254, 180, 186, 179, 110, 197, 233}, // 全景軟體
		LastSeen:     1600681288,
	}
)
