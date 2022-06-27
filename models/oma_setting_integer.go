package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OmaSettingInteger 
type OmaSettingInteger struct {
    OmaSetting
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Value.
    value *int32
}
// NewOmaSettingInteger instantiates a new OmaSettingInteger and sets the default values.
func NewOmaSettingInteger()(*OmaSettingInteger) {
    m := &OmaSettingInteger{
        OmaSetting: *NewOmaSetting(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOmaSettingIntegerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOmaSettingIntegerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOmaSettingInteger(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OmaSettingInteger) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OmaSettingInteger) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OmaSetting.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. Value.
func (m *OmaSettingInteger) GetValue()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// Serialize serializes information the current object
func (m *OmaSettingInteger) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OmaSetting.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("value", m.GetValue())
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
func (m *OmaSettingInteger) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetValue sets the value property value. Value.
func (m *OmaSettingInteger) SetValue(value *int32)() {
    if m != nil {
        m.value = value
    }
}
