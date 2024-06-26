package validators

type NewTodoBody struct {
	Title string `json:"title" binding:"required"`
}
