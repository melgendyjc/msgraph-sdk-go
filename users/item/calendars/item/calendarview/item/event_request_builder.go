package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i2541624cc2b7c2f09531c622ffbb42f4a195d101ad0ad14960f1f7e657102dbc "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/extensions"
    i2a0da227fb026925fc8b27a16ecbf8cc55f0b9a76e5549734baa76307a0bb83e "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/multivalueextendedproperties"
    i3d8ae6cb74af0972786738f46478813a75154a5fe06967bf203a2e0918eabee5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/dismissreminder"
    i4ea184a6331e61d7549b8b34404e661ecdeb5805dfd8c57e0c1b3b015dffc03c "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/attachments"
    i5a19924c15817bb8517dee2b633e6419ee084968aef5fcaf92f58cd7172a31ed "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/decline"
    i73cc31439204381dc3d6dc91ac06813562311395172df52ccc39264157757675 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/snoozereminder"
    i76172f12b34f51d0a11592d73ad0199c16d4ef8ebf8706eb0c1a27110f6d1f45 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/accept"
    i8432ac922f3623c62732b850934ae00f8014ebaabc1bb287cef0cef6433702f1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/singlevalueextendedproperties"
    i8631ad4f49ade0f7640c89e9d89155294a2c76a8fb7649c43e851a03b4311aee "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/forward"
    i8a87a0c791b5a48fd3b0ec426dab5c995b07ba2ee236b6f6b8a734cc5bf6c13d "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/cancel"
    i8d260fb96002fb446c710037c6bf1fb37f6ab3fbde3346f2af2a637e926d6baf "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/instances"
    ia4474e31f9f6f53023bfbed80dd5f789939a7bbae9e7a395838fa9a4bae43623 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/calendar"
    ie4908e0e67bcb1af8d1ee415aaefee36f4683dd9af729a4ffdc661090e90eb73 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/tentativelyaccept"
    i0a374527704358b2f886240ad18b1b8556ada3a535e6ffcd229d15b9decf0a3e "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/attachments/item"
    i2c1f1efe93a31892c62f2d27507ea8bc2b85dd0749dc109db8a9be7bc576746c "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/multivalueextendedproperties/item"
    i361917cefafc439600ed28778b8652cd53d5430c792eb398c7d41102efe1a81e "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/extensions/item"
    ib6f4ea864b364faa0da40819efa0bbe6984fece1d20e43e11a87d621c451a0b0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/instances/item"
    id28d63750d1254314b44fe7205991f1cd9bc0de925b89d62436af5889d20a03a "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item/singlevalueextendedproperties/item"
)

type EventRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type EventRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    EndDateTime *string;
    Expand []string;
    Select_escpaped []string;
    StartDateTime *string;
}
func (m *EventRequestBuilder) Accept()(*i76172f12b34f51d0a11592d73ad0199c16d4ef8ebf8706eb0c1a27110f6d1f45.AcceptRequestBuilder) {
    return i76172f12b34f51d0a11592d73ad0199c16d4ef8ebf8706eb0c1a27110f6d1f45.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i4ea184a6331e61d7549b8b34404e661ecdeb5805dfd8c57e0c1b3b015dffc03c.AttachmentsRequestBuilder) {
    return i4ea184a6331e61d7549b8b34404e661ecdeb5805dfd8c57e0c1b3b015dffc03c.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) AttachmentsById(id string)(*i0a374527704358b2f886240ad18b1b8556ada3a535e6ffcd229d15b9decf0a3e.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i0a374527704358b2f886240ad18b1b8556ada3a535e6ffcd229d15b9decf0a3e.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*ia4474e31f9f6f53023bfbed80dd5f789939a7bbae9e7a395838fa9a4bae43623.CalendarRequestBuilder) {
    return ia4474e31f9f6f53023bfbed80dd5f789939a7bbae9e7a395838fa9a4bae43623.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i8a87a0c791b5a48fd3b0ec426dab5c995b07ba2ee236b6f6b8a734cc5bf6c13d.CancelRequestBuilder) {
    return i8a87a0c791b5a48fd3b0ec426dab5c995b07ba2ee236b6f6b8a734cc5bf6c13d.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/calendars/{calendar_id}/calendarView/{event_id}{?startDateTime,endDateTime,select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewEventRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *EventRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *EventRequestBuilder) CreateGetRequestInformation(q func (value *EventRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(EventRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *EventRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *EventRequestBuilder) Decline()(*i5a19924c15817bb8517dee2b633e6419ee084968aef5fcaf92f58cd7172a31ed.DeclineRequestBuilder) {
    return i5a19924c15817bb8517dee2b633e6419ee084968aef5fcaf92f58cd7172a31ed.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventRequestBuilder) DismissReminder()(*i3d8ae6cb74af0972786738f46478813a75154a5fe06967bf203a2e0918eabee5.DismissReminderRequestBuilder) {
    return i3d8ae6cb74af0972786738f46478813a75154a5fe06967bf203a2e0918eabee5.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*i2541624cc2b7c2f09531c622ffbb42f4a195d101ad0ad14960f1f7e657102dbc.ExtensionsRequestBuilder) {
    return i2541624cc2b7c2f09531c622ffbb42f4a195d101ad0ad14960f1f7e657102dbc.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExtensionsById(id string)(*i361917cefafc439600ed28778b8652cd53d5430c792eb398c7d41102efe1a81e.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i361917cefafc439600ed28778b8652cd53d5430c792eb398c7d41102efe1a81e.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i8631ad4f49ade0f7640c89e9d89155294a2c76a8fb7649c43e851a03b4311aee.ForwardRequestBuilder) {
    return i8631ad4f49ade0f7640c89e9d89155294a2c76a8fb7649c43e851a03b4311aee.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Get(q func (value *EventRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEvent() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event), nil
}
func (m *EventRequestBuilder) Instances()(*i8d260fb96002fb446c710037c6bf1fb37f6ab3fbde3346f2af2a637e926d6baf.InstancesRequestBuilder) {
    return i8d260fb96002fb446c710037c6bf1fb37f6ab3fbde3346f2af2a637e926d6baf.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) InstancesById(id string)(*ib6f4ea864b364faa0da40819efa0bbe6984fece1d20e43e11a87d621c451a0b0.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return ib6f4ea864b364faa0da40819efa0bbe6984fece1d20e43e11a87d621c451a0b0.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*i2a0da227fb026925fc8b27a16ecbf8cc55f0b9a76e5549734baa76307a0bb83e.MultiValueExtendedPropertiesRequestBuilder) {
    return i2a0da227fb026925fc8b27a16ecbf8cc55f0b9a76e5549734baa76307a0bb83e.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i2c1f1efe93a31892c62f2d27507ea8bc2b85dd0749dc109db8a9be7bc576746c.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i2c1f1efe93a31892c62f2d27507ea8bc2b85dd0749dc109db8a9be7bc576746c.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*i8432ac922f3623c62732b850934ae00f8014ebaabc1bb287cef0cef6433702f1.SingleValueExtendedPropertiesRequestBuilder) {
    return i8432ac922f3623c62732b850934ae00f8014ebaabc1bb287cef0cef6433702f1.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*id28d63750d1254314b44fe7205991f1cd9bc0de925b89d62436af5889d20a03a.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return id28d63750d1254314b44fe7205991f1cd9bc0de925b89d62436af5889d20a03a.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i73cc31439204381dc3d6dc91ac06813562311395172df52ccc39264157757675.SnoozeReminderRequestBuilder) {
    return i73cc31439204381dc3d6dc91ac06813562311395172df52ccc39264157757675.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*ie4908e0e67bcb1af8d1ee415aaefee36f4683dd9af729a4ffdc661090e90eb73.TentativelyAcceptRequestBuilder) {
    return ie4908e0e67bcb1af8d1ee415aaefee36f4683dd9af729a4ffdc661090e90eb73.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}