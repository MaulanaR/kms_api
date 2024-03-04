package src

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanar/kms/app"
	"github.com/maulanar/kms/src/accesstoken"
	"github.com/maulanar/kms/src/advisanalytic"
	"github.com/maulanar/kms/src/adviskategori"
	"github.com/maulanar/kms/src/advislistdata"
	"github.com/maulanar/kms/src/advissubkategori"
	"github.com/maulanar/kms/src/advissumberdata"
	"github.com/maulanar/kms/src/akademi"
	"github.com/maulanar/kms/src/attachment"
	"github.com/maulanar/kms/src/bannercarousel"
	"github.com/maulanar/kms/src/dislike"
	"github.com/maulanar/kms/src/dokumen"
	"github.com/maulanar/kms/src/dokumenmap"
	"github.com/maulanar/kms/src/elibrary"
	"github.com/maulanar/kms/src/event"
	"github.com/maulanar/kms/src/eventmateri"
	"github.com/maulanar/kms/src/forum"
	"github.com/maulanar/kms/src/hadiah"
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
	"github.com/maulanar/kms/src/notifikasi"
	"github.com/maulanar/kms/src/orang"
	"github.com/maulanar/kms/src/pedoman"
	"github.com/maulanar/kms/src/pencapaian"
	"github.com/maulanar/kms/src/penerbit"
	"github.com/maulanar/kms/src/pengetahuan"
	"github.com/maulanar/kms/src/provinsi"
	"github.com/maulanar/kms/src/pulau"
	"github.com/maulanar/kms/src/referensi"
	"github.com/maulanar/kms/src/statuspengetahuan"
	"github.com/maulanar/kms/src/subjenispengetahuan"
	"github.com/maulanar/kms/src/tag"
	"github.com/maulanar/kms/src/totalsummary"
	"github.com/maulanar/kms/src/user"
	// import : DONT REMOVE THIS COMMENT
)

func Router() *routerUtil {
	if router == nil {
		router = &routerUtil{}
		router.Configure()
		router.isConfigured = true
	}
	return router
}

var router *routerUtil

type routerUtil struct {
	isConfigured bool
}

func (r *routerUtil) Configure() {
	app.Server().AddRoute("/api/v1/version", "GET", app.VersionHandler, nil)
	app.Server().AddStatic("/api/storages", "./storages", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		Index:         "index.html",
		CacheDuration: 1 * time.Hour,
		MaxAge:        3600,
	})

	app.Server().AddRoute("/api/v1/user", "POST", user.REST().Create, user.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/user", "GET", user.REST().Get, user.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/user/{id}", "GET", user.REST().GetByID, user.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/user/{id}", "PUT", user.REST().UpdateByID, user.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/user/{id}", "PATCH", user.REST().PartiallyUpdateByID, user.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/user/{id}", "DELETE", user.REST().DeleteByID, user.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/orang", "POST", orang.REST().Create, orang.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/orang", "GET", orang.REST().Get, orang.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/orang/{id}", "GET", orang.REST().GetByID, orang.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/orang/{id}", "PUT", orang.REST().UpdateByID, orang.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/orang/{id}", "PATCH", orang.REST().PartiallyUpdateByID, orang.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/orang/{id}", "DELETE", orang.REST().DeleteByID, orang.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/akademi", "POST", akademi.REST().Create, akademi.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/akademi", "GET", akademi.REST().Get, akademi.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/akademi/{id}", "GET", akademi.REST().GetByID, akademi.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/akademi/{id}", "PUT", akademi.REST().UpdateByID, akademi.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/akademi/{id}", "PATCH", akademi.REST().PartiallyUpdateByID, akademi.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/akademi/{id}", "DELETE", akademi.REST().DeleteByID, akademi.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/jenis_pengetahuan", "POST", jenispengetahuan.REST().Create, jenispengetahuan.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/jenis_pengetahuan", "GET", jenispengetahuan.REST().Get, jenispengetahuan.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/jenis_pengetahuan/{id}", "GET", jenispengetahuan.REST().GetByID, jenispengetahuan.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/jenis_pengetahuan/{id}", "PUT", jenispengetahuan.REST().UpdateByID, jenispengetahuan.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/jenis_pengetahuan/{id}", "PATCH", jenispengetahuan.REST().PartiallyUpdateByID, jenispengetahuan.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/jenis_pengetahuan/{id}", "DELETE", jenispengetahuan.REST().DeleteByID, jenispengetahuan.OpenAPI().DeleteByID())

	// app.Server().AddRoute("/api/v1/kompetensi", "POST", kompetensi.REST().Create, kompetensi.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/kompetensi", "GET", kompetensi.REST().Get, kompetensi.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/kompetensi/{id}", "GET", kompetensi.REST().GetByID, kompetensi.OpenAPI().GetByID())
	// app.Server().AddRoute("/api/v1/kompetensi/{id}", "PUT", kompetensi.REST().UpdateByID, kompetensi.OpenAPI().UpdateByID())
	// app.Server().AddRoute("/api/v1/kompetensi/{id}", "PATCH", kompetensi.REST().PartiallyUpdateByID, kompetensi.OpenAPI().PartiallyUpdateByID())
	// app.Server().AddRoute("/api/v1/kompetensi/{id}", "DELETE", kompetensi.REST().DeleteByID, kompetensi.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/lingkup_pengetahuan", "POST", lingkuppengetahuan.REST().Create, lingkuppengetahuan.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/lingkup_pengetahuan", "GET", lingkuppengetahuan.REST().Get, lingkuppengetahuan.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/lingkup_pengetahuan/{id}", "GET", lingkuppengetahuan.REST().GetByID, lingkuppengetahuan.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/lingkup_pengetahuan/{id}", "PUT", lingkuppengetahuan.REST().UpdateByID, lingkuppengetahuan.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/lingkup_pengetahuan/{id}", "PATCH", lingkuppengetahuan.REST().PartiallyUpdateByID, lingkuppengetahuan.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/lingkup_pengetahuan/{id}", "DELETE", lingkuppengetahuan.REST().DeleteByID, lingkuppengetahuan.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/status_pengetahuan", "POST", statuspengetahuan.REST().Create, statuspengetahuan.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/status_pengetahuan", "GET", statuspengetahuan.REST().Get, statuspengetahuan.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/status_pengetahuan/{id}", "GET", statuspengetahuan.REST().GetByID, statuspengetahuan.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/status_pengetahuan/{id}", "PUT", statuspengetahuan.REST().UpdateByID, statuspengetahuan.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/status_pengetahuan/{id}", "PATCH", statuspengetahuan.REST().PartiallyUpdateByID, statuspengetahuan.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/status_pengetahuan/{id}", "DELETE", statuspengetahuan.REST().DeleteByID, statuspengetahuan.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/subjenis_pengetahuan", "POST", subjenispengetahuan.REST().Create, subjenispengetahuan.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/subjenis_pengetahuan", "GET", subjenispengetahuan.REST().Get, subjenispengetahuan.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/subjenis_pengetahuan/{id}", "GET", subjenispengetahuan.REST().GetByID, subjenispengetahuan.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/subjenis_pengetahuan/{id}", "PUT", subjenispengetahuan.REST().UpdateByID, subjenispengetahuan.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/subjenis_pengetahuan/{id}", "PATCH", subjenispengetahuan.REST().PartiallyUpdateByID, subjenispengetahuan.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/subjenis_pengetahuan/{id}", "DELETE", subjenispengetahuan.REST().DeleteByID, subjenispengetahuan.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/tag", "POST", tag.REST().Create, tag.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/tag", "GET", tag.REST().Get, tag.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/tag/{id}", "GET", tag.REST().GetByID, tag.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/tag/{id}", "PUT", tag.REST().UpdateByID, tag.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/tag/{id}", "PATCH", tag.REST().PartiallyUpdateByID, tag.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/tag/{id}", "DELETE", tag.REST().DeleteByID, tag.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/pengetahuan", "POST", pengetahuan.REST().Create, pengetahuan.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/pengetahuan", "GET", pengetahuan.REST().Get, pengetahuan.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/pengetahuan/{id}", "GET", pengetahuan.REST().GetByID, pengetahuan.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/pengetahuan/{id}", "PUT", pengetahuan.REST().UpdateByID, pengetahuan.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/pengetahuan/{id}", "PATCH", pengetahuan.REST().PartiallyUpdateByID, pengetahuan.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/pengetahuan/{id}", "DELETE", pengetahuan.REST().DeleteByID, pengetahuan.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/slider_pengetahuan", "GET", pengetahuan.REST().GetSlider, pengetahuan.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/mix_slider", "GET", pengetahuan.REST().GetMixSlider, pengetahuan.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/populer_pengetahuan", "GET", pengetahuan.REST().GetPopuler, pengetahuan.OpenAPI().Get())

	//like & dislike
	app.Server().AddRoute("/api/v1/pengetahuan/{id}/like", "POST", like.REST().UpdateByPengetahuanID, like.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/pengetahuan/{id}/dislike", "POST", dislike.REST().UpdateByPengetahuanID, dislike.OpenAPI().PartiallyUpdateByID())

	app.Server().AddRoute("/api/v1/referensi", "POST", referensi.REST().Create, referensi.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/referensi", "GET", referensi.REST().Get, referensi.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/referensi/{id}", "GET", referensi.REST().GetByID, referensi.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/referensi/{id}", "PUT", referensi.REST().UpdateByID, referensi.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/referensi/{id}", "PATCH", referensi.REST().PartiallyUpdateByID, referensi.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/referensi/{id}", "DELETE", referensi.REST().DeleteByID, referensi.OpenAPI().DeleteByID())

	// app.Server().AddRoute("/api/v1/narasumber", "POST", narasumber.REST().Create, narasumber.OpenAPI().Create())
	// app.Server().AddRoute("/api/v1/narasumber", "GET", narasumber.REST().Get, narasumber.OpenAPI().Get())
	// app.Server().AddRoute("/api/v1/narasumber/{id}", "GET", narasumber.REST().GetByID, narasumber.OpenAPI().GetByID())
	// app.Server().AddRoute("/api/v1/narasumber/{id}", "PUT", narasumber.REST().UpdateByID, narasumber.OpenAPI().UpdateByID())
	// app.Server().AddRoute("/api/v1/narasumber/{id}", "PATCH", narasumber.REST().PartiallyUpdateByID, narasumber.OpenAPI().PartiallyUpdateByID())
	// app.Server().AddRoute("/api/v1/narasumber/{id}", "DELETE", narasumber.REST().DeleteByID, narasumber.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/penerbit", "POST", penerbit.REST().Create, penerbit.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/penerbit", "GET", penerbit.REST().Get, penerbit.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/penerbit/{id}", "GET", penerbit.REST().GetByID, penerbit.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/penerbit/{id}", "PUT", penerbit.REST().UpdateByID, penerbit.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/penerbit/{id}", "PATCH", penerbit.REST().PartiallyUpdateByID, penerbit.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/penerbit/{id}", "DELETE", penerbit.REST().DeleteByID, penerbit.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/login", "POST", accesstoken.REST().Login, accesstoken.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/session", "GET", accesstoken.REST().GetByID, accesstoken.OpenAPI().Create())

	app.Server().AddRoute("/api/v1/attachments", "POST", attachment.REST().Create, attachment.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/attachments", "GET", attachment.REST().Get, attachment.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/attachments/{id}", "GET", attachment.REST().GetByID, attachment.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/attachments/{id}", "DELETE", attachment.REST().DeleteByID, attachment.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/caches", "DELETE", attachment.REST().ClearCaches, attachment.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/komentar", "POST", komentar.REST().Create, komentar.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/komentar", "GET", komentar.REST().Get, komentar.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/komentar/{id}", "GET", komentar.REST().GetByID, komentar.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/komentar/{id}", "PUT", komentar.REST().UpdateByID, komentar.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/komentar/{id}", "PATCH", komentar.REST().PartiallyUpdateByID, komentar.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/komentar/{id}", "DELETE", komentar.REST().DeleteByID, komentar.OpenAPI().DeleteByID())

	// app.Server().AddRoute("/api/v1/like", "POST", like.REST().Create, like.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/like", "GET", like.REST().Get, like.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/like/{id}", "GET", like.REST().GetByID, like.OpenAPI().GetByID())
	// app.Server().AddRoute("/api/v1/like/{id}", "PUT", like.REST().UpdateByID, like.OpenAPI().UpdateByID())
	// app.Server().AddRoute("/api/v1/like/{id}", "PATCH", like.REST().PartiallyUpdateByID, like.OpenAPI().PartiallyUpdateByID())
	// app.Server().AddRoute("/api/v1/like/{id}", "DELETE", like.REST().DeleteByID, like.OpenAPI().DeleteByID())

	// app.Server().AddRoute("/api/v1/dislike", "POST", dislike.REST().Create, dislike.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/dislike", "GET", dislike.REST().Get, dislike.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/dislike/{id}", "GET", dislike.REST().GetByID, dislike.OpenAPI().GetByID())
	// app.Server().AddRoute("/api/v1/dislike/{id}", "PUT", dislike.REST().UpdateByID, dislike.OpenAPI().UpdateByID())
	// app.Server().AddRoute("/api/v1/dislike/{id}", "PATCH", dislike.REST().PartiallyUpdateByID, dislike.OpenAPI().PartiallyUpdateByID())
	// app.Server().AddRoute("/api/v1/dislike/{id}", "DELETE", dislike.REST().DeleteByID, dislike.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/search_pengetahuan", "GET", pengetahuan.REST().GetSearch, pengetahuan.OpenAPI().Get())

	app.Server().AddRoute("/api/v1/events/{id}/live_komentar", "GET", event.REST().GetLiveKomentar, event.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/events/{id}/live_komentar", "POST", event.REST().CreateLiveKomentar, event.OpenAPI().Create())

	app.Server().AddRoute("/api/v1/events", "POST", event.REST().Create, event.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/events", "GET", event.REST().Get, event.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/events/{id}", "GET", event.REST().GetByID, event.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/events/{id}", "PUT", event.REST().UpdateByID, event.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/events/{id}", "PATCH", event.REST().PartiallyUpdateByID, event.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/events/{id}", "DELETE", event.REST().DeleteByID, event.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/event_materi", "POST", eventmateri.REST().Create, eventmateri.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/event_materi", "GET", eventmateri.REST().Get, eventmateri.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/event_materi/{id}", "GET", eventmateri.REST().GetByID, eventmateri.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/event_materi/{id}", "PUT", eventmateri.REST().UpdateByID, eventmateri.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/event_materi/{id}", "PATCH", eventmateri.REST().PartiallyUpdateByID, eventmateri.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/event_materi/{id}", "DELETE", eventmateri.REST().DeleteByID, eventmateri.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/forum", "POST", forum.REST().Create, forum.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/forum", "GET", forum.REST().Get, forum.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/forum/{id}", "GET", forum.REST().GetByID, forum.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/forum/{id}", "PUT", forum.REST().UpdateByID, forum.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/forum/{id}", "PATCH", forum.REST().PartiallyUpdateByID, forum.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/forum/{id}", "DELETE", forum.REST().DeleteByID, forum.OpenAPI().DeleteByID())
	app.Server().AddRoute("/api/v1/forum/{id}/like", "POST", like.REST().UpdateByforumID, like.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/forum/{id}/dislike", "POST", dislike.REST().UpdateByforumID, dislike.OpenAPI().PartiallyUpdateByID())

	app.Server().AddRoute("/api/v1/search_forum", "GET", forum.REST().GetSearch, forum.OpenAPI().Get())

	app.Server().AddRoute("/api/v1/leader_talk", "POST", leadertalk.REST().Create, leadertalk.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/leader_talk", "GET", leadertalk.REST().Get, leadertalk.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/leader_talk/{id}", "GET", leadertalk.REST().GetByID, leadertalk.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/leader_talk/{id}", "PUT", leadertalk.REST().UpdateByID, leadertalk.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/leader_talk/{id}", "PATCH", leadertalk.REST().PartiallyUpdateByID, leadertalk.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/leader_talk/{id}", "DELETE", leadertalk.REST().DeleteByID, leadertalk.OpenAPI().DeleteByID())
	app.Server().AddRoute("/api/v1/leader_talk/{id}/like", "POST", like.REST().UpdateByLeaderTalkID, like.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/leader_talk/{id}/dislike", "POST", dislike.REST().UpdateByLeaderTalkID, dislike.OpenAPI().PartiallyUpdateByID())

	app.Server().AddRoute("/api/v1/library_cafe", "POST", librarycafe.REST().Create, librarycafe.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/library_cafe", "GET", librarycafe.REST().Get, librarycafe.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/library_cafe/{id}", "GET", librarycafe.REST().GetByID, librarycafe.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/library_cafe/{id}", "PUT", librarycafe.REST().UpdateByID, librarycafe.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/library_cafe/{id}", "PATCH", librarycafe.REST().PartiallyUpdateByID, librarycafe.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/library_cafe/{id}", "DELETE", librarycafe.REST().DeleteByID, librarycafe.OpenAPI().DeleteByID())
	app.Server().AddRoute("/api/v1/library_cafe/{id}/like", "POST", like.REST().UpdateByLibraryCafeID, like.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/library_cafe/{id}/dislike", "POST", dislike.REST().UpdateByLibraryCafeID, dislike.OpenAPI().PartiallyUpdateByID())

	app.Server().AddRoute("/api/v1/pedoman", "POST", pedoman.REST().Create, pedoman.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/pedoman", "GET", pedoman.REST().Get, pedoman.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/pedoman/{id}", "GET", pedoman.REST().GetByID, pedoman.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/pedoman/{id}", "PUT", pedoman.REST().UpdateByID, pedoman.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/pedoman/{id}", "PATCH", pedoman.REST().PartiallyUpdateByID, pedoman.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/pedoman/{id}", "DELETE", pedoman.REST().DeleteByID, pedoman.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/kelompok_dokumen", "POST", kelompokdokumen.REST().Create, kelompokdokumen.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/kelompok_dokumen", "GET", kelompokdokumen.REST().Get, kelompokdokumen.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/kelompok_dokumen/{id}", "GET", kelompokdokumen.REST().GetByID, kelompokdokumen.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/kelompok_dokumen/{id}", "PUT", kelompokdokumen.REST().UpdateByID, kelompokdokumen.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/kelompok_dokumen/{id}", "PATCH", kelompokdokumen.REST().PartiallyUpdateByID, kelompokdokumen.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/kelompok_dokumen/{id}", "DELETE", kelompokdokumen.REST().DeleteByID, kelompokdokumen.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/kategori_pengetahuan", "POST", kategoripengetahuan.REST().Create, kategoripengetahuan.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/kategori_pengetahuan", "GET", kategoripengetahuan.REST().Get, kategoripengetahuan.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/kategori_pengetahuan/{id}", "GET", kategoripengetahuan.REST().GetByID, kategoripengetahuan.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/kategori_pengetahuan/{id}", "PUT", kategoripengetahuan.REST().UpdateByID, kategoripengetahuan.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/kategori_pengetahuan/{id}", "PATCH", kategoripengetahuan.REST().PartiallyUpdateByID, kategoripengetahuan.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/kategori_pengetahuan/{id}", "DELETE", kategoripengetahuan.REST().DeleteByID, kategoripengetahuan.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/dokumen", "POST", dokumen.REST().Create, dokumen.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/dokumen", "GET", dokumen.REST().Get, dokumen.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/dokumen/{id}", "GET", dokumen.REST().GetByID, dokumen.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/dokumen/{id}", "PUT", dokumen.REST().UpdateByID, dokumen.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/dokumen/{id}", "PATCH", dokumen.REST().PartiallyUpdateByID, dokumen.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/dokumen/{id}", "DELETE", dokumen.REST().DeleteByID, dokumen.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/kategori_buku", "POST", kategoribuku.REST().Create, kategoribuku.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/kategori_buku", "GET", kategoribuku.REST().Get, kategoribuku.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/kategori_buku/{id}", "GET", kategoribuku.REST().GetByID, kategoribuku.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/kategori_buku/{id}", "PUT", kategoribuku.REST().UpdateByID, kategoribuku.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/kategori_buku/{id}", "PATCH", kategoribuku.REST().PartiallyUpdateByID, kategoribuku.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/kategori_buku/{id}", "DELETE", kategoribuku.REST().DeleteByID, kategoribuku.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/elibrary", "POST", elibrary.REST().Create, elibrary.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/elibrary", "GET", elibrary.REST().Get, elibrary.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/elibrary/{id}", "GET", elibrary.REST().GetByID, elibrary.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/elibrary/{id}", "PUT", elibrary.REST().UpdateByID, elibrary.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/elibrary/{id}", "PATCH", elibrary.REST().PartiallyUpdateByID, elibrary.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/elibrary/{id}", "DELETE", elibrary.REST().DeleteByID, elibrary.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/history_points", "POST", historypoint.REST().Create, historypoint.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/history_points", "GET", historypoint.REST().Get, historypoint.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/history_points/{id}", "GET", historypoint.REST().GetByID, historypoint.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/history_points/{id}", "PUT", historypoint.REST().UpdateByID, historypoint.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/history_points/{id}", "PATCH", historypoint.REST().PartiallyUpdateByID, historypoint.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/history_points/{id}", "DELETE", historypoint.REST().DeleteByID, historypoint.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/dokumen_map", "POST", dokumenmap.REST().Create, dokumenmap.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/dokumen_map", "GET", dokumenmap.REST().Get, dokumenmap.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/dokumen_map/{id}", "GET", dokumenmap.REST().GetByID, dokumenmap.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/dokumen_map/{id}", "PUT", dokumenmap.REST().UpdateByID, dokumenmap.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/dokumen_map/{id}", "PATCH", dokumenmap.REST().PartiallyUpdateByID, dokumenmap.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/dokumen_map/{id}", "DELETE", dokumenmap.REST().DeleteByID, dokumenmap.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/total_summaries", "GET", totalsummary.REST().Get, totalsummary.OpenAPI().Get())

	app.Server().AddRoute("/api/v1/advis_analytics", "POST", advisanalytic.REST().Create, advisanalytic.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/advis_analytics", "GET", advisanalytic.REST().Get, advisanalytic.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/advis_analytics/template_csv", "GET", advisanalytic.REST().DownloadTemplateCSV, advisanalytic.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/advis_analytics/upload_csv", "POST", advisanalytic.REST().UploadCSV, advisanalytic.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/advis_analytics/{id}", "GET", advisanalytic.REST().GetByID, advisanalytic.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/advis_analytics/{id}", "PUT", advisanalytic.REST().UpdateByID, advisanalytic.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/advis_analytics/{id}", "PATCH", advisanalytic.REST().PartiallyUpdateByID, advisanalytic.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/advis_analytics/{id}", "DELETE", advisanalytic.REST().DeleteByID, advisanalytic.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/advis_list_data", "POST", advislistdata.REST().Create, advislistdata.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/advis_list_data", "GET", advislistdata.REST().Get, advislistdata.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/advis_list_data/template_csv", "GET", advislistdata.REST().DownloadTemplateCSV, advislistdata.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/advis_list_data/upload_csv", "POST", advislistdata.REST().UploadCSV, advislistdata.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/advis_list_data/{id}", "GET", advislistdata.REST().GetByID, advislistdata.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/advis_list_data/{id}", "PUT", advislistdata.REST().UpdateByID, advislistdata.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/advis_list_data/{id}", "PATCH", advislistdata.REST().PartiallyUpdateByID, advislistdata.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/advis_list_data/{id}", "DELETE", advislistdata.REST().DeleteByID, advislistdata.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/advis_kategori", "POST", adviskategori.REST().Create, adviskategori.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/advis_kategori", "GET", adviskategori.REST().Get, adviskategori.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/advis_kategori/{id}", "GET", adviskategori.REST().GetByID, adviskategori.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/advis_kategori/{id}", "PUT", adviskategori.REST().UpdateByID, adviskategori.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/advis_kategori/{id}", "PATCH", adviskategori.REST().PartiallyUpdateByID, adviskategori.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/advis_kategori/{id}", "DELETE", adviskategori.REST().DeleteByID, adviskategori.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/advis_sub_kategori", "POST", advissubkategori.REST().Create, advissubkategori.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/advis_sub_kategori", "GET", advissubkategori.REST().Get, advissubkategori.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/advis_sub_kategori/{id}", "GET", advissubkategori.REST().GetByID, advissubkategori.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/advis_sub_kategori/{id}", "PUT", advissubkategori.REST().UpdateByID, advissubkategori.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/advis_sub_kategori/{id}", "PATCH", advissubkategori.REST().PartiallyUpdateByID, advissubkategori.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/advis_sub_kategori/{id}", "DELETE", advissubkategori.REST().DeleteByID, advissubkategori.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/advis_sumber_data", "POST", advissumberdata.REST().Create, advissumberdata.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/advis_sumber_data", "GET", advissumberdata.REST().Get, advissumberdata.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/advis_sumber_data/{id}", "GET", advissumberdata.REST().GetByID, advissumberdata.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/advis_sumber_data/{id}", "PUT", advissumberdata.REST().UpdateByID, advissumberdata.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/advis_sumber_data/{id}", "PATCH", advissumberdata.REST().PartiallyUpdateByID, advissumberdata.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/advis_sumber_data/{id}", "DELETE", advissumberdata.REST().DeleteByID, advissumberdata.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/pulau", "POST", pulau.REST().Create, pulau.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/pulau", "GET", pulau.REST().Get, pulau.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/pulau/{id}", "GET", pulau.REST().GetByID, pulau.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/pulau/{id}", "PUT", pulau.REST().UpdateByID, pulau.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/pulau/{id}", "PATCH", pulau.REST().PartiallyUpdateByID, pulau.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/pulau/{id}", "DELETE", pulau.REST().DeleteByID, pulau.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/provinsi", "POST", provinsi.REST().Create, provinsi.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/provinsi", "GET", provinsi.REST().Get, provinsi.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/provinsi/{id}", "GET", provinsi.REST().GetByID, provinsi.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/provinsi/{id}", "PUT", provinsi.REST().UpdateByID, provinsi.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/provinsi/{id}", "PATCH", provinsi.REST().PartiallyUpdateByID, provinsi.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/provinsi/{id}", "DELETE", provinsi.REST().DeleteByID, provinsi.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/notifikasi", "POST", notifikasi.REST().Create, notifikasi.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/notifikasi", "GET", notifikasi.REST().Get, notifikasi.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/notifikasi/{id}", "GET", notifikasi.REST().GetByID, notifikasi.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/notifikasi/{id}", "PUT", notifikasi.REST().UpdateByID, notifikasi.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/notifikasi/{id}", "PATCH", notifikasi.REST().PartiallyUpdateByID, notifikasi.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/notifikasi/{id}", "DELETE", notifikasi.REST().DeleteByID, notifikasi.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/ranking_point", "GET", hadiah.REST().GetRanking, hadiah.OpenAPI().Get())

	app.Server().AddRoute("/api/v1/hadiah", "POST", hadiah.REST().Create, hadiah.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/hadiah", "GET", hadiah.REST().Get, hadiah.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/hadiah/{id}", "GET", hadiah.REST().GetByID, hadiah.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/hadiah/{id}", "PUT", hadiah.REST().UpdateByID, hadiah.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/hadiah/{id}", "PATCH", hadiah.REST().PartiallyUpdateByID, hadiah.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/hadiah/{id}", "DELETE", hadiah.REST().DeleteByID, hadiah.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/pencapaian", "POST", pencapaian.REST().Create, pencapaian.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/pencapaian", "GET", pencapaian.REST().Get, pencapaian.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/pencapaian/{id}", "GET", pencapaian.REST().GetByID, pencapaian.OpenAPI().GetByID())
	// app.Server().AddRoute("/api/v1/pencapaian/{id}", "PUT", pencapaian.REST().UpdateByID, pencapaian.OpenAPI().UpdateByID())
	// app.Server().AddRoute("/api/v1/pencapaian/{id}", "PATCH", pencapaian.REST().PartiallyUpdateByID, pencapaian.OpenAPI().PartiallyUpdateByID())
	// app.Server().AddRoute("/api/v1/pencapaian/{id}", "DELETE", pencapaian.REST().DeleteByID, pencapaian.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/banner_carousel", "POST", bannercarousel.REST().Create, bannercarousel.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/banner_carousel", "GET", bannercarousel.REST().Get, bannercarousel.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/banner_carousel/{id}", "GET", bannercarousel.REST().GetByID, bannercarousel.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/banner_carousel/{id}", "PUT", bannercarousel.REST().UpdateByID, bannercarousel.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/banner_carousel/{id}", "PATCH", bannercarousel.REST().PartiallyUpdateByID, bannercarousel.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/banner_carousel/{id}", "DELETE", bannercarousel.REST().DeleteByID, bannercarousel.OpenAPI().DeleteByID())

	// AddRoute : DONT REMOVE THIS COMMENT
}
