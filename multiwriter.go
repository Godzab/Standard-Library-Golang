package main

import (
	"encoding/json"
	"io"
	"os"
)

type Company struct{
	Name string
	ID int
}

func main(){
	companies := []Company{
		{Name: "Apple", ID: 1},
		{Name: "Google", ID: 2},
		{Name: "Amazon", ID: 3},
		{Name: "Tesla", ID: 4},
	}

	f, err := os.Create("companies.json")

	if err != nil{
		panic(err)
	}
	f2, err := os.Create("companies.bkp")
	if err != nil{
		panic(err)
	}
	defer f.Close()

	// AS MANY WRITERS AS YOU WANT
	mw := io.MultiWriter(f, f2, os.Stdout)

	enc := json.NewEncoder(mw)
	for _, cp := range companies{
		if err := enc.Encode(cp); err != nil{
			panic(err)
		}
	}
}
