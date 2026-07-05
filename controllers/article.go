package controllers

import (
	"net/http"
	"sharingvisionbe/ent"
	"sharingvisionbe/ent/post"
	"sharingvisionbe/requests"
	"sharingvisionbe/responses"
	"sharingvisionbe/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateArticle(db *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body requests.CreateArticleRequest

		if err := c.ShouldBindJSON(&body); err != nil {
			responses.JSONError(c.Writer, http.StatusBadRequest, err.Error(), nil)
			return
		}

		article, err := services.CreateArticle(
			c.Request.Context(),
			db,
			&body,
		)

		if err != nil {
			responses.JSONError(c.Writer, http.StatusBadRequest, err.Error(), nil)
			return
		}

		responses.JSONSuccess(c.Writer, article, nil, nil)
	}
}

func GetAll(db *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "2"))
		if err != nil {
			responses.JSONError(c.Writer, http.StatusBadRequest, "invalid limit", nil)
			return
		}

		offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
		if err != nil {
			responses.JSONError(c.Writer, http.StatusBadRequest, "invalid offset", nil)
			return
		}

		statusQuery := c.Query("status")

		var status *post.Status
		if statusQuery != "" {
			s := post.Status(statusQuery)

			switch s {
			case post.StatusPublish, post.StatusDraft, post.StatusTrash:
				status = &s
			default:
				responses.JSONError(c.Writer, http.StatusBadRequest, "invalid status", nil)
				return
			}
		}

		articles, total, err := services.GetArticles(
			c.Request.Context(),
			db,
			limit,
			offset,
			status,
		)
		if err != nil {
			responses.JSONError(c.Writer, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		pagination := map[string]any{
			"limit":  limit,
			"offset": offset,
			"total":  total,
		}

		meta := map[string]any{
			"count": len(articles),
		}

		responses.JSONSuccess(
			c.Writer,
			articles,
			pagination,
			meta,
		)
	}
}

func GetArticleById(db *ent.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responses.JSONError(ctx.Writer, http.StatusBadRequest, "invalid id", nil)
			return
		}

		result, err := services.GetArticleById(ctx, db, id)
		if err != nil {
			responses.JSONError(ctx.Writer, http.StatusNotFound, "not found", nil)
		}

		responses.JSONSuccess(ctx.Writer, result, nil, nil)
	}
}

func ModifyArticleById(db *ent.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responses.JSONError(ctx.Writer, http.StatusBadRequest, "invalid id", nil)
			return
		}

		var body requests.ModifyArticleRequest

		if err := ctx.ShouldBindJSON(&body); err != nil {
			responses.JSONError(ctx.Writer, http.StatusBadRequest, err.Error(), nil)
			return
		}

		article, err := services.ModifyArticleById(
			ctx.Request.Context(),
			db,
			id,
			&body,
		)

		if err != nil {
			responses.JSONError(ctx.Writer, http.StatusBadRequest, err.Error(), nil)
			return
		}

		responses.JSONSuccess(ctx.Writer, article, nil, nil)
	}
}

func DeleteArticleById(db *ent.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responses.JSONError(ctx.Writer, http.StatusBadRequest, "invalid id", nil)
			return
		}

		status := post.StatusTrash.String()

		article, err := services.ModifyArticleById(
			ctx.Request.Context(),
			db,
			id,
			&requests.ModifyArticleRequest{
				Status: &status,
			},
		)
		if err != nil {
			responses.JSONError(ctx.Writer, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		responses.JSONSuccess(ctx.Writer, article, nil, nil)
	}
}
