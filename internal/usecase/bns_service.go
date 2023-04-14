package usecase

import (
	"context"
	"dapp-moderator/external/bns_service"
	"dapp-moderator/internal/delivery/http/request"
	"fmt"
	"net/url"
)

func (u *Usecase) BnsNames(ctx context.Context, filter request.FilterBNSNames) ([]*bns_service.NameResp, error) {
	params := url.Values{}
	if filter.Limit != nil {
		params.Set("limit", fmt.Sprintf("%d", *filter.Limit))
	}

	if filter.Offset != nil {
		params.Set("offset", fmt.Sprintf("%d", *filter.Offset))
	}

	resp, err := u.BnsService.Names(params)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *Usecase) BnsName(ctx context.Context, name string) (*bns_service.NameResp, error) {
	resp, err := u.BnsService.Name(name)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *Usecase) BnsNameAvailable(ctx context.Context, name string) (*bool, error) {
	resp, err := u.BnsService.NameAvailable(name)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *Usecase) BnsNamesOnwedByWalletAddress(ctx context.Context, walletAdress string, filter request.FilterBNSNames) ([]*bns_service.NameResp, error) {
	params := url.Values{}
	if filter.Limit != nil {
		params.Set("limit", fmt.Sprintf("%d", *filter.Limit))
	}

	if filter.Offset != nil {
		params.Set("offset", fmt.Sprintf("%d", *filter.Offset))
	}

	resp, err := u.BnsService.NameOnwedByWalletAddress(walletAdress, params)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
