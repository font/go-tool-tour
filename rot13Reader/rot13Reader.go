package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {

	n, err = rot.r.Read(p)


	if err != nil {
		if err != io.EOF {
			fmt.Fprintf(os.Stderr, "read error: %v\n", err)
		}
		return n, err
	}

	for i, v := range p[:n] {
		if v >= 'A' && v <= 'Z' {
			// upper case
			if v <= 'M' {
				v += 13
			} else {
				v -= 65 // normalize
				v %= 13
				v += 65 // de-normalize
			}
		} else if v >= 'a' && v <= 'z' {
			// lower case
			if v <= 'm' {
				v += 13
			} else {
				v -= 97 // normalize
				v %= 13
				v += 97 // de-normalize
			}
		} else {
			continue
		}
		p[i] = v
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
