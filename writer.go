package main

import (
	"bytes"
	"encoding/json"
	"os"
)

type user struct{
	Name string `json:"name"`
	Age int `json:"age"`
}



func main(){
	fl, err := os.OpenFile("file.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil{
		panic(err)
	}
	defer fl.Close()
	fl.Write([]byte("This is a test\n"))

	// json.encoder writer
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	u := user{Name: "Godfrey", Age: 30}

	if err := enc.Encode(u); err != nil{
		panic(err)
	}
	// {"name": "Godfrey", "age": 30}
	fl.Write(buf.Bytes())
}