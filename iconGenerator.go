package main

import (
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

// icon will be 500 x 500
// boarder of 10
func main() {
	fmt.Println("Icon Generator start")

	hexColor := flag.String("background", "a0a0a0", "Should be a hex string, can include alpha value")
	flag.Parse()

	fmt.Printf("Background color: '%s'\n", *hexColor)

	backgroundColor, err := hexToColor(*hexColor)
	if err != nil {
		// there isn't much I can do here, anyway this is mostly just for me anyway
		panic(err)
	}
	mainColor, err := hexToColor("00add8")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Using color %+v for the background\n", backgroundColor)

	img := initialImage(backgroundColor)
	img = changeImage(10, 10, 100, 100, mainColor, img)

	// write the icon out
	// playing with jpg and comparing against png
	f1, err := os.Create("icon.png")
	f2, err2 := os.Create("icon.jpg")
	if err != nil || err2 != nil {
		// can do nothing but panic, sorry
		panic(err)
	}
	png.Encode(f1, img)
	var opt jpeg.Options
	opt.Quality = 43
	jpeg.Encode(f2, img, &opt)

}

func hexToColor(hexColor string) (color.RGBA, error) {
	nilRGBA := color.RGBA{0, 0, 0, 0}
	rgbaArray, err := hex.DecodeString(hexColor)
	if err != nil {
		return nilRGBA, err
	}

	if len(rgbaArray) != 3 && len(rgbaArray) != 4 {
		return nilRGBA, errors.New("The hex value given was too big")
	}

	if len(rgbaArray) == 3 {
		rgbaArray = append(rgbaArray, 255)
	}
	return color.RGBA{
		R: rgbaArray[0],
		G: rgbaArray[1],
		B: rgbaArray[2],
		A: rgbaArray[3],
	}, nil
}

// initialImage icon for now will be 500 by 500
func initialImage(bg color.RGBA) draw.Image {
	img := image.NewRGBA(image.Rect(0, 0, 500, 500))

	for x := 0; x < 500; x++ {
		for y := 0; y < 500; y++ {
			img.SetRGBA(x, y, bg)
		}
	}

	return img
}

func changeImage(x0, y0, x1, y1 int, color color.RGBA, img draw.Image) draw.Image {
	for x := x0; x < x1; x++ {
		for y := y0; y < y1; y++ {
			img.Set(x, y, color)
		}
	}
	return img
}
