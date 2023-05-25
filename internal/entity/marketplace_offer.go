package entity

import "dapp-moderator/utils"

type MarketplaceOffers struct {
	BaseEntity         `bson:",inline"`
	ID                 string            `bson:"id"`
	OfferingId         string            `bson:"offering_id"`
	CollectionContract string            `bson:"collection_contract"`
	TokenId            string            `bson:"token_id"`
	Buyer              string            `bson:"buyer"`
	Erc20Token         string            `bson:"erc_20_token"`
	Price              string            `bson:"price"`
	Status             MarketplaceStatus `bson:"status"`
	DurationTime       string            `bson:"duration_time"`
	BlockNumber        uint64            `bson:"block_number"`
	OwnerAddress       *string           `bson:"owner_address"`
}

func (u MarketplaceOffers) CollectionName() string {
	return utils.COLLECTION_MARKETPLACE_OFFER
}
