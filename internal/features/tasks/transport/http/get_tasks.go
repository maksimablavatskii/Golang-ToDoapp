package tasks_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/maksimablavatskii/Golang-ToDoap/internal/core/logger"
	core_http_request "github.com/maksimablavatskii/Golang-ToDoap/internal/core/transport/http/request"
	core_http_response "github.com/maksimablavatskii/Golang-ToDoap/internal/core/transport/http/response"
)

type GetTasksResponse []TaskDTOResponse

// GetTasks 	godoc
// @Summary 	Список задач
// @Description Просмотр списка задач с опциональной пагинацией и/или фильтрацией по ID автора задачи
// @Tags 		tasks
// @Produce 	json
// @Param 		user_id query int false "Фильтрация задач по ID автора"
// @Param 		limit query int false "Размер страницы с задачами"
// @Param 		offset query int false "Смещение страницы с задачами"
// @Success 	200 {object} GetTasksResponse "Список задач"
// @Failure 	400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 	500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		/tasks [get]
func (h *TasksHTTPHandler) GetTasks(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userId, limit, offset, err := getUserIdLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get 'userId'/'limit'/'offset' query param")

		return
	}

	tasksDomains, err := h.tasksService.GetTasks(ctx, userId, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get tasks")

		return
	}

	response := GetTasksResponse(tasksDTOFromdomains(tasksDomains))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getUserIdLimitOffsetQueryParams(r *http.Request) (*int, *int, *int, error) {
	const (
		userIdQueryParamKey = "user_id"
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)
	userId, err := core_http_request.GetIntQueryParam(r, userIdQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get userId query param: %w", err)
	}

	limit, err := core_http_request.GetIntQueryParam(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get limit query param: %w", err)
	}

	offset, err := core_http_request.GetIntQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get offset query param: %w", err)
	}

	return userId, limit, offset, nil
}
