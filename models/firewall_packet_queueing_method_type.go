package models
import (
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type FirewallPacketQueueingMethodType int

const (
    // No value configured by Intune, do not override the user-configured device default value
    DEVICEDEFAULT_FIREWALLPACKETQUEUEINGMETHODTYPE FirewallPacketQueueingMethodType = iota
    // Disable packet queuing
    DISABLED_FIREWALLPACKETQUEUEINGMETHODTYPE
    // Queue inbound encrypted packets
    QUEUEINBOUND_FIREWALLPACKETQUEUEINGMETHODTYPE
    // Queue decrypted outbound packets for forwarding
    QUEUEOUTBOUND_FIREWALLPACKETQUEUEINGMETHODTYPE
    // Queue both inbound and outbound packets
    QUEUEBOTH_FIREWALLPACKETQUEUEINGMETHODTYPE
)

func (i FirewallPacketQueueingMethodType) String() string {
    return []string{"deviceDefault", "disabled", "queueInbound", "queueOutbound", "queueBoth"}[i]
}
func ParseFirewallPacketQueueingMethodType(v string) (interface{}, error) {
    result := DEVICEDEFAULT_FIREWALLPACKETQUEUEINGMETHODTYPE
    switch v {
        case "deviceDefault":
            result = DEVICEDEFAULT_FIREWALLPACKETQUEUEINGMETHODTYPE
        case "disabled":
            result = DISABLED_FIREWALLPACKETQUEUEINGMETHODTYPE
        case "queueInbound":
            result = QUEUEINBOUND_FIREWALLPACKETQUEUEINGMETHODTYPE
        case "queueOutbound":
            result = QUEUEOUTBOUND_FIREWALLPACKETQUEUEINGMETHODTYPE
        case "queueBoth":
            result = QUEUEBOTH_FIREWALLPACKETQUEUEINGMETHODTYPE
        default:
            return 0, errors.New("Unknown FirewallPacketQueueingMethodType value: " + v)
    }
    return &result, nil
}
func SerializeFirewallPacketQueueingMethodType(values []FirewallPacketQueueingMethodType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
