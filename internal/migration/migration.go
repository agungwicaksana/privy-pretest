package migration

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type migration struct {
	db      *sqlx.DB
	schemas []string
	seeders []Seeder
}

type Seeder struct {
	Query,
	Table string
}

func NewMigration(db *sqlx.DB) *migration {
	if db == nil {
		panic("db is nil")
	}
	schemas := make([]string, 0)
	seeders := make([]Seeder, 0)
	return &migration{db, schemas, seeders}
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

func (m *migration) RegisterSeeder(seeder ...Seeder) *migration {
	m.seeders = append(m.seeders, seeder...)
	return m
}

func (m *migration) Seed() *migration {
	for i := 0; i < len(m.schemas); i++ {
		data := []interface{}{}
		err := m.db.Select(&data, fmt.Sprintf("SELECT id FROM %s", m.seeders[i].Table))
		if err != nil {
			panic(err)
		}
		if len(data) < 1 {
			_, err := m.db.Exec(m.seeders[i].Query)
			if err != nil {
				panic(err)
			}
		}
	}
	return m
}
