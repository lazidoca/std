package post

import (
	"context"

	"4d63.com/optional"

	"github.com/Southclaws/storyden/api/src/resources/user"
)

type Repository interface {
	CreatePost(
		ctx context.Context,
		body string,
		authorID user.UserID,
		parentID PostID,
		replyToID optional.Optional[PostID],
	) (*Post, error)

	// EditPost(ctx context.Context, authorID, postID string, title *string, body *string) (*Post, error)
	// DeletePost(ctx context.Context, authorID, postID string, force bool) (*Post, error)
}
