// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

// run from the root of the repository
// go run ./lecture3/examples/file/main.go
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

	data, err := fs.ReadFile(fsys, "result.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
