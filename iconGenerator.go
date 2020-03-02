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
	"math/rand"
	"os"
	"time"
)

// default icon will be 740 x 740
// boarder of 10
// 740 because 720 is a highly composie number
func main() {
	fmt.Println("Icon Generator start")
	var iconSize int
	var sections int

	hexColor := flag.String("background", "707070", "Should be a hex string, can include alpha value")
	flag.IntVar(&iconSize, "iconSize", 720, "size of the icon, icons generated will be square")
	flag.IntVar(&sections, "sections", 6, "Number of sections per side")
	flag.Parse()
	stepSize := (iconSize - 20) / sections

	rand.Seed(time.Now().Unix())

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

	img := initialImage(iconSize, backgroundColor)
	for x := 10; x < iconSize-10-stepSize; x += stepSize {
		for y := 10; y < iconSize-10-stepSize; y += stepSize {
			if rand.Int()%2 == 0 {
				img = changeImage(x, y, x+stepSize, y+stepSize, mainColor, img)
			}
		}
	}

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
	opt.Quality = 17
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
func initialImage(iconSize int, bg color.RGBA) draw.Image {
	img := image.NewRGBA(image.Rect(0, 0, iconSize, iconSize))

	for x := 0; x < iconSize; x++ {
		for y := 0; y < iconSize; y++ {
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
