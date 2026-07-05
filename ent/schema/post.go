package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("title").MaxLen(200),
		field.Text("content"),
		field.String("category").MaxLen(100),
		field.Enum("status").Values(
			"Publish",
			"Draft",
			"Trash",
		).Default("Draft"),
		field.Time("created_date").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return nil
}
