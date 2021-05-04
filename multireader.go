package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	fl1, err := os.Open("companies.json")
	if err != nil{
		panic(err)
	}
	fl2, err := os.Open("companies.bkp")
	if err != nil{
		panic(err)
	}

	out, err := os.OpenFile("Final.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil{
		panic(err)
	}

	//Declare the multi-Reader here
	mr := io.MultiReader(fl1, fl2)
	mw := io.MultiWriter(os.Stdout, out)
	parseData(mr, mw)
}

func parseData(reader io.Reader, writer io.Writer){
	bf := make([]byte, 100)
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
			nn, _ := fmt.Fprint(writer, string(bf[:n]))
			fmt.Println(nn)
		}
	}
}