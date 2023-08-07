package main

import (
	"bufio"
	"os"
	"strings"
)

func ReadFirstLine() string {
	open, err := os.Open("log")
	//defer func(open *os.File) {
	//	err := open.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(open)
	defer func(open *os.File) {
		err := open.Close()
		if err != nil {

		}
	}(open)
	if err != nil {
		return ""
	}

	scanner := bufio.NewScanner(open)
	for scanner.Scan() {
		return scanner.Text()
	}
	return ""

}

func ProcessFirstLine() string {
	line := ReadFirstLine()
	destLine := strings.ReplaceAll(line, "11", "00")
	return destLine
}

func main() {
	ProcessFirstLine()
}
