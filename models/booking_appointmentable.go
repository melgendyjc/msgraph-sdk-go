package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingAppointmentable 
type BookingAppointmentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalInformation()(*string)
    GetCustomers()([]BookingCustomerInformationBaseable)
    GetCustomerTimeZone()(*string)
    GetDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetEndDateTime()(DateTimeTimeZoneable)
    GetFilledAttendeesCount()(*int32)
    GetIsLocationOnline()(*bool)
    GetJoinWebUrl()(*string)
    GetMaximumAttendeesCount()(*int32)
    GetOptOutOfCustomerEmail()(*bool)
    GetPostBuffer()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetPreBuffer()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetPrice()(*float64)
    GetPriceType()(*BookingPriceType)
    GetReminders()([]BookingReminderable)
    GetSelfServiceAppointmentId()(*string)
    GetServiceId()(*string)
    GetServiceLocation()(Locationable)
    GetServiceName()(*string)
    GetServiceNotes()(*string)
    GetSmsNotificationsEnabled()(*bool)
    GetStaffMemberIds()([]string)
    GetStartDateTime()(DateTimeTimeZoneable)
    SetAdditionalInformation(value *string)()
    SetCustomers(value []BookingCustomerInformationBaseable)()
    SetCustomerTimeZone(value *string)()
    SetDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetEndDateTime(value DateTimeTimeZoneable)()
    SetFilledAttendeesCount(value *int32)()
    SetIsLocationOnline(value *bool)()
    SetJoinWebUrl(value *string)()
    SetMaximumAttendeesCount(value *int32)()
    SetOptOutOfCustomerEmail(value *bool)()
    SetPostBuffer(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetPreBuffer(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetPrice(value *float64)()
    SetPriceType(value *BookingPriceType)()
    SetReminders(value []BookingReminderable)()
    SetSelfServiceAppointmentId(value *string)()
    SetServiceId(value *string)()
    SetServiceLocation(value Locationable)()
    SetServiceName(value *string)()
    SetServiceNotes(value *string)()
    SetSmsNotificationsEnabled(value *bool)()
    SetStaffMemberIds(value []string)()
    SetStartDateTime(value DateTimeTimeZoneable)()
}
