package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type NonAlphabetError int

func (e *NonAlphabetError) Error() string {
	return fmt.Sprintf("%d %c isn't an alphabetical charatcer ", *e, *e)
}

type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(b []byte) (int, error) {
	n, err := rr.r.Read(b)
	if err != nil {
		return n, err
	}

	for i := 0; i < n; i++ {
		rotated, _ := rot13(b[i])
		b[i] = rotated
	}
	return n, nil

}

func rot13(b byte) (byte, error) {
	switch {
	case 'a' <= b && b <= 'z':
		return 'a' + (b-'a'+13)%26, nil
	case 'A' <= b && b <= 'Z':
		return 'A' + (b-'A'+13)%26, nil
	default:
		err := NonAlphabetError(b)
		return b, &err
	}
}

func mainRot13Reader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

}
