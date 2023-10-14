// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/ent/account"
	"github.com/Southclaws/storyden/internal/ent/asset"
	"github.com/Southclaws/storyden/internal/ent/cluster"
	"github.com/Southclaws/storyden/internal/ent/item"
	"github.com/Southclaws/storyden/internal/ent/tag"
	"github.com/rs/xid"
)

// ItemCreate is the builder for creating a Item entity.
type ItemCreate struct {
	config
	mutation *ItemMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ic *ItemCreate) SetCreatedAt(t time.Time) *ItemCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ic *ItemCreate) SetNillableCreatedAt(t *time.Time) *ItemCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the "updated_at" field.
func (ic *ItemCreate) SetUpdatedAt(t time.Time) *ItemCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ic *ItemCreate) SetNillableUpdatedAt(t *time.Time) *ItemCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetDeletedAt sets the "deleted_at" field.
func (ic *ItemCreate) SetDeletedAt(t time.Time) *ItemCreate {
	ic.mutation.SetDeletedAt(t)
	return ic
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ic *ItemCreate) SetNillableDeletedAt(t *time.Time) *ItemCreate {
	if t != nil {
		ic.SetDeletedAt(*t)
	}
	return ic
}

// SetName sets the "name" field.
func (ic *ItemCreate) SetName(s string) *ItemCreate {
	ic.mutation.SetName(s)
	return ic
}

// SetSlug sets the "slug" field.
func (ic *ItemCreate) SetSlug(s string) *ItemCreate {
	ic.mutation.SetSlug(s)
	return ic
}

// SetImageURL sets the "image_url" field.
func (ic *ItemCreate) SetImageURL(s string) *ItemCreate {
	ic.mutation.SetImageURL(s)
	return ic
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (ic *ItemCreate) SetNillableImageURL(s *string) *ItemCreate {
	if s != nil {
		ic.SetImageURL(*s)
	}
	return ic
}

// SetDescription sets the "description" field.
func (ic *ItemCreate) SetDescription(s string) *ItemCreate {
	ic.mutation.SetDescription(s)
	return ic
}

// SetAccountID sets the "account_id" field.
func (ic *ItemCreate) SetAccountID(x xid.ID) *ItemCreate {
	ic.mutation.SetAccountID(x)
	return ic
}

// SetProperties sets the "properties" field.
func (ic *ItemCreate) SetProperties(a any) *ItemCreate {
	ic.mutation.SetProperties(a)
	return ic
}

// SetID sets the "id" field.
func (ic *ItemCreate) SetID(x xid.ID) *ItemCreate {
	ic.mutation.SetID(x)
	return ic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ic *ItemCreate) SetNillableID(x *xid.ID) *ItemCreate {
	if x != nil {
		ic.SetID(*x)
	}
	return ic
}

// SetOwnerID sets the "owner" edge to the Account entity by ID.
func (ic *ItemCreate) SetOwnerID(id xid.ID) *ItemCreate {
	ic.mutation.SetOwnerID(id)
	return ic
}

// SetOwner sets the "owner" edge to the Account entity.
func (ic *ItemCreate) SetOwner(a *Account) *ItemCreate {
	return ic.SetOwnerID(a.ID)
}

// AddClusterIDs adds the "clusters" edge to the Cluster entity by IDs.
func (ic *ItemCreate) AddClusterIDs(ids ...xid.ID) *ItemCreate {
	ic.mutation.AddClusterIDs(ids...)
	return ic
}

// AddClusters adds the "clusters" edges to the Cluster entity.
func (ic *ItemCreate) AddClusters(c ...*Cluster) *ItemCreate {
	ids := make([]xid.ID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ic.AddClusterIDs(ids...)
}

// AddAssetIDs adds the "assets" edge to the Asset entity by IDs.
func (ic *ItemCreate) AddAssetIDs(ids ...string) *ItemCreate {
	ic.mutation.AddAssetIDs(ids...)
	return ic
}

// AddAssets adds the "assets" edges to the Asset entity.
func (ic *ItemCreate) AddAssets(a ...*Asset) *ItemCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ic.AddAssetIDs(ids...)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (ic *ItemCreate) AddTagIDs(ids ...xid.ID) *ItemCreate {
	ic.mutation.AddTagIDs(ids...)
	return ic
}

// AddTags adds the "tags" edges to the Tag entity.
func (ic *ItemCreate) AddTags(t ...*Tag) *ItemCreate {
	ids := make([]xid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ic.AddTagIDs(ids...)
}

// Mutation returns the ItemMutation object of the builder.
func (ic *ItemCreate) Mutation() *ItemMutation {
	return ic.mutation
}

// Save creates the Item in the database.
func (ic *ItemCreate) Save(ctx context.Context) (*Item, error) {
	ic.defaults()
	return withHooks[*Item, ItemMutation](ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *ItemCreate) SaveX(ctx context.Context) *Item {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *ItemCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *ItemCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *ItemCreate) defaults() {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := item.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		v := item.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
	if _, ok := ic.mutation.ID(); !ok {
		v := item.DefaultID()
		ic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *ItemCreate) check() error {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Item.created_at"`)}
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Item.updated_at"`)}
	}
	if _, ok := ic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Item.name"`)}
	}
	if _, ok := ic.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "Item.slug"`)}
	}
	if _, ok := ic.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Item.description"`)}
	}
	if _, ok := ic.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account_id", err: errors.New(`ent: missing required field "Item.account_id"`)}
	}
	if v, ok := ic.mutation.ID(); ok {
		if err := item.IDValidator(v.String()); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Item.id": %w`, err)}
		}
	}
	if _, ok := ic.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Item.owner"`)}
	}
	return nil
}

func (ic *ItemCreate) sqlSave(ctx context.Context) (*Item, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*xid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *ItemCreate) createSpec() (*Item, *sqlgraph.CreateSpec) {
	var (
		_node = &Item{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(item.Table, sqlgraph.NewFieldSpec(item.FieldID, field.TypeString))
	)
	_spec.OnConflict = ic.conflict
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.SetField(item.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.SetField(item.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ic.mutation.DeletedAt(); ok {
		_spec.SetField(item.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := ic.mutation.Name(); ok {
		_spec.SetField(item.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ic.mutation.Slug(); ok {
		_spec.SetField(item.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if value, ok := ic.mutation.ImageURL(); ok {
		_spec.SetField(item.FieldImageURL, field.TypeString, value)
		_node.ImageURL = &value
	}
	if value, ok := ic.mutation.Description(); ok {
		_spec.SetField(item.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ic.mutation.Properties(); ok {
		_spec.SetField(item.FieldProperties, field.TypeJSON, value)
		_node.Properties = value
	}
	if nodes := ic.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.OwnerTable,
			Columns: []string{item.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AccountID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.ClustersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   item.ClustersTable,
			Columns: item.ClustersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cluster.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.AssetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   item.AssetsTable,
			Columns: item.AssetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   item.TagsTable,
			Columns: item.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeString),
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
//	client.Item.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ItemUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ic *ItemCreate) OnConflict(opts ...sql.ConflictOption) *ItemUpsertOne {
	ic.conflict = opts
	return &ItemUpsertOne{
		create: ic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Item.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ic *ItemCreate) OnConflictColumns(columns ...string) *ItemUpsertOne {
	ic.conflict = append(ic.conflict, sql.ConflictColumns(columns...))
	return &ItemUpsertOne{
		create: ic,
	}
}

type (
	// ItemUpsertOne is the builder for "upsert"-ing
	//  one Item node.
	ItemUpsertOne struct {
		create *ItemCreate
	}

	// ItemUpsert is the "OnConflict" setter.
	ItemUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *ItemUpsert) SetUpdatedAt(v time.Time) *ItemUpsert {
	u.Set(item.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ItemUpsert) UpdateUpdatedAt() *ItemUpsert {
	u.SetExcluded(item.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ItemUpsert) SetDeletedAt(v time.Time) *ItemUpsert {
	u.Set(item.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ItemUpsert) UpdateDeletedAt() *ItemUpsert {
	u.SetExcluded(item.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ItemUpsert) ClearDeletedAt() *ItemUpsert {
	u.SetNull(item.FieldDeletedAt)
	return u
}

// SetName sets the "name" field.
func (u *ItemUpsert) SetName(v string) *ItemUpsert {
	u.Set(item.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ItemUpsert) UpdateName() *ItemUpsert {
	u.SetExcluded(item.FieldName)
	return u
}

// SetSlug sets the "slug" field.
func (u *ItemUpsert) SetSlug(v string) *ItemUpsert {
	u.Set(item.FieldSlug, v)
	return u
}

// UpdateSlug sets the "slug" field to the value that was provided on create.
func (u *ItemUpsert) UpdateSlug() *ItemUpsert {
	u.SetExcluded(item.FieldSlug)
	return u
}

// SetImageURL sets the "image_url" field.
func (u *ItemUpsert) SetImageURL(v string) *ItemUpsert {
	u.Set(item.FieldImageURL, v)
	return u
}

// UpdateImageURL sets the "image_url" field to the value that was provided on create.
func (u *ItemUpsert) UpdateImageURL() *ItemUpsert {
	u.SetExcluded(item.FieldImageURL)
	return u
}

// ClearImageURL clears the value of the "image_url" field.
func (u *ItemUpsert) ClearImageURL() *ItemUpsert {
	u.SetNull(item.FieldImageURL)
	return u
}

// SetDescription sets the "description" field.
func (u *ItemUpsert) SetDescription(v string) *ItemUpsert {
	u.Set(item.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ItemUpsert) UpdateDescription() *ItemUpsert {
	u.SetExcluded(item.FieldDescription)
	return u
}

// SetAccountID sets the "account_id" field.
func (u *ItemUpsert) SetAccountID(v xid.ID) *ItemUpsert {
	u.Set(item.FieldAccountID, v)
	return u
}

// UpdateAccountID sets the "account_id" field to the value that was provided on create.
func (u *ItemUpsert) UpdateAccountID() *ItemUpsert {
	u.SetExcluded(item.FieldAccountID)
	return u
}

// SetProperties sets the "properties" field.
func (u *ItemUpsert) SetProperties(v any) *ItemUpsert {
	u.Set(item.FieldProperties, v)
	return u
}

// UpdateProperties sets the "properties" field to the value that was provided on create.
func (u *ItemUpsert) UpdateProperties() *ItemUpsert {
	u.SetExcluded(item.FieldProperties)
	return u
}

// ClearProperties clears the value of the "properties" field.
func (u *ItemUpsert) ClearProperties() *ItemUpsert {
	u.SetNull(item.FieldProperties)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Item.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(item.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ItemUpsertOne) UpdateNewValues() *ItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(item.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(item.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Item.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ItemUpsertOne) Ignore() *ItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ItemUpsertOne) DoNothing() *ItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ItemCreate.OnConflict
// documentation for more info.
func (u *ItemUpsertOne) Update(set func(*ItemUpsert)) *ItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ItemUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ItemUpsertOne) SetUpdatedAt(v time.Time) *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ItemUpsertOne) UpdateUpdatedAt() *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ItemUpsertOne) SetDeletedAt(v time.Time) *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ItemUpsertOne) UpdateDeletedAt() *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ItemUpsertOne) ClearDeletedAt() *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.ClearDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *ItemUpsertOne) SetName(v string) *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ItemUpsertOne) UpdateName() *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateName()
	})
}

// SetSlug sets the "slug" field.
func (u *ItemUpsertOne) SetSlug(v string) *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.SetSlug(v)
	})
}

// UpdateSlug sets the "slug" field to the value that was provided on create.
func (u *ItemUpsertOne) UpdateSlug() *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateSlug()
	})
}

// SetImageURL sets the "image_url" field.
func (u *ItemUpsertOne) SetImageURL(v string) *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.SetImageURL(v)
	})
}

// UpdateImageURL sets the "image_url" field to the value that was provided on create.
func (u *ItemUpsertOne) UpdateImageURL() *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateImageURL()
	})
}

// ClearImageURL clears the value of the "image_url" field.
func (u *ItemUpsertOne) ClearImageURL() *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.ClearImageURL()
	})
}

// SetDescription sets the "description" field.
func (u *ItemUpsertOne) SetDescription(v string) *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ItemUpsertOne) UpdateDescription() *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateDescription()
	})
}

// SetAccountID sets the "account_id" field.
func (u *ItemUpsertOne) SetAccountID(v xid.ID) *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.SetAccountID(v)
	})
}

// UpdateAccountID sets the "account_id" field to the value that was provided on create.
func (u *ItemUpsertOne) UpdateAccountID() *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateAccountID()
	})
}

// SetProperties sets the "properties" field.
func (u *ItemUpsertOne) SetProperties(v any) *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.SetProperties(v)
	})
}

// UpdateProperties sets the "properties" field to the value that was provided on create.
func (u *ItemUpsertOne) UpdateProperties() *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateProperties()
	})
}

// ClearProperties clears the value of the "properties" field.
func (u *ItemUpsertOne) ClearProperties() *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.ClearProperties()
	})
}

// Exec executes the query.
func (u *ItemUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ItemCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ItemUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ItemUpsertOne) ID(ctx context.Context) (id xid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ItemUpsertOne.ID is not supported by MySQL driver. Use ItemUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ItemUpsertOne) IDX(ctx context.Context) xid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ItemCreateBulk is the builder for creating many Item entities in bulk.
type ItemCreateBulk struct {
	config
	builders []*ItemCreate
	conflict []sql.ConflictOption
}

// Save creates the Item entities in the database.
func (icb *ItemCreateBulk) Save(ctx context.Context) ([]*Item, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Item, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ItemMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = icb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *ItemCreateBulk) SaveX(ctx context.Context) []*Item {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *ItemCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *ItemCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Item.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ItemUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (icb *ItemCreateBulk) OnConflict(opts ...sql.ConflictOption) *ItemUpsertBulk {
	icb.conflict = opts
	return &ItemUpsertBulk{
		create: icb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Item.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (icb *ItemCreateBulk) OnConflictColumns(columns ...string) *ItemUpsertBulk {
	icb.conflict = append(icb.conflict, sql.ConflictColumns(columns...))
	return &ItemUpsertBulk{
		create: icb,
	}
}

// ItemUpsertBulk is the builder for "upsert"-ing
// a bulk of Item nodes.
type ItemUpsertBulk struct {
	create *ItemCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Item.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(item.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ItemUpsertBulk) UpdateNewValues() *ItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(item.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(item.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Item.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ItemUpsertBulk) Ignore() *ItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ItemUpsertBulk) DoNothing() *ItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ItemCreateBulk.OnConflict
// documentation for more info.
func (u *ItemUpsertBulk) Update(set func(*ItemUpsert)) *ItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ItemUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ItemUpsertBulk) SetUpdatedAt(v time.Time) *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ItemUpsertBulk) UpdateUpdatedAt() *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ItemUpsertBulk) SetDeletedAt(v time.Time) *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ItemUpsertBulk) UpdateDeletedAt() *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ItemUpsertBulk) ClearDeletedAt() *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.ClearDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *ItemUpsertBulk) SetName(v string) *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ItemUpsertBulk) UpdateName() *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateName()
	})
}

// SetSlug sets the "slug" field.
func (u *ItemUpsertBulk) SetSlug(v string) *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.SetSlug(v)
	})
}

// UpdateSlug sets the "slug" field to the value that was provided on create.
func (u *ItemUpsertBulk) UpdateSlug() *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateSlug()
	})
}

// SetImageURL sets the "image_url" field.
func (u *ItemUpsertBulk) SetImageURL(v string) *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.SetImageURL(v)
	})
}

// UpdateImageURL sets the "image_url" field to the value that was provided on create.
func (u *ItemUpsertBulk) UpdateImageURL() *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateImageURL()
	})
}

// ClearImageURL clears the value of the "image_url" field.
func (u *ItemUpsertBulk) ClearImageURL() *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.ClearImageURL()
	})
}

// SetDescription sets the "description" field.
func (u *ItemUpsertBulk) SetDescription(v string) *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ItemUpsertBulk) UpdateDescription() *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateDescription()
	})
}

// SetAccountID sets the "account_id" field.
func (u *ItemUpsertBulk) SetAccountID(v xid.ID) *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.SetAccountID(v)
	})
}

// UpdateAccountID sets the "account_id" field to the value that was provided on create.
func (u *ItemUpsertBulk) UpdateAccountID() *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateAccountID()
	})
}

// SetProperties sets the "properties" field.
func (u *ItemUpsertBulk) SetProperties(v any) *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.SetProperties(v)
	})
}

// UpdateProperties sets the "properties" field to the value that was provided on create.
func (u *ItemUpsertBulk) UpdateProperties() *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateProperties()
	})
}

// ClearProperties clears the value of the "properties" field.
func (u *ItemUpsertBulk) ClearProperties() *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.ClearProperties()
	})
}

// Exec executes the query.
func (u *ItemUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ItemCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ItemCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ItemUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
