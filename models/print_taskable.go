package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintTaskable 
type PrintTaskable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefinition()(PrintTaskDefinitionable)
    GetParentUrl()(*string)
    GetStatus()(PrintTaskStatusable)
    GetTrigger()(PrintTaskTriggerable)
    SetDefinition(value PrintTaskDefinitionable)()
    SetParentUrl(value *string)()
    SetStatus(value PrintTaskStatusable)()
    SetTrigger(value PrintTaskTriggerable)()
}
