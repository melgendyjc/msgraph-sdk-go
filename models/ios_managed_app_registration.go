package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosManagedAppRegistration 
type IosManagedAppRegistration struct {
    ManagedAppRegistration
}
// NewIosManagedAppRegistration instantiates a new IosManagedAppRegistration and sets the default values.
func NewIosManagedAppRegistration()(*IosManagedAppRegistration) {
    m := &IosManagedAppRegistration{
        ManagedAppRegistration: *NewManagedAppRegistration(),
    }
    return m
}
// CreateIosManagedAppRegistrationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosManagedAppRegistrationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosManagedAppRegistration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosManagedAppRegistration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ManagedAppRegistration.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *IosManagedAppRegistration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ManagedAppRegistration.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
