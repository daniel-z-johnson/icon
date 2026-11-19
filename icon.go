package icon

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math/rand/v2"
	"strings"
)

// marginDivisor sets the empty border around the pattern as a fraction of the
// icon size: each side gets Size/marginDivisor pixels (10px at 720px).
const marginDivisor = 72

// Options describes an icon. Colors are hex strings: "rgb", "rgba",
// "rrggbb" or "rrggbbaa", with or without a leading "#".
type Options struct {
	Size       int    // width and height in pixels
	Sections   int    // grid cells per side
	Background string // background color
	Color      string // icon color
	MirrorX    bool   // mirror left to right
	MirrorY    bool   // mirror top to bottom
	Seed       string // if set, hashed so the same seed gives the same icon; if empty, random
}

// Generate returns a square icon described by opts.
func Generate(opts Options) (image.Image, error) {
	if opts.Sections < 1 {
		return nil, fmt.Errorf("sections must be at least 1, got %d", opts.Sections)
	}
	margin := opts.Size / marginDivisor
	stepSize := (opts.Size - 2*margin) / opts.Sections
	if stepSize < 1 {
		return nil, fmt.Errorf("icon size %d is too small for %d sections", opts.Size, opts.Sections)
	}
	bgColor, err := hexToColor(opts.Background)
	if err != nil {
		return nil, fmt.Errorf("background: %w", err)
	}
	mainColor, err := hexToColor(opts.Color)
	if err != nil {
		return nil, fmt.Errorf("color: %w", err)
	}

	grid := pattern(source(opts.Seed), opts.Sections, opts.MirrorX, opts.MirrorY)

	// Center the grid so any leftover pixels are split between both sides.
	offset := (opts.Size - stepSize*opts.Sections) / 2

	img := image.NewRGBA(image.Rect(0, 0, opts.Size, opts.Size))
	fillRect(img, img.Bounds(), bgColor)
	for i := range opts.Sections {
		x0, x1 := cellRange(i, opts.Sections, stepSize, offset, opts.Size, opts.MirrorX)
		for j := range opts.Sections {
			if !grid[i][j] {
				continue
			}
			y0, y1 := cellRange(j, opts.Sections, stepSize, offset, opts.Size, opts.MirrorY)
			fillRect(img, image.Rect(x0, y0, x1, y1), mainColor)
		}
	}
	return img, nil
}

// source returns a random source, seeded from a SHA-256 hash of seed when
// seed is not empty.
func source(seed string) rand.Source {
	if seed == "" {
		return rand.NewPCG(rand.Uint64(), rand.Uint64())
	}
	sum := sha256.Sum256([]byte(seed))
	return rand.NewPCG(binary.LittleEndian.Uint64(sum[:8]), binary.LittleEndian.Uint64(sum[8:16]))
}

// pattern decides which cells are filled. Cells on a mirrored half copy their
// counterpart, so only the canonical half (or quadrant) is random.
func pattern(src rand.Source, n int, mirrorX, mirrorY bool) [][]bool {
	bits := bitReader{src: src}
	grid := make([][]bool, n)
	for i := range grid {
		grid[i] = make([]bool, n)
	}
	for i := range n {
		for j := range n {
			ci, cj := i, j
			if mirrorX {
				ci = min(i, n-1-i)
			}
			if mirrorY {
				cj = min(j, n-1-j)
			}
			if ci == i && cj == j {
				grid[i][j] = bits.next()
			} else {
				// The canonical cell has smaller indices, so it is already set.
				grid[i][j] = grid[ci][cj]
			}
		}
	}
	return grid
}

// cellRange returns the pixel span [lo, hi) of cell i along one axis. On a
// mirrored axis, cells in the second half are exact pixel mirrors of the first
// half and a middle cell (odd n) stretches to stay symmetric about the center.
func cellRange(i, n, step, offset, size int, mirrored bool) (lo, hi int) {
	m := n - 1 - i
	if mirrored && i > m {
		return size - (offset + (m+1)*step), size - (offset + m*step)
	}
	lo = offset + i*step
	hi = lo + step
	if mirrored && i == m {
		hi = size - lo
	}
	return lo, hi
}

// bitReader hands out random bits one at a time. Pulling raw bits from the
// source (rather than rand.Rand helpers) keeps seeded output stable.
type bitReader struct {
	src  rand.Source
	buf  uint64
	left int
}

func (b *bitReader) next() bool {
	if b.left == 0 {
		b.buf = b.src.Uint64()
		b.left = 64
	}
	bit := b.buf&1 == 1
	b.buf >>= 1
	b.left--
	return bit
}

func hexToColor(hexColor string) (color.RGBA, error) {
	h := strings.TrimPrefix(hexColor, "#")
	switch len(h) {
	case 3, 4: // shorthand: "fa0" means "ffaa00"
		long := make([]byte, 0, 2*len(h))
		for i := range len(h) {
			long = append(long, h[i], h[i])
		}
		h = string(long)
	case 6, 8:
	default:
		return color.RGBA{}, fmt.Errorf("hex color must be 3, 4, 6 or 8 characters, got %d", len(h))
	}

	rgba, err := hex.DecodeString(h)
	if err != nil {
		return color.RGBA{}, fmt.Errorf("invalid hex color %q: %w", hexColor, err)
	}

	if len(rgba) == 3 {
		rgba = append(rgba, 255)
	}
	return color.RGBA{R: rgba[0], G: rgba[1], B: rgba[2], A: rgba[3]}, nil
}

func fillRect(img draw.Image, r image.Rectangle, c color.Color) {
	draw.Draw(img, r, image.NewUniform(c), image.Point{}, draw.Src)
}
