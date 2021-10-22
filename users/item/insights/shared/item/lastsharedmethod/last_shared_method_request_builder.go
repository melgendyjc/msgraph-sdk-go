package lastsharedmethod

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i14988d2192ea3bfd876a73dc2fa29e9c461099775af2daa51550f01232119e17 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/mobileappcontentfile"
    i1c5e952cbcca5bb0561f018e495d472c32de97497ee9f2dcea5c10787062433c "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/calendarsharingmessage"
    i2013aba7e11d64583041b9ec3c14bb138531f49988b7fd4c16494ec690243282 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/managedappprotection"
    i2f35c4ffb6ad8416681d8d7927b754175812d1f6041468af8c5487826b458546 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrangefill"
    i35ead6f6b167894f8abf0266d628da9322792994e7578f2bf43f61edfdc9ac4b "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/schedulechangerequest"
    i7da25b7e9ee77c6e557103353c3815635abb71927408417ca27f7bff195a8f16 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/windowsinformationprotection"
    i81ee7201d40affa2819c85c7e237c667196e5ef9ea6a5e283dda2ebe3867ea5a "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrange"
    i84c2d38d04706be0941cc9c59b2e39f93592c499b136c85fd6ce5f5e98c076d0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrangeview"
    i8eb07e2f7d7178edc6434ab76c55df72cdc2df0543bc76dd1d3b154c9b9a8033 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrangeformat"
    i8ed57c30df939e8f36b03587b38e3aee9c462b27d5b70818a755e551fcb16cd1 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/printdocument"
    iaf57896b88d0934211fbf881a2361740c9f6794e739c1ffc20fa7d6bc71c5094 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/printjob"
    ib7cb7a354159d58663d9e865787ccb50cd54059c0690453d12d4d30f13c42e4b "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/targetedmanagedappprotection"
    ieb46782fb9bc2b505892f37aff6b97bb1db3556585144877b6695d5cb70d8a20 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/ref"
    if00fbcd158c4e9a59d7f77fd23ba8ec76eac01483cef335d8ea76e4fdaf651fb "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrangesort"
)

type LastSharedMethodRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type LastSharedMethodRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *LastSharedMethodRequestBuilder) CalendarSharingMessage()(*i1c5e952cbcca5bb0561f018e495d472c32de97497ee9f2dcea5c10787062433c.CalendarSharingMessageRequestBuilder) {
    return i1c5e952cbcca5bb0561f018e495d472c32de97497ee9f2dcea5c10787062433c.NewCalendarSharingMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewLastSharedMethodRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*LastSharedMethodRequestBuilder) {
    m := &LastSharedMethodRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/insights/shared/{sharedInsight_id}/lastSharedMethod{?select,expand}";
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
func NewLastSharedMethodRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*LastSharedMethodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLastSharedMethodRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *LastSharedMethodRequestBuilder) CreateGetRequestInformation(q func (value *LastSharedMethodRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(LastSharedMethodRequestBuilderGetQueryParameters)
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
func (m *LastSharedMethodRequestBuilder) Get(q func (value *LastSharedMethodRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Entity, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEntity() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Entity), nil
}
func (m *LastSharedMethodRequestBuilder) ManagedAppProtection()(*i2013aba7e11d64583041b9ec3c14bb138531f49988b7fd4c16494ec690243282.ManagedAppProtectionRequestBuilder) {
    return i2013aba7e11d64583041b9ec3c14bb138531f49988b7fd4c16494ec690243282.NewManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) MobileAppContentFile()(*i14988d2192ea3bfd876a73dc2fa29e9c461099775af2daa51550f01232119e17.MobileAppContentFileRequestBuilder) {
    return i14988d2192ea3bfd876a73dc2fa29e9c461099775af2daa51550f01232119e17.NewMobileAppContentFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) PrintDocument()(*i8ed57c30df939e8f36b03587b38e3aee9c462b27d5b70818a755e551fcb16cd1.PrintDocumentRequestBuilder) {
    return i8ed57c30df939e8f36b03587b38e3aee9c462b27d5b70818a755e551fcb16cd1.NewPrintDocumentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) PrintJob()(*iaf57896b88d0934211fbf881a2361740c9f6794e739c1ffc20fa7d6bc71c5094.PrintJobRequestBuilder) {
    return iaf57896b88d0934211fbf881a2361740c9f6794e739c1ffc20fa7d6bc71c5094.NewPrintJobRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) Ref()(*ieb46782fb9bc2b505892f37aff6b97bb1db3556585144877b6695d5cb70d8a20.RefRequestBuilder) {
    return ieb46782fb9bc2b505892f37aff6b97bb1db3556585144877b6695d5cb70d8a20.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) ScheduleChangeRequest()(*i35ead6f6b167894f8abf0266d628da9322792994e7578f2bf43f61edfdc9ac4b.ScheduleChangeRequestRequestBuilder) {
    return i35ead6f6b167894f8abf0266d628da9322792994e7578f2bf43f61edfdc9ac4b.NewScheduleChangeRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) TargetedManagedAppProtection()(*ib7cb7a354159d58663d9e865787ccb50cd54059c0690453d12d4d30f13c42e4b.TargetedManagedAppProtectionRequestBuilder) {
    return ib7cb7a354159d58663d9e865787ccb50cd54059c0690453d12d4d30f13c42e4b.NewTargetedManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) WindowsInformationProtection()(*i7da25b7e9ee77c6e557103353c3815635abb71927408417ca27f7bff195a8f16.WindowsInformationProtectionRequestBuilder) {
    return i7da25b7e9ee77c6e557103353c3815635abb71927408417ca27f7bff195a8f16.NewWindowsInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) WorkbookRange()(*i81ee7201d40affa2819c85c7e237c667196e5ef9ea6a5e283dda2ebe3867ea5a.WorkbookRangeRequestBuilder) {
    return i81ee7201d40affa2819c85c7e237c667196e5ef9ea6a5e283dda2ebe3867ea5a.NewWorkbookRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) WorkbookRangeFill()(*i2f35c4ffb6ad8416681d8d7927b754175812d1f6041468af8c5487826b458546.WorkbookRangeFillRequestBuilder) {
    return i2f35c4ffb6ad8416681d8d7927b754175812d1f6041468af8c5487826b458546.NewWorkbookRangeFillRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) WorkbookRangeFormat()(*i8eb07e2f7d7178edc6434ab76c55df72cdc2df0543bc76dd1d3b154c9b9a8033.WorkbookRangeFormatRequestBuilder) {
    return i8eb07e2f7d7178edc6434ab76c55df72cdc2df0543bc76dd1d3b154c9b9a8033.NewWorkbookRangeFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) WorkbookRangeSort()(*if00fbcd158c4e9a59d7f77fd23ba8ec76eac01483cef335d8ea76e4fdaf651fb.WorkbookRangeSortRequestBuilder) {
    return if00fbcd158c4e9a59d7f77fd23ba8ec76eac01483cef335d8ea76e4fdaf651fb.NewWorkbookRangeSortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LastSharedMethodRequestBuilder) WorkbookRangeView()(*i84c2d38d04706be0941cc9c59b2e39f93592c499b136c85fd6ce5f5e98c076d0.WorkbookRangeViewRequestBuilder) {
    return i84c2d38d04706be0941cc9c59b2e39f93592c499b136c85fd6ce5f5e98c076d0.NewWorkbookRangeViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}