package main

import (
	"flag"
	"fmt"
	"strings"
)

// Run: image_splitter -image somerepo/test:0.0.1
// Returns: 0.0.1
func main() {

	image := flag.String("image", "", "Provide full image name with version")
	flag.Parse()

	list := strings.Split(*image, ":")

	fmt.Println(list[len(list)-1])

}
