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

func (h *TaskHandler) CreateTask(c echo.Context) error {
	userID := c.Get("user_id").(string)

	var input task.CreateTaskInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := h.uc.CreateTask(userID, input); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	userID := c.Get("user_id").(string)
	taskID := c.Param("id")

	var input task.UpdateTaskInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	input.ID = taskID

	if err := h.uc.UpdateTask(userID, input); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
