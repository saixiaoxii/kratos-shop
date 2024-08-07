// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-shop/app/user/internal/data/ent/predicate"
	"kratos-shop/app/user/internal/data/ent/user"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUser = "User"
)

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int64
	mobile        *string
	password      *string
	nickname      *string
	birthday      *time.Time
	gender        *string
	role          *int
	addrole       *int
	created_at    *time.Time
	updated_at    *time.Time
	is_deleted    *bool
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int64) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of User entities.
func (m *UserMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetMobile sets the "mobile" field.
func (m *UserMutation) SetMobile(s string) {
	m.mobile = &s
}

// Mobile returns the value of the "mobile" field in the mutation.
func (m *UserMutation) Mobile() (r string, exists bool) {
	v := m.mobile
	if v == nil {
		return
	}
	return *v, true
}

// OldMobile returns the old "mobile" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldMobile(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMobile is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMobile requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMobile: %w", err)
	}
	return oldValue.Mobile, nil
}

// ResetMobile resets all changes to the "mobile" field.
func (m *UserMutation) ResetMobile() {
	m.mobile = nil
}

// SetPassword sets the "password" field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
}

// SetNickname sets the "nickname" field.
func (m *UserMutation) SetNickname(s string) {
	m.nickname = &s
}

// Nickname returns the value of the "nickname" field in the mutation.
func (m *UserMutation) Nickname() (r string, exists bool) {
	v := m.nickname
	if v == nil {
		return
	}
	return *v, true
}

// OldNickname returns the old "nickname" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldNickname(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldNickname is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldNickname requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNickname: %w", err)
	}
	return oldValue.Nickname, nil
}

// ResetNickname resets all changes to the "nickname" field.
func (m *UserMutation) ResetNickname() {
	m.nickname = nil
}

// SetBirthday sets the "birthday" field.
func (m *UserMutation) SetBirthday(t time.Time) {
	m.birthday = &t
}

// Birthday returns the value of the "birthday" field in the mutation.
func (m *UserMutation) Birthday() (r time.Time, exists bool) {
	v := m.birthday
	if v == nil {
		return
	}
	return *v, true
}

// OldBirthday returns the old "birthday" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldBirthday(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBirthday is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBirthday requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBirthday: %w", err)
	}
	return oldValue.Birthday, nil
}

// ClearBirthday clears the value of the "birthday" field.
func (m *UserMutation) ClearBirthday() {
	m.birthday = nil
	m.clearedFields[user.FieldBirthday] = struct{}{}
}

// BirthdayCleared returns if the "birthday" field was cleared in this mutation.
func (m *UserMutation) BirthdayCleared() bool {
	_, ok := m.clearedFields[user.FieldBirthday]
	return ok
}

// ResetBirthday resets all changes to the "birthday" field.
func (m *UserMutation) ResetBirthday() {
	m.birthday = nil
	delete(m.clearedFields, user.FieldBirthday)
}

// SetGender sets the "gender" field.
func (m *UserMutation) SetGender(s string) {
	m.gender = &s
}

// Gender returns the value of the "gender" field in the mutation.
func (m *UserMutation) Gender() (r string, exists bool) {
	v := m.gender
	if v == nil {
		return
	}
	return *v, true
}

// OldGender returns the old "gender" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldGender(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGender is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGender requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGender: %w", err)
	}
	return oldValue.Gender, nil
}

// ResetGender resets all changes to the "gender" field.
func (m *UserMutation) ResetGender() {
	m.gender = nil
}

// SetRole sets the "role" field.
func (m *UserMutation) SetRole(i int) {
	m.role = &i
	m.addrole = nil
}

// Role returns the value of the "role" field in the mutation.
func (m *UserMutation) Role() (r int, exists bool) {
	v := m.role
	if v == nil {
		return
	}
	return *v, true
}

// OldRole returns the old "role" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldRole(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRole is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRole requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRole: %w", err)
	}
	return oldValue.Role, nil
}

// AddRole adds i to the "role" field.
func (m *UserMutation) AddRole(i int) {
	if m.addrole != nil {
		*m.addrole += i
	} else {
		m.addrole = &i
	}
}

// AddedRole returns the value that was added to the "role" field in this mutation.
func (m *UserMutation) AddedRole() (r int, exists bool) {
	v := m.addrole
	if v == nil {
		return
	}
	return *v, true
}

// ResetRole resets all changes to the "role" field.
func (m *UserMutation) ResetRole() {
	m.role = nil
	m.addrole = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ClearCreatedAt clears the value of the "created_at" field.
func (m *UserMutation) ClearCreatedAt() {
	m.created_at = nil
	m.clearedFields[user.FieldCreatedAt] = struct{}{}
}

// CreatedAtCleared returns if the "created_at" field was cleared in this mutation.
func (m *UserMutation) CreatedAtCleared() bool {
	_, ok := m.clearedFields[user.FieldCreatedAt]
	return ok
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *UserMutation) ResetCreatedAt() {
	m.created_at = nil
	delete(m.clearedFields, user.FieldCreatedAt)
}

// SetUpdatedAt sets the "updated_at" field.
func (m *UserMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *UserMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (m *UserMutation) ClearUpdatedAt() {
	m.updated_at = nil
	m.clearedFields[user.FieldUpdatedAt] = struct{}{}
}

// UpdatedAtCleared returns if the "updated_at" field was cleared in this mutation.
func (m *UserMutation) UpdatedAtCleared() bool {
	_, ok := m.clearedFields[user.FieldUpdatedAt]
	return ok
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *UserMutation) ResetUpdatedAt() {
	m.updated_at = nil
	delete(m.clearedFields, user.FieldUpdatedAt)
}

// SetIsDeleted sets the "is_deleted" field.
func (m *UserMutation) SetIsDeleted(b bool) {
	m.is_deleted = &b
}

// IsDeleted returns the value of the "is_deleted" field in the mutation.
func (m *UserMutation) IsDeleted() (r bool, exists bool) {
	v := m.is_deleted
	if v == nil {
		return
	}
	return *v, true
}

// OldIsDeleted returns the old "is_deleted" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldIsDeleted(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIsDeleted is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIsDeleted requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsDeleted: %w", err)
	}
	return oldValue.IsDeleted, nil
}

// ResetIsDeleted resets all changes to the "is_deleted" field.
func (m *UserMutation) ResetIsDeleted() {
	m.is_deleted = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 9)
	if m.mobile != nil {
		fields = append(fields, user.FieldMobile)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	if m.nickname != nil {
		fields = append(fields, user.FieldNickname)
	}
	if m.birthday != nil {
		fields = append(fields, user.FieldBirthday)
	}
	if m.gender != nil {
		fields = append(fields, user.FieldGender)
	}
	if m.role != nil {
		fields = append(fields, user.FieldRole)
	}
	if m.created_at != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, user.FieldUpdatedAt)
	}
	if m.is_deleted != nil {
		fields = append(fields, user.FieldIsDeleted)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldMobile:
		return m.Mobile()
	case user.FieldPassword:
		return m.Password()
	case user.FieldNickname:
		return m.Nickname()
	case user.FieldBirthday:
		return m.Birthday()
	case user.FieldGender:
		return m.Gender()
	case user.FieldRole:
		return m.Role()
	case user.FieldCreatedAt:
		return m.CreatedAt()
	case user.FieldUpdatedAt:
		return m.UpdatedAt()
	case user.FieldIsDeleted:
		return m.IsDeleted()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldMobile:
		return m.OldMobile(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	case user.FieldNickname:
		return m.OldNickname(ctx)
	case user.FieldBirthday:
		return m.OldBirthday(ctx)
	case user.FieldGender:
		return m.OldGender(ctx)
	case user.FieldRole:
		return m.OldRole(ctx)
	case user.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case user.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case user.FieldIsDeleted:
		return m.OldIsDeleted(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldMobile:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMobile(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case user.FieldNickname:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNickname(v)
		return nil
	case user.FieldBirthday:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBirthday(v)
		return nil
	case user.FieldGender:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGender(v)
		return nil
	case user.FieldRole:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRole(v)
		return nil
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case user.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case user.FieldIsDeleted:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsDeleted(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.addrole != nil {
		fields = append(fields, user.FieldRole)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case user.FieldRole:
		return m.AddedRole()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case user.FieldRole:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddRole(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldBirthday) {
		fields = append(fields, user.FieldBirthday)
	}
	if m.FieldCleared(user.FieldCreatedAt) {
		fields = append(fields, user.FieldCreatedAt)
	}
	if m.FieldCleared(user.FieldUpdatedAt) {
		fields = append(fields, user.FieldUpdatedAt)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldBirthday:
		m.ClearBirthday()
		return nil
	case user.FieldCreatedAt:
		m.ClearCreatedAt()
		return nil
	case user.FieldUpdatedAt:
		m.ClearUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldMobile:
		m.ResetMobile()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	case user.FieldNickname:
		m.ResetNickname()
		return nil
	case user.FieldBirthday:
		m.ResetBirthday()
		return nil
	case user.FieldGender:
		m.ResetGender()
		return nil
	case user.FieldRole:
		m.ResetRole()
		return nil
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case user.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case user.FieldIsDeleted:
		m.ResetIsDeleted()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
