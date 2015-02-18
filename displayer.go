package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/coddingtonbear/go-jsonselect"
)

// Displayer provides helper methods to display command responses back to the user
// This includes optionally extracting values with JSON:select and pretty-printing
type Displayer struct {
	response *http.Response
	body     string
	output   interface{}
	prettify bool
}

// Factory method for displayer, reads body out of response
func NewDisplayer(resp *http.Response) (*Displayer, error) {
	defer resp.Body.Close()
	var js, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response (%s)", err.Error())
	}
	var disp = Displayer{response: resp, body: string(js)}
	err = json.Unmarshal(js, &disp.output)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal response JSON: %s, response body was:\n%s",
			err.Error(), js)
	}
	return &disp, nil
}

// Apply a single extract JSON selector, it's an error if selector yields more than one value
func (d *Displayer) ApplySingleExtract(extract string) error {
	if err := d.ApplyExtract(extract, true); err != nil {
		return err
	}
	var outputs = d.output.([]interface{})
	if len(outputs) > 1 {
		return fmt.Errorf("JSON selector '%s' returned more than one value, returned values are:\n%v\nOriginal JSON:\n%s",
			extract, outputs, d.body)
	}
	d.output = outputs[0]
	return nil
}

// Apply JSON selector
func (d *Displayer) ApplyExtract(selector string, json bool) error {
	var parser, err = jsonselect.CreateParserFromString(d.body)
	if err != nil {
		return fmt.Errorf("Failed to load response JSON: %s, JSON was:\n%s",
			err.Error(), d.body)
	}
	outputs, err := parser.GetValues(selector)
	if !json {
		var strs = make([]string, len(outputs))
		for i, o := range outputs {
			strs[i] = fmt.Sprintf("%v", o)
		}
		d.output = strings.Join(strs, " ")
	} else {
		d.output = outputs
	}
	if err != nil {
		return fmt.Errorf("Failed to apply JSON selector '%s' to response: %s, JSON was:\n%s",
			selector, err.Error(), d.body)
	}
	return nil
}

// Apply JSON selector to locate link
func (d *Displayer) ApplyLinkExtract(link string) error {
	// TBD
	return nil
}

// Apply header extraction
func (d *Displayer) ApplyHeaderExtract(header string) error {
	d.output = d.response.Header.Get(header)
	if d.output == "" {
		return fmt.Errorf("Response does not contain the '%s' header", header)
	}
	return nil
}

// Prettify output
func (d *Displayer) Pretty() {
	d.prettify = true
}

// Return output
func (d *Displayer) Output() string {
	var output = d.output
	if outputStr, ok := d.output.(string); ok {
		return outputStr
	}
	var b []byte
	var err error
	if d.prettify {
		b, err = json.MarshalIndent(output, "", "    ")
	} else {
		b, err = json.Marshal(output)
	}
	if err != nil {
		return fmt.Sprintf("%v", output)
	}
	return string(b)
}
