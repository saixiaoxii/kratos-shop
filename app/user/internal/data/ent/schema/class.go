package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Class holds the schema definition for the Class entity.
type Class struct {
	ent.Schema
}

// Fields of the Class.
func (Class) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(50).Comment("名称"),
		field.Int("level").Comment("级别"),
	}
}

// Edges of the Class.
func (Class) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("student", Student.Type),
	}
}

func (Class) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "class"},
	}
}
