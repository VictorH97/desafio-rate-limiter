package database

import "github.com/VictorH97/devfullcycle/goexpert/desafio-rate-limiter/internal/entity"

type LimiterInfoRepositoryInterface interface {
	Save(r *entity.LimiterInfo) error
	GetByIP(ip string) (*entity.LimiterInfo, error)
	Update(r *entity.LimiterInfo) error
}
