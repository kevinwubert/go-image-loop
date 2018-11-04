package main

import (
	"fmt"
	"os"
)

func usage() string {
	return "./go-image-loop <image1-path> <image2-path> <output-path>"
}

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) != 3 {
		fmt.Println("Usage error.")
		fmt.Println(usage())
		return
	}

	image1Path := argsWithoutProg[0]
	image2Path := argsWithoutProg[1]
	outputPath := argsWithoutProg[2]

	imageloop.CreateImageLoop(image1Path, image2Path, outputPath)
}
