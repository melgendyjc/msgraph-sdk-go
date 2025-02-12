package models
import (
    "errors"
)
// Provides operations to manage the deviceAppManagement singleton.
type ManagedAppAvailability int

const (
    // A globally available app to all tenants.
    GLOBAL_MANAGEDAPPAVAILABILITY ManagedAppAvailability = iota
    // A line of business apps private to an organization.
    LINEOFBUSINESS_MANAGEDAPPAVAILABILITY
)

func (i ManagedAppAvailability) String() string {
    return []string{"global", "lineOfBusiness"}[i]
}
func ParseManagedAppAvailability(v string) (interface{}, error) {
    result := GLOBAL_MANAGEDAPPAVAILABILITY
    switch v {
        case "global":
            result = GLOBAL_MANAGEDAPPAVAILABILITY
        case "lineOfBusiness":
            result = LINEOFBUSINESS_MANAGEDAPPAVAILABILITY
        default:
            return 0, errors.New("Unknown ManagedAppAvailability value: " + v)
    }
    return &result, nil
}
func SerializeManagedAppAvailability(values []ManagedAppAvailability) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
