package repository

import (
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// All the created collections and the collections which have the owned nfts
func (r *Repository) UserCollections(filter request.CollectionsFilter) ([]entity.Collections, error) {
	res := []entity.Collections{}
	f := bson.D{}

	collectionIDs := []string{}
	data, err := r.CollectionsByNfts(*filter.Owner)
	if err == nil {
		for _, item := range data {
			collectionIDs = append(collectionIDs, item.ID.CollectionAddress)
		}
	}

	if filter.Address != nil && *filter.Address != "" {
		f = append(f, bson.E{"contract", primitive.Regex{Pattern: *filter.Address, Options: "i"}})
	}

	if filter.Name != nil && *filter.Name != "" {
		f = append(f, bson.E{"name", primitive.Regex{Pattern: *filter.Name, Options: "i"}})
	}

	if filter.Owner != nil && *filter.Owner != "" {
		f = append(f, bson.E{"contract", bson.M{"$in": collectionIDs}})
	}

	sortBy := "deployed_at_block"
	if filter.SortBy != nil && *filter.SortBy != "" {
		sortBy = *filter.SortBy
	}

	sort := 1
	if filter.Sort != nil {
		sort = *filter.Sort
	}

	s := bson.D{{sortBy, sort}, {"index", 1}}
	err = r.Find(utils.COLLECTION_COLLECTIONS, f, int64(*filter.Limit), int64(*filter.Offset), &res, s)
	if err != nil {
		return nil, err
	}

	return res, nil

}
