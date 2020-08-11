package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"regexp"
)

const red = "\u001b[31m"
const cyan = "\u001b[36;1m"
const reset = "\u001b[0m"

func processText(r io.Reader, w io.Writer) error {

	inBuff := make([]byte, 32768)
	var outBuff bytes.Buffer

	reNumbers := regexp.MustCompile(`\d`)
	reCaps := regexp.MustCompile(`[A-Z]`)

	for {
		n, err := io.ReadFull(r, inBuff)
		if err != nil && err != io.ErrUnexpectedEOF && err != io.EOF {
			return err
		}

		outBuff.Reset()

		partial := reNumbers.ReplaceAll(inBuff[0:n], []byte(red+"$0"+reset))
		full := reCaps.ReplaceAll(partial, []byte(cyan+"$0"+reset))
		w.Write(full)

		if err == io.ErrUnexpectedEOF || err == io.EOF {
			break
		}
	}

	return nil

}

func main() {
	if err := processText(bufio.NewReader(os.Stdin), os.Stdout); err != nil {
		panic(err)
	}
}
