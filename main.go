package main

import (
	"log"

	"github.com/skip2/go-qrcode"
)

func main() {
	// file, err := os.Open("file.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer file.Close()

	// readFile, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fileStr := string(readFile)
	fileStr := "https://go.dev/"

	newQRLow := "file-qr-low.png"
	newQRMed := "file-qr-medium.png"
	newQRHigh := "file-qr-high.png"

	err := qrcode.WriteFile(fileStr, qrcode.Low, 512, newQRLow)
	if err != nil {
		log.Fatal(err)
	}

	err = qrcode.WriteFile(fileStr, qrcode.Medium, 512, newQRMed)
	if err != nil {
		log.Fatal(err)
	}

	err = qrcode.WriteFile(fileStr, qrcode.High, 512, newQRHigh)
	if err != nil {
		log.Fatal(err)
	}
}
