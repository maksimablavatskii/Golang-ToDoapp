package statistics_postgres_repository

import core_postrgres_pool "github.com/maksimablavatskii/Golang-ToDoap/internal/core/repository/postgres/pool"

type StatisticsRepository struct {
	pool core_postrgres_pool.Pool
}

func NewStatisticsRepository(pool core_postrgres_pool.Pool) *StatisticsRepository{
	return &StatisticsRepository{
		pool: pool,
	}
}