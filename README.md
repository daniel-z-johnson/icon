# Icon Generator

Makes square icons. Random ones, or the same one every time if you give it a seed.

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
	img, err := icon.Generate(icon.Options{
		Size:       720,
		Sections:   8,
		Background: "#aaa",
		Color:      "#0000a0",
		MirrorX:    true,
		MirrorY:    true,
	})
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

Want the same icon each time? Set `Seed`.

```go
img, err := icon.Generate(icon.Options{Size: 720, Sections: 8, Background: "aaa", Color: "00a", Seed: "alice"})
```

Same seed, same options, same icon. No seed, a new icon each time.

## API

```go
func Generate(opts Options) (image.Image, error)
```

`Options`:

- `Size`: width and height, in pixels.
- `Sections`: grid cells per side.
- `Background`: hex color. `"rgb"`, `"rgba"`, `"rrggbb"` or `"rrggbbaa"`. `#` optional.
- `Color`: icon color, same rules.
- `MirrorX`: mirror left to right.
- `MirrorY`: mirror top to bottom.
- `Seed`: any string. It is hashed. Empty means random.

A small border, about 1/72 of the size, surrounds the pattern.

## Errors

You get an error, not an icon, when:

- `Sections` is less than 1.
- `Size` is too small to give each section at least one pixel.
- A color is missing, not valid hex, or the wrong length.

## Testing

```sh
go test ./...
```

No external dependencies. Tests write no files.

## License

MIT. See `LICENSE`.
