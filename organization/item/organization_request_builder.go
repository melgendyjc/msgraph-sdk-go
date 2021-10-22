package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2c3e7e24e2730178382d4adade068738ce6025ae386d5724c24a3a268a81525c "github.com/microsoftgraph/msgraph-sdk-go/organization/item/getmembergroups"
    i3d6521f9c6af801263d8177b58bd3f14c44b32d31606bcefa9d9284325e9ae05 "github.com/microsoftgraph/msgraph-sdk-go/organization/item/checkmemberobjects"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i62fcec4c98d8f7e9caa7bd4b260ea9d20f54a28de6683ff87d4101316c41f38a "github.com/microsoftgraph/msgraph-sdk-go/organization/item/extensions"
    i6be0b4e2753e0c9bfc4e1158ef7d0165dbba9087c2342ddf87c85eb2609049e8 "github.com/microsoftgraph/msgraph-sdk-go/organization/item/branding"
    i6d71976a15ac6f607366bcc5307ea5cc985e145be0c3586755194917f7fb4169 "github.com/microsoftgraph/msgraph-sdk-go/organization/item/checkmembergroups"
    i8885d5cee4c4d28a727da7b8bbf4e0a830b801820a872d899634b98dea40da8c "github.com/microsoftgraph/msgraph-sdk-go/organization/item/getmemberobjects"
    id705df728183d1caa2cd0a8e71b54c1f25c4bf34951aa4d58d74debd1d27c3d8 "github.com/microsoftgraph/msgraph-sdk-go/organization/item/restore"
    ieb9078c75a28c74bd2df523a3384718424b13d8e1f2c2fb11f57ab22d0557625 "github.com/microsoftgraph/msgraph-sdk-go/organization/item/setmobiledevicemanagementauthority"
    iee043428b7b7f62f2e5471e1c8a277c3f71551ba1ea8cb7e6ba5dbb3e2b0607d "github.com/microsoftgraph/msgraph-sdk-go/organization/item/certificatebasedauthconfiguration"
    iead5f0a6000982961cff6f0e297f9427177afc718378c21027082d54c9ed6105 "github.com/microsoftgraph/msgraph-sdk-go/organization/item/extensions/item"
)

type OrganizationRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type OrganizationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *OrganizationRequestBuilder) Branding()(*i6be0b4e2753e0c9bfc4e1158ef7d0165dbba9087c2342ddf87c85eb2609049e8.BrandingRequestBuilder) {
    return i6be0b4e2753e0c9bfc4e1158ef7d0165dbba9087c2342ddf87c85eb2609049e8.NewBrandingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrganizationRequestBuilder) CertificateBasedAuthConfiguration()(*iee043428b7b7f62f2e5471e1c8a277c3f71551ba1ea8cb7e6ba5dbb3e2b0607d.CertificateBasedAuthConfigurationRequestBuilder) {
    return iee043428b7b7f62f2e5471e1c8a277c3f71551ba1ea8cb7e6ba5dbb3e2b0607d.NewCertificateBasedAuthConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrganizationRequestBuilder) CheckMemberGroups()(*i6d71976a15ac6f607366bcc5307ea5cc985e145be0c3586755194917f7fb4169.CheckMemberGroupsRequestBuilder) {
    return i6d71976a15ac6f607366bcc5307ea5cc985e145be0c3586755194917f7fb4169.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrganizationRequestBuilder) CheckMemberObjects()(*i3d6521f9c6af801263d8177b58bd3f14c44b32d31606bcefa9d9284325e9ae05.CheckMemberObjectsRequestBuilder) {
    return i3d6521f9c6af801263d8177b58bd3f14c44b32d31606bcefa9d9284325e9ae05.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewOrganizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OrganizationRequestBuilder) {
    m := &OrganizationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/organization/{organization_id}{?select,expand}";
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
func NewOrganizationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OrganizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrganizationRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OrganizationRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *OrganizationRequestBuilder) CreateGetRequestInformation(q func (value *OrganizationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(OrganizationRequestBuilderGetQueryParameters)
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
func (m *OrganizationRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Organization, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *OrganizationRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *OrganizationRequestBuilder) Extensions()(*i62fcec4c98d8f7e9caa7bd4b260ea9d20f54a28de6683ff87d4101316c41f38a.ExtensionsRequestBuilder) {
    return i62fcec4c98d8f7e9caa7bd4b260ea9d20f54a28de6683ff87d4101316c41f38a.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrganizationRequestBuilder) ExtensionsById(id string)(*iead5f0a6000982961cff6f0e297f9427177afc718378c21027082d54c9ed6105.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return iead5f0a6000982961cff6f0e297f9427177afc718378c21027082d54c9ed6105.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OrganizationRequestBuilder) Get(q func (value *OrganizationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Organization, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewOrganization() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Organization), nil
}
func (m *OrganizationRequestBuilder) GetMemberGroups()(*i2c3e7e24e2730178382d4adade068738ce6025ae386d5724c24a3a268a81525c.GetMemberGroupsRequestBuilder) {
    return i2c3e7e24e2730178382d4adade068738ce6025ae386d5724c24a3a268a81525c.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrganizationRequestBuilder) GetMemberObjects()(*i8885d5cee4c4d28a727da7b8bbf4e0a830b801820a872d899634b98dea40da8c.GetMemberObjectsRequestBuilder) {
    return i8885d5cee4c4d28a727da7b8bbf4e0a830b801820a872d899634b98dea40da8c.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrganizationRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Organization, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *OrganizationRequestBuilder) Restore()(*id705df728183d1caa2cd0a8e71b54c1f25c4bf34951aa4d58d74debd1d27c3d8.RestoreRequestBuilder) {
    return id705df728183d1caa2cd0a8e71b54c1f25c4bf34951aa4d58d74debd1d27c3d8.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrganizationRequestBuilder) SetMobileDeviceManagementAuthority()(*ieb9078c75a28c74bd2df523a3384718424b13d8e1f2c2fb11f57ab22d0557625.SetMobileDeviceManagementAuthorityRequestBuilder) {
    return ieb9078c75a28c74bd2df523a3384718424b13d8e1f2c2fb11f57ab22d0557625.NewSetMobileDeviceManagementAuthorityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}