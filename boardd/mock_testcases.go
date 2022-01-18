package boardd

import "github.com/Ptt-official-app/go-pttbbs/ptttype"

var (
	testBoard6 = &Board{ // nolint
		Parent:     5,
		Bid:        6,
		Attributes: uint32(ptttype.BRD_POSTMASK),
		Name:       string([]byte{'A', 'L', 'L', 'P', 'O', 'S', 'T', 0x00, 0x2e, 0x2e, 0x2e, 0x2e}),
		Title: string([]byte{
			0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xb8, 0xf3, 0xaa,
			0x4f, 0xa6, 0xa1, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0xb7, 0x73,
			0xa4, 0xe5, 0xb3, 0xb9, 0x00, 0x20, 0xaf, 0xb8, 0xaa, 0xf8,
			0x20, 0x20, 0xa3, 0xad, 0xa1, 0x49, 0x00, 0x6e,
		}),
		RawModerators: "",
	}

	testBoard7 = &Board{ // nolint
		Parent: 5,
		Bid:    7,
		Name:   string([]byte{0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x00, 0x2e, 0x2e, 0x2e, 0x2e}),
		Title: string([]byte{
			0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xb8, 0xea, 0xb7,
			0xbd, 0xa6, 0x5e, 0xa6, 0xac, 0xb5, 0xa9, 0x00, 0xb7, 0x73,
			0xa4, 0xe5, 0xb3, 0xb9, 0x00, 0x20, 0xaf, 0xb8, 0xaa, 0xf8,
			0x20, 0x20, 0xa3, 0xad, 0xa1, 0x49, 0x00, 0x6e,
		}),
		RawModerators: "",
	}

	testBoard11 = &Board{ // nolint
		Parent: 5,
		Bid:    11,
		Name:   string([]byte{0x45, 0x64, 0x69, 0x74, 0x45, 0x78, 0x70, 0x00, 0x2e, 0x2e, 0x2e, 0x2e}),
		Title: string([]byte{
			0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xbd, 0x64, 0xa5,
			0xbb, 0xba, 0xeb, 0xc6, 0x46, 0xa7, 0xeb, 0xbd, 0x5a, 0xb0,
			0xcf, 0x00, 0xd6, 0xa1, 0x49, 0x00, 0xaf, 0xb8, 0xaa, 0xf8,
			0x20, 0x20, 0xa3, 0xad, 0xa1, 0x49, 0x00, 0x6e,
		}),
		RawModerators: "",
	}

	testBoard8 = &Board{
		Parent: 5,
		Bid:    8,
		Name:   string([]byte{0x4e, 0x6f, 0x74, 0x65, 0x00, 0x65, 0x64, 0x00, 0x2e, 0x2e, 0x2e, 0x2e}),
		Title: string([]byte{
			0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xb0, 0xca, 0xba,
			0x41, 0xac, 0xdd, 0xaa, 0x4f, 0xa4, 0xce, 0xba, 0x71, 0xa6,
			0xb1, 0xa7, 0xeb, 0xbd, 0x5a, 0x00, 0xaf, 0xb8, 0xaa, 0xf8,
			0x20, 0x20, 0xa3, 0xad, 0xa1, 0x49, 0x00, 0x6e,
		}),
		RawModerators: "",
	}

	testBoard1 = &Board{
		Parent:     2,
		Bid:        1,
		Attributes: uint32(ptttype.BRD_POSTMASK),
		Name:       string([]byte{'S', 'Y', 'S', 'O', 'P'}),
		Title: string([]byte{
			0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xaf, 0xb8, 0xaa,
			0xf8, 0xa6, 0x6e, 0x21,
		}),
		RawModerators: "",
	}

	testBoard9 = &Board{ // nolint
		Parent:     5,
		Bid:        9,
		Attributes: uint32(ptttype.BRD_POSTMASK),
		Name:       string([]byte{'R', 'e', 'c', 'o', 'r', 'd', 0x00, 0x00, 0x2e, 0x2e, 0x2e, 0x2e}),
		Title: string([]byte{
			0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xa7, 0xda, 0xad,
			0xcc, 0xaa, 0xba, 0xa6, 0xa8, 0xaa, 0x47, 0x00, 0x71, 0xa6,
			0xb1, 0xa7, 0xeb, 0xbd, 0x5a, 0x00, 0xaf, 0xb8, 0xaa, 0xf8,
			0x20, 0x20, 0xa3, 0xad, 0xa1, 0x49, 0x00, 0x6e, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		}),
		RawModerators: "",
	}

	testBoard10 = &Board{
		Parent: 5,
		Bid:    10,
		Name:   string([]byte{'W', 'h', 'o', 'A', 'm', 'I', 0x00, 0x00, 0x2e, 0x2e, 0x2e, 0x2e}),
		Title: string([]byte{
			0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xa8, 0xfe, 0xa8,
			0xfe, 0xa1, 0x41, 0xb2, 0x71, 0xb2, 0x71, 0xa7, 0xda, 0xac,
			0x4f, 0xbd, 0xd6, 0xa1, 0x49, 0x00, 0xaf, 0xb8, 0xaa, 0xf8,
			0x20, 0x20, 0xa3, 0xad, 0xa1, 0x49, 0x0, 0x6e,
		}),
		RawModerators: "",
	}
)