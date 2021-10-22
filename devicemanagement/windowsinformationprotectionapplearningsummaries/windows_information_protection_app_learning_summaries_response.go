package windowsinformationprotectionapplearningsummaries

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

type WindowsInformationProtectionAppLearningSummariesResponse struct {
    additionalData map[string]interface{};
    nextLink *string;
    value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WindowsInformationProtectionAppLearningSummary;
}
func NewWindowsInformationProtectionAppLearningSummariesResponse()(*WindowsInformationProtectionAppLearningSummariesResponse) {
    m := &WindowsInformationProtectionAppLearningSummariesResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *WindowsInformationProtectionAppLearningSummariesResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *WindowsInformationProtectionAppLearningSummariesResponse) GetNextLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.nextLink
    }
}
func (m *WindowsInformationProtectionAppLearningSummariesResponse) GetValue()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WindowsInformationProtectionAppLearningSummary) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *WindowsInformationProtectionAppLearningSummariesResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["@odata.nextLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNextLink(val)
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewWindowsInformationProtectionAppLearningSummary() })
        if err != nil {
            return err
        }
        res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WindowsInformationProtectionAppLearningSummary, len(val))
        for i, v := range val {
            res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WindowsInformationProtectionAppLearningSummary))
        }
        m.SetValue(res)
        return nil
    }
    return res
}
func (m *WindowsInformationProtectionAppLearningSummariesResponse) IsNil()(bool) {
    return m == nil
}
func (m *WindowsInformationProtectionAppLearningSummariesResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.nextLink", m.GetNextLink())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("value", cast)
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
func (m *WindowsInformationProtectionAppLearningSummariesResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *WindowsInformationProtectionAppLearningSummariesResponse) SetNextLink(value *string)() {
    m.nextLink = value
}
func (m *WindowsInformationProtectionAppLearningSummariesResponse) SetValue(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.WindowsInformationProtectionAppLearningSummary)() {
    m.value = value
}