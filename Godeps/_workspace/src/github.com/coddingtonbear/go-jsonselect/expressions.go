package jsonselect

import (
    "strings"
)

type exprElement struct {
    value interface{}
    typ jsonType
}

var precedenceMap = map[string]int{
    "*": 1,
    "/": 1,
    "%": 1,
    "+": 2,
    "-": 2,
    "<=": 3,
    "<": 3,
    ">": 3,
    "$=": 3,
    "^=": 3,
    "*=": 3,
    "=": 3,
    "!=": 3,
    "&&": 4,
    "||": 5,
}

var comparatorMap = map[string]func(lhs exprElement, rhs exprElement)exprElement{
    "*": func(lhs exprElement, rhs exprElement)exprElement {
        value := getFloat64(lhs.value) * getFloat64(rhs.value)
        return exprElement{value, J_NUMBER}
    },
    "/": func(lhs exprElement, rhs exprElement)exprElement {
        value := getFloat64(lhs.value) / getFloat64(rhs.value)
        return exprElement{value, J_NUMBER}
    },
    "%": func(lhs exprElement, rhs exprElement)exprElement {
        value := float64(getInt32(lhs.value) % getInt32(rhs.value))
        return exprElement{value, J_NUMBER}
    },
    "+": func(lhs exprElement, rhs exprElement)exprElement {
        value := getFloat64(lhs.value) + getFloat64(rhs.value)
        return exprElement{value, J_NUMBER}
    },
    "-": func(lhs exprElement, rhs exprElement)exprElement {
        value := getFloat64(lhs.value) - getFloat64(rhs.value)
        return exprElement{value, J_NUMBER}
    },
    "<=": func(lhs exprElement, rhs exprElement)exprElement {
        if getFloat64(lhs.value) <= getFloat64(rhs.value) {
            return exprElement{true, J_BOOLEAN}
        }
        return exprElement{false, J_BOOLEAN}
    },
    "<": func(lhs exprElement, rhs exprElement)exprElement {
        if getFloat64(lhs.value) < getFloat64(rhs.value) {
            return exprElement{true, J_BOOLEAN}
        }
        return exprElement{false, J_BOOLEAN}
    },
    ">=": func(lhs exprElement, rhs exprElement)exprElement {
        if getFloat64(lhs.value) > getFloat64(rhs.value) {
            return exprElement{true, J_BOOLEAN}
        }
        return exprElement{false, J_BOOLEAN}
    },
    ">": func(lhs exprElement, rhs exprElement)exprElement {
        if getFloat64(lhs.value) > getFloat64(rhs.value) {
            return exprElement{true, J_BOOLEAN}
        }
        return exprElement{false, J_BOOLEAN}
    },
    "$=": func(lhs exprElement, rhs exprElement)exprElement {
        lhs_str := getJsonString(lhs.value)
        rhs_str := getJsonString(rhs.value)
        if strings.HasSuffix(lhs_str, rhs_str) {
            return exprElement{true, J_BOOLEAN}
        }
        return exprElement{false, J_BOOLEAN}
    },
    "^=": func(lhs exprElement, rhs exprElement)exprElement {
        lhs_str := getJsonString(lhs.value)
        rhs_str := getJsonString(rhs.value)
        if strings.Index(lhs_str, rhs_str) == 0 {
            return exprElement{true, J_BOOLEAN}
        }
        return exprElement{false, J_BOOLEAN}
    },
    "*=": func(lhs exprElement, rhs exprElement)exprElement {
        lhs_str := getJsonString(lhs.value)
        rhs_str := getJsonString(rhs.value)
        if strings.Index(lhs_str, rhs_str) != -1 {
            return exprElement{true, J_BOOLEAN}
        }
        return exprElement{false, J_BOOLEAN}
    },
    "=": func(lhs exprElement, rhs exprElement)exprElement {
        lhs_string := getJsonString(lhs.value)
        rhs_string := getJsonString(rhs.value)
        if lhs_string == rhs_string {
            return exprElement{true, J_BOOLEAN}
        }
        return exprElement{false, J_BOOLEAN}
    },
    "!=": func(lhs exprElement, rhs exprElement)exprElement {
        lhs_string := getJsonString(lhs.value)
        rhs_string := getJsonString(rhs.value)
        if lhs_string != rhs_string {
            return exprElement{true, J_BOOLEAN}
        }
        return exprElement{false, J_BOOLEAN}
    },
    "&&": func(lhs exprElement, rhs exprElement)exprElement {
        if lhs.value.(bool) && rhs.value.(bool) {
            return exprElement{true, J_BOOLEAN}
        }
        return exprElement{false, J_BOOLEAN}
    },
    "||": func(lhs exprElement, rhs exprElement)exprElement {
        if lhs.value.(bool) || rhs.value.(bool) {
            return exprElement{true, J_BOOLEAN}
        }
        return exprElement{false, J_BOOLEAN}
    },
}

func (p *Parser) evaluateExpressionWithPrecedence(elements []*exprElement, precedenceLevel int) []*exprElement {
    var newExpression []*exprElement
    var i int
    var element *exprElement
    for i = 0; i < len(elements); i++ {
        element = elements[i]
        if element.typ == J_OPER && precedenceMap[element.value.(string)] == precedenceLevel {
            lhs := newExpression[len(newExpression)-1]
            rhs := elements[i+1]
            newExpression = newExpression[0:len(newExpression)-1]
            if exprElementsMatch(*lhs, *rhs) {
                result := comparatorMap[element.value.(string)](*lhs, *rhs)
                newExpression = append(newExpression, &result)
            } else {
                logger.Print("Cannot compare ", lhs.value, " and ", rhs.value, "; types differ: ", rhs.typ, " != ", lhs.typ)
                newExpression = append(newExpression, &exprElement{false, J_BOOLEAN})
            }
            // Skip ahead an additional step more than normal
            // since we just pulled off one more element than normal.
            i += 1
        } else {
            newExpression = append(newExpression, element)
        }
    }
    return newExpression
}

func (p *Parser) evaluateExpression(elements []*exprElement) exprElement {
    for i := 1; i < 5; i++ {
        elements = p.evaluateExpressionWithPrecedence(elements, i)
    }
    if len(elements) > 1 {
        panic("More than one expression result")
    }
    return *elements[0]
}

func (p *Parser) parseExpression(tokens []*token, node *jsonNode) exprElement {
    var finalTokens []*exprElement

    logger.Print("Parsing expression ", getFormattedTokens(tokens))

    // If this expression is wrapped with parentheses, let's remove them.
    if tokens[0].typ == S_PAREN && tokens[0].val == "(" {
        tokens = tokens[1:len(tokens)-1]
    }

    for i := 0; i < len(tokens); i++ {
        thisToken := *tokens[i]
        if thisToken.typ == S_PAREN && thisToken.val == "(" {
            var leftCount = 0
            var rightCount = 0
            var j = 0
            var subExprToken *token
            for j, subExprToken = range tokens[i:] {
                if subExprToken.typ == S_PAREN && subExprToken.val == "(" {
                    leftCount++
                } else if subExprToken.typ == S_PAREN && subExprToken.val == ")" {
                    rightCount++
                }
                if leftCount == rightCount {
                    break
                }
            }
            subexprResult := p.parseExpression(tokens[i:i+j+1], node)
            finalTokens = append(finalTokens, &subexprResult)
            // Let's move the cursor to the end of the subexpression above.
            i = i + j
        } else if thisToken.typ == S_PVAR {
            // This is for the value of the node being processed.
            finalTokens = append(finalTokens, &exprElement{node.value, node.typ})
        } else if thisToken.typ == S_STRING || thisToken.typ == S_BOOL || thisToken.typ == S_NIL || thisToken.typ == S_NUMBER {
            // Let's copy these kinds of tokens down as expression elements.
            var jType jsonType
            switch thisToken.typ {
                case S_BOOL:
                    jType = J_BOOLEAN
                case S_NIL:
                    jType = J_NULL
                case S_NUMBER:
                    jType = J_NUMBER
                default:
                    jType = J_STRING
            }
            finalTokens = append(finalTokens, &exprElement{thisToken.val, jType})
        } else if thisToken.typ == S_BINOP {
            finalTokens = append(finalTokens, &exprElement{thisToken.val, J_OPER})
        } else {
            logger.Print("Unexpected token ", thisToken, " ignored.")
        }
    }

    return p.evaluateExpression(finalTokens)
}
