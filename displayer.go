package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/coddingtonbear/go-jsonselect"
)

// Displayer provides helper methods to display command responses back to the user
// This includes optionally extracting values with JSON:select and pretty-printing
type Displayer struct {
	Response        *http.Response
	responseBody    string
	output          interface{}
	extractedValues []interface{}
	prettify        bool
}

// Factory method for displayer, reads body out of response
func NewDisplayer(resp *http.Response) (*Displayer, error) {
	defer resp.Body.Close()
	var js, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response (%s)", err.Error())
	}
	var disp = Displayer{resp, string(js), nil, nil, false}
	err = json.Unmarshal(js, &disp.output)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal response JSON: %s, response body was:\n%s",
			err.Error(), js)
	}
	return &disp, nil
}

// Apply a single extract JSON selector, it's an error if selector yields more than one value
func (d *Displayer) ApplySingleExtract(extract string) error {
	var err = d.ApplyExtract(extract)
	if err == nil && len(d.extractedValues) > 1 {
		err = fmt.Errorf("JSON selector '%s' returned more than one value, returned values are:\n%v\nOriginal JSON:\n%s",
			extract, d.extractedValues, d.responseBody)
		d.extractedValues = nil
	}
	return nil
}

// Apply JSON selector
func (d *Displayer) ApplyExtract(selector string) error {
	parser, err := jsonselect.CreateParserFromString(d.responseBody)
	if err != nil {
		return fmt.Errorf("Failed to load response JSON: %s, JSON was:\n%s",
			err.Error(), d.responseBody)
	}
	d.extractedValues, err = parser.GetValues(selector)
	if err != nil {
		return fmt.Errorf("Failed to apply JSON selector '%s' to response: %s, JSON was:\n%s",
			selector, err.Error(), d.responseBody)
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
	d.output = d.Response.Header.Get(header)
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
	if d.extractedValues != nil {
		output = d.extractedValues
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
