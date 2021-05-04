package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const Layout = "2006-01-02"

type User struct{
	Name string 	`json: "name"`
	Surname string  `json: "surname"`
	CreatedAt CustomTime `json:"created_at"`
}

type CustomTime struct{
	time.Time
}

func (c CustomTime) MarshalJSON()([]byte, error){
	return []byte(fmt.Sprintf("\"%s\"", c.Format(Layout))), nil
}

func (c *CustomTime) UnmarshalJSON(v []byte)error{
	var err error
	c.Time, err =  time.Parse(Layout, strings.ReplaceAll(string(v), "\"", "" ))
	return err
}

func main(){
	us := User{
		Name:      "Godfrey",
		Surname:   "Bafana",
		CreatedAt: CustomTime{time.Now()},
	}

	out, err := json.MarshalIndent(us, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	unmarshalExample()

	decodedJson()
}


func unmarshalExample(){
	jsonBytes, err := ioutil.ReadFile("out-single.json")
	if err != nil {
		panic(err)
	}

	u := User{}

	if err := json.Unmarshal(jsonBytes, &u); err !=nil{
		panic(err)
	}
	fmt.Printf("%+v\n", u)
}

func decodedJson(){
	js, err := os.Open("out-single.json")
	if err != nil {
		panic(err)
	}
	defer js.Close()
	dc := json.NewDecoder(js)
	u := &User{}
	if err := dc.Decode(u); err != nil{
		panic(err)
	}
	fmt.Printf("%+v\n", u)
}