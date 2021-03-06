package branch

import (
	"local/biz"
	"local/biz/mdl"

	// "log"

	"github.com/go-pg/pg"
)

// Module info
var (
	Module = biz.Module{
		Provider: func(db *pg.DB) RepoI {
			return repoImpl{
				db: db,
			}
		},
	}
)

// RepoI branch repository
type RepoI interface {
	Create(*mdl.Branch) (int, error)
	SelectByID(int) (*mdl.Branch, error)
	Update(*mdl.Branch) error
	SelectAll() (*[]mdl.Branch, error)
	DeleteByID(int) error
}
