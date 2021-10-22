package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type InvitationParticipantInfo struct {
    additionalData map[string]interface{};
    identity *IdentitySet;
    replacesCallId *string;
}
func NewInvitationParticipantInfo()(*InvitationParticipantInfo) {
    m := &InvitationParticipantInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *InvitationParticipantInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *InvitationParticipantInfo) GetIdentity()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.identity
    }
}
func (m *InvitationParticipantInfo) GetReplacesCallId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.replacesCallId
    }
}
func (m *InvitationParticipantInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["identity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetIdentity(val.(*IdentitySet))
        return nil
    }
    res["replacesCallId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReplacesCallId(val)
        return nil
    }
    return res
}
func (m *InvitationParticipantInfo) IsNil()(bool) {
    return m == nil
}
func (m *InvitationParticipantInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("identity", m.GetIdentity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("replacesCallId", m.GetReplacesCallId())
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
func (m *InvitationParticipantInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *InvitationParticipantInfo) SetIdentity(value *IdentitySet)() {
    m.identity = value
}
func (m *InvitationParticipantInfo) SetReplacesCallId(value *string)() {
    m.replacesCallId = value
}