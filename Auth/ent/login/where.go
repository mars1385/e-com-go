// Code generated by ent, DO NOT EDIT.

package login

import (
	"github.com/mars1385/e-com-go/auth/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Login {
	return predicate.Login(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Login {
	return predicate.Login(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Login {
	return predicate.Login(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Login {
	return predicate.Login(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Login {
	return predicate.Login(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Login {
	return predicate.Login(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Login {
	return predicate.Login(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Login {
	return predicate.Login(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Login {
	return predicate.Login(sql.FieldLTE(FieldID, id))
}

// IP applies equality check predicate on the "ip" field. It's identical to IPEQ.
func IP(v string) predicate.Login {
	return predicate.Login(sql.FieldEQ(FieldIP, v))
}

// Device applies equality check predicate on the "device" field. It's identical to DeviceEQ.
func Device(v string) predicate.Login {
	return predicate.Login(sql.FieldEQ(FieldDevice, v))
}

// Access applies equality check predicate on the "access" field. It's identical to AccessEQ.
func Access(v string) predicate.Login {
	return predicate.Login(sql.FieldEQ(FieldAccess, v))
}

// Refresh applies equality check predicate on the "refresh" field. It's identical to RefreshEQ.
func Refresh(v string) predicate.Login {
	return predicate.Login(sql.FieldEQ(FieldRefresh, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Login {
	return predicate.Login(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Login {
	return predicate.Login(sql.FieldEQ(FieldUpdatedAt, v))
}

// IPEQ applies the EQ predicate on the "ip" field.
func IPEQ(v string) predicate.Login {
	return predicate.Login(sql.FieldEQ(FieldIP, v))
}

// IPNEQ applies the NEQ predicate on the "ip" field.
func IPNEQ(v string) predicate.Login {
	return predicate.Login(sql.FieldNEQ(FieldIP, v))
}

// IPIn applies the In predicate on the "ip" field.
func IPIn(vs ...string) predicate.Login {
	return predicate.Login(sql.FieldIn(FieldIP, vs...))
}

// IPNotIn applies the NotIn predicate on the "ip" field.
func IPNotIn(vs ...string) predicate.Login {
	return predicate.Login(sql.FieldNotIn(FieldIP, vs...))
}

// IPGT applies the GT predicate on the "ip" field.
func IPGT(v string) predicate.Login {
	return predicate.Login(sql.FieldGT(FieldIP, v))
}

// IPGTE applies the GTE predicate on the "ip" field.
func IPGTE(v string) predicate.Login {
	return predicate.Login(sql.FieldGTE(FieldIP, v))
}

// IPLT applies the LT predicate on the "ip" field.
func IPLT(v string) predicate.Login {
	return predicate.Login(sql.FieldLT(FieldIP, v))
}

// IPLTE applies the LTE predicate on the "ip" field.
func IPLTE(v string) predicate.Login {
	return predicate.Login(sql.FieldLTE(FieldIP, v))
}

// IPContains applies the Contains predicate on the "ip" field.
func IPContains(v string) predicate.Login {
	return predicate.Login(sql.FieldContains(FieldIP, v))
}

// IPHasPrefix applies the HasPrefix predicate on the "ip" field.
func IPHasPrefix(v string) predicate.Login {
	return predicate.Login(sql.FieldHasPrefix(FieldIP, v))
}

// IPHasSuffix applies the HasSuffix predicate on the "ip" field.
func IPHasSuffix(v string) predicate.Login {
	return predicate.Login(sql.FieldHasSuffix(FieldIP, v))
}

// IPEqualFold applies the EqualFold predicate on the "ip" field.
func IPEqualFold(v string) predicate.Login {
	return predicate.Login(sql.FieldEqualFold(FieldIP, v))
}

// IPContainsFold applies the ContainsFold predicate on the "ip" field.
func IPContainsFold(v string) predicate.Login {
	return predicate.Login(sql.FieldContainsFold(FieldIP, v))
}

// DeviceEQ applies the EQ predicate on the "device" field.
func DeviceEQ(v string) predicate.Login {
	return predicate.Login(sql.FieldEQ(FieldDevice, v))
}

// DeviceNEQ applies the NEQ predicate on the "device" field.
func DeviceNEQ(v string) predicate.Login {
	return predicate.Login(sql.FieldNEQ(FieldDevice, v))
}

// DeviceIn applies the In predicate on the "device" field.
func DeviceIn(vs ...string) predicate.Login {
	return predicate.Login(sql.FieldIn(FieldDevice, vs...))
}

// DeviceNotIn applies the NotIn predicate on the "device" field.
func DeviceNotIn(vs ...string) predicate.Login {
	return predicate.Login(sql.FieldNotIn(FieldDevice, vs...))
}

// DeviceGT applies the GT predicate on the "device" field.
func DeviceGT(v string) predicate.Login {
	return predicate.Login(sql.FieldGT(FieldDevice, v))
}

// DeviceGTE applies the GTE predicate on the "device" field.
func DeviceGTE(v string) predicate.Login {
	return predicate.Login(sql.FieldGTE(FieldDevice, v))
}

// DeviceLT applies the LT predicate on the "device" field.
func DeviceLT(v string) predicate.Login {
	return predicate.Login(sql.FieldLT(FieldDevice, v))
}

// DeviceLTE applies the LTE predicate on the "device" field.
func DeviceLTE(v string) predicate.Login {
	return predicate.Login(sql.FieldLTE(FieldDevice, v))
}

// DeviceContains applies the Contains predicate on the "device" field.
func DeviceContains(v string) predicate.Login {
	return predicate.Login(sql.FieldContains(FieldDevice, v))
}

// DeviceHasPrefix applies the HasPrefix predicate on the "device" field.
func DeviceHasPrefix(v string) predicate.Login {
	return predicate.Login(sql.FieldHasPrefix(FieldDevice, v))
}

// DeviceHasSuffix applies the HasSuffix predicate on the "device" field.
func DeviceHasSuffix(v string) predicate.Login {
	return predicate.Login(sql.FieldHasSuffix(FieldDevice, v))
}

// DeviceEqualFold applies the EqualFold predicate on the "device" field.
func DeviceEqualFold(v string) predicate.Login {
	return predicate.Login(sql.FieldEqualFold(FieldDevice, v))
}

// DeviceContainsFold applies the ContainsFold predicate on the "device" field.
func DeviceContainsFold(v string) predicate.Login {
	return predicate.Login(sql.FieldContainsFold(FieldDevice, v))
}

// AccessEQ applies the EQ predicate on the "access" field.
func AccessEQ(v string) predicate.Login {
	return predicate.Login(sql.FieldEQ(FieldAccess, v))
}

// AccessNEQ applies the NEQ predicate on the "access" field.
func AccessNEQ(v string) predicate.Login {
	return predicate.Login(sql.FieldNEQ(FieldAccess, v))
}

// AccessIn applies the In predicate on the "access" field.
func AccessIn(vs ...string) predicate.Login {
	return predicate.Login(sql.FieldIn(FieldAccess, vs...))
}

// AccessNotIn applies the NotIn predicate on the "access" field.
func AccessNotIn(vs ...string) predicate.Login {
	return predicate.Login(sql.FieldNotIn(FieldAccess, vs...))
}

// AccessGT applies the GT predicate on the "access" field.
func AccessGT(v string) predicate.Login {
	return predicate.Login(sql.FieldGT(FieldAccess, v))
}

// AccessGTE applies the GTE predicate on the "access" field.
func AccessGTE(v string) predicate.Login {
	return predicate.Login(sql.FieldGTE(FieldAccess, v))
}

// AccessLT applies the LT predicate on the "access" field.
func AccessLT(v string) predicate.Login {
	return predicate.Login(sql.FieldLT(FieldAccess, v))
}

// AccessLTE applies the LTE predicate on the "access" field.
func AccessLTE(v string) predicate.Login {
	return predicate.Login(sql.FieldLTE(FieldAccess, v))
}

// AccessContains applies the Contains predicate on the "access" field.
func AccessContains(v string) predicate.Login {
	return predicate.Login(sql.FieldContains(FieldAccess, v))
}

// AccessHasPrefix applies the HasPrefix predicate on the "access" field.
func AccessHasPrefix(v string) predicate.Login {
	return predicate.Login(sql.FieldHasPrefix(FieldAccess, v))
}

// AccessHasSuffix applies the HasSuffix predicate on the "access" field.
func AccessHasSuffix(v string) predicate.Login {
	return predicate.Login(sql.FieldHasSuffix(FieldAccess, v))
}

// AccessEqualFold applies the EqualFold predicate on the "access" field.
func AccessEqualFold(v string) predicate.Login {
	return predicate.Login(sql.FieldEqualFold(FieldAccess, v))
}

// AccessContainsFold applies the ContainsFold predicate on the "access" field.
func AccessContainsFold(v string) predicate.Login {
	return predicate.Login(sql.FieldContainsFold(FieldAccess, v))
}

// RefreshEQ applies the EQ predicate on the "refresh" field.
func RefreshEQ(v string) predicate.Login {
	return predicate.Login(sql.FieldEQ(FieldRefresh, v))
}

// RefreshNEQ applies the NEQ predicate on the "refresh" field.
func RefreshNEQ(v string) predicate.Login {
	return predicate.Login(sql.FieldNEQ(FieldRefresh, v))
}

// RefreshIn applies the In predicate on the "refresh" field.
func RefreshIn(vs ...string) predicate.Login {
	return predicate.Login(sql.FieldIn(FieldRefresh, vs...))
}

// RefreshNotIn applies the NotIn predicate on the "refresh" field.
func RefreshNotIn(vs ...string) predicate.Login {
	return predicate.Login(sql.FieldNotIn(FieldRefresh, vs...))
}

// RefreshGT applies the GT predicate on the "refresh" field.
func RefreshGT(v string) predicate.Login {
	return predicate.Login(sql.FieldGT(FieldRefresh, v))
}

// RefreshGTE applies the GTE predicate on the "refresh" field.
func RefreshGTE(v string) predicate.Login {
	return predicate.Login(sql.FieldGTE(FieldRefresh, v))
}

// RefreshLT applies the LT predicate on the "refresh" field.
func RefreshLT(v string) predicate.Login {
	return predicate.Login(sql.FieldLT(FieldRefresh, v))
}

// RefreshLTE applies the LTE predicate on the "refresh" field.
func RefreshLTE(v string) predicate.Login {
	return predicate.Login(sql.FieldLTE(FieldRefresh, v))
}

// RefreshContains applies the Contains predicate on the "refresh" field.
func RefreshContains(v string) predicate.Login {
	return predicate.Login(sql.FieldContains(FieldRefresh, v))
}

// RefreshHasPrefix applies the HasPrefix predicate on the "refresh" field.
func RefreshHasPrefix(v string) predicate.Login {
	return predicate.Login(sql.FieldHasPrefix(FieldRefresh, v))
}

// RefreshHasSuffix applies the HasSuffix predicate on the "refresh" field.
func RefreshHasSuffix(v string) predicate.Login {
	return predicate.Login(sql.FieldHasSuffix(FieldRefresh, v))
}

// RefreshEqualFold applies the EqualFold predicate on the "refresh" field.
func RefreshEqualFold(v string) predicate.Login {
	return predicate.Login(sql.FieldEqualFold(FieldRefresh, v))
}

// RefreshContainsFold applies the ContainsFold predicate on the "refresh" field.
func RefreshContainsFold(v string) predicate.Login {
	return predicate.Login(sql.FieldContainsFold(FieldRefresh, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Login {
	return predicate.Login(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Login {
	return predicate.Login(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Login {
	return predicate.Login(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Login {
	return predicate.Login(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Login {
	return predicate.Login(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Login {
	return predicate.Login(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Login {
	return predicate.Login(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Login {
	return predicate.Login(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Login {
	return predicate.Login(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Login {
	return predicate.Login(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Login {
	return predicate.Login(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Login {
	return predicate.Login(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Login {
	return predicate.Login(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Login {
	return predicate.Login(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Login {
	return predicate.Login(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Login {
	return predicate.Login(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Login {
	return predicate.Login(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Login {
	return predicate.Login(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Login) predicate.Login {
	return predicate.Login(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Login) predicate.Login {
	return predicate.Login(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Login) predicate.Login {
	return predicate.Login(sql.NotPredicates(p))
}
