package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"bitbucket.org/pkg/inflect"
	"github.com/kr/text"
)

func main() {
	destDefault, _ := filepath.Abs("codegen.go")
	var metadata = flag.String("metadata", "",
		"Path to API 1.5 metadata file, defaults to './api_data.json'")
	var dest = flag.String("output", destDefault,
		"Path to output file, defaults to './codegen.go'")
	flag.Parse()
	if len(*metadata) == 0 {
		check(fmt.Errorf("Specify path to metadata json with '-metadata /path/to/api_data.json'"))
	}
	js, err := ioutil.ReadFile(*metadata)
	check(err)
	var content map[string]interface{}
	err = json.Unmarshal(js, &content)
	f, err := os.Create(*dest)
	check(err)
	generateTopComment(f)
	generateCommonTypes(f)
	for name, resource := range content {
		check(generateResource(name, resource, f))
	}
}

func generateTopComment(dest io.Writer) {
	dest.Write([]byte("/**************************************************************************/\n"))
	dest.Write([]byte("/*                     RightScale API 1.5 go client                       */\n"))
	dest.Write([]byte("/*                                                                        */\n"))
	dest.Write(pad("/* Generated ", time.Now().Format("Jan 2, 2006 at 3:04pm (PST)")))
	dest.Write(pad("/* Command: `api15gen ", strings.Join(os.Args[1:], " "), "`"))
	dest.Write([]byte("/*                                                                        */\n"))
	dest.Write([]byte("/* The content of this file is auto-generated, DO NOT MODIFY              */\n"))
	dest.Write([]byte("/**************************************************************************/\n\n"))
}

func pad(items ...string) []byte {
	l := 0
	buffer := bytes.NewBuffer(make([]byte, 80))
	for _, i := range items {
		buffer.WriteString(i)
		l = l + len(i)
	}
	numSpaces := 74 - l
	if numSpaces < 1 {
		numSpaces = 1
	}
	buffer.WriteString(strings.Repeat(" ", numSpaces))
	buffer.WriteString("*/\n")
	return buffer.Bytes()
}

func generateCommonTypes(dest io.Writer) {
	dest.Write([]byte(`
// Href type
type Href string
`))
}

func generateResource(name string, resource interface{}, dest io.Writer) error {
	res := resource.(map[string]interface{})
	dest.Write([]byte(fmt.Sprintf("\n// == %s ==\n\n", name)))
	methods := res["methods"].(map[string]interface{})
	var mediaType map[string]interface{}
	if raw, _ := res["media_type"]; raw != nil {
		mediaType = raw.(map[string]interface{})
	}
	for n, m := range methods {
		//fmt.Printf("GENERATING %s of %s\n", n, name)
		kind := n
		n = methodName(n, name)
		meth := m.(map[string]interface{})
		params := meth["parameters"].(map[string]interface{})
		body := generateBody(kind, mediaType)
		description := "<no description>"
		if d, _ := meth["description"]; d != nil {
			description = d.(string)
		}
		generateHeader(n, description, meth["route"].(string), dest)
		generateSignature(n, name, params, kind, dest)
		dest.Write(body)
		dest.Write([]byte("}\n"))
		generateGenericSignature(n, name, kind, dest)
		dest.Write(body)
		dest.Write([]byte("}\n\n"))
	}
	return nil
}

func generateBody(kind string, mediaType map[string]interface{}) []byte {
	return []byte{}
}

func generateHeader(name string, description string, route string, dest io.Writer) {
	dest.Write([]byte("// "))
	dest.Write(parseRoute(route))
	// There's a bug in the text package that will cause an infinite loop if the number below
	// is less than 104 (has to do with the comment on the 'create CloudAccount' action)
	dest.Write([]byte(text.Indent(text.Wrap(description, 104), "// ")))
	dest.Write([]byte("\n"))
}

func generateSignature(name, resName string, params map[string]interface{}, kind string, dest io.Writer) {
	dest.Write([]byte("func (c *Client) "))
	dest.Write([]byte(name))
	dest.Write(parseParams(params))
	dest.Write(parseReturn(kind, resName))
	dest.Write([]byte(" {\n"))
}

func generateGenericSignature(name, resName, kind string, dest io.Writer) {
	dest.Write([]byte("func (c *Client) "))
	dest.Write([]byte(name))
	dest.Write([]byte("(p *Params)"))
	dest.Write(parseReturn(kind, resName))
	dest.Write([]byte(" {\n"))
}

func parseRoute(route string) []byte {
	if len(route) < 7 {
		return []byte("<empty route>") // :(((( some routes are empty
	}
	verb := strings.TrimRight(route[0:7], " ")
	path := strings.TrimRight(strings.Split(route[8:], "{")[0], " ")
	buffer := bytes.NewBuffer(make([]byte, 128))
	buffer.WriteString(verb)
	buffer.WriteString(" ")
	buffer.WriteString(path)
	buffer.WriteString("\n")
	return buffer.Bytes()
}

func parseParams(params map[string]interface{}) []byte {
	buffer := bytes.NewBuffer(make([]byte, 256))
	buffer.WriteString("(")
	idx, max := 0, len(params)
	for name, param := range params {
		p := param.(map[string]interface{})
		class, ok := p["class"]
		if !ok {
			// Assume string, seems to happen e.g. with parameter 'limit' of audit entries index
			class = "String"
		}
		switch class { // Possible values: "String", "Hash", "Array", "Integer", "Enumerable"
		case "Integer":
			buffer.Write(parseParamName(name))
			buffer.WriteString(" int")
		case "String":
			buffer.Write(parseParamName(name))
			buffer.WriteString(" string")
		case "Hash":
			continue // skip - one parameter per field
		case "Array":
			buffer.Write(parseParamName(name))
			buffer.WriteString(" []string")
		case "Enumerable": // inputs
			buffer.Write(parseParamName(name))
			buffer.WriteString(" map[string]string")
		}
		idx += 1
		if idx < max {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString(")")
	return buffer.Bytes()
}

func parseParamName(name string) []byte {
	r := regexp.MustCompile("[^[:alnum:]]+")
	p := r.ReplaceAllString(name, "_")
	return []byte(inflect.CamelizeDownFirst(p))
}

func parseReturn(kind, resName string) []byte {
	switch kind {
	case "show":
		b := bytes.NewBuffer(make([]byte, len(resName)+11))
		b.WriteString(" (*")
		b.WriteString(resName)
		b.WriteString(", error) ")
		return b.Bytes()
	case "index":
		b := bytes.NewBuffer(make([]byte, len(resName)+11))
		b.WriteString(" ([]")
		b.WriteString(resName)
		b.WriteString(", error) ")
		return b.Bytes()
	case "create":
		return []byte(" (href, error) ")
	case "update", "destroy":
		return []byte(" error ")
	default:
		return []byte(" ") // TBD: Look at media type
	}

}

// Heuristic to create method name from action and resource names
func methodName(action, resource string) string {
	if action != "index" && strings.Index(action, "multi") != 0 {
		resource = inflect.Singularize(resource)
	}
	return inflect.Camelize(action) + inflect.Camelize(resource)
}

// Panic if error is not nil
func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "api15gen: %s\n", err.Error())
		os.Exit(1)
	}
}
