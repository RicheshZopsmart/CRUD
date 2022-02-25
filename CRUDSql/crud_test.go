package CRUDSql

import (
	"database/sql"
	"errors"
	"log"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

var emp = Empl{
	id:    1,
	Name:  "Richesh",
	Email: "Richesh@zopsmart.com",
	role:  "sde",
}

// func TestGetDetailsById(t *testing.T) {
// 	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

// 	if err != nil {
// 		t.Errorf("Error while creating SQLMock")
// 	}
// 	defer db.Close()

// 	rows := sqlmock.NewRows([]string{"id", "Name", "Email", "role"}).AddRow(1, "Richesh", "r@r.com", "SDE-I")

// 	testcases := []struct {
// 		id            int
// 		user          Empl
// 		mockQ         interface{}
// 		expectedError error
// 	}{
// 		{
// 			id:            1,
// 			user:          Empl{1, "Richesh", "r@r.com", "SDE-I"},
// 			mockQ:         mock.ExpectQuery("select * from user where id=?").WithArgs(1).WillReturnRows(rows),
// 			expectedError: nil,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run("", func(t *testing.T) {
// 			user, err := GetDetailsById(db, tc.id)

// 			if err != nil && err.Error() != tc.expectedError.Error() {
// 				t.Errorf("expected %v Got: %v", tc.expectedError, err)
// 			}

// 			if !reflect.DeepEqual(user, tc.user) {
// 				t.Errorf("expected %v Got: %v", tc.user, user)
// 			}
// 		})
// 	}

// }

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestUpdateById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf(err.Error())
	}
	defer db.Close()

	query := "update Employee_Details set Name=?, Email=?, role=? where id=?"
	kquery := "update Employee_Details1 set Name=?, Email=?, role=? where id=?"

	tc := []struct {
		id    int
		Name  string
		Email string
		role  string
		err   error
		mockQ interface{}
	}{
		{
			id:    1,
			Name:  "Richesh",
			Email: "r@r.com",
			role:  "SDE-I",
			err:   nil,
			mockQ: mock.ExpectPrepare(query).ExpectExec().WithArgs("Richesh", "r@r.com", "SDE-I", 1).WillReturnResult(sqlmock.NewResult(1, 0)),
		},
		{
			id:    2,
			Name:  "KAKA",
			Email: "k@k.com",
			role:  "SDE-I",
			err:   sql.ErrNoRows,
			mockQ: mock.ExpectPrepare(query).ExpectExec().WithArgs("KAKA", "k@k.com", "SDE-I", 2).WillReturnError(sql.ErrNoRows),
		},
		{
			id:    2,
			Name:  "KAKA",
			Email: "k@k.com",
			role:  "SDE-I",
			err:   errors.New("Prepare Error"),
			mockQ: mock.ExpectPrepare(kquery).ExpectExec().WithArgs("KAKA", "k@k.com", "SDE-I", 3).WillReturnError(errors.New("Prepare Table")),
		},
	}

	for _, tt := range tc {
		t.Run("", func(t *testing.T) {

			err := UpdateById(db, tt.id, tt.Name, tt.Email, tt.role)
			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf("expected error:%v, got:%v", tt.err, err)
			}

		})
	}

}

func TestDeleteById(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf(err.Error())
	}
	defer db.Close()

	query := "delete from Employee_Details where id=?"
	kquery := "delete from Employee_Details1 where id=?"
	// prep := mock.ExpectPrepare(query)
	// prep.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
	tc := []struct {
		id    int
		Name  string
		Email string
		role  string
		err   error
		mockQ interface{}
	}{
		{
			id:    1,
			Name:  "Richesh",
			Email: "r@r.com",
			role:  "SDE-I",
			err:   nil,
			mockQ: mock.ExpectPrepare(query).ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 0)),
		},
		{
			id:    2,
			Name:  "KAKA",
			Email: "k@k.com",
			role:  "SDE-I",
			err:   sql.ErrNoRows,
			mockQ: mock.ExpectPrepare(query).ExpectExec().WithArgs(2).WillReturnError(sql.ErrNoRows),
		},
		{
			id:    2,
			Name:  "KAKA",
			Email: "k@k.com",
			role:  "SDE-I",
			err:   errors.New("Prepare Error"),
			mockQ: mock.ExpectPrepare(kquery).ExpectExec().WithArgs("KAKA", "k@k.com", "SDE-I").WillReturnError(errors.New("Wrong Table")),
		},
	}

	for _, tt := range tc {
		t.Run("", func(t *testing.T) {

			err := DeleteById(db, tt.id)
			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf("expected error:%v, got:%v", tt.err, err)
			}

		})
	}

}

func TestInsertData(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf(err.Error())
	}
	defer db.Close()

	query := "INSERT INTO Employee_Details (Name,Email,role) values (?,?,?);"
	kquery := "INSERT INTO Employeee_Details1 (Name,Email,role) values (?,?,?);"
	tc := []struct {
		id    int
		Name  string
		Email string
		role  string
		err   error
		mockQ interface{}
	}{
		{
			id:    1,
			Name:  "Richesh",
			Email: "r@r.com",
			role:  "SDE-I",
			err:   nil,
			mockQ: mock.ExpectPrepare(query).ExpectExec().WithArgs("Richesh", "r@r.com", "SDE-I").WillReturnResult(sqlmock.NewResult(0, 1)),
		},
		{
			id:    2,
			Name:  "KAKA",
			Email: "k@k.com",
			role:  "SDE-I",
			err:   sql.ErrNoRows,
			mockQ: mock.ExpectPrepare(query).ExpectExec().WithArgs("KAKA", "k@k.com", "SDE-I").WillReturnError(sql.ErrNoRows),
		},
		{
			id:    2,
			Name:  "KAKA",
			Email: "k@k.com",
			role:  "SDE-I",
			err:   errors.New("Prepare Error"),
			mockQ: mock.ExpectPrepare(kquery).ExpectExec().WithArgs("KAKA", "k@k.com", "SDE-I").WillReturnError(errors.New("Wrong Table")),
		},
	}

	for _, tt := range tc {
		t.Run("", func(t *testing.T) {

			err := InsertData("Employee_Details", db, tt.Name, tt.Email, tt.role)
			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf("expected error:%v, got:%v", tt.err, err)
			}

		})
	}

}

func TestGetDetailsById(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Errorf(err.Error())
	}

	defer db.Close()

	query := "select * from Employee_Details where id=?"
	kquery := "select * from Employee_Details1 where id=?"
	// prep := mock.ExpectPrepare(query)
	tc := []struct {
		id          int
		emp         *Empl
		mockQuery   interface{}
		expectError error
	}{
		{
			id:          1,
			emp:         &Empl{id: 1, Name: "Richesh", Email: "r@r.com", role: "SDE-I"},
			mockQuery:   mock.ExpectPrepare(query).ExpectQuery().WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "Name", "Email", "role"}).AddRow(1, "Richesh", "r@r.com", "SDE-I")),
			expectError: nil,
		},
		{
			id:          2,
			emp:         &Empl{id: 2, Name: "Jane", Email: "j@j.com", role: "SDE-II"},
			mockQuery:   mock.ExpectPrepare(query).ExpectQuery().WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"id", "Name", "Email", "role"}).AddRow(2, "Jane", "j@j.com", "SDE-II")),
			expectError: nil,
		},
		// Failure
		{
			id:          3,
			emp:         nil,
			mockQuery:   mock.ExpectPrepare(query).ExpectQuery().WithArgs(3).WillReturnError(errors.New("Prepare Error")),
			expectError: errors.New("Prepare Error"),
		},

		{
			id:          4,
			emp:         nil,
			expectError: errors.New(""),
			mockQuery:   mock.ExpectPrepare(kquery).WillReturnError(errors.New("")),
		},
	}

	for _, tt := range tc {
		t.Run("", func(t *testing.T) {

			emp, err := GetById(db, tt.id)
			if err != nil && err.Error() != tt.expectError.Error() {
				t.Errorf("expected error:%v, got:%v", tt.expectError, err)
			} else {

				if !reflect.DeepEqual(tt.emp, emp) {
					t.Errorf("expected users %v, got: %v", tt.emp.Name, emp)
				}
			}

		})
	}
}
