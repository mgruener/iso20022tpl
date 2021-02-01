// Modified copy of https://github.com/dminGod/SepaToGo to make it usable as a go generator.
// All credit goes to https://github.com/dminGod

package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"aqwari.net/xml/xsdgen"
)

func main() {

	var files []string
	var err error

	path := fmt.Sprintf("%s/*.xsd", os.Args[1])
	files, err = filepath.Glob(path)
	if err != nil {
		fmt.Printf("Could not get details of the files - Error:  %v\n", err.Error())
		os.Exit(1)
	}

	for _, f := range files {
		fmt.Printf("Finished: %v \n", f)
		generateFile(f)
	}
}

func generateFile(f string) {

	var cfg xsdgen.Config

	// Windows:
	fn := strings.Replace(f, "\\", "/", -1)

	// File and folder name:
	_, fn = path.Split(fn)
	parts := strings.Split(fn, ".")
	file := fmt.Sprintf("%s_%s_%s_%s.gen.go", parts[0], parts[1], parts[2], parts[3])
	folder := fmt.Sprintf("%s/%s/%s/%s", parts[0], parts[1], parts[2], parts[3])
	pt := fmt.Sprintf("%s/%s", folder, file)

	err := os.MkdirAll(folder, 0777)
	if err != nil {
		fmt.Printf("Failed to create directory '%s': %s", folder, err)
	}

	cfg.Option(
		xsdgen.IgnoreAttributes("id", "href", "offset"),
		xsdgen.IgnoreElements("comment"),
		xsdgen.PackageName(fmt.Sprintf(fmt.Sprintf("%s_%s_%s_%s", parts[0], parts[1], parts[2], parts[3]))),
		xsdgen.Replace("_", ""),
		xsdgen.HandleSOAPArrayType(),
		xsdgen.SOAPArrayAsSlice(),
	)

	if err := cfg.GenCLI("-o", pt, f); err != nil {
		fmt.Printf("%v", err)
	}
}
