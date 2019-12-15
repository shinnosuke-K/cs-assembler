package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	corresTable "github.com/shinnosuke-K/cs-assembler/corres-table"
	symbolTable "github.com/shinnosuke-K/cs-assembler/symbol-table"
)

// First Pass
// Add the label symbols
func firstPass(file io.Reader) {

	var text string

	scanFile := bufio.NewScanner(file)
	for lineNum := 0; scanFile.Scan(); lineNum++ {
		text = strings.Replace(scanFile.Text(), " ", "", -1)
		if len(text) == 0 {
			continue
		}

		if string(text[0]) == "(" {
			symbolTable.InputSymbol(text[1:len(text)-1], lineNum)
		}
	}
	secondPass(file)
}

// Second Pass
// Add the var symbols
func secondPass(file io.Reader) {

	var val = 16
	scanFile := bufio.NewScanner(file)

	for scanFile.Scan() {
		text := strings.Replace(scanFile.Text(), " ", "", -1)
		if len(text) == 0 {
			continue
		}

		_, err := strconv.Atoi(text[1:])

		if string(text[0]) == "@" && err != nil {
			if _, ok := symbolTable.GetSymbolValue(text[2:]); !ok {
				symbolTable.InputSymbol(text[1:], val)
				val++
			}
		}
	}
}

func processString(text string) string {
	req := regexp.MustCompile(`\s{3,}`)
	processText := req.ReplaceAllString(text, "")

	req = regexp.MustCompile(`//.*`)
	return req.ReplaceAllString(processText, "")
}

func writeBinary(file io.Reader) {
	scanFile := bufio.NewScanner(file)

	for scanFile.Scan() {

		text := processString(scanFile.Text())

		if len(text) == 0 {
			continue
		}

		out := "111"

		if string(text[0]) == "@" {
			if v, err := strconv.Atoi(text[1:]); err != nil {
				value, _ := symbolTable.GetSymbolValue(text[1:])
				fmt.Println(encodeBinary(value))

			} else {
				fmt.Println(encodeBinary(v))
			}
			fmt.Println()
			continue
		}

		var comp, dest, jump string

		switch string(text[1]) {
		case "=":
			dest = corresTable.GetDestBinary(string(text[0]))

			if index := strings.Index(text, ";"); index != -1 && len(text[index+1:]) > 0 {
				comp = corresTable.GetCompBinary(text[2:index])
				jump = corresTable.GetJumpBinary(text[index+1:])
			} else {
				comp = corresTable.GetCompBinary(text[2:])
				jump = corresTable.GetJumpBinary("null")
			}
			out += comp + dest + jump

		case ";":
			comp = corresTable.GetCompBinary(string(text[0]))
			dest = corresTable.GetDestBinary("null")

			if len(text[1:]) > 0 {
				jump = corresTable.GetJumpBinary(text[2:])
			} else {
				jump = corresTable.GetDestBinary("null")
			}

			out += comp + dest + jump

		}
		fmt.Println(out)
		fmt.Println()
	}
}

func encodeBinary(num int) string {
	bin := ""
	for n := 0; n < 16; n++ {
		bin = strconv.Itoa(num&1) + bin
		num = num >> 1
	}
	return bin
}

func main() {

	passFile, err := os.Open(os.Args[1])
	defer passFile.Close()
	if err != nil {
		panic(err)
	}
	firstPass(passFile)

	encodeFile, err := os.Open(os.Args[1])
	defer encodeFile.Close()
	if err != nil {
		panic(err)
	}
	writeBinary(encodeFile)
}
