package models
import (
    "errors"
)
// Provides operations to manage the collection of application entities.
type MeetingRequestType int

const (
    NONE_MEETINGREQUESTTYPE MeetingRequestType = iota
    NEWMEETINGREQUEST_MEETINGREQUESTTYPE
    FULLUPDATE_MEETINGREQUESTTYPE
    INFORMATIONALUPDATE_MEETINGREQUESTTYPE
    SILENTUPDATE_MEETINGREQUESTTYPE
    OUTDATED_MEETINGREQUESTTYPE
    PRINCIPALWANTSCOPY_MEETINGREQUESTTYPE
)

func (i MeetingRequestType) String() string {
    return []string{"none", "newMeetingRequest", "fullUpdate", "informationalUpdate", "silentUpdate", "outdated", "principalWantsCopy"}[i]
}
func ParseMeetingRequestType(v string) (interface{}, error) {
    result := NONE_MEETINGREQUESTTYPE
    switch v {
        case "none":
            result = NONE_MEETINGREQUESTTYPE
        case "newMeetingRequest":
            result = NEWMEETINGREQUEST_MEETINGREQUESTTYPE
        case "fullUpdate":
            result = FULLUPDATE_MEETINGREQUESTTYPE
        case "informationalUpdate":
            result = INFORMATIONALUPDATE_MEETINGREQUESTTYPE
        case "silentUpdate":
            result = SILENTUPDATE_MEETINGREQUESTTYPE
        case "outdated":
            result = OUTDATED_MEETINGREQUESTTYPE
        case "principalWantsCopy":
            result = PRINCIPALWANTSCOPY_MEETINGREQUESTTYPE
        default:
            return 0, errors.New("Unknown MeetingRequestType value: " + v)
    }
    return &result, nil
}
func SerializeMeetingRequestType(values []MeetingRequestType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
