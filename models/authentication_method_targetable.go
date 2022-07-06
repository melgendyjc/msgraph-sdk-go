package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationMethodTargetable 
type AuthenticationMethodTargetable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsRegistrationRequired()(*bool)
    GetTargetType()(*AuthenticationMethodTargetType)
    GetType()(*string)
    SetIsRegistrationRequired(value *bool)()
    SetTargetType(value *AuthenticationMethodTargetType)()
    SetType(value *string)()
}