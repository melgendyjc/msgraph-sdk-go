package directoryroles

import (
    ia4917d7d2d6bbb2d83014b7111cd4d79a23adb1d36bc4e4126b6166e1d29afa8 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/getavailableextensionproperties"
    ia87332b05b41e0506ca217464387e3ed5c7d91dea72449a3a8af52c972c79b03 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/getbyids"
    id178ea9892b7517ea8099dc2841e11c78f3dfdc11a8a6c2abd1433ac224172e1 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/delta"
    ida2bdcfe4d6bf2bb544752c17e1ca5059ff3edc052c3080051a9c874aeb8aa1a "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/validateproperties"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

type DirectoryRolesRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type DirectoryRolesRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Count *bool;
    Expand []string;
    Filter *string;
    Orderby []string;
    Search *string;
    Select_escpaped []string;
    Skip *int32;
    Top *int32;
}
func NewDirectoryRolesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRolesRequestBuilder) {
    m := &DirectoryRolesRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/directoryRoles{?top,skip,search,filter,count,orderby,select,expand}";
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
func NewDirectoryRolesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRolesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRolesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DirectoryRolesRequestBuilder) CreateGetRequestInformation(q func (value *DirectoryRolesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(DirectoryRolesRequestBuilderGetQueryParameters)
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
func (m *DirectoryRolesRequestBuilder) CreatePostRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRole, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
func (m *DirectoryRolesRequestBuilder) Delta()(*id178ea9892b7517ea8099dc2841e11c78f3dfdc11a8a6c2abd1433ac224172e1.DeltaRequestBuilder) {
    return id178ea9892b7517ea8099dc2841e11c78f3dfdc11a8a6c2abd1433ac224172e1.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRolesRequestBuilder) Get(q func (value *DirectoryRolesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*DirectoryRolesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryRolesResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*DirectoryRolesResponse), nil
}
func (m *DirectoryRolesRequestBuilder) GetAvailableExtensionProperties()(*ia4917d7d2d6bbb2d83014b7111cd4d79a23adb1d36bc4e4126b6166e1d29afa8.GetAvailableExtensionPropertiesRequestBuilder) {
    return ia4917d7d2d6bbb2d83014b7111cd4d79a23adb1d36bc4e4126b6166e1d29afa8.NewGetAvailableExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRolesRequestBuilder) GetByIds()(*ia87332b05b41e0506ca217464387e3ed5c7d91dea72449a3a8af52c972c79b03.GetByIdsRequestBuilder) {
    return ia87332b05b41e0506ca217464387e3ed5c7d91dea72449a3a8af52c972c79b03.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRolesRequestBuilder) Post(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRole, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRole, error) {
    requestInfo, err := m.CreatePostRequestInformation(body, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDirectoryRole() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRole), nil
}
func (m *DirectoryRolesRequestBuilder) ValidateProperties()(*ida2bdcfe4d6bf2bb544752c17e1ca5059ff3edc052c3080051a9c874aeb8aa1a.ValidatePropertiesRequestBuilder) {
    return ida2bdcfe4d6bf2bb544752c17e1ca5059ff3edc052c3080051a9c874aeb8aa1a.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}