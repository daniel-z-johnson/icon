package icon

import (
	"encoding/hex"
	"errors"
	"image"
	"image/color"
	"image/draw"
	"math/rand"
	"time"
)

func IconGen(iconSize, sections int, background, iconColor string, horizontal, vertical bool) (image.Image, error) {
	rand.Seed(time.Now().Unix())
	bgColor, err := hexToColor(background)
	if err != nil {
		return nil, err
	}
	mainColor, err := hexToColor(iconColor)
	if err != nil {
		return nil, err
	}
	img := initialImage(iconSize, bgColor)
	stepSize := (iconSize - 20) / sections
	maxX := iconSize - stepSize
	if vertical {
		maxX /= 2
	}
	maxY := iconSize - stepSize
	if horizontal {
		maxY /= 2
	}
	for x := 10; x < maxX; x += stepSize {
		for y := 10; y < maxY; y += stepSize {
			if rand.Int()%2 == 0 {
				img = changeImage(x, y, x+stepSize, y+stepSize, mainColor, img)
			}
		}
	}

	for x := 0; x <= maxX + stepSize; x++ {
		for y := 0; y <= maxY + stepSize; y++ {
			c := img.At(x, y)
			if horizontal {
				img.Set(x, iconSize - y, c)
			}
			if vertical {
				img.Set(iconSize - x, y, c)
			}
			if horizontal && vertical {
				img.Set(iconSize - x, iconSize - y, c)
			}
		}
	}

	return img, err
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
