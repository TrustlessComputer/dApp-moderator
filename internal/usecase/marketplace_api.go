package usecase

import (
	"context"
	"dapp-moderator/internal/entity"
)

func (u *Usecase) FilterMKListing(ctx context.Context, filter entity.FilterMarketplaceListings) ([]entity.MarketplaceListings, error) {
	return u.Repo.FilterMarketplaceListings(filter)
}

func (u *Usecase) FilterMKOffers(ctx context.Context, filter entity.FilterMarketplaceOffer) ([]entity.MarketplaceOffers, error) {
	return u.Repo.FilterMarketplaceOffer(filter)
}

func (u *Usecase) FilterTokenActivities(ctx context.Context, filter entity.FilterTokenActivities) ([]entity.MarketplaceTokenActivity, error) {
	return u.Repo.FilterTokenActivites(filter)
}
