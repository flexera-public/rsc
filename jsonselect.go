package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rightscale/go-jsonselect"
)

// Apply given jsonselect expression to STDIN, see http://jsonselect.org
// Output result to STDOUT and any error to STDERR
func ApplyJsonselect(exp string) error {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	parser, err := jsonselect.CreateParserFromString(string(bytes))
	if err != nil {
		return err
	}
	results, err := parser.GetValues(exp)
	if err != nil {
		return err
	}
	res := interface{}(results)
	if len(results) == 1 {
		res = results[0]
	}
	m, err := json.MarshalIndent(res, "", "    ")
	if err != nil {
		return err
	}
	fmt.Print(string(m))
	return nil
}
