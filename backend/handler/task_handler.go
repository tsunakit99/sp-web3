package handler

import (
	"net/http"

	"github.com/tsunakit99/sp-web3/usecase/task"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	uc task.TaskUsecase
}

func NewTaskHandler(uc task.TaskUsecase) *TaskHandler {
	return &TaskHandler{uc}
}

func (h *TaskHandler) GetTasks(c echo.Context) error {
	userID := c.Get("user_id").(string)
	tasks, err := h.uc.GetTasks(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tasks)
}