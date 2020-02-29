package main

import (
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"image/color"
)

func main() {
	fmt.Println("Icon Generator start")

	hexColor := flag.String("background", "f0f0f0", "Should be a hex string, can include alpha value")
	flag.Parse()

	fmt.Printf("Background color: '%s'\n", *hexColor)

	backgroundColor, err := rgbaValueToColor(hexColor)
	if err != nil {
		// there isn't much I can do here, anyway this is mostly just for me anyway
		panic(err)
	}
	fmt.Printf("Using color %+v for the background\n", backgroundColor)

}

func rgbaValueToColor(hexColor *string) (color.Color, error) {
	rgbaArray, err := hex.DecodeString(*hexColor)
	if err != nil {
		return nil, err
	}

	if len(rgbaArray) != 3 && len(rgbaArray) != 4 {
		return nil, errors.New("The hex value given was too big")
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
