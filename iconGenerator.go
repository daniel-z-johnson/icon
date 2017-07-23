package main

import (
    "fmt"
    "flag"

    "github.com/prime23/hexcolor"
)

func main() {

    var bkgcString string
    var fgString string
    // var height int
    var width int
    // var outFile string

    flag.IntVar(&width, "width", 1024, "width to be used when generating icon")
    flag.StringVar(&bkgcString, "background", "f0f0f0", "background color of the icon in form of a hex string. e.x. 2ae9df")
    flag.StringVar(&fgString, "color", "c75c93", "coloor of the icon")
    flag.Parse()

    background, _ := hexcolor.HexRgb24BitToColor(bkgcString)
    color, _ := hexcolor.HexRgb24BitToColor(fgString)

    fmt.Println("--- Icon Generator ---")
    fmt.Printf("width: %d\n", width)
    fmt.Printf("Background: %t\n", background)
    fmt.Printf("Color: %t\n", color)
}
