package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleManagementPolicyRuleTargetable 
type UnifiedRoleManagementPolicyRuleTargetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCaller()(*string)
    GetEnforcedSettings()([]string)
    GetInheritableSettings()([]string)
    GetLevel()(*string)
    GetOperations()([]string)
    GetTargetObjects()([]DirectoryObjectable)
    SetCaller(value *string)()
    SetEnforcedSettings(value []string)()
    SetInheritableSettings(value []string)()
    SetLevel(value *string)()
    SetOperations(value []string)()
    SetTargetObjects(value []DirectoryObjectable)()
}
