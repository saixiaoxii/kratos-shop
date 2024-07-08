// +build tools
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"user/internal/data/ent/class"
	"user/internal/data/ent/schema"
	"user/internal/data/ent/student"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	classFields := schema.Class{}.Fields()
	_ = classFields
	// classDescName is the schema descriptor for name field.
	classDescName := classFields[0].Descriptor()
	// class.NameValidator is a validator for the "name" field. It is called by the builders before save.
	class.NameValidator = classDescName.Validators[0].(func(string) error)
	studentFields := schema.Student{}.Fields()
	_ = studentFields
	// studentDescName is the schema descriptor for name field.
	studentDescName := studentFields[0].Descriptor()
	// student.NameValidator is a validator for the "name" field. It is called by the builders before save.
	student.NameValidator = studentDescName.Validators[0].(func(string) error)
}