package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// App holds the schema definition for the App entity.
type App struct {
	ent.Schema
}

// Fields of the App.
func (App) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Annotations(entsql.Annotation{Collation: "utf8_general_ci"}),
		field.String("version").Annotations(entsql.Annotation{Collation: "utf8_general_ci"}),
		field.String("publisher").Optional().Annotations(entsql.Annotation{Collation: "utf8_general_ci"}),
		field.String("install_date").Optional(),
	}
}

// Edges of the App.
func (App) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Agent.Type).Unique().Ref("apps").Required(),
	}
}
