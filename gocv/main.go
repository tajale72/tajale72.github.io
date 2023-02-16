package main

import (
	"fmt"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	// Create a new client for OCR
	client := gosseract.NewClient()

	// Set the path to the business card image file
	client.SetImage("/Users/romittajale/Documents/tajale72.github.io/gocv/images/test4.png")

	// Set the language for OCR (e.g. "eng" for English)
	client.SetLanguage("eng")

	// Preprocess the image
	//client.SetPageSegMode(1)

	//fmt.Println("ocr", ocr)

	// Run OCR on the image
	text, err := client.Text()
	if err != nil {
		panic(err)
	}

	// Print the extracted text from the business card
	fmt.Println(text)

	// Close the client
	client.Close()
}
