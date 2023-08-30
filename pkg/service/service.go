package service

import (
	tech "github.com/LittleMikle/avito_tech_2023"
	"github.com/LittleMikle/avito_tech_2023/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Segmentation interface {
	CreateSegment(segment tech.Segment) (int, error)
	DeleteSegment(id int) error
}

type UsersSeg interface {
	CreateUsersSeg(userId int, segment tech.Segment) error
	DeleteUsersSeg(userId int, segment tech.Segment) error
	GetUserSeg(userId int) ([]tech.USegments, error)
	GetHistory(userId int) error
	ScheduleDelete(userId, days int, segment tech.Segment) error
	RandomSegments(segment tech.Segment, percent float64) error
	CountRows() (int, error)
}

type Service struct {
	Segmentation
	UsersSeg
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Segmentation: NewSegmenationService(repos.Segmentation),
		UsersSeg:     NewUsersSegService(repos.UsersSeg),
	}
}
