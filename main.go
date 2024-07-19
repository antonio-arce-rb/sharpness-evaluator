package main

import (
	"fmt"
	"image"
	"os"
)

// PSEUDOCODE
// what is image data
// extract image data in go
// add the package "blurry"
// figure out what image data "blurry" requires
// add gaussian-blur to image
// make image greyscale
// apply laplacian matrix
// convolute and return laplacian variance

const TestFilepath = "images/best-seller-poster/13060191.jpg"

func main() {
	_, err := decodeImage(TestFilepath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("decoded the image")
}

func decodeImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}
