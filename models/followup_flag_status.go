package models
import (
    "errors"
)
// Provides operations to manage the collection of application entities.
type FollowupFlagStatus int

const (
    NOTFLAGGED_FOLLOWUPFLAGSTATUS FollowupFlagStatus = iota
    COMPLETE_FOLLOWUPFLAGSTATUS
    FLAGGED_FOLLOWUPFLAGSTATUS
)

func (i FollowupFlagStatus) String() string {
    return []string{"notFlagged", "complete", "flagged"}[i]
}
func ParseFollowupFlagStatus(v string) (interface{}, error) {
    result := NOTFLAGGED_FOLLOWUPFLAGSTATUS
    switch v {
        case "notFlagged":
            result = NOTFLAGGED_FOLLOWUPFLAGSTATUS
        case "complete":
            result = COMPLETE_FOLLOWUPFLAGSTATUS
        case "flagged":
            result = FLAGGED_FOLLOWUPFLAGSTATUS
        default:
            return 0, errors.New("Unknown FollowupFlagStatus value: " + v)
    }
    return &result, nil
}
func SerializeFollowupFlagStatus(values []FollowupFlagStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
