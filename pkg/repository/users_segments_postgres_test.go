package repository

import (
	tech "github.com/LittleMikle/avito_tech_2023"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func TestUsersSegPostgres_CreateUsersSeg(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal().Err(err).Msg("failed with sqlmock.Newx")
	}
	defer db.Close()

	r := NewUsersSegPostgres(db)

	type args struct {
		userId  int
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
				userId: 1,
				segment: tech.Segment{
					Title: "Avito_Melushev",
				},
			},
			id: 1,
			mockBehavior: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("SELECT id FROM segments").
					WithArgs(args.segment.Id).WillReturnRows(rows)

				mock.ExpectQuery("INSERT INTO user_segment").
					WithArgs(args.userId, args.segment.Id).WillReturnRows(rows)

				mock.ExpectExec("INSERT INTO user_history").
					WithArgs(args.userId, args.segment.Id, "ADD")

				mock.ExpectCommit()
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args, testCase.id)

			err := r.CreateUsersSeg(testCase.args.userId, testCase.args.segment)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUsersSegPostgres_DeleteUsersSeg(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal().Err(err).Msg("failed with sqlmock.Newx")
	}
	defer db.Close()

	r := NewUsersSegPostgres(db)

	type args struct {
		userId  int
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
				userId: 1,
				segment: tech.Segment{
					Title: "Avito_Melushev",
				},
			},
			id: 1,
			mockBehavior: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("SELECT id FROM segments").
					WithArgs(args.segment.Id).WillReturnRows(rows)

				mock.ExpectQuery("DELETE FROM user_segment").
					WithArgs(args.userId, args.segment.Id)

				mock.ExpectExec("INSERT INTO user_history").
					WithArgs(args.userId, args.segment.Id, "REMOVE")

				mock.ExpectCommit()
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args, testCase.id)

			err := r.CreateUsersSeg(testCase.args.userId, testCase.args.segment)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
