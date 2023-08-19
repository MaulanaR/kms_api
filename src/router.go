package src

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanar/kms/app"
	"github.com/maulanar/kms/src/accesstoken"
	"github.com/maulanar/kms/src/akademi"
	"github.com/maulanar/kms/src/attachment"
	"github.com/maulanar/kms/src/dislike"
	"github.com/maulanar/kms/src/jenispengetahuan"
	"github.com/maulanar/kms/src/komentar"
	"github.com/maulanar/kms/src/kompetensi"
	"github.com/maulanar/kms/src/like"
	"github.com/maulanar/kms/src/lingkuppengetahuan"
	"github.com/maulanar/kms/src/orang"
	"github.com/maulanar/kms/src/pengetahuan"
	"github.com/maulanar/kms/src/referensi"
	"github.com/maulanar/kms/src/statuspengetahuan"
	"github.com/maulanar/kms/src/subjenispengetahuan"
	"github.com/maulanar/kms/src/tag"
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

	app.Server().AddRoute("/api/v1/kompetensi", "POST", kompetensi.REST().Create, kompetensi.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/kompetensi", "GET", kompetensi.REST().Get, kompetensi.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/kompetensi/{id}", "GET", kompetensi.REST().GetByID, kompetensi.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/kompetensi/{id}", "PUT", kompetensi.REST().UpdateByID, kompetensi.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/kompetensi/{id}", "PATCH", kompetensi.REST().PartiallyUpdateByID, kompetensi.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/kompetensi/{id}", "DELETE", kompetensi.REST().DeleteByID, kompetensi.OpenAPI().DeleteByID())

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

	//like & dislike
	app.Server().AddRoute("/api/v1/pengetahuan/{id}/like", "POST", like.REST().UpdateByPengetahuanID, like.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/pengetahuan/{id}/dislike", "POST", dislike.REST().UpdateByPengetahuanID, dislike.OpenAPI().PartiallyUpdateByID())

	app.Server().AddRoute("/api/v1/referensi", "POST", referensi.REST().Create, referensi.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/referensi", "GET", referensi.REST().Get, referensi.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/referensi/{id}", "GET", referensi.REST().GetByID, referensi.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/referensi/{id}", "PUT", referensi.REST().UpdateByID, referensi.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/referensi/{id}", "PATCH", referensi.REST().PartiallyUpdateByID, referensi.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/referensi/{id}", "DELETE", referensi.REST().DeleteByID, referensi.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/login", "POST", accesstoken.REST().Login, accesstoken.OpenAPI().Create())

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

	app.Server().AddRoute("/api/v1/like", "POST", like.REST().Create, like.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/like", "GET", like.REST().Get, like.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/like/{id}", "GET", like.REST().GetByID, like.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/like/{id}", "PUT", like.REST().UpdateByID, like.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/like/{id}", "PATCH", like.REST().PartiallyUpdateByID, like.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/like/{id}", "DELETE", like.REST().DeleteByID, like.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/dislike", "POST", dislike.REST().Create, dislike.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/dislike", "GET", dislike.REST().Get, dislike.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/dislike/{id}", "GET", dislike.REST().GetByID, dislike.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/dislike/{id}", "PUT", dislike.REST().UpdateByID, dislike.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/dislike/{id}", "PATCH", dislike.REST().PartiallyUpdateByID, dislike.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/dislike/{id}", "DELETE", dislike.REST().DeleteByID, dislike.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/search_pengetahuan", "GET", pengetahuan.REST().GetSearch, pengetahuan.OpenAPI().Get())

	// AddRoute : DONT REMOVE THIS COMMENT
}
