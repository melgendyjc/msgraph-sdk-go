package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatMessagePolicyViolationPolicyTip 
type ChatMessagePolicyViolationPolicyTip struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The URL a user can visit to read about the data loss prevention policies for the organization. (ie, policies about what users shouldn't say in chats)
    complianceUrl *string
    // Explanatory text shown to the sender of the message.
    generalText *string
    // The list of improper data in the message that was detected by the data loss prevention app. Each DLP app defines its own conditions, examples include 'Credit Card Number' and 'Social Security Number'.
    matchedConditionDescriptions []string
}
// NewChatMessagePolicyViolationPolicyTip instantiates a new chatMessagePolicyViolationPolicyTip and sets the default values.
func NewChatMessagePolicyViolationPolicyTip()(*ChatMessagePolicyViolationPolicyTip) {
    m := &ChatMessagePolicyViolationPolicyTip{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateChatMessagePolicyViolationPolicyTipFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatMessagePolicyViolationPolicyTipFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatMessagePolicyViolationPolicyTip(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChatMessagePolicyViolationPolicyTip) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetComplianceUrl gets the complianceUrl property value. The URL a user can visit to read about the data loss prevention policies for the organization. (ie, policies about what users shouldn't say in chats)
func (m *ChatMessagePolicyViolationPolicyTip) GetComplianceUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.complianceUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatMessagePolicyViolationPolicyTip) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["complianceUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceUrl(val)
        }
        return nil
    }
    res["generalText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeneralText(val)
        }
        return nil
    }
    res["matchedConditionDescriptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMatchedConditionDescriptions(res)
        }
        return nil
    }
    return res
}
// GetGeneralText gets the generalText property value. Explanatory text shown to the sender of the message.
func (m *ChatMessagePolicyViolationPolicyTip) GetGeneralText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.generalText
    }
}
// GetMatchedConditionDescriptions gets the matchedConditionDescriptions property value. The list of improper data in the message that was detected by the data loss prevention app. Each DLP app defines its own conditions, examples include 'Credit Card Number' and 'Social Security Number'.
func (m *ChatMessagePolicyViolationPolicyTip) GetMatchedConditionDescriptions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.matchedConditionDescriptions
    }
}
// Serialize serializes information the current object
func (m *ChatMessagePolicyViolationPolicyTip) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("complianceUrl", m.GetComplianceUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("generalText", m.GetGeneralText())
        if err != nil {
            return err
        }
    }
    if m.GetMatchedConditionDescriptions() != nil {
        err := writer.WriteCollectionOfStringValues("matchedConditionDescriptions", m.GetMatchedConditionDescriptions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChatMessagePolicyViolationPolicyTip) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetComplianceUrl sets the complianceUrl property value. The URL a user can visit to read about the data loss prevention policies for the organization. (ie, policies about what users shouldn't say in chats)
func (m *ChatMessagePolicyViolationPolicyTip) SetComplianceUrl(value *string)() {
    if m != nil {
        m.complianceUrl = value
    }
}
// SetGeneralText sets the generalText property value. Explanatory text shown to the sender of the message.
func (m *ChatMessagePolicyViolationPolicyTip) SetGeneralText(value *string)() {
    if m != nil {
        m.generalText = value
    }
}
// SetMatchedConditionDescriptions sets the matchedConditionDescriptions property value. The list of improper data in the message that was detected by the data loss prevention app. Each DLP app defines its own conditions, examples include 'Credit Card Number' and 'Social Security Number'.
func (m *ChatMessagePolicyViolationPolicyTip) SetMatchedConditionDescriptions(value []string)() {
    if m != nil {
        m.matchedConditionDescriptions = value
    }
}
