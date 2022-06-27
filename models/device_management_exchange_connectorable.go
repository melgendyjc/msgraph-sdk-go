package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementExchangeConnectorable 
type DeviceManagementExchangeConnectorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnectorServerName()(*string)
    GetExchangeAlias()(*string)
    GetExchangeConnectorType()(*DeviceManagementExchangeConnectorType)
    GetExchangeOrganization()(*string)
    GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPrimarySmtpAddress()(*string)
    GetServerName()(*string)
    GetStatus()(*DeviceManagementExchangeConnectorStatus)
    GetVersion()(*string)
    SetConnectorServerName(value *string)()
    SetExchangeAlias(value *string)()
    SetExchangeConnectorType(value *DeviceManagementExchangeConnectorType)()
    SetExchangeOrganization(value *string)()
    SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPrimarySmtpAddress(value *string)()
    SetServerName(value *string)()
    SetStatus(value *DeviceManagementExchangeConnectorStatus)()
    SetVersion(value *string)()
}
