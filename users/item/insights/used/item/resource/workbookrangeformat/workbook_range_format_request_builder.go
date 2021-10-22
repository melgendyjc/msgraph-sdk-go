package workbookrangeformat

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i1163a064374e5b4998d872e16f60b44ac2f77534f08c7e662a0174d9fa7b7b3e "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrangeformat/autofitcolumns"
    i9224fac1a0bf495158a80afcaab0329aa2ef9cab5bc4bde72cf8a43fb8469bf2 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrangeformat/autofitrows"
)

type WorkbookRangeFormatRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *WorkbookRangeFormatRequestBuilder) AutofitColumns()(*i1163a064374e5b4998d872e16f60b44ac2f77534f08c7e662a0174d9fa7b7b3e.AutofitColumnsRequestBuilder) {
    return i1163a064374e5b4998d872e16f60b44ac2f77534f08c7e662a0174d9fa7b7b3e.NewAutofitColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeFormatRequestBuilder) AutofitRows()(*i9224fac1a0bf495158a80afcaab0329aa2ef9cab5bc4bde72cf8a43fb8469bf2.AutofitRowsRequestBuilder) {
    return i9224fac1a0bf495158a80afcaab0329aa2ef9cab5bc4bde72cf8a43fb8469bf2.NewAutofitRowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewWorkbookRangeFormatRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFormatRequestBuilder) {
    m := &WorkbookRangeFormatRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/insights/used/{usedInsight_id}/resource/microsoft.graph.workbookRangeFormat";
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
func NewWorkbookRangeFormatRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFormatRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookRangeFormatRequestBuilderInternal(urlParams, requestAdapter)
}