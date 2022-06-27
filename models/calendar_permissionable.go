package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CalendarPermissionable 
type CalendarPermissionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedRoles()([]string)
    GetEmailAddress()(EmailAddressable)
    GetIsInsideOrganization()(*bool)
    GetIsRemovable()(*bool)
    GetRole()(*CalendarRoleType)
    SetAllowedRoles(value []string)()
    SetEmailAddress(value EmailAddressable)()
    SetIsInsideOrganization(value *bool)()
    SetIsRemovable(value *bool)()
    SetRole(value *CalendarRoleType)()
}
