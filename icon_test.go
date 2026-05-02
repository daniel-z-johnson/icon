package icon

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
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

	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
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

func TestHexToColorInvalidLength(t *testing.T) {
	_, err := hexToColor("fff")

	assert.EqualError(t, err, "hex color must be 6 or 8 characters, got 3")
}

func TestHexToColorInvalidValue(t *testing.T) {
	_, err := hexToColor("zzzzzz")

	assert.EqualError(t, err, "invalid hex color \"zzzzzz\": encoding/hex: invalid byte: U+007A 'z'")
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

func TestChangeImage(t *testing.T) {
	expected := baseImage()
	actual := baseImage()
	rgba := color.RGBA{
		R: 0x0f,
		G: 0x0f,
		B: 0xff,
		A: 0xff,
	}
	expected.Set(0, 0, rgba)
	expected.Set(0, 1, rgba)
	expected.Set(1, 0, rgba)
	expected.Set(1, 1, rgba)

	changeImage(0, 0, 2, 2, rgba, actual)

	assert.Equal(t, actual, expected)
}

func TestIconGenMirrorsWithinBounds(t *testing.T) {
	icon, err := IconGen(40, 4, "aaaaaa", "0000a0", true, true)

	assert.Nil(t, err)
	assert.Equal(t, image.Rect(0, 0, 40, 40), icon.Bounds())
	for x := 0; x < 40; x++ {
		for y := 0; y < 40; y++ {
			assert.Equal(t, icon.At(x, y), icon.At(39-x, y))
			assert.Equal(t, icon.At(x, y), icon.At(x, 39-y))
		}
	}
}

func TestIconGen(t *testing.T) {
	icon, err := IconGen(720, 8, "aaaaaa", "0000a0", false, false)
	assert.Nil(t, err)
	f1, err := os.Create("icon.png")
	png.Encode(f1, icon)
}
