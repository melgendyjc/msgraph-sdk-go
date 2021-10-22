package logteleconferencedevicequality

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

type LogTeleconferenceDeviceQualityRequestBody struct {
    additionalData map[string]interface{};
    quality *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TeleconferenceDeviceQuality;
}
func NewLogTeleconferenceDeviceQualityRequestBody()(*LogTeleconferenceDeviceQualityRequestBody) {
    m := &LogTeleconferenceDeviceQualityRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *LogTeleconferenceDeviceQualityRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *LogTeleconferenceDeviceQualityRequestBody) GetQuality()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TeleconferenceDeviceQuality) {
    if m == nil {
        return nil
    } else {
        return m.quality
    }
}
func (m *LogTeleconferenceDeviceQualityRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["quality"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewTeleconferenceDeviceQuality() })
        if err != nil {
            return err
        }
        m.SetQuality(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TeleconferenceDeviceQuality))
        return nil
    }
    return res
}
func (m *LogTeleconferenceDeviceQualityRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *LogTeleconferenceDeviceQualityRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("quality", m.GetQuality())
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
func (m *LogTeleconferenceDeviceQualityRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *LogTeleconferenceDeviceQualityRequestBody) SetQuality(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TeleconferenceDeviceQuality)() {
    m.quality = value
}