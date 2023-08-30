package repository

import (
	"fmt"
	tech "github.com/LittleMikle/avito_tech_2023"
	"github.com/jmoiron/sqlx"
	"github.com/joho/sqltocsv"
	"strconv"
	"time"
)

type UsersSegPostgres struct {
	db *sqlx.DB
}

func NewUsersSegPostgres(db *sqlx.DB) *UsersSegPostgres {
	return &UsersSegPostgres{
		db: db,
	}
}

func (r *UsersSegPostgres) CreateUsersSeg(userId int, segment tech.Segment) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	getIdQuery := fmt.Sprintf("SELECT id FROM %s WHERE title =$1", segmentsTable)
	err = r.db.Get(&segment, getIdQuery, segment.Title)

	createUsersSegQuery := fmt.Sprintf("INSERT INTO %s (user_id, segment_id) SELECT %d, %d WHERE NOT EXISTS (SELECT id FROM user_segment WHERE user_id = $1 AND segment_id = $2)", userSegmentsTable, userId, segment.Id)
	fmt.Println(createUsersSegQuery)
	_, err = r.db.Exec(createUsersSegQuery, userId, segment.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	operationType := "ADD"
	historyQuery := fmt.Sprintf("INSERT INTO %s (user_id, segment_id, operation_type) VALUES ($1, $2, $3)", userHistoryTable)
	_, err = r.db.Exec(historyQuery, userId, segment.Id, operationType)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *UsersSegPostgres) GetUserSeg(userId int) ([]tech.USegments, error) {
	var segments []tech.USegments

	query := fmt.Sprintf("SELECT us.id, us.user_id, us.segment_id FROM %s us INNER JOIN %s s ON us.segment_id = s.id WHERE user_id =$1", userSegmentsTable, segmentsTable)
	if err := r.db.Select(&segments, query, userId); err != nil {
		return nil, err
	}
	fmt.Println(segments)
	return segments, nil
}

func (r *UsersSegPostgres) GetHistory(userId int) error {
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", userHistoryTable)

	rows, _ := r.db.Query(query, userId)
	strNum := strconv.Itoa(userId)
	err := sqltocsv.WriteFile("resultUserId"+strNum+".csv", rows)
	if err != nil {
		return err
	}

	return err
}

func (r *UsersSegPostgres) ScheduleDelete(userId, days int, segment tech.Segment) error {
	getIdQuery := fmt.Sprintf("SELECT id FROM %s WHERE title =$1", segmentsTable)
	err := r.db.Get(&segment, getIdQuery, segment.Title)

	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 AND segment_id=$2", userSegmentsTable)
	_, err = r.db.Exec(deleteQuery, userId, segment.Id)

	operationType := "REMOVED"
	operationTime := time.Now()
	historyQuery := fmt.Sprintf("INSERT INTO %s (user_id, segment_id, operation_type, execution_time) VALUES ($1, $2, $3, ($4::timestamptz))", userHistoryTable)
	_, err = r.db.Exec(historyQuery, userId, segment.Id, operationType, operationTime)

	return err
}

func (r *UsersSegPostgres) CountRows() (int, error) {
	var countRows int
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s", usersTable)
	err := r.db.Get(&countRows, query)
	if err != nil {
		return 0, err
	}
	return countRows, err
}

func (r *UsersSegPostgres) RandomSegments(segment tech.Segment, val int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	getIdQuery := fmt.Sprintf("SELECT id FROM %s WHERE title =$1", segmentsTable)
	err = r.db.Get(&segment, getIdQuery, segment.Title)

	createUsersSegQuery := fmt.Sprintf("INSERT INTO %s (user_id, segment_id) SELECT %d, %d WHERE NOT EXISTS (SELECT id FROM user_segment WHERE user_id = $1 AND segment_id = $2)", userSegmentsTable, val, segment.Id)
	_, err = r.db.Exec(createUsersSegQuery, val, segment.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	operationType := "ADD"
	historyQuery := fmt.Sprintf("INSERT INTO %s (user_id, segment_id, operation_type) VALUES ($1, $2, $3)", userHistoryTable)
	_, err = r.db.Exec(historyQuery, val, segment.Id, operationType)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
