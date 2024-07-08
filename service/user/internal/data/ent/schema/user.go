package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").StorageKey("ID").Unique(),
		field.String("mobile").StorageKey("Mobile").Unique().NotEmpty().MaxLen(11).Comment("手机号码，用户唯一标识"),
		field.String("password").StorageKey("Password").NotEmpty().MaxLen(100).Comment("用户密码的保存需要注意是否加密"),
		field.String("nickname").StorageKey("NickName").MaxLen(25).Comment("用户昵称"),
		field.Time("birthday").StorageKey("Birthday").Nillable().Optional().Comment("出生日期"),
		field.String("gender").StorageKey("Gender").Default("male").MaxLen(16).Comment("female:女,male:男"),
		field.Int("role").StorageKey("Role").Default(1).Comment("1:普通用户,2:管理员"),
		field.Time("created_at").StorageKey("add_time").Optional(),
		field.Time("updated_at").StorageKey("update_time").Optional(),
		field.Bool("is_deleted").StorageKey("IsDeletedAt").Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("mobile").Unique(),
	}
}
