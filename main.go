package main

import (
	"flag"
	"log"
	"os"
	//"fmt"
)

const (
	API_URL = "https://api.vk.com/method/"
)

var (
	token     string
	outputDir string
)

func main() {

	var access_token, output string
	flag.StringVar(&access_token, "ac", "", "access_token")
	flag.StringVar(&output, "o", "", "output dir")
	flag.Parse()

	if access_token != "" {
		token = access_token

		if output != "" {
			outputDir = output
		} else {
			outputDir = "output"
		}

		if _, err := os.Stat(outputDir); os.IsNotExist(err) {
			err := createDir(outputDir)

			if err != nil {
				log.Fatal(err)
			}
		}



		//if list, err := getAllPhotos(); err != nil {
		//	fmt.Errorf("%s", err)
		//} else {
		//	fmt.Println(len(list))
		//}

		importPhotos()

	} else {
		log.Fatalf("%s", "Param -ac needed, access token required")
	}
}
