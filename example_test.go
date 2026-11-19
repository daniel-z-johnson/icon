package icon_test

import (
	"fmt"
	"image/png"
	"io"
	"log"

	"github.com/daniel-z-johnson/icon"
)

func ExampleGenerate() {
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

	// Write to a file instead of io.Discard to keep it.
	if err := png.Encode(io.Discard, img); err != nil {
		log.Fatal(err)
	}
	fmt.Println(img.Bounds())
	// Output: (0,0)-(720,720)
}

func ExampleGenerate_seed() {
	opts := icon.Options{Size: 64, Sections: 6, Background: "aaa", Color: "00a", MirrorX: true, Seed: "alice"}
	a, _ := icon.Generate(opts)
	b, _ := icon.Generate(opts)

	same := true
	for x := range 64 {
		for y := range 64 {
			if a.At(x, y) != b.At(x, y) {
				same = false
			}
		}
	}
	fmt.Println("same icon:", same)
	// Output: same icon: true
}
