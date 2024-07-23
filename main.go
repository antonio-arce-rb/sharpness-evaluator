package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"math"
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
// const TestFilepath = "images/best-seller-poster/52897657.png"

// const TestFilepath = "images/best-seller-poster/71891485.jpg" // be kind snow poster
// const TestFilepath = "images/best-seller-poster/162821896.png" // be kind snow poster
// const TestFilepath = "images/best-seller-poster/49043302.jpg" // be kind snow poster
// const TestFilepath = "images/best-seller-poster/13060191.jpg"
// const TestFilepath = "images/with-defects/104679315.jpg"
const Sigma = 1.0

//const Sigma = 0.3

var kernel3x3 = [9]float64{0, 1, 0, 1, -4, 1, 0, 1, 0}

// var kernel3x3 = [9]float64{1, 4, 1, 4, -20, 4, 1, 4, 1}

func main() {
	// items, _ := os.ReadDir("images/with-defects")
	theFiles := []string{"108644963.png", "104679315.jpg", "150966500.jpg", "143703300.jpg", "160918245.jpeg", "109072446.jpg", "157928170.jpg", "159667086.jpg", "104144236.jpg", "141598070.jpg", "161327344.jpg", "128473276.jpg", "122032158.jpg", "106593432.jpg", "81454929.jpg", "74825964.jpg", "144139218.jpg", "90756016.jpg", "111652919.jpg", "42464649.jpg", "145734216.jpg", "74783467.jpg", "148063673.jpg", "48216883.jpeg", "41935294.jpg"}

	for _, item := range theFiles {
		// fmt.Print(item + ": ")

		doTheThing("images/with-defects/" + item)
	}
	// doTheThing("images/with-defects/104679315.jpg")
}

func doTheThing(filePath string) {
	myImage, err := decodeImage(filePath)
	if err != nil {
		fmt.Println("Dis not working")
		fmt.Println(err)
		return
	}
	// fmt.Println("decoded the image")

	croppedImage := cropImage(myImage)
	blurredImage := blurrifyImage(croppedImage)
	grayscaledImage := grayscaleImage(blurredImage)
	laplacianImage := convolveImage(grayscaledImage)
	// fmt.Println(laplacianImage.At(40, 41))
	// fmt.Println(rgba.Pix(laplacianImage))
	// saveImage(laplacianImage)
	arrayOfPixels := buildArray(laplacianImage)
	mean := calculateMean(arrayOfPixels)
	variance := calculateVariance(arrayOfPixels, mean)

	fmt.Println((variance))
}

func saveImage(laplacianImage *image.NRGBA) {
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

func calculateMean(imageData []uint8) float64 {
	sum := 0
	for i := 0; i < len(imageData); i++ {
		sum += int(imageData[i])
	}

	// fmt.Println(sum)
	// fmt.Println(float32(sum))
	// formattedValueWithPrecision := fmt.Sprintf("%.2f", float64(sum))

	// fmt.Println("Formatted float with precision:", formattedValueWithPrecision)

	return float64(sum) / float64(len(imageData))
}

func calculateVariance(imageData []uint8, mean float64) float64 {
	sum := 0.0
	for i := 0; i < len(imageData); i++ {
		value := float64(imageData[i]) - mean
		sum += float64(imageData[i]) + math.Pow(value, 2)
	}

	return sum / float64(len(imageData))
}

//const variance = laplacianImageData.reduce((sum, value) => sum + Math.pow(value - mean, 2), 0) / laplacianImageData.length;

func buildArray(laplacianImage *image.NRGBA) []uint8 {
	bounds := laplacianImage.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	myArray := []uint8{}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, _, _, _ := laplacianImage.At(x, y).RGBA()
			c := uint8(r >> 8)
			myArray = append(myArray, c)
		}
	}

	return myArray
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
