package src

import (
	"github.com/maulanar/kms/app"
	"github.com/maulanar/kms/src/accesstoken"
	"github.com/maulanar/kms/src/akademi"
	"github.com/maulanar/kms/src/attachment"
	"github.com/maulanar/kms/src/jenispengetahuan"
	"github.com/maulanar/kms/src/kompetensi"
	"github.com/maulanar/kms/src/lingkuppengetahuan"
	"github.com/maulanar/kms/src/orang"
	"github.com/maulanar/kms/src/pengetahuan"
	"github.com/maulanar/kms/src/referensi"
	"github.com/maulanar/kms/src/statuspengetahuan"
	"github.com/maulanar/kms/src/subjenispengetahuan"
	"github.com/maulanar/kms/src/tag"
	"github.com/maulanar/kms/src/tpengetahuanrelation"
	"github.com/maulanar/kms/src/user"
	// import : DONT REMOVE THIS COMMENT
)

func Migrator() *migratorUtil {
	if migrator == nil {
		migrator = &migratorUtil{}
		migrator.Configure()
		if app.APP_ENV == "local" || app.IS_MAIN_SERVER {
			migrator.Run()
		}
		migrator.isConfigured = true
	}
	return migrator
}

var migrator *migratorUtil

type migratorUtil struct {
	isConfigured bool
}

func (*migratorUtil) Configure() {
	app.DB().RegisterTable("main", user.User{})
	app.DB().RegisterTable("main", orang.Orang{})
	app.DB().RegisterTable("main", akademi.Akademi{})
	app.DB().RegisterTable("main", jenispengetahuan.JenisPengetahuan{})
	app.DB().RegisterTable("main", kompetensi.Kompetensi{})
	app.DB().RegisterTable("main", lingkuppengetahuan.LingkupPengetahuan{})
	app.DB().RegisterTable("main", statuspengetahuan.StatusPengetahuan{})
	app.DB().RegisterTable("main", subjenispengetahuan.SubjenisPengetahuan{})
	app.DB().RegisterTable("main", tag.Tag{})
	app.DB().RegisterTable("main", pengetahuan.Pengetahuan{})
	app.DB().RegisterTable("main", tpengetahuanrelation.TPengetahuanAkademi{})
	app.DB().RegisterTable("main", tpengetahuanrelation.TPengetahuanKapitalisasi{})
	app.DB().RegisterTable("main", tpengetahuanrelation.TPengetahuanPenulisExternal{})
	app.DB().RegisterTable("main", tpengetahuanrelation.TPengetahuanTag{})
	app.DB().RegisterTable("main", tpengetahuanrelation.TPengetahuanTugas{})
	app.DB().RegisterTable("main", tpengetahuanrelation.TPengetahuanReferensi{})
	app.DB().RegisterTable("main", tpengetahuanrelation.TPengetahuanDokumen{})
	app.DB().RegisterTable("main", referensi.Referensi{})
	app.DB().RegisterTable("main", accesstoken.AccessToken{})
	app.DB().RegisterTable("main", attachment.Attachment{})
	// RegisterTable : DONT REMOVE THIS COMMENT
}

func (*migratorUtil) Run() {
	tx, err := app.DB().Conn("main")
	if err != nil {
		app.Logger().Fatal().Err(err).Send()
	} else {
		err = app.DB().MigrateTable(tx, "main", app.Setting{})
	}
	if err != nil {
		app.Logger().Fatal().Err(err).Send()
	}
}
