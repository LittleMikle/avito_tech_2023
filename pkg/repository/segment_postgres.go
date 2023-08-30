package repository

import (
	"fmt"
	tech "github.com/LittleMikle/avito_tech_2023"
	"github.com/jmoiron/sqlx"
)

type SegmentationPostgres struct {
	db *sqlx.DB
}

func NewSegmentationPostgres(db *sqlx.DB) *SegmentationPostgres {
	return &SegmentationPostgres{
		db: db,
	}
}

func (r *SegmentationPostgres) CreateSegment(segment tech.Segment) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createSegmentQuery := fmt.Sprintf("INSERT INTO %s (title) VALUES ($1) RETURNING id", segmentsTable)
	row := tx.QueryRow(createSegmentQuery, segment.Title)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *SegmentationPostgres) DeleteSegment(id int) error {
	deleteSegmentQuery := fmt.Sprintf("DELETE FROM %s WHERE id =$1", segmentsTable)
	_, err := r.db.Exec(deleteSegmentQuery, id)

	return err
}
