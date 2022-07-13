//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd
// +build darwin dragonfly freebsd linux netbsd openbsd

package englishgen

import (
	"github.com/puzpuzpuz/xsync"
	"github.com/samber/lo"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var wordlist []string

var wordMapByLen = xsync.NewMapOf[[]string]()
var wordLenMax = 0

func init() {
	if wordlist == nil {
		file, err := os.Open("/usr/share/dict/words")
		if err != nil {
			panic(err)
		}

		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}

		wordlist = strings.Split(string(bytes), "\n")

		wm := lo.GroupBy[string, int](wordlist, func(word string) int {
			wordLen := len(word)

			if wordLen > wordLenMax {
				wordLenMax = wordLen
			}

			return wordLen
		})

		for k, v := range wm {
			wordMapByLen.Store(strconv.Itoa(k), v)
		}

		for i := 1; i <= wordLenMax; i++ {
			lenStr := strconv.Itoa(i)

			if _, ok := wordMapByLen.Load(lenStr); !ok {
				panic(`no word with its length is ` + lenStr)
			}
		}
	}
}

// ref: https://github.com/drhodes/golorem/blob/master/wordlist.go
