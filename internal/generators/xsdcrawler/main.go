/*
Copyright © 2021 Michael Gruener

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"archive/zip"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gocolly/colly/v2"
)

func main() {
	crawl()
}

func crawl() {
	dir := os.Args[1]
	os.MkdirAll(dir, 0777)

	c := colly.NewCollector(
		colly.AllowedDomains("www.iso20022.org"),
	)

	// Visit all download links of the current ISO 20022 message catalogue
	c.OnHTML("a[href$=download][aria-label=Download]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	// Visit all download links of the ISO 20022 message catalogue archive
	c.OnHTML("a[href$=download][aria-label='Download Complete Message Set']", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	// Handle pageing
	c.OnHTML("a[href][title^='Go to page']", func(e *colly.HTMLElement) {
		fmt.Printf("Visiting next page\n")
		e.Request.Visit(e.Attr("href"))
	})

	// Download message catalogue zip files
	c.OnResponse(func(r *colly.Response) {
		if strings.Index(r.Headers.Get("Content-Type"), "application/zip") > -1 {
			fmt.Printf("Processing %s\n", r.FileName())
			if err := unzipWriteFile(r.Body, dir); err != nil {
				fmt.Printf("Failed to unzip file '%s': %s\n", r.FileName(), err)
			}
		}
	})

	c.Visit("https://www.iso20022.org/iso-20022-message-definitions?page=0")
	c.Visit("https://www.iso20022.org/catalogue-messages/iso-20022-messages-archive?page=0")
}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func unzipWriteFile(data []byte, dir string) error {
	mime, err := mimetype.DetectReader(bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("failed to detect mimetype")
	}
	if !mime.Is("application/zip") {
		return fmt.Errorf("data is not a zipfile")
	}

	zipReader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return fmt.Errorf("failed to read zipped data: %s (%s)", err, mime)
	}

	// Read all the files from zip archive
	for _, zipFile := range zipReader.File {
		// we are only interested in xml schema and zip files
		if path.Ext(zipFile.Name) != "xsd" && path.Ext(zipFile.Name) != "zip" {
			continue
		}
		fmt.Printf("  extracting %s\n", zipFile.Name)
		unzippedFileBytes, err := readZipFile(zipFile)
		if err != nil {
			return fmt.Errorf("failed to unzip file '%s': %s", zipFile.Name, err)
		}

		if path.Ext(zipFile.Name) == "zip" {
			if err := unzipWriteFile(unzippedFileBytes, dir); err != nil {
				return err
			}
			continue
		}

		// store the extracted file
		if err := ioutil.WriteFile(fmt.Sprintf("%s/%s", dir, zipFile.Name), unzippedFileBytes, 0644); err != nil {
			return fmt.Errorf("failed to save extracted file as '%s/%s': %s", dir, zipFile.Name, err)
		}
	}

	return nil
}
