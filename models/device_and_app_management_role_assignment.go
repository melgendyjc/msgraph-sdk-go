package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAndAppManagementRoleAssignment 
type DeviceAndAppManagementRoleAssignment struct {
    RoleAssignment
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The list of ids of role member security groups. These are IDs from Azure Active Directory.
    members []string
}
// NewDeviceAndAppManagementRoleAssignment instantiates a new DeviceAndAppManagementRoleAssignment and sets the default values.
func NewDeviceAndAppManagementRoleAssignment()(*DeviceAndAppManagementRoleAssignment) {
    m := &DeviceAndAppManagementRoleAssignment{
        RoleAssignment: *NewRoleAssignment(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceAndAppManagementRoleAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAndAppManagementRoleAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAndAppManagementRoleAssignment(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceAndAppManagementRoleAssignment) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAndAppManagementRoleAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RoleAssignment.GetFieldDeserializers()
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMembers(res)
        }
        return nil
    }
    return res
}
// GetMembers gets the members property value. The list of ids of role member security groups. These are IDs from Azure Active Directory.
func (m *DeviceAndAppManagementRoleAssignment) GetMembers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// Serialize serializes information the current object
func (m *DeviceAndAppManagementRoleAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RoleAssignment.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMembers() != nil {
        err = writer.WriteCollectionOfStringValues("members", m.GetMembers())
        if err != nil {
            return err
        }
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
func (m *DeviceAndAppManagementRoleAssignment) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMembers sets the members property value. The list of ids of role member security groups. These are IDs from Azure Active Directory.
func (m *DeviceAndAppManagementRoleAssignment) SetMembers(value []string)() {
    if m != nil {
        m.members = value
    }
}
