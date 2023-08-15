package rdb

import (
	migratepkg "github.com/rubenv/sql-migrate"
)

func (dbh *DBHandler) Migrate(dir string) error {
	db, err := dbh.Conn.DB()
	if err != nil {
		return err
	}

	migrate := migratepkg.FileMigrationSource{
		Dir: dir,
	}

	_, err = migratepkg.Exec(db, "mysql", migrate, migratepkg.Up)
	return err
}
