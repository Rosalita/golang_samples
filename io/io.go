package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"bytes"
)

func main() {

	// reader is an io.Reader
	var reader io.Reader

	fmt.Println(reader) // nil

	//  io.Reader is an interface with a method Read
	//  type Reader interface {
	//	  Read(p []byte) (n int, err error)
	//  }
	//
	// it returns an int which is the number of bytes read
	// and an error which is any error encountered

	////////////////////////
	// Examples of readers
	////////////////////////

	r1, _ := os.Open("file.txt")
	// r1 is a *file
	reader = r1
	// *file satisfies the io.Reader interface because it has a .Read() method

	r2 := strings.NewReader("This reader will return this string as []bytes")
	// r2 is a *Reader
	reader = r2
	// *Reader satisfies the io.Reader interface because it has a .Read() method

	r3, _ := http.Get("http://google.com/")
	// r3 is a *Response, which does not satisfy io.Reader
	reader = r3.Body // However, r3.Body is an io.ReadCloser
	// r3.Body satisfies the io.Reader interface

	r4, _ := http.NewRequest("GET", "http://example.com", nil)
	// r4 is a *Request, which does not satisfy io.Reader
	reader = r4.Body // However, r4.Body is an io.ReadCloser
	// r4.Body satifies the io.Reader interface

	r5 := bytes.Buffer{} // bytes.Buffer has .Read() and .Write() methods
	reader = &r5
	// r5 satisfies the io.Reader interface

}
