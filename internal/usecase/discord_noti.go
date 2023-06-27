package usecase

import (
	"context"
	"dapp-moderator/external/bns_service"
	"dapp-moderator/utils"
	discordclient "dapp-moderator/utils/discord"
	"dapp-moderator/utils/helpers"
	"encoding/json"
	"errors"
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
							Value:  fmt.Sprintf("**Collection Name: [%s](https://trustlessnfts.com/collection/%s)**", utils.NameOrAddress(collection.Name, collection.Contract), collection.Contract),
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
		Content:   fmt.Sprintf("**NEW SMART INSCRIPTION**"),
		Username:  "Satoshi 27",
		AvatarUrl: "",
		Embeds: []discordclient.Embed{
			{
				Fields: []discordclient.Field{
					{
						Value:  fmt.Sprintf("**Owner:** \n%s", utils.ShortenBlockAddress(nfts.Owner)),
						Inline: true,
					},
					{
						Value:  fmt.Sprintf("**Type:** \n%s", nfts.ContentType),
						Inline: true,
					},
				},
			},
		},
	}
	if nfts.Image != "" {
		if strings.HasPrefix(nfts.Image, "/dapp/api/nft-explorer/collections/") {
			message.Embeds[0].Image.Url = "https://dapp.trustless.computer" + nfts.Image
		} else {
			message.Embeds[0].Image.Url = nfts.Image
		}

	}
	notify := &entity.DiscordNotification{
		Message: message,
		Status:  entity.PENDING,
		Event:   entity.EventNewArtifact,
	}

	if nfts.Image != "" && strings.HasPrefix(nfts.Image, "/dapp/api/nft-explorer/collections/") {
		notify.Message.Embeds[0].Image.Url = "https://dapp.trustless.computer" + nfts.Image
	} else {
		notify.Message.Embeds[0].Image.Url = nfts.Image
	}
	notify.Message.Embeds[0].Title = fmt.Sprintf("Smart Inscription #%s", nfts.TokenID)
	notify.Message.Embeds[0].Url = fmt.Sprintf("https://smartinscription.xyz/%s", nfts.TokenID)

	return u.CreateDiscordNotify(notify)
}

// SOUL notifications
func (u *Usecase) NewAuctionCreatedNotify(auction *entity.Auction) (*entity.DiscordNotification, error) {
	nft, err := u.Repo.GetNft(auction.CollectionAddress, auction.TokenID)
	if err != nil {
		return nil, err
	}

	message := discordclient.Message{
		Content:   fmt.Sprintf("**Create Adopt**"),
		Username:  "Satoshi 27",
		AvatarUrl: "",
		Embeds: []discordclient.Embed{
			{
				Fields: []discordclient.Field{
					{
						Value:  fmt.Sprintf("**Start Block:** \n%s", auction.StartTimeBlock),
						Inline: true,
					},
					{
						Value:  fmt.Sprintf("**End Block:** \n%s", auction.EndTimeBlock),
						Inline: true,
					},
				},
			},
		},
	}

	if nft.ImageCapture != "" {
		message.Embeds[0].Image.Url = nft.ImageCapture
	}

	notify := &entity.DiscordNotification{
		Message: message,
		Status:  entity.PENDING,
		Event:   entity.EventAuctionCreated,
	}

	notify.Message.Embeds[0].Title = fmt.Sprintf("Soul #%s", nft.TokenID)
	notify.Message.Embeds[0].Url = fmt.Sprintf("%s/souls/%s", os.Getenv("SOUL_DOMAIN"), nft.TokenID)

	if os.Getenv("ENV") == "production" {
		err = u.CreateDiscordNotify(notify)
		if err != nil {
			return nil, err
		}
	}

	return notify, nil
}

func (u *Usecase) NewAuctionSettledNotify(auction *entity.Auction) (*entity.DiscordNotification, error) {
	nft, err := u.Repo.GetNft(auction.CollectionAddress, auction.TokenID)
	if err != nil {
		return nil, err
	}

	wn := ""
	if auction.Winner == nil {
		return nil, errors.New("Auction doesn't have winner")
	}

	wn = *auction.Winner

	message := discordclient.Message{
		Content:   fmt.Sprintf("**Settle**"),
		Username:  "Satoshi 27",
		AvatarUrl: "",
		Embeds: []discordclient.Embed{
			{
				Fields: []discordclient.Field{
					{
						Value:  fmt.Sprintf("**Winner:** \n%s", utils.ShortenBlockAddress(wn)),
						Inline: true,
					},
					{
						Value:  fmt.Sprintf("**Amount:** \n%.5f", helpers.GetValue(auction.TotalAmount, 18)),
						Inline: true,
					},
				},
			},
		},
	}

	if nft.ImageCapture != "" {
		message.Embeds[0].Image.Url = nft.ImageCapture
	}

	notify := &entity.DiscordNotification{
		Message: message,
		Status:  entity.PENDING,
		Event:   entity.EventAuctionSettled,
	}

	notify.Message.Embeds[0].Title = fmt.Sprintf("Soul #%s", nft.TokenID)
	notify.Message.Embeds[0].Url = fmt.Sprintf("%s/souls/%s", os.Getenv("SOUL_DOMAIN"), nft.TokenID)

	if os.Getenv("ENV") == "production" {
		err = u.CreateDiscordNotify(notify)
		if err != nil {
			return nil, err
		}
	}

	return notify, nil
}

func (u *Usecase) NewBidCreatedNotify(auctionBid *entity.AuctionBid) (*entity.DiscordNotification, error) {
	nft, err := u.Repo.GetNft(auctionBid.CollectionAddress, auctionBid.TokenID)
	if err != nil {
		return nil, err
	}

	message := discordclient.Message{
		Content:   fmt.Sprintf("**Create Bid**"),
		Username:  "Satoshi 27",
		AvatarUrl: "",
		Embeds: []discordclient.Embed{
			{
				Fields: []discordclient.Field{
					{
						Value:  fmt.Sprintf("**Sender:** \n%s", utils.ShortenBlockAddress(auctionBid.Sender)),
						Inline: true,
					},
					{
						Value:  fmt.Sprintf("**Amount:** \n%.5f GM", helpers.GetValue(auctionBid.Amount, 18)),
						Inline: true,
					},
				},
			},
		},
	}

	if nft.ImageCapture != "" {
		message.Embeds[0].Image.Url = nft.ImageCapture
	}

	notify := &entity.DiscordNotification{
		Message: message,
		Status:  entity.PENDING,
		Event:   entity.EventBidCreated,
	}

	notify.Message.Embeds[0].Title = fmt.Sprintf("Soul #%s", nft.TokenID)
	notify.Message.Embeds[0].Url = fmt.Sprintf("%s/souls/%s", os.Getenv("SOUL_DOMAIN"), nft.TokenID)

	if os.Getenv("ENV") == "production" {
		err = u.CreateDiscordNotify(notify)
		if err != nil {
			return nil, err
		}
	}

	return notify, nil
}

// it's disabled by order
func (u *Usecase) NewMintTokenNotify(nfts *entity.Nfts) error {
	message := discordclient.Message{
		Content:   fmt.Sprintf("**MINT**"),
		Username:  "Satoshi 27",
		AvatarUrl: "",
		Embeds: []discordclient.Embed{
			{
				Fields: []discordclient.Field{
					{
						Value:  fmt.Sprintf("**Owner:** \n%s", utils.ShortenBlockAddress(nfts.Owner)),
						Inline: true,
					},
					{
						Value:  fmt.Sprintf("**Type:** \n%s", nfts.ContentType),
						Inline: true,
					},
				},
			},
		},
	}
	if nfts.Image != "" {
		if strings.HasPrefix(nfts.Image, "/dapp/api/nft-explorer/collections/") {
			message.Embeds[0].Image.Url = "https://dapp.trustless.computer" + nfts.Image
		} else {
			message.Embeds[0].Image.Url = nfts.Image
		}

	}
	notify := &entity.DiscordNotification{
		Message: message,
		Status:  entity.PENDING,
		Event:   entity.EventNewArtifact,
	}

	if nfts.Image != "" && strings.HasPrefix(nfts.Image, "/dapp/api/nft-explorer/collections/") {
		notify.Message.Embeds[0].Image.Url = "https://dapp.trustless.computer" + nfts.Image
	} else {
		notify.Message.Embeds[0].Image.Url = nfts.Image
	}
	notify.Message.Embeds[0].Title = fmt.Sprintf("Smart Inscription #%s", nfts.TokenID)
	notify.Message.Embeds[0].Url = fmt.Sprintf("https://smartinscription.xyz/%s", nfts.TokenID)

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

				err = u.Repo.UpdateDiscord(context.TODO(), notify.Id(), map[string]interface{}{
					"num_retried": notify.NumRetried + 1,
				})

				if err != nil {
					logger.AtLog.Logger.Error(fmt.Sprintf("Send discord message failed - %s", notify.Message.Content), zap.Error(err))
				}

				if notify.NumRetried+1 == entity.MaxSendDiscordRetryTimes {
					err = u.Repo.UpdateDiscord(context.TODO(), notify.Id(), map[string]interface{}{
						"status": entity.FAILED,
						"note":   fmt.Sprintf("failed after %d times", entity.MaxSendDiscordRetryTimes),
					})
					if err != nil {
						logger.AtLog.Logger.Error(fmt.Sprintf("Send discord message failed - %s", notify.Message.Content), zap.Error(err))
					}
				}
			} else {
				err = u.Repo.UpdateDiscord(context.TODO(), notify.Id(), map[string]interface{}{
					"status": entity.DONE,
					"note":   "messaged is sent at " + time.Now().Format(time.RFC3339),
				})
				if err != nil {
					logger.AtLog.Logger.Error(fmt.Sprintf("Send discord message failed - %s", notify.Message.Content), zap.Error(err))
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

func (u *Usecase) NewListForSaleNotify(listing *entity.MarketplaceListings) (*entity.DiscordNotification, error) {
	logger.AtLog.Logger.Info(fmt.Sprintf("NewListForSaleNotify %s %s", listing.CollectionContract, listing.TokenId), zap.String("offering id", listing.OfferingId))

	nfts, err := u.Repo.GetNft(listing.CollectionContract, listing.TokenId)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("NewListForSaleNotify %s %s", listing.CollectionContract, listing.TokenId), zap.Error(err), zap.String("offering id", listing.OfferingId))
		return nil, err
	}

	collection, err := u.Repo.GetCollection(nfts.ContractAddress)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("NewListForSaleNotify %s %s", listing.CollectionContract, listing.TokenId), zap.Error(err), zap.String("offering id", listing.OfferingId))
		return nil, err
	}

	p := helpers.GetValue(listing.Price, 18)
	r := listing.Erc20Token
	token := ""
	if strings.ToLower(r) == strings.ToLower(os.Getenv("WETH_ADDRESS")) {
		token = "WETH"
	} else if strings.ToLower(r) == strings.ToLower(os.Getenv("WBTC_ADDRESS")) {
		token = "WBTC"
	} else {
		return nil, errors.New("Cannot detect ERC20")
	}

	owner := listing.Seller
	price := fmt.Sprintf("%f %s", p, token)

	message := discordclient.Message{
		Content:   fmt.Sprintf("**NEW LISTING**"),
		Username:  "Satoshi 27",
		AvatarUrl: "",
		Embeds: []discordclient.Embed{
			{
				Fields: []discordclient.Field{
					{
						Value:  fmt.Sprintf("**List Price:** \n%s", price),
						Inline: true,
					},
					{
						Value:  fmt.Sprintf("**Seller:** \n%s", utils.ShortenBlockAddress(owner)),
						Inline: true,
					},
				},
			},
		},
	}

	notify := &entity.DiscordNotification{
		Message: message,
		Status:  entity.PENDING,
		Event:   entity.EventListForSale,
	}

	image := fmt.Sprintf("https://dapp.trustless.computer/dapp/api/nft-explorer/collections/%s/nfts/%s/content", nfts.ContractAddress, nfts.TokenID)
	notify.Message.Embeds[0].Image.Url = u.ParseSvgImage(image)

	notify.Message.Embeds[0].Title = fmt.Sprintf("%s #%s", collection.Name, nfts.TokenID)
	notify.Message.Embeds[0].Url = fmt.Sprintf("https://trustlessnfts.com/collection/%s/token/%s", nfts.ContractAddress, nfts.TokenID)

	if os.Getenv("ENV") == "production" {
		err = u.CreateDiscordNotify(notify)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("ListForSaleNotify %s %s", listing.CollectionContract, listing.TokenId), zap.Error(err), zap.String("offering id", listing.OfferingId))
			return nil, err
		}
	}

	return notify, nil
}

func (u *Usecase) NewPurchaseTokenNotify(offeringID string) (*entity.DiscordNotification, error) {
	logger.AtLog.Logger.Info(fmt.Sprintf("NewPurchaseTokenNotify %s", offeringID), zap.String("offeringID", offeringID))

	ml, err := u.Repo.GetMarketplaceListing(offeringID)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("NewListForSaleNotify %s %s", ml.CollectionContract, ml.TokenId), zap.Error(err), zap.String("offering id", ml.OfferingId))
		return nil, err
	}

	nfts, err := u.Repo.GetNft(ml.CollectionContract, ml.TokenId)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("NewListForSaleNotify %s %s", ml.CollectionContract, ml.TokenId), zap.Error(err), zap.String("offering id", ml.OfferingId), zap.Any("nfts", nfts))
		return nil, err
	}

	collection, err := u.Repo.GetCollection(nfts.ContractAddress)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("NewListForSaleNotify %s %s", ml.CollectionContract, ml.TokenId), zap.Error(err), zap.String("offering id", ml.OfferingId), zap.Any("nfts", nfts))
		return nil, err
	}

	mkpActivities, err := u.Repo.PurchaseMKPActivity(offeringID)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("NewListForSaleNotify %s %s", ml.CollectionContract, ml.TokenId), zap.Error(err), zap.String("offering id", ml.OfferingId), zap.Any("nfts", nfts))
		return nil, err
	}

	p := helpers.GetValue(ml.Price, 18)
	r := ml.Erc20Token
	token := ""
	if strings.ToLower(r) == strings.ToLower(os.Getenv("WETH_ADDRESS")) {
		token = "WETH"
	} else if strings.ToLower(r) == strings.ToLower(os.Getenv("WBTC_ADDRESS")) {
		token = "WBTC"
	} else {
		return nil, errors.New("Cannot detect ERC20")
	}

	price := fmt.Sprintf("%f %s", p, token)

	message := discordclient.Message{
		Content:   fmt.Sprintf("**NEW SALE**"),
		Username:  "Satoshi 27",
		AvatarUrl: "",
		Embeds: []discordclient.Embed{
			{
				Fields: []discordclient.Field{
					{
						Value:  fmt.Sprintf("**Sale Price:** \n%s", price),
						Inline: true,
					},
					{
						Value:  fmt.Sprintf("**Buyer:** \n%s", utils.ShortenBlockAddress(mkpActivities.UserBAddress)),
						Inline: true,
					},
					{
						Value:  fmt.Sprintf("**Seller:** \n%s", utils.ShortenBlockAddress(mkpActivities.UserAAddress)),
						Inline: true,
					},
				},
			},
		},
	}

	notify := &entity.DiscordNotification{
		Message: message,
		Status:  entity.PENDING,
		Event:   entity.EventPurchaseListing,
	}

	image := fmt.Sprintf("https://dapp.trustless.computer/dapp/api/nft-explorer/collections/%s/nfts/%s/content", nfts.ContractAddress, nfts.TokenID)
	notify.Message.Embeds[0].Image.Url = u.ParseSvgImage(image)

	notify.Message.Embeds[0].Title = fmt.Sprintf("%s #%s", collection.Name, nfts.TokenID)
	notify.Message.Embeds[0].Url = fmt.Sprintf("https://trustlessnfts.com/collection/%s/token/%s", nfts.ContractAddress, nfts.TokenID)

	if os.Getenv("ENV") == "production" {
		err = u.CreateDiscordNotify(notify)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("ListForSaleNotify %s %s", ml.CollectionContract, ml.TokenId), zap.Error(err), zap.String("offering id", ml.OfferingId))
			return nil, err
		}
	}

	return notify, nil
}

func (u *Usecase) ParseSvgImage(imageURL string) string {
	parseImageUrl := "https://devnet.generative.xyz/generative/api/photo/pare-svg"

	postData := make(map[string]interface{})
	postData["display_url"] = imageURL
	postData["delay_time"] = 1
	postData["app_id"] = "dapp"

	resp, _, _, err := helpers.HttpRequest(parseImageUrl, "POST", make(map[string]string), postData)
	if err != nil {
		return imageURL
	}

	type respdata struct {
		Err    error  `json:"error"`
		Status bool   `json:"status"`
		Data   string `json:"data"`
	}

	response := &respdata{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return imageURL
	}

	if !response.Status {
		return imageURL
	}

	if response.Err != nil {
		return imageURL
	}

	return strings.ReplaceAll(response.Data, "https", "http")
}

func (u *Usecase) ParseHtmlImage(imageURL string) (string, map[string]string, error) {
	parseImageUrl := "https://devnet.generative.xyz/generative/api/photo/pare-html"

	now := time.Now().UTC().UnixNano()
	postData := make(map[string]interface{})
	postData["display_url"] = imageURL
	postData["delay_time"] = 20
	postData["app_id"] = fmt.Sprintf("dapp-%d", now)

	resp, _, _, err := helpers.HttpRequest(parseImageUrl, "POST", make(map[string]string), postData)
	if err != nil {
		return "", nil, err
	}

	type respdata struct {
		Err    error `json:"error"`
		Status bool  `json:"status"`
		Data   struct {
			Image  string            `json:"image"`
			Traits map[string]string `json:"traits"`
		} `json:"data"`
	}

	response := &respdata{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return "", nil, err
	}

	if !response.Status {
		return "", nil, err
	}

	if response.Err != nil {
		return "", nil, err
	}

	return response.Data.Image, response.Data.Traits, nil
}

func (u *Usecase) TestSendNotify() {
	env := os.Getenv("ENVIRONMENT")
	if env == "production" {
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
