package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AllDevicesAssignmentTarget 
type AllDevicesAssignmentTarget struct {
    DeviceAndAppManagementAssignmentTarget
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
}
// NewAllDevicesAssignmentTarget instantiates a new AllDevicesAssignmentTarget and sets the default values.
func NewAllDevicesAssignmentTarget()(*AllDevicesAssignmentTarget) {
    m := &AllDevicesAssignmentTarget{
        DeviceAndAppManagementAssignmentTarget: *NewDeviceAndAppManagementAssignmentTarget(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAllDevicesAssignmentTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAllDevicesAssignmentTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAllDevicesAssignmentTarget(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AllDevicesAssignmentTarget) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AllDevicesAssignmentTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceAndAppManagementAssignmentTarget.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AllDevicesAssignmentTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceAndAppManagementAssignmentTarget.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AllDevicesAssignmentTarget) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
