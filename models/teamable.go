package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Teamable 
type Teamable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllChannels()([]Channelable)
    GetChannels()([]Channelable)
    GetClassification()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetFunSettings()(TeamFunSettingsable)
    GetGroup()(Groupable)
    GetGuestSettings()(TeamGuestSettingsable)
    GetIncomingChannels()([]Channelable)
    GetInstalledApps()([]TeamsAppInstallationable)
    GetInternalId()(*string)
    GetIsArchived()(*bool)
    GetMembers()([]ConversationMemberable)
    GetMemberSettings()(TeamMemberSettingsable)
    GetMessagingSettings()(TeamMessagingSettingsable)
    GetOperations()([]TeamsAsyncOperationable)
    GetPrimaryChannel()(Channelable)
    GetSchedule()(Scheduleable)
    GetSpecialization()(*TeamSpecialization)
    GetTemplate()(TeamsTemplateable)
    GetTenantId()(*string)
    GetVisibility()(*TeamVisibilityType)
    GetWebUrl()(*string)
    SetAllChannels(value []Channelable)()
    SetChannels(value []Channelable)()
    SetClassification(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetFunSettings(value TeamFunSettingsable)()
    SetGroup(value Groupable)()
    SetGuestSettings(value TeamGuestSettingsable)()
    SetIncomingChannels(value []Channelable)()
    SetInstalledApps(value []TeamsAppInstallationable)()
    SetInternalId(value *string)()
    SetIsArchived(value *bool)()
    SetMembers(value []ConversationMemberable)()
    SetMemberSettings(value TeamMemberSettingsable)()
    SetMessagingSettings(value TeamMessagingSettingsable)()
    SetOperations(value []TeamsAsyncOperationable)()
    SetPrimaryChannel(value Channelable)()
    SetSchedule(value Scheduleable)()
    SetSpecialization(value *TeamSpecialization)()
    SetTemplate(value TeamsTemplateable)()
    SetTenantId(value *string)()
    SetVisibility(value *TeamVisibilityType)()
    SetWebUrl(value *string)()
}
