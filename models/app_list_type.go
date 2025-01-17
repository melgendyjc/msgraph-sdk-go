package models
import (
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type AppListType int

const (
    // Default value, no intent.
    NONE_APPLISTTYPE AppListType = iota
    // The list represents the apps that will be considered compliant (only apps on the list are compliant).
    APPSINLISTCOMPLIANT_APPLISTTYPE
    // The list represents the apps that will be considered non compliant (all apps are compliant except apps on the list).
    APPSNOTINLISTCOMPLIANT_APPLISTTYPE
)

func (i AppListType) String() string {
    return []string{"none", "appsInListCompliant", "appsNotInListCompliant"}[i]
}
func ParseAppListType(v string) (interface{}, error) {
    result := NONE_APPLISTTYPE
    switch v {
        case "none":
            result = NONE_APPLISTTYPE
        case "appsInListCompliant":
            result = APPSINLISTCOMPLIANT_APPLISTTYPE
        case "appsNotInListCompliant":
            result = APPSNOTINLISTCOMPLIANT_APPLISTTYPE
        default:
            return 0, errors.New("Unknown AppListType value: " + v)
    }
    return &result, nil
}
func SerializeAppListType(values []AppListType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
