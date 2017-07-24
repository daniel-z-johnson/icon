package main

import (
    "fmt"
    "flag"

    "github.com/prime23/hexcolor"
)

func main() {

    var bkgcString string
    var fgString string
    var side int
    // var outFile string

    flag.IntVar(&side, "side", 1024, "side to be used when generating icon")
    flag.StringVar(&bkgcString, "background", "f0f0f0", "background color of the icon in form of a hex string. e.x. 2ae9df")
    flag.StringVar(&fgString, "color", "c75c93", "coloor of the icon")
    flag.Parse()

    background, _ := hexcolor.HexRgb24BitToColor(bkgcString)
    color, _ := hexcolor.HexRgb24BitToColor(fgString)

    fmt.Println("--- Icon Generator ---")
    fmt.Printf("side: %d\n", side)
    fmt.Printf("Background: %t\n", background)
    fmt.Printf("Color: %t\n", color)
}
