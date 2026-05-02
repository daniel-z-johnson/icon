# Icon Generator

A small Go package for generating randomized square icons.

The package creates an `image.Image` by filling a square canvas with a
background color, randomly coloring grid sections, and optionally mirroring the
pattern horizontally and/or vertically.

## Installation

```sh
go get github.com/daniel-z-johnson/icon
```

## Usage

```go
package main

import (
	"image/png"
	"log"
	"os"

	"github.com/daniel-z-johnson/icon"
)

func main() {
	img, err := icon.IconGen(
		720,      // icon size in pixels
		8,        // number of grid sections
		"aaaaaa", // background color, hex RGB or RGBA
		"0000a0", // icon color, hex RGB or RGBA
		true,     // mirror horizontally
		true,     // mirror vertically
	)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("icon.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		log.Fatal(err)
	}
}
```

## API

```go
func IconGen(iconSize, sections int, background, iconColor string, horizontal, vertical bool) (image.Image, error)
```

Parameters:

- `iconSize`: width and height of the generated square image, in pixels.
- `sections`: number of grid sections used to divide the drawable area.
- `background`: background color as a hex string, either RGB (`"ffffff"`) or RGBA (`"ffffffff"`).
- `iconColor`: foreground color as a hex string, either RGB (`"000000"`) or RGBA (`"000000ff"`).
- `horizontal`: mirrors the generated pattern across the horizontal axis.
- `vertical`: mirrors the generated pattern across the vertical axis.

Hex colors should not include a leading `#`.

## Testing

```sh
go test ./...
```

The current `TestIconGen` test writes an `icon.png` file in the repository root.
