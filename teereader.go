package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type counter struct{
	total uint64
}

func (c *counter) Write(p []byte) (n int, err error) {
	c.total += uint64(len(p)) //32kb at a time
	progress := float64(c.total) / (1024*1024)
	fmt.Printf("\rDownloading %f MB...", progress)
	return len(p), nil
}

// r: Download a file from https://www.stats.govt.nz/assets/Uploads/Annual-enterprise-survey/Annual-enterprise-survey-2019-financial-year-provisional/Download-data/annual-enterprise-survey-2019-financial-year-provisional-csv.csv
// w: Measure in realtime the number of MB downloaded
// w: Write file into our local filesystem
// w: Write the file into our archive

// r1 : <file from the internet>
// w1 : <progress counter>
// r2 : TeeReader(r1, counter)
// r2.Read()
//   -- TeeReader
//      r1.Read(b)
//      w1.Write(b)
//      return

const fileUrl = "https://www.stats.govt.nz/assets/Uploads/Annual-enterprise-survey/Annual-enterprise-survey-2019-financial-year-provisional/Download-data/annual-enterprise-survey-2019-financial-year-provisional-csv.csv"

func main(){
	res, err := http.Get(fileUrl)
	if err != nil{
		panic(err)
	}
	// Download the file into filesystem
	fl, err := os.OpenFile("annual-enterprise-survey-2019-financial-year-provisional-csv.csv", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil{
		panic(err)
	}
	defer fl.Close()

	if _, err := io.Copy(fl, io.TeeReader(res.Body, &counter{})); err != nil{
		panic(err)
	}

	defer res.Body.Close()
}
