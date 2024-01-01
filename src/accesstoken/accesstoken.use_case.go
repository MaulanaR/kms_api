package accesstoken

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/maulanar/kms/app"
	"github.com/maulanar/kms/src/user"
	"grest.dev/grest"
)

// UseCase returns a UseCaseHandler for expected use case functional.
func UseCase(ctx app.Ctx, query ...url.Values) UseCaseHandler {
	u := UseCaseHandler{
		Ctx:   &ctx,
		Query: url.Values{},
	}
	if len(query) > 0 {
		u.Query = query[0]
	}
	return u
}

// UseCaseHandler provides a convenient interface for AccessToken use case, use UseCase to access UseCaseHandler.
type UseCaseHandler struct {
	AccessToken

	// injectable dependencies
	Ctx   *app.Ctx   `json:"-" db:"-" gorm:"-"`
	Query url.Values `json:"-" db:"-" gorm:"-"`
}

// Async return UseCaseHandler with async process.
func (u UseCaseHandler) Async(ctx app.Ctx, query ...url.Values) UseCaseHandler {
	ctx.IsAsync = true
	return UseCase(ctx, query...)
}

// GetByID returns the AccessToken data for the specified ID.
func (u UseCaseHandler) GetByID(id string) (AccessToken, error) {
	res := AccessToken{}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// get from db
	key := "access_token"
	if id == "" {
		id = u.Ctx.Token.AccessToken.String
	}
	u.Query.Add(key, id)
	err = app.Query().First(tx, &res, u.Query)
	if err != nil {
		return res, u.Ctx.NotFoundError(err, u.EndPoint(), key, id)
	}

	return res, err
}

// Get returns the list of AccessToken data.
func (u UseCaseHandler) Get() (app.ListModel, error) {
	res := app.ListModel{}

	// check permission
	err := u.Ctx.ValidatePermission("access_token.list")
	if err != nil {
		return res, err
	}
	// get from cache and return if exists
	cacheKey := u.EndPoint() + "?" + u.Query.Encode()
	err = app.Cache().Get(cacheKey, &res)
	if err == nil {
		return res, err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// set pagination info
	res.Count,
		res.PageContext.Page,
		res.PageContext.PerPage,
		res.PageContext.PageCount,
		err = app.Query().PaginationInfo(tx, &AccessToken{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	// return data count if $per_page set to 0
	if res.PageContext.PerPage == 0 {
		return res, err
	}

	// find data
	data, err := app.Query().Find(tx, &AccessToken{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	res.SetData(data, u.Query)

	return res, err
}

// Create creates a new data AccessToken with specified parameters.
func (u UseCaseHandler) Login(p *ParamCreate) (AccessToken, error) {
	res := AccessToken{}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// validasi data
	if p.Password.String == "" {
		return res, app.Error().New(http.StatusUnauthorized, "Password wajib diisikan", map[string]string{"Password": "Password is required"})
	}
	if p.Username.String == "" {
		return res, app.Error().New(http.StatusUnauthorized, "Username wajib diisikan", map[string]string{"Password": "Password is required"})
	}

	//CHECK apakah user/ pass ada dari api stara
	//jika dari api stara responsenya ok, maka cek di db apakah datanya ada
	//jika ada, maka update detailnya sesuai yg di api
	///jika blm ada. maka insert

	loginBPKP := false
	dataUser := user.User{}

	endPoint := "http://api-stara.bpkp.go.id/api/auth/login"
	body := map[string]any{
		"username": p.Username.String,
		"password": p.Password.String,
	}

	c := app.HttpClient("POST", endPoint)
	c.Debug()
	c.AddJsonBody(body)
	_, errAPI := c.Send()
	if errAPI == nil {
		staraAuth := LoginStara{}
		errAPI = grest.NewJSON(c.BodyResponse).ToFlat().Unmarshal(&staraAuth)
		if errAPI == nil {
			if staraAuth.HttpCode.Int64 == 200 {
				// maka sukses login
				//cek apakah user sudah ada, by username
				err = tx.Table("m_user").Where("username", staraAuth.Username.String).First(&dataUser).Error
				if err != nil {
					//maka data user blm ada, lanjut insert
					dataBaru := user.ParamCreate{}
					dataBaru.OrangNama.NullString = staraAuth.NamaGelar.NullString
					dataBaru.OrangNamaPanggilan.NullString = staraAuth.Name.NullString
					dataBaru.OrangNip.NullString = staraAuth.Nipbaru.NullString
					dataBaru.OrangNik.NullString = staraAuth.Nik.NullString
					if staraAuth.JenisKelamin.String == "Laki-laki" {
						dataBaru.OrangJenisKelamin.Set("pria")
					} else {
						dataBaru.OrangJenisKelamin.Set("wanita")
					}
					dataBaru.OrangTelp.NullString = staraAuth.Nomorhp.NullString
					dataBaru.OrangJabatan.NullString = staraAuth.Jabatan.NullString
					dataBaru.OrangUnitKerja.NullString = staraAuth.Namaunit.NullString
					dataBaru.Username.NullString = staraAuth.Username.NullString
					dataBaru.Kategori.Set("BPKP")
					dataBaru.UsernameStara.Set(p.Username.String)
					//dataBaru.OrangTempatLahir = staraAuth.
					// dataBaru.OrangTglLahir = staraAuth.
					// dataBaru.OrangAlamat = staraAuth.
					// dataBaru.OrangEmail = staraAuth.
					// dataBaru.OrangFotoID = staraAuth.
					// dataBaru.OrangFotoUrl = staraAuth.
					// dataBaru.OrangFotoNama = staraAuth.
					// dataBaru.OrangUserLevel = staraAuth.
					// dataBaru.OrangStatusLevel = staraAuth.
					// dataBaru.Level = staraAuth.
					// dataBaru.Points = staraAuth.
					dataBaru.Password.Set("")
					err = user.UseCase(*u.Ctx).Create(&dataBaru)
					if err != nil {
						return res, err
					}

					dataUser = dataBaru.User
				} else {
					//maka data user ditemukan, lanjut update data terbaru.
					//maka data user blm ada, lanjut insert
					dataUpdate := user.ParamPartiallyUpdate{}
					dataUpdate.OrangNama.NullString = staraAuth.NamaGelar.NullString
					dataUpdate.OrangNamaPanggilan.NullString = staraAuth.Name.NullString
					dataUpdate.OrangNip.NullString = staraAuth.Nipbaru.NullString
					dataUpdate.OrangNik.NullString = staraAuth.Nik.NullString
					if staraAuth.JenisKelamin.String == "Laki-laki" {
						dataUpdate.OrangJenisKelamin.Set("pria")
					} else {
						dataUpdate.OrangJenisKelamin.Set("wanita")
					}
					dataUpdate.OrangTelp.NullString = staraAuth.Nomorhp.NullString
					dataUpdate.OrangJabatan.NullString = staraAuth.Jabatan.NullString
					dataUpdate.OrangUnitKerja.NullString = staraAuth.Namaunit.NullString
					dataUpdate.Username.NullString = staraAuth.Username.NullString
					dataUpdate.Kategori.Set("BPKP")
					dataUpdate.UsernameStara.Set(p.Username.String)
					err = user.UseCase(*u.Ctx).PartiallyUpdateByID(strconv.Itoa(int(dataUser.ID.Int64)), &dataUpdate)
					if err != nil {
						return res, err
					}
				}
				loginBPKP = true
			}
		}
	}

	if !loginBPKP {
		//query db
		err = tx.Table("m_user").Where("username", p.Username.String).First(&dataUser).Error
		if err != nil {
			return res, app.Error().New(http.StatusBadRequest, app.Translator().Trans(u.Ctx.Lang, "username_wrong", map[string]string{}))
		}

		//validate password
		decrypt, err := app.Crypto().Decrypt(dataUser.Password.String)
		if err != nil {
			return res, err
		}

		if p.Password.String != decrypt {
			return res, app.Error().New(http.StatusBadRequest, app.Translator().Trans(u.Ctx.Lang, "password_wrong", map[string]string{}))
		}
	}

	//generate access token
	res.AccessToken.Set(app.Crypto().GenerateAccessToken(25))
	res.ExpiredAt.Set(time.Now().AddDate(0, 0, 1))
	res.UserId = dataUser.ID
	res.CreatedAt.Set(time.Now())

	err = tx.Create(&res).Error
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	return res, nil
}
