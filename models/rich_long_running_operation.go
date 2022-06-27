package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RichLongRunningOperation 
type RichLongRunningOperation struct {
    LongRunningOperation
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Error due to which the operation failed.
    error PublicErrorable
    // A value between 0 and 100 that indicates the progress of the operation.
    percentageComplete *int32
    // A unique identifier for the result.
    resourceId *string
    // Type of the operation.
    type_escaped *string
}
// NewRichLongRunningOperation instantiates a new RichLongRunningOperation and sets the default values.
func NewRichLongRunningOperation()(*RichLongRunningOperation) {
    m := &RichLongRunningOperation{
        LongRunningOperation: *NewLongRunningOperation(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRichLongRunningOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRichLongRunningOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRichLongRunningOperation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RichLongRunningOperation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetError gets the error property value. Error due to which the operation failed.
func (m *RichLongRunningOperation) GetError()(PublicErrorable) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RichLongRunningOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LongRunningOperation.GetFieldDeserializers()
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePublicErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(PublicErrorable))
        }
        return nil
    }
    res["percentageComplete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPercentageComplete(val)
        }
        return nil
    }
    res["resourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetPercentageComplete gets the percentageComplete property value. A value between 0 and 100 that indicates the progress of the operation.
func (m *RichLongRunningOperation) GetPercentageComplete()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.percentageComplete
    }
}
// GetResourceId gets the resourceId property value. A unique identifier for the result.
func (m *RichLongRunningOperation) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
// GetType gets the type property value. Type of the operation.
func (m *RichLongRunningOperation) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *RichLongRunningOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LongRunningOperation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("percentageComplete", m.GetPercentageComplete())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type", m.GetType())
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
func (m *RichLongRunningOperation) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetError sets the error property value. Error due to which the operation failed.
func (m *RichLongRunningOperation) SetError(value PublicErrorable)() {
    if m != nil {
        m.error = value
    }
}
// SetPercentageComplete sets the percentageComplete property value. A value between 0 and 100 that indicates the progress of the operation.
func (m *RichLongRunningOperation) SetPercentageComplete(value *int32)() {
    if m != nil {
        m.percentageComplete = value
    }
}
// SetResourceId sets the resourceId property value. A unique identifier for the result.
func (m *RichLongRunningOperation) SetResourceId(value *string)() {
    if m != nil {
        m.resourceId = value
    }
}
// SetType sets the type property value. Type of the operation.
func (m *RichLongRunningOperation) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
