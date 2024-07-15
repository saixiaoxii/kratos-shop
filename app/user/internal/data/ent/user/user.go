// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "ID"
	// FieldMobile holds the string denoting the mobile field in the database.
	FieldMobile = "Mobile"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "Password"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "NickName"
	// FieldBirthday holds the string denoting the birthday field in the database.
	FieldBirthday = "Birthday"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "Gender"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "Role"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "add_time"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "update_time"
	// FieldIsDeleted holds the string denoting the is_deleted field in the database.
	FieldIsDeleted = "IsDeletedAt"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldMobile,
	FieldPassword,
	FieldNickname,
	FieldBirthday,
	FieldGender,
	FieldRole,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldIsDeleted,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// MobileValidator is a validator for the "mobile" field. It is called by the builders before save.
	MobileValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	NicknameValidator func(string) error
	// DefaultGender holds the default value on creation for the "gender" field.
	DefaultGender string
	// GenderValidator is a validator for the "gender" field. It is called by the builders before save.
	GenderValidator func(string) error
	// DefaultRole holds the default value on creation for the "role" field.
	DefaultRole int
	// DefaultIsDeleted holds the default value on creation for the "is_deleted" field.
	DefaultIsDeleted bool
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByMobile orders the results by the mobile field.
func ByMobile(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMobile, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByNickname orders the results by the nickname field.
func ByNickname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNickname, opts...).ToFunc()
}

// ByBirthday orders the results by the birthday field.
func ByBirthday(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBirthday, opts...).ToFunc()
}

// ByGender orders the results by the gender field.
func ByGender(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGender, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByIsDeleted orders the results by the is_deleted field.
func ByIsDeleted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsDeleted, opts...).ToFunc()
}
