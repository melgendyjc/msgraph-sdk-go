package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ParticipantLeftNotification 
type ParticipantLeftNotification struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The call property
    call Callable
    // ID of the participant under the policy who has left the meeting.
    participantId *string
}
// NewParticipantLeftNotification instantiates a new ParticipantLeftNotification and sets the default values.
func NewParticipantLeftNotification()(*ParticipantLeftNotification) {
    m := &ParticipantLeftNotification{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateParticipantLeftNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateParticipantLeftNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewParticipantLeftNotification(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ParticipantLeftNotification) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCall gets the call property value. The call property
func (m *ParticipantLeftNotification) GetCall()(Callable) {
    if m == nil {
        return nil
    } else {
        return m.call
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ParticipantLeftNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["call"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCallFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCall(val.(Callable))
        }
        return nil
    }
    res["participantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipantId(val)
        }
        return nil
    }
    return res
}
// GetParticipantId gets the participantId property value. ID of the participant under the policy who has left the meeting.
func (m *ParticipantLeftNotification) GetParticipantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.participantId
    }
}
// Serialize serializes information the current object
func (m *ParticipantLeftNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("call", m.GetCall())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("participantId", m.GetParticipantId())
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
func (m *ParticipantLeftNotification) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCall sets the call property value. The call property
func (m *ParticipantLeftNotification) SetCall(value Callable)() {
    if m != nil {
        m.call = value
    }
}
// SetParticipantId sets the participantId property value. ID of the participant under the policy who has left the meeting.
func (m *ParticipantLeftNotification) SetParticipantId(value *string)() {
    if m != nil {
        m.participantId = value
    }
}
