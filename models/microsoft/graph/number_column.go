package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type NumberColumn struct {
    additionalData map[string]interface{};
    decimalPlaces *string;
    displayAs *string;
    maximum *float64;
    minimum *float64;
}
func NewNumberColumn()(*NumberColumn) {
    m := &NumberColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *NumberColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *NumberColumn) GetDecimalPlaces()(*string) {
    if m == nil {
        return nil
    } else {
        return m.decimalPlaces
    }
}
func (m *NumberColumn) GetDisplayAs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayAs
    }
}
func (m *NumberColumn) GetMaximum()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.maximum
    }
}
func (m *NumberColumn) GetMinimum()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.minimum
    }
}
func (m *NumberColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["decimalPlaces"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDecimalPlaces(val)
        return nil
    }
    res["displayAs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayAs(val)
        return nil
    }
    res["maximum"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetMaximum(val)
        return nil
    }
    res["minimum"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetMinimum(val)
        return nil
    }
    return res
}
func (m *NumberColumn) IsNil()(bool) {
    return m == nil
}
func (m *NumberColumn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("decimalPlaces", m.GetDecimalPlaces())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayAs", m.GetDisplayAs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("maximum", m.GetMaximum())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("minimum", m.GetMinimum())
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
func (m *NumberColumn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *NumberColumn) SetDecimalPlaces(value *string)() {
    m.decimalPlaces = value
}
func (m *NumberColumn) SetDisplayAs(value *string)() {
    m.displayAs = value
}
func (m *NumberColumn) SetMaximum(value *float64)() {
    m.maximum = value
}
func (m *NumberColumn) SetMinimum(value *float64)() {
    m.minimum = value
}