package main

import (
	"flag"
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Icon Generator start")

	rgbaValue := flag.String("background", "240,240,240", "The value for the background color, in quotes comma seperated 3 or 4 values, between 0-255")
	flag.Parse()

	fmt.Printf("Background color: '%s'\n", *rgbaValue)

	backgroundColor, err := rgbaValueToColor(rgbaValue)
	if err != nil {
		// there isn't much I can do here, anyway this is mostly just for me anyway
		panic(err)
	}
	fmt.Printf("Using color %+v for the background\n", backgroundColor)

}

func rgbaValueToColor(rgbaValue *string) (color.Color, error) {
	rgbaArrayValue := strings.Split(*rgbaValue, ",")
	if len(rgbaArrayValue) != 3 && len(rgbaArrayValue) != 4 {
		return nil, fmt.Errorf("There shoudl be 3 for 4 values, found '%d' instead", len(rgbaArrayValue))
	}
	rgbaArray := make([]uint8, 0, 0)
	for _, value := range rgbaArrayValue {
		part, err := strconv.ParseUint(value, 10, 8)
		if err != nil {
			return nil, fmt.Errorf("Could cast '%s' into an int8", value)
		}
		rgbaArray = append(rgbaArray, uint8(part))
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
