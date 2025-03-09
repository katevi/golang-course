// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	fsys := os.DirFS(".")
	envVars := os.Environ()
	file, err := os.Create("result.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, envVar := range envVars {
		file.WriteString(envVar)
	}

	data, err := fs.ReadFile(fsys, "README.md")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
