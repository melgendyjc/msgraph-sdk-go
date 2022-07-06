package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsAppInstallationable 
type TeamsAppInstallationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTeamsApp()(TeamsAppable)
    GetTeamsAppDefinition()(TeamsAppDefinitionable)
    GetType()(*string)
    SetTeamsApp(value TeamsAppable)()
    SetTeamsAppDefinition(value TeamsAppDefinitionable)()
    SetType(value *string)()
}
