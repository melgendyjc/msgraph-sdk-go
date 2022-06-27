package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdgeSearchEngineCustom 
type EdgeSearchEngineCustom struct {
    EdgeSearchEngineBase
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Points to a https link containing the OpenSearch xml file that contains, at minimum, the short name and the URL to the search Engine.
    edgeSearchEngineOpenSearchXmlUrl *string
}
// NewEdgeSearchEngineCustom instantiates a new EdgeSearchEngineCustom and sets the default values.
func NewEdgeSearchEngineCustom()(*EdgeSearchEngineCustom) {
    m := &EdgeSearchEngineCustom{
        EdgeSearchEngineBase: *NewEdgeSearchEngineBase(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEdgeSearchEngineCustomFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdgeSearchEngineCustomFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdgeSearchEngineCustom(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EdgeSearchEngineCustom) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEdgeSearchEngineOpenSearchXmlUrl gets the edgeSearchEngineOpenSearchXmlUrl property value. Points to a https link containing the OpenSearch xml file that contains, at minimum, the short name and the URL to the search Engine.
func (m *EdgeSearchEngineCustom) GetEdgeSearchEngineOpenSearchXmlUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.edgeSearchEngineOpenSearchXmlUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdgeSearchEngineCustom) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EdgeSearchEngineBase.GetFieldDeserializers()
    res["edgeSearchEngineOpenSearchXmlUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeSearchEngineOpenSearchXmlUrl(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *EdgeSearchEngineCustom) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EdgeSearchEngineBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("edgeSearchEngineOpenSearchXmlUrl", m.GetEdgeSearchEngineOpenSearchXmlUrl())
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
func (m *EdgeSearchEngineCustom) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEdgeSearchEngineOpenSearchXmlUrl sets the edgeSearchEngineOpenSearchXmlUrl property value. Points to a https link containing the OpenSearch xml file that contains, at minimum, the short name and the URL to the search Engine.
func (m *EdgeSearchEngineCustom) SetEdgeSearchEngineOpenSearchXmlUrl(value *string)() {
    if m != nil {
        m.edgeSearchEngineOpenSearchXmlUrl = value
    }
}
