package user

import (
	"context"
	"fmt"

	"local/biz"
	"local/biz/mdl"
	"local/biz/modules/group"

	vld "gopkg.in/go-playground/validator.v9"
)

type svsImpl struct {
	repo RepoI
	groupRepo group.RepoI
	vld vld.Validate
}

func (s svsImpl) Register(ctx context.Context,p *RegisterUserParam) error {
	//TODO
	return nil
}

func (s svsImpl) SetGroups4User(ctx context.Context, p *SetGroups4UserParam) error {
	err := s.vld.Struct(p)
	if err != nil {
		return err
	}
	
	// caller, _ := biz.GetUsrFromContext(ctx)
	//TODO ACL
	allGroups, err := s.groupRepo.ListAll()
	if err != nil {
		return err
	}

	outer:
	for _, id := range *p.GroupIDs {
		for _, g := range *allGroups {
			if g.ID == id {
				continue outer
			}
		}
		return biz.NewErr(biz.CodeBadRequest, fmt.Sprintf("Group with id: %s not exists", id))
	}

	u, err := s.repo.FindByID(p.UserID)
	if err != nil {
		return err
	}
	if u == nil {
		return biz.ErrNotExists
	}

	return s.repo.SetGroups4User(p.UserID, p.GroupIDs)
}

func (s svsImpl) AddUser(ctx context.Context, user *mdl.User) (id uint32, err error) {
	return s.repo.Create(user)
}

func (s svsImpl) FindByID(ctx context.Context, id uint32) (*mdl.User, error) {
	return s.repo.FindByID(id)
}