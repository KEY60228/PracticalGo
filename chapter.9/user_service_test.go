package main

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_FetchUser(t *testing.T) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", TestHost, TestPort, TestUser, TestDB, TestPassword)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	sqlBytes, err := ioutil.ReadFile("./create_users.sql")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := db.ExecContext(context.TODO(), string(sqlBytes)); err != nil {
		t.Fatal(err)
	}

	tests := map[string]struct {
		userID       string
		inputTestSQL string
		want         *User
		hasErr       bool
	}{
		"1件取得": {
			userID:       "0003",
			inputTestSQL: "./testdata/input_user_1.sql",
			want:         &User{UserID: "0003", UserName: "gopher01"},
			hasErr:       false,
		},
		"0件取得": {
			userID:       "9999",
			inputTestSQL: "./testdata/input_user_2.sql",
			want:         nil,
			hasErr:       true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			sqlBytes, err := ioutil.ReadFile(tt.inputTestSQL)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := db.ExecContext(context.TODO(), string(sqlBytes)); err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() {
				if _, err := db.ExecContext(context.TODO(), `TRUNCATE users;`); err != nil {
					t.Fatal(err)
				}
			})

			s := NewUserService(db)
			got, err := s.FetchUser(context.TODO(), tt.userID)
			if (err != nil) != tt.hasErr {
				t.Fatalf("FetchUser() error = %v, hasError %v", err, tt.hasErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_FetchUser_Mock(t *testing.T) {
	type mock struct {
		db      *sql.DB
		sqlmock sqlmock.Sqlmock
	}

	tests := map[string]struct {
		userID string
		mock   mock
		want   *User
		hasErr bool
	}{
		"1件取得": {
			userID: "0001",
			mock: func() mock {
				db, m, err := sqlmock.New()
				if err != nil {
					t.Fatal(err)
				}
				m.ExpectQuery(regexp.QuoteMeta(`SELECT user_id, user_name FROM users WHERE user_id = $1;`)).
					WithArgs("0001").
					WillReturnRows(sqlmock.NewRows([]string{"user_id", "user_name"}).AddRow("0003", "gopher01"))
				return mock{db, m}
			}(),
			want:   &User{UserID: "0003", UserName: "gopher01"},
			hasErr: false,
		},
		"0件取得": {
			userID: "9999",
			mock: func() mock {
				db, m, err := sqlmock.New()
				if err != nil {
					t.Fatal(err)
				}
				m.ExpectQuery(regexp.QuoteMeta(`SELECT user_id, user_name FROM users WHERE user_id = $1;`)).
					WithArgs("9999").
					WillReturnError(sql.ErrNoRows)
				return mock{db, m}
			}(),
			want:   nil,
			hasErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			db := tt.mock.db
			s := NewUserService(db)
			got, err := s.FetchUser(context.TODO(), tt.userID)
			if (err != nil) != tt.hasErr {
				t.Fatalf("FetchUser() error = %v, hasErr %v", err, tt.hasErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
