package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10NetworkProxyServerable 
type Windows10NetworkProxyServerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(*string)
    GetExceptions()([]string)
    GetUseForLocalAddresses()(*bool)
    SetAddress(value *string)()
    SetExceptions(value []string)()
    SetUseForLocalAddresses(value *bool)()
}
