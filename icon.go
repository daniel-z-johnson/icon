package icon

import (
	"encoding/hex"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math/rand/v2"
)

func IconGen(iconSize, sections int, background, iconColor string, horizontal, vertical bool) (image.Image, error) {
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
			if rand.IntN(2) == 0 {
				img = changeImage(x, y, x+stepSize, y+stepSize, mainColor, img)
			}
		}
	}

	if !horizontal && !vertical {
		return img, nil
	}

	for x := 0; x < minInt(maxX+stepSize, iconSize); x++ {
		for y := 0; y < minInt(maxY+stepSize, iconSize); y++ {
			c := img.At(x, y)
			if horizontal {
				img.Set(x, iconSize-1-y, c)
			}
			if vertical {
				img.Set(iconSize-1-x, y, c)
			}
			if horizontal && vertical {
				img.Set(iconSize-1-x, iconSize-1-y, c)
			}
		}
	}

	return img, err
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func hexToColor(hexColor string) (color.RGBA, error) {
	nilRGBA := color.RGBA{0, 0, 0, 0}
	if len(hexColor) != 6 && len(hexColor) != 8 {
		return nilRGBA, fmt.Errorf("hex color must be 6 or 8 characters, got %d", len(hexColor))
	}

	rgbaArray, err := hex.DecodeString(hexColor)
	if err != nil {
		return nilRGBA, fmt.Errorf("invalid hex color %q: %w", hexColor, err)
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
