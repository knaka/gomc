package gomc

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	RED = "\033[31m"
	END = "\033[0m"
)

func isPrintable(ch byte) bool {
	return 0x20 <= ch && ch < 0x7F
}

func DmMain(args []string) {
	var ab []byte
	if len(args) > 0 {
		var err error
		ab, err = ioutil.ReadFile(args[0])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		ab, _ = ioutil.ReadAll(os.Stdin)
	}
	lineIndex := 0
	lineString := ""
	for i, ch := range ab {
		if lineIndex == 0 {
			_, _ = fmt.Fprintf(os.Stdout, "%08X |", i)
		}
		_, _ = fmt.Fprintf(os.Stdout, " %02X", ch)
		if isPrintable(ch) {
			lineString += string(ch)
		} else {
			lineString += RED + "." + END
		}
		lineIndex++
		if lineIndex == 16 {
			_, _ = fmt.Fprintln(os.Stdout, " |", lineString)
			lineIndex = 0
			lineString = ""
		}
	}
	if lineIndex != 0 {
		for i := 0; i < (16 - lineIndex); i ++ {
			_, _ = fmt.Fprint(os.Stdout, "   ")
		}
		_, _ = fmt.Fprintln(os.Stdout, " |", lineString)
	}
}
