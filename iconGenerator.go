package main

import (
    "fmt"
    "flag"
)

func main() {
    fmt.Println("--- Icon Generator ---")

    var bkgcString string
    // var fgString string
    // var height int
    var width int
    // var outFile string

    flag.IntVar(&width, "width", 1024, "width to be used when generating icon")
    flag.StringVar(&bkgcString, "background", "f0f0f0", "background color of the icon in form of a hex string. e.x. 2ae9df")
    flag.Parse()
    fmt.Printf("width: %d\n", width)
    fmt.Printf("Color: %s\n", bkgcString)
}
