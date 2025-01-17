package models
import (
    "errors"
)
// Provides operations to manage the collection of application entities.
type WebsiteType int

const (
    OTHER_WEBSITETYPE WebsiteType = iota
    HOME_WEBSITETYPE
    WORK_WEBSITETYPE
    BLOG_WEBSITETYPE
    PROFILE_WEBSITETYPE
)

func (i WebsiteType) String() string {
    return []string{"other", "home", "work", "blog", "profile"}[i]
}
func ParseWebsiteType(v string) (interface{}, error) {
    result := OTHER_WEBSITETYPE
    switch v {
        case "other":
            result = OTHER_WEBSITETYPE
        case "home":
            result = HOME_WEBSITETYPE
        case "work":
            result = WORK_WEBSITETYPE
        case "blog":
            result = BLOG_WEBSITETYPE
        case "profile":
            result = PROFILE_WEBSITETYPE
        default:
            return 0, errors.New("Unknown WebsiteType value: " + v)
    }
    return &result, nil
}
func SerializeWebsiteType(values []WebsiteType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
