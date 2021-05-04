package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main(){
	//fileReader()
	//stringsReader()
	connReader()
}

func connReader(){
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil{
		panic(err)
	}
	defer conn.Close()
	fmt.Fprint(conn, "GET HTTP 1.0\r\n\r\n")
	readerToStdout(conn, 25)
}

func stringsReader(){
	s := strings.NewReader("We are now using a strings reader here now")
	readerToStdout(s, 1)
}

func fileReader(){
	fl, err := os.Open("file.txt")
	if err != nil{
		panic(err)
	}
	defer fl.Close()

	readerToStdout(fl, 20)
}

func readerToStdout(reader io.Reader, bs int){
	bf := make([]byte, bs)
	for {
		n, err := reader.Read(bf)
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Println(err)
			break
		}
		if n > 0 {
			fmt.Println(string(bf[:n]))
		}
	}
}