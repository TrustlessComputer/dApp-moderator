package usecase

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"errors"
)

func (u *Usecase) ApiCreateDappInfo(ctx context.Context, req *request.CreateDappInfoReq) (*entity.DappInfo, error) {
	if len(req.Desc) == 0 || len(req.Name) == 0 || len(req.Image) == 0 || len(req.Link) == 0 || len(req.Creator) == 0 {
		return nil, errors.New("name, image, creator, desc, link are required")
	}

	dapp := &entity.DappInfo{
		Name:        req.Name,
		Image:       req.Image,
		Link:        req.Link,
		Creator:     req.Creator,
		Description: req.Desc,
		Status:      0,
	}

	err := u.Repo.InsertDappInfo(dapp)
	if err != nil {
		return nil, err
	}
	return dapp, nil

}

func (u *Usecase) ApiListDappInfo() ([]*entity.DappInfo, error) {
	return u.Repo.ListDappInfo()
}
