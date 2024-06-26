// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/mars1385/e-com-go/auth/ent/login"
	"github.com/mars1385/e-com-go/auth/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LoginDelete is the builder for deleting a Login entity.
type LoginDelete struct {
	config
	hooks    []Hook
	mutation *LoginMutation
}

// Where appends a list predicates to the LoginDelete builder.
func (ld *LoginDelete) Where(ps ...predicate.Login) *LoginDelete {
	ld.mutation.Where(ps...)
	return ld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ld *LoginDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ld.sqlExec, ld.mutation, ld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ld *LoginDelete) ExecX(ctx context.Context) int {
	n, err := ld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ld *LoginDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(login.Table, sqlgraph.NewFieldSpec(login.FieldID, field.TypeInt))
	if ps := ld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ld.mutation.done = true
	return affected, err
}

// LoginDeleteOne is the builder for deleting a single Login entity.
type LoginDeleteOne struct {
	ld *LoginDelete
}

// Where appends a list predicates to the LoginDelete builder.
func (ldo *LoginDeleteOne) Where(ps ...predicate.Login) *LoginDeleteOne {
	ldo.ld.mutation.Where(ps...)
	return ldo
}

// Exec executes the deletion query.
func (ldo *LoginDeleteOne) Exec(ctx context.Context) error {
	n, err := ldo.ld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{login.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ldo *LoginDeleteOne) ExecX(ctx context.Context) {
	if err := ldo.Exec(ctx); err != nil {
		panic(err)
	}
}
