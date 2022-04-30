// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldBio holds the string denoting the bio field in the database.
	FieldBio = "bio"
	// FieldAdmin holds the string denoting the admin field in the database.
	FieldAdmin = "admin"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deletedat field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// EdgeReacts holds the string denoting the reacts edge name in mutations.
	EdgeReacts = "reacts"
	// EdgeSubscriptions holds the string denoting the subscriptions edge name in mutations.
	EdgeSubscriptions = "subscriptions"
	// EdgeAuthentication holds the string denoting the authentication edge name in mutations.
	EdgeAuthentication = "authentication"
	// Table holds the table name of the user in the database.
	Table = "users"
	// PostsTable is the table that holds the posts relation/edge.
	PostsTable = "posts"
	// PostsInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostsInverseTable = "posts"
	// PostsColumn is the table column denoting the posts relation/edge.
	PostsColumn = "user_posts"
	// ReactsTable is the table that holds the reacts relation/edge.
	ReactsTable = "reacts"
	// ReactsInverseTable is the table name for the React entity.
	// It exists in this package in order to avoid circular dependency with the "react" package.
	ReactsInverseTable = "reacts"
	// ReactsColumn is the table column denoting the reacts relation/edge.
	ReactsColumn = "user_reacts"
	// SubscriptionsTable is the table that holds the subscriptions relation/edge.
	SubscriptionsTable = "subscriptions"
	// SubscriptionsInverseTable is the table name for the Subscription entity.
	// It exists in this package in order to avoid circular dependency with the "subscription" package.
	SubscriptionsInverseTable = "subscriptions"
	// SubscriptionsColumn is the table column denoting the subscriptions relation/edge.
	SubscriptionsColumn = "user_subscriptions"
	// AuthenticationTable is the table that holds the authentication relation/edge.
	AuthenticationTable = "authentications"
	// AuthenticationInverseTable is the table name for the Authentication entity.
	// It exists in this package in order to avoid circular dependency with the "authentication" package.
	AuthenticationInverseTable = "authentications"
	// AuthenticationColumn is the table column denoting the authentication relation/edge.
	AuthenticationColumn = "user_authentication"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldName,
	FieldBio,
	FieldAdmin,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultAdmin holds the default value on creation for the "admin" field.
	DefaultAdmin bool
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updatedAt" field.
	DefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
