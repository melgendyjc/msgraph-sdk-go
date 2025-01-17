package models
import (
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type SharedPCAllowedAccountType int

const (
    // Only guest accounts.
    GUEST_SHAREDPCALLOWEDACCOUNTTYPE SharedPCAllowedAccountType = iota
    // Only domain-joined accounts.
    DOMAIN_SHAREDPCALLOWEDACCOUNTTYPE
)

func (i SharedPCAllowedAccountType) String() string {
    return []string{"guest", "domain"}[i]
}
func ParseSharedPCAllowedAccountType(v string) (interface{}, error) {
    result := GUEST_SHAREDPCALLOWEDACCOUNTTYPE
    switch v {
        case "guest":
            result = GUEST_SHAREDPCALLOWEDACCOUNTTYPE
        case "domain":
            result = DOMAIN_SHAREDPCALLOWEDACCOUNTTYPE
        default:
            return 0, errors.New("Unknown SharedPCAllowedAccountType value: " + v)
    }
    return &result, nil
}
func SerializeSharedPCAllowedAccountType(values []SharedPCAllowedAccountType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
