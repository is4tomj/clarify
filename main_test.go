package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {

	var tests = []struct {
		name  string
		input []byte
		want  []byte
	}{
		{
			"lower",
			[]byte("a"),
			[]byte("a"),
		},
		{
			"number",
			[]byte("1"),
			[]byte(red + "1" + reset),
		},
		{
			"number and lower",
			[]byte("1a"),
			[]byte(red + "1" + reset + "a"),
		},
		{
			"lower and number",
			[]byte("a1"),
			[]byte("a" + red + "1" + reset),
		},
		{
			"upper",
			[]byte("A"),
			[]byte(cyan + "A" + reset),
		},
		{
			"upper and lower",
			[]byte("Aa"),
			[]byte(cyan + "A" + reset + "a"),
		},
		{
			"lower and upper",
			[]byte("aA"),
			[]byte("a" + cyan + "A" + reset),
		},
		{
			"lower, number, and upper",
			[]byte("a1A"),
			[]byte("a" + red + "1" + reset + cyan + "A" + reset),
		},
		{
			"number, lower, upper",
			[]byte("2bB"),
			[]byte(red + "2" + reset + "b" + cyan + "B" + reset),
		},
		{
			"upper, number, lower",
			[]byte("C3c"),
			[]byte(cyan + "C" + reset + red + "3" + reset + "c"),
		},
		{
			"complex",
			[]byte("abc123defXYZabc"),
			[]byte("abc" + red + "1" + reset + red + "2" + reset + red + "3" + reset + "def" + cyan + "X" + reset + cyan + "Y" + reset + cyan + "Z" + reset + "abc"),
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			var buff bytes.Buffer

			if err := processText(bytes.NewReader(test.input), &buff); err != nil {
				t.Errorf("processText failed: %v", err)
			}

			output := buff.Bytes()

			if !(bytes.Equal(output, test.want)) {
				t.Errorf("expected bytes of len %d, but got len %d", len(test.want), len(output))
				t.Errorf("expected \"%s\", but got \"%s\"", test.want, output)
			}
		})
	}

}
