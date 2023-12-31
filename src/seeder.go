package src

import (
	"github.com/maulanar/kms/app"
	"github.com/maulanar/kms/src/jenispengetahuan"
)

func Seeder() *seederUtil {
	if seeder == nil {
		seeder = &seederUtil{}
		seeder.Configure()
		if app.APP_ENV == "local" || app.IS_MAIN_SERVER {
			seeder.Run()
		}
		seeder.isConfigured = true
	}
	return seeder
}

var seeder *seederUtil

type seederUtil struct {
	isConfigured bool
}

func (s *seederUtil) Configure() {
	app.DB().RegisterSeeder("main", "jenis_pengetahuan", jenispengetahuan.SeederData)
}

func (s *seederUtil) Run() {
	tx, err := app.DB().Conn("main")
	if err != nil {
		app.Logger().Fatal().Err(err).Send()
	} else {
		err = app.DB().RunSeeder(tx, "main", app.Setting{})
	}
	if err != nil {
		app.Logger().Fatal().Err(err).Send()
	}
}
