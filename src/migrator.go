package src

import (
	"github.com/maulanar/kms/app"
	"github.com/maulanar/kms/src/accesstoken"
	"github.com/maulanar/kms/src/advisanalytic"
	"github.com/maulanar/kms/src/adviskategori"
	"github.com/maulanar/kms/src/advislistdata"
	"github.com/maulanar/kms/src/advissubkategori"
	"github.com/maulanar/kms/src/advissumberdata"
	"github.com/maulanar/kms/src/akademi"
	"github.com/maulanar/kms/src/attachment"
	"github.com/maulanar/kms/src/dislike"
	"github.com/maulanar/kms/src/dokumen"
	"github.com/maulanar/kms/src/dokumenmap"
	"github.com/maulanar/kms/src/elibrary"
	"github.com/maulanar/kms/src/event"
	"github.com/maulanar/kms/src/eventmateri"
	"github.com/maulanar/kms/src/forum"
	"github.com/maulanar/kms/src/historypoint"
	"github.com/maulanar/kms/src/jenispengetahuan"
	"github.com/maulanar/kms/src/kategoribuku"
	"github.com/maulanar/kms/src/kategoripengetahuan"
	"github.com/maulanar/kms/src/kelompokdokumen"
	"github.com/maulanar/kms/src/komentar"
	"github.com/maulanar/kms/src/kompetensi"
	"github.com/maulanar/kms/src/leadertalk"
	"github.com/maulanar/kms/src/librarycafe"
	"github.com/maulanar/kms/src/like"
	"github.com/maulanar/kms/src/lingkuppengetahuan"
	"github.com/maulanar/kms/src/orang"
	"github.com/maulanar/kms/src/pedoman"
	"github.com/maulanar/kms/src/penerbit"
	"github.com/maulanar/kms/src/pengetahuan"
	"github.com/maulanar/kms/src/provinsi"
	"github.com/maulanar/kms/src/pulau"
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
	// app.DB().RegisterTable("main", tpengetahuanrelation.TPengetahuanDokumen{})
	app.DB().RegisterTable("main", tpengetahuanrelation.TPengetahuanKompetensi{})
	app.DB().RegisterTable("main", tpengetahuanrelation.TPengetahuanKiat{})
	app.DB().RegisterTable("main", tpengetahuanrelation.TPengetahuanTenagaAhli{})
	app.DB().RegisterTable("main", tpengetahuanrelation.TPengetahuanPedoman{})
	app.DB().RegisterTable("main", tpengetahuanrelation.TPengetahuanResensi{})
	app.DB().RegisterTable("main", tpengetahuanrelation.TPengetahuanNarsum{})
	app.DB().RegisterTable("main", tpengetahuanrelation.TPengetahuanPenerbit{})

	app.DB().RegisterTable("main", referensi.Referensi{})
	app.DB().RegisterTable("main", accesstoken.AccessToken{})
	app.DB().RegisterTable("main", attachment.Attachment{})
	// app.DB().RegisterTable("main", narasumber.Narasumber{})
	app.DB().RegisterTable("main", penerbit.Penerbit{})
	app.DB().RegisterTable("main", komentar.Komentar{})
	app.DB().RegisterTable("main", like.Like{})
	app.DB().RegisterTable("main", dislike.Dislike{})
	app.DB().RegisterTable("main", event.Event{})
	app.DB().RegisterTable("main", event.OtherAttachment{})
	app.DB().RegisterTable("main", event.LiveComment{})
	app.DB().RegisterTable("main", eventmateri.EventMateri{})
	app.DB().RegisterTable("main", eventmateri.MateriAttachment{})
	app.DB().RegisterTable("main", forum.Forum{})
	app.DB().RegisterTable("main", leadertalk.LeaderTalk{})
	app.DB().RegisterTable("main", librarycafe.LibraryCafe{})
	app.DB().RegisterTable("main", pedoman.Pedoman{})
	app.DB().RegisterTable("main", kelompokdokumen.KelompokDokumen{})
	app.DB().RegisterTable("main", kategoripengetahuan.KategoriPengetahuan{})
	app.DB().RegisterTable("main", dokumen.Dokumen{})
	app.DB().RegisterTable("main", kategoribuku.KategoriBuku{})
	app.DB().RegisterTable("main", elibrary.Elibrary{})
	app.DB().RegisterTable("main", historypoint.HistoryPoint{})
	app.DB().RegisterTable("main", dokumenmap.DokumenMap{})
	// app.DB().RegisterTable("main", totalsummary.TotalSummary{})
	app.DB().RegisterTable("main", advisanalytic.AdvisAnalytic{})
	app.DB().RegisterTable("main", advislistdata.AdvisListData{})
	app.DB().RegisterTable("main", adviskategori.AdvisKategori{})
	app.DB().RegisterTable("main", advissubkategori.AdvisSubKategori{})
	app.DB().RegisterTable("main", advissumberdata.AdvisSumberData{})
	app.DB().RegisterTable("main", pulau.Pulau{})
	app.DB().RegisterTable("main", provinsi.Provinsi{})
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
