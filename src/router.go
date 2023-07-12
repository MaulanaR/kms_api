package src

import (
	"github.com/maulanar/kms/app"
	"github.com/maulanar/kms/src/akademi"
	"github.com/maulanar/kms/src/jenispengetahuan"
	"github.com/maulanar/kms/src/kompetensi"
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
	app.Server().AddRoute("/api/version", "GET", app.VersionHandler, nil)

	app.Server().AddRoute("/src/user", "POST", user.REST().Create, user.OpenAPI().Create())
	app.Server().AddRoute("/src/user", "GET", user.REST().Get, user.OpenAPI().Get())
	app.Server().AddRoute("/src/user/{id}", "GET", user.REST().GetByID, user.OpenAPI().GetByID())
	app.Server().AddRoute("/src/user/{id}", "PUT", user.REST().UpdateByID, user.OpenAPI().UpdateByID())
	app.Server().AddRoute("/src/user/{id}", "PATCH", user.REST().PartiallyUpdateByID, user.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/src/user/{id}", "DELETE", user.REST().DeleteByID, user.OpenAPI().DeleteByID())

	app.Server().AddRoute("/src/orang", "POST", orang.REST().Create, orang.OpenAPI().Create())
	app.Server().AddRoute("/src/orang", "GET", orang.REST().Get, orang.OpenAPI().Get())
	app.Server().AddRoute("/src/orang/{id}", "GET", orang.REST().GetByID, orang.OpenAPI().GetByID())
	app.Server().AddRoute("/src/orang/{id}", "PUT", orang.REST().UpdateByID, orang.OpenAPI().UpdateByID())
	app.Server().AddRoute("/src/orang/{id}", "PATCH", orang.REST().PartiallyUpdateByID, orang.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/src/orang/{id}", "DELETE", orang.REST().DeleteByID, orang.OpenAPI().DeleteByID())

	app.Server().AddRoute("/src/akademi", "POST", akademi.REST().Create, akademi.OpenAPI().Create())
	app.Server().AddRoute("/src/akademi", "GET", akademi.REST().Get, akademi.OpenAPI().Get())
	app.Server().AddRoute("/src/akademi/{id}", "GET", akademi.REST().GetByID, akademi.OpenAPI().GetByID())
	app.Server().AddRoute("/src/akademi/{id}", "PUT", akademi.REST().UpdateByID, akademi.OpenAPI().UpdateByID())
	app.Server().AddRoute("/src/akademi/{id}", "PATCH", akademi.REST().PartiallyUpdateByID, akademi.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/src/akademi/{id}", "DELETE", akademi.REST().DeleteByID, akademi.OpenAPI().DeleteByID())

	app.Server().AddRoute("/src/jenis_pengetahuan", "POST", jenispengetahuan.REST().Create, jenispengetahuan.OpenAPI().Create())
	app.Server().AddRoute("/src/jenis_pengetahuan", "GET", jenispengetahuan.REST().Get, jenispengetahuan.OpenAPI().Get())
	app.Server().AddRoute("/src/jenis_pengetahuan/{id}", "GET", jenispengetahuan.REST().GetByID, jenispengetahuan.OpenAPI().GetByID())
	app.Server().AddRoute("/src/jenis_pengetahuan/{id}", "PUT", jenispengetahuan.REST().UpdateByID, jenispengetahuan.OpenAPI().UpdateByID())
	app.Server().AddRoute("/src/jenis_pengetahuan/{id}", "PATCH", jenispengetahuan.REST().PartiallyUpdateByID, jenispengetahuan.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/src/jenis_pengetahuan/{id}", "DELETE", jenispengetahuan.REST().DeleteByID, jenispengetahuan.OpenAPI().DeleteByID())

	app.Server().AddRoute("/src/kompetensi", "POST", kompetensi.REST().Create, kompetensi.OpenAPI().Create())
	app.Server().AddRoute("/src/kompetensi", "GET", kompetensi.REST().Get, kompetensi.OpenAPI().Get())
	app.Server().AddRoute("/src/kompetensi/{id}", "GET", kompetensi.REST().GetByID, kompetensi.OpenAPI().GetByID())
	app.Server().AddRoute("/src/kompetensi/{id}", "PUT", kompetensi.REST().UpdateByID, kompetensi.OpenAPI().UpdateByID())
	app.Server().AddRoute("/src/kompetensi/{id}", "PATCH", kompetensi.REST().PartiallyUpdateByID, kompetensi.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/src/kompetensi/{id}", "DELETE", kompetensi.REST().DeleteByID, kompetensi.OpenAPI().DeleteByID())

	app.Server().AddRoute("/src/lingkup_pengetahuan", "POST", lingkuppengetahuan.REST().Create, lingkuppengetahuan.OpenAPI().Create())
	app.Server().AddRoute("/src/lingkup_pengetahuan", "GET", lingkuppengetahuan.REST().Get, lingkuppengetahuan.OpenAPI().Get())
	app.Server().AddRoute("/src/lingkup_pengetahuan/{id}", "GET", lingkuppengetahuan.REST().GetByID, lingkuppengetahuan.OpenAPI().GetByID())
	app.Server().AddRoute("/src/lingkup_pengetahuan/{id}", "PUT", lingkuppengetahuan.REST().UpdateByID, lingkuppengetahuan.OpenAPI().UpdateByID())
	app.Server().AddRoute("/src/lingkup_pengetahuan/{id}", "PATCH", lingkuppengetahuan.REST().PartiallyUpdateByID, lingkuppengetahuan.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/src/lingkup_pengetahuan/{id}", "DELETE", lingkuppengetahuan.REST().DeleteByID, lingkuppengetahuan.OpenAPI().DeleteByID())

	app.Server().AddRoute("/src/status_pengetahuan", "POST", statuspengetahuan.REST().Create, statuspengetahuan.OpenAPI().Create())
	app.Server().AddRoute("/src/status_pengetahuan", "GET", statuspengetahuan.REST().Get, statuspengetahuan.OpenAPI().Get())
	app.Server().AddRoute("/src/status_pengetahuan/{id}", "GET", statuspengetahuan.REST().GetByID, statuspengetahuan.OpenAPI().GetByID())
	app.Server().AddRoute("/src/status_pengetahuan/{id}", "PUT", statuspengetahuan.REST().UpdateByID, statuspengetahuan.OpenAPI().UpdateByID())
	app.Server().AddRoute("/src/status_pengetahuan/{id}", "PATCH", statuspengetahuan.REST().PartiallyUpdateByID, statuspengetahuan.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/src/status_pengetahuan/{id}", "DELETE", statuspengetahuan.REST().DeleteByID, statuspengetahuan.OpenAPI().DeleteByID())

	app.Server().AddRoute("/src/subjenis_pengetahuan", "POST", subjenispengetahuan.REST().Create, subjenispengetahuan.OpenAPI().Create())
	app.Server().AddRoute("/src/subjenis_pengetahuan", "GET", subjenispengetahuan.REST().Get, subjenispengetahuan.OpenAPI().Get())
	app.Server().AddRoute("/src/subjenis_pengetahuan/{id}", "GET", subjenispengetahuan.REST().GetByID, subjenispengetahuan.OpenAPI().GetByID())
	app.Server().AddRoute("/src/subjenis_pengetahuan/{id}", "PUT", subjenispengetahuan.REST().UpdateByID, subjenispengetahuan.OpenAPI().UpdateByID())
	app.Server().AddRoute("/src/subjenis_pengetahuan/{id}", "PATCH", subjenispengetahuan.REST().PartiallyUpdateByID, subjenispengetahuan.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/src/subjenis_pengetahuan/{id}", "DELETE", subjenispengetahuan.REST().DeleteByID, subjenispengetahuan.OpenAPI().DeleteByID())

	app.Server().AddRoute("/src/tag", "POST", tag.REST().Create, tag.OpenAPI().Create())
	app.Server().AddRoute("/src/tag", "GET", tag.REST().Get, tag.OpenAPI().Get())
	app.Server().AddRoute("/src/tag/{id}", "GET", tag.REST().GetByID, tag.OpenAPI().GetByID())
	app.Server().AddRoute("/src/tag/{id}", "PUT", tag.REST().UpdateByID, tag.OpenAPI().UpdateByID())
	app.Server().AddRoute("/src/tag/{id}", "PATCH", tag.REST().PartiallyUpdateByID, tag.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/src/tag/{id}", "DELETE", tag.REST().DeleteByID, tag.OpenAPI().DeleteByID())

	app.Server().AddRoute("/src/pengetahuan", "POST", pengetahuan.REST().Create, pengetahuan.OpenAPI().Create())
	app.Server().AddRoute("/src/pengetahuan", "GET", pengetahuan.REST().Get, pengetahuan.OpenAPI().Get())
	app.Server().AddRoute("/src/pengetahuan/{id}", "GET", pengetahuan.REST().GetByID, pengetahuan.OpenAPI().GetByID())
	app.Server().AddRoute("/src/pengetahuan/{id}", "PUT", pengetahuan.REST().UpdateByID, pengetahuan.OpenAPI().UpdateByID())
	app.Server().AddRoute("/src/pengetahuan/{id}", "PATCH", pengetahuan.REST().PartiallyUpdateByID, pengetahuan.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/src/pengetahuan/{id}", "DELETE", pengetahuan.REST().DeleteByID, pengetahuan.OpenAPI().DeleteByID())

	app.Server().AddRoute("/src/referensi", "POST", referensi.REST().Create, referensi.OpenAPI().Create())
	app.Server().AddRoute("/src/referensi", "GET", referensi.REST().Get, referensi.OpenAPI().Get())
	app.Server().AddRoute("/src/referensi/{id}", "GET", referensi.REST().GetByID, referensi.OpenAPI().GetByID())
	app.Server().AddRoute("/src/referensi/{id}", "PUT", referensi.REST().UpdateByID, referensi.OpenAPI().UpdateByID())
	app.Server().AddRoute("/src/referensi/{id}", "PATCH", referensi.REST().PartiallyUpdateByID, referensi.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/src/referensi/{id}", "DELETE", referensi.REST().DeleteByID, referensi.OpenAPI().DeleteByID())

	// AddRoute : DONT REMOVE THIS COMMENT
}
