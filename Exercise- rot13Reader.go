package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (a *rot13Reader) Read(b []byte) (int, error){
	
	n, err := a.r.Read(b)

	for x:=range b {
		if (b[x]>='a' && b[x]<='m') || (b[x]>='A' && b[x]<='M')  {
		b[x] = b[x] + 13
		}else if (b[x]>='n' && b[x]<='z') || (b[x]>='N' && b[x]<='Z')  {
		b[x] = b[x] - 13
		}
	}
	
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
