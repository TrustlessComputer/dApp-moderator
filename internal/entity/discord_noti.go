package entity

import (
	"dapp-moderator/utils"
	discordclient "dapp-moderator/utils/discord"
	"dapp-moderator/utils/helpers"
	"go.mongodb.org/mongo-driver/bson"
)

type DiscordStatus int
type DiscordEvent string
type DiscordImageSourceType int
type DiscordImagePosition int

const (
	PENDING DiscordStatus = 0
	DONE    DiscordStatus = 1
	FAILED  DiscordStatus = 2

	EventNewCollection   DiscordEvent = "new_collection"
	EventNewArtifact     DiscordEvent = "new_artifact"
	EventNewToken        DiscordEvent = "new_token"
	EventNewName         DiscordEvent = "new_name"
	EventListForSale     DiscordEvent = "list_for_sale"
	EventPurchaseListing DiscordEvent = "purchase_listing"

	ImageFromInscriptionID DiscordImageSourceType = 1
	ThumbNailPosition      DiscordImagePosition   = 1
	FullImagePosition      DiscordImagePosition   = 2

	MaxSendDiscordRetryTimes = 3
	MaxFindImageRetryTimes   = 20
)

type GetDiscordReq struct {
	Page   int64
	Limit  int64
	Status DiscordStatus
}

type DiscordMeta struct {
	SendTo string `bson:"send_to"`
}

type DiscordNotification struct {
	BaseEntity `bson:",inline"`
	UUID       string
	Message    discordclient.Message `bson:"message"` // message to send, base on client

	Status     DiscordStatus `bson:"status"`
	NumRetried int           `bson:"num_retried"`
	Webhook    string        `bson:"webhook"`
	Event      DiscordEvent  `bson:"event"`
	Meta       DiscordMeta   `bson:"meta"`

	RequireImage    bool                   `bson:"require_image"`
	ImageSourceID   string                 `bson:"image_source_id"`
	Note            string                 `bson:"note"`
	ImageSourceType DiscordImageSourceType `bson:"image_source_type"`
	ImagePosition   DiscordImagePosition   `bson:"image_position"`
}

func (u *DiscordNotification) TableName() string {
	return utils.COLLECTION_DISCORD_NOTIFICATION
}

func (u *DiscordNotification) ToBson() (*bson.D, error) {
	return helpers.ToDoc(u)
}

type DiscordPartner struct {
	BaseEntity            `bson:",inline"`
	Webhooks              map[string]string `bson:"webhooks"`
	ProjectIDs            []string          `bson:"project_ids"`
	Name                  string            `bson:"name"`
	AmountGreaterThanZero bool              `bson:"greater_than_zero"`
	Categories            []string          `bson:"categories"`
}

func (u *DiscordPartner) TableName() string {
	return utils.COLLECTION_DISCORD_PARTNERS
}

func (u *DiscordPartner) ToBson() (*bson.D, error) {
	return helpers.ToDoc(u)
}

func (u *DiscordPartner) MatchProject(t string) bool {
	if len(u.ProjectIDs) == 0 {
		return true
	}

	for _, p := range u.ProjectIDs {
		if p == t {
			return true
		}
	}

	return false
}

func (u *DiscordPartner) MatchCategory(t string) bool {
	if len(u.Categories) == 0 {
		return true
	}

	for _, p := range u.Categories {
		if p == t {
			return true
		}
	}

	return false
}

func (u *DiscordPartner) MatchAmountGreaterThanZero(amount uint64) bool {
	if !u.AmountGreaterThanZero {
		return true
	}

	if u.AmountGreaterThanZero && amount > 0 {
		return true
	}

	return false
}
