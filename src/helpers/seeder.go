package helpers

import (
	"markisak/src/configs"
	"markisak/src/seeders"

	gorm_seeder "github.com/kachit/gorm-seeder"
)

func Seeder() {
	//Build seeders stack
	usersSeeder := seeders.NewUsersSeeder(gorm_seeder.SeederConfiguration{Rows: 10})
	seedersStack := gorm_seeder.NewSeedersStack(configs.DB)
	seedersStack.AddSeeder(&usersSeeder)

	//Apply seed
	seedersStack.Seed()
}
