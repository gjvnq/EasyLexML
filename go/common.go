package easyLexML

import (
	"fmt"
	"strconv"
)

const VERSION = "0.4.4"

var Debug bool = false

var SubCounterSymbols []string = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func counter2string(counter, subcounter int) string {
	ans := strconv.Itoa(counter)
	if subcounter > 0 {
		ans += "-"
	}
	max_symb := len(SubCounterSymbols)
	for subcounter > 0 {
		if subcounter > max_symb {
			ans += SubCounterSymbols[0]
			subcounter -= max_symb
		} else {
			ans += SubCounterSymbols[subcounter-1]
			subcounter = 0
		}
	}
	return ans
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Debugln(args ...interface{}) {
	if Debug {
		fmt.Println(args...)
	}
}

func Debugf(format string, args ...interface{}) {
	if Debug {
		fmt.Printf(format, args...)
	}
}
