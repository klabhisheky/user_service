package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/klabhisheky/user_service/model"
)

type userStore struct {
	database *sql.DB
}

func New(db *sql.DB) *userStore {
	return &userStore{database: db}
}

func (store *userStore) Find(ctx context.Context, userId int) (*model.User, error) {
	q := `SELECT user_id, name, address, phone FROM users WHERE user_id = ?`
	//process the quesry for the perticular
	if userId < 0 {
		return nil, errors.New("invalid userid")
	}
	rows, err := store.database.QueryContext(ctx, q, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usr := model.User{}

	for rows.Next() {
		err := rows.Scan(&usr.UserId, &usr.Name, &usr.Address, &usr.Phone)
		if err != nil {
			return nil, err
		}
	}

	return &usr, nil
}

func (store *userStore) Create(ctx context.Context, model *model.User) (*model.User, error) {

	q := `INSERT into USERS VALUES (null, ?, ?, ?)`
	_, err := store.database.Exec(q, model.Name, model.Address, model.Phone)
	if err != nil {
		return nil, err
	}
	//fetch userId for the newly created record
	usrId, err := store.GetLastInsertedRowID(ctx)
	if err != nil {
		return nil, err
	}

	//fetch the entire newly created row
	newUser, err := store.Find(ctx, usrId)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (store *userStore) GetLastInsertedRowID(ctx context.Context) (int, error) {
	q := `SELECT LAST_INSERT_ID()`
	rows, err := store.database.QueryContext(ctx, q)
	if err != nil {
		return 0, err //TODO: create constant NULL
	}
	defer rows.Close()

	var id int

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return 0, err //TODO: create constant NULL
		}
	}
	return id, nil
}
