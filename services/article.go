package services

import (
	"context"
	"fmt"

	"sharingvisionbe/ent"
	"sharingvisionbe/ent/post"
	"sharingvisionbe/requests"
)

func CreateArticle(ctx context.Context, db *ent.Client, object *requests.CreateArticleRequest) (*ent.Post, error) {
	builder := db.Post.
		Create().
		SetTitle(object.Title).
		SetContent(object.Content).
		SetCategory(object.Category)

	switch object.Status {
	case "Publish":
		builder.SetStatus(post.StatusPublish)
	case "Draft":
		builder.SetStatus(post.StatusDraft)
	case "Trash":
		builder.SetStatus(post.StatusTrash)
	default:
		return nil, fmt.Errorf("invalid status")
	}

	return builder.Save(ctx)
}

func GetArticles(
	ctx context.Context,
	db *ent.Client,
	limit int,
	offset int,
) ([]*ent.Post, int, error) {

	total, err := db.Post.
		Query().
		Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	articles, err := db.Post.
		Query().
		Limit(limit).
		Offset(offset).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}

func GetArticleById(
	ctx context.Context,
	db *ent.Client,
	id int,
) (*ent.Post, error) {
	article, err := db.Post.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func ModifyArticleById(
	ctx context.Context,
	db *ent.Client,
	id int,
	object *requests.ModifyArticleRequest,
) (*ent.Post, error) {

	builder := db.Post.UpdateOneID(id)

	if object.Title != nil {
		builder.SetTitle(*object.Title)
	}

	if object.Content != nil {
		builder.SetContent(*object.Content)
	}

	if object.Category != nil {
		builder.SetCategory(*object.Category)
	}

	if object.Status != nil {
		switch *object.Status {
		case "Publish":
			builder.SetStatus(post.StatusPublish)
		case "Draft":
			builder.SetStatus(post.StatusDraft)
		case "Trash":
			builder.SetStatus(post.StatusTrash)
		default:
			return nil, fmt.Errorf("invalid status")
		}
	}

	return builder.Save(ctx)
}
