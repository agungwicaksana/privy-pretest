package migration

import (
	"github.com/jmoiron/sqlx"
)

type migration struct {
	db      *sqlx.DB
	schemas []string
}

func NewMigration(db *sqlx.DB) *migration {
	if db == nil {
		panic("db is nil")
	}
	schemas := make([]string, 0)
	return &migration{db, schemas}
}

func (m *migration) RegisterSchemas(regSchemas ...string) *migration {
	m.schemas = append(m.schemas, regSchemas...)
	return m
}

func (m *migration) Migrate() *migration {
	for i := 0; i < len(m.schemas); i++ {
		_, err := m.db.Exec(m.schemas[i])
		if err != nil {
			panic(err)
		}
	}
	return m
}
