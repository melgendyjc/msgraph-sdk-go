package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookTableColumn provides operations to manage the collection of agreementAcceptance entities.
type WorkbookTableColumn struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Retrieve the filter applied to the column. Read-only.
    filter WorkbookFilterable
    // Returns the index number of the column within the columns collection of the table. Zero-indexed. Read-only.
    index *int32
    // Returns the name of the table column.
    name *string
    // Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
    values Jsonable
}
// NewWorkbookTableColumn instantiates a new workbookTableColumn and sets the default values.
func NewWorkbookTableColumn()(*WorkbookTableColumn) {
    m := &WorkbookTableColumn{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWorkbookTableColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookTableColumnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookTableColumn(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbookTableColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookTableColumn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["filter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookFilterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilter(val.(WorkbookFilterable))
        }
        return nil
    }
    res["index"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndex(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValues(val.(Jsonable))
        }
        return nil
    }
    return res
}
// GetFilter gets the filter property value. Retrieve the filter applied to the column. Read-only.
func (m *WorkbookTableColumn) GetFilter()(WorkbookFilterable) {
    if m == nil {
        return nil
    } else {
        return m.filter
    }
}
// GetIndex gets the index property value. Returns the index number of the column within the columns collection of the table. Zero-indexed. Read-only.
func (m *WorkbookTableColumn) GetIndex()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.index
    }
}
// GetName gets the name property value. Returns the name of the table column.
func (m *WorkbookTableColumn) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetValues gets the values property value. Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
func (m *WorkbookTableColumn) GetValues()(Jsonable) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
// Serialize serializes information the current object
func (m *WorkbookTableColumn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("filter", m.GetFilter())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("index", m.GetIndex())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("values", m.GetValues())
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
func (m *WorkbookTableColumn) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFilter sets the filter property value. Retrieve the filter applied to the column. Read-only.
func (m *WorkbookTableColumn) SetFilter(value WorkbookFilterable)() {
    if m != nil {
        m.filter = value
    }
}
// SetIndex sets the index property value. Returns the index number of the column within the columns collection of the table. Zero-indexed. Read-only.
func (m *WorkbookTableColumn) SetIndex(value *int32)() {
    if m != nil {
        m.index = value
    }
}
// SetName sets the name property value. Returns the name of the table column.
func (m *WorkbookTableColumn) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetValues sets the values property value. Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
func (m *WorkbookTableColumn) SetValues(value Jsonable)() {
    if m != nil {
        m.values = value
    }
}
