package src

import (
	"github.com/jinzhu/gorm"
	"github.com/loopcontext/go-graphql-orm/test/gen"
	"gopkg.in/gormigrate.v1"
)

// GetMigrations migrations
func GetMigrations(db *gen.DB) []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "INIT",
			Migrate: func(tx *gorm.DB) error {
				return db.AutoMigrate()
			},
			Rollback: func(tx *gorm.DB) error {
				// there's not much we can do if initialization/automigration fails
				return nil
			},
		},
		{
			ID: "01create_user_blah",
			Migrate: func(tx *gorm.DB) error {
				type User struct {
					Blah string
				}
				return tx.AutoMigrate(&User{}).Error
			},
		},
		{
			ID: "02drop_user_blah",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}
				type User struct {
					blah string
				}
				return tx.Model(User{blah: "blah"}).DropColumn("blah").Error
			},
		},
	}
}
