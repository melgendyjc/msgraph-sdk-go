package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MediaContentRatingGermany 
type MediaContentRatingGermany struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Movies rating selected for Germany. Possible values are: allAllowed, allBlocked, general, agesAbove6, agesAbove12, agesAbove16, adults.
    movieRating *RatingGermanyMoviesType
    // TV rating selected for Germany. Possible values are: allAllowed, allBlocked, general, agesAbove6, agesAbove12, agesAbove16, adults.
    tvRating *RatingGermanyTelevisionType
}
// NewMediaContentRatingGermany instantiates a new mediaContentRatingGermany and sets the default values.
func NewMediaContentRatingGermany()(*MediaContentRatingGermany) {
    m := &MediaContentRatingGermany{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMediaContentRatingGermanyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMediaContentRatingGermanyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMediaContentRatingGermany(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MediaContentRatingGermany) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MediaContentRatingGermany) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["movieRating"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRatingGermanyMoviesType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMovieRating(val.(*RatingGermanyMoviesType))
        }
        return nil
    }
    res["tvRating"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRatingGermanyTelevisionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTvRating(val.(*RatingGermanyTelevisionType))
        }
        return nil
    }
    return res
}
// GetMovieRating gets the movieRating property value. Movies rating selected for Germany. Possible values are: allAllowed, allBlocked, general, agesAbove6, agesAbove12, agesAbove16, adults.
func (m *MediaContentRatingGermany) GetMovieRating()(*RatingGermanyMoviesType) {
    if m == nil {
        return nil
    } else {
        return m.movieRating
    }
}
// GetTvRating gets the tvRating property value. TV rating selected for Germany. Possible values are: allAllowed, allBlocked, general, agesAbove6, agesAbove12, agesAbove16, adults.
func (m *MediaContentRatingGermany) GetTvRating()(*RatingGermanyTelevisionType) {
    if m == nil {
        return nil
    } else {
        return m.tvRating
    }
}
// Serialize serializes information the current object
func (m *MediaContentRatingGermany) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetMovieRating() != nil {
        cast := (*m.GetMovieRating()).String()
        err := writer.WriteStringValue("movieRating", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTvRating() != nil {
        cast := (*m.GetTvRating()).String()
        err := writer.WriteStringValue("tvRating", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MediaContentRatingGermany) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMovieRating sets the movieRating property value. Movies rating selected for Germany. Possible values are: allAllowed, allBlocked, general, agesAbove6, agesAbove12, agesAbove16, adults.
func (m *MediaContentRatingGermany) SetMovieRating(value *RatingGermanyMoviesType)() {
    if m != nil {
        m.movieRating = value
    }
}
// SetTvRating sets the tvRating property value. TV rating selected for Germany. Possible values are: allAllowed, allBlocked, general, agesAbove6, agesAbove12, agesAbove16, adults.
func (m *MediaContentRatingGermany) SetTvRating(value *RatingGermanyTelevisionType)() {
    if m != nil {
        m.tvRating = value
    }
}