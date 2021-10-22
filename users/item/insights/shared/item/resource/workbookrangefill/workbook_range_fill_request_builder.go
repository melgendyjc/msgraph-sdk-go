package workbookrangefill

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i03a4e62fe0911d308f4bbc2f8d1f5d1e957b3f5632591c3b27e3052886940abc "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrangefill/clear"
)

type WorkbookRangeFillRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *WorkbookRangeFillRequestBuilder) Clear()(*i03a4e62fe0911d308f4bbc2f8d1f5d1e957b3f5632591c3b27e3052886940abc.ClearRequestBuilder) {
    return i03a4e62fe0911d308f4bbc2f8d1f5d1e957b3f5632591c3b27e3052886940abc.NewClearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewWorkbookRangeFillRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFillRequestBuilder) {
    m := &WorkbookRangeFillRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/insights/shared/{sharedInsight_id}/resource/microsoft.graph.workbookRangeFill";
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
func NewWorkbookRangeFillRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFillRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookRangeFillRequestBuilderInternal(urlParams, requestAdapter)
}