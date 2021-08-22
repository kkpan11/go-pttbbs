package cmsys

import "github.com/Ptt-official-app/go-pttbbs/ptttype"

var (
	testArticleSummary0 = &ptttype.ArticleSummaryRaw{
		Aid: 1,
		FileHeaderRaw: &ptttype.FileHeaderRaw{
			Filename: ptttype.Filename_t{ // M.1607937366.A.F6C
				0x4d, 0x2e, 0x31, 0x36, 0x30, 0x37, 0x39, 0x33,
				0x37, 0x33, 0x36, 0x36, 0x2e, 0x41, 0x2e, 0x46,
				0x36, 0x43,
			},
			Modified: 1607937365,
			Owner:    ptttype.Owner_t{0x74, 0x65, 0x73, 0x74},      // test
			Date:     ptttype.Date_t{0x31, 0x32, 0x2f, 0x31, 0x34}, // 12/14
			Title: ptttype.Title_t{ //[問題] p 幣要怎樣子賺呢？～
				0x5b, 0xb0, 0xdd, 0xc3, 0x44, 0x5d, 0x20, 0x70,
				0x20, 0xb9, 0xf4, 0xad, 0x6e, 0xab, 0xe7, 0xbc,
				0xcb, 0xa4, 0x6c, 0xc1, 0xc8, 0xa9, 0x4f, 0xa1,
				0x48, 0xa1, 0xe3,
			},
			Multi: [4]byte{0x05},
		},
	}
	testArticleSummary1 = &ptttype.ArticleSummaryRaw{
		Aid: 1,
		FileHeaderRaw: &ptttype.FileHeaderRaw{
			Filename: ptttype.Filename_t{ // M.1607202239.A.30D
				0x4d, 0x2e, 0x31, 0x36, 0x30, 0x37, 0x32, 0x30,
				0x32, 0x32, 0x33, 0x39, 0x2e, 0x41, 0x2e, 0x33,
				0x30, 0x44,
			},
			Modified: 1607202238,
			Owner:    ptttype.Owner_t{0x53, 0x59, 0x53, 0x4f, 0x50}, // SYSOP
			Date:     ptttype.Date_t{0x31, 0x32, 0x2f, 0x30, 0x36},  // 12/06
			Title: ptttype.Title_t{ //[問題] 我是誰？～
				0x5b, 0xb0, 0xdd, 0xc3, 0x44, 0x5d, 0x20, 0xa7,
				0xda, 0xac, 0x4f, 0xbd, 0xd6, 0xa1, 0x48, 0xa1,
				0xe3,
			},
		},
	}
	testArticleSummary2 = &ptttype.ArticleSummaryRaw{
		Aid: 2,
		FileHeaderRaw: &ptttype.FileHeaderRaw{
			Filename: ptttype.Filename_t{ // M.1607203395.A.F6C
				0x4d, 0x2e, 0x31, 0x36, 0x30, 0x37, 0x32, 0x30,
				0x33, 0x33, 0x39, 0x35, 0x2e, 0x41, 0x2e, 0x30,
				0x30, 0x44,
			},
			Modified: 1607203394,
			Owner:    ptttype.Owner_t{0x53, 0x59, 0x53, 0x4f, 0x50}, // SYSOP
			Date:     ptttype.Date_t{0x31, 0x32, 0x2f, 0x30, 0x36},  // 12/06
			Title: ptttype.Title_t{ //[心得] 然後呢？～
				0x5b, 0xa4, 0xdf, 0xb1, 0x6f, 0x5d, 0x20, 0xb5,
				0x4d, 0xab, 0xe1, 0xa9, 0x4f, 0xa1, 0x48, 0xa1,
				0xe3,
			},
			Filemode: ptttype.FILE_MARKED,
		},
	}
	testArticleSummary3 = &ptttype.ArticleSummaryRaw{
		Aid: 3,
		FileHeaderRaw: &ptttype.FileHeaderRaw{
			Filename: ptttype.Filename_t{ // M.1607203395.A.F6D
				0x4d, 0x2e, 0x31, 0x36, 0x30, 0x37, 0x32, 0x30,
				0x33, 0x33, 0x39, 0x35, 0x2e, 0x41, 0x2e, 0x30,
				0x30, 0x45,
			},
			Modified: 1607203394,
			Owner:    ptttype.Owner_t{0x53, 0x59, 0x53, 0x4f, 0x51}, // SYSOQ
			Date:     ptttype.Date_t{0x31, 0x32, 0x2f, 0x30, 0x36},  // 12/06
			Title: ptttype.Title_t{ //[心得] 然後呢？～
				0x5b, 0xa4, 0xdf, 0xb1, 0x6f, 0x5d, 0x20, 0xb5,
				0x4d, 0xab, 0xe1, 0xa9, 0x4f, 0xa1, 0x48, 0xa1,
				0xe3,
			},
			Filemode: ptttype.FILE_MARKED,
		},
	}
	testArticleSummary4 = &ptttype.ArticleSummaryRaw{
		Aid: 4,
		FileHeaderRaw: &ptttype.FileHeaderRaw{
			Filename: ptttype.Filename_t{ // M.1607203395.A.F6A
				0x4d, 0x2e, 0x31, 0x36, 0x30, 0x37, 0x32, 0x30,
				0x33, 0x33, 0x39, 0x35, 0x2e, 0x41, 0x2e, 0x30,
				0x30, 0x41,
			},
			Modified: 1607203394,
			Owner:    ptttype.Owner_t{0x53, 0x59, 0x53, 0x4f, 0x41}, // SYSOA
			Date:     ptttype.Date_t{0x31, 0x32, 0x2f, 0x30, 0x36},  // 12/06
			Title: ptttype.Title_t{ //[心得] 然後呢？～
				0x5b, 0xa4, 0xdf, 0xb1, 0x6f, 0x5d, 0x20, 0xb5,
				0x4d, 0xab, 0xe1, 0xa9, 0x4f, 0xa1, 0x48, 0xa1,
				0xe3,
			},
			Filemode: ptttype.FILE_MARKED,
		},
	}
	testArticleSummary5 = &ptttype.ArticleSummaryRaw{
		Aid: 5,
		FileHeaderRaw: &ptttype.FileHeaderRaw{
			Filename: ptttype.Filename_t{ // M.1607203396.A.F6A
				0x4d, 0x2e, 0x31, 0x36, 0x30, 0x37, 0x32, 0x30,
				0x33, 0x33, 0x39, 0x36, 0x2e, 0x41, 0x2e, 0x30,
				0x30, 0x41,
			},
			Modified: 1607203394,
			Owner:    ptttype.Owner_t{0x53, 0x59, 0x53, 0x4f, 0x41}, // SYSOA
			Date:     ptttype.Date_t{0x31, 0x32, 0x2f, 0x30, 0x36},  // 12/06
			Title: ptttype.Title_t{ //[心得] 然後呢？～
				0x5b, 0xa4, 0xdf, 0xb1, 0x6f, 0x5d, 0x20, 0xb5,
				0x4d, 0xab, 0xe1, 0xa9, 0x4f, 0xa1, 0x48, 0xa1,
				0xe3,
			},
			Filemode: ptttype.FILE_MARKED,
		},
	}
)
