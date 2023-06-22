package utils

import (
	"time"

	"github.com/jinzhu/now"
)

type QuerySort struct {
	Sort   int
	SortBy string
}

type AggregateDexBTCListing struct {
	FromDate time.Time
	ToDate   time.Time
}

func ParseAggregation(key string) AggregateDexBTCListing {
	sortParams := make(map[string]AggregateDexBTCListing)
	to := time.Now().UTC()
	sortParams["week"] = AggregateDexBTCListing{FromDate: now.BeginningOfDay().AddDate(0, 0, -7), ToDate: to}
	sortParams["month"] = AggregateDexBTCListing{FromDate: now.BeginningOfDay().AddDate(0, 0, -30), ToDate: to}
	filter, ok := sortParams[key]
	if !ok {
		return sortParams["custom"]
	}
	return filter
}

func ParseSort(key string) QuerySort {
	sortParams := make(map[string]QuerySort)
	sortParams["custom"] = QuerySort{Sort: -1, SortBy: ""}
	sortParams["newest"] = QuerySort{Sort: -1, SortBy: "created_at"}
	sortParams["oldest"] = QuerySort{Sort: 1, SortBy: "created_at"}
	sortParams["priority-asc"] = QuerySort{Sort: 1, SortBy: "priority"}
	sortParams["priority-desc"] = QuerySort{Sort: -1, SortBy: "priority"}
	sortParams["price-asc"] = QuerySort{Sort: 1, SortBy: "price"}
	sortParams["price-desc"] = QuerySort{Sort: -1, SortBy: "price"}
	sortParams["minted-newest"] = QuerySort{Sort: -1, SortBy: "minted_time"}
	sortParams["token-price-desc"] = QuerySort{Sort: -1, SortBy: "stats.price_int"}
	sortParams["token-price-asc"] = QuerySort{Sort: 1, SortBy: "stats.price_int"}
	sortParams["trending-score"] = QuerySort{Sort: -1, SortBy: "stats.trending_score"}

	sort, ok := sortParams[key]
	if !ok {
		return sortParams["custom"]
	}

	return sort
}

func ParseSortNew(key string) QuerySort {
	sortParams := make(map[string]QuerySort)
	sortParams["custom"] = QuerySort{Sort: -1, SortBy: ""}
	sortParams["newest"] = QuerySort{Sort: -1, SortBy: "created_at"}
	sortParams["oldest"] = QuerySort{Sort: 1, SortBy: "created_at"}
	sortParams["priority-asc"] = QuerySort{Sort: 1, SortBy: "priority"}
	sortParams["priority-desc"] = QuerySort{Sort: -1, SortBy: "priority"}
	sortParams["price-asc"] = QuerySort{Sort: 1, SortBy: "priceBTC"}
	sortParams["price-desc"] = QuerySort{Sort: -1, SortBy: "priceBTC"}
	sortParams["minted-newest"] = QuerySort{Sort: -1, SortBy: "minted_time"}
	sortParams["token-price-desc"] = QuerySort{Sort: -1, SortBy: "stats.price_int"}
	sortParams["token-price-asc"] = QuerySort{Sort: 1, SortBy: "stats.price_int"}
	sortParams["trending-score"] = QuerySort{Sort: -1, SortBy: "stats.trending_score"}

	sort, ok := sortParams[key]
	if !ok {
		return sortParams["custom"]
	}

	return sort
}

const (
	MAX_CHECK_BALANCE int    = 15
	PERCENT_EARNING   int    = 900
	API_KEY           string = "Api-Key"
	//AUTH_TOKEN           string = "Authorization" //token will be save in this variable
	AUTH_TOKEN                string = "Authorization" //token will be save in this variable
	REDIS_VERIFIED_TOKEN      string = "verified_token"
	REDIS_PROFILE             string = "profile"
	REDIS_INSCRIPTION         string = "inscription"
	REDIS_NFT_METADATA_KEY    string = "nfts_metadata_%s_%s"
	REDIS_PAGINATION_KEY      string = "pagination_%s"
	REDIS_PAGINATION_DATA_KEY string = "pagination_data_%s"
	SIGNED_USER_ID            string = "signed_user_id"
	SIGNED_ADMIN_KEY          string = "admin_user_id_%s"
	SIGNED_WALLET_ADDRESS     string = "signed_wallet_address"
	SIGNED_EMAIL              string = "signed_email"
	SERVICE_API_KEY           string = "service_key"
	TRACER_EMAIL              string = "email"
	XRECAPTCHA                string = "x-recaptcha"
	RECAPTCHA                 string = "recaptcha"

	CODE_LENGTH         int    = 3
	WORK_STATION        string = "working_place"
	WORK_STATION_PREFIX        = "SD4"
	OTHER_TYPE          string = "other"
	MODIFIED_TOKEN      string = "modified-token-%s"
	VERIFY_TOKEN        string = "verify-token-%s"

	EMAIL_TAG              string = "email"
	TOKEN_ID_TAG           string = "tokenID"
	PROJECT_ID_TAG         string = "projectID"
	WALLET_ADDRESS_TAG     string = "wallet_address"
	ORD_WALLET_ADDRESS_TAG string = "ord_wallet_address"
	GEN_NFT_ADDRESS_TAG    string = "gen_nft_address"

	PubsubCreateDeviceType           string = "Device:PubsubCreateDeviceType"
	PubsubUpdateDeviceType           string = "Device:PubsubUpdateDeviceType"
	PubsubDeleteDeviceType           string = "Device:PubsubDeleteDeviceType"
	PubsubSendMessageToSlack         string = "Device:PubsubSendMessageToSlack"
	PUBSUB_SEND_OTP                  string = "Hybrid:SendOtp"
	PUBSUB_REGISTER                  string = "WorkspaceGateway::PubsubRegister"
	PUBSUB_FORGOT_PASSWORD           string = "Hybrid:ResetPasswordEmail"
	NFT_CACHE_EXPIRED_TIME           int    = 86400
	TOKEN_CACHE_EXPIRED_TIME         int    = 86400 * 30       //a month (second)
	REFRESH_TOKEN_CACHE_EXPIRED_TIME int    = 86400 * 360      //a year (second)
	DB_CACHE_EXPIRED_TIME            int    = 86400            //a day
	REDIS_CACHE_EXPIRED_TIME         int    = 86400            //a day
	DB_CACHE_KEY                     string = "db.cache.%s.%s" //a week
	NONCE_MESSAGE_FORMAT             string = "%s"

	KEY_UUID                       string = "uuid"
	KEY_BASE_PRODUCT_KEY           string = "product_key"
	KEY_ORDER_ID                   string = "order_id"
	KEY_AUTO_USERID                string = "user_id"
	KEY_WALLET_ADDRESS             string = "wallet_address"
	KEY_WALLET_ADDRESS_BTC         string = "wallet_address_btc"
	KEY_WALLET_ADDRESS_BTC_TAPROOT string = "wallet_address_btc_taproot"
	KEY_DELETED_AT                 string = "deleted_at"
	KEY_PROJECT_ID                 string = "project_id"
	KEY_LISTING_CONTRACT           string = "collection_contract"
	KEY_BTC_WALLET_INFO            string = "btc_wallet_info"

	COLLECTION_COLLECTIONS                string = "collections"           //ERC721, a collection has many nfts
	COLLECTION_NFTS                       string = "nfts"                  //ERC721
	COLLECTION_SOUL_IMAGE_HISTORIES       string = "soul_image_histories"  //ERC721
	COLLECTION_NFT_AUCTIONS               string = "nft_auction_available" //ERC721
	COLLECTION_BNS                        string = "bns"                   //ERC721
	COLLECTION_BNS_DEFAULT                string = "bns_default"           //ERC721
	VIEW_BNS                              string = "bns_view"              //ERC721
	COLLECTION_NFT_HISTORIES              string = "nft_histories"         //ERC721 - histories
	COLLECTION_ALLOW_LIST                 string = "allow_list"
	COLLECTION_INSERTED_ALLOW_LIST        string = "inserted_allow_list"
	COLLECTION_FAUCETS                    string = "faucets"
	COLLECTION_MARKETPLACE_LISTING        string = "marketplace_listings"
	COLLECTION_MARKETPLACE_OFFER          string = "marketplace_offers"
	COLLECTION_MARKETPLACE_TOKEN_ACTIVITY string = "marketplace_token_activities"

	COLLECTION_TOKEN_URI            string = "token_uris"
	COLLECTION_TOKENS               string = "tokens"
	COLLECTION_USERS                string = "users"
	COLLECTION_USER_HISTORIES       string = "user_histories"
	COLLECTION_UPLOADED_FILES       string = "uploaded_files"
	COLLECTION_UPLOADED_FILE_CHUNKS string = "uploaded_file_chunks"

	COLLECTION_SWAP_CONFIGS                string = "swap_configs"
	COLLECTION_SWAP_FE_LOGS                string = "swap_fe_logs"
	COLLECTION_SWAP_PAIR                   string = "swap_pair"
	COLLECTION_SWAP_PAIR_SYNC              string = "swap_pair_sync"
	COLLECTION_SWAP_PAIR_EVENTS            string = "swap_pair_events"
	COLLECTION_SWAP_HISTORIES              string = "swap_swap_histories"
	COLLECTION_SWAP_IDO                    string = "swap_swap_ido"
	COLLECTION_SWAP_PENDING_TRANSACTION    string = "swap_pending_transaction"
	COLLECTION_SWAP_TOKEN_TRANSFER_HISTORY string = "swap_token_transfer_histories"
	COLLECTION_SWAP_WALLET_ADDRESS         string = "swap_wallet_address"
	COLLECTION_SWAP_BOT_CONFIG             string = "swap_bot_config"
	COLLECTION_SWAP_BOT_TRANSACTION        string = "swap_bot_transaction"
	COLLECTION_SWAP_USER_GM_BALANCE        string = "swap_user_gm_balance"
	COLLECTION_SWAP_USER_GM_PAID           string = "swap_user_gm_paid"
	COLLECTION_SWAP_BLACKLIST_TOKENS       string = "swap_blacklist_tokens"

	VIEW_MARKETPLACE_NFT_WITH_ATTRIBUTES           string = "marketplace_nft_with_attributes_view"
	VIEW_SWAP_LIQUIDITY_REPORT_SLACK               string = "swap_liquidity_report_view"
	VIEW_SWAP_REPORT_SLACK                         string = "swap_token_report_slack_view"
	VIEW_SWAP_REPORT_FINAL                         string = "swap_token_report_final_view"
	VIEW_SWAP_IDO_TOKEN                            string = "swap_ido_tokens_view"
	VIEW_SWAP_PAIR_VOLUME                          string = "swap_pair_volume_view"
	VIEW_SWAP_PAIR_CURRENT_RESERVE                 string = "swap_pair_current_reserve_view"
	VIEW_SWAP_PAIR_APR                             string = "swap_pair_apr_view"
	VIEW_MARKETPLACE_COLLECTION_ATTRIBUTES         string = "marketplace_collection_attributes_view"
	VIEW_MARKETPLACE_COLLECTION_ATTRIBUTES_PERCENT string = "marketplace_collection_attribute_percent_view"
	VIEW_MARKETPLACE_AGGREGATED_COLLECTIONS        string = "marketplace_aggregated_collections_view"
	VIEW_MARKETPLACE_NFTS                          string = "marketplace_nfts_view"
	VIEW_NEW_MARKETPLACE_NFTS                      string = "marketplace_new_nfts_view"
	VIEW_SOUL_MARKETPLACE_NFTS_AUCTION_RARITY      string = "soul_nfts_with_aution_data_rarity_view"
	VIEW_NFT_AUCTION_AVAILABLE                     string = "nft_auction_available_view"
	VIEW_MARKETPLACE_COLLECTION_CHART              string = "marketplace_collection_chart"
	VIEW_MARKETPLACE_COUNT_COLLECTION_OWNER        string = "marketplace_count_collection_owner"
	COLLECTION_MARKETPLACE_AGGREGATED_COLLECTIONS  string = "marketplace_aggregated_collections"
	VIEW_NFTS_WITH_SIZE                            string = "nft_with_size_view"
	TRANSACTIONS                                   string = "transactions"

	REDIS_KEY_LOCK_TX_CONSUMER_CONSUMER_BLOCK string = "lock-tx-consumer-update-last-processed-block"
	EVM_NULL_ADDRESS                          string = "0x0000000000000000000000000000000000000000"
	PUBSUB_TOKEN_THUMBNAIL                    string = "token_thumbnail"
	PUBSUB_PROJECT_UNZIP                      string = "project_unzip"

	BTCConfirmationThreshold = 1
	FirstScannedBTCBlkHeight = 697200
	BUY_NFT_CHARGE           = 0      // 0%
	MIN_BTC_TO_LIST_BTC      = 500000 // 0.005 btc

	FEE_BTC_SEND_AGV = 8000 // fee send btc
	MIN_FILE_SIZE    = 4096 // min file size (for linux system)

	FEE_ETH_SEND_MASTER = 0.0007
	FEE_BTC_SEND_NFT    = 10000

	DEVELOPER_INSCRIBE_MAX_REQUEST = 200

	INSCRIBE_TIMEOUT = 3

	MASTER_ADDRESS = "bc1p8ts7h86jgduat5v98cwlurngeyasqrd5c6ch2my8qwen3ykpagyswv2sy8"

	NETWORK_BTC = "btc"
	NETWORK_ETH = "eth"

	PLATFORM_ORDINAL = "ordinal"
	PLATFORM_TC      = "tc"

	AIRDROP_MAGIC  = "https://storage.googleapis.com/generative-static-prod/airdrop/magickey.html"
	AIRDROP_GOLDEN = "https://storage.googleapis.com/generative-static-prod/airdrop/goldenkey.html"
	AIRDROP_SILVER = "https://storage.googleapis.com/generative-static-prod/airdrop/silverkey.html"

	PAGINATION = "pagination"

	COLLECTION_DISCORD_NOTIFICATION = "discord_notifications"
	COLLECTION_DISCORD_PARTNERS     = "discord_partners"
)

type PubSubSendOtp struct {
	Email   string `json:"email"`
	Code    string `json:"code"`
	AppName string `json:"app_name"`
}

const HttpContextTimeOut = time.Second * 15
