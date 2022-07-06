package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ShiftItemable 
type ShiftItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ScheduleEntityable
    GetActivities()([]ShiftActivityable)
    GetDisplayName()(*string)
    GetNotes()(*string)
    GetType()(*string)
    SetActivities(value []ShiftActivityable)()
    SetDisplayName(value *string)()
    SetNotes(value *string)()
    SetType(value *string)()
}
