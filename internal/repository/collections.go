package repository

import (
	"context"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"dapp-moderator/utils/helpers"
	"errors"
	"os"
	"strings"

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

func (r *Repository) CollectionThumbnailByNfts() ([]*entity.CollectionNftThumbnail, error) {
	f := bson.A{
		bson.D{
			{"$match",
				bson.D{
					{"thumbnail", ""},
					{"total_items", bson.D{{"$gt", 0}}},
				},
			},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "nfts"},
					{"localField", "contract"},
					{"foreignField", "collection_address"},
					{"let", bson.D{{"token_id_int", "$token_id_int"}}},
					{"pipeline",
						bson.A{
							bson.D{{"$sort", bson.D{{"token_id_int", -1}}}},
							bson.D{{"$limit", 1}},
						},
					},
					{"as", "nfts"},
				},
			},
		},
		bson.D{
			{"$unwind",
				bson.D{
					{"path", "$nfts"},
					{"includeArrayIndex", "string"},
					{"preserveNullAndEmptyArrays", true},
				},
			},
		},
		bson.D{
			{"$project",
				bson.D{
					{"contract", 1},
					{"thumbnail", 1},
					{"nft_image", "$nfts.image"},
					{"nft_token_uri", "$nfts.token_uri"},
					{"nft_token_id_int", "$nfts.token_id_int"},
				},
			},
		},
	}

	cursor, err := r.DB.Collection(entity.Collections{}.CollectionName()).Aggregate(context.TODO(), f, nil)
	if err != nil {
		return nil, err
	}

	// display the results
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	resp := []*entity.CollectionNftThumbnail{}

	for _, item := range results {
		res := &entity.CollectionNftThumbnail{}
		err = helpers.Transform(item, res)
		if err != nil {
			return nil, err
		}

		resp = append(resp, res)
	}

	return resp, err
}

func (r *Repository) UpdateCollectionThumbnail(ctx context.Context, contract string, thumbnail string) error {
	filter := bson.M{
		"contract": contract,
	}

	update := bson.M{
		"thumbnail": thumbnail,
	}

	result, err := r.DB.Collection(entity.Collections{}.CollectionName()).UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("no document")
	}

	return nil
}

func (r *Repository) UpdateCollectionIndex(ctx context.Context, contract string, index int) error {
	filter := bson.M{
		"contract": contract,
	}

	update := bson.M{
		"index": index,
	}

	result, err := r.DB.Collection(entity.Collections{}.CollectionName()).UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("no document")
	}

	return nil
}

func (r *Repository) AllCollections() ([]entity.FilteredCollections, error) {
	result := []entity.FilteredCollections{}

	f := bson.A{
		bson.D{
			{"$project",
				bson.D{
					{"contract", 1},
					{"deployed_at_block", 1},
					{"contract_type", 1},
				},
			},
		},
	}

	cursor, err := r.DB.Collection(entity.Collections{}.CollectionName()).Aggregate(context.TODO(), f)
	if err != nil {
		return nil, err
	}
	if err = cursor.All((context.TODO()), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) GetCollection(contractAddress string) (*entity.Collections, error) {
	res := &entity.Collections{}
	f := bson.D{{
		"contract", contractAddress,
	}}

	s, err := r.FindOne(utils.COLLECTION_COLLECTIONS, f)
	if err != nil {
		return nil, err
	}

	err = s.Decode(res)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (r *Repository) GetSoulCollection() (*entity.Collections, error) {
	result := &entity.Collections{}

	f := bson.D{
		{"contract", strings.ToLower(os.Getenv("SOUL_CONTRACT"))},
	}

	cursor := r.DB.Collection(entity.Collections{}.CollectionName()).FindOne(context.TODO(), f)

	if err := cursor.Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) CollectionNftOwner(f entity.FilterCollectionNftOwners) ([]*entity.CollectionNftOwner, error) {
	result := []*entity.CollectionNftOwner{}

	filter := bson.A{
		bson.D{
			{"$match",
				bson.D{
					{"collection_address", strings.ToLower(*f.ContractAddress)},
				},
			},
		},
		bson.D{{"$skip", f.Offset}},
		bson.D{{"$limit", f.Limit}},
		bson.D{{"$sort", bson.D{
			{"count", entity.SORT_DESC},
		}}},
	}

	cursor, err := r.DB.Collection(utils.VIEW_MARKETPLACE_COUNT_COLLECTION_OWNER).Aggregate(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	if err = cursor.All((context.TODO()), &result); err != nil {
		return nil, err
	}

	return result, nil
}
