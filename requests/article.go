package requests

type CreateArticleRequest struct {
	Title    string `json:"title" binding:"required,min=20,max=200"`
	Content  string `json:"content" binding:"required,min=200"`
	Category string `json:"category" binding:"required,min=3,max=100"`
	Status   string `json:"status" binding:"required,oneof=Publish Draft Trash"`
}

type ModifyArticleRequest struct {
	Title    *string `json:"title" binding:"omitempty,min=20,max=200"`
	Content  *string `json:"content" binding:"omitempty,min=200"`
	Category *string `json:"category" binding:"omitempty,min=3,max=100"`
	Status   *string `json:"status" binding:"omitempty,oneof=Publish Draft Trash"`
}
