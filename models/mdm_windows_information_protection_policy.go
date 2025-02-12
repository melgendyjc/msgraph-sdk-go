package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MdmWindowsInformationProtectionPolicy 
type MdmWindowsInformationProtectionPolicy struct {
    WindowsInformationProtection
}
// NewMdmWindowsInformationProtectionPolicy instantiates a new MdmWindowsInformationProtectionPolicy and sets the default values.
func NewMdmWindowsInformationProtectionPolicy()(*MdmWindowsInformationProtectionPolicy) {
    m := &MdmWindowsInformationProtectionPolicy{
        WindowsInformationProtection: *NewWindowsInformationProtection(),
    }
    return m
}
// CreateMdmWindowsInformationProtectionPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMdmWindowsInformationProtectionPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMdmWindowsInformationProtectionPolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MdmWindowsInformationProtectionPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsInformationProtection.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *MdmWindowsInformationProtectionPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsInformationProtection.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
