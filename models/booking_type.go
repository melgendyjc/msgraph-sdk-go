package models
import (
    "errors"
)
// Provides operations to manage the collection of place entities.
type BookingType int

const (
    UNKNOWN_BOOKINGTYPE BookingType = iota
    STANDARD_BOOKINGTYPE
    RESERVED_BOOKINGTYPE
)

func (i BookingType) String() string {
    return []string{"unknown", "standard", "reserved"}[i]
}
func ParseBookingType(v string) (interface{}, error) {
    result := UNKNOWN_BOOKINGTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_BOOKINGTYPE
        case "standard":
            result = STANDARD_BOOKINGTYPE
        case "reserved":
            result = RESERVED_BOOKINGTYPE
        default:
            return 0, errors.New("Unknown BookingType value: " + v)
    }
    return &result, nil
}
func SerializeBookingType(values []BookingType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
