package service

import (
	tech "github.com/LittleMikle/avito_tech_2023"
	"github.com/LittleMikle/avito_tech_2023/pkg/repository"
)

type SegmentationService struct {
	repo repository.Segmentation
}

func NewSegmenationService(repo repository.Segmentation) *SegmentationService {
	return &SegmentationService{repo: repo}
}

func (s *SegmentationService) CreateSegment(segment tech.Segment) (int, error) {
	return s.repo.CreateSegment(segment)
}

func (s *SegmentationService) DeleteSegment(id int) error {
	return s.repo.DeleteSegment(id)
}
