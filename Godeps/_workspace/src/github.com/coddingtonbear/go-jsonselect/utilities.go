package jsonselect

import (
    "encoding/json"
    "log"
    "strconv"
    "github.com/coddingtonbear/go-simplejson"
)

func nodeIsMemberOfHaystack(needle *jsonNode, haystack map[*simplejson.Json]*jsonNode) bool {
    _, ok := haystack[needle.json]
    return ok
}

func nodeIsMemberOfList(needle *jsonNode, haystack []*jsonNode) bool {
    for _, element := range haystack {
        if element.json == needle.json {
            return true
        }
    }
    return false
}

func appendAncestorsToHaystack(node *jsonNode, haystack map[*simplejson.Json]*jsonNode) {
    if node.parent != nil {
        haystack[node.parent.json] = node.parent
        appendAncestorsToHaystack(node.parent, haystack)
    }
}

func nodeIsChildOfHaystackMember(needle *jsonNode, haystack map[*simplejson.Json]*jsonNode) bool {
    if nodeIsMemberOfHaystack(needle, haystack) {
        return true
    }
    if needle.parent == nil {
        return false
    }
    return nodeIsChildOfHaystackMember(needle.parent, haystack)
}

func parents(lhs []*jsonNode, rhs []*jsonNode) []*jsonNode {
    var results []*jsonNode

    lhsHaystack := getHaystackFromNodeList(lhs)

    for _, element := range rhs {
        if nodeIsMemberOfHaystack(element.parent, lhsHaystack) {
            results = append(results, element)
        }
    }

    return results
}

func ancestors(lhs []*jsonNode, rhs []*jsonNode) []*jsonNode {
    var results []*jsonNode
    haystack := getHaystackFromNodeList(lhs)

    for _, element := range rhs {
        if nodeIsChildOfHaystackMember(element, haystack) {
            results = append(results, element)
        }
    }

    return results
}

func siblings(lhs []*jsonNode, rhs []*jsonNode) []*jsonNode {
    var results []*jsonNode
    parents := make(map[*simplejson.Json]*jsonNode, len(lhs))

    for _, element := range lhs {
        parents[element.parent.json] = element.parent
    }

    for _, element := range rhs {
        if nodeIsMemberOfHaystack(element.parent, parents){
            results = append(results, element)
        }
    }

    return results
}

func getFloat64(in interface{}) float64 {
    as_float, ok := in.(float64)
    if ok {
        return as_float
    }
    as_int, ok := in.(int64)
    if ok {
        value := float64(as_int)
        return value
    }
    as_string, ok := in.(string)
    if ok {
        parsed_float_string, err := strconv.ParseFloat(as_string, 64)
        if err == nil {
            value := parsed_float_string
            return value
        }
        parsed_int_string, err := strconv.ParseInt(as_string, 10, 32)
        if err == nil {
            value := float64(parsed_int_string)
            return value
        }
    }
    panic("ERR")
    result := float64(-1)
    log.Print("Error transforming ", in, " into Float64")
    return result
}

func getInt32(in interface{}) int32 {
    value := int32(getFloat64(in))
    if value == -1 {
        log.Print("Error transforming ", in, " into Int32")
    }
    return value
}

func getJsonString(in interface{}) string {
    as_string, ok := in.(string)
    if ok {
        return as_string
    }
    marshaled_result, err := json.Marshal(in)
    if err != nil {
        log.Print("Error transforming ", in, " into JSON string")
    }
    result := string(marshaled_result)
    return result
}

func exprElementIsTruthy(e exprElement) bool {
    switch e.typ {
        case J_STRING:
            return len(e.value.(string)) > 0
        case J_NUMBER:
            return e.value.(float64) > 0
        case J_OBJECT:
            return true
        case J_ARRAY:
            return true
        case J_BOOLEAN:
            return e.value.(bool)
        case J_NULL:
            return false
        default:
            return false
    }
}

func exprElementsMatch(lhs exprElement, rhs exprElement) bool {
    return lhs.typ == rhs.typ
}

func getHaystackFromNodeList(nodes []*jsonNode) map[*simplejson.Json]*jsonNode {
    hashmap := make(map[*simplejson.Json]*jsonNode, len(nodes))
    for _, node := range nodes {
        hashmap[node.json] = node
    }
    return hashmap
}
