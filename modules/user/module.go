package user

import (
	"context"
	"local/biz"
	"local/biz/ac"
	"local/biz/mdl"
	"local/biz/modules/group"

	"github.com/go-pg/pg"
	// "github.com/go-pg/pg/orm"
	vld "gopkg.in/go-playground/validator.v9"
)

// Module injection provider and bootstrap
var Module = biz.Module{
	Provider: func(db *pg.DB, groupRepo group.RepoI, validator *vld.Validate) (RepoI, SvsI, error) {
		var impl RepoI = repoImpl{
			db: db,
		}
		var svs SvsI = svsImpl{
			repo:      impl,
			groupRepo: groupRepo,
			vld:       validator, // add custom rule when need
		}
		return impl, svs, nil
	},
}

// RepoI .
type RepoI interface {
	Create(*mdl.User) (int, error)
	Update(*mdl.User) error
	FindByUsername(string) (*mdl.User, error)
	FindByID(int) (*mdl.User, error)
	SetGroups4User(userID int, groupIDs *[]string) error
}

// SvsI user service
type SvsI interface {
	Register(context.Context, *RegisterUserParam) error
	SetGroups4User(context.Context, *SetGroups4UserParam) error
	AddUser(context.Context, *mdl.User) (id int, err error)
	FindByID(context.Context, int) (*mdl.User, error)
	GetUserAsSub(int) (ac.Sub, error)
}

// SetGroups4UserParam .
type SetGroups4UserParam struct {
	GroupIDs *[]string `validate:"required,gt=1"`
	UserID   int       `validate:"required"`
}

// RegisterUserParam 注册参数
type RegisterUserParam struct {
	Phone    string `validate:"required"`
	Captch   string `validate:"required"`
	Username string `validate:"required"`
	Password string `validate:"required"`
}
