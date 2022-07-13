//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd
// +build darwin dragonfly freebsd linux netbsd openbsd

package englishgen

import (
	"io/ioutil"
	"os"
	"strings"
)

var wordlist []string

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
	}
}

// ref: https://github.com/drhodes/golorem/blob/master/wordlist.go
