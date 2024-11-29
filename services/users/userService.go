package services

import (
	"database/sql"
	"fmt"
	"rest-api-demo/types"
)

type Store struct {
	dbConnection *sql.DB
}

func NewUserStore(db *sql.DB) *Store {
	return &Store{
		dbConnection: db,
	}
}

func (s *Store) GetAllUsers() ([]types.UserModel, error){
	result := []types.UserModel{}
	fmt.Println("start to get all users")
	rows, err := s.dbConnection.Query("select * from User")
	if (err  != nil){
		return nil, err
	}

	defer rows.Close()

	for rows.Next(){
		var user types.UserModel
		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
            return nil, err
        }
		result = append(result, user)
	}

	if err := rows.Err(); err != nil {
        return nil, err
    }

	return result, nil
}
