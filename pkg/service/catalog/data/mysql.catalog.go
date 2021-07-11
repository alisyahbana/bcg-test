package data

import (
	"database/sql"
	"github.com/alisyahbana/bcg-test/pkg/common/database"
	"github.com/jmoiron/sqlx"
)

type MysqlCatalogData struct {
}

const (
	GetItemQuery = `SELECT * FROM catalog WHERE name = ? LIMIT 1;`
)

type MysqlCatalogStatement struct {
	GetItem *sqlx.Stmt
}

var stmt MysqlCatalogStatement

func init() {
	stmt.GetItem = database.Prepare(database.GetDBMaster(), GetItemQuery)
}

func (m MysqlCatalogData) GetItem(name string) (*Item, error) {
	var item Item
	err := stmt.GetItem.Get(
		&item,
		name,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &item, nil
}
