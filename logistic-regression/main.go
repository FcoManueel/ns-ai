package main

import (
	"os"
	"log"
	"github.com/kniren/gota/dataframe"
)

const defaultDataSetFilename = "admision.txt"

func main() {
	f, err := os.Open(defaultDataSetFilename)
	if err != nil {
		log.Fatalf("Error while loading data set!!\n Error: %s \n\n", err.Error())
	}
	defer f.Close()

	dataFrame := dataframe.ReadCSV(f)
	_ = dataFrame
}