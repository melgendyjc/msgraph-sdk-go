package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConversationThreadable 
type ConversationThreadable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCcRecipients()([]Recipientable)
    GetHasAttachments()(*bool)
    GetIsLocked()(*bool)
    GetLastDeliveredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPosts()([]Postable)
    GetPreview()(*string)
    GetTopic()(*string)
    GetToRecipients()([]Recipientable)
    GetUniqueSenders()([]string)
    SetCcRecipients(value []Recipientable)()
    SetHasAttachments(value *bool)()
    SetIsLocked(value *bool)()
    SetLastDeliveredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPosts(value []Postable)()
    SetPreview(value *string)()
    SetTopic(value *string)()
    SetToRecipients(value []Recipientable)()
    SetUniqueSenders(value []string)()
}
