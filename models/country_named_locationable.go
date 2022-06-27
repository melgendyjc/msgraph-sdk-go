package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CountryNamedLocationable 
type CountryNamedLocationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    NamedLocationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCountriesAndRegions()([]string)
    GetCountryLookupMethod()(*CountryLookupMethodType)
    GetIncludeUnknownCountriesAndRegions()(*bool)
    SetCountriesAndRegions(value []string)()
    SetCountryLookupMethod(value *CountryLookupMethodType)()
    SetIncludeUnknownCountriesAndRegions(value *bool)()
}
