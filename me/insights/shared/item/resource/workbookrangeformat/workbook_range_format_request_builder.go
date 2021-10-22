package workbookrangeformat

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i59be4ff9e5cfa84f12c10e3f2b052ab0062298e7a5f88045e660ddf3fd33193b "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrangeformat/autofitrows"
    ic7d62f81fffb8eeca19a8b0ef08616810723885319cdf8d99c7263146be1edcd "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrangeformat/autofitcolumns"
)

type WorkbookRangeFormatRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *WorkbookRangeFormatRequestBuilder) AutofitColumns()(*ic7d62f81fffb8eeca19a8b0ef08616810723885319cdf8d99c7263146be1edcd.AutofitColumnsRequestBuilder) {
    return ic7d62f81fffb8eeca19a8b0ef08616810723885319cdf8d99c7263146be1edcd.NewAutofitColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeFormatRequestBuilder) AutofitRows()(*i59be4ff9e5cfa84f12c10e3f2b052ab0062298e7a5f88045e660ddf3fd33193b.AutofitRowsRequestBuilder) {
    return i59be4ff9e5cfa84f12c10e3f2b052ab0062298e7a5f88045e660ddf3fd33193b.NewAutofitRowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewWorkbookRangeFormatRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFormatRequestBuilder) {
    m := &WorkbookRangeFormatRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/shared/{sharedInsight_id}/resource/microsoft.graph.workbookRangeFormat";
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