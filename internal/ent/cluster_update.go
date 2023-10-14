// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Southclaws/storyden/internal/ent/account"
	"github.com/Southclaws/storyden/internal/ent/asset"
	"github.com/Southclaws/storyden/internal/ent/cluster"
	"github.com/Southclaws/storyden/internal/ent/item"
	"github.com/Southclaws/storyden/internal/ent/predicate"
	"github.com/Southclaws/storyden/internal/ent/tag"
	"github.com/rs/xid"
)

// ClusterUpdate is the builder for updating Cluster entities.
type ClusterUpdate struct {
	config
	hooks     []Hook
	mutation  *ClusterMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ClusterUpdate builder.
func (cu *ClusterUpdate) Where(ps ...predicate.Cluster) *ClusterUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *ClusterUpdate) SetUpdatedAt(t time.Time) *ClusterUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetDeletedAt sets the "deleted_at" field.
func (cu *ClusterUpdate) SetDeletedAt(t time.Time) *ClusterUpdate {
	cu.mutation.SetDeletedAt(t)
	return cu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cu *ClusterUpdate) SetNillableDeletedAt(t *time.Time) *ClusterUpdate {
	if t != nil {
		cu.SetDeletedAt(*t)
	}
	return cu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cu *ClusterUpdate) ClearDeletedAt() *ClusterUpdate {
	cu.mutation.ClearDeletedAt()
	return cu
}

// SetName sets the "name" field.
func (cu *ClusterUpdate) SetName(s string) *ClusterUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetSlug sets the "slug" field.
func (cu *ClusterUpdate) SetSlug(s string) *ClusterUpdate {
	cu.mutation.SetSlug(s)
	return cu
}

// SetImageURL sets the "image_url" field.
func (cu *ClusterUpdate) SetImageURL(s string) *ClusterUpdate {
	cu.mutation.SetImageURL(s)
	return cu
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (cu *ClusterUpdate) SetNillableImageURL(s *string) *ClusterUpdate {
	if s != nil {
		cu.SetImageURL(*s)
	}
	return cu
}

// ClearImageURL clears the value of the "image_url" field.
func (cu *ClusterUpdate) ClearImageURL() *ClusterUpdate {
	cu.mutation.ClearImageURL()
	return cu
}

// SetDescription sets the "description" field.
func (cu *ClusterUpdate) SetDescription(s string) *ClusterUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetParentClusterID sets the "parent_cluster_id" field.
func (cu *ClusterUpdate) SetParentClusterID(x xid.ID) *ClusterUpdate {
	cu.mutation.SetParentClusterID(x)
	return cu
}

// SetNillableParentClusterID sets the "parent_cluster_id" field if the given value is not nil.
func (cu *ClusterUpdate) SetNillableParentClusterID(x *xid.ID) *ClusterUpdate {
	if x != nil {
		cu.SetParentClusterID(*x)
	}
	return cu
}

// ClearParentClusterID clears the value of the "parent_cluster_id" field.
func (cu *ClusterUpdate) ClearParentClusterID() *ClusterUpdate {
	cu.mutation.ClearParentClusterID()
	return cu
}

// SetAccountID sets the "account_id" field.
func (cu *ClusterUpdate) SetAccountID(x xid.ID) *ClusterUpdate {
	cu.mutation.SetAccountID(x)
	return cu
}

// SetProperties sets the "properties" field.
func (cu *ClusterUpdate) SetProperties(a any) *ClusterUpdate {
	cu.mutation.SetProperties(a)
	return cu
}

// ClearProperties clears the value of the "properties" field.
func (cu *ClusterUpdate) ClearProperties() *ClusterUpdate {
	cu.mutation.ClearProperties()
	return cu
}

// SetOwnerID sets the "owner" edge to the Account entity by ID.
func (cu *ClusterUpdate) SetOwnerID(id xid.ID) *ClusterUpdate {
	cu.mutation.SetOwnerID(id)
	return cu
}

// SetOwner sets the "owner" edge to the Account entity.
func (cu *ClusterUpdate) SetOwner(a *Account) *ClusterUpdate {
	return cu.SetOwnerID(a.ID)
}

// SetParentID sets the "parent" edge to the Cluster entity by ID.
func (cu *ClusterUpdate) SetParentID(id xid.ID) *ClusterUpdate {
	cu.mutation.SetParentID(id)
	return cu
}

// SetNillableParentID sets the "parent" edge to the Cluster entity by ID if the given value is not nil.
func (cu *ClusterUpdate) SetNillableParentID(id *xid.ID) *ClusterUpdate {
	if id != nil {
		cu = cu.SetParentID(*id)
	}
	return cu
}

// SetParent sets the "parent" edge to the Cluster entity.
func (cu *ClusterUpdate) SetParent(c *Cluster) *ClusterUpdate {
	return cu.SetParentID(c.ID)
}

// AddClusterIDs adds the "clusters" edge to the Cluster entity by IDs.
func (cu *ClusterUpdate) AddClusterIDs(ids ...xid.ID) *ClusterUpdate {
	cu.mutation.AddClusterIDs(ids...)
	return cu
}

// AddClusters adds the "clusters" edges to the Cluster entity.
func (cu *ClusterUpdate) AddClusters(c ...*Cluster) *ClusterUpdate {
	ids := make([]xid.ID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddClusterIDs(ids...)
}

// AddItemIDs adds the "items" edge to the Item entity by IDs.
func (cu *ClusterUpdate) AddItemIDs(ids ...xid.ID) *ClusterUpdate {
	cu.mutation.AddItemIDs(ids...)
	return cu
}

// AddItems adds the "items" edges to the Item entity.
func (cu *ClusterUpdate) AddItems(i ...*Item) *ClusterUpdate {
	ids := make([]xid.ID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cu.AddItemIDs(ids...)
}

// AddAssetIDs adds the "assets" edge to the Asset entity by IDs.
func (cu *ClusterUpdate) AddAssetIDs(ids ...string) *ClusterUpdate {
	cu.mutation.AddAssetIDs(ids...)
	return cu
}

// AddAssets adds the "assets" edges to the Asset entity.
func (cu *ClusterUpdate) AddAssets(a ...*Asset) *ClusterUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cu.AddAssetIDs(ids...)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (cu *ClusterUpdate) AddTagIDs(ids ...xid.ID) *ClusterUpdate {
	cu.mutation.AddTagIDs(ids...)
	return cu
}

// AddTags adds the "tags" edges to the Tag entity.
func (cu *ClusterUpdate) AddTags(t ...*Tag) *ClusterUpdate {
	ids := make([]xid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.AddTagIDs(ids...)
}

// Mutation returns the ClusterMutation object of the builder.
func (cu *ClusterUpdate) Mutation() *ClusterMutation {
	return cu.mutation
}

// ClearOwner clears the "owner" edge to the Account entity.
func (cu *ClusterUpdate) ClearOwner() *ClusterUpdate {
	cu.mutation.ClearOwner()
	return cu
}

// ClearParent clears the "parent" edge to the Cluster entity.
func (cu *ClusterUpdate) ClearParent() *ClusterUpdate {
	cu.mutation.ClearParent()
	return cu
}

// ClearClusters clears all "clusters" edges to the Cluster entity.
func (cu *ClusterUpdate) ClearClusters() *ClusterUpdate {
	cu.mutation.ClearClusters()
	return cu
}

// RemoveClusterIDs removes the "clusters" edge to Cluster entities by IDs.
func (cu *ClusterUpdate) RemoveClusterIDs(ids ...xid.ID) *ClusterUpdate {
	cu.mutation.RemoveClusterIDs(ids...)
	return cu
}

// RemoveClusters removes "clusters" edges to Cluster entities.
func (cu *ClusterUpdate) RemoveClusters(c ...*Cluster) *ClusterUpdate {
	ids := make([]xid.ID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveClusterIDs(ids...)
}

// ClearItems clears all "items" edges to the Item entity.
func (cu *ClusterUpdate) ClearItems() *ClusterUpdate {
	cu.mutation.ClearItems()
	return cu
}

// RemoveItemIDs removes the "items" edge to Item entities by IDs.
func (cu *ClusterUpdate) RemoveItemIDs(ids ...xid.ID) *ClusterUpdate {
	cu.mutation.RemoveItemIDs(ids...)
	return cu
}

// RemoveItems removes "items" edges to Item entities.
func (cu *ClusterUpdate) RemoveItems(i ...*Item) *ClusterUpdate {
	ids := make([]xid.ID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cu.RemoveItemIDs(ids...)
}

// ClearAssets clears all "assets" edges to the Asset entity.
func (cu *ClusterUpdate) ClearAssets() *ClusterUpdate {
	cu.mutation.ClearAssets()
	return cu
}

// RemoveAssetIDs removes the "assets" edge to Asset entities by IDs.
func (cu *ClusterUpdate) RemoveAssetIDs(ids ...string) *ClusterUpdate {
	cu.mutation.RemoveAssetIDs(ids...)
	return cu
}

// RemoveAssets removes "assets" edges to Asset entities.
func (cu *ClusterUpdate) RemoveAssets(a ...*Asset) *ClusterUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cu.RemoveAssetIDs(ids...)
}

// ClearTags clears all "tags" edges to the Tag entity.
func (cu *ClusterUpdate) ClearTags() *ClusterUpdate {
	cu.mutation.ClearTags()
	return cu
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (cu *ClusterUpdate) RemoveTagIDs(ids ...xid.ID) *ClusterUpdate {
	cu.mutation.RemoveTagIDs(ids...)
	return cu
}

// RemoveTags removes "tags" edges to Tag entities.
func (cu *ClusterUpdate) RemoveTags(t ...*Tag) *ClusterUpdate {
	ids := make([]xid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.RemoveTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ClusterUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks[int, ClusterMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ClusterUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ClusterUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ClusterUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ClusterUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := cluster.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ClusterUpdate) check() error {
	if _, ok := cu.mutation.OwnerID(); cu.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Cluster.owner"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *ClusterUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ClusterUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *ClusterUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(cluster.Table, cluster.Columns, sqlgraph.NewFieldSpec(cluster.FieldID, field.TypeString))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(cluster.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.DeletedAt(); ok {
		_spec.SetField(cluster.FieldDeletedAt, field.TypeTime, value)
	}
	if cu.mutation.DeletedAtCleared() {
		_spec.ClearField(cluster.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(cluster.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Slug(); ok {
		_spec.SetField(cluster.FieldSlug, field.TypeString, value)
	}
	if value, ok := cu.mutation.ImageURL(); ok {
		_spec.SetField(cluster.FieldImageURL, field.TypeString, value)
	}
	if cu.mutation.ImageURLCleared() {
		_spec.ClearField(cluster.FieldImageURL, field.TypeString)
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(cluster.FieldDescription, field.TypeString, value)
	}
	if value, ok := cu.mutation.Properties(); ok {
		_spec.SetField(cluster.FieldProperties, field.TypeJSON, value)
	}
	if cu.mutation.PropertiesCleared() {
		_spec.ClearField(cluster.FieldProperties, field.TypeJSON)
	}
	if cu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cluster.OwnerTable,
			Columns: []string{cluster.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cluster.OwnerTable,
			Columns: []string{cluster.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cluster.ParentTable,
			Columns: []string{cluster.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cluster.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cluster.ParentTable,
			Columns: []string{cluster.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cluster.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ClustersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cluster.ClustersTable,
			Columns: []string{cluster.ClustersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cluster.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedClustersIDs(); len(nodes) > 0 && !cu.mutation.ClustersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cluster.ClustersTable,
			Columns: []string{cluster.ClustersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cluster.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ClustersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cluster.ClustersTable,
			Columns: []string{cluster.ClustersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cluster.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cluster.ItemsTable,
			Columns: cluster.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedItemsIDs(); len(nodes) > 0 && !cu.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cluster.ItemsTable,
			Columns: cluster.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cluster.ItemsTable,
			Columns: cluster.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.AssetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cluster.AssetsTable,
			Columns: cluster.AssetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedAssetsIDs(); len(nodes) > 0 && !cu.mutation.AssetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cluster.AssetsTable,
			Columns: cluster.AssetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.AssetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cluster.AssetsTable,
			Columns: cluster.AssetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   cluster.TagsTable,
			Columns: cluster.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedTagsIDs(); len(nodes) > 0 && !cu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   cluster.TagsTable,
			Columns: cluster.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   cluster.TagsTable,
			Columns: cluster.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(cu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cluster.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ClusterUpdateOne is the builder for updating a single Cluster entity.
type ClusterUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ClusterMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *ClusterUpdateOne) SetUpdatedAt(t time.Time) *ClusterUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cuo *ClusterUpdateOne) SetDeletedAt(t time.Time) *ClusterUpdateOne {
	cuo.mutation.SetDeletedAt(t)
	return cuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cuo *ClusterUpdateOne) SetNillableDeletedAt(t *time.Time) *ClusterUpdateOne {
	if t != nil {
		cuo.SetDeletedAt(*t)
	}
	return cuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cuo *ClusterUpdateOne) ClearDeletedAt() *ClusterUpdateOne {
	cuo.mutation.ClearDeletedAt()
	return cuo
}

// SetName sets the "name" field.
func (cuo *ClusterUpdateOne) SetName(s string) *ClusterUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetSlug sets the "slug" field.
func (cuo *ClusterUpdateOne) SetSlug(s string) *ClusterUpdateOne {
	cuo.mutation.SetSlug(s)
	return cuo
}

// SetImageURL sets the "image_url" field.
func (cuo *ClusterUpdateOne) SetImageURL(s string) *ClusterUpdateOne {
	cuo.mutation.SetImageURL(s)
	return cuo
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (cuo *ClusterUpdateOne) SetNillableImageURL(s *string) *ClusterUpdateOne {
	if s != nil {
		cuo.SetImageURL(*s)
	}
	return cuo
}

// ClearImageURL clears the value of the "image_url" field.
func (cuo *ClusterUpdateOne) ClearImageURL() *ClusterUpdateOne {
	cuo.mutation.ClearImageURL()
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *ClusterUpdateOne) SetDescription(s string) *ClusterUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetParentClusterID sets the "parent_cluster_id" field.
func (cuo *ClusterUpdateOne) SetParentClusterID(x xid.ID) *ClusterUpdateOne {
	cuo.mutation.SetParentClusterID(x)
	return cuo
}

// SetNillableParentClusterID sets the "parent_cluster_id" field if the given value is not nil.
func (cuo *ClusterUpdateOne) SetNillableParentClusterID(x *xid.ID) *ClusterUpdateOne {
	if x != nil {
		cuo.SetParentClusterID(*x)
	}
	return cuo
}

// ClearParentClusterID clears the value of the "parent_cluster_id" field.
func (cuo *ClusterUpdateOne) ClearParentClusterID() *ClusterUpdateOne {
	cuo.mutation.ClearParentClusterID()
	return cuo
}

// SetAccountID sets the "account_id" field.
func (cuo *ClusterUpdateOne) SetAccountID(x xid.ID) *ClusterUpdateOne {
	cuo.mutation.SetAccountID(x)
	return cuo
}

// SetProperties sets the "properties" field.
func (cuo *ClusterUpdateOne) SetProperties(a any) *ClusterUpdateOne {
	cuo.mutation.SetProperties(a)
	return cuo
}

// ClearProperties clears the value of the "properties" field.
func (cuo *ClusterUpdateOne) ClearProperties() *ClusterUpdateOne {
	cuo.mutation.ClearProperties()
	return cuo
}

// SetOwnerID sets the "owner" edge to the Account entity by ID.
func (cuo *ClusterUpdateOne) SetOwnerID(id xid.ID) *ClusterUpdateOne {
	cuo.mutation.SetOwnerID(id)
	return cuo
}

// SetOwner sets the "owner" edge to the Account entity.
func (cuo *ClusterUpdateOne) SetOwner(a *Account) *ClusterUpdateOne {
	return cuo.SetOwnerID(a.ID)
}

// SetParentID sets the "parent" edge to the Cluster entity by ID.
func (cuo *ClusterUpdateOne) SetParentID(id xid.ID) *ClusterUpdateOne {
	cuo.mutation.SetParentID(id)
	return cuo
}

// SetNillableParentID sets the "parent" edge to the Cluster entity by ID if the given value is not nil.
func (cuo *ClusterUpdateOne) SetNillableParentID(id *xid.ID) *ClusterUpdateOne {
	if id != nil {
		cuo = cuo.SetParentID(*id)
	}
	return cuo
}

// SetParent sets the "parent" edge to the Cluster entity.
func (cuo *ClusterUpdateOne) SetParent(c *Cluster) *ClusterUpdateOne {
	return cuo.SetParentID(c.ID)
}

// AddClusterIDs adds the "clusters" edge to the Cluster entity by IDs.
func (cuo *ClusterUpdateOne) AddClusterIDs(ids ...xid.ID) *ClusterUpdateOne {
	cuo.mutation.AddClusterIDs(ids...)
	return cuo
}

// AddClusters adds the "clusters" edges to the Cluster entity.
func (cuo *ClusterUpdateOne) AddClusters(c ...*Cluster) *ClusterUpdateOne {
	ids := make([]xid.ID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddClusterIDs(ids...)
}

// AddItemIDs adds the "items" edge to the Item entity by IDs.
func (cuo *ClusterUpdateOne) AddItemIDs(ids ...xid.ID) *ClusterUpdateOne {
	cuo.mutation.AddItemIDs(ids...)
	return cuo
}

// AddItems adds the "items" edges to the Item entity.
func (cuo *ClusterUpdateOne) AddItems(i ...*Item) *ClusterUpdateOne {
	ids := make([]xid.ID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cuo.AddItemIDs(ids...)
}

// AddAssetIDs adds the "assets" edge to the Asset entity by IDs.
func (cuo *ClusterUpdateOne) AddAssetIDs(ids ...string) *ClusterUpdateOne {
	cuo.mutation.AddAssetIDs(ids...)
	return cuo
}

// AddAssets adds the "assets" edges to the Asset entity.
func (cuo *ClusterUpdateOne) AddAssets(a ...*Asset) *ClusterUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cuo.AddAssetIDs(ids...)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (cuo *ClusterUpdateOne) AddTagIDs(ids ...xid.ID) *ClusterUpdateOne {
	cuo.mutation.AddTagIDs(ids...)
	return cuo
}

// AddTags adds the "tags" edges to the Tag entity.
func (cuo *ClusterUpdateOne) AddTags(t ...*Tag) *ClusterUpdateOne {
	ids := make([]xid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.AddTagIDs(ids...)
}

// Mutation returns the ClusterMutation object of the builder.
func (cuo *ClusterUpdateOne) Mutation() *ClusterMutation {
	return cuo.mutation
}

// ClearOwner clears the "owner" edge to the Account entity.
func (cuo *ClusterUpdateOne) ClearOwner() *ClusterUpdateOne {
	cuo.mutation.ClearOwner()
	return cuo
}

// ClearParent clears the "parent" edge to the Cluster entity.
func (cuo *ClusterUpdateOne) ClearParent() *ClusterUpdateOne {
	cuo.mutation.ClearParent()
	return cuo
}

// ClearClusters clears all "clusters" edges to the Cluster entity.
func (cuo *ClusterUpdateOne) ClearClusters() *ClusterUpdateOne {
	cuo.mutation.ClearClusters()
	return cuo
}

// RemoveClusterIDs removes the "clusters" edge to Cluster entities by IDs.
func (cuo *ClusterUpdateOne) RemoveClusterIDs(ids ...xid.ID) *ClusterUpdateOne {
	cuo.mutation.RemoveClusterIDs(ids...)
	return cuo
}

// RemoveClusters removes "clusters" edges to Cluster entities.
func (cuo *ClusterUpdateOne) RemoveClusters(c ...*Cluster) *ClusterUpdateOne {
	ids := make([]xid.ID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveClusterIDs(ids...)
}

// ClearItems clears all "items" edges to the Item entity.
func (cuo *ClusterUpdateOne) ClearItems() *ClusterUpdateOne {
	cuo.mutation.ClearItems()
	return cuo
}

// RemoveItemIDs removes the "items" edge to Item entities by IDs.
func (cuo *ClusterUpdateOne) RemoveItemIDs(ids ...xid.ID) *ClusterUpdateOne {
	cuo.mutation.RemoveItemIDs(ids...)
	return cuo
}

// RemoveItems removes "items" edges to Item entities.
func (cuo *ClusterUpdateOne) RemoveItems(i ...*Item) *ClusterUpdateOne {
	ids := make([]xid.ID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cuo.RemoveItemIDs(ids...)
}

// ClearAssets clears all "assets" edges to the Asset entity.
func (cuo *ClusterUpdateOne) ClearAssets() *ClusterUpdateOne {
	cuo.mutation.ClearAssets()
	return cuo
}

// RemoveAssetIDs removes the "assets" edge to Asset entities by IDs.
func (cuo *ClusterUpdateOne) RemoveAssetIDs(ids ...string) *ClusterUpdateOne {
	cuo.mutation.RemoveAssetIDs(ids...)
	return cuo
}

// RemoveAssets removes "assets" edges to Asset entities.
func (cuo *ClusterUpdateOne) RemoveAssets(a ...*Asset) *ClusterUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cuo.RemoveAssetIDs(ids...)
}

// ClearTags clears all "tags" edges to the Tag entity.
func (cuo *ClusterUpdateOne) ClearTags() *ClusterUpdateOne {
	cuo.mutation.ClearTags()
	return cuo
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (cuo *ClusterUpdateOne) RemoveTagIDs(ids ...xid.ID) *ClusterUpdateOne {
	cuo.mutation.RemoveTagIDs(ids...)
	return cuo
}

// RemoveTags removes "tags" edges to Tag entities.
func (cuo *ClusterUpdateOne) RemoveTags(t ...*Tag) *ClusterUpdateOne {
	ids := make([]xid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.RemoveTagIDs(ids...)
}

// Where appends a list predicates to the ClusterUpdate builder.
func (cuo *ClusterUpdateOne) Where(ps ...predicate.Cluster) *ClusterUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ClusterUpdateOne) Select(field string, fields ...string) *ClusterUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Cluster entity.
func (cuo *ClusterUpdateOne) Save(ctx context.Context) (*Cluster, error) {
	cuo.defaults()
	return withHooks[*Cluster, ClusterMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ClusterUpdateOne) SaveX(ctx context.Context) *Cluster {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ClusterUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ClusterUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ClusterUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := cluster.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ClusterUpdateOne) check() error {
	if _, ok := cuo.mutation.OwnerID(); cuo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Cluster.owner"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *ClusterUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ClusterUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *ClusterUpdateOne) sqlSave(ctx context.Context) (_node *Cluster, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(cluster.Table, cluster.Columns, sqlgraph.NewFieldSpec(cluster.FieldID, field.TypeString))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Cluster.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cluster.FieldID)
		for _, f := range fields {
			if !cluster.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cluster.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(cluster.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.DeletedAt(); ok {
		_spec.SetField(cluster.FieldDeletedAt, field.TypeTime, value)
	}
	if cuo.mutation.DeletedAtCleared() {
		_spec.ClearField(cluster.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(cluster.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Slug(); ok {
		_spec.SetField(cluster.FieldSlug, field.TypeString, value)
	}
	if value, ok := cuo.mutation.ImageURL(); ok {
		_spec.SetField(cluster.FieldImageURL, field.TypeString, value)
	}
	if cuo.mutation.ImageURLCleared() {
		_spec.ClearField(cluster.FieldImageURL, field.TypeString)
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(cluster.FieldDescription, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Properties(); ok {
		_spec.SetField(cluster.FieldProperties, field.TypeJSON, value)
	}
	if cuo.mutation.PropertiesCleared() {
		_spec.ClearField(cluster.FieldProperties, field.TypeJSON)
	}
	if cuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cluster.OwnerTable,
			Columns: []string{cluster.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cluster.OwnerTable,
			Columns: []string{cluster.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cluster.ParentTable,
			Columns: []string{cluster.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cluster.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cluster.ParentTable,
			Columns: []string{cluster.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cluster.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ClustersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cluster.ClustersTable,
			Columns: []string{cluster.ClustersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cluster.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedClustersIDs(); len(nodes) > 0 && !cuo.mutation.ClustersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cluster.ClustersTable,
			Columns: []string{cluster.ClustersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cluster.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ClustersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cluster.ClustersTable,
			Columns: []string{cluster.ClustersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cluster.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cluster.ItemsTable,
			Columns: cluster.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedItemsIDs(); len(nodes) > 0 && !cuo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cluster.ItemsTable,
			Columns: cluster.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cluster.ItemsTable,
			Columns: cluster.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.AssetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cluster.AssetsTable,
			Columns: cluster.AssetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedAssetsIDs(); len(nodes) > 0 && !cuo.mutation.AssetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cluster.AssetsTable,
			Columns: cluster.AssetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.AssetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cluster.AssetsTable,
			Columns: cluster.AssetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   cluster.TagsTable,
			Columns: cluster.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !cuo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   cluster.TagsTable,
			Columns: cluster.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   cluster.TagsTable,
			Columns: cluster.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(cuo.modifiers...)
	_node = &Cluster{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cluster.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
