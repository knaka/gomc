package gomc

import (
	"fmt"
	"io/ioutil"
	"log"
)

const (
	RED = "\033[31m"
	END = "\033[0m"
)

func DmMain(args []string) {
	for _, arg := range args {
		var ab, err = ioutil.ReadFile(arg)
		if err != nil {
			log.Fatal(err)
		}
		count := 0
		s := ""
		for i, ch := range ab {
			if count == 0 {
				print(fmt.Sprintf("%08X |", i))
			}
			print(fmt.Sprintf(" %02X", ch))
			if 0x20 <= ch && ch < 0x7F {
				s += string(ch)
			} else {
				s += RED + "." + END
			}
			count ++
			if count == 16 {
				println(" |", s)
				count = 0
				s = ""
			}
		}
		if count != 0 {
			for i := 0; i < 16 - count; i ++ {
				print("   ")
			}
			println(" |", s)
		}
	}
}
