package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	imaging "github.com/disintegration/imaging"
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

const TestFilepath string = "images/best-seller-poster/test3.jpeg"

func main() {
	myImage, err := decodeImage(TestFilepath)
	if err != nil {
		fmt.Println("Dis not working")
		fmt.Println(err)
		return
	}
	fmt.Println("decoded the image")

	blurredImage := blurrifyImage(myImage)
	//if err != nil {
	//	fmt.Println("Error blurring image")
	//	fmt.Println(err)
	//	return
	//}
	grayscaledImage := grayscaleImage(blurredImage)
	laplacianImage := convolveImage(grayscaledImage)
	output_path := "images/results/foobar.jpeg"
	outFile, err := os.Create(output_path)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "could not create output file: %v\n", err)
		os.Exit(2)
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, laplacianImage, &jpeg.Options{Quality: 100})
	if err != nil {
		panic(err) // Handle error properly in real code
	}
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

func blurrifyImage(imageToBlur image.Image) *image.NRGBA {
	img := imaging.Blur(imageToBlur, 1.0)

	return img
}

func grayscaleImage(imageToBlur image.Image) *image.NRGBA {
	img := imaging.Grayscale(imageToBlur)

	return img
}

func convolveImage(imageToBlur image.Image) *image.NRGBA {
	img := imaging.Convolve3x3(
		imageToBlur,
		[9]float64{
			0, -1, 0,
			-1, 4, -1,
			0, -1, 0,
		},
		nil,
	)

	return img
}
