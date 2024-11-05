package seeders

import (
	"fmt"
	"markisak/src/configs"
	"markisak/src/models"

	gorm_seeder "github.com/kachit/gorm-seeder"

	"gorm.io/gorm"
)

var DB = configs.DB

type UsersSeeder struct {
	gorm_seeder.SeederAbstract
}

func NewUsersSeeder(cfg gorm_seeder.SeederConfiguration) UsersSeeder {
	return UsersSeeder{gorm_seeder.NewSeederAbstract(cfg)}
}

// Implement Seed method
func (s *UsersSeeder) Seed(DB *gorm.DB) error {
	fmt.Printf("Configuration: %+v\n", s.Configuration)

	if s.Configuration.Rows <= 0 {
		return fmt.Errorf("invalid number of rows: %d", s.Configuration.Rows)
	}

	var users []models.User
	for i := 0; i < s.Configuration.Rows; i++ {
		indexStr := fmt.Sprint(i)
		user := models.User{
			Name:     "Name LastName" + indexStr,
			Email:    "foo" + indexStr + "@bar.gov",
			Password: "password-hash-string" + indexStr,
		}
		users = append(users, user)
	}

	fmt.Printf("Users to be created: %+v\n", users)

	if DB == nil {
		return fmt.Errorf("database connection is nil")
	}

	if err := DB.CreateInBatches(users, s.Configuration.Rows).Error; err != nil {
		fmt.Printf("Error creating users: %v\n", err)
		return fmt.Errorf("failed to create users: %w", err)
	}

	fmt.Printf("Successfully created %d users\n", len(users))
	return nil
}

// Implement Clear method
func (s *UsersSeeder) Clear(DB *gorm.DB) error {
	entity := models.User{}
	return s.SeederAbstract.Delete(DB, entity.TableName())
}
