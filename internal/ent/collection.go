// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Southclaws/storyden/internal/ent/account"
	"github.com/Southclaws/storyden/internal/ent/collection"
	"github.com/rs/xid"
)

// Collection is the model entity for the Collection schema.
type Collection struct {
	config `json:"-"`
	// ID of the ent.
	ID xid.ID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CollectionQuery when eager-loading is set.
	Edges               CollectionEdges `json:"edges"`
	account_collections *xid.ID
}

// CollectionEdges holds the relations/edges for other nodes in the graph.
type CollectionEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Account `json:"owner,omitempty"`
	// Posts holds the value of the posts edge.
	Posts []*Post `json:"posts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CollectionEdges) OwnerOrErr() (*Account, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: account.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e CollectionEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[1] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Collection) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case collection.FieldName, collection.FieldDescription:
			values[i] = new(sql.NullString)
		case collection.FieldCreatedAt, collection.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case collection.FieldID:
			values[i] = new(xid.ID)
		case collection.ForeignKeys[0]: // account_collections
			values[i] = &sql.NullScanner{S: new(xid.ID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Collection", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Collection fields.
func (c *Collection) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case collection.FieldID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case collection.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case collection.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case collection.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case collection.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case collection.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field account_collections", values[i])
			} else if value.Valid {
				c.account_collections = new(xid.ID)
				*c.account_collections = *value.S.(*xid.ID)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Collection entity.
func (c *Collection) QueryOwner() *AccountQuery {
	return NewCollectionClient(c.config).QueryOwner(c)
}

// QueryPosts queries the "posts" edge of the Collection entity.
func (c *Collection) QueryPosts() *PostQuery {
	return NewCollectionClient(c.config).QueryPosts(c)
}

// Update returns a builder for updating this Collection.
// Note that you need to call Collection.Unwrap() before calling this method if this Collection
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Collection) Update() *CollectionUpdateOne {
	return NewCollectionClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Collection entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Collection) Unwrap() *Collection {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Collection is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Collection) String() string {
	var builder strings.Builder
	builder.WriteString("Collection(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(c.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Collections is a parsable slice of Collection.
type Collections []*Collection
