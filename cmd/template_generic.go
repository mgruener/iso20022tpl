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
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"text/template"

	"github.com/mgruener/iso20022tpl/internal/iso20022"
	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"
)

func newTemplateGenericCmd(c *Cli) *cobra.Command {
	var cmd = &cobra.Command{
		Use:                   "generic [iso20022 file] [template]",
		Short:                 "Render a generic template based on a ISO 20022 file",
		Args:                  cobra.ExactArgs(2),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  c.wrap(runTemplateGeneric),
	}

	cmd.Flags().StringP("output", "o", "", "Write output to file")
	cmd.Flags().String("validate", "", "Validate the result as (json,yaml,csv)")
	cmd.Flags().String("csv-separator", ",", "CSV separator used when validating csv")

	return cmd
}

func runTemplateGeneric(c *Cli, cmd *cobra.Command, args []string) error {
	isoFile := args[0]
	templateFile := args[1]

	doc, err := iso20022.New(isoFile)
	if err != nil {
		return fmt.Errorf("failed to load iso20022 document from '%s': %s", isoFile, err)
	}

	td, err := template.ParseFiles(templateFile)
	if err != nil {
		return fmt.Errorf("failed to load template from '%s': %s", templateFile, err)
	}

	var buf bytes.Buffer
	if err := td.Execute(&buf, doc); err != nil {
		return fmt.Errorf("failed to render template: %s", err)
	}

	if err := validate(c, buf.Bytes()); err != nil {
		return fmt.Errorf("failed to validate result as '%s': %s", c.viper.GetString("validate"), err)
	}

	return output(buf.Bytes(), c.viper.GetString("output"))
}

func validate(c *Cli, data []byte) error {
	as := c.viper.GetString("validate")
	if as == "" {
		return nil
	}

	var doc interface{}
	var err error
	switch as {
	case "json":
		err = json.Unmarshal(data, &doc)
	case "yaml":
		err = yaml.Unmarshal(data, &doc)
	case "csv":
		err = validateCSV(c, data)
	default:
		return fmt.Errorf("unknown validation format: %s", as)
	}

	return err
}

func validateCSV(c *Cli, data []byte) error {
	sep := c.viper.GetString("csv-separator")
	csvReader := csv.NewReader(bytes.NewReader(data))
	csvReader.Comma = ([]rune(sep))[0]

	for {
		if _, err := csvReader.Read(); err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
	}
}

func output(data []byte, dest string) error {
	if dest != "" {
		if err := ioutil.WriteFile(dest, data, 0644); err != nil {
			return fmt.Errorf("failed to write output to '%s': %s", dest, err)
		}
		return nil
	}

	fmt.Printf("%s", string(data))
	return nil
}
