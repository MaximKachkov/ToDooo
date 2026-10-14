package statistics_repository_postgres

import core_postgres_pool "github.com/MaximKachkov/ToDooo/internal/core/repository/postgres/pool"

type StatisticsRepositry struct {
	pool core_postgres_pool.Pool
}

func NewTasksRepository(pool core_postgres_pool.Pool) *StatisticsRepositry {
	return &StatisticsRepositry{
		pool: pool,
	}
}
