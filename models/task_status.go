package models
import (
    "errors"
)
// Provides operations to manage the collection of application entities.
type TaskStatus int

const (
    NOTSTARTED_TASKSTATUS TaskStatus = iota
    INPROGRESS_TASKSTATUS
    COMPLETED_TASKSTATUS
    WAITINGONOTHERS_TASKSTATUS
    DEFERRED_TASKSTATUS
)

func (i TaskStatus) String() string {
    return []string{"notStarted", "inProgress", "completed", "waitingOnOthers", "deferred"}[i]
}
func ParseTaskStatus(v string) (interface{}, error) {
    result := NOTSTARTED_TASKSTATUS
    switch v {
        case "notStarted":
            result = NOTSTARTED_TASKSTATUS
        case "inProgress":
            result = INPROGRESS_TASKSTATUS
        case "completed":
            result = COMPLETED_TASKSTATUS
        case "waitingOnOthers":
            result = WAITINGONOTHERS_TASKSTATUS
        case "deferred":
            result = DEFERRED_TASKSTATUS
        default:
            return 0, errors.New("Unknown TaskStatus value: " + v)
    }
    return &result, nil
}
func SerializeTaskStatus(values []TaskStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
