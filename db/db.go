package db

import (
	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgresDB struct {
	connStr string
	conn    *gorm.DB
	dialect string
}

//NewPostgresDB NewPostgresDB
func NewPostgresDB(conn string, dialect string) *PostgresDB {
	return &PostgresDB{
		connStr: conn,
		dialect: dialect,
	}
}

//EstablishConnection EstablishConnection
func (pdb *PostgresDB) EstablishConnection() *gorm.DB {
	var err error
	pdb.conn, err = gorm.Open(pdb.dialect, pdb.connStr)

	if err != nil {
		panic(err)
	}
	return pdb.conn
}

//Migrate Migrate
func (pdb *PostgresDB) Migrate(models []interface{}) {

	for _, model := range models {
		if !pdb.conn.HasTable(model) {
			pdb.conn.AutoMigrate(model)
		}

	}
}