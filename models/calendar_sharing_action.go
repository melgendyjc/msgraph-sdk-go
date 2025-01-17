package models
import (
    "errors"
)
// Provides operations to manage the collection of application entities.
type CalendarSharingAction int

const (
    ACCEPT_CALENDARSHARINGACTION CalendarSharingAction = iota
    ACCEPTANDVIEWCALENDAR_CALENDARSHARINGACTION
    VIEWCALENDAR_CALENDARSHARINGACTION
    ADDTHISCALENDAR_CALENDARSHARINGACTION
)

func (i CalendarSharingAction) String() string {
    return []string{"accept", "acceptAndViewCalendar", "viewCalendar", "addThisCalendar"}[i]
}
func ParseCalendarSharingAction(v string) (interface{}, error) {
    result := ACCEPT_CALENDARSHARINGACTION
    switch v {
        case "accept":
            result = ACCEPT_CALENDARSHARINGACTION
        case "acceptAndViewCalendar":
            result = ACCEPTANDVIEWCALENDAR_CALENDARSHARINGACTION
        case "viewCalendar":
            result = VIEWCALENDAR_CALENDARSHARINGACTION
        case "addThisCalendar":
            result = ADDTHISCALENDAR_CALENDARSHARINGACTION
        default:
            return 0, errors.New("Unknown CalendarSharingAction value: " + v)
    }
    return &result, nil
}
func SerializeCalendarSharingAction(values []CalendarSharingAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
