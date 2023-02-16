package main

import (
	"image/jpeg"
	"os"

	"github.com/disintegration/imaging"
)

func main() {
	ImproveContrast()
	FiftyPercent()
}

func ImproveContrast() {

	// Open the JPEG file
	file, err := os.Open("/Users/romittajale/Documents/tajale72.github.io/tesaract/images/b1.jpeg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Decode the JPEG data
	img, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}

	// Adjust the contrast of the image
	contrast := imaging.AdjustContrast(img, 30)

	// Save the new image as a new JPEG file
	newFile, err := os.Create("/Users/romittajale/Documents/tajale72.github.io/tesaract/images/contrast.jpeg")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	jpeg.Encode(newFile, contrast, nil)

}

func FiftyPercent() {
	// Open the JPEG file
	file, err := os.Open("/Users/romittajale/Documents/tajale72.github.io/tesaract/images/b1.jpeg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Decode the JPEG data
	img, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}

	// Resize the image to 50% of its original size
	resized := imaging.Resize(img, img.Bounds().Dx()*2, img.Bounds().Dy()*2, imaging.Lanczos)

	// Save the resized image as a new file
	newFile, err := os.Create("/Users/romittajale/Documents/tajale72.github.io/tesaract/images/resized.jpg")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	jpeg.Encode(newFile, resized, nil)
}
