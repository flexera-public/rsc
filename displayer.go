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
	response  *http.Response
	body      string
	RawOutput interface{}
	prettify  bool
}

// Factory method for displayer, reads body out of response
func NewDisplayer(resp *http.Response) (*Displayer, error) {
	defer resp.Body.Close()
	js, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response (%s)", err.Error())
	}
	disp := Displayer{response: resp, body: string(js)}
	if len(js) > 2 {
		err = json.Unmarshal(js, &disp.RawOutput)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal response JSON: %s, response body was:\n%s",
				err.Error(), js)
		}
	}
	return &disp, nil
}

// Apply a single extract JSON selector, it's an error if selector yields more than one value
func (d *Displayer) ApplySingleExtract(extract string) error {
	if err := d.ApplyExtract(extract, true); err != nil {
		return err
	}
	outputs := d.RawOutput.([]interface{})
	if len(outputs) > 1 {
		return fmt.Errorf("JSON selector '%s' returned more than one value, returned values are:\n%v\nOriginal JSON:\n%s",
			extract, outputs, d.body)
	}
	if len(outputs) == 0 {
		d.RawOutput = ""
	} else {
		d.RawOutput = outputs[0]
	}
	return nil
}

// Apply JSON selector
func (d *Displayer) ApplyExtract(selector string, json bool) error {
	parser, err := jsonselect.CreateParserFromString(d.body)
	if err != nil {
		return fmt.Errorf("Failed to load response JSON: %s, JSON was:\n%s",
			err.Error(), d.body)
	}
	outputs, err := parser.GetValues(selector)
	if !json {
		strs := make([]string, len(outputs))
		for i, o := range outputs {
			strs[i] = fmt.Sprintf("%v", o)
		}
		d.RawOutput = strings.Join(strs, " ")
	} else {
		d.RawOutput = outputs
	}
	if err != nil {
		return fmt.Errorf("Failed to apply JSON selector '%s' to response: %s, JSON was:\n%s",
			selector, err.Error(), d.body)
	}
	return nil
}

// Apply header extraction
func (d *Displayer) ApplyHeaderExtract(header string) error {
	d.RawOutput = d.response.Header.Get(header)
	if d.RawOutput == "" {
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
	output := d.RawOutput
	if output == nil {
		return ""
	}
	if outputStr, ok := d.RawOutput.(string); ok {
		return outputStr
	}
	var out string
	var err error
	if d.prettify {
		var b []byte
		b, err = json.MarshalIndent(output, "", "    ")
		if err == nil {
			out = string(b) + "\n"
		}
	} else {
		var b []byte
		b, err = json.Marshal(output)
		out = string(b)
	}
	if err != nil {
		return fmt.Sprintf("%v", output)
	}
	return out
}
