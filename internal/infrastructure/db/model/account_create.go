// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/infrastructure/db/model/account"
	"github.com/Southclaws/storyden/internal/infrastructure/db/model/authentication"
	"github.com/Southclaws/storyden/internal/infrastructure/db/model/post"
	"github.com/Southclaws/storyden/internal/infrastructure/db/model/react"
	"github.com/Southclaws/storyden/internal/infrastructure/db/model/role"
	"github.com/Southclaws/storyden/internal/infrastructure/db/model/subscription"
	"github.com/google/uuid"
)

// AccountCreate is the builder for creating a Account entity.
type AccountCreate struct {
	config
	mutation *AccountMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetEmail sets the "email" field.
func (ac *AccountCreate) SetEmail(s string) *AccountCreate {
	ac.mutation.SetEmail(s)
	return ac
}

// SetName sets the "name" field.
func (ac *AccountCreate) SetName(s string) *AccountCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetBio sets the "bio" field.
func (ac *AccountCreate) SetBio(s string) *AccountCreate {
	ac.mutation.SetBio(s)
	return ac
}

// SetNillableBio sets the "bio" field if the given value is not nil.
func (ac *AccountCreate) SetNillableBio(s *string) *AccountCreate {
	if s != nil {
		ac.SetBio(*s)
	}
	return ac
}

// SetAdmin sets the "admin" field.
func (ac *AccountCreate) SetAdmin(b bool) *AccountCreate {
	ac.mutation.SetAdmin(b)
	return ac
}

// SetNillableAdmin sets the "admin" field if the given value is not nil.
func (ac *AccountCreate) SetNillableAdmin(b *bool) *AccountCreate {
	if b != nil {
		ac.SetAdmin(*b)
	}
	return ac
}

// SetCreatedAt sets the "createdAt" field.
func (ac *AccountCreate) SetCreatedAt(t time.Time) *AccountCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (ac *AccountCreate) SetNillableCreatedAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updatedAt" field.
func (ac *AccountCreate) SetUpdatedAt(t time.Time) *AccountCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (ac *AccountCreate) SetNillableUpdatedAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetDeletedAt sets the "deletedAt" field.
func (ac *AccountCreate) SetDeletedAt(t time.Time) *AccountCreate {
	ac.mutation.SetDeletedAt(t)
	return ac
}

// SetNillableDeletedAt sets the "deletedAt" field if the given value is not nil.
func (ac *AccountCreate) SetNillableDeletedAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetDeletedAt(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AccountCreate) SetID(u uuid.UUID) *AccountCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AccountCreate) SetNillableID(u *uuid.UUID) *AccountCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (ac *AccountCreate) AddPostIDs(ids ...uuid.UUID) *AccountCreate {
	ac.mutation.AddPostIDs(ids...)
	return ac
}

// AddPosts adds the "posts" edges to the Post entity.
func (ac *AccountCreate) AddPosts(p ...*Post) *AccountCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ac.AddPostIDs(ids...)
}

// AddReactIDs adds the "reacts" edge to the React entity by IDs.
func (ac *AccountCreate) AddReactIDs(ids ...uuid.UUID) *AccountCreate {
	ac.mutation.AddReactIDs(ids...)
	return ac
}

// AddReacts adds the "reacts" edges to the React entity.
func (ac *AccountCreate) AddReacts(r ...*React) *AccountCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ac.AddReactIDs(ids...)
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (ac *AccountCreate) AddRoleIDs(ids ...uuid.UUID) *AccountCreate {
	ac.mutation.AddRoleIDs(ids...)
	return ac
}

// AddRoles adds the "roles" edges to the Role entity.
func (ac *AccountCreate) AddRoles(r ...*Role) *AccountCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ac.AddRoleIDs(ids...)
}

// AddSubscriptionIDs adds the "subscriptions" edge to the Subscription entity by IDs.
func (ac *AccountCreate) AddSubscriptionIDs(ids ...uuid.UUID) *AccountCreate {
	ac.mutation.AddSubscriptionIDs(ids...)
	return ac
}

// AddSubscriptions adds the "subscriptions" edges to the Subscription entity.
func (ac *AccountCreate) AddSubscriptions(s ...*Subscription) *AccountCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ac.AddSubscriptionIDs(ids...)
}

// AddAuthenticationIDs adds the "authentication" edge to the Authentication entity by IDs.
func (ac *AccountCreate) AddAuthenticationIDs(ids ...uuid.UUID) *AccountCreate {
	ac.mutation.AddAuthenticationIDs(ids...)
	return ac
}

// AddAuthentication adds the "authentication" edges to the Authentication entity.
func (ac *AccountCreate) AddAuthentication(a ...*Authentication) *AccountCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddAuthenticationIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (ac *AccountCreate) Mutation() *AccountMutation {
	return ac.mutation
}

// Save creates the Account in the database.
func (ac *AccountCreate) Save(ctx context.Context) (*Account, error) {
	var (
		err  error
		node *Account
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Account)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AccountMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AccountCreate) SaveX(ctx context.Context) *Account {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AccountCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AccountCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AccountCreate) defaults() {
	if _, ok := ac.mutation.Admin(); !ok {
		v := account.DefaultAdmin
		ac.mutation.SetAdmin(v)
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := account.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := account.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := account.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AccountCreate) check() error {
	if _, ok := ac.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`model: missing required field "Account.email"`)}
	}
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`model: missing required field "Account.name"`)}
	}
	if v, ok := ac.mutation.Name(); ok {
		if err := account.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`model: validator failed for field "Account.name": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Admin(); !ok {
		return &ValidationError{Name: "admin", err: errors.New(`model: missing required field "Account.admin"`)}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`model: missing required field "Account.createdAt"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`model: missing required field "Account.updatedAt"`)}
	}
	return nil
}

func (ac *AccountCreate) sqlSave(ctx context.Context) (*Account, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (ac *AccountCreate) createSpec() (*Account, *sqlgraph.CreateSpec) {
	var (
		_node = &Account{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: account.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: account.FieldID,
			},
		}
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ac.mutation.Bio(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldBio,
		})
		_node.Bio = value
	}
	if value, ok := ac.mutation.Admin(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: account.FieldAdmin,
		})
		_node.Admin = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if nodes := ac.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.PostsTable,
			Columns: []string{account.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.ReactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.ReactsTable,
			Columns: []string{account.ReactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: react.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   account.RolesTable,
			Columns: account.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.SubscriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.SubscriptionsTable,
			Columns: []string{account.SubscriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: subscription.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.AuthenticationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.AuthenticationTable,
			Columns: []string{account.AuthenticationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: authentication.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Account.Create().
//		SetEmail(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AccountUpsert) {
//			SetEmail(v+v).
//		}).
//		Exec(ctx)
//
func (ac *AccountCreate) OnConflict(opts ...sql.ConflictOption) *AccountUpsertOne {
	ac.conflict = opts
	return &AccountUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ac *AccountCreate) OnConflictColumns(columns ...string) *AccountUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AccountUpsertOne{
		create: ac,
	}
}

type (
	// AccountUpsertOne is the builder for "upsert"-ing
	//  one Account node.
	AccountUpsertOne struct {
		create *AccountCreate
	}

	// AccountUpsert is the "OnConflict" setter.
	AccountUpsert struct {
		*sql.UpdateSet
	}
)

// SetEmail sets the "email" field.
func (u *AccountUpsert) SetEmail(v string) *AccountUpsert {
	u.Set(account.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *AccountUpsert) UpdateEmail() *AccountUpsert {
	u.SetExcluded(account.FieldEmail)
	return u
}

// SetName sets the "name" field.
func (u *AccountUpsert) SetName(v string) *AccountUpsert {
	u.Set(account.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AccountUpsert) UpdateName() *AccountUpsert {
	u.SetExcluded(account.FieldName)
	return u
}

// SetBio sets the "bio" field.
func (u *AccountUpsert) SetBio(v string) *AccountUpsert {
	u.Set(account.FieldBio, v)
	return u
}

// UpdateBio sets the "bio" field to the value that was provided on create.
func (u *AccountUpsert) UpdateBio() *AccountUpsert {
	u.SetExcluded(account.FieldBio)
	return u
}

// ClearBio clears the value of the "bio" field.
func (u *AccountUpsert) ClearBio() *AccountUpsert {
	u.SetNull(account.FieldBio)
	return u
}

// SetAdmin sets the "admin" field.
func (u *AccountUpsert) SetAdmin(v bool) *AccountUpsert {
	u.Set(account.FieldAdmin, v)
	return u
}

// UpdateAdmin sets the "admin" field to the value that was provided on create.
func (u *AccountUpsert) UpdateAdmin() *AccountUpsert {
	u.SetExcluded(account.FieldAdmin)
	return u
}

// SetCreatedAt sets the "createdAt" field.
func (u *AccountUpsert) SetCreatedAt(v time.Time) *AccountUpsert {
	u.Set(account.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "createdAt" field to the value that was provided on create.
func (u *AccountUpsert) UpdateCreatedAt() *AccountUpsert {
	u.SetExcluded(account.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updatedAt" field.
func (u *AccountUpsert) SetUpdatedAt(v time.Time) *AccountUpsert {
	u.Set(account.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updatedAt" field to the value that was provided on create.
func (u *AccountUpsert) UpdateUpdatedAt() *AccountUpsert {
	u.SetExcluded(account.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deletedAt" field.
func (u *AccountUpsert) SetDeletedAt(v time.Time) *AccountUpsert {
	u.Set(account.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deletedAt" field to the value that was provided on create.
func (u *AccountUpsert) UpdateDeletedAt() *AccountUpsert {
	u.SetExcluded(account.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deletedAt" field.
func (u *AccountUpsert) ClearDeletedAt() *AccountUpsert {
	u.SetNull(account.FieldDeletedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(account.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AccountUpsertOne) UpdateNewValues() *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(account.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Account.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AccountUpsertOne) Ignore() *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AccountUpsertOne) DoNothing() *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AccountCreate.OnConflict
// documentation for more info.
func (u *AccountUpsertOne) Update(set func(*AccountUpsert)) *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AccountUpsert{UpdateSet: update})
	}))
	return u
}

// SetEmail sets the "email" field.
func (u *AccountUpsertOne) SetEmail(v string) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateEmail() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateEmail()
	})
}

// SetName sets the "name" field.
func (u *AccountUpsertOne) SetName(v string) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateName() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateName()
	})
}

// SetBio sets the "bio" field.
func (u *AccountUpsertOne) SetBio(v string) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetBio(v)
	})
}

// UpdateBio sets the "bio" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateBio() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateBio()
	})
}

// ClearBio clears the value of the "bio" field.
func (u *AccountUpsertOne) ClearBio() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.ClearBio()
	})
}

// SetAdmin sets the "admin" field.
func (u *AccountUpsertOne) SetAdmin(v bool) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetAdmin(v)
	})
}

// UpdateAdmin sets the "admin" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateAdmin() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateAdmin()
	})
}

// SetCreatedAt sets the "createdAt" field.
func (u *AccountUpsertOne) SetCreatedAt(v time.Time) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "createdAt" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateCreatedAt() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updatedAt" field.
func (u *AccountUpsertOne) SetUpdatedAt(v time.Time) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updatedAt" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateUpdatedAt() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deletedAt" field.
func (u *AccountUpsertOne) SetDeletedAt(v time.Time) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deletedAt" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateDeletedAt() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deletedAt" field.
func (u *AccountUpsertOne) ClearDeletedAt() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *AccountUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for AccountCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AccountUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AccountUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("model: AccountUpsertOne.ID is not supported by MySQL driver. Use AccountUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AccountUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AccountCreateBulk is the builder for creating many Account entities in bulk.
type AccountCreateBulk struct {
	config
	builders []*AccountCreate
	conflict []sql.ConflictOption
}

// Save creates the Account entities in the database.
func (acb *AccountCreateBulk) Save(ctx context.Context) ([]*Account, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Account, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AccountCreateBulk) SaveX(ctx context.Context) []*Account {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AccountCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AccountCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Account.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AccountUpsert) {
//			SetEmail(v+v).
//		}).
//		Exec(ctx)
//
func (acb *AccountCreateBulk) OnConflict(opts ...sql.ConflictOption) *AccountUpsertBulk {
	acb.conflict = opts
	return &AccountUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (acb *AccountCreateBulk) OnConflictColumns(columns ...string) *AccountUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AccountUpsertBulk{
		create: acb,
	}
}

// AccountUpsertBulk is the builder for "upsert"-ing
// a bulk of Account nodes.
type AccountUpsertBulk struct {
	create *AccountCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(account.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AccountUpsertBulk) UpdateNewValues() *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(account.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AccountUpsertBulk) Ignore() *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AccountUpsertBulk) DoNothing() *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AccountCreateBulk.OnConflict
// documentation for more info.
func (u *AccountUpsertBulk) Update(set func(*AccountUpsert)) *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AccountUpsert{UpdateSet: update})
	}))
	return u
}

// SetEmail sets the "email" field.
func (u *AccountUpsertBulk) SetEmail(v string) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateEmail() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateEmail()
	})
}

// SetName sets the "name" field.
func (u *AccountUpsertBulk) SetName(v string) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateName() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateName()
	})
}

// SetBio sets the "bio" field.
func (u *AccountUpsertBulk) SetBio(v string) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetBio(v)
	})
}

// UpdateBio sets the "bio" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateBio() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateBio()
	})
}

// ClearBio clears the value of the "bio" field.
func (u *AccountUpsertBulk) ClearBio() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.ClearBio()
	})
}

// SetAdmin sets the "admin" field.
func (u *AccountUpsertBulk) SetAdmin(v bool) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetAdmin(v)
	})
}

// UpdateAdmin sets the "admin" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateAdmin() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateAdmin()
	})
}

// SetCreatedAt sets the "createdAt" field.
func (u *AccountUpsertBulk) SetCreatedAt(v time.Time) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "createdAt" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateCreatedAt() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updatedAt" field.
func (u *AccountUpsertBulk) SetUpdatedAt(v time.Time) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updatedAt" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateUpdatedAt() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deletedAt" field.
func (u *AccountUpsertBulk) SetDeletedAt(v time.Time) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deletedAt" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateDeletedAt() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deletedAt" field.
func (u *AccountUpsertBulk) ClearDeletedAt() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *AccountUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the AccountCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for AccountCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AccountUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
