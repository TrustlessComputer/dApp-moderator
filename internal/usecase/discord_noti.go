package usecase

import (
	"context"
	"dapp-moderator/external/bns_service"
	"dapp-moderator/utils"
	discordclient "dapp-moderator/utils/discord"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"

	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/logger"
)

func (u *Usecase) NewTokenNotify(token *entity.Token) error {

	return u.CreateDiscordNotify(&entity.DiscordNotification{
		Message: discordclient.Message{
			Content:   fmt.Sprintf("**NEW BRC-20 #%d**", token.Index),
			Username:  "Satoshi 27",
			AvatarUrl: "",
			Embeds: []discordclient.Embed{
				{
					Url: fmt.Sprintf("https://explorer.trustless.computer/token/%s/token-transfers", token.Address),
					Fields: []discordclient.Field{
						{
							Value:  fmt.Sprintf("**Token Name: [%s](https://explorer.trustless.computer/token/%s/token-transfers)**", token.Name, token.Address),
							Inline: false,
						},
						{
							Value:  fmt.Sprintf("**Token Symbal: %v**", token.Symbol),
							Inline: false,
						},
						{
							Value:  fmt.Sprintf("**Token Supply: %s**", utils.FormatStringNumber(token.TotalSupply)),
							Inline: false,
						},
						{
							Value:  fmt.Sprintf("**Owner: [%s](https://explorer.trustless.computer/token/%s/token-transfers)**", utils.ShortenBlockAddress(token.Owner), token.Owner),
							Inline: false,
						},
					},
				},
			},
		},
		Status: entity.PENDING,
		Event:  entity.EventNewToken,
	})
}

func (u *Usecase) NewCollectionNotify(collection *entity.Collections) error {
	return u.CreateDiscordNotify(&entity.DiscordNotification{
		Message: discordclient.Message{
			Content:   fmt.Sprintf("**NEW BRC-721 #%d**", collection.Index),
			Username:  "Satoshi 27",
			AvatarUrl: "",
			Embeds: []discordclient.Embed{
				{
					Url: fmt.Sprintf("https://explorer.trustless.computer/token/%s/token-transfers", collection.Contract),
					Fields: []discordclient.Field{
						{
							Value:  fmt.Sprintf("**Collection Name: [%s](https://trustless.computer/collection?contrac=%s)**", collection.Name, collection.Contract),
							Inline: false,
						},
						{
							Value:  fmt.Sprintf("**Suply: %d**", collection.TotalItems),
							Inline: false,
						},
						{
							Value:  fmt.Sprintf("**Contract: %s**", utils.ShortenBlockAddress(collection.Contract)),
							Inline: false,
						},
						{
							Value:  fmt.Sprintf("**Owner: [%s](https://explorer.trustless.computer/address/%s/token-transfers)**", utils.ShortenBlockAddress(collection.Creator), collection.Creator),
							Inline: false,
						},
					},
				},
			},
		},
		Status: entity.PENDING,
		Event:  entity.EventNewCollection,
	})
}

func (u *Usecase) NewNameNotify(bns *bns_service.NameResp) error {
	return u.CreateDiscordNotify(&entity.DiscordNotification{
		Message: discordclient.Message{
			Content:   fmt.Sprintf("**NEW BNS #%s**", bns.ID),
			Username:  "Satoshi 27",
			AvatarUrl: "",
			Embeds: []discordclient.Embed{
				{
					Fields: []discordclient.Field{
						{
							Value:  fmt.Sprintf("**Name: [%s](https://trustless.computer/dapp?tab=name&key=%s)**", bns.Name, bns.Name),
							Inline: false,
						},
						{
							Value:  fmt.Sprintf("**Owner: [%s](https://explorer.trustless.computer/address/%s/token-transfers)**", utils.ShortenBlockAddress(bns.Owner), bns.Owner),
							Inline: false,
						},
					},
				},
			},
		},
		Status: entity.PENDING,
		Event:  entity.EventNewName,
	})
}

func (u *Usecase) NewArtifactNotify(nfts *entity.Nfts) error {
	notify := &entity.DiscordNotification{
		Message: discordclient.Message{
			Content:   fmt.Sprintf("**NEW ARTIFACT #%s**", nfts.TokenID),
			Username:  "Satoshi 27",
			AvatarUrl: "",
			Embeds: []discordclient.Embed{
				{
					Fields: []discordclient.Field{
						{
							Value:  fmt.Sprintf("**Owner: [%s](https://explorer.trustless.computer/address/%s/token-transfers)**", utils.ShortenBlockAddress(nfts.Owner), nfts.Owner),
							Inline: false,
						},
						{
							Value:  fmt.Sprintf("**Type: %s**", nfts.ContentType),
							Inline: false,
						},
					},
				},
			},
		},
		Status: entity.PENDING,
		Event:  entity.EventNewArtifact,
	}

	if nfts.Image != "" && strings.HasPrefix(nfts.Image, "/dapp/api/nft-explorer/collections/") {
		notify.Message.Embeds[0].Image.Url = "https://dapp.trustless.computer" + nfts.Image
	}

	return u.CreateDiscordNotify(notify)
}

func (u *Usecase) JobSendDiscord() error {
	logger.AtLog.Logger.Info("JobSendDiscord Starting ...")
	for page := int64(1); ; page++ {

		notifications, err := u.Repo.FindDiscordNotifications(context.TODO(), entity.GetDiscordReq{
			Page:   page,
			Limit:  10,
			Status: entity.PENDING,
		})
		if err != nil {
			return err
		}

		if len(notifications) == 0 {
			break
		}

		for _, notify := range notifications {
			if err = u.DiscordClient.SendMessage(context.TODO(), notify.Webhook, notify.Message); err != nil {
				logger.AtLog.Logger.Info("Send discord message failed", zap.Error(err))
				err = u.Repo.UpdateDiscord(context.TODO(), notify.Id(), map[string]interface{}{
					"num_retried": notify.NumRetried + 1,
				})
				if err != nil {
					logger.AtLog.Logger.Info("Update discord status failed", zap.Error(err))
				}

				if notify.NumRetried+1 == entity.MaxSendDiscordRetryTimes {
					err = u.Repo.UpdateDiscord(context.TODO(), notify.Id(), map[string]interface{}{
						"status": entity.FAILED,
						"note":   fmt.Sprintf("failed after %d times", entity.MaxSendDiscordRetryTimes),
					})
					if err != nil {
						logger.AtLog.Logger.Info("Update discord status failed", zap.Error(err))
					}
				}
			} else {
				err = u.Repo.UpdateDiscord(context.TODO(), notify.Id(), map[string]interface{}{
					"status": entity.DONE,
					"note":   "messaged is sent at " + time.Now().Format(time.RFC3339),
				})
				if err != nil {
					logger.AtLog.Logger.Info("Update discord status failed", zap.Error(err))
				}
			}
		}
	}

	return nil
}

func (u *Usecase) CreateDiscordNotify(notify *entity.DiscordNotification) error {
	partners, err := u.Repo.GetAllDiscordPartner()
	if err != nil {
		return err
	}
	for _, partner := range partners {
		webhook := partner.Webhooks[string(notify.Event)]
		if webhook == "" {
			continue
		}
		notify.Webhook = webhook
		notify.Meta.SendTo = partner.Name
		err = u.Repo.CreateDiscordNotification(context.TODO(), notify)
		if err != nil {
			logger.AtLog.Error("CreateDiscordNotification", zap.Error(err))
		}
	}

	return nil
}

func (u *Usecase) TestSendNotify() {
	env := os.Getenv("ENVIRONMENT")
	if env == "local" {
		token, _ := u.Repo.FindToken(context.TODO(), entity.TokenFilter{
			Address: "0xF3368Ba95164A9dFbd8E1F7D33F222D565C10bf6",
		})
		err := u.NewTokenNotify(token)
		logger.AtLog.Info("err", zap.Error(err))

		collection := &entity.Collections{}
		sr, err := u.Repo.FindOne(utils.COLLECTION_COLLECTIONS, bson.D{
			{"contract", primitive.Regex{Pattern: "0xbecb1bd2e2849532520da6d36b1684ce63b68617", Options: "i"}},
		})
		sr.Decode(collection)

		err = u.NewCollectionNotify(collection)
		logger.AtLog.Info("err", zap.Error(err))

		u.NewNameNotify(&bns_service.NameResp{
			Name:  "annhien",
			Owner: "0x82268af8207117ddbcd8ce4e444263ccd8d1bf87",
			ID:    "1166",
		})

		nftResp, _ := u.Repo.FindOne(entity.Nfts{}.CollectionName(), bson.D{
			{"token_id", "1"},
		})
		nfts := &entity.Nfts{}
		nftResp.Decode(nfts)

		u.NewArtifactNotify(nfts)

		u.JobSendDiscord()
		fmt.Println("done")
	}
}