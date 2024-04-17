// Code generated by ent, DO NOT EDIT.

package ent

import (
	"AUTH/ent/login"
	"AUTH/ent/predicate"
	"AUTH/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableName(s *string) *UserUpdate {
	if s != nil {
		uu.SetName(*s)
	}
	return uu
}

// SetIP sets the "ip" field.
func (uu *UserUpdate) SetIP(s string) *UserUpdate {
	uu.mutation.SetIP(s)
	return uu
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIP(s *string) *UserUpdate {
	if s != nil {
		uu.SetIP(*s)
	}
	return uu
}

// SetDevice sets the "device" field.
func (uu *UserUpdate) SetDevice(s string) *UserUpdate {
	uu.mutation.SetDevice(s)
	return uu
}

// SetNillableDevice sets the "device" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDevice(s *string) *UserUpdate {
	if s != nil {
		uu.SetDevice(*s)
	}
	return uu
}

// SetVerified sets the "verified" field.
func (uu *UserUpdate) SetVerified(b bool) *UserUpdate {
	uu.mutation.SetVerified(b)
	return uu
}

// SetNillableVerified sets the "verified" field if the given value is not nil.
func (uu *UserUpdate) SetNillableVerified(b *bool) *UserUpdate {
	if b != nil {
		uu.SetVerified(*b)
	}
	return uu
}

// SetBlocked sets the "blocked" field.
func (uu *UserUpdate) SetBlocked(b bool) *UserUpdate {
	uu.mutation.SetBlocked(b)
	return uu
}

// SetNillableBlocked sets the "blocked" field if the given value is not nil.
func (uu *UserUpdate) SetNillableBlocked(b *bool) *UserUpdate {
	if b != nil {
		uu.SetBlocked(*b)
	}
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUsername(s *string) *UserUpdate {
	if s != nil {
		uu.SetUsername(*s)
	}
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UserUpdate) SetCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUpdatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetUpdatedAt(*t)
	}
	return uu
}

// AddLoginIDs adds the "logins" edge to the Login entity by IDs.
func (uu *UserUpdate) AddLoginIDs(ids ...int) *UserUpdate {
	uu.mutation.AddLoginIDs(ids...)
	return uu
}

// AddLogins adds the "logins" edges to the Login entity.
func (uu *UserUpdate) AddLogins(l ...*Login) *UserUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uu.AddLoginIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearLogins clears all "logins" edges to the Login entity.
func (uu *UserUpdate) ClearLogins() *UserUpdate {
	uu.mutation.ClearLogins()
	return uu
}

// RemoveLoginIDs removes the "logins" edge to Login entities by IDs.
func (uu *UserUpdate) RemoveLoginIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveLoginIDs(ids...)
	return uu
}

// RemoveLogins removes "logins" edges to Login entities.
func (uu *UserUpdate) RemoveLogins(l ...*Login) *UserUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uu.RemoveLoginIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.IP(); ok {
		_spec.SetField(user.FieldIP, field.TypeString, value)
	}
	if value, ok := uu.mutation.Device(); ok {
		_spec.SetField(user.FieldDevice, field.TypeString, value)
	}
	if value, ok := uu.mutation.Verified(); ok {
		_spec.SetField(user.FieldVerified, field.TypeBool, value)
	}
	if value, ok := uu.mutation.Blocked(); ok {
		_spec.SetField(user.FieldBlocked, field.TypeBool, value)
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if uu.mutation.LoginsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.LoginsTable,
			Columns: []string{user.LoginsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(login.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedLoginsIDs(); len(nodes) > 0 && !uu.mutation.LoginsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.LoginsTable,
			Columns: []string{user.LoginsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(login.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.LoginsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.LoginsTable,
			Columns: []string{user.LoginsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(login.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetName(*s)
	}
	return uuo
}

// SetIP sets the "ip" field.
func (uuo *UserUpdateOne) SetIP(s string) *UserUpdateOne {
	uuo.mutation.SetIP(s)
	return uuo
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIP(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetIP(*s)
	}
	return uuo
}

// SetDevice sets the "device" field.
func (uuo *UserUpdateOne) SetDevice(s string) *UserUpdateOne {
	uuo.mutation.SetDevice(s)
	return uuo
}

// SetNillableDevice sets the "device" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDevice(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetDevice(*s)
	}
	return uuo
}

// SetVerified sets the "verified" field.
func (uuo *UserUpdateOne) SetVerified(b bool) *UserUpdateOne {
	uuo.mutation.SetVerified(b)
	return uuo
}

// SetNillableVerified sets the "verified" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableVerified(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetVerified(*b)
	}
	return uuo
}

// SetBlocked sets the "blocked" field.
func (uuo *UserUpdateOne) SetBlocked(b bool) *UserUpdateOne {
	uuo.mutation.SetBlocked(b)
	return uuo
}

// SetNillableBlocked sets the "blocked" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBlocked(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetBlocked(*b)
	}
	return uuo
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUsername(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUsername(*s)
	}
	return uuo
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UserUpdateOne) SetCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUpdatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetUpdatedAt(*t)
	}
	return uuo
}

// AddLoginIDs adds the "logins" edge to the Login entity by IDs.
func (uuo *UserUpdateOne) AddLoginIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddLoginIDs(ids...)
	return uuo
}

// AddLogins adds the "logins" edges to the Login entity.
func (uuo *UserUpdateOne) AddLogins(l ...*Login) *UserUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uuo.AddLoginIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearLogins clears all "logins" edges to the Login entity.
func (uuo *UserUpdateOne) ClearLogins() *UserUpdateOne {
	uuo.mutation.ClearLogins()
	return uuo
}

// RemoveLoginIDs removes the "logins" edge to Login entities by IDs.
func (uuo *UserUpdateOne) RemoveLoginIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveLoginIDs(ids...)
	return uuo
}

// RemoveLogins removes "logins" edges to Login entities.
func (uuo *UserUpdateOne) RemoveLogins(l ...*Login) *UserUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uuo.RemoveLoginIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.IP(); ok {
		_spec.SetField(user.FieldIP, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Device(); ok {
		_spec.SetField(user.FieldDevice, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Verified(); ok {
		_spec.SetField(user.FieldVerified, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.Blocked(); ok {
		_spec.SetField(user.FieldBlocked, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if uuo.mutation.LoginsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.LoginsTable,
			Columns: []string{user.LoginsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(login.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedLoginsIDs(); len(nodes) > 0 && !uuo.mutation.LoginsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.LoginsTable,
			Columns: []string{user.LoginsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(login.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.LoginsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.LoginsTable,
			Columns: []string{user.LoginsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(login.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}