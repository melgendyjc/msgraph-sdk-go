package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceEnrollmentPlatformRestrictionable 
type DeviceEnrollmentPlatformRestrictionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOsMaximumVersion()(*string)
    GetOsMinimumVersion()(*string)
    GetPersonalDeviceEnrollmentBlocked()(*bool)
    GetPlatformBlocked()(*bool)
    SetOsMaximumVersion(value *string)()
    SetOsMinimumVersion(value *string)()
    SetPersonalDeviceEnrollmentBlocked(value *bool)()
    SetPlatformBlocked(value *bool)()
}
