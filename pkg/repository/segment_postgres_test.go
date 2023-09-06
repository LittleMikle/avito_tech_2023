package repository

import (
	"errors"
	tech "github.com/LittleMikle/avito_tech_2023"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func TestSegmentationPostgres_CreateSegment(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal().Err(err).Msg("failed with sqlmock.Newx")
	}
	defer db.Close()

	r := NewSegmentationPostgres(db)

	type args struct {
		segment tech.Segment
	}

	type mockBehavior func(args args, id int)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		id           int
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				segment: tech.Segment{
					Title: "Avito_Melushev",
				},
			},
			id: 1,
			mockBehavior: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO segments").
					WithArgs(args.segment.Title).WillReturnRows(rows)

				mock.ExpectCommit()
			},
		},
		{
			name: "Empty Field",
			args: args{
				segment: tech.Segment{
					Title: "",
				},
			},
			mockBehavior: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id).RowError(1, errors.New("some error"))
				mock.ExpectQuery("INSERT INTO segments").
					WithArgs(args.segment.Title).WillReturnRows(rows)

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args, testCase.id)

			got, err := r.CreateSegment(testCase.args.segment)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.id, got)
			}
		})
	}
}

func TestSegmentationPostgres_DeleteSegment(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal().Err(err).Msg("failed with sqlmock.Newx")
	}
	defer db.Close()

	r := NewSegmentationPostgres(db)

	type args struct {
		id int
	}

	type mockBehavior func(args args)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				id: 1,
			},
			mockBehavior: func(args args) {

				mock.ExpectExec("DELETE FROM segments WHERE (.+)").
					WithArgs(args.id)

			},
			wantErr: true,
		},
		{
			name: "BAD QUERY",
			args: args{
				id: 1,
			},
			mockBehavior: func(args args) {

				mock.ExpectExec("INSERT INTO segments ").
					WithArgs(args.id)

			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args)

			err := r.DeleteSegment(testCase.args.id)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
