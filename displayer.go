package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rightscale/go-jsonselect"
)

// Displayer provides helper methods to display command responses back to the user
// This includes optionally extracting values with JSON:select and pretty-printing
type Displayer struct {
	response  *http.Response
	body      string
	RawOutput interface{}
	prettify  bool
}

// NewDisplayer creates a new displayer using the response body.
func NewDisplayer(resp *http.Response) (*Displayer, error) {
	defer resp.Body.Close()
	js, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response (%s)", err)
	}
	disp := Displayer{response: resp, body: string(js)}
	if len(js) > 2 {
		err = json.Unmarshal(js, &disp.RawOutput)
		if err != nil {
			disp.RawOutput = string(js)
		}
	}
	return &disp, nil
}

// ApplySingleExtract applies the given JSON selector and returns the results.
// It's an error if the selector yields more than one value.
func (d *Displayer) ApplySingleExtract(extract string) error {
	if err := d.ApplyExtract(extract, true); err != nil {
		return err
	}
	outputs := d.RawOutput.([]interface{})
	if len(outputs) != 1 {
		d.RawOutput = nil
		return fmt.Errorf("JSON selector '%s' returned %d instead of one value",
			extract, len(outputs))
	}
	if len(outputs) == 0 {
		d.RawOutput = ""
	} else {
		switch v := outputs[0].(type) {
		case nil:
			d.RawOutput = ""
		case float64, bool:
			d.RawOutput = fmt.Sprint(v)
		case string:
			d.RawOutput = v
		default:
			d.RawOutput = v
		}
		d.RawOutput = outputs[0]
	}
	return nil
}

// ApplyExtract applies selector to js.
func (d *Displayer) ApplyExtract(selector string, js bool) error {
	parser, err := jsonselect.CreateParserFromString(d.body)
	if err != nil {
		return fmt.Errorf("Failed to load response JSON: %s, JSON was:\n%s", err, d.body)
	}
	outputs, err := parser.GetValues(selector)
	if !js {
		out := ""
		for _, o := range outputs {
			b, _ := json.Marshal(o)
			out += string(b) + "\n"
		}
		d.RawOutput = out
	} else {
		d.RawOutput = outputs
	}
	if err != nil {
		return fmt.Errorf("Failed to apply JSON selector '%s' to response: %s, JSON was:\n%s",
			selector, err, d.body)
	}
	return nil
}

// ApplyHeaderExtract reads the value of the given header.
func (d *Displayer) ApplyHeaderExtract(header string) error {
	d.RawOutput = d.response.Header.Get(header)
	if d.RawOutput == "" {
		return fmt.Errorf("Response does not contain the '%s' header", header)
	}
	return nil
}

// Pretty switches the display mode to produce human friendly output.
func (d *Displayer) Pretty() {
	d.prettify = true
}

// Output returns the current output.
func (d *Displayer) Output() string {
	output := d.RawOutput
	if output == nil {
		return ""
	}
	if outputStr, ok := d.RawOutput.(string); ok {
		suffix := ""
		if d.prettify {
			suffix = "\n"
		}
		return outputStr + suffix
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
		fm := "%v"
		if d.prettify {
			fm += "\n"
		}
		return fmt.Sprintf(fm, output)
	}
	return out
}
