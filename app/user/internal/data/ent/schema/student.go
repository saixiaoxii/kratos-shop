package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(50).Comment("名称"),
		field.Bool("sex").Comment("性别"),
		field.Int("age").Comment("年龄"),
		field.Int("class_id").Comment("班级ID"),
	}
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("class", Class.Type).
			Ref("student").
			Unique().
			Field("class_id").
			Required(),
	}
}

func (Student) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "student",
		},
	}
}
