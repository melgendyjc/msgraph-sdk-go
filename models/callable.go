package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Callable 
type Callable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAudioRoutingGroups()([]AudioRoutingGroupable)
    GetCallbackUri()(*string)
    GetCallChainId()(*string)
    GetCallOptions()(CallOptionsable)
    GetCallRoutes()([]CallRouteable)
    GetChatInfo()(ChatInfoable)
    GetDirection()(*CallDirection)
    GetIncomingContext()(IncomingContextable)
    GetMediaConfig()(MediaConfigable)
    GetMediaState()(CallMediaStateable)
    GetMeetingInfo()(MeetingInfoable)
    GetMyParticipantId()(*string)
    GetOperations()([]CommsOperationable)
    GetParticipants()([]Participantable)
    GetRequestedModalities()([]string)
    GetResultInfo()(ResultInfoable)
    GetSource()(ParticipantInfoable)
    GetState()(*CallState)
    GetSubject()(*string)
    GetTargets()([]InvitationParticipantInfoable)
    GetTenantId()(*string)
    GetToneInfo()(ToneInfoable)
    GetTranscription()(CallTranscriptionInfoable)
    SetAudioRoutingGroups(value []AudioRoutingGroupable)()
    SetCallbackUri(value *string)()
    SetCallChainId(value *string)()
    SetCallOptions(value CallOptionsable)()
    SetCallRoutes(value []CallRouteable)()
    SetChatInfo(value ChatInfoable)()
    SetDirection(value *CallDirection)()
    SetIncomingContext(value IncomingContextable)()
    SetMediaConfig(value MediaConfigable)()
    SetMediaState(value CallMediaStateable)()
    SetMeetingInfo(value MeetingInfoable)()
    SetMyParticipantId(value *string)()
    SetOperations(value []CommsOperationable)()
    SetParticipants(value []Participantable)()
    SetRequestedModalities(value []string)()
    SetResultInfo(value ResultInfoable)()
    SetSource(value ParticipantInfoable)()
    SetState(value *CallState)()
    SetSubject(value *string)()
    SetTargets(value []InvitationParticipantInfoable)()
    SetTenantId(value *string)()
    SetToneInfo(value ToneInfoable)()
    SetTranscription(value CallTranscriptionInfoable)()
}
