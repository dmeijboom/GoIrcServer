package assert;

import (
    "strconv"
)

const (
    COLOR_BLACK = "\x1b[30m"
    COLOR_RED = "\x1b[31m"
    COLOR_GREEN = "\x1b[32m"
    BOLD_ON = "\x1b[1m"
    BOLD_OFF = "\x1b[22m"
    UNDERLINE_ON = "\x1b[4m"
    UNDERLINE_OFF = "\x1b[24m"
    BLINK_ON = "\x1b[5m"
    BLINK_OFF = "\x1b[25m"
)

// EqualsString checks if two strings match
func EqualsString(expected string, actual string) string {
    if expected == actual {
        return ""
    }
    
    left := ""
    diff := ""
    
    diff += "expected "
    diff += BOLD_ON
    diff += expected
    diff += BOLD_OFF
    diff += " actual "
    
    if len(actual) == 0 {
        diff += UNDERLINE_ON
        diff += "empty string"
        diff += UNDERLINE_OFF
        return diff
    }
    
    for i := range expected {
        if len(actual) <= i || expected[i] != actual[i] {
            break
        }
        
        left += string(actual[i])
    }
    
    if len(left) > 0 {
        diff += COLOR_GREEN
        diff += left
        diff += COLOR_RED
        
        if len(expected) > len(actual) {
            diff += expected[len(left):]
        } else {
            diff += actual[len(left):]
        }
        
        diff += COLOR_BLACK
    } else {
        diff += COLOR_RED
        diff += actual
        diff += COLOR_BLACK
    }
    
    return diff
}

func EqualsInt(expected int, actual int) string {
    return EqualsString(strconv.Itoa(expected), strconv.Itoa(actual))
}

func EqualsBool(expected bool, actual bool) string {
    if expected == actual {
        return ""
    }
    
    msg := "Expected assertion to be "
    
    if expected {
        msg += "true"
    } else {
        msg += "false"
    }
    
    msg += " instead it was "
    
    if actual {
        msg += "true"
    } else {
        msg += "false"
    }
    
    return msg
}