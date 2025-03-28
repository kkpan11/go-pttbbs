package ptt

import (
	"reflect"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/cache"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/go-pttbbs/types"
)

func TestLoadGeneralArticles(t *testing.T) {
	setupTest(t.Name())
	defer teardownTest(t.Name())

	cache.ReloadBCache()

	boardID := &ptttype.BoardID_t{}
	copy(boardID[:], []byte("WhoAmI"))
	bid := ptttype.Bid(10)

	boardID1 := &ptttype.BoardID_t{}
	copy(boardID1[:], []byte("SYSOP"))
	bid1 := ptttype.Bid(1)

	type args struct {
		user       *ptttype.UserecRaw
		uid        ptttype.UID
		boardIDRaw *ptttype.BoardID_t
		bid        ptttype.Bid
		startAid   ptttype.SortIdx
		nArticles  int
		isDesc     bool
	}
	tests := []struct {
		name                string
		args                args
		expectedSummaries   []*ptttype.ArticleSummaryRaw
		expectedIsNewest    bool
		expectedNextSummary *ptttype.ArticleSummaryRaw
		expectedStartNumIdx ptttype.SortIdx
		wantErr             bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				user:       testUserecRaw1,
				uid:        1,
				boardIDRaw: boardID,
				bid:        bid,
				startAid:   2,
				nArticles:  1,
				isDesc:     true,
			},
			expectedSummaries:   []*ptttype.ArticleSummaryRaw{testArticleSummary1},
			expectedNextSummary: testArticleSummary0,
			expectedStartNumIdx: 2,
			expectedIsNewest:    true,
		},
		{
			args: args{
				user:       testUserecRaw1,
				uid:        1,
				boardIDRaw: boardID,
				bid:        bid,
				startAid:   2,
				nArticles:  100,
				isDesc:     true,
			},
			expectedSummaries:   []*ptttype.ArticleSummaryRaw{testArticleSummary1, testArticleSummary0},
			expectedStartNumIdx: 2,
			expectedIsNewest:    true,
		},
		{
			args: args{
				user:       testUserecRaw1,
				uid:        1,
				boardIDRaw: boardID,
				bid:        bid,
				startAid:   2,
				nArticles:  2,
				isDesc:     true,
			},
			expectedSummaries:   []*ptttype.ArticleSummaryRaw{testArticleSummary1, testArticleSummary0},
			expectedStartNumIdx: 2,
			expectedIsNewest:    true,
		},
		{
			args: args{
				user:       testUserecRaw1,
				uid:        1,
				boardIDRaw: boardID,
				bid:        bid,
				startAid:   2,
				nArticles:  1,
				isDesc:     true,
			},
			expectedSummaries:   []*ptttype.ArticleSummaryRaw{testArticleSummary1},
			expectedIsNewest:    true,
			expectedStartNumIdx: 2,
			expectedNextSummary: testArticleSummary0,
		},
		{
			args: args{
				user:       testUserecRaw1,
				uid:        1,
				boardIDRaw: boardID,
				bid:        bid,
				startAid:   1,
				nArticles:  1,
				isDesc:     true,
			},
			expectedSummaries:   []*ptttype.ArticleSummaryRaw{testArticleSummary0},
			expectedStartNumIdx: 1,
			expectedIsNewest:    false,
		},
		{
			args: args{
				user:       testUserecRaw1,
				uid:        1,
				boardIDRaw: boardID1,
				bid:        bid1,
				startAid:   2,
				nArticles:  1,
				isDesc:     true,
			},
			expectedSummaries: nil,
			expectedIsNewest:  true,
		},
		{
			args: args{
				user:       testUserecRaw1,
				uid:        1,
				boardIDRaw: boardID,
				bid:        bid,
				startAid:   2,
				nArticles:  1,
				isDesc:     false,
			},
			expectedSummaries:   []*ptttype.ArticleSummaryRaw{testArticleSummary1},
			expectedNextSummary: nil,
			expectedStartNumIdx: 2,
			expectedIsNewest:    true,
		},
		{
			args: args{
				user:       testUserecRaw1,
				uid:        1,
				boardIDRaw: boardID,
				bid:        bid,
				startAid:   1,
				nArticles:  100,
				isDesc:     false,
			},
			expectedSummaries:   []*ptttype.ArticleSummaryRaw{testArticleSummary0, testArticleSummary1},
			expectedStartNumIdx: 1,
			expectedIsNewest:    true,
		},
		{
			args: args{
				user:       testUserecRaw1,
				uid:        1,
				boardIDRaw: boardID,
				bid:        bid,
				startAid:   1,
				nArticles:  2,
				isDesc:     false,
			},
			expectedSummaries:   []*ptttype.ArticleSummaryRaw{testArticleSummary0, testArticleSummary1},
			expectedStartNumIdx: 1,
			expectedIsNewest:    true,
		},
		{
			args: args{
				user:       testUserecRaw1,
				uid:        1,
				boardIDRaw: boardID,
				bid:        bid,
				startAid:   1,
				nArticles:  1,
				isDesc:     false,
			},
			expectedSummaries:   []*ptttype.ArticleSummaryRaw{testArticleSummary0},
			expectedIsNewest:    false,
			expectedStartNumIdx: 1,
			expectedNextSummary: testArticleSummary1,
		},
		{
			args: args{
				user:       testUserecRaw1,
				uid:        1,
				boardIDRaw: boardID,
				bid:        bid,
				startAid:   2,
				nArticles:  1,
				isDesc:     false,
			},
			expectedSummaries:   []*ptttype.ArticleSummaryRaw{testArticleSummary1},
			expectedStartNumIdx: 2,
			expectedIsNewest:    true,
		},
		{
			args: args{
				user:       testUserecRaw1,
				uid:        1,
				boardIDRaw: boardID1,
				bid:        bid1,
				startAid:   1,
				nArticles:  1,
				isDesc:     false,
			},
			expectedSummaries: nil,
			expectedIsNewest:  true,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotSummaries, gotIsNewest, gotNextSummary, gotStartNumIdx, err := LoadGeneralArticles(tt.args.user, tt.args.uid, tt.args.boardIDRaw, tt.args.bid, tt.args.startAid, tt.args.nArticles, tt.args.isDesc)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadGeneralArticles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for _, each := range gotSummaries {
				each.BoardID = nil
			}

			testutil.TDeepEqual(t, "summaries", gotSummaries, tt.expectedSummaries)

			if gotNextSummary != nil {
				gotNextSummary.BoardID = nil
			}
			testutil.TDeepEqual(t, "nextSummary", gotNextSummary, tt.expectedNextSummary)

			if !reflect.DeepEqual(gotIsNewest, tt.expectedIsNewest) {
				t.Errorf("LoadGeneralArticles() isNewest = %v, want %v", gotIsNewest, tt.expectedIsNewest)
			}

			if gotStartNumIdx != tt.expectedStartNumIdx {
				t.Errorf("LoadGeneralArticles() startNumIdx = %v, want %v", gotStartNumIdx, tt.expectedStartNumIdx)
			}
		})
		wg.Wait()
	}
}

func TestLoadGeneralArticlesAllGuest(t *testing.T) {
	setupTest(t.Name())
	defer teardownTest(t.Name())

	boardID := &ptttype.BoardID_t{}
	copy(boardID[:], []byte("WhoAmI"))

	boardID1 := &ptttype.BoardID_t{}
	copy(boardID1[:], []byte("SYSOP"))

	type args struct {
		boardIDRaw *ptttype.BoardID_t
		startIdx   ptttype.SortIdx
		nArticles  int
		isDesc     bool
	}
	tests := []struct {
		name            string
		args            args
		wantSummaries   []*ptttype.ArticleSummaryRaw
		wantIsNewest    bool
		wantNextSummary *ptttype.ArticleSummaryRaw
		wantStartNumIdx ptttype.SortIdx
		wantErr         bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				boardIDRaw: boardID,
				startIdx:   2,
				nArticles:  1,
				isDesc:     true,
			},
			wantSummaries:   []*ptttype.ArticleSummaryRaw{testArticleSummary1},
			wantNextSummary: testArticleSummary0,
			wantStartNumIdx: 2,
			wantIsNewest:    true,
		},
		{
			args: args{
				boardIDRaw: boardID,
				startIdx:   2,
				nArticles:  100,
				isDesc:     true,
			},
			wantSummaries:   []*ptttype.ArticleSummaryRaw{testArticleSummary1, testArticleSummary0},
			wantStartNumIdx: 2,
			wantIsNewest:    true,
		},
		{
			args: args{
				boardIDRaw: boardID,
				startIdx:   2,
				nArticles:  2,
				isDesc:     true,
			},
			wantSummaries:   []*ptttype.ArticleSummaryRaw{testArticleSummary1, testArticleSummary0},
			wantStartNumIdx: 2,
			wantIsNewest:    true,
		},
		{
			args: args{
				boardIDRaw: boardID,
				startIdx:   2,
				nArticles:  1,
				isDesc:     true,
			},
			wantSummaries:   []*ptttype.ArticleSummaryRaw{testArticleSummary1},
			wantIsNewest:    true,
			wantStartNumIdx: 2,
			wantNextSummary: testArticleSummary0,
		},
		{
			args: args{
				boardIDRaw: boardID,
				startIdx:   1,
				nArticles:  1,
				isDesc:     true,
			},
			wantSummaries:   []*ptttype.ArticleSummaryRaw{testArticleSummary0},
			wantStartNumIdx: 1,
			wantIsNewest:    false,
		},
		{
			args: args{
				boardIDRaw: boardID1,
				startIdx:   2,
				nArticles:  1,
				isDesc:     true,
			},
			wantSummaries: nil,
			wantIsNewest:  true,
		},
		{
			args: args{
				boardIDRaw: boardID,
				startIdx:   2,
				nArticles:  1,
				isDesc:     false,
			},
			wantSummaries:   []*ptttype.ArticleSummaryRaw{testArticleSummary1},
			wantNextSummary: nil,
			wantStartNumIdx: 2,
			wantIsNewest:    true,
		},
		{
			args: args{
				boardIDRaw: boardID,
				startIdx:   1,
				nArticles:  100,
				isDesc:     false,
			},
			wantSummaries:   []*ptttype.ArticleSummaryRaw{testArticleSummary0, testArticleSummary1},
			wantStartNumIdx: 1,
			wantIsNewest:    true,
		},
		{
			args: args{
				boardIDRaw: boardID,
				startIdx:   1,
				nArticles:  2,
				isDesc:     false,
			},
			wantSummaries:   []*ptttype.ArticleSummaryRaw{testArticleSummary0, testArticleSummary1},
			wantStartNumIdx: 1,
			wantIsNewest:    true,
		},
		{
			args: args{
				boardIDRaw: boardID,
				startIdx:   1,
				nArticles:  1,
				isDesc:     false,
			},
			wantSummaries:   []*ptttype.ArticleSummaryRaw{testArticleSummary0},
			wantIsNewest:    false,
			wantStartNumIdx: 1,
			wantNextSummary: testArticleSummary1,
		},
		{
			args: args{
				boardIDRaw: boardID,
				startIdx:   2,
				nArticles:  1,
				isDesc:     false,
			},
			wantSummaries:   []*ptttype.ArticleSummaryRaw{testArticleSummary1},
			wantStartNumIdx: 2,
			wantIsNewest:    true,
		},
		{
			args: args{
				boardIDRaw: boardID1,
				nArticles:  1,
				isDesc:     false,
			},
			wantSummaries: nil,
			wantIsNewest:  true,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotSummaries, gotIsNewest, gotNextSummary, gotStartNumIdx, err := LoadGeneralArticlesAllGuest(tt.args.boardIDRaw, tt.args.startIdx, tt.args.nArticles, tt.args.isDesc)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadGeneralArticlesAllGuest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for _, each := range gotSummaries {
				each.BoardID = nil
			}

			if !reflect.DeepEqual(gotSummaries, tt.wantSummaries) {
				t.Errorf("LoadGeneralArticlesAllGuest() gotSummaries = %v, want %v", gotSummaries[0], tt.wantSummaries[0])
			}
			if gotIsNewest != tt.wantIsNewest {
				t.Errorf("LoadGeneralArticlesAllGuest() gotIsNewest = %v, want %v", gotIsNewest, tt.wantIsNewest)
			}

			if gotNextSummary != nil {
				gotNextSummary.BoardID = nil
			}

			if !reflect.DeepEqual(gotNextSummary, tt.wantNextSummary) {
				t.Errorf("LoadGeneralArticlesAllGuest() gotNextSummary = %v, want %v", gotNextSummary, tt.wantNextSummary)
			}
			if !reflect.DeepEqual(gotStartNumIdx, tt.wantStartNumIdx) {
				t.Errorf("LoadGeneralArticlesAllGuest() gotStartNumIdx = %v, want %v", gotStartNumIdx, tt.wantStartNumIdx)
			}
		})
		wg.Wait()
	}
}

func TestFindArticleStartIdx(t *testing.T) {
	setupTest(t.Name())
	defer teardownTest(t.Name())

	cache.ReloadBCache()

	boardID := &ptttype.BoardID_t{}
	copy(boardID[:], []byte("WhoAmI"))
	bid := ptttype.Bid(10)

	/*
		boardID1 := &ptttype.BoardID_t{}
		copy(boardID1[:], []byte("SYSOP"))
		bid1 := ptttype.Bid(1)
	*/

	type args struct {
		user       *ptttype.UserecRaw
		uid        ptttype.UID
		boardID    *ptttype.BoardID_t
		bid        ptttype.Bid
		createTime types.Time4
		filename   *ptttype.Filename_t
		isDesc     bool
	}
	tests := []struct {
		name             string
		args             args
		expectedStartIdx ptttype.SortIdx
		wantErr          bool
	}{
		// TODO: Add test cases.
		{
			name: "most basic: filename: beginning, asc (not isDesc)",
			args: args{
				user:       testUserecRaw1,
				uid:        1,
				boardID:    boardID,
				bid:        bid,
				createTime: 1607202239,
				filename:   &testArticleSummary0.FileHeaderRaw.Filename,
				isDesc:     false,
			},
			expectedStartIdx: 1,
		},
		{
			name: "most basic: filename: last, asc (not isDesc)",
			args: args{
				user:       testUserecRaw1,
				uid:        1,
				boardID:    boardID,
				bid:        bid,
				createTime: 1607203395,
				filename:   &testArticleSummary1.FileHeaderRaw.Filename,
				isDesc:     false,
			},
			expectedStartIdx: 2,
		},
		{
			name: "most basic: filename: beginning, desc",
			args: args{
				user:       testUserecRaw1,
				uid:        1,
				boardID:    boardID,
				bid:        bid,
				createTime: 1607202239,
				filename:   &testArticleSummary0.FileHeaderRaw.Filename,
				isDesc:     true,
			},
			expectedStartIdx: 1,
		},
		{
			name: "most basic: filename: last, desc",
			args: args{
				user:       testUserecRaw1,
				uid:        1,
				boardID:    boardID,
				bid:        bid,
				createTime: 1607203395,
				filename:   &testArticleSummary1.FileHeaderRaw.Filename,
				isDesc:     true,
			},
			expectedStartIdx: 2,
		},
		{
			name: "createTime: beginning, filename: nil, desc",
			args: args{
				user:       testUserecRaw1,
				uid:        1,
				boardID:    boardID,
				bid:        bid,
				createTime: 1607202239,
				filename:   nil,
				isDesc:     true,
			},
			expectedStartIdx: 1,
		},
		{
			name: "createTime: last, filename: nil, desc",
			args: args{
				user:       testUserecRaw1,
				uid:        1,
				boardID:    boardID,
				bid:        bid,
				createTime: 1607203395,
				filename:   nil,
				isDesc:     true,
			},
			expectedStartIdx: 2,
		},
		{
			name: "createTime: beginning, filename: nil, asc (not isDesc)",
			args: args{
				user:       testUserecRaw1,
				uid:        1,
				boardID:    boardID,
				bid:        bid,
				createTime: 1607202239,
				filename:   nil,
				isDesc:     false,
			},
			expectedStartIdx: 1,
		},
		{
			name: "createTime: last, filename: nil, asc (not isDesc)",
			args: args{
				user:       testUserecRaw1,
				uid:        1,
				boardID:    boardID,
				bid:        bid,
				createTime: 1607203395,
				filename:   nil,
				isDesc:     false,
			},
			expectedStartIdx: 2,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotStartAid, err := FindArticleStartIdx(tt.args.user, tt.args.uid, tt.args.boardID, tt.args.bid, tt.args.createTime, tt.args.filename, tt.args.isDesc)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindArticleStartAid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStartAid, tt.expectedStartIdx) {
				t.Errorf("FindArticleStartAid() = %v, want %v", gotStartAid, tt.expectedStartIdx)
			}
		})
		wg.Wait()
	}
}

func TestFindArticleStartIdxAllGuest(t *testing.T) {
	setupTest(t.Name())
	defer teardownTest(t.Name())

	boardID := &ptttype.BoardID_t{}
	copy(boardID[:], []byte("WhoAmI"))

	type args struct {
		boardID    *ptttype.BoardID_t
		createTime types.Time4
		filename   *ptttype.Filename_t
		isDesc     bool
	}
	tests := []struct {
		name         string
		args         args
		wantStartIdx ptttype.SortIdx
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name: "most basic: filename: beginning, asc (not isDesc)",
			args: args{
				boardID:    boardID,
				createTime: 1607202239,
				filename:   &testArticleSummary0.FileHeaderRaw.Filename,
				isDesc:     false,
			},
			wantStartIdx: 1,
		},
		{
			name: "most basic: filename: last, asc (not isDesc)",
			args: args{
				boardID:    boardID,
				createTime: 1607203395,
				filename:   &testArticleSummary1.FileHeaderRaw.Filename,
				isDesc:     false,
			},
			wantStartIdx: 2,
		},
		{
			name: "most basic: filename: beginning, desc",
			args: args{
				boardID:    boardID,
				createTime: 1607202239,
				filename:   &testArticleSummary0.FileHeaderRaw.Filename,
				isDesc:     true,
			},
			wantStartIdx: 1,
		},
		{
			name: "most basic: filename: last, desc",
			args: args{
				boardID:    boardID,
				createTime: 1607203395,
				filename:   &testArticleSummary1.FileHeaderRaw.Filename,
				isDesc:     true,
			},
			wantStartIdx: 2,
		},
		{
			name: "createTime: beginning, filename: nil, desc",
			args: args{
				boardID:    boardID,
				createTime: 1607202239,
				filename:   nil,
				isDesc:     true,
			},
			wantStartIdx: 1,
		},
		{
			name: "createTime: last, filename: nil, desc",
			args: args{
				boardID:    boardID,
				createTime: 1607203395,
				filename:   nil,
				isDesc:     true,
			},
			wantStartIdx: 2,
		},
		{
			name: "createTime: beginning, filename: nil, asc (not isDesc)",
			args: args{
				boardID:    boardID,
				createTime: 1607202239,
				filename:   nil,
				isDesc:     false,
			},
			wantStartIdx: 1,
		},
		{
			name: "createTime: last, filename: nil, asc (not isDesc)",
			args: args{
				boardID:    boardID,
				createTime: 1607203395,
				filename:   nil,
				isDesc:     false,
			},
			wantStartIdx: 2,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotStartIdx, err := FindArticleStartIdxAllGuest(tt.args.boardID, tt.args.createTime, tt.args.filename, tt.args.isDesc)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindArticleStartIdxAllGuest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStartIdx, tt.wantStartIdx) {
				t.Errorf("FindArticleStartIdxAllGuest() = %v, want %v", gotStartIdx, tt.wantStartIdx)
			}
		})
		wg.Wait()
	}
}

func TestLoadGeneralArticlesSameCreateTime(t *testing.T) {
	setupTest(t.Name())
	defer teardownTest(t.Name())

	cache.ReloadBCache()

	boardID := &ptttype.BoardID_t{}
	copy(boardID[:], []byte("WhoAmI"))
	bid := ptttype.Bid(10)

	type args struct {
		boardIDRaw *ptttype.BoardID_t
		bid        ptttype.Bid
		startAid   ptttype.SortIdx
		endAid     ptttype.SortIdx
		createTime types.Time4
	}
	tests := []struct {
		name                string
		args                args
		expectedSummaries   []*ptttype.ArticleSummaryRaw
		expectedStartNumIdx ptttype.SortIdx
		expectedEndNumIdx   ptttype.SortIdx
		wantErr             bool
	}{
		// TODO: Add test cases.
		{
			args:                args{boardIDRaw: boardID, bid: bid, startAid: 1, endAid: 0, createTime: 1607202239},
			expectedSummaries:   []*ptttype.ArticleSummaryRaw{testArticleSummary0},
			expectedStartNumIdx: 1,
			expectedEndNumIdx:   1,
		},
		{
			args:                args{boardIDRaw: boardID, bid: bid, startAid: 1, endAid: 0, createTime: 1607203395},
			expectedSummaries:   []*ptttype.ArticleSummaryRaw{testArticleSummary1},
			expectedStartNumIdx: 2,
			expectedEndNumIdx:   2,
		},
		{
			args:              args{boardIDRaw: boardID, bid: bid, startAid: 1, endAid: 0, createTime: 1607203394},
			expectedSummaries: nil,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotSummaries, gotStartNumIdx, gotEndNumIdx, err := LoadGeneralArticlesSameCreateTime(tt.args.boardIDRaw, tt.args.bid, tt.args.startAid, tt.args.endAid, tt.args.createTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadGeneralArticlesSameCreateTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, each := range gotSummaries {
				each.BoardID = nil
			}
			testutil.TDeepEqual(t, "summaries", gotSummaries, tt.expectedSummaries)

			if gotStartNumIdx != tt.expectedStartNumIdx {
				t.Errorf("LoadGeneralArticlesSameCreateTime() startNumIdx = %v, want %v", gotStartNumIdx, tt.expectedStartNumIdx)
			}

			if gotEndNumIdx != tt.expectedEndNumIdx {
				t.Errorf("LoadGeneralArticlesSameCreateTime() endNumIdx = %v, want %v", gotEndNumIdx, tt.expectedEndNumIdx)
			}
		})
		wg.Wait()
	}
}

func TestLoadBottomArticles(t *testing.T) {
	setupTest(t.Name())
	defer teardownTest(t.Name())

	cache.ReloadBCache()

	boardID := &ptttype.BoardID_t{}
	copy(boardID[:], []byte("WhoAmI"))
	bid := ptttype.Bid(10)

	boardID1 := &ptttype.BoardID_t{}
	copy(boardID1[:], []byte("SYSOP"))
	bid1 := ptttype.Bid(1)

	type args struct {
		user       *ptttype.UserecRaw
		uid        ptttype.UID
		boardIDRaw *ptttype.BoardID_t
		bid        ptttype.Bid
	}
	tests := []struct {
		name              string
		args              args
		expectedSummaries []*ptttype.ArticleSummaryRaw
		wantErr           bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				user:       testUserecRaw1,
				uid:        1,
				boardIDRaw: boardID,
				bid:        bid,
			},
			expectedSummaries: []*ptttype.ArticleSummaryRaw{testBottomSummary1},
		},
		{
			args: args{
				user:       testUserecRaw1,
				uid:        1,
				boardIDRaw: boardID1,
				bid:        bid1,
			},
			expectedSummaries: nil,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotSummaries, err := LoadBottomArticles(tt.args.user, tt.args.uid, tt.args.boardIDRaw, tt.args.bid)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadBottomArticles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for _, each := range gotSummaries {
				each.BoardID = nil
				copy(each.Multi[:], []byte{0, 0, 0, 0})
			}

			testutil.TDeepEqual(t, "summaries", gotSummaries, tt.expectedSummaries)
		})
		wg.Wait()
	}
}

func TestLoadBottomArticlesAllGuest(t *testing.T) {
	setupTest(t.Name())
	defer teardownTest(t.Name())

	boardID := &ptttype.BoardID_t{}
	copy(boardID[:], []byte("WhoAmI"))

	boardID1 := &ptttype.BoardID_t{}
	copy(boardID1[:], []byte("SYSOP"))

	type args struct {
		boardIDRaw *ptttype.BoardID_t
	}
	tests := []struct {
		name          string
		args          args
		wantSummaries []*ptttype.ArticleSummaryRaw
		wantErr       bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				boardIDRaw: boardID,
			},
			wantSummaries: []*ptttype.ArticleSummaryRaw{testBottomSummary1},
		},
		{
			args: args{
				boardIDRaw: boardID1,
			},
			wantSummaries: nil,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotSummaries, err := LoadBottomArticlesAllGuest(tt.args.boardIDRaw)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadBottomArticlesAllGuest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for _, each := range gotSummaries {
				each.BoardID = nil
				copy(each.Multi[:], []byte{0, 0, 0, 0})
			}

			if !reflect.DeepEqual(gotSummaries, tt.wantSummaries) {
				t.Errorf("LoadBottomArticlesAllGuest() = %v, want %v", gotSummaries, tt.wantSummaries)
			}
		})
		wg.Wait()
	}
}
