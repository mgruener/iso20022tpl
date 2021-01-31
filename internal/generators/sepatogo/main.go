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
	pkg := os.Args[2]
	files, err = filepath.Glob(path)
	if err != nil {
		fmt.Printf("Could not get details of the files - Error:  %v\n", err.Error())
		os.Exit(1)
	}

	for _, f := range files {
		fmt.Printf("Finished: %v \n", f)
		generateFile(f, pkg)
	}
}

func generateFile(f string, pkg string) {

	var cfg xsdgen.Config

	// Windows:
	fn := strings.Replace(f, "\\", "/", -1)

	// File and folder name:
	_, fn = path.Split(fn)
	fn = strings.Replace(fn, ".", "_", -1)
	fn = strings.TrimRight(fn, "_xsd")
	pt := fmt.Sprintf("%s.gen.go", fn)

	cfg.Option(
		xsdgen.IgnoreAttributes("id", "href", "offset"),
		xsdgen.IgnoreElements("comment"),
		xsdgen.PackageName(pkg),
		xsdgen.Replace("_", ""),
		xsdgen.HandleSOAPArrayType(),
		xsdgen.SOAPArrayAsSlice(),
	)

	if err := cfg.GenCLI("-o", pt, f); err != nil {
		fmt.Printf("%v", err)
	}
}
