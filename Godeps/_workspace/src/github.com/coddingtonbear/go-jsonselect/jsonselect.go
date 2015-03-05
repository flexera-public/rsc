package jsonselect

import (
    "errors"
    "io/ioutil"
    "log"
    "regexp"
    "strconv"
    "strings"
    "github.com/coddingtonbear/go-simplejson"
)

type Parser struct {
    Data *simplejson.Json
    nodes []*jsonNode
}

func CreateParserFromString(body string) (*Parser, error) {
    json, err := simplejson.NewJson([]byte(body))
    if err != nil {
        return nil, err
    }
    parser := Parser{json, nil}
    parser.mapDocument()
    return &parser, err
}

func CreateParser(json *simplejson.Json) (*Parser, error) {
    log.SetOutput(ioutil.Discard)
    parser := Parser{json, nil}
    parser.mapDocument()
    return &parser, nil
}

func (p *Parser) evaluateSelector(selector string) ([]*jsonNode, error) {
    tokens, err := lex(selector, selectorScanner)
    if err != nil {
        return nil, err
    }

    nodes, err := p.selectorProduction(tokens, p.nodes, 1)
    if err != nil {
        return nil, err
    }

    logger.Print(len(nodes), " matches found")

    return nodes, nil
}

func (p *Parser) GetJsonElements(selector string) ([]*simplejson.Json, error) {
    nodes, err := p.evaluateSelector(selector)
    if err != nil {
        return nil, err
    }

    var results = make([]*simplejson.Json, 0, len(nodes))
    for _, node := range nodes {
        results = append(
            results,
            node.json,
        )
    }
    return results, nil
}

func (p *Parser) GetValues(selector string) ([]interface{}, error) {
    nodes, err := p.evaluateSelector(selector)
    if err != nil {
        return nil, err
    }

    var results = make([]interface{}, 0, len(nodes))
    for _, node := range nodes {
        results = append(
            results,
            node.value,
        )
    }

    return results, nil
}

func (p *Parser) selectorProduction(tokens []*token, documentMap []*jsonNode, recursionDepth int) ([]*jsonNode, error) {
    var results []*jsonNode
    var matched bool
    var value interface{}
    var validator func(*jsonNode)bool
    var validators = make([]func(*jsonNode)bool, 0, 10)
    logger.Print("selectorProduction(", recursionDepth, ") starting with ", tokens[0], " - ", len(tokens), " tokens remaining.")

    _, matched, _ = p.peek(tokens, S_TYPE)
    if matched {
        value, tokens, _ = p.match(tokens, S_TYPE)
        validators = append(
            validators,
            p.typeProduction(value),
        )
    }
    _, matched, _ = p.peek(tokens, S_IDENTIFIER)
    if matched {
        value, tokens, _ = p.match(tokens, S_IDENTIFIER)
        validators = append(
            validators,
            p.keyProduction(value),
        )
    }
    _, matched, _ = p.peek(tokens, S_PCLASS)
    if matched {
        value, tokens, _ = p.match(tokens, S_PCLASS)
        validators = append(
            validators,
            p.pclassProduction(value),
        )
    }
    _, matched, _ = p.peek(tokens, S_NTH_FUNC)
    if matched {
        value, tokens, _ = p.match(tokens, S_NTH_FUNC)
        validator, tokens = p.nthChildProduction(value, tokens)
        validators = append(validators, validator)
    }
    _, matched, _ = p.peek(tokens, S_PCLASS_FUNC)
    if matched {
        value, tokens, _ = p.match(tokens, S_PCLASS_FUNC)
        validator, tokens = p.pclassFuncProduction(value, tokens, documentMap)
        validators = append(validators, validator)
    }
    result, matched, _ := p.peek(tokens, S_OPER)
    if matched && result.(string) == "*" {
        value, tokens, _ = p.match(tokens, S_OPER)
        validator = p.universalProduction(value)
        validators = append(validators, validator)
    }

    if len(validators) < 1 {
        return nil, errors.New("No selector recognized")
    }

    results, err := p.matchNodes(validators, documentMap)
    if err != nil {
        return nil, err
    }
    logger.Print("Applying ", len(validators), " validators to document resulted in ", len(results), " matches")

    _, matched, _ = p.peek(tokens, S_OPER)
    if matched {
        value, tokens, _ = p.match(tokens, S_OPER)
        logger.Print("Recursing selectorProduction(", recursionDepth, ") via operator ", value, " starting with ", tokens[0], " ;", len(tokens), " tokens remaining")
        logger.IncreaseDepth()
        rvals, err := p.selectorProduction(tokens, documentMap, recursionDepth+1)
        logger.DecreaseDepth()
        logger.Print("Recursion completed; returned control to selectorProduction(", recursionDepth, ") via operator ", value, " with ", len(rvals), " matches.")
        if err != nil {
            return nil, err
        }
        switch value {
            case ",":
                logger.Print("(", recursionDepth, ") Operator ',': ", len(results), " => ", len(results) + len(rvals))
                for _, val := range rvals {
                    // TODO: This is quite slow
                    // it seems like it's probably quite easy to expand
                    // the list just once
                    results = append(results, val)
                }
            case ">":
                originalLength := len(results)
                results = parents(results, rvals)
                logger.Print("(", recursionDepth, ") Operator '>': ", originalLength, " => ", len(results))
            case "~":
                originalLength := len(results)
                results = siblings(results, rvals)
                logger.Print("(", recursionDepth, ") Operator '~': ", originalLength, " => ", len(results))
            case " ":
                originalLength := len(results)
                results = ancestors(results, rvals)
                logger.Print("(", recursionDepth, ") Operator ' ': ", originalLength, " => ", len(results))
            default:
                return nil, errors.New("Unrecognized operator")
        }
    } else if len(tokens) > 0 {
        logger.Print("Recursing selectorProduction(", recursionDepth, ") for excess tokens starting with ", tokens[0], " ;", len(tokens), " tokens remaining")
        logger.IncreaseDepth()
        rvals, err := p.selectorProduction(tokens, documentMap, recursionDepth+1)
        logger.DecreaseDepth()
        logger.Print("Recursion completed; returned control to selectorProduction(", recursionDepth, ") for excess tokens with ", len(rvals), " matches.")
        if err != nil {
            return nil, err
        }
        results = ancestors(results, rvals)
    }

    logger.Print("selectorProduction(", recursionDepth, ") returning ", len(results), " matches.")
    return results, nil
}

func (p *Parser) peek(tokens []*token, typ tokenType) (interface{}, bool, error) {
    if len(tokens) < 1 {
        return nil, false, errors.New("No more tokens")
    }
    if tokens[0].typ == typ {
        return tokens[0].val, true, nil
    }
    return nil, false, nil
}

func (p *Parser) match(tokens []*token, typ tokenType) (interface{}, []*token, error) {
    value, matched, _ := p.peek(tokens, typ)
    if !matched {
        return nil, tokens, errors.New("Match not successful")
    }
    _, tokens = tokens[0], tokens[1:]
    return value, tokens, nil
}

func (p *Parser) matchNodes(validators []func(*jsonNode)bool, documentMap []*jsonNode) ([]*jsonNode, error) {
    var matches []*jsonNode
    nodeCount := 0
    if logger.Enabled {
        nodeCount = len(documentMap)
    }
    for idx, node := range documentMap {
        var passed = true
        if logger.Enabled {
            logger.SetPrefix("[Node ", idx, "/", nodeCount, "] ")
        }
        for _, validator := range validators {
            if !validator(node) {
                passed = false
                break
            }
        }
        if passed {
            logger.Print("MATCHED: ", node)
            matches = append(matches, node)
        }
    }
    return matches, nil
}

func (p *Parser) typeProduction(value interface{}) func(*jsonNode)bool {
    logger.Print("Creating typeProduction validator ", value)
    return func(node *jsonNode) bool {
        logger.Print("typeProduction ? ", node.typ, " == ", value)
        return string(node.typ) == value
    }
}

func (p *Parser) keyProduction(value interface{}) func(*jsonNode)bool {
    logger.Print("Creating keyProduction validator ", value)
    return func(node *jsonNode) bool {
        logger.Print("keyProduction ? ", node.parent_key, " == ", value)
        if node.parent_key == ""{
            return false
        }
        return string(node.parent_key) == value
    }
}

func (p *Parser) universalProduction(value interface{}) func(*jsonNode)bool {
    operator := value.(string)
    if operator == "*" {
        return func(node *jsonNode) bool {
            logger.Print("universalProduction ? true")
            return true
        }
    } else {
        logger.Print("Error: Unexpected operator: ", operator)
        return func(node *jsonNode) bool {
            logger.Print("Asserting false due to failed universalProduction")
            return false
        }
    }
}

func (p *Parser) pclassProduction(value interface{}) func(*jsonNode)bool {
    pclass := value.(string)
    logger.Print("Creating pclassProduction validator ", pclass)
    if pclass == "first-child" {
        return func(node *jsonNode) bool {
            logger.Print("pclassProduction first-child ? ", node.idx, " == 1")
            return node.idx == 1
        }
    } else if pclass == "last-child" {
        return func(node *jsonNode) bool {
            logger.Print("pclassProduction last-child ? ", node.siblings, " > 0 AND ", node.idx, " == ", node.siblings)
            return node.siblings > 0 && node.idx == node.siblings
        }
    } else if pclass == "only-child" {
        return func(node *jsonNode) bool {
            logger.Print("pclassProduction ony-child ? ", node.siblings, " == 1")
            return node.siblings == 1
        }
    } else if pclass == "root" {
        return func(node *jsonNode) bool {
            logger.Print("pclassProduction root ? ", node.parent, " == nil")
            return node.parent == nil
        }
    } else if pclass == "empty" {
        return func(node *jsonNode) bool {
            logger.Print("pclassProduction empty ? ", node.typ, " == ", J_ARRAY, " AND ", len(node.value.(string)), " < 1")
            return node.typ == J_ARRAY && len(node.value.(string)) < 1
        }
    }
    logger.Print("Error: Unknown pclass: ", pclass)
    return func(node *jsonNode) bool {
        logger.Print("Asserting false due to failed pclassProduction")
        return false
    }
}

func (p *Parser) nthChildProduction(value interface{}, tokens []*token) (func(*jsonNode)bool, []*token) {
    nthChildRegexp := regexp.MustCompile(`^\s*\(\s*(?:([+\-]?)([0-9]*)n\s*(?:([+\-])\s*([0-9]))?|(odd|even)|([+\-]?[0-9]+))\s*\)`)
    args, tokens, _ := p.match(tokens, S_EXPR)
    var a int
    var b int
    var reverse bool = false

    pattern := nthChildRegexp.FindStringSubmatch(args.(string))

    logger.Print("Creating nthChildProduction validator ", pattern)
    if logger.Enabled {
        for idx, pat := range pattern {
            logger.Print("[", idx, "] ", pat)
        }
    }

    if pattern[5] != "" {
        a = 2
        if pattern[5] == "odd" {
            b = 1
        } else {
            b = 0
        }
    } else if pattern[6] != ""{
        a = 0
        b_parsed, _ := strconv.ParseInt(pattern[6], 10, 64)
        b = int(b_parsed)
    } else {
        // Expression like +n-3
        sign := "+"
        if pattern[1] != "" {
            sign = pattern[1]
        }
        coeff := "1"
        if pattern[2] != "" {
            coeff = pattern[2]
        }
        a_parsed, _ := strconv.ParseInt(coeff, 10, 64)
        a = int(a_parsed)
        if sign == "-" {
            a = -1 * a
        }

        if pattern[3] != "" {
            b_sign := pattern[3]
            b_parsed, _ := strconv.ParseInt(pattern[4], 10, 64)
            b = int(b_parsed)
            if b_sign == "-" {
                b = -1 * b
            }
        } else {
            b = 0
        }
    }

    if value.(string) == "nth-last-child" {
        reverse = true
    }

    return func(node *jsonNode)bool {
        var b_str string
        if b > 0 {
            b_str = "+"
        } else {
            b_str = "-"
        }
        logger.Print("nthChildProduction ", a, "n", b_str, b)
        logger.Print("nthChildProduction ? ", node.siblings, " == 0")
        if node.siblings == 0 {
            return false
        }

        idx := node.idx - 1
        if reverse {
            idx = node.siblings - idx
        } else {
            idx++
        }

        logger.Print("nthChildProduction (continued-1) ? ", a, " == 0")
        if a == 0 {
            return b == idx
        } else {
            logger.Print("nthChildProduction (continued-2) ? ", idx-b % a, " == 0 AND ", idx*a+b, " >= 0")
            return ((idx - b) % a) == 0 && (idx * a + b) >= 0
        }
    }, tokens
}

func (p *Parser) pclassFuncProduction(value interface{}, tokens []*token, documentMap []*jsonNode) (func(*jsonNode)bool, []*token) {
    sargs, tokens, _ := p.match(tokens, S_EXPR)
    pclass := value.(string)

    logger.Print("Creating pclassFuncProduction validator ", pclass)

    if pclass == "expr" {
        tokens, err := lex(sargs.(string), expressionScanner)
        if err != nil {
            panic(err)
        }
        var tokens_to_return []*token
        return func(node *jsonNode)bool {
            result := p.parseExpression(tokens, node)
            logger.Print("pclassFuncProduction expr ? ", result)
            return exprElementIsTruthy(result)
        }, tokens_to_return
    }

    lexString := sargs.(string)[1:len(sargs.(string))-1]
    args, _ := lex(lexString, selectorScanner)

    logger.Print("pclassFuncProduction lex results for [", lexString, "]: (follow)")
    if logger.Enabled {
        for i, arg := range args {
            logger.Print("[", i, "]: ", arg)
        }
    }

    if pclass == "has" {
        return func(node *jsonNode)bool {
            newMap := p.getFlooredDocumentMap(node)
            logger.Print("pclassFuncProduction recursing into selectorProduction(-100) starting with ", args[0], "; ", len(args), " tokens remaining.")
            logger.IncreaseDepth()
            rvals, _ := p.selectorProduction(args, newMap, -100)
            logger.DecreaseDepth()
            logger.Print("pclassFuncProduction resursion completed with ", len(rvals), " results.")
            ancestors := make(map[*simplejson.Json]*jsonNode, len(rvals))
            for _, node := range rvals {
                if node.parent != nil {
                    ancestors[node.parent.json] = node.parent
                }
            }
            logger.Print("pclassFuncProduction has ? ", node, " ∈ ", getFormattedNodeMap(ancestors))
            return nodeIsMemberOfHaystack(node, ancestors)
        }, tokens
    } else if pclass == "contains" {
        return func(node *jsonNode)bool {
            logger.Print("pclassFuncProduction contains ? ", node.typ, " == ", J_STRING, " AND ", strings.Count(node.value.(string), args[0].val.(string)), " > 0")
            return node.typ == J_STRING && strings.Count(node.value.(string), args[0].val.(string)) > 0
        }, tokens
    } else if pclass == "val" {
        return func(node *jsonNode)bool {
            lhsString := getJsonString(node.value)
            rhsString := getJsonString(args[0].val)
            logger.Print("pclassFuncProduction val ? ", lhsString, " == ", rhsString)
            return lhsString == rhsString
        }, tokens
    }

    // If we didn't find a known pclass, do not match anything.
    logger.Print("Error: Unknown pclass: ", pclass)
    return func(node *jsonNode)bool {
        logger.Print("Asserting false due to failed pclassFuncProduction")
        return false
    }, tokens
}
