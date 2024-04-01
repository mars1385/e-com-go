package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Login holds the schema definition for the Login entity.
type Login struct {
	ent.Schema
}

// Fields of the Login.
func (Login) Fields() []ent.Field {
	return []ent.Field{
		field.String("ip"),
		field.String("device"),
		field.String("access"),
		field.String("refresh"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the Login.
func (Login) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Unique().Ref("logins"),
	}
}
