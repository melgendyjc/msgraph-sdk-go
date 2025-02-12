package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessClientApplications 
type ConditionalAccessClientApplications struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Service principal IDs excluded from the policy scope.
    excludeServicePrincipals []string
    // Service principal IDs included in the policy scope, or ServicePrincipalsInMyTenant.
    includeServicePrincipals []string
}
// NewConditionalAccessClientApplications instantiates a new conditionalAccessClientApplications and sets the default values.
func NewConditionalAccessClientApplications()(*ConditionalAccessClientApplications) {
    m := &ConditionalAccessClientApplications{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateConditionalAccessClientApplicationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessClientApplicationsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessClientApplications(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessClientApplications) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExcludeServicePrincipals gets the excludeServicePrincipals property value. Service principal IDs excluded from the policy scope.
func (m *ConditionalAccessClientApplications) GetExcludeServicePrincipals()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeServicePrincipals
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessClientApplications) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["excludeServicePrincipals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludeServicePrincipals(res)
        }
        return nil
    }
    res["includeServicePrincipals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIncludeServicePrincipals(res)
        }
        return nil
    }
    return res
}
// GetIncludeServicePrincipals gets the includeServicePrincipals property value. Service principal IDs included in the policy scope, or ServicePrincipalsInMyTenant.
func (m *ConditionalAccessClientApplications) GetIncludeServicePrincipals()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeServicePrincipals
    }
}
// Serialize serializes information the current object
func (m *ConditionalAccessClientApplications) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetExcludeServicePrincipals() != nil {
        err := writer.WriteCollectionOfStringValues("excludeServicePrincipals", m.GetExcludeServicePrincipals())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeServicePrincipals() != nil {
        err := writer.WriteCollectionOfStringValues("includeServicePrincipals", m.GetIncludeServicePrincipals())
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
func (m *ConditionalAccessClientApplications) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExcludeServicePrincipals sets the excludeServicePrincipals property value. Service principal IDs excluded from the policy scope.
func (m *ConditionalAccessClientApplications) SetExcludeServicePrincipals(value []string)() {
    if m != nil {
        m.excludeServicePrincipals = value
    }
}
// SetIncludeServicePrincipals sets the includeServicePrincipals property value. Service principal IDs included in the policy scope, or ServicePrincipalsInMyTenant.
func (m *ConditionalAccessClientApplications) SetIncludeServicePrincipals(value []string)() {
    if m != nil {
        m.includeServicePrincipals = value
    }
}
