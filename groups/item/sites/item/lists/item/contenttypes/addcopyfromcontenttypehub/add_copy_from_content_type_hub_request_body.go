package addcopyfromcontenttypehub

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AddCopyFromContentTypeHubRequestBody provides operations to call the addCopyFromContentTypeHub method.
type AddCopyFromContentTypeHubRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The contentTypeId property
    contentTypeId *string
}
// NewAddCopyFromContentTypeHubRequestBody instantiates a new addCopyFromContentTypeHubRequestBody and sets the default values.
func NewAddCopyFromContentTypeHubRequestBody()(*AddCopyFromContentTypeHubRequestBody) {
    m := &AddCopyFromContentTypeHubRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAddCopyFromContentTypeHubRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAddCopyFromContentTypeHubRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAddCopyFromContentTypeHubRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AddCopyFromContentTypeHubRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContentTypeId gets the contentTypeId property value. The contentTypeId property
func (m *AddCopyFromContentTypeHubRequestBody) GetContentTypeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentTypeId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AddCopyFromContentTypeHubRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentTypeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentTypeId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AddCopyFromContentTypeHubRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("contentTypeId", m.GetContentTypeId())
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
func (m *AddCopyFromContentTypeHubRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContentTypeId sets the contentTypeId property value. The contentTypeId property
func (m *AddCopyFromContentTypeHubRequestBody) SetContentTypeId(value *string)() {
    if m != nil {
        m.contentTypeId = value
    }
}