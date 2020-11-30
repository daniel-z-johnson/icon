package icon

import (
	"image"
	"image/color"
	"image/draw"
	"testing"

	"github.com/stretchr/testify/assert"
)

func baseImage() draw.Image {
	img := image.NewRGBA(image.Rect(0, 0, 4, 4))
	rgba := color.RGBA{
		R: 0xff,
		G: 0xaf,
		B: 0xaf,
		A: 0xff,
	}

	for x := 0; x <= 4; x++ {
		for y := 0; y <= 4; y++ {
			img.Set(x, y, rgba)
		}
	}

	return img
}

func TestHexToColor(t *testing.T) {

	assert := assert.New(t)

	expected := color.RGBA{
		R: 0xff,
		G: 0xcc,
		B: 0x1f,
		A: 0xff,
	}

	actual, err := hexToColor("ffcc1f")

	assert.Nil(err)
	assert.Equal(expected, actual)

	expected = color.RGBA{
		R: 0xff,
		G: 0xcc,
		B: 0x1f,
		A: 0xfa,
	}

	actual, err = hexToColor("ffcc1ffa")

	assert.Nil(err)
	assert.Equal(expected, actual)
}

func TestInitialImage(t *testing.T) {
	expected := baseImage()
	rgba := color.RGBA{
		R: 0xff,
		G: 0xaf,
		B: 0xaf,
		A: 0xff,
	}

	actual := initialImage(4, rgba)

	assert.Equal(t, expected, actual)
}
