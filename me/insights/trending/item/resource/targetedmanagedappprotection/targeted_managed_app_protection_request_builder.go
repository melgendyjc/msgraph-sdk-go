package targetedmanagedappprotection

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i3272a439ec24aa01bb0c5c5eab90b0412bea2231615fa2033f9ed7c454b9a111 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/targetedmanagedappprotection/assign"
)

type TargetedManagedAppProtectionRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *TargetedManagedAppProtectionRequestBuilder) Assign()(*i3272a439ec24aa01bb0c5c5eab90b0412bea2231615fa2033f9ed7c454b9a111.AssignRequestBuilder) {
    return i3272a439ec24aa01bb0c5c5eab90b0412bea2231615fa2033f9ed7c454b9a111.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewTargetedManagedAppProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppProtectionRequestBuilder) {
    m := &TargetedManagedAppProtectionRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/trending/{trending_id}/resource/microsoft.graph.targetedManagedAppProtection";
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
func NewTargetedManagedAppProtectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTargetedManagedAppProtectionRequestBuilderInternal(urlParams, requestAdapter)
}