package cache

import (
	"reflect"
	"sync"
	"testing"
	"unsafe"

	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/types"
)

func TestGetBCache(t *testing.T) {
	setupTest()
	defer teardownTest()

	boards := [3]ptttype.BoardHeaderRaw{testBoardHeader0, testBoardHeader1, testBoardHeader2}
	Shm.WriteAt(
		unsafe.Offsetof(Shm.Raw.BCache),
		unsafe.Sizeof(boards),
		unsafe.Pointer(&boards),
	)

	type args struct {
		bidInCache ptttype.Bid
	}
	tests := []struct {
		name          string
		args          args
		expectedBoard *ptttype.BoardHeaderRaw
		wantErr       bool
	}{
		// TODO: Add test cases.
		{
			args:          args{1},
			expectedBoard: &testBoardHeader0,
		},
		{
			args:          args{2},
			expectedBoard: &testBoardHeader1,
		},
		{
			args:          args{3},
			expectedBoard: &testBoardHeader2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBoard, err := GetBCache(tt.args.bidInCache)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBoard, tt.expectedBoard) {
				t.Errorf("GetBCache() = %v, want %v", gotBoard, tt.expectedBoard)
			}
		})
	}
}

func TestIsHiddenBoardFriend(t *testing.T) {
	setupTest()
	defer teardownTest()

	_ = LoadUHash()

	ReloadBCache()

	type args struct {
		bidInCache ptttype.BidInStore
		uidInCache ptttype.UidInStore
	}
	tests := []struct {
		name     string
		args     args
		expected bool
	}{
		// TODO: Add test cases.
		{
			args:     args{0, 0}, //board: SYSOP user: SYSOP
			expected: true,
		},
		{
			args:     args{0, 1}, //board: SYSOP user: CodingMan
			expected: false,
		},
		{
			args:     args{0, 2}, //board: SYSOP user: pichu
			expected: false,
		},
		{
			args:     args{0, 3}, //board: SYSOP user: Kahou
			expected: true,
		},
		{
			args:     args{0, 4}, //board: SYSOP user: Kahou2
			expected: false,
		},
		{
			args:     args{0, 5}, //board: SYSOP user: (non-exist)
			expected: false,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if got := IsHiddenBoardFriend(tt.args.bidInCache, tt.args.uidInCache); got != tt.expected {
				t.Errorf("IsHiddenBoardFriend() = %v, want %v", got, tt.expected)
			}
		})
		wg.Wait()
	}
}

func TestNumBoards(t *testing.T) {
	setupTest()
	defer teardownTest()

	ReloadBCache()

	tests := []struct {
		name     string
		expected int32
	}{
		// TODO: Add test cases.
		{
			expected: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumBoards(); got != tt.expected {
				t.Errorf("NumBoards() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestReloadBCache(t *testing.T) {
	setupTest()
	defer teardownTest()

	tests := []struct {
		name                  string
		expectedNBoard        int32
		expectedSortedByName  []ptttype.BidInStore
		expectedSortedByClass []ptttype.BidInStore
		expectedBCacheName    []ptttype.BoardID_t
		expectedBCacheTitle   []ptttype.BoardTitle_t
	}{
		// TODO: Add test cases.
		{
			expectedNBoard:        int32(12),
			expectedSortedByName:  testSortedByName,
			expectedSortedByClass: testSortedByClass,
			expectedBCacheName:    testBCacheName,
			expectedBCacheTitle:   testBCacheTitle,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReloadBCache()

			nBoard := int32(0)
			Shm.ReadAt(
				unsafe.Offsetof(Shm.Raw.BNumber),
				types.INT32_SZ,
				unsafe.Pointer(&nBoard),
			)

			if !reflect.DeepEqual(nBoard, tt.expectedNBoard) {
				t.Errorf("ReloadBCache() = %v, want %v", nBoard, tt.expectedNBoard)
			}

			bsorted := [ptttype.BSORT_BY_MAX][ptttype.MAX_BOARD]ptttype.BidInStore{}
			Shm.ReadAt(
				unsafe.Offsetof(Shm.Raw.BSorted),
				unsafe.Sizeof(bsorted),
				unsafe.Pointer(&bsorted),
			)

			for idx := int32(0); idx < nBoard; idx++ {
				board, _ := GetBCache(ptttype.Bid(idx + 1))
				if types.Cstrcmp(board.Brdname[:], tt.expectedBCacheName[idx][:]) != 0 {
					t.Errorf("bcacheName (%v/%v) = %v, want %v", idx, nBoard, string(board.Brdname[:]), string(tt.expectedBCacheName[idx][:]))
				}
				if types.Cstrcmp(board.Title[:], tt.expectedBCacheTitle[idx][:]) != 0 {
					t.Errorf("bcacheTitle (%v/%v) = %v, want %v", idx, nBoard, string(board.Title[:]), string(tt.expectedBCacheTitle[idx][:]))
				}
			}

			bsortedByName := bsorted[ptttype.BSORT_BY_NAME][:nBoard]
			bsortedByClass := bsorted[ptttype.BSORT_BY_CLASS][:nBoard]
			if !reflect.DeepEqual(bsortedByName, tt.expectedSortedByName) {
				t.Errorf("bsorted-by-name = %v, want %v", bsortedByName, tt.expectedSortedByName)
			}
			if !reflect.DeepEqual(bsortedByClass, tt.expectedSortedByClass) {
				t.Errorf("bsorted-by-class = %v, want %v", bsortedByClass, tt.expectedSortedByClass)
			}

		})
	}
}

func Test_reloadCacheLoadBottom(t *testing.T) {
	setupTest()
	defer teardownTest()

	ReloadBCache()

	tests := []struct {
		name     string
		expected uint8
	}{
		// TODO: Add test cases.
		{
			expected: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reloadCacheLoadBottom()
			nBottom := uint8(0)
			Shm.ReadAt(
				unsafe.Offsetof(Shm.Raw.NBottom)+unsafe.Sizeof(nBottom)*9,
				unsafe.Sizeof(nBottom),
				unsafe.Pointer(&nBottom),
			)

			if nBottom != tt.expected {
				t.Errorf("nBottom: %v want: %v", nBottom, tt.expected)
			}
		})
	}
}

func TestGetBTotal(t *testing.T) {
	setupTest()
	defer teardownTest()

	ReloadBCache()

	bid1 := ptttype.Bid(3)
	bid1InCache := bid1.ToBidInStore()
	total1 := int32(5)

	Shm.WriteAt(
		unsafe.Offsetof(Shm.Raw.Total)+types.INT32_SZ*uintptr(bid1InCache),
		types.INT32_SZ,
		unsafe.Pointer(&total1),
	)

	type args struct {
		bid ptttype.Bid
	}
	tests := []struct {
		name          string
		args          args
		expectedTotal int32
	}{
		// TODO: Add test cases.
		{
			args:          args{bid1},
			expectedTotal: total1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTotal := GetBTotal(tt.args.bid); gotTotal != tt.expectedTotal {
				t.Errorf("GetBTotal() = %v, want %v", gotTotal, tt.expectedTotal)
			}
		})
	}
}

func TestSetBTotal(t *testing.T) {
	setupTest()
	defer teardownTest()

	ReloadBCache()

	type args struct {
		bid ptttype.Bid
	}
	tests := []struct {
		name                 string
		args                 args
		wantErr              bool
		expectedTotal        int32
		expectedLastPostTime types.Time4
	}{
		// TODO: Add test cases.
		{
			args:                 args{10},
			expectedTotal:        2,
			expectedLastPostTime: 1607203395,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetBTotal(tt.args.bid); (err != nil) != tt.wantErr {
				t.Errorf("SetBTotal() error = %v, wantErr %v", err, tt.wantErr)
			}

			total := GetBTotal(tt.args.bid)
			if total != tt.expectedTotal {
				t.Errorf("SetBTotal: total: %v want: %v", total, tt.expectedTotal)
			}

			bidInCache := tt.args.bid.ToBidInStore()
			lastPostTime := types.Time4(0)
			Shm.ReadAt(
				unsafe.Offsetof(Shm.Raw.LastPostTime)+types.TIME4_SZ*uintptr(bidInCache),
				types.TIME4_SZ,
				unsafe.Pointer(&lastPostTime),
			)
			if lastPostTime != tt.expectedLastPostTime {
				t.Errorf("SetBTotal: lastPostTime: %v want: %v", lastPostTime, tt.expectedLastPostTime)
			}
		})
	}
}

func TestSetBottomTotal(t *testing.T) {
	setupTest()
	defer teardownTest()

	ReloadBCache()

	type args struct {
		bid ptttype.Bid
	}
	tests := []struct {
		name          string
		args          args
		wantErr       bool
		expectedTotal uint8
	}{
		// TODO: Add test cases.
		{
			args:          args{10},
			expectedTotal: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetBottomTotal(tt.args.bid); (err != nil) != tt.wantErr {
				t.Errorf("SetBottomTotal() error = %v, wantErr %v", err, tt.wantErr)
			}

			bidInCache := tt.args.bid.ToBidInStore()
			total := uint8(0)
			const uint8sz = unsafe.Sizeof(total)

			Shm.ReadAt(
				unsafe.Offsetof(Shm.Raw.NBottom)+uint8sz*uintptr(bidInCache),
				uint8sz,
				unsafe.Pointer(&total),
			)
			if total != tt.expectedTotal {
				t.Errorf("SetBottomTotal: total: %v want: %v", total, tt.expectedTotal)
			}

		})
	}
}

func TestGetBid(t *testing.T) {
	setupTest()
	defer teardownTest()

	ReloadBCache()

	boardID0 := &ptttype.BoardID_t{}
	copy(boardID0[:], []byte("WhoAmI"))

	boardID1 := &ptttype.BoardID_t{}
	copy(boardID1[:], []byte("SYSOP"))

	type args struct {
		boardID *ptttype.BoardID_t
	}
	tests := []struct {
		name        string
		args        args
		expectedBid ptttype.Bid
		wantErr     bool
	}{
		// TODO: Add test cases.
		{
			args:        args{boardID: boardID0},
			expectedBid: 10,
		},
		{
			args:        args{boardID: boardID1},
			expectedBid: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBid, err := GetBid(tt.args.boardID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBid, tt.expectedBid) {
				t.Errorf("GetBid() = %v, want %v", gotBid, tt.expectedBid)
			}
		})
	}
}

func TestFindBoardIdxByName(t *testing.T) {
	setupTest()
	defer teardownTest()

	ReloadBCache()

	boardID0 := &ptttype.BoardID_t{}
	copy(boardID0[:], []byte("WhoAmI"))

	boardID1 := &ptttype.BoardID_t{}
	copy(boardID1[:], []byte("SYSOP"))

	boardID2 := &ptttype.BoardID_t{}
	copy(boardID2[:], []byte("2..........."))

	boardID3 := &ptttype.BoardID_t{}
	copy(boardID3[:], []byte("EditExp"))

	boardID4 := &ptttype.BoardID_t{}
	copy(boardID4[:], []byte("WhoAmA"))

	boardID5 := &ptttype.BoardID_t{}
	copy(boardID5[:], []byte("junk"))

	boardID6 := &ptttype.BoardID_t{}
	copy(boardID6[:], []byte("Note"))

	boardID7 := &ptttype.BoardID_t{}
	copy(boardID7[:], []byte("Record"))

	boardID8 := &ptttype.BoardID_t{}
	copy(boardID8[:], []byte("WhoAmJ"))

	boardID9 := &ptttype.BoardID_t{}
	copy(boardID9[:], []byte("a0"))

	type args struct {
		boardID *ptttype.BoardID_t
		isAsc   bool
	}
	tests := []struct {
		name        string
		args        args
		expectedIdx ptttype.SortIdx
		wantErr     bool
	}{
		// TODO: Add test cases.
		{
			args:        args{boardID: boardID0, isAsc: true},
			expectedIdx: 12,
		},
		{
			args:        args{boardID: boardID1, isAsc: true},
			expectedIdx: 11,
		},
		{
			args:        args{boardID: boardID2, isAsc: true},
			expectedIdx: 2,
		},
		{
			args:        args{boardID: boardID3, isAsc: true},
			expectedIdx: 6,
		},
		{
			args:        args{boardID: boardID4, isAsc: true},
			expectedIdx: 12,
		},
		{
			args:        args{boardID: boardID5, isAsc: true},
			expectedIdx: 7,
		},
		{
			args:        args{boardID: boardID6, isAsc: true},
			expectedIdx: 8,
		},
		{
			args:        args{boardID: boardID7, isAsc: true},
			expectedIdx: 9,
		},
		{
			args:        args{boardID: boardID8, isAsc: true},
			expectedIdx: -1,
		},
		{
			args:        args{boardID: boardID9, isAsc: true},
			expectedIdx: 3,
		},

		{
			args:        args{boardID: boardID0, isAsc: false},
			expectedIdx: 12,
		},
		{
			args:        args{boardID: boardID1, isAsc: false},
			expectedIdx: 11,
		},
		{
			args:        args{boardID: boardID2, isAsc: false},
			expectedIdx: 2,
		},
		{
			args:        args{boardID: boardID3, isAsc: false},
			expectedIdx: 6,
		},
		{
			args:        args{boardID: boardID4, isAsc: false},
			expectedIdx: 11,
		},
		{
			args:        args{boardID: boardID5, isAsc: false},
			expectedIdx: 7,
		},
		{
			args:        args{boardID: boardID6, isAsc: false},
			expectedIdx: 8,
		},
		{
			args:        args{boardID: boardID7, isAsc: false},
			expectedIdx: 9,
		},
		{
			args:        args{boardID: boardID8, isAsc: false},
			expectedIdx: 12,
		},
		{
			args:        args{boardID: boardID9, isAsc: false},
			expectedIdx: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIdx, err := FindBoardIdxByName(tt.args.boardID, tt.args.isAsc)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindBoardIdxByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotIdx, tt.expectedIdx) {
				t.Errorf("FindBoardIdxByName() = %v, want %v", gotIdx, tt.expectedIdx)
			}
		})
	}
}

func TestFindBoardIdxByClass(t *testing.T) {
	setupTest()
	defer teardownTest()

	ReloadBCache()

	testTitle0 := make([]byte, 4)
	copy(testTitle0, testBCacheTitle[0][:])

	testBrdname0 := &ptttype.BoardID_t{}
	copy(testBrdname0[:], []byte("WhoAMA"))

	testTitle1 := make([]byte, 4)
	copy(testTitle1, testBCacheTitle[0][:])
	testTitle1[3]--

	testTitle2 := make([]byte, 4)
	copy(testTitle2, testBCacheTitle[0][:])
	testTitle2[3]++

	type args struct {
		cls     []byte
		boardID *ptttype.BoardID_t
		isAsc   bool
	}
	tests := []struct {
		name        string
		args        args
		expectedIdx ptttype.SortIdx
		wantErr     bool
	}{
		// TODO: Add test cases.
		{
			args:        args{cls: testBCacheTitle[0][:4], boardID: &testBCacheName[0], isAsc: true},
			expectedIdx: 11,
		},
		{
			args:        args{cls: testBCacheTitle[1][:4], boardID: &testBCacheName[1], isAsc: true},
			expectedIdx: 1,
		},
		{
			args:        args{cls: testBCacheTitle[2][:4], boardID: &testBCacheName[2], isAsc: true},
			expectedIdx: 3,
		},
		{
			args:        args{cls: testBCacheTitle[3][:4], boardID: &testBCacheName[3], isAsc: true},
			expectedIdx: 4,
		},
		{
			args:        args{cls: testBCacheTitle[4][:4], boardID: &testBCacheName[4], isAsc: true},
			expectedIdx: 2,
		},
		{
			args:        args{cls: testBCacheTitle[5][:4], boardID: &testBCacheName[5], isAsc: true},
			expectedIdx: 6,
		},
		{
			args:        args{cls: testBCacheTitle[6][:4], boardID: &testBCacheName[6], isAsc: true},
			expectedIdx: 7,
		},
		{
			args:        args{cls: testBCacheTitle[7][:4], boardID: &testBCacheName[7], isAsc: true},
			expectedIdx: 9,
		},
		{
			args:        args{cls: testBCacheTitle[8][:4], boardID: &testBCacheName[8], isAsc: true},
			expectedIdx: 10,
		},
		{
			args:        args{cls: testBCacheTitle[9][:4], boardID: &testBCacheName[9], isAsc: true},
			expectedIdx: 12,
		},
		{
			args:        args{cls: testBCacheTitle[10][:4], boardID: &testBCacheName[10], isAsc: true},
			expectedIdx: 8,
		},
		{
			args:        args{cls: testBCacheTitle[11][:4], boardID: &testBCacheName[11], isAsc: true},
			expectedIdx: 5,
		},
		{
			name:        "title0-brdname0, WhoAMA, asc",
			args:        args{cls: testTitle0, boardID: testBrdname0, isAsc: true},
			expectedIdx: 12,
		},
		{
			name:        "title1-brdname0, between Security/AllHIDPOST, asc",
			args:        args{cls: testTitle1, boardID: testBrdname0, isAsc: true},
			expectedIdx: 5,
		},
		{
			name:        "title2-brdname0, after WhoAmI, asc",
			args:        args{cls: testTitle2, boardID: testBrdname0, isAsc: true},
			expectedIdx: -1,
		},

		{
			args:        args{cls: testBCacheTitle[0][:4], boardID: &testBCacheName[0], isAsc: false},
			expectedIdx: 11,
		},
		{
			args:        args{cls: testBCacheTitle[1][:4], boardID: &testBCacheName[1], isAsc: false},
			expectedIdx: 1,
		},
		{
			args:        args{cls: testBCacheTitle[2][:4], boardID: &testBCacheName[2], isAsc: false},
			expectedIdx: 3,
		},
		{
			args:        args{cls: testBCacheTitle[3][:4], boardID: &testBCacheName[3], isAsc: false},
			expectedIdx: 4,
		},
		{
			args:        args{cls: testBCacheTitle[4][:4], boardID: &testBCacheName[4], isAsc: false},
			expectedIdx: 2,
		},
		{
			args:        args{cls: testBCacheTitle[5][:4], boardID: &testBCacheName[5], isAsc: false},
			expectedIdx: 6,
		},
		{
			args:        args{cls: testBCacheTitle[6][:4], boardID: &testBCacheName[6], isAsc: false},
			expectedIdx: 7,
		},
		{
			args:        args{cls: testBCacheTitle[7][:4], boardID: &testBCacheName[7], isAsc: false},
			expectedIdx: 9,
		},
		{
			args:        args{cls: testBCacheTitle[8][:4], boardID: &testBCacheName[8], isAsc: false},
			expectedIdx: 10,
		},
		{
			args:        args{cls: testBCacheTitle[9][:4], boardID: &testBCacheName[9], isAsc: false},
			expectedIdx: 12,
		},
		{
			args:        args{cls: testBCacheTitle[10][:4], boardID: &testBCacheName[10], isAsc: false},
			expectedIdx: 8,
		},
		{
			args:        args{cls: testBCacheTitle[11][:4], boardID: &testBCacheName[11], isAsc: false},
			expectedIdx: 5,
		},
		{
			name:        "title0-brdname0, WhoAMA, desc",
			args:        args{cls: testTitle0, boardID: testBrdname0, isAsc: false},
			expectedIdx: 11,
		},
		{
			name:        "title1-brdname0, between Security/AllHIDPOST, desc",
			args:        args{cls: testTitle1, boardID: testBrdname0, isAsc: false},
			expectedIdx: 4,
		},
		{
			name:        "title2-brdname0, after WhoAmI, desc",
			args:        args{cls: testTitle2, boardID: testBrdname0, isAsc: false},
			expectedIdx: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIdx, err := FindBoardIdxByClass(tt.args.cls, tt.args.boardID, tt.args.isAsc)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindBoardIdxByClass() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotIdx, tt.expectedIdx) {
				t.Errorf("FindBoardIdxByClass() = %v, want %v", gotIdx, tt.expectedIdx)
			}
		})
	}
}

func TestFindBoardAutoCompleteStartIdx(t *testing.T) {
	setupTest()
	defer teardownTest()

	ReloadBCache()

	type args struct {
		keyword []byte
		isAsc   bool
	}
	tests := []struct {
		name             string
		args             args
		expectedStartIdx ptttype.SortIdx
		wantErr          bool
	}{
		// TODO: Add test cases.
		{
			name:             "a, asc (ALLHIDPOST)",
			args:             args{keyword: []byte{'a'}, isAsc: true},
			expectedStartIdx: 3,
		},
		{
			name:             "a, desc (ALLPOST)",
			args:             args{keyword: []byte{'a'}, isAsc: false},
			expectedStartIdx: 4,
		},
		{
			args:             args{keyword: []byte{'b'}, isAsc: true},
			expectedStartIdx: -1,
		},
		{
			args:             args{keyword: []byte{'r'}, isAsc: true},
			expectedStartIdx: 9,
		},
		{
			args:             args{keyword: []byte{'r'}, isAsc: false},
			expectedStartIdx: 9,
		},
		{
			name:             "s, asc (Security)",
			args:             args{keyword: []byte{'s'}, isAsc: true},
			expectedStartIdx: 10,
		},
		{
			name:             "s, desc (SYSOP)",
			args:             args{keyword: []byte{'s'}, isAsc: false},
			expectedStartIdx: 11,
		},
		{
			args:             args{keyword: []byte{'t'}, isAsc: true},
			expectedStartIdx: -1,
		},
		{
			args:             args{keyword: []byte{'w'}, isAsc: true},
			expectedStartIdx: 12,
		},
		{
			args:             args{keyword: []byte{'w'}, isAsc: false},
			expectedStartIdx: 12,
		},
		{
			args:             args{keyword: []byte{'y'}, isAsc: true},
			expectedStartIdx: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStartIdx, err := FindBoardAutoCompleteStartIdx(tt.args.keyword, tt.args.isAsc)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindBoardAutoCompleteStartIdx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStartIdx, tt.expectedStartIdx) {
				t.Errorf("FindBoardAutoCompleteStartIdx() = %v, want %v", gotStartIdx, tt.expectedStartIdx)
			}
		})
	}
}
