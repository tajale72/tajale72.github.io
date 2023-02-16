package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"time"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	// Replace "image.jpg" with the path to your image file

	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// client := gosseract.NewClient()
	// defer client.Close()

	// client.SetImage("/Users/romittajale/Documents/tajale72.github.io/tesaract/images/b1.jpeg")
	// text, err := client.Text()
	// if err != nil {
	// 	fmt.Println("Error extracting text from image:", err)
	// } else {
	// 	fmt.Println(text)
	// 	data := []byte(text)
	// 	_, err = file.Write(data)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	time.Sleep(time.Second * 1)

	ReadImageText()
	NewImage()
}

func ReadImageText() {

	// Open the JPEG file
	file, err := os.Open("/Users/romittajale/Documents/tajale72.github.io/improvequality/contrast.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Decode the JPEG data
	img, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}

	// Create a new text file to write the pixel data to
	txtFile, err := os.Create("image_data.txt")
	if err != nil {
		panic(err)
	}
	defer txtFile.Close()

	// Get the size of the image
	bounds := img.Bounds()

	// Loop over each pixel in the image
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Get the color of the pixel at (x, y)
			pixel := img.At(x, y)

			// Get the RGB values of the pixel color
			r, g, b, _ := pixel.RGBA()

			// Write the pixel data to the text file
			fmt.Fprintf(txtFile, "%d,%d,%d\n", r>>8, g>>8, b>>8)
		}
	}
}

func NewImage() {

	// Create a new client for OCR
	client := gosseract.NewClient()

	// Set the path to the business card image file
	client.SetImage("/Users/romittajale/Documents/tajale72.github.io/tesaract/images/test6.jpeg")

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


package main

import (
    "fmt"
    "image"
    _ "image/jpeg" // import JPEG support
    _ "image/png"  // import PNG support
    "os"
)

func main() {
    // set the path to the image file
    imagePath := "path/to/image.png"

    // open the image file and decode it
    file, err := os.Open(imagePath)
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
    defer file.Close()

    img, _, err := image.Decode(file)
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }

    // get the width and height of the image
    bounds := img.Bounds()
    width, height := bounds.Max.X, bounds.Max.Y

    // loop over the pixels in the image and extract their color data
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            r, g, b, a := img.At(x, y).RGBA()
            fmt.Printf("Pixel at (%d,%d) has color RGBA(%d,%d,%d,%d)\n", x, y, r>>8, g>>8, b>>8, a>>8)
        }
    }
}
