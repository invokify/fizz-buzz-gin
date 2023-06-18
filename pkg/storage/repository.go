package storage

import (
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DBPath string = "fizzbuzz.db"
	create string = `
		CREATE TABLE IF NOT EXISTS stats (
			id INTEGER NOT NULL PRIMARY KEY,
			createdat TEXT NOT NULL,
			str1 TEXT,
			str2 TEXT,
			int1 INTEGER,
			int2 INTEGER,
			lim INTEGER,
			result TEXT
		);`
)

type Stats struct {
	mu sync.Mutex
	db *sql.DB
}

func NewStatsClient() (*Stats, error) {
	// create SQLite file connection
	db, err := sql.Open("sqlite3", DBPath)
	if err != nil {
		return nil, err
	}

	// Create table if not exists
	if _, err := db.Exec(create); err != nil {
		return nil, err
	}

	return &Stats{db: db}, nil
}

func (c *Stats) Insert(dto StatsDto) (int, error) {
	// convert dto to model
	model := dto.ToModel()

	// Insert stats into db
	res, err := c.db.Exec("insert into stats values (NULL, ?, ?, ?, ?, ?, ?, ?);",
		model.CreatedAt, // RFC3339
		model.Str1,
		model.Str2,
		model.Int1,
		model.Int2,
		model.Limit,
		model.Result)
	if err != nil {
		return 0, err
	}

	// get last inserted id
	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}

	return int(id), nil
}

func (c *Stats) GetAll() ([]StatsDto, error) {
	// Query DB rows based on ID
	rows, err := c.db.Query("SELECT id, createdat, str1, str2, int1, int2, lim, result FROM stats", nil)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	models := []statsModel{}
	for rows.Next() {
		// Parse row into StatsDto struct
		model := statsModel{}
		if err := rows.Scan(&model.ID, &model.CreatedAt, &model.Str1, &model.Str2, &model.Int1, &model.Int2, &model.Limit, &model.Result); err != nil {
			return nil, err
		}

		// populate stats slice
		models = append(models, model)
	}

	// convert model to dto
	dtos := []StatsDto{}
	for _, m := range models {
		dtos = append(dtos, m.ToDto())
	}

	return dtos, err
}

func (c *Stats) Close() {
	c.db.Close()
}
