package graph
import (
    "strings"
    "errors"
)
type ManagedDevicePartnerReportedHealthState int

const (
    UNKNOWN_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE ManagedDevicePartnerReportedHealthState = iota
    ACTIVATED_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
    DEACTIVATED_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
    SECURED_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
    LOWSEVERITY_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
    MEDIUMSEVERITY_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
    HIGHSEVERITY_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
    UNRESPONSIVE_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
    COMPROMISED_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
    MISCONFIGURED_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE
)

func (i ManagedDevicePartnerReportedHealthState) String() string {
    return []string{"UNKNOWN", "ACTIVATED", "DEACTIVATED", "SECURED", "LOWSEVERITY", "MEDIUMSEVERITY", "HIGHSEVERITY", "UNRESPONSIVE", "COMPROMISED", "MISCONFIGURED"}[i]
}
func ParseManagedDevicePartnerReportedHealthState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE, nil
        case "ACTIVATED":
            return ACTIVATED_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE, nil
        case "DEACTIVATED":
            return DEACTIVATED_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE, nil
        case "SECURED":
            return SECURED_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE, nil
        case "LOWSEVERITY":
            return LOWSEVERITY_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE, nil
        case "MEDIUMSEVERITY":
            return MEDIUMSEVERITY_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE, nil
        case "HIGHSEVERITY":
            return HIGHSEVERITY_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE, nil
        case "UNRESPONSIVE":
            return UNRESPONSIVE_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE, nil
        case "COMPROMISED":
            return COMPROMISED_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE, nil
        case "MISCONFIGURED":
            return MISCONFIGURED_MANAGEDDEVICEPARTNERREPORTEDHEALTHSTATE, nil
    }
    return 0, errors.New("Unknown ManagedDevicePartnerReportedHealthState value: " + v)
}
func SerializeManagedDevicePartnerReportedHealthState(values []ManagedDevicePartnerReportedHealthState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}