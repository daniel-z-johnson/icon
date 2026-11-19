package icon

import (
	"bytes"
	"encoding/hex"
	"errors"
	"image"
	"image/color"
	"math/rand/v2"
	"strings"
	"testing"
)

func baseImage() *image.RGBA {
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

// opts returns valid Options for the given size and sections.
func opts(size, sections int, mirrorX, mirrorY bool) Options {
	return Options{
		Size:       size,
		Sections:   sections,
		Background: "aaaaaa",
		Color:      "0000a0",
		MirrorX:    mirrorX,
		MirrorY:    mirrorY,
	}
}

func sameImage(a, b image.Image) bool {
	ra, ok1 := a.(*image.RGBA)
	rb, ok2 := b.(*image.RGBA)
	return ok1 && ok2 && ra.Rect == rb.Rect && bytes.Equal(ra.Pix, rb.Pix)
}

func mustNoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func wantErr(t *testing.T, err error, want string) {
	t.Helper()
	if err == nil {
		t.Fatalf("want error %q, got nil", want)
	}
	if err.Error() != want {
		t.Fatalf("want error %q, got %q", want, err.Error())
	}
}

func TestHexToColor(t *testing.T) {
	cases := []struct {
		in   string
		want color.RGBA
	}{
		{"ffcc1f", color.RGBA{R: 0xff, G: 0xcc, B: 0x1f, A: 0xff}},
		{"ffcc1ffa", color.RGBA{R: 0xff, G: 0xcc, B: 0x1f, A: 0xfa}},
		{"#ffcc1f", color.RGBA{R: 0xff, G: 0xcc, B: 0x1f, A: 0xff}},
		{"fc1", color.RGBA{R: 0xff, G: 0xcc, B: 0x11, A: 0xff}},
		{"#fc18", color.RGBA{R: 0xff, G: 0xcc, B: 0x11, A: 0x88}},
	}
	for _, tc := range cases {
		got, err := hexToColor(tc.in)
		mustNoErr(t, err)
		if got != tc.want {
			t.Errorf("hexToColor(%q) = %v, want %v", tc.in, got, tc.want)
		}
	}
}

func TestHexToColorInvalidLength(t *testing.T) {
	for _, in := range []string{"", "ff", "#fffff", "fffffffff"} {
		_, err := hexToColor(in)
		if err == nil || !strings.HasPrefix(err.Error(), "hex color must be 3, 4, 6 or 8 characters") {
			t.Errorf("hexToColor(%q): want length error, got %v", in, err)
		}
	}
}

func TestHexToColorInvalidValue(t *testing.T) {
	_, err := hexToColor("zzzzzz")

	if err == nil {
		t.Fatal("want error, got nil")
	}
	if !strings.HasPrefix(err.Error(), `invalid hex color "zzzzzz": `) {
		t.Errorf("unexpected error message: %q", err)
	}
	var hexErr hex.InvalidByteError
	if !errors.As(err, &hexErr) || byte(hexErr) != 'z' {
		t.Errorf("want wrapped hex.InvalidByteError('z'), got %v", err)
	}
}

func TestFillRectWholeImage(t *testing.T) {
	expected := baseImage()
	rgba := color.RGBA{
		R: 0xff,
		G: 0xaf,
		B: 0xaf,
		A: 0xff,
	}

	actual := image.NewRGBA(image.Rect(0, 0, 4, 4))
	fillRect(actual, actual.Bounds(), rgba)

	if !sameImage(expected, actual) {
		t.Error("filled image does not match expected")
	}
}

func TestFillRect(t *testing.T) {
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

	fillRect(actual, image.Rect(0, 0, 2, 2), rgba)

	if !sameImage(expected, actual) {
		t.Error("filled image does not match expected")
	}
}

func TestGenerateRejectsBadOptions(t *testing.T) {
	badBg := opts(64, 4, false, false)
	badBg.Background = "12345"
	noColor := opts(64, 4, false, false)
	noColor.Color = ""

	cases := []struct {
		name    string
		opts    Options
		wantErr string
	}{
		{"zero sections", opts(64, 0, false, false), "sections must be at least 1, got 0"},
		{"negative sections", opts(64, -3, false, false), "sections must be at least 1, got -3"},
		{"too small", opts(3, 4, false, false), "icon size 3 is too small for 4 sections"},
		{"zero size", opts(0, 1, false, false), "icon size 0 is too small for 1 sections"},
		{"bad background", badBg, "background: hex color must be 3, 4, 6 or 8 characters, got 5"},
		{"missing color", noColor, "color: hex color must be 3, 4, 6 or 8 characters, got 0"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			icon, err := Generate(tc.opts)
			if icon != nil {
				t.Errorf("want nil image, got %v", icon.Bounds())
			}
			wantErr(t, err, tc.wantErr)
		})
	}
}

func TestGenerateSmallSizes(t *testing.T) {
	// Sizes that used to hang (stepSize <= 0 with the old fixed 10px margin).
	for _, size := range []int{4, 10, 20, 24} {
		icon, err := Generate(opts(size, 4, true, true))
		mustNoErr(t, err)
		if want := image.Rect(0, 0, size, size); icon.Bounds() != want {
			t.Errorf("size %d: bounds %v, want %v", size, icon.Bounds(), want)
		}
	}
}

func assertSymmetric(t *testing.T, img image.Image, mirrorX, mirrorY bool) {
	t.Helper()
	n := img.Bounds().Dx()
	for x := range n {
		for y := range n {
			if mirrorX && img.At(x, y) != img.At(n-1-x, y) {
				t.Fatalf("not mirrored left/right at (%d,%d)", x, y)
			}
			if mirrorY && img.At(x, y) != img.At(x, n-1-y) {
				t.Fatalf("not mirrored top/bottom at (%d,%d)", x, y)
			}
		}
	}
}

func TestGenerateMirrorOneAxis(t *testing.T) {
	for range 20 {
		lr, err := Generate(opts(101, 5, true, false))
		mustNoErr(t, err)
		assertSymmetric(t, lr, true, false)

		tb, err := Generate(opts(101, 5, false, true))
		mustNoErr(t, err)
		assertSymmetric(t, tb, false, true)
	}
}

func TestGenerateOddSectionsMirrored(t *testing.T) {
	// Odd section counts and odd leftover pixels must still be symmetric.
	for _, size := range []int{100, 101, 720, 723} {
		for _, sections := range []int{1, 3, 5, 7} {
			icon, err := Generate(opts(size, sections, true, true))
			mustNoErr(t, err)
			assertSymmetric(t, icon, true, true)
		}
	}
}

func TestPatternOddSectionsHasCenterColumn(t *testing.T) {
	// With 5 sections and mirroring, all 5 columns are drawn: the center
	// column must be independently random, not always empty.
	sawCenter := false
	for seed := range 50 {
		grid := pattern(rand.NewPCG(uint64(seed), 0), 5, true, true)
		for j := range 5 {
			sawCenter = sawCenter || grid[2][j]
		}
	}
	if !sawCenter {
		t.Error("center column never filled")
	}
}

func TestGenerateSeededIsDeterministic(t *testing.T) {
	seeded := func(seed string) image.Image {
		o := opts(64, 6, false, true)
		o.Seed = seed
		img, err := Generate(o)
		mustNoErr(t, err)
		return img
	}

	a, b, c := seeded("alice"), seeded("alice"), seeded("bob")

	if !sameImage(a, b) {
		t.Error("same seed produced different images")
	}
	if sameImage(a, c) {
		t.Error("different seeds produced the same image")
	}
}

func TestGenerateReturnsSquareImage(t *testing.T) {
	icon, err := Generate(opts(64, 4, false, false))
	mustNoErr(t, err)

	if want := image.Rect(0, 0, 64, 64); icon.Bounds() != want {
		t.Errorf("bounds %v, want %v", icon.Bounds(), want)
	}
}
