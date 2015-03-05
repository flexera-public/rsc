package jsonselect

import (
    "io/ioutil"
    "reflect"
    "strings"
    "testing"
    "github.com/coddingtonbear/go-simplejson"
)

// Used for storing the results of the benchmarking tests below
// to avoid compiler optimizations.
var parser *Parser
var values []interface{}

func getTestParser(testDocuments map[string]*simplejson.Json, testName string) (*Parser, error) {
    jsonDocument := testDocuments[testName[0:strings.Index(testName, "_")]]
    return CreateParser(jsonDocument)
}

func runTestsInDirectory(t *testing.T, baseDirectory string) {
    var testDocuments = make(map[string]*simplejson.Json)
    var testSelectors = make(map[string]string)
    var testOutput = make(map[string][]string)

    files, err := ioutil.ReadDir(baseDirectory)
    if err != nil {
        t.Error("Error encountered while loading conformance tests ", err)
    }

    for _, fileInfo := range(files) {
        name := fileInfo.Name()
        if strings.HasSuffix(name, ".json") {
            json_document, err := ioutil.ReadFile(baseDirectory + name)
            if err != nil {
                t.Error("Error encountered while reading ", name, ": ", err)
                continue
            }
            parsed_document, err := simplejson.NewJson(json_document)
            if err != nil {
                t.Error("Error encountered while deserializing ", name, ": ", err)
                continue
            }
            testDocuments[name[0:len(name)-len(".json")]] = parsed_document
        } else if (strings.HasSuffix(name, ".output")) {
            output_document, err := ioutil.ReadFile(baseDirectory + name)
            if err != nil {
                t.Error("Error encountered while reading ", name, ": ", err)
                continue
            }
            if strings.HasPrefix(string(output_document), "Error") {
                // We won't be handling errors in the same way.
                continue
            }
            var actualOutput []string
            var stringTemporary string
            for _, str := range strings.Split(string(output_document), "\n") {
                stringTemporary = stringTemporary + str
                // Try to parse -- if it works, we have the whole object
                if strings.Index(stringTemporary, "{") == 0 {
                    if strings.Count(stringTemporary, "{") != strings.Count(stringTemporary, "}") {
                        continue
                    }
                    actualOutput = append(actualOutput, stringTemporary)
                    stringTemporary = ""
                } else if strings.Index(stringTemporary, "[") == 0 {
                    if strings.Count(stringTemporary, "[") != strings.Count(stringTemporary, "]") {
                        continue
                    }
                    actualOutput = append(actualOutput, stringTemporary)
                    stringTemporary = ""
                } else if len(stringTemporary) > 0 {
                    actualOutput = append(actualOutput, stringTemporary)
                    stringTemporary = ""
                }
            }
            testOutput[name[0:len(name)-len(".output")]] = actualOutput
        } else if (strings.HasSuffix(name, ".selector")) {
            selector_document, err := ioutil.ReadFile(baseDirectory + name)
            if err != nil {
                t.Error("Error encountered while reading ", name, ": ", err)
                continue
            }
            testSelectors[name[0:len(name)-len(".selector")]] = string(selector_document)
        }
    }

    for testName := range testOutput {
        var passed bool = true
        t.Log("Running test ", testName)
        parser, err := getTestParser(testDocuments, testName)
        if err != nil {
            t.Error("Test ", testName, "failed: ", err)
            passed = false
        }
        selectorString := testSelectors[testName]
        expectedOutput := testOutput[testName]

        results, err := parser.GetJsonElements(selectorString)
        if err != nil {
            t.Error("Test ", testName, "failed: ", err)
            passed = false
        }
        var stringResults []string
        for _, result := range results {
            encoded, err := result.Encode()
            if err != nil {
                t.Error("Test ", testName, "failed: ", err)
                passed = false
            }
            stringResults = append(stringResults, string(encoded))
        }

        if len(stringResults) != len(expectedOutput) {
            t.Error("Test ", testName, " failed due to number of results being mismatched; ", len(stringResults), " != ", len(expectedOutput), ": [Actual] ", stringResults, " != [Expected] ", expectedOutput)
            passed = false
        } else {
            var expected = make([]*simplejson.Json, 0, 10)
            var actual = make([]*simplejson.Json, 0, 10)
            matchType := "string"

            for idx, result := range stringResults {
                expectedEncoded := expectedOutput[idx]

                expectedJson, err := simplejson.NewJson([]byte(expectedEncoded))
                if err != nil {
                    t.Error(
                        "Test ", testName, " failed due to a JSON decoding error while decoding expectation: ", err,
                    )
                    passed = false
                }
                resultJson, err := simplejson.NewJson([]byte(result))
                if err != nil {
                    t.Error(
                        "Test ", testName, " failed due to a JSON decoding error while decoding result: ", err,
                    )
                    passed = false
                }

                expected = append(expected, expectedJson)
                actual = append(actual, resultJson)

                if strings.Index(strings.TrimSpace(expectedEncoded), "{") == 0 {
                    matchType = "map"
                } else if strings.Index(strings.TrimSpace(expectedEncoded), "[") == 0 {
                    matchType = "array"
                } else if expectedEncoded != result {
                    matchType = "string"
                }
            }

            // Iterate over each of the actual elements; if:
            // * We find a match, remove the match from the expected.
            // * We do not find a match, report an error
            for _, actualElement := range actual {
                var matched bool = false
                for expectedIdx, expectedElement := range expected {
                    if matchType == "map" {
                        matched = reflect.DeepEqual(
                            expectedElement.MustMap(),
                            actualElement.MustMap(),
                        )
                    } else if matchType == "array" {
                        matched = reflect.DeepEqual(
                            expectedElement.MustArray(),
                            actualElement.MustArray(),
                        )
                    } else if matchType == "string" {
                        matched = reflect.DeepEqual(
                            expectedElement.MustString(),
                            actualElement.MustString(),
                        )
                    }
                    if matched {
                        expected = append(
                            expected[:expectedIdx],
                            expected[expectedIdx+1:]...
                        )
                        break
                    }
                }
                if !matched {
                    t.Error("Actual element", actualElement, "not found in expected.")
                    passed = false
                    break;
                }
            }
            if len(expected) > 0 {
                for _, value := range expected {
                    t.Error("Expected element", value, "not found in actual.")
                }
                passed = false
            }
        }
        if passed {
            t.Log("Test ", testName, " PASSED")
        } else {
            t.Error("Test ", testName, " FAILED")
        }
    }
}

func TestLevel1(t *testing.T) {
    runTestsInDirectory(t, "./conformance_tests/level_1/")
}

func TestLevel2(t *testing.T) {
    runTestsInDirectory(t, "./conformance_tests/level_2/")
}

func TestLevel3(t *testing.T) {
    runTestsInDirectory(t, "./conformance_tests/level_3/")
}

func BenchmarkParseDocument(b *testing.B) {
    json_ast, _ := ioutil.ReadFile("./test_data/example_json_ast.json")
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        parser, _ = CreateParserFromString(string(json_ast))
    }
}

func BenchmarkBasicSelector(b *testing.B) {
    json_ast, _ := ioutil.ReadFile("./test_data/example_json_ast.json")
    parser, _ = CreateParserFromString(string(json_ast))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        values, _ = parser.GetValues(`.Link`)
    }
}

func BenchmarkComplexSelector(b *testing.B) {
    json_ast, _ := ioutil.ReadFile("./test_data/example_json_ast.json")
    parser, _ = CreateParserFromString(string(json_ast))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        values, _ = parser.GetValues(`.Link object:has(.Str:val("News"))`)
    }
}

func BenchmarkAncestorSelector(b *testing.B) {
    json_ast, _ := ioutil.ReadFile("./test_data/example_json_ast.json")
    parser, _ = CreateParserFromString(string(json_ast))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        values, _ = parser.GetValues(`.Link object`)
    }
}

func BenchmarkHasExpression(b *testing.B) {
    json_ast, _ := ioutil.ReadFile("./test_data/example_json_ast.json")
    parser, _ = CreateParserFromString(string(json_ast))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        values, _ = parser.GetValues(`object:has(.Str:val("News"))`)
    }
}
