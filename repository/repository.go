package repository

import (
	"database/sql"
	db "github.com/SawitProRecruitment/UserService/db"
)

type UserRepository struct {
	DB *sql.DB
}

type NewUserRepositoryOptions struct {
	Dsn string
}

func NewUserRepository(opts NewUserRepositoryOptions) UserRepository {
	dbCon := db.OpenDBConnection(opts.Dsn)
	return UserRepository{
		DB: dbCon,
	}
}
