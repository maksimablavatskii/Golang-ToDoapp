package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/maksimablavatskii/Golang-ToDoap/internal/core/domain"
	core_logger "github.com/maksimablavatskii/Golang-ToDoap/internal/core/logger"
	core_http_request "github.com/maksimablavatskii/Golang-ToDoap/internal/core/transport/http/request"
	core_http_response "github.com/maksimablavatskii/Golang-ToDoap/internal/core/transport/http/response"
)

type GetStatisticsResponse struct {
	TasksCreated               int      `json:"tasks_created" example:"50"`
	TasksCompleted             int      `json:"tasks_completed" example:"10"`
	TasksCompletedRate         *float64 `json:"tasks_completed_rate" example:"20"`
	TasksAverageCompletionTime *string  `json:"tasks_average_completion_time" example:"1m30s"`
}

// GetStatistics 	godoc
// @Summary 	Список статистики
// @Description Просмотр списка статистики с опциональной фильтрацией по user_id и/или по временому промежутку
// @Tags 		statistics
// @Produce 	json
// @Param 		user_id query int false "Фильтрация статистики по ID пользователя"
// @Param 		from query string false "Начало рассмотрети статистики (Включительно), формат: YYYY-MM-DD"
// @Param 		to query string false "Конец рассмотрения промежутка статистики(не включительно), формат: YYYY-MM-DD"
// @Success 	200 {object} GetStatisticsResponse "Список статистики"
// @Failure 	400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 	500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		/statistics [get]
func (h *StatisticsHTTPHandler) GetStatistics(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userId, from, to, err := getUserIDFromToQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userID/from/to query params")
	}

	statisctics, err := h.statisticsService.GetStatistics(ctx, userId, from, to)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get staticstics")

		return
	}

	response := toDTOFromDomain(statisctics)

	responseHandler.JSONResponse(response, http.StatusOK)

}

func toDTOFromDomain(statistics domain.Statistics) GetStatisticsResponse {
	var avgTime *string
	if statistics.TasksAverageCompletionTime != nil {
		duration := statistics.TasksAverageCompletionTime.String()
		avgTime = &duration
	}
	return GetStatisticsResponse{
		TasksCreated:               statistics.TasksCreated,
		TasksCompleted:             statistics.TasksCompleted,
		TasksCompletedRate:         statistics.TasksCompletedRate,
		TasksAverageCompletionTime: avgTime,
	}
}

func getUserIDFromToQueryParams(r *http.Request) (*int, *time.Time, *time.Time, error) {
	const (
		userIdQueryParamKey = "user_id"
		FromQueryParamKey   = "from"
		ToQueryParamKey     = "to"
	)

	userID, err := core_http_request.GetIntQueryParam(r, userIdQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'user_id' query param: %w", err)
	}

	from, err := core_http_request.GetDateQueryParam(r, FromQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'from' query param: %w", err)
	}

	to, err := core_http_request.GetDateQueryParam(r, ToQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'to' query param: %w", err)
	}

	return userID, from, to, nil
}
