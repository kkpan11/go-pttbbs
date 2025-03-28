package types

import (
	"sync"
)

var (
	origBig5ToUtf8 = ""
	origUtf8ToBig5 = ""

	lock sync.Mutex
)

func SetIsTest(name string) {
	lock.Lock()

	origBig5ToUtf8 = BIG5_TO_UTF8
	origUtf8ToBig5 = UTF8_TO_BIG5

	BIG5_TO_UTF8 = "../docs/etc/uao250-b2u.big5.txt"
	UTF8_TO_BIG5 = "../docs/etc/uao250-u2b.big5.txt"

	_ = initBig5()
}

func UnsetIsTest(name string) {
	defer lock.Unlock()

	BIG5_TO_UTF8 = origBig5ToUtf8
	UTF8_TO_BIG5 = origUtf8ToBig5
}
