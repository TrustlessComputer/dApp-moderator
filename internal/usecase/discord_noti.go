package usecase

import (
	"context"
	discordclient "dapp-moderator/utils/discord"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"

	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/logger"
)

func (u *Usecase) NewTokenNotify(token entity.Token) error {
	return u.CreateDiscordNotify(&entity.DiscordNotification{
		Message: discordclient.Message{
			Content: fmt.Sprintf("**NEW BRC-20 #%d**", token.Index),
			Embeds: []discordclient.Embed{
				{
					Fields: []discordclient.Field{
						{
							Name:   fmt.Sprintf("Token Name: [%s](https://explorer.trustless.computer/token/%s/token-transfers)", token.Name, token.Address),
							Inline: false,
						},
						{
							Name:   fmt.Sprintf("Token Symbal: %v", token.Symbol),
							Inline: false,
						},
						{
							Name:   fmt.Sprintf("Token Supply: %v", token.TotalSupply),
							Inline: false,
						},
						{
							Name:   fmt.Sprintf("Owner: [%s](https://explorer.trustless.computer/token/%s/token-transfers)", token.Name, token.Address),
							Inline: false,
						},
					},
				},
			},
		},
	})
}
func (u *Usecase) NewCollectionNotify(collection entity.Collections) error {
	panic("not implemented")
}
func (u *Usecase) NewNameNotify(airdrop entity.Bns) error {
	panic("not implemented")
}
func (u *Usecase) NewArtifactNotify(airdrop entity.Nfts) error {
	panic("not implemented")
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
			discordMsg := &notify.Message
			logger.AtLog.Logger.Info("sending new airdrop message to discord", zap.Any("discordMsg", discordMsg))
			if err := u.DiscordClient.SendMessage(context.TODO(), notify.Webhook, *discordMsg); err != nil {
				u.Repo.UpdateDiscord(context.TODO(), notify.ID.String(), map[string]interface{}{
					"num_retried": notify.NumRetried + 1,
				})

				if notify.NumRetried+1 == entity.MaxSendDiscordRetryTimes {
					u.Repo.UpdateDiscord(context.TODO(), notify.ID.String(), map[string]interface{}{
						"status": entity.FAILED,
						"note":   fmt.Sprintf("failed after %d times", entity.MaxSendDiscordRetryTimes),
					})
				}
			} else {
				u.Repo.UpdateDiscord(context.TODO(), notify.ID.String(), map[string]interface{}{
					"status": entity.DONE,
					"note":   "messaged is sent at " + time.Now().Format(time.RFC3339),
				})
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
		if partner.MatchProject(notify.Meta.ProjectID) && partner.MatchCategory(notify.Meta.Category) && partner.MatchAmountGreaterThanZero(notify.Meta.Amount) {
			notify.Webhook = webhook
			notify.Meta.SendTo = partner.Name
			u.Repo.CreateDiscordNotification(context.TODO(), notify)
		}
	}

	return nil
}

func (u *Usecase) TestSendNotify() {
	domain := os.Getenv("DOMAIN")
	if domain == "https://devnet.generative.xyz" {
		u.JobSendDiscord()
		fmt.Println("done")
	}
}
