package service

import (
	tech "github.com/LittleMikle/avito_tech_2023"
	"github.com/LittleMikle/avito_tech_2023/pkg/repository"
	"github.com/rs/zerolog/log"
	"math/rand"
)

type UsersSegService struct {
	repo repository.UsersSeg
}

func NewUsersSegService(repo repository.UsersSeg) *UsersSegService {
	return &UsersSegService{repo: repo}
}
func (s *UsersSegService) CreateUsersSeg(userId int, segment tech.Segment) error {
	return s.repo.CreateUsersSeg(userId, segment)
}

func (s *UsersSegService) DeleteUsersSeg(userId int, segment tech.Segment) error {
	return s.repo.DeleteUsersSeg(userId, segment)
}

func (s *UsersSegService) GetUserSeg(userId int) ([]tech.USegments, error) {
	return s.repo.GetUserSeg(userId)
}

func (s *UsersSegService) GetHistory(userId int) error {
	return s.repo.GetHistory(userId)
}

func (s *UsersSegService) ScheduleDelete(userId, days int, segment tech.Segment) error {
	return s.repo.ScheduleDelete(userId, days, segment)
}

func (s *UsersSegService) RandomSegments(segment tech.Segment, percent float64) error {
	var countRows int
	countRows, err := s.repo.CountRows()
	if err != nil {
		return err
	}

	randNums := 0.01 * percent * float64(countRows)
	randNum := int(randNums)
	randMap := make(map[int]struct{}, randNum)
	for len(randMap) < randNum {
		val := rand.Intn(countRows)
		if val == 0 {
			continue
		}
		if _, ok := randMap[val]; ok {
			continue
		}
		randMap[val] = struct{}{}
		err := s.repo.RandomSegments(segment, val)
		if err != nil {
			log.Error().Err(err).Msg("failed with RandomSegments: ")
		}
	}
	return nil
}

func (s *UsersSegService) CountRows() (int, error) {
	return s.repo.CountRows()
}
