package repository

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) CreateDiscordNotification(ctx context.Context, notify *entity.DiscordNotification) error {
	now := time.Now().UTC()
	notify.ID = primitive.NewObjectID()
	notify.CreatedAt = &now

	if notify.UUID == "" {
		notify.UUID = notify.Id()
	}

	_, err := r.DB.Collection(utils.COLLECTION_DISCORD_NOTIFICATION).InsertOne(ctx, notify)
	return err
}

func (r *Repository) FindDiscordNotifications(ctx context.Context, req entity.GetDiscordReq) ([]entity.DiscordNotification, error) {

	var results []entity.DiscordNotification
	filter := bson.M{"status": req.Status}

	// paging
	numToSkip := (req.Page - 1) * req.Limit
	options := options.Find()
	options.SetSkip(numToSkip)
	options.SetLimit(req.Limit)

	// sort
	options.Sort = bson.D{{"created_at", 1}}

	cursor, err := r.DB.Collection(utils.COLLECTION_DISCORD_NOTIFICATION).Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var msg entity.DiscordNotification
		err = cursor.Decode(&msg)
		if err != nil {
			return nil, err
		}
		results = append(results, msg)
	}

	return results, nil
}

func (r *Repository) UpdateDiscord(ctx context.Context, uuid string, fields map[string]interface{}) error {
	filter := bson.M{
		"uuid": uuid,
	}

	update := bson.M{}
	for k, v := range fields {
		update[k] = v
	}
	update["updated_at"] = time.Now().UTC()

	result, err := r.DB.Collection(utils.COLLECTION_DISCORD_NOTIFICATION).UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (r *Repository) GetAllDiscordPartner() ([]entity.DiscordPartner, error) {
	var partners []entity.DiscordPartner
	f := bson.M{}

	cursor, err := r.DB.Collection(utils.COLLECTION_DISCORD_PARTNERS).Find(context.TODO(), f)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &partners); err != nil {
		return nil, err
	}
	return partners, nil
}
