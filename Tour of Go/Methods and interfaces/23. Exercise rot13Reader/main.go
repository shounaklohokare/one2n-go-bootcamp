package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(x byte) byte {
	switch {
	// If the letter is in the first half of the alphabet (A-M or a-m) shift it forward by 13
	case x >= 65 && x <= 77:
		fallthrough
	case x >= 97 && x <= 109:
		x = x + 13
	//	If the letter is in the second half (N-Z or n-z) shift it backward by 13
	case x >= 78 && x <= 90:
		fallthrough
	case x >= 110 && x <= 122:
		x = x - 13
	}
	return x
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	for i := 0; i <= n; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
