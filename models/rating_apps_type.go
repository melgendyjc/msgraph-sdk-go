package models
import (
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type RatingAppsType int

const (
    // Default value, allow all apps content
    ALLALLOWED_RATINGAPPSTYPE RatingAppsType = iota
    // Do not allow any apps content
    ALLBLOCKED_RATINGAPPSTYPE
    // 4+, age 4 and above
    AGESABOVE4_RATINGAPPSTYPE
    // 9+, age 9 and above
    AGESABOVE9_RATINGAPPSTYPE
    // 12+, age 12 and above 
    AGESABOVE12_RATINGAPPSTYPE
    // 17+, age 17 and above
    AGESABOVE17_RATINGAPPSTYPE
)

func (i RatingAppsType) String() string {
    return []string{"allAllowed", "allBlocked", "agesAbove4", "agesAbove9", "agesAbove12", "agesAbove17"}[i]
}
func ParseRatingAppsType(v string) (interface{}, error) {
    result := ALLALLOWED_RATINGAPPSTYPE
    switch v {
        case "allAllowed":
            result = ALLALLOWED_RATINGAPPSTYPE
        case "allBlocked":
            result = ALLBLOCKED_RATINGAPPSTYPE
        case "agesAbove4":
            result = AGESABOVE4_RATINGAPPSTYPE
        case "agesAbove9":
            result = AGESABOVE9_RATINGAPPSTYPE
        case "agesAbove12":
            result = AGESABOVE12_RATINGAPPSTYPE
        case "agesAbove17":
            result = AGESABOVE17_RATINGAPPSTYPE
        default:
            return 0, errors.New("Unknown RatingAppsType value: " + v)
    }
    return &result, nil
}
func SerializeRatingAppsType(values []RatingAppsType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
