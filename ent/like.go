// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"hyphen-hellog/ent/author"
	"hyphen-hellog/ent/like"
	"hyphen-hellog/ent/post"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Like is the model entity for the Like schema.
type Like struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LikeQuery when eager-loading is set.
	Edges        LikeEdges `json:"edges"`
	author_likes *int
	post_likes   *int
	selectValues sql.SelectValues
}

// LikeEdges holds the relations/edges for other nodes in the graph.
type LikeEdges struct {
	// Author holds the value of the author edge.
	Author *Author `json:"author,omitempty"`
	// Post holds the value of the post edge.
	Post *Post `json:"post,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LikeEdges) AuthorOrErr() (*Author, error) {
	if e.loadedTypes[0] {
		if e.Author == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: author.Label}
		}
		return e.Author, nil
	}
	return nil, &NotLoadedError{edge: "author"}
}

// PostOrErr returns the Post value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LikeEdges) PostOrErr() (*Post, error) {
	if e.loadedTypes[1] {
		if e.Post == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: post.Label}
		}
		return e.Post, nil
	}
	return nil, &NotLoadedError{edge: "post"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Like) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case like.FieldID:
			values[i] = new(sql.NullInt64)
		case like.ForeignKeys[0]: // author_likes
			values[i] = new(sql.NullInt64)
		case like.ForeignKeys[1]: // post_likes
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Like fields.
func (l *Like) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case like.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		case like.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field author_likes", value)
			} else if value.Valid {
				l.author_likes = new(int)
				*l.author_likes = int(value.Int64)
			}
		case like.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field post_likes", value)
			} else if value.Valid {
				l.post_likes = new(int)
				*l.post_likes = int(value.Int64)
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Like.
// This includes values selected through modifiers, order, etc.
func (l *Like) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// QueryAuthor queries the "author" edge of the Like entity.
func (l *Like) QueryAuthor() *AuthorQuery {
	return NewLikeClient(l.config).QueryAuthor(l)
}

// QueryPost queries the "post" edge of the Like entity.
func (l *Like) QueryPost() *PostQuery {
	return NewLikeClient(l.config).QueryPost(l)
}

// Update returns a builder for updating this Like.
// Note that you need to call Like.Unwrap() before calling this method if this Like
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Like) Update() *LikeUpdateOne {
	return NewLikeClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Like entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Like) Unwrap() *Like {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Like is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Like) String() string {
	var builder strings.Builder
	builder.WriteString("Like(")
	builder.WriteString(fmt.Sprintf("id=%v", l.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Likes is a parsable slice of Like.
type Likes []*Like
