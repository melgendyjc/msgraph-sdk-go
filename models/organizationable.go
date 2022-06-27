package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Organizationable 
type Organizationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedPlans()([]AssignedPlanable)
    GetBranding()(OrganizationalBrandingable)
    GetBusinessPhones()([]string)
    GetCertificateBasedAuthConfiguration()([]CertificateBasedAuthConfigurationable)
    GetCity()(*string)
    GetCountry()(*string)
    GetCountryLetterCode()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDisplayName()(*string)
    GetExtensions()([]Extensionable)
    GetMarketingNotificationEmails()([]string)
    GetMobileDeviceManagementAuthority()(*MdmAuthority)
    GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOnPremisesSyncEnabled()(*bool)
    GetPostalCode()(*string)
    GetPreferredLanguage()(*string)
    GetPrivacyProfile()(PrivacyProfileable)
    GetProvisionedPlans()([]ProvisionedPlanable)
    GetSecurityComplianceNotificationMails()([]string)
    GetSecurityComplianceNotificationPhones()([]string)
    GetState()(*string)
    GetStreet()(*string)
    GetTechnicalNotificationMails()([]string)
    GetTenantType()(*string)
    GetVerifiedDomains()([]VerifiedDomainable)
    SetAssignedPlans(value []AssignedPlanable)()
    SetBranding(value OrganizationalBrandingable)()
    SetBusinessPhones(value []string)()
    SetCertificateBasedAuthConfiguration(value []CertificateBasedAuthConfigurationable)()
    SetCity(value *string)()
    SetCountry(value *string)()
    SetCountryLetterCode(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDisplayName(value *string)()
    SetExtensions(value []Extensionable)()
    SetMarketingNotificationEmails(value []string)()
    SetMobileDeviceManagementAuthority(value *MdmAuthority)()
    SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOnPremisesSyncEnabled(value *bool)()
    SetPostalCode(value *string)()
    SetPreferredLanguage(value *string)()
    SetPrivacyProfile(value PrivacyProfileable)()
    SetProvisionedPlans(value []ProvisionedPlanable)()
    SetSecurityComplianceNotificationMails(value []string)()
    SetSecurityComplianceNotificationPhones(value []string)()
    SetState(value *string)()
    SetStreet(value *string)()
    SetTechnicalNotificationMails(value []string)()
    SetTenantType(value *string)()
    SetVerifiedDomains(value []VerifiedDomainable)()
}
