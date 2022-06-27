package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AadUserConversationMemberResult 
type AadUserConversationMemberResult struct {
    ActionResultPart
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The user object ID of the Azure AD user that was being added as part of the bulk operation.
    userId *string
}
// NewAadUserConversationMemberResult instantiates a new AadUserConversationMemberResult and sets the default values.
func NewAadUserConversationMemberResult()(*AadUserConversationMemberResult) {
    m := &AadUserConversationMemberResult{
        ActionResultPart: *NewActionResultPart(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAadUserConversationMemberResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAadUserConversationMemberResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAadUserConversationMemberResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AadUserConversationMemberResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AadUserConversationMemberResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ActionResultPart.GetFieldDeserializers()
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetUserId gets the userId property value. The user object ID of the Azure AD user that was being added as part of the bulk operation.
func (m *AadUserConversationMemberResult) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// Serialize serializes information the current object
func (m *AadUserConversationMemberResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ActionResultPart.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
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
func (m *AadUserConversationMemberResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetUserId sets the userId property value. The user object ID of the Azure AD user that was being added as part of the bulk operation.
func (m *AadUserConversationMemberResult) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
