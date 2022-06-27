package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDeviceable 
type ManagedDeviceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivationLockBypassCode()(*string)
    GetAndroidSecurityPatchLevel()(*string)
    GetAzureADDeviceId()(*string)
    GetAzureADRegistered()(*bool)
    GetComplianceGracePeriodExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetComplianceState()(*ComplianceState)
    GetConfigurationManagerClientEnabledFeatures()(ConfigurationManagerClientEnabledFeaturesable)
    GetDeviceActionResults()([]DeviceActionResultable)
    GetDeviceCategory()(DeviceCategoryable)
    GetDeviceCategoryDisplayName()(*string)
    GetDeviceCompliancePolicyStates()([]DeviceCompliancePolicyStateable)
    GetDeviceConfigurationStates()([]DeviceConfigurationStateable)
    GetDeviceEnrollmentType()(*DeviceEnrollmentType)
    GetDeviceHealthAttestationState()(DeviceHealthAttestationStateable)
    GetDeviceName()(*string)
    GetDeviceRegistrationState()(*DeviceRegistrationState)
    GetEasActivated()(*bool)
    GetEasActivationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEasDeviceId()(*string)
    GetEmailAddress()(*string)
    GetEnrolledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEthernetMacAddress()(*string)
    GetExchangeAccessState()(*DeviceManagementExchangeAccessState)
    GetExchangeAccessStateReason()(*DeviceManagementExchangeAccessStateReason)
    GetExchangeLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFreeStorageSpaceInBytes()(*int64)
    GetIccid()(*string)
    GetImei()(*string)
    GetIsEncrypted()(*bool)
    GetIsSupervised()(*bool)
    GetJailBroken()(*string)
    GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagedDeviceName()(*string)
    GetManagedDeviceOwnerType()(*ManagedDeviceOwnerType)
    GetManagementAgent()(*ManagementAgentType)
    GetManufacturer()(*string)
    GetMeid()(*string)
    GetModel()(*string)
    GetNotes()(*string)
    GetOperatingSystem()(*string)
    GetOsVersion()(*string)
    GetPartnerReportedThreatState()(*ManagedDevicePartnerReportedHealthState)
    GetPhoneNumber()(*string)
    GetPhysicalMemoryInBytes()(*int64)
    GetRemoteAssistanceSessionErrorDetails()(*string)
    GetRemoteAssistanceSessionUrl()(*string)
    GetSerialNumber()(*string)
    GetSubscriberCarrier()(*string)
    GetTotalStorageSpaceInBytes()(*int64)
    GetUdid()(*string)
    GetUserDisplayName()(*string)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    GetWiFiMacAddress()(*string)
    SetActivationLockBypassCode(value *string)()
    SetAndroidSecurityPatchLevel(value *string)()
    SetAzureADDeviceId(value *string)()
    SetAzureADRegistered(value *bool)()
    SetComplianceGracePeriodExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetComplianceState(value *ComplianceState)()
    SetConfigurationManagerClientEnabledFeatures(value ConfigurationManagerClientEnabledFeaturesable)()
    SetDeviceActionResults(value []DeviceActionResultable)()
    SetDeviceCategory(value DeviceCategoryable)()
    SetDeviceCategoryDisplayName(value *string)()
    SetDeviceCompliancePolicyStates(value []DeviceCompliancePolicyStateable)()
    SetDeviceConfigurationStates(value []DeviceConfigurationStateable)()
    SetDeviceEnrollmentType(value *DeviceEnrollmentType)()
    SetDeviceHealthAttestationState(value DeviceHealthAttestationStateable)()
    SetDeviceName(value *string)()
    SetDeviceRegistrationState(value *DeviceRegistrationState)()
    SetEasActivated(value *bool)()
    SetEasActivationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEasDeviceId(value *string)()
    SetEmailAddress(value *string)()
    SetEnrolledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEthernetMacAddress(value *string)()
    SetExchangeAccessState(value *DeviceManagementExchangeAccessState)()
    SetExchangeAccessStateReason(value *DeviceManagementExchangeAccessStateReason)()
    SetExchangeLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFreeStorageSpaceInBytes(value *int64)()
    SetIccid(value *string)()
    SetImei(value *string)()
    SetIsEncrypted(value *bool)()
    SetIsSupervised(value *bool)()
    SetJailBroken(value *string)()
    SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagedDeviceName(value *string)()
    SetManagedDeviceOwnerType(value *ManagedDeviceOwnerType)()
    SetManagementAgent(value *ManagementAgentType)()
    SetManufacturer(value *string)()
    SetMeid(value *string)()
    SetModel(value *string)()
    SetNotes(value *string)()
    SetOperatingSystem(value *string)()
    SetOsVersion(value *string)()
    SetPartnerReportedThreatState(value *ManagedDevicePartnerReportedHealthState)()
    SetPhoneNumber(value *string)()
    SetPhysicalMemoryInBytes(value *int64)()
    SetRemoteAssistanceSessionErrorDetails(value *string)()
    SetRemoteAssistanceSessionUrl(value *string)()
    SetSerialNumber(value *string)()
    SetSubscriberCarrier(value *string)()
    SetTotalStorageSpaceInBytes(value *int64)()
    SetUdid(value *string)()
    SetUserDisplayName(value *string)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
    SetWiFiMacAddress(value *string)()
}
