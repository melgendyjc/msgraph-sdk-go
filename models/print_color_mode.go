package models
import (
    "errors"
)
// Provides operations to manage the print singleton.
type PrintColorMode int

const (
    BLACKANDWHITE_PRINTCOLORMODE PrintColorMode = iota
    GRAYSCALE_PRINTCOLORMODE
    COLOR_PRINTCOLORMODE
    AUTO_PRINTCOLORMODE
    UNKNOWNFUTUREVALUE_PRINTCOLORMODE
)

func (i PrintColorMode) String() string {
    return []string{"blackAndWhite", "grayscale", "color", "auto", "unknownFutureValue"}[i]
}
func ParsePrintColorMode(v string) (interface{}, error) {
    result := BLACKANDWHITE_PRINTCOLORMODE
    switch v {
        case "blackAndWhite":
            result = BLACKANDWHITE_PRINTCOLORMODE
        case "grayscale":
            result = GRAYSCALE_PRINTCOLORMODE
        case "color":
            result = COLOR_PRINTCOLORMODE
        case "auto":
            result = AUTO_PRINTCOLORMODE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PRINTCOLORMODE
        default:
            return 0, errors.New("Unknown PrintColorMode value: " + v)
    }
    return &result, nil
}
func SerializePrintColorMode(values []PrintColorMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
