package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Alertable 
type Alertable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivityGroupName()(*string)
    GetAlertDetections()([]AlertDetectionable)
    GetAssignedTo()(*string)
    GetAzureSubscriptionId()(*string)
    GetAzureTenantId()(*string)
    GetCategory()(*string)
    GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCloudAppStates()([]CloudAppSecurityStateable)
    GetComments()([]string)
    GetConfidence()(*int32)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDetectionIds()([]string)
    GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFeedback()(*AlertFeedback)
    GetFileStates()([]FileSecurityStateable)
    GetHistoryStates()([]AlertHistoryStateable)
    GetHostStates()([]HostSecurityStateable)
    GetIncidentIds()([]string)
    GetInvestigationSecurityStates()([]InvestigationSecurityStateable)
    GetLastEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMalwareStates()([]MalwareStateable)
    GetMessageSecurityStates()([]MessageSecurityStateable)
    GetNetworkConnections()([]NetworkConnectionable)
    GetProcesses()([]Processable)
    GetRecommendedActions()([]string)
    GetRegistryKeyStates()([]RegistryKeyStateable)
    GetSecurityResources()([]SecurityResourceable)
    GetSeverity()(*AlertSeverity)
    GetSourceMaterials()([]string)
    GetStatus()(*AlertStatus)
    GetTags()([]string)
    GetTitle()(*string)
    GetTriggers()([]AlertTriggerable)
    GetUriClickSecurityStates()([]UriClickSecurityStateable)
    GetUserStates()([]UserSecurityStateable)
    GetVendorInformation()(SecurityVendorInformationable)
    GetVulnerabilityStates()([]VulnerabilityStateable)
    SetActivityGroupName(value *string)()
    SetAlertDetections(value []AlertDetectionable)()
    SetAssignedTo(value *string)()
    SetAzureSubscriptionId(value *string)()
    SetAzureTenantId(value *string)()
    SetCategory(value *string)()
    SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCloudAppStates(value []CloudAppSecurityStateable)()
    SetComments(value []string)()
    SetConfidence(value *int32)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDetectionIds(value []string)()
    SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFeedback(value *AlertFeedback)()
    SetFileStates(value []FileSecurityStateable)()
    SetHistoryStates(value []AlertHistoryStateable)()
    SetHostStates(value []HostSecurityStateable)()
    SetIncidentIds(value []string)()
    SetInvestigationSecurityStates(value []InvestigationSecurityStateable)()
    SetLastEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMalwareStates(value []MalwareStateable)()
    SetMessageSecurityStates(value []MessageSecurityStateable)()
    SetNetworkConnections(value []NetworkConnectionable)()
    SetProcesses(value []Processable)()
    SetRecommendedActions(value []string)()
    SetRegistryKeyStates(value []RegistryKeyStateable)()
    SetSecurityResources(value []SecurityResourceable)()
    SetSeverity(value *AlertSeverity)()
    SetSourceMaterials(value []string)()
    SetStatus(value *AlertStatus)()
    SetTags(value []string)()
    SetTitle(value *string)()
    SetTriggers(value []AlertTriggerable)()
    SetUriClickSecurityStates(value []UriClickSecurityStateable)()
    SetUserStates(value []UserSecurityStateable)()
    SetVendorInformation(value SecurityVendorInformationable)()
    SetVulnerabilityStates(value []VulnerabilityStateable)()
}
