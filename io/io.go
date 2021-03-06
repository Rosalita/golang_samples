package main

import (
	"encoding/json"
	"bytes"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type readingThing struct {
}

func (rt readingThing) Read(stuff []byte)(n int, err error){
	return len(stuff), nil
}

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

	// Go supports buffered IO via package bufio which makes it easy to work with text.
	// bufio.NewReader returns a new reader with default buffer size
	bufioReader := bufio.NewReader(r1)
	line, _ := bufioReader.ReadString('\n') // can read line by line
	fmt.Println(line)

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

	r6 := bytes.NewReader([]byte{20,21,22})
	// r6 is a *Reader
	reader = r6
	// *Reader satisfies the io.Reader interface because it has a .Read() method

	///////////////////
	// Using Readers
	///////////////////

	var someReader io.Reader
	someReader = readingThing{} // reading thing has a method .Read() that returns number of bytes read and an error

	// Probably the least useful thing is to read from a reader directly
	p := []byte("this is some bytes")
	numberOfBytesRead, err := someReader.Read(p)
	// this returns how many bytes were read and any error
	fmt.Println(numberOfBytesRead, err)

	// can get all the raw []byte data out of reader using ioutil.ReadAll
	reader = strings.NewReader("Hello")
	b, err := ioutil.ReadAll(reader)
	fmt.Println(b, string(b), err)

	reader = bytes.NewReader([]byte{20,21,22})
	b, err = ioutil.ReadAll(reader)
	fmt.Println(b, string(b), err)

	reader, _ = os.Open("file.txt")
	b, err = ioutil.ReadAll(reader)
	fmt.Println(b, string(b), err)

	// However beware using ioutil.ReadAll on files if they are large, or have the potential to become large
	// ioutil.ReadAll loads the whole file into memory, which can become bad if multiple users do this
	// eventually crashing due to lack of memory
	// When you have a reader and are going to write to to a io.Writer, prefer io.Copy over ioutil.ReadAll

	// io.Copy reads ALL bytes from an io.Reader, and writes it to an io.Writer
	buffer := bytes.NewBuffer([]byte("something already in the buffer..."))
	// buffer is a *Buffer which has .Read() and .Write() methods
	file, err := os.Open("file.txt")
	io.Copy(buffer, file)
	fmt.Println(buffer)

	// If building a package or utility (even if it’s an internal thing that nobody will ever see) 
	// rather than taking in strings or []byte slices, consider taking in an io.Reader for data sources. 
	// By doing so, your code will work with every type that implements io.Reader.

	
    // Can get data as a []byte from a file using ioutil.ReadFile
	jsonStream, _ := ioutil.ReadFile("example.json")

	//can create a new json decoder from an io.Reader
	// json.NewDecoder takes an io.Reader as an input and returns a *Decoder
	decoder := json.NewDecoder(bytes.NewReader(jsonStream))

	foo := map[string]float64{}
	decoder.Decode(&foo) // can decode without a struct to unmarshal into
	fmt.Println(foo)

	reader = bytes.NewReader(jsonStream)

	doStuffWithReader(reader)

}

func doStuffWithReader(r io.Reader){

	decoder := json.NewDecoder(r)
	fmt.Println(decoder)

}
