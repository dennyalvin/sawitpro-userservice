package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/SawitProRecruitment/UserService/db"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/model"
	"time"
)

func (repository UserRepository) Create(ctx context.Context, user model.User) (*model.User, error) {
	var lastInsertedId int

	//Prepare the values argument to insert
	values := []any{
		user.FullName,
		user.Password,
		user.Phone,
		time.Now().UTC(), //created_at
		0,                //login_success counter
	}

	//Auto Generate Sql Argument like $1,$2,$3,etc..
	sqlArgs := db.GeneratePsqlArgument(values)

	//Concatenate the Insert Query and the Sql Argument
	query := fmt.Sprintf(`INSERT INTO users(full_name, password, phone, created_at, login_success ) VALUES (%s) RETURNING id`, sqlArgs)

	//Execute the Insert Query, and Scan the returned inserted ID
	err := repository.DB.QueryRowContext(ctx, query, values...).Scan(&lastInsertedId)
	if err != nil {
		return nil, err
	}

	user.Id = lastInsertedId

	return &user, nil
}

func (repository UserRepository) Update(ctx context.Context, userId int, payload *generated.ProfileUpdateParams) error {
	var setQuery string
	var values []any
	var columnUpdate map[string]any

	if payload.FullName != nil {
		columnUpdate = map[string]any{"full_name": payload.FullName}
	}

	if payload.Phone != nil {
		columnUpdate = map[string]any{"full_name": payload.FullName}
	}

	setQuery, values = db.GenerateSqlUpdateAndArgument(columnUpdate)

	query := fmt.Sprintf(`UPDATE users SET %s WHERE id=$%d`, setQuery, len(values)+1)
	values = append(values, userId)

	_, err := repository.DB.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}

func (repository UserRepository) UpdateLoginSuccess(ctx context.Context, userId int) error {

	query := `UPDATE users SET login_success = login_success + 1 WHERE id=$1`

	_, err := repository.DB.ExecContext(ctx, query, userId)
	if err != nil {
		return err
	}

	return nil
}

func (repository UserRepository) FindBy(ctx context.Context, column string, value any) (*model.User, error) {
	var user model.User

	//Prepare the query
	query := fmt.Sprintf(`SELECT * FROM users WHERE deleted_at IS NULL AND %s = $1`, column)

	//execute query and scan the result to model
	err := repository.DB.QueryRowContext(ctx, query, value).Scan(&user.Id, &user.FullName,
		&user.Password, &user.Phone, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
		&user.LoginSuccess)

	if err != nil {
		//return no error when no result found
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil

}
