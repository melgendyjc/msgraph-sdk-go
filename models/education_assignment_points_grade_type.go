package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationAssignmentPointsGradeType 
type EducationAssignmentPointsGradeType struct {
    EducationAssignmentGradeType
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Max points possible for this assignment.
    maxPoints *float32
}
// NewEducationAssignmentPointsGradeType instantiates a new EducationAssignmentPointsGradeType and sets the default values.
func NewEducationAssignmentPointsGradeType()(*EducationAssignmentPointsGradeType) {
    m := &EducationAssignmentPointsGradeType{
        EducationAssignmentGradeType: *NewEducationAssignmentGradeType(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEducationAssignmentPointsGradeTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationAssignmentPointsGradeTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationAssignmentPointsGradeType(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationAssignmentPointsGradeType) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationAssignmentPointsGradeType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationAssignmentGradeType.GetFieldDeserializers()
    res["maxPoints"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxPoints(val)
        }
        return nil
    }
    return res
}
// GetMaxPoints gets the maxPoints property value. Max points possible for this assignment.
func (m *EducationAssignmentPointsGradeType) GetMaxPoints()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.maxPoints
    }
}
// Serialize serializes information the current object
func (m *EducationAssignmentPointsGradeType) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationAssignmentGradeType.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteFloat32Value("maxPoints", m.GetMaxPoints())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationAssignmentPointsGradeType) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMaxPoints sets the maxPoints property value. Max points possible for this assignment.
func (m *EducationAssignmentPointsGradeType) SetMaxPoints(value *float32)() {
    if m != nil {
        m.maxPoints = value
    }
}
