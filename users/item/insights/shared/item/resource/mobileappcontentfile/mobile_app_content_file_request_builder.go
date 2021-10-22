package mobileappcontentfile

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i13081911cac76e5ad9bf9ad5eef8e74bf6a8ef099b7a1eb663beaadcee7a222f "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/mobileappcontentfile/commit"
    i1498cecb70176119b7d7e17d3fabf65d01f3bd3b91651e805cf5918ccd6ccb9d "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/mobileappcontentfile/renewupload"
)

type MobileAppContentFileRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *MobileAppContentFileRequestBuilder) Commit()(*i13081911cac76e5ad9bf9ad5eef8e74bf6a8ef099b7a1eb663beaadcee7a222f.CommitRequestBuilder) {
    return i13081911cac76e5ad9bf9ad5eef8e74bf6a8ef099b7a1eb663beaadcee7a222f.NewCommitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewMobileAppContentFileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppContentFileRequestBuilder) {
    m := &MobileAppContentFileRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/insights/shared/{sharedInsight_id}/resource/microsoft.graph.mobileAppContentFile";
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
func NewMobileAppContentFileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppContentFileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppContentFileRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MobileAppContentFileRequestBuilder) RenewUpload()(*i1498cecb70176119b7d7e17d3fabf65d01f3bd3b91651e805cf5918ccd6ccb9d.RenewUploadRequestBuilder) {
    return i1498cecb70176119b7d7e17d3fabf65d01f3bd3b91651e805cf5918ccd6ccb9d.NewRenewUploadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}