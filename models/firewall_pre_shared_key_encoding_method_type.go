package models
import (
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type FirewallPreSharedKeyEncodingMethodType int

const (
    // No value configured by Intune, do not override the user-configured device default value
    DEVICEDEFAULT_FIREWALLPRESHAREDKEYENCODINGMETHODTYPE FirewallPreSharedKeyEncodingMethodType = iota
    // Preshared key is not encoded. Instead, it is kept in its wide-character format
    NONE_FIREWALLPRESHAREDKEYENCODINGMETHODTYPE
    // Encode the preshared key using UTF-8
    UTF8_FIREWALLPRESHAREDKEYENCODINGMETHODTYPE
)

func (i FirewallPreSharedKeyEncodingMethodType) String() string {
    return []string{"deviceDefault", "none", "utF8"}[i]
}
func ParseFirewallPreSharedKeyEncodingMethodType(v string) (interface{}, error) {
    result := DEVICEDEFAULT_FIREWALLPRESHAREDKEYENCODINGMETHODTYPE
    switch v {
        case "deviceDefault":
            result = DEVICEDEFAULT_FIREWALLPRESHAREDKEYENCODINGMETHODTYPE
        case "none":
            result = NONE_FIREWALLPRESHAREDKEYENCODINGMETHODTYPE
        case "utF8":
            result = UTF8_FIREWALLPRESHAREDKEYENCODINGMETHODTYPE
        default:
            return 0, errors.New("Unknown FirewallPreSharedKeyEncodingMethodType value: " + v)
    }
    return &result, nil
}
func SerializeFirewallPreSharedKeyEncodingMethodType(values []FirewallPreSharedKeyEncodingMethodType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
