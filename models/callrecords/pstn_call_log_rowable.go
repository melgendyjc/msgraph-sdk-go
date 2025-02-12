package callrecords

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PstnCallLogRowable 
type PstnCallLogRowable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCallDurationSource()(*PstnCallDurationSource)
    GetCalleeNumber()(*string)
    GetCallerNumber()(*string)
    GetCallId()(*string)
    GetCallType()(*string)
    GetCharge()(*float64)
    GetConferenceId()(*string)
    GetConnectionCharge()(*float64)
    GetCurrency()(*string)
    GetDestinationContext()(*string)
    GetDestinationName()(*string)
    GetDuration()(*int32)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetId()(*string)
    GetInventoryType()(*string)
    GetLicenseCapability()(*string)
    GetOperator()(*string)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTenantCountryCode()(*string)
    GetUsageCountryCode()(*string)
    GetUserDisplayName()(*string)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    SetCallDurationSource(value *PstnCallDurationSource)()
    SetCalleeNumber(value *string)()
    SetCallerNumber(value *string)()
    SetCallId(value *string)()
    SetCallType(value *string)()
    SetCharge(value *float64)()
    SetConferenceId(value *string)()
    SetConnectionCharge(value *float64)()
    SetCurrency(value *string)()
    SetDestinationContext(value *string)()
    SetDestinationName(value *string)()
    SetDuration(value *int32)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetId(value *string)()
    SetInventoryType(value *string)()
    SetLicenseCapability(value *string)()
    SetOperator(value *string)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTenantCountryCode(value *string)()
    SetUsageCountryCode(value *string)()
    SetUserDisplayName(value *string)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
}
