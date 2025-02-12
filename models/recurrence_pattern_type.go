package models
import (
    "errors"
)
// Provides operations to manage the collection of application entities.
type RecurrencePatternType int

const (
    DAILY_RECURRENCEPATTERNTYPE RecurrencePatternType = iota
    WEEKLY_RECURRENCEPATTERNTYPE
    ABSOLUTEMONTHLY_RECURRENCEPATTERNTYPE
    RELATIVEMONTHLY_RECURRENCEPATTERNTYPE
    ABSOLUTEYEARLY_RECURRENCEPATTERNTYPE
    RELATIVEYEARLY_RECURRENCEPATTERNTYPE
)

func (i RecurrencePatternType) String() string {
    return []string{"daily", "weekly", "absoluteMonthly", "relativeMonthly", "absoluteYearly", "relativeYearly"}[i]
}
func ParseRecurrencePatternType(v string) (interface{}, error) {
    result := DAILY_RECURRENCEPATTERNTYPE
    switch v {
        case "daily":
            result = DAILY_RECURRENCEPATTERNTYPE
        case "weekly":
            result = WEEKLY_RECURRENCEPATTERNTYPE
        case "absoluteMonthly":
            result = ABSOLUTEMONTHLY_RECURRENCEPATTERNTYPE
        case "relativeMonthly":
            result = RELATIVEMONTHLY_RECURRENCEPATTERNTYPE
        case "absoluteYearly":
            result = ABSOLUTEYEARLY_RECURRENCEPATTERNTYPE
        case "relativeYearly":
            result = RELATIVEYEARLY_RECURRENCEPATTERNTYPE
        default:
            return 0, errors.New("Unknown RecurrencePatternType value: " + v)
    }
    return &result, nil
}
func SerializeRecurrencePatternType(values []RecurrencePatternType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
