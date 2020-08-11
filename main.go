package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"regexp"
)

const black = "\u001b[30m"
const red = "\u001b[31m"
const cyan = "\u001b[36;1m"
const reset = "\u001b[0m"

func main() {
	r := bufio.NewReader(os.Stdin)

	inBuff := make([]byte, 32768)
	var outBuff bytes.Buffer

	reNumbers := regexp.MustCompile(`\d`)
	reCaps := regexp.MustCompile(`[A-Z]`)

	for {
		n, err := io.ReadFull(r, inBuff)
		if err != nil && err != io.ErrUnexpectedEOF && err != io.EOF {
			panic(err)
		}

		outBuff.Reset()

		os.Stdout.Write([]byte(black))
		partial := reNumbers.ReplaceAll(inBuff[0:n], []byte(red+"$0"+reset))
		full := reCaps.ReplaceAll(partial, []byte(cyan+"$0"+reset))
		os.Stdout.Write(full)

		if err == io.ErrUnexpectedEOF || err == io.EOF {
			break
		}
	}

}
