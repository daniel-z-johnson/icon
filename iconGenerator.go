package main

import (
	"os"
    "fmt"
    "flag"

    "github.com/prime23/hexcolor"
)

func main() {

    var bkgcString string
    var fgString string
    var side int
    var outFile string
    var numSquares int
    var padding int

    flag.IntVar(&side, "side", 1024, "side to be used when generating icon")
    flag.StringVar(&bkgcString, "background", "f0f0f0", "background color of the icon in form of a hex string. e.x. 2ae9df")
    flag.StringVar(&fgString, "color", "c75c93", "coloor of the icon")
	flag.StringVar(&outFile, "file", "icon.png", "name of file for icon file")
    flag.IntVar(&numSquares, "squares", 5, "number of squares for row and column")
    flag.IntVar(&padding, "padding", 12, "amount of padding, in pixels, for the icon")
    flag.Parse()

    background, _ := hexcolor.HexRgb24BitToColor(bkgcString)
    color, _ := hexcolor.HexRgb24BitToColor(fgString)
	f1, _ := os.OpenFile(outFile, os.O_CREATE | os.O_WRONLY, 0644)

    fmt.Println("--- Icon Generator ---")
    fmt.Printf("side: %d\n", side)
    fmt.Printf("Background: %t\n", background)
    fmt.Printf("Color: %t\n", color)
    fmt.Printf("f1: %t\n", f1)
    fmt.Printf("squares: %t\n", numSquares)
    fmt.Printf("padding: %t\n", padding)
}
