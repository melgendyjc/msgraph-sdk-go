package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerTaskable 
type PlannerTaskable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveChecklistItemCount()(*int32)
    GetAppliedCategories()(PlannerAppliedCategoriesable)
    GetAssignedToTaskBoardFormat()(PlannerAssignedToTaskBoardTaskFormatable)
    GetAssigneePriority()(*string)
    GetAssignments()(PlannerAssignmentsable)
    GetBucketId()(*string)
    GetBucketTaskBoardFormat()(PlannerBucketTaskBoardTaskFormatable)
    GetChecklistItemCount()(*int32)
    GetCompletedBy()(IdentitySetable)
    GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetConversationThreadId()(*string)
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDetails()(PlannerTaskDetailsable)
    GetDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHasDescription()(*bool)
    GetOrderHint()(*string)
    GetPercentComplete()(*int32)
    GetPlanId()(*string)
    GetPreviewType()(*PlannerPreviewType)
    GetPriority()(*int32)
    GetProgressTaskBoardFormat()(PlannerProgressTaskBoardTaskFormatable)
    GetReferenceCount()(*int32)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTitle()(*string)
    SetActiveChecklistItemCount(value *int32)()
    SetAppliedCategories(value PlannerAppliedCategoriesable)()
    SetAssignedToTaskBoardFormat(value PlannerAssignedToTaskBoardTaskFormatable)()
    SetAssigneePriority(value *string)()
    SetAssignments(value PlannerAssignmentsable)()
    SetBucketId(value *string)()
    SetBucketTaskBoardFormat(value PlannerBucketTaskBoardTaskFormatable)()
    SetChecklistItemCount(value *int32)()
    SetCompletedBy(value IdentitySetable)()
    SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetConversationThreadId(value *string)()
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDetails(value PlannerTaskDetailsable)()
    SetDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHasDescription(value *bool)()
    SetOrderHint(value *string)()
    SetPercentComplete(value *int32)()
    SetPlanId(value *string)()
    SetPreviewType(value *PlannerPreviewType)()
    SetPriority(value *int32)()
    SetProgressTaskBoardFormat(value PlannerProgressTaskBoardTaskFormatable)()
    SetReferenceCount(value *int32)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTitle(value *string)()
}
