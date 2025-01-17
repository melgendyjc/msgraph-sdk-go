package models
import (
    "errors"
)
// Provides operations to manage the identityGovernance singleton.
type AccessReviewExpirationBehavior int

const (
    KEEPACCESS_ACCESSREVIEWEXPIRATIONBEHAVIOR AccessReviewExpirationBehavior = iota
    REMOVEACCESS_ACCESSREVIEWEXPIRATIONBEHAVIOR
    ACCEPTACCESSRECOMMENDATION_ACCESSREVIEWEXPIRATIONBEHAVIOR
    UNKNOWNFUTUREVALUE_ACCESSREVIEWEXPIRATIONBEHAVIOR
)

func (i AccessReviewExpirationBehavior) String() string {
    return []string{"keepAccess", "removeAccess", "acceptAccessRecommendation", "unknownFutureValue"}[i]
}
func ParseAccessReviewExpirationBehavior(v string) (interface{}, error) {
    result := KEEPACCESS_ACCESSREVIEWEXPIRATIONBEHAVIOR
    switch v {
        case "keepAccess":
            result = KEEPACCESS_ACCESSREVIEWEXPIRATIONBEHAVIOR
        case "removeAccess":
            result = REMOVEACCESS_ACCESSREVIEWEXPIRATIONBEHAVIOR
        case "acceptAccessRecommendation":
            result = ACCEPTACCESSRECOMMENDATION_ACCESSREVIEWEXPIRATIONBEHAVIOR
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ACCESSREVIEWEXPIRATIONBEHAVIOR
        default:
            return 0, errors.New("Unknown AccessReviewExpirationBehavior value: " + v)
    }
    return &result, nil
}
func SerializeAccessReviewExpirationBehavior(values []AccessReviewExpirationBehavior) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
