package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
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

// const TestFilepath string = "images/best-seller-poster/test3.jpeg"
// const TestFilepath = "images/best-seller-poster/test3.jpeg"
// const TestFilepath = "images/best-seller-poster/test4.jpg"
const TestFilepath = "images/best-seller-poster/52897657.png"

// const TestFilepath = "images/best-seller-poster/71891485.jpg" // be kind snow poster
// const TestFilepath = "images/best-seller-poster/162821896.png" // be kind snow poster
// const TestFilepath = "images/best-seller-poster/49043302.jpg" // be kind snow poster
const Sigma = 1.3

//const Sigma = 0.3

//var kernel3x3 = [9]float64{0, 1, 0, 1, -4, 1, 0, 1, 0}

var kernel3x3 = [9]float64{1, 4, 1, 4, -20, 4, 1, 4, 1}

func main() {
	myImage, err := decodeImage(TestFilepath)
	if err != nil {
		fmt.Println("Dis not working")
		fmt.Println(err)
		return
	}
	fmt.Println("decoded the image")

	croppedImage := cropImage(myImage)
	blurredImage := blurrifyImage(croppedImage)
	grayscaledImage := grayscaleImage(blurredImage)
	laplacianImage := convolveImage(grayscaledImage)
	fmt.Println(laplacianImage.At(40, 41))
	output_path := "images/results/foobar3.png"
	outFile, err := os.Create(output_path)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "could not create output file: %v\n", err)
		os.Exit(2)
	}
	defer outFile.Close()

	//err = jpeg.Encode(outFile, laplacianImage, &jpeg.Options{Quality: 100})
	//if err != nil {
	//	panic(err) // Handle error properly in real code
	//}
	err = png.Encode(outFile, laplacianImage)
	if err != nil {
		panic(err) // Handle error properly in real code
	}
}

//func calculateMean(imageData *image.NRGBA) (float64){
//	imageData.reduce
//}
//
//func calculateVariance(imageData *image.NRGBA) (float64) {
//
//}

//const mean = laplacianImageData.reduce((sum, value) => sum + value, 0) / laplacianImageData.length;
//const variance = laplacianImageData.reduce((sum, value) => sum + Math.pow(value - mean, 2), 0) / laplacianImageData.length;

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
		return loadedImage, nil
	}

	if imageType == "png" {
		loadedImage, err := png.Decode(existingImageFile)
		if err != nil {
			fmt.Println(err)
		}
		return loadedImage, nil
	}

	return nil, fmt.Errorf("image type not supported")
}

func blurrifyImage(imageToBlur image.Image) *image.NRGBA {
	img := imaging.Blur(imageToBlur, Sigma)

	return img
}

func grayscaleImage(imageToBlur image.Image) *image.NRGBA {
	img := imaging.Grayscale(imageToBlur)

	return img
}

func cropImage(imageToCrop image.Image) *image.NRGBA {
	img := imaging.CropCenter(imageToCrop, 1000, 1000)

	return img
}

func convolveImage(imageToBlur image.Image) *image.NRGBA {
	img := imaging.Convolve3x3(
		imageToBlur,
		kernel3x3,
		nil,
	)

	return img
}
