package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConvertIdResultable 
type ConvertIdResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetErrorDetails()(GenericErrorable)
    GetSourceId()(*string)
    GetTargetId()(*string)
    SetErrorDetails(value GenericErrorable)()
    SetSourceId(value *string)()
    SetTargetId(value *string)()
}
