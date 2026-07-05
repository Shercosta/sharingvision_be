package routes

import (
	"sharingvisionbe/controllers"
	"sharingvisionbe/ent"

	"github.com/gin-gonic/gin"
)

func RouteInit(r *gin.Engine, db *ent.Client) {
	article := r.Group("/article")
	{
		article.POST("", controllers.CreateArticle(db))
		article.GET("", controllers.GetAll(db))

		articleById := article.Group("/:id")
		{
			articleById.GET("", controllers.GetArticleById(db))
			articleById.PUT("", controllers.ModifyArticleById(db))
			articleById.DELETE("", controllers.DeleteArticleById(db))
		}
	}
}
