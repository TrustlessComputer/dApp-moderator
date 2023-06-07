package usecase

import (
	"context"
	"dapp-moderator/external/bns_service"
	"dapp-moderator/utils"
	discordclient "dapp-moderator/utils/discord"
	"fmt"
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
							Value:  fmt.Sprintf("**Token Symbol: %v**", token.Symbol),
							Inline: false,
						},
						{
							Value:  fmt.Sprintf("**Token Supply: %s**", utils.FormatStringNumber(token.TotalSupply, token.Decimal)),
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
							Value:  fmt.Sprintf("**Collection Name: [%s](https://trustlessnfts.com/collection?contract=%s)**", utils.NameOrAddress(collection.Name, collection.Contract), collection.Contract),
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
							Value:  fmt.Sprintf("**Name: [%s](https://trustless.domains)**", bns.Name),
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
	message := discordclient.Message{
		Content:   fmt.Sprintf("**NEW SMART INSCRIPTION #%s**", nfts.TokenID),
		Username:  "Satoshi 27",
		AvatarUrl: "",
		Embeds: []discordclient.Embed{
			{
				Fields: []discordclient.Field{
					{
						Value:  fmt.Sprintf("**Owner: [%s](https://smartinscription.xyz/token?contract=%s&id=%s)**", utils.ShortenBlockAddress(nfts.Owner), strings.ToLower(nfts.ContractAddress), nfts.TokenID),
						Inline: false,
					},
					{
						Value:  fmt.Sprintf("**Type: %s**", nfts.ContentType),
						Inline: false,
					},
				},
			},
		},
	}
	if nfts.Image != "" {
		if strings.HasPrefix(nfts.Image, "/dapp/api/nft-explorer/collections/") {
			message.Embeds[0].Thumbnail.Url = "https://dapp.trustless.computer" + nfts.Image
		} else {
			message.Embeds[0].Thumbnail.Url = nfts.Image
		}

	}
	notify := &entity.DiscordNotification{
		Message: message,
		Status:  entity.PENDING,
		Event:   entity.EventNewArtifact,
	}

	if nfts.Image != "" && strings.HasPrefix(nfts.Image, "/dapp/api/nft-explorer/collections/") {
		notify.Message.Embeds[0].Thumbnail.Url = "https://dapp.trustless.computer" + nfts.Image
	} else {
		notify.Message.Embeds[0].Thumbnail.Url = nfts.Image
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
		nft := &entity.Nfts{
			TokenID:     "1",
			Owner:       "0xb764d696aa52f6a7e6ec6647c7d7a7736f3aff59",
			Image:       "https://dapp.trustless.computer/dapp/api/nft-explorer/collections/0x7c3cee22435461f3d858af2aa1770ea75403b63e/nfts/1/content",
			Name:        "test",
			ContentType: "image/png",
		}
		u.NewArtifactNotify(nft)
		u.JobSendDiscord()
	}
}
