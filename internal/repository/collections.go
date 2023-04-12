package repository

import (
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (r *Repository) UserCollections(filter request.CollectionsFilter) ([]entity.Collections,  error) {

	res := []entity.Collections{}
	f := bson.D{}

	// if filter.AllowEmpty != nil && *filter.AllowEmpty == false {
	// 	f = append(f, bson.E{"total_items", bson.M{"$gt": 0}})
	// }

	if filter.Address != nil && *filter.Address != "" {
		f = append(f, bson.E{"contract", primitive.Regex{Pattern: *filter.Address, Options: "i"}})
	}

	if filter.Name != nil && *filter.Name != "" {
		f = append(f, bson.E{"name", primitive.Regex{Pattern: *filter.Name, Options: "i"}})
	}

	if filter.Owner != nil && *filter.Owner != "" {
		f = append(f, bson.E{"creator", primitive.Regex{Pattern: *filter.Owner, Options: "i"}})
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
	err := r.Find(utils.COLLECTION_COLLECTIONS, f, int64(*filter.Limit), int64(*filter.Offset), &res, s)
	if err != nil {
		return nil, err
	}

	
	return res, nil

}