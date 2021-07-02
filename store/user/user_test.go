package user

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestStoreFind(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error '%s' encounterd when oppening mock database", err)
	}
	defer db.Close()

	//create usrStore, tests will be called through these objects
	usrstr := New(db)

	//args to pass the func under test, Find(ctx, userId)
	type args struct {
		ctx   context.Context
		usrId int
	}

	//multiple test cases
	tests := []struct {
		name             string
		funcFindArgs     args
		shouldThrowError bool
	}{
		//testcases might not be valid for production scenario, but they might give insights on how sqlmock behave
		{name: "Test#1:Valid UserID fetch", funcFindArgs: args{context.Background(), 350}, shouldThrowError: false},
		{name: "Test#2:Invalid UserId fetch", funcFindArgs: args{context.Background(), 230}, shouldThrowError: true},
		{name: "Test#3:Invalid Query Statement", funcFindArgs: args{context.Background(), 230}, shouldThrowError: true},
	}

	for _, tc := range tests {
		q := `SELECT user_id, name, address, phone FROM users WHERE user_id = ?`
		t.Run(tc.name, func(t *testing.T) {

			var rows *sqlmock.Rows
			switch tc.name {
			case "Test#2:Invalid UserId fetch":
				//test with invalid userID type : bool
				rows = mock.NewRows([]string{"user_id", "name", "address", "phone"}).AddRow(true, "Sameer Khanna", "vasant vihar, delhi, IN", "8767869887")
				mock.ExpectQuery(q).WithArgs(tc.funcFindArgs.usrId).WillReturnRows(rows)
			case "Test#3:Invalid Query Statement":
				q = `SELECT user_id, name, address, phone FROM users WHERE user_id_wrong_column_name = ?`
				mock.ExpectQuery(q).WithArgs(tc.funcFindArgs.usrId).WillReturnError(errors.New("invalid query"))
			default:
				rows = mock.NewRows([]string{"user_id", "name", "address", "phone"}).AddRow(355, "Sameer Khanna", "vasant vihar, delhi, IN", "8767869887")
				mock.ExpectQuery(q).WithArgs(tc.funcFindArgs.usrId).WillReturnRows(rows)
			}

			_, err = usrstr.Find(tc.funcFindArgs.ctx, tc.funcFindArgs.usrId)

			//if there is an error but we don't expecct it, than log the error for FAIL
			//vice-versa, if there is not an error but we expect one, then also log and FAIL
			if (err == nil) == tc.shouldThrowError {
				t.Errorf("\n\n Error expected : %v, but Error enountered : %v \n Error: %v", tc.shouldThrowError, bool(err != nil), err)
			}
		})
	}
}
