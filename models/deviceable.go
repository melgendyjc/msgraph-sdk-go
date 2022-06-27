package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deviceable 
type Deviceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountEnabled()(*bool)
    GetAlternativeSecurityIds()([]AlternativeSecurityIdable)
    GetApproximateLastSignInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetComplianceExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeviceId()(*string)
    GetDeviceMetadata()(*string)
    GetDeviceVersion()(*int32)
    GetDisplayName()(*string)
    GetExtensions()([]Extensionable)
    GetIsCompliant()(*bool)
    GetIsManaged()(*bool)
    GetMdmAppId()(*string)
    GetMemberOf()([]DirectoryObjectable)
    GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOnPremisesSyncEnabled()(*bool)
    GetOperatingSystem()(*string)
    GetOperatingSystemVersion()(*string)
    GetPhysicalIds()([]string)
    GetProfileType()(*string)
    GetRegisteredOwners()([]DirectoryObjectable)
    GetRegisteredUsers()([]DirectoryObjectable)
    GetSystemLabels()([]string)
    GetTransitiveMemberOf()([]DirectoryObjectable)
    GetTrustType()(*string)
    SetAccountEnabled(value *bool)()
    SetAlternativeSecurityIds(value []AlternativeSecurityIdable)()
    SetApproximateLastSignInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetComplianceExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeviceId(value *string)()
    SetDeviceMetadata(value *string)()
    SetDeviceVersion(value *int32)()
    SetDisplayName(value *string)()
    SetExtensions(value []Extensionable)()
    SetIsCompliant(value *bool)()
    SetIsManaged(value *bool)()
    SetMdmAppId(value *string)()
    SetMemberOf(value []DirectoryObjectable)()
    SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOnPremisesSyncEnabled(value *bool)()
    SetOperatingSystem(value *string)()
    SetOperatingSystemVersion(value *string)()
    SetPhysicalIds(value []string)()
    SetProfileType(value *string)()
    SetRegisteredOwners(value []DirectoryObjectable)()
    SetRegisteredUsers(value []DirectoryObjectable)()
    SetSystemLabels(value []string)()
    SetTransitiveMemberOf(value []DirectoryObjectable)()
    SetTrustType(value *string)()
}
