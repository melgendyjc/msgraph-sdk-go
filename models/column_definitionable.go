package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ColumnDefinitionable 
type ColumnDefinitionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBoolean()(BooleanColumnable)
    GetCalculated()(CalculatedColumnable)
    GetChoice()(ChoiceColumnable)
    GetColumnGroup()(*string)
    GetContentApprovalStatus()(ContentApprovalStatusColumnable)
    GetCurrency()(CurrencyColumnable)
    GetDateTime()(DateTimeColumnable)
    GetDefaultValue()(DefaultColumnValueable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetEnforceUniqueValues()(*bool)
    GetGeolocation()(GeolocationColumnable)
    GetHidden()(*bool)
    GetHyperlinkOrPicture()(HyperlinkOrPictureColumnable)
    GetIndexed()(*bool)
    GetIsDeletable()(*bool)
    GetIsReorderable()(*bool)
    GetIsSealed()(*bool)
    GetLookup()(LookupColumnable)
    GetName()(*string)
    GetNumber()(NumberColumnable)
    GetPersonOrGroup()(PersonOrGroupColumnable)
    GetPropagateChanges()(*bool)
    GetReadOnly()(*bool)
    GetRequired()(*bool)
    GetSourceColumn()(ColumnDefinitionable)
    GetSourceContentType()(ContentTypeInfoable)
    GetTerm()(TermColumnable)
    GetText()(TextColumnable)
    GetThumbnail()(ThumbnailColumnable)
    GetType()(*ColumnTypes)
    GetValidation()(ColumnValidationable)
    SetBoolean(value BooleanColumnable)()
    SetCalculated(value CalculatedColumnable)()
    SetChoice(value ChoiceColumnable)()
    SetColumnGroup(value *string)()
    SetContentApprovalStatus(value ContentApprovalStatusColumnable)()
    SetCurrency(value CurrencyColumnable)()
    SetDateTime(value DateTimeColumnable)()
    SetDefaultValue(value DefaultColumnValueable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetEnforceUniqueValues(value *bool)()
    SetGeolocation(value GeolocationColumnable)()
    SetHidden(value *bool)()
    SetHyperlinkOrPicture(value HyperlinkOrPictureColumnable)()
    SetIndexed(value *bool)()
    SetIsDeletable(value *bool)()
    SetIsReorderable(value *bool)()
    SetIsSealed(value *bool)()
    SetLookup(value LookupColumnable)()
    SetName(value *string)()
    SetNumber(value NumberColumnable)()
    SetPersonOrGroup(value PersonOrGroupColumnable)()
    SetPropagateChanges(value *bool)()
    SetReadOnly(value *bool)()
    SetRequired(value *bool)()
    SetSourceColumn(value ColumnDefinitionable)()
    SetSourceContentType(value ContentTypeInfoable)()
    SetTerm(value TermColumnable)()
    SetText(value TextColumnable)()
    SetThumbnail(value ThumbnailColumnable)()
    SetType(value *ColumnTypes)()
    SetValidation(value ColumnValidationable)()
}
