package task_model

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type UpdateTaskRequest struct {
	ID          string `json:"id" binding:"required,uuid"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type GetTaskIDRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}
