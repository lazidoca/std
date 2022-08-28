package category

import (
	"github.com/rs/xid"

	"github.com/Southclaws/storyden/internal/infrastructure/db/model"
)

type CategoryID xid.ID

func (i CategoryID) String() string { return xid.ID(i).String() }

type PostMeta struct {
	Author string
	PostID xid.ID
	Slug   string
	Title  string
	Short  string
}

type Category struct {
	ID          CategoryID
	Name        string
	Description string
	Colour      string
	Sort        int
	Admin       bool
	Recent      []PostMeta
	PostCount   int
}

func PostMetaFromModel(p *model.Post) *PostMeta {
	slug := p.Slug

	title := p.Title

	return &PostMeta{
		Author: p.Edges.Author.Name,
		PostID: p.ID,
		Slug:   slug,
		Title:  title,
		Short:  p.Short,
	}
}

func FromModel(c *model.Category) *Category {
	recent := []PostMeta{}
	for _, p := range c.Edges.Posts {
		recent = append(recent, *PostMetaFromModel(p))
	}

	return &Category{
		ID:          CategoryID(c.ID),
		Name:        c.Name,
		Description: c.Description,
		Colour:      c.Colour,
		Sort:        c.Sort,
		Admin:       c.Admin,
		Recent:      recent,
	}
}