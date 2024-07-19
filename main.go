package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

// PSEUDOCODE
// what is image data ✅
// extract image data in go ✅
// add the package "blurry"
// figure out what image data "blurry" requires
// add gaussian-blur to image
// make image greyscale
// apply laplacian matrix
// convolute and return laplacian variance

const TestFilepath string = "images/best-seller-poster/13060191.jpg"

func main() {
	myImage, err := decodeImage(TestFilepath)
	if err != nil {
		fmt.Println("Dis not working")
		fmt.Println(err)
		return
	}
	fmt.Println("decoded the image")
	fmt.Println(myImage)
}

func decodeImage(filePath string) (image.Image, error) {
	existingImageFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening files")
	}
	defer existingImageFile.Close()

	_, imageType, err := image.Decode(existingImageFile)
	if err != nil {
		fmt.Println("Error decoding image")
		fmt.Println(err)
	}

	// We only need this because we already read from the file
	// We have to reset the file pointer back to beginning
	existingImageFile.Seek(0, 0)

	if imageType == "jpeg" {
		loadedImage, err := jpeg.Decode(existingImageFile)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(loadedImage)
		return loadedImage, nil
	}

	return nil, fmt.Errorf("image type not supported")
}
