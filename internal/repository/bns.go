package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) UpdateBnsResolver(tokenID string, resolver string) (*mongo.UpdateResult, error) {
	f := bson.D{
		{"token_id", tokenID},
	}

	now := time.Now()
	update := bson.M{"$set": bson.M{
		"resolver":   strings.ToLower(resolver),
		"updated_at": now,
	}}
	updated, err := r.UpdateOne(utils.COLLECTION_BNS, f, update)

	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (r *Repository) UpdateBnsPfp(tokenID string, pfp string) (*mongo.UpdateResult, error) {
	f := bson.D{
		{"token_id", tokenID},
	}

	update := bson.M{"$set": bson.M{"pfp": strings.ToLower(pfp)}}
	updated, err := r.UpdateOne(utils.COLLECTION_BNS, f, update)

	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (r *Repository) UpdateBnsPfpData(tokenID string, pfp *entity.BnsPfpData) (*mongo.UpdateResult, error) {
	f := bson.D{
		{"token_id", tokenID},
	}

	update := bson.M{"$set": bson.M{"pfp_data": pfp}}
	updated, err := r.UpdateOne(utils.COLLECTION_BNS, f, update)

	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (r *Repository) FilterBNS(filter entity.FilterBns, fromCollection ...string) ([]*entity.FilteredBNS, error) {
	resp := []*entity.FilteredBNS{}
	f := bson.A{}
	match := bson.D{}
	if filter.Resolver != nil && *filter.Resolver != "" {
		match = append(match, bson.E{"resolver", strings.ToLower(*filter.Resolver)})
	}

	if filter.PFP != nil && *filter.PFP != "" {
		match = append(match, bson.E{"pfp", strings.ToLower(*filter.PFP)})
	}

	if filter.Owner != nil && *filter.Owner != "" {
		match = append(match, bson.E{"owner", strings.ToLower(*filter.Owner)})
	}

	if filter.Name != nil && *filter.Name != "" {
		match = append(match, bson.E{"name", primitive.Regex{Pattern: *filter.Name, Options: "i"}})
	}

	if filter.TokenID != nil && *filter.TokenID != "" {
		match = append(match, bson.E{"token_id", *filter.TokenID})
	}

	if filter.Limit <= 0 {
		filter.Limit = 100
	}

	if len(match) > 0 {
		f = append(f, bson.D{{"$match", match}})
	}

	f = append(f, bson.D{{"$addFields", bson.D{{"id", "$token_id"}}}})
	sort := bson.D{{"token_id_int", entity.SORT_DESC}}
	if filter.SortBy != "" && filter.Sort != 0 {
		sort = bson.D{{filter.SortBy, filter.Sort}}
	}
	f = append(f, bson.D{{"$sort", sort}})
	f = append(f, bson.D{{"$skip", filter.Offset}})
	f = append(f, bson.D{{"$limit", filter.Limit}})

	collectionName := utils.VIEW_BNS // default from bns_view
	if len(fromCollection) > 0 && fromCollection[0] != "" {
		collectionName = fromCollection[0]
	}
	cursor, err := r.DB.Collection(collectionName).Aggregate(context.TODO(), f, nil)
	if err != nil {
		return nil, err
	}

	err = cursor.All((context.TODO()), &resp)
	if err != nil {
		return nil, err
	}

	if filter.Resolver != nil && *filter.Resolver != "" {
		r.SortBNSDefaultOfResolverFirst(resp, strings.ToLower(*filter.Resolver))
	}
	return resp, nil
}

func (r *Repository) SortBNSDefaultOfResolverFirst(bns []*entity.FilteredBNS, resolver string) {
	if len(bns) < 2 {
		return
	}

	result, err := r.FindOne(utils.COLLECTION_BNS_DEFAULT, bson.D{{"resolver", resolver}})
	if err != nil {
		return
	}
	bnsDefault := &entity.BNSDefault{}
	if err := result.Decode(bnsDefault); err != nil {
		return
	}
	bnsEntity := &entity.Bns{}
	result, err = r.FindOne(utils.COLLECTION_BNS, bson.D{{"_id", bnsDefault.BNSDefaultID}})
	if err != nil {
		return
	}
	if err := result.Decode(bnsEntity); err != nil {
		return
	}

	if bnsDefault != nil {
		for index, item := range bns {
			if strings.ToLower(item.TokenID) == strings.ToLower(bnsEntity.TokenID) {
				// swap item to first index
				bns[0], bns[index] = bns[index], bns[0]
			}
		}

		return
	}

	if bns[0].PfpData == nil {
		// Ưu tiên set phần tử đầu tiên là item có pfp_data
		for index, item := range bns {
			if item.PfpData != nil {
				bns[0], bns[index] = bns[index], bns[0]
			}
		}
	}

	return
}

func (r *Repository) FindBNS(tokenID string) (*entity.FilteredBNS, error) {
	resp := &entity.FilteredBNS{}
	f := bson.D{{"token_id", tokenID}}

	cursor := r.DB.Collection(utils.VIEW_BNS).FindOne(context.TODO(), f)

	err := cursor.Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
