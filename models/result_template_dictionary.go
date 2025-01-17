package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ResultTemplateDictionary provides operations to call the query method.
type ResultTemplateDictionary struct {
    Dictionary
}
// NewResultTemplateDictionary instantiates a new resultTemplateDictionary and sets the default values.
func NewResultTemplateDictionary()(*ResultTemplateDictionary) {
    m := &ResultTemplateDictionary{
        Dictionary: *NewDictionary(),
    }
    return m
}
// CreateResultTemplateDictionaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResultTemplateDictionaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResultTemplateDictionary(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResultTemplateDictionary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Dictionary.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ResultTemplateDictionary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Dictionary.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
