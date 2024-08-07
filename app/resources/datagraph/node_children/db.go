package node_children

import (
	"context"
	"fmt"

	"github.com/Southclaws/dt"
	"github.com/Southclaws/fault"
	"github.com/Southclaws/fault/fctx"
	"github.com/rs/xid"

	"github.com/Southclaws/storyden/app/resources/datagraph"
	"github.com/Southclaws/storyden/app/resources/datagraph/node"
	"github.com/Southclaws/storyden/internal/ent"
	node_model "github.com/Southclaws/storyden/internal/ent/node"
)

type database struct {
	db *ent.Client
	nr node.Repository
}

func New(db *ent.Client, nr node.Repository) Repository {
	return &database{db, nr}
}

func (d *database) Move(ctx context.Context, fromSlug datagraph.NodeSlug, toSlug datagraph.NodeSlug) (*datagraph.Node, error) {
	fromNode, err := d.nr.Get(ctx, fromSlug)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	toNode, err := d.nr.Get(ctx, toSlug)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	tx, err := d.db.Tx(ctx)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	nodes, err := d.db.Node.Query().
		Select(node_model.FieldID).
		Where(node_model.ParentNodeID(xid.ID(fromNode.ID))).
		All(ctx)
	if err != nil {
		return nil, fault.Wrap(err)
	}
	childNodeIDs := dt.Map(nodes, func(c *ent.Node) xid.ID { return c.ID })

	err = d.db.Node.Update().
		SetParentID(xid.ID(toNode.ID)).
		Where(node_model.IDIn(childNodeIDs...)).
		Exec(ctx)
	if err != nil {
		terr := tx.Rollback()
		if terr != nil {
			panic(fmt.Errorf("while handling error: %w, rollback error: %s", err, terr))
		}

		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	if err := tx.Commit(); err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return toNode, nil
}
