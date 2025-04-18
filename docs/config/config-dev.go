//go:build dev

package ptttype

const (
	// These are special config variables requiring to be const.
	// copy to ptttype/00-config-production.go to setup production config.
	//
	// type struct requires const.
	MAX_USERS = 50 /* 最高註冊人數 */

	MAX_ACTIVE = 31 /* 最多同時上站人數 */

	MAX_BOARD = 100 /* 最大開板個數 */

	//////////
	//config.h
	//////////
	HASH_BITS = 16 /* userid->uid hashing bits */

	MAX_FRIEND = 100 /* 載入 cache 之最多朋友數目 */

	MAX_REJECT = 32 /* 載入 cache 之最多壞人數目 */

	MAX_MSGS = 10 /* 水球(熱訊)忍耐上限 */

	MAX_ADBANNER = 100 /* 最多動態看板數 */

	HOTBOARDCACHE = 10 /* 熱門看板快取 */

	MAX_FROM = 30 /* 最多故鄉數 */

	MAX_REVIEW = 7 /* 最多水球回顧 */

	NUMVIEWFILE = 14 /* 進站畫面最多數 */

	MAX_ADBANNER_SECTION = 10 /* 最多動態看板類別 */

	MAX_ADBANNER_HEIGHT = 11 /* 最大動態看板內容高度 */

	/////////////////////////////////////////////////////////////////////////////
	// OS Settings 作業系統相關設定

	MAXPATHLEN = 256

	PATHLEN = 256
)
