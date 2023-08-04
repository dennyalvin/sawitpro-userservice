package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strings"
	"time"
)

func OpenDBConnection(dsn string) *sql.DB {
	if dsn == "" {
		panic("Please set ENV DATABASE_URL")
	}

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

// GeneratePsqlArgument - Generate something like $1,$2,$3 etc...
func GeneratePsqlArgument(params []any) string {

	var prep []string
	for i, _ := range params {
		prep = append(prep, fmt.Sprintf("$%d", i+1))
	}

	return fmt.Sprintf(strings.Join(prep, ","))
}

// GenerateSqlUpdateAndArgument - Auto Generate Sql SQL SET statement and argument values
func GenerateSqlUpdateAndArgument(columns map[string]any) (string, []any) {
	var setSql []string
	var values []any

	i := 1
	for col, val := range columns {
		values = append(values, val)
		setSql = append(setSql, fmt.Sprintf("%s=$%d", col, i))
	}

	return strings.Join(setSql, " , "), values
}
