// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIP holds the string denoting the ip field in the database.
	FieldIP = "ip"
	// FieldDevice holds the string denoting the device field in the database.
	FieldDevice = "device"
	// FieldVerified holds the string denoting the verified field in the database.
	FieldVerified = "verified"
	// FieldBlocked holds the string denoting the blocked field in the database.
	FieldBlocked = "blocked"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeLogins holds the string denoting the logins edge name in mutations.
	EdgeLogins = "logins"
	// Table holds the table name of the user in the database.
	Table = "users"
	// LoginsTable is the table that holds the logins relation/edge.
	LoginsTable = "logins"
	// LoginsInverseTable is the table name for the Login entity.
	// It exists in this package in order to avoid circular dependency with the "login" package.
	LoginsInverseTable = "logins"
	// LoginsColumn is the table column denoting the logins relation/edge.
	LoginsColumn = "user_logins"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldPassword,
	FieldName,
	FieldIP,
	FieldDevice,
	FieldVerified,
	FieldBlocked,
	FieldUsername,
	FieldCreatedAt,
	FieldUpdatedAt,
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
	// DefaultVerified holds the default value on creation for the "verified" field.
	DefaultVerified bool
	// DefaultBlocked holds the default value on creation for the "blocked" field.
	DefaultBlocked bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByIP orders the results by the ip field.
func ByIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIP, opts...).ToFunc()
}

// ByDevice orders the results by the device field.
func ByDevice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDevice, opts...).ToFunc()
}

// ByVerified orders the results by the verified field.
func ByVerified(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVerified, opts...).ToFunc()
}

// ByBlocked orders the results by the blocked field.
func ByBlocked(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBlocked, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByLoginsCount orders the results by logins count.
func ByLoginsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLoginsStep(), opts...)
	}
}

// ByLogins orders the results by logins terms.
func ByLogins(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLoginsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newLoginsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LoginsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LoginsTable, LoginsColumn),
	)
}
