package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i04a148e32be31a86cd21b897c8b55a4508d63dbae47a02eaeb122662dbd2ff9b "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/template"
    i5cc953f63726e53531e5d00c609bd14469e1ce23944c22b6aea0986c18fb0043 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/primarychannel"
    i8a1cdbeac728d5d9d3409d0d7085c53384ad37435e0292d966ed94bbc4155a05 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/group"
    i9fa0e9d329dc2b42ce0cc0330991bb8f8e864efaaef5061789d895e28321a6b2 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/operations"
    icf925d6e8373dd15bb408b246595c1d2598e1881b555a62a02a76300fc5b7cd2 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/installedapps"
    id463d65124ba412b3980ec713bebe8eb90e4a925515f5f993cd79d5b01b70907 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/channels"
    ifc923048eff969ef232f17cdaf6c11e18676c5c9e2918bc19001d30cbdc4b5c1 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/members"
    ifd53534f50d40567e607c2213e794582e29aa46a0c07e2d406db231a42a0140a "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/schedule"
    i03b2a35397701f29dd92bcf0dd549e252ab75a1bd539b1b1ca11a58738ba2c51 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/members/item"
    i77d75bcb3611254a7012d361f29a795e840693affe1de5861c58d310feff299c "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/operations/item"
    id12792e5e45c250abd3ecec9bbc023cd6dc63edea08efc7907dcad0da46a8a0d "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/channels/item"
    ifaff06990f5e17731e64f75f4cb575f1ada76645d84341378dccf3aef4f66fab "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/installedapps/item"
)

// TeamItemRequestBuilder provides operations to manage the joinedTeams property of the microsoft.graph.user entity.
type TeamItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamItemRequestBuilderGetQueryParameters the Microsoft Teams teams that the user is a member of. Read-only. Nullable.
type TeamItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamItemRequestBuilderGetQueryParameters
}
// TeamItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Channels the channels property
func (m *TeamItemRequestBuilder) Channels()(*id463d65124ba412b3980ec713bebe8eb90e4a925515f5f993cd79d5b01b70907.ChannelsRequestBuilder) {
    return id463d65124ba412b3980ec713bebe8eb90e4a925515f5f993cd79d5b01b70907.NewChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChannelsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.channels.item collection
func (m *TeamItemRequestBuilder) ChannelsById(id string)(*id12792e5e45c250abd3ecec9bbc023cd6dc63edea08efc7907dcad0da46a8a0d.ChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return id12792e5e45c250abd3ecec9bbc023cd6dc63edea08efc7907dcad0da46a8a0d.NewChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTeamItemRequestBuilderInternal instantiates a new TeamItemRequestBuilder and sets the default values.
func NewTeamItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamItemRequestBuilder) {
    m := &TeamItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedTeams/{team%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamItemRequestBuilder instantiates a new TeamItemRequestBuilder and sets the default values.
func NewTeamItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property joinedTeams for me
func (m *TeamItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property joinedTeams for me
func (m *TeamItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *TeamItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformationWithRequestConfiguration the Microsoft Teams teams that the user is a member of. Read-only. Nullable.
func (m *TeamItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the Microsoft Teams teams that the user is a member of. Read-only. Nullable.
func (m *TeamItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *TeamItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property joinedTeams in me
func (m *TeamItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property joinedTeams in me
func (m *TeamItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, requestConfiguration *TeamItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DeleteWithResponseHandler delete navigation property joinedTeams for me
func (m *TeamItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *TeamItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property joinedTeams for me
func (m *TeamItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *TeamItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// GetWithResponseHandler the Microsoft Teams teams that the user is a member of. Read-only. Nullable.
func (m *TeamItemRequestBuilder) GetWithResponseHandler(requestConfiguration *TeamItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler the Microsoft Teams teams that the user is a member of. Read-only. Nullable.
func (m *TeamItemRequestBuilder) GetWithResponseHandler(requestConfiguration *TeamItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable), nil
}
// Group the group property
func (m *TeamItemRequestBuilder) Group()(*i8a1cdbeac728d5d9d3409d0d7085c53384ad37435e0292d966ed94bbc4155a05.GroupRequestBuilder) {
    return i8a1cdbeac728d5d9d3409d0d7085c53384ad37435e0292d966ed94bbc4155a05.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledApps the installedApps property
func (m *TeamItemRequestBuilder) InstalledApps()(*icf925d6e8373dd15bb408b246595c1d2598e1881b555a62a02a76300fc5b7cd2.InstalledAppsRequestBuilder) {
    return icf925d6e8373dd15bb408b246595c1d2598e1881b555a62a02a76300fc5b7cd2.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.installedApps.item collection
func (m *TeamItemRequestBuilder) InstalledAppsById(id string)(*ifaff06990f5e17731e64f75f4cb575f1ada76645d84341378dccf3aef4f66fab.TeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return ifaff06990f5e17731e64f75f4cb575f1ada76645d84341378dccf3aef4f66fab.NewTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Members the members property
func (m *TeamItemRequestBuilder) Members()(*ifc923048eff969ef232f17cdaf6c11e18676c5c9e2918bc19001d30cbdc4b5c1.MembersRequestBuilder) {
    return ifc923048eff969ef232f17cdaf6c11e18676c5c9e2918bc19001d30cbdc4b5c1.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.members.item collection
func (m *TeamItemRequestBuilder) MembersById(id string)(*i03b2a35397701f29dd92bcf0dd549e252ab75a1bd539b1b1ca11a58738ba2c51.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return i03b2a35397701f29dd92bcf0dd549e252ab75a1bd539b1b1ca11a58738ba2c51.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *TeamItemRequestBuilder) Operations()(*i9fa0e9d329dc2b42ce0cc0330991bb8f8e864efaaef5061789d895e28321a6b2.OperationsRequestBuilder) {
    return i9fa0e9d329dc2b42ce0cc0330991bb8f8e864efaaef5061789d895e28321a6b2.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.operations.item collection
func (m *TeamItemRequestBuilder) OperationsById(id string)(*i77d75bcb3611254a7012d361f29a795e840693affe1de5861c58d310feff299c.TeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation%2Did"] = id
    }
    return i77d75bcb3611254a7012d361f29a795e840693affe1de5861c58d310feff299c.NewTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property joinedTeams in me
func (m *TeamItemRequestBuilder) PatchWithResponseHandler(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, requestConfiguration *TeamItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property joinedTeams in me
func (m *TeamItemRequestBuilder) PatchWithResponseHandler(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, requestConfiguration *TeamItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// PrimaryChannel the primaryChannel property
func (m *TeamItemRequestBuilder) PrimaryChannel()(*i5cc953f63726e53531e5d00c609bd14469e1ce23944c22b6aea0986c18fb0043.PrimaryChannelRequestBuilder) {
    return i5cc953f63726e53531e5d00c609bd14469e1ce23944c22b6aea0986c18fb0043.NewPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schedule the schedule property
func (m *TeamItemRequestBuilder) Schedule()(*ifd53534f50d40567e607c2213e794582e29aa46a0c07e2d406db231a42a0140a.ScheduleRequestBuilder) {
    return ifd53534f50d40567e607c2213e794582e29aa46a0c07e2d406db231a42a0140a.NewScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Template the template property
func (m *TeamItemRequestBuilder) Template()(*i04a148e32be31a86cd21b897c8b55a4508d63dbae47a02eaeb122662dbd2ff9b.TemplateRequestBuilder) {
    return i04a148e32be31a86cd21b897c8b55a4508d63dbae47a02eaeb122662dbd2ff9b.NewTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
