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

package cmd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"

	camt053v4 "github.com/mgruener/iso20022tocsv/internal/iso20022/camt/053/001/04"
	"github.com/spf13/cobra"
)

func newJsonCmd(c *Cli) *cobra.Command {
	var cmd = &cobra.Command{
		Use:                   "json [file]",
		Short:                 "Convert ISO 20022 file to json",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  c.wrap(runJson),
	}

	return cmd
}

func runJson(c *Cli, cmd *cobra.Command, args []string) error {
	file := args[0]

	var fcBy []byte
	var err error

	// Open the file and fetch the contents:
	fcBy, err = ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("There was an error reading the file : %v", err.Error())
	}

	var doc camt053v4.Document

	// Get the XML object:
	err = xml.Unmarshal(fcBy, &doc)
	if err != nil {
		return fmt.Errorf("There was an error in parsing the XML : %v", err.Error())
	}

	// JSON format:
	js, _ := json.MarshalIndent(doc, "", "  ")
	fmt.Printf("%+v", string(js))

	return nil
}
