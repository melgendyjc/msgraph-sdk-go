package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityResource 
type SecurityResource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Name of the resource that is related to current alert. Required.
    resource *string
    // Represents type of security resources related to an alert. Possible values are: attacked, related.
    resourceType *SecurityResourceType
}
// NewSecurityResource instantiates a new securityResource and sets the default values.
func NewSecurityResource()(*SecurityResource) {
    m := &SecurityResource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSecurityResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityResource(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityResource) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val)
        }
        return nil
    }
    res["resourceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityResourceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceType(val.(*SecurityResourceType))
        }
        return nil
    }
    return res
}
// GetResource gets the resource property value. Name of the resource that is related to current alert. Required.
func (m *SecurityResource) GetResource()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// GetResourceType gets the resourceType property value. Represents type of security resources related to an alert. Possible values are: attacked, related.
func (m *SecurityResource) GetResourceType()(*SecurityResourceType) {
    if m == nil {
        return nil
    } else {
        return m.resourceType
    }
}
// Serialize serializes information the current object
func (m *SecurityResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    if m.GetResourceType() != nil {
        cast := (*m.GetResourceType()).String()
        err := writer.WriteStringValue("resourceType", &cast)
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
func (m *SecurityResource) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetResource sets the resource property value. Name of the resource that is related to current alert. Required.
func (m *SecurityResource) SetResource(value *string)() {
    if m != nil {
        m.resource = value
    }
}
// SetResourceType sets the resourceType property value. Represents type of security resources related to an alert. Possible values are: attacked, related.
func (m *SecurityResource) SetResourceType(value *SecurityResourceType)() {
    if m != nil {
        m.resourceType = value
    }
}
