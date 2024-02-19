/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	URL        string
	selector   string
	structName string
)

type attribute struct {
	jsonName    string
	simpleType  string
	description string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "get-struct",
	Short: "Get a table from Harvest's documentation and output it as a Go struct",
	Long:  "Get a table from Harvest's documentation and output it as a Go struct",
	Run: func(cmd *cobra.Command, args []string) {
		attributes := []attribute{}
		col := colly.NewCollector()
		col.OnHTML(selector, func(e *colly.HTMLElement) {
			isRequestBody := false
			e.ForEach("tr", func(i int, element *colly.HTMLElement) {
				cols := strings.Split(strings.Trim(element.Text, " \n"), "\n")
				for j, c := range cols {
					cols[j] = strings.Trim(c, " \n")
				}
				if i == 0 {
					isRequestBody = reflect.DeepEqual(cols, []string{"Parameter", "Type", "Required", "Description"})
				} else if i > 0 {
					if isRequestBody {
						attributes = append(attributes, attribute{cols[0], cols[1], cols[3] + " - " + cols[2]})
					} else {
						attributes = append(attributes, attribute{cols[0], cols[1], cols[2]})
					}
				}
			})
		})
		col.Visit(URL)

		structString := writeStruct(attributes)
		err := writeStructFile(structString)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing struct: %s", err.Error())
		} else {
			fmt.Println("Success!")
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&URL, "url", "u", "", "URL of the Harvest documentation page")
	rootCmd.Flags().StringVarP(&selector, "selector", "s", "", "CSS selector for the table")
	rootCmd.Flags().StringVarP(&structName, "struct-name", "n", "", "Name of the Go struct")
}

func writeStruct(attributes []attribute) string {
	typeMap := map[string]string{
		"boolean":  "bool",
		"string":   "string",
		"integer":  "int",
		"decimal":  "float64",
		"date":     "Date",
		"datetime": "time.Time",
		"time":     "time.Time",
		"array":    "[]int",
		"object":   "struct{}",
	}
	structString := strings.Builder{}
	structString.WriteString(fmt.Sprintf("type %s struct {\n", structName))
	for _, a := range attributes {
		attributeName := snakeToPascal(strings.Trim(a.jsonName, " "))
		structString.WriteString("\t")
		structString.WriteString(attributeName)
		structString.WriteString(" ")
		structString.WriteString(typeMap[strings.Trim(a.simpleType, " ")])
		structString.WriteString(" ")
		structString.WriteString(fmt.Sprintf("`json:\"%s\" url:\"%s,omitempty\"`", strings.Trim(a.jsonName, " "), strings.Trim(a.jsonName, " ")))
		structString.WriteString(" ")
		structString.WriteString("// ")
		structString.WriteString(strings.Trim(a.description, " "))
		structString.WriteString("\n")
	}
	structString.WriteString("}\n")
	return structString.String()
}

func snakeToPascal(s string) string {
	exceptions := map[string]string{
		"Id":  "ID",
		"Url": "URL",
	}
	words := strings.Split(s, "_")
	var out string
	for _, word := range words {
		w := cases.Title(language.AmericanEnglish).String(word)
		if v, ok := exceptions[w]; ok {
			out += v
		} else {
			out += w
		}
	}
	return out
}

func writeStructFile(structString string) error {
	if _, err := os.Stat("./dist"); err != nil && !os.IsExist(err) {
		err := os.Mkdir("./dist", 0755)
		if err != nil {
			return err
		}
	}
	return os.WriteFile(fmt.Sprintf("./dist/%s.go", structName), []byte(structString), 0644)
}
