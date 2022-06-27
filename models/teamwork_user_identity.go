package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkUserIdentity 
type TeamworkUserIdentity struct {
    Identity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Type of user. Possible values are: aadUser, onPremiseAadUser, anonymousGuest, federatedUser, personalMicrosoftAccountUser, skypeUser, phoneUser, and emailUser.
    userIdentityType *TeamworkUserIdentityType
}
// NewTeamworkUserIdentity instantiates a new TeamworkUserIdentity and sets the default values.
func NewTeamworkUserIdentity()(*TeamworkUserIdentity) {
    m := &TeamworkUserIdentity{
        Identity: *NewIdentity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkUserIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkUserIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkUserIdentity(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkUserIdentity) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkUserIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["userIdentityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkUserIdentityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserIdentityType(val.(*TeamworkUserIdentityType))
        }
        return nil
    }
    return res
}
// GetUserIdentityType gets the userIdentityType property value. Type of user. Possible values are: aadUser, onPremiseAadUser, anonymousGuest, federatedUser, personalMicrosoftAccountUser, skypeUser, phoneUser, and emailUser.
func (m *TeamworkUserIdentity) GetUserIdentityType()(*TeamworkUserIdentityType) {
    if m == nil {
        return nil
    } else {
        return m.userIdentityType
    }
}
// Serialize serializes information the current object
func (m *TeamworkUserIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetUserIdentityType() != nil {
        cast := (*m.GetUserIdentityType()).String()
        err = writer.WriteStringValue("userIdentityType", &cast)
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
func (m *TeamworkUserIdentity) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetUserIdentityType sets the userIdentityType property value. Type of user. Possible values are: aadUser, onPremiseAadUser, anonymousGuest, federatedUser, personalMicrosoftAccountUser, skypeUser, phoneUser, and emailUser.
func (m *TeamworkUserIdentity) SetUserIdentityType(value *TeamworkUserIdentityType)() {
    if m != nil {
        m.userIdentityType = value
    }
}
