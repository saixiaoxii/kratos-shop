// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ClassColumns holds the columns for the "class" table.
	ClassColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 50},
		{Name: "level", Type: field.TypeInt},
	}
	// ClassTable holds the schema information for the "class" table.
	ClassTable = &schema.Table{
		Name:       "class",
		Columns:    ClassColumns,
		PrimaryKey: []*schema.Column{ClassColumns[0]},
	}
	// StudentColumns holds the columns for the "student" table.
	StudentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 50},
		{Name: "sex", Type: field.TypeBool},
		{Name: "age", Type: field.TypeInt},
		{Name: "class_id", Type: field.TypeInt},
	}
	// StudentTable holds the schema information for the "student" table.
	StudentTable = &schema.Table{
		Name:       "student",
		Columns:    StudentColumns,
		PrimaryKey: []*schema.Column{StudentColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "student_class_student",
				Columns:    []*schema.Column{StudentColumns[4]},
				RefColumns: []*schema.Column{ClassColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ClassTable,
		StudentTable,
		UsersTable,
	}
)

func init() {
	ClassTable.Annotation = &entsql.Annotation{
		Table: "class",
	}
	StudentTable.ForeignKeys[0].RefTable = ClassTable
	StudentTable.Annotation = &entsql.Annotation{
		Table: "student",
	}
}
