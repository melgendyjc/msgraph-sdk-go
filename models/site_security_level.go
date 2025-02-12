package models
import (
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type SiteSecurityLevel int

const (
    // User Defined, default value, no intent.
    USERDEFINED_SITESECURITYLEVEL SiteSecurityLevel = iota
    // Low.
    LOW_SITESECURITYLEVEL
    // Medium-low.
    MEDIUMLOW_SITESECURITYLEVEL
    // Medium.
    MEDIUM_SITESECURITYLEVEL
    // Medium-high.
    MEDIUMHIGH_SITESECURITYLEVEL
    // High.
    HIGH_SITESECURITYLEVEL
)

func (i SiteSecurityLevel) String() string {
    return []string{"userDefined", "low", "mediumLow", "medium", "mediumHigh", "high"}[i]
}
func ParseSiteSecurityLevel(v string) (interface{}, error) {
    result := USERDEFINED_SITESECURITYLEVEL
    switch v {
        case "userDefined":
            result = USERDEFINED_SITESECURITYLEVEL
        case "low":
            result = LOW_SITESECURITYLEVEL
        case "mediumLow":
            result = MEDIUMLOW_SITESECURITYLEVEL
        case "medium":
            result = MEDIUM_SITESECURITYLEVEL
        case "mediumHigh":
            result = MEDIUMHIGH_SITESECURITYLEVEL
        case "high":
            result = HIGH_SITESECURITYLEVEL
        default:
            return 0, errors.New("Unknown SiteSecurityLevel value: " + v)
    }
    return &result, nil
}
func SerializeSiteSecurityLevel(values []SiteSecurityLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
