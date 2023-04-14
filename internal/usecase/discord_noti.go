package usecase

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"

	"dapp-moderator/internal/entity"
	"dapp-moderator/utils/logger"
)

func (c *Usecase) NewTokenNotify(token entity.Token) error {
	panic("not implemented")
}
func (c *Usecase) NewCollectionNotify(collection entity.Collections) error {
	panic("not implemented")
}
func (c *Usecase) NewNameNotify(airdrop entity.Bns) error {
	panic("not implemented")
}
func (c *Usecase) NewArtifactNotify(airdrop entity.Nfts) error {
	panic("not implemented")
}

func (c *Usecase) JobSendDiscord() error {
	logger.AtLog.Logger.Info("JobSendDiscord Starting ...")
	for page := int64(1); ; page++ {

		notifications, err := c.Repo.FindDiscordNotifications(context.TODO(), entity.GetDiscordReq{
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
			if err := c.DiscordClient.SendMessage(context.TODO(), notify.Webhook, *discordMsg); err != nil {
				c.Repo.UpdateDiscord(context.TODO(), notify.ID.String(), map[string]interface{}{
					"num_retried": notify.NumRetried + 1,
				})

				if notify.NumRetried+1 == entity.MaxSendDiscordRetryTimes {
					c.Repo.UpdateDiscord(context.TODO(), notify.ID.String(), map[string]interface{}{
						"status": entity.FAILED,
						"note":   fmt.Sprintf("failed after %d times", entity.MaxSendDiscordRetryTimes),
					})
				}
			} else {
				c.Repo.UpdateDiscord(context.TODO(), notify.ID.String(), map[string]interface{}{
					"status": entity.DONE,
					"note":   "messaged is sent at " + time.Now().Format(time.RFC3339),
				})
			}
		}
	}

	return nil
}

func (c *Usecase) CreateDiscordNotify(notify *entity.DiscordNotification) error {
	partners, err := c.Repo.GetAllDiscordPartner()
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
			c.Repo.CreateDiscordNotification(context.TODO(), notify)
		}
	}

	return nil
}

func (c *Usecase) TestSendNotify() {
	domain := os.Getenv("DOMAIN")
	if domain == "https://devnet.generative.xyz" {
		c.JobSendDiscord()
		fmt.Println("done")
	}
}
