package usecase

import (
	"context"
	"crypto/rand"
	"dapp-moderator/internal/delivery/http/request"
	"dapp-moderator/internal/entity"
	"dapp-moderator/internal/usecase/structure"
	"dapp-moderator/utils"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"dapp-moderator/utils/oauth2service"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (u Usecase) GenerateMessage(ctx context.Context, data *structure.GenerateMessage) (*string, error) {
	addrr := data.Address
	addrr = strings.ToLower(addrr)

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		logger.AtLog.Error("GenerateMessage", zap.String("walletAddress", data.Address), zap.String("WalletType", data.WalletType), zap.Error(err))
		return nil, err
	}
	message := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	message = fmt.Sprintf(utils.NONCE_MESSAGE_FORMAT, message)
	now := time.Now().UTC()
	_, err = u.Repo.FindUserByWalletAddress(addrr)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			//insert
			user := &entity.Users{}
			user.WalletType = data.WalletType
			user.WalletAddress = addrr
			user.Message = message
			user.CreatedAt = &now

			_, err = u.Repo.InsertOne(user)
			if err != nil {
				logger.AtLog.Error("GenerateMessage", zap.String("walletAddress", data.Address), zap.String("WalletType", data.WalletType), zap.Error(err))
				return nil, err
			}

			return &message, nil

		} else {
			logger.AtLog.Error("GenerateMessage", zap.String("walletAddress", data.Address), zap.String("WalletType", data.WalletType), zap.Error(err))
			return nil, err
		}
	}

	_, err = u.Repo.UpdateUserMessage(addrr, message)
	if err != nil {
		logger.AtLog.Error("GenerateMessage", zap.String("walletAddress", data.Address), zap.String("WalletType", data.WalletType), zap.Error(err))
		return nil, err
	}

	return &message, nil
}

func (u Usecase) VerifyMessage(ctx context.Context, data *structure.VerifyMessage) (*structure.VerifyResponse, error) {
	logger.AtLog.Info("VerifyMessage", zap.Any("walletAddress", data.Address))
	if data.Signature == "" {
		return nil, errors.New("invalid params: Signature")
	}

	if data.Address == "" {
		return nil, errors.New("invalid params: Address")
	}

	addrr := strings.ToLower(data.Address)
	user, err := u.Repo.FindUserByWalletAddress(addrr)
	if err != nil {
		logger.AtLog.Error("VerifyMessage", zap.Any("walletAddress", data.Address), zap.Error(err))
		return nil, err
	}
	userID := user.ID.Hex()
	isVeried, err := u.verify(data.Signature, data.Address, user.Message)
	if err != nil {
		logger.AtLog.Error("VerifyMessage", zap.Any("walletAddress", data.Address), zap.Error(err))
		return nil, err
	}

	if !isVeried {
		err := errors.New("Cannot verify wallet address")
		logger.AtLog.Error("VerifyMessage", zap.Any("walletAddress", data.Address), zap.Error(err))
		return nil, err
	}

	token, refreshToken, err := u.Auth2.GenerateAllTokens(user.WalletAddress, "", "", "", userID)
	if err != nil {
		logger.AtLog.Error("VerifyMessage", zap.Any("walletAddress", data.Address), zap.Error(err))
		return nil, err
	}

	tokenMd5 := helpers.GenerateMd5String(token)
	err = u.Cache.SetDataWithExpireTime(tokenMd5, userID, int(utils.TOKEN_CACHE_EXPIRED_TIME))
	if err != nil {
		logger.AtLog.Error("VerifyMessage", zap.Any("walletAddress", data.Address), zap.Error(err))
		return nil, err
	}

	u.Repo.UpdateUserLastLoggedIn(user.WalletAddress)
	verified := structure.VerifyResponse{
		Token:        token,
		RefreshToken: refreshToken,
		IsVerified:   isVeried,
	}

	return &verified, nil
}

func buildMsgETH(taprootAddress, segwitAddress, nonceMessage string) string {
	msg := "GM.\n\nPlease sign this message to confirm your Generative wallet addresses generated by your Ethereum address.\n\n"
	msg += "Taproot address:\n" + taprootAddress
	msg += "\n\nSegwit address:\n" + segwitAddress
	msg += "\n\nNonce:\n" + nonceMessage
	msg += "\n\nThe Generative Core Team"
	return msg
}

func (u Usecase) verifyBTCSegwit(msgStr string, data structure.VerifyMessage) (bool, error) {
	// // verify BTC segwit signature
	// // Reconstruct the pubkey
	// publicKey, wasCompressed, err := helpers.PubKeyFromSignature(data.Signature, msgStr, *data.MessagePrefix)
	// if err != nil {
	// 	return false, err
	// }

	// // Get the address
	// var addressWitnessPubKeyHash *btcutil.AddressPubKeyHash
	// if addressWitnessPubKeyHash, err = helpers.GetAddressFromPubKey(publicKey, wasCompressed); err != nil {
	// 	return false, err
	// }

	// // Return nil if addresses match.
	// temp := addressWitnessPubKeyHash.String()
	// if temp != *data.AddressBTCSegwit {
	// 	return false, fmt.Errorf(
	// 		"address (%s) not found - compressed: %t\n%s was found instead",
	// 		*data.AddressBTCSegwit,
	// 		wasCompressed,
	// 		addressWitnessPubKeyHash.String(),
	// 	)
	// }

	// // verify ETH signature
	// msg2 := buildMsgETH(*data.AddressBTC, *data.AddressBTCSegwit, msgStr)
	// return u.verify(data.ETHSignature, data.Address, msg2)

	//TODO - implement these code lines above
	return true, nil
}

func (u Usecase) verify(signatureHex string, signer string, msgStr string) (bool, error) {
	logger.AtLog.Info("verify", zap.String("signatureHex", signatureHex), zap.String("signer", signer), zap.String("msgStr", msgStr))
	sig := hexutil.MustDecode(signatureHex)

	msgBytes := []byte(msgStr)
	msgHash := accounts.TextHash(msgBytes)

	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	recovered, err := crypto.SigToPub(msgHash, sig)
	if err != nil {
		logger.AtLog.Error(err)
		return false, err
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)
	signerHex := recoveredAddr.Hex()
	isVerified := strings.ToLower(signer) == strings.ToLower(signerHex)

	logger.AtLog.Info("verify", zap.Bool("isVerified", isVerified), zap.String("signerHex", signerHex), zap.String("signatureHex", signatureHex), zap.String("signer", signer), zap.String("msgStr", msgStr), zap.Any("recoveredAddr", recoveredAddr))
	return isVerified, nil
}

func (u Usecase) ValidateAccessToken(accessToken string) (*oauth2service.SignedDetails, error) {

	//tokenMd5 := helpers.GenerateMd5String(accessToken)
	//logger.AtLog.Logger.Info("ValidateAccessToken", zap.String("ValidateAccessToken", zap.Any("accessToken)", accessToken)))

	// userID, err := u.Cache.GetData(tokenMd5)
	// if err != nil {
	// 	err = errors.New("Access token is invaild")
	// 	logger.AtLog.Logger.Error("ValidateAccessToken", zap.String("GetData", accessToken), zap.Error(err))
	// 	return nil, err

	// }

	//Claim wallet Address
	claim, err := u.Auth2.ValidateToken(accessToken)
	if err != nil {
		logger.AtLog.Logger.Error("ValidateAccessToken", zap.String("ValidateToken", accessToken), zap.Error(err))
		return nil, err
	}

	userID := &claim.Uid
	if userID == nil {
		err := errors.New("Cannot find userID")
		logger.AtLog.Logger.Error("ValidateAccessToken", zap.String("userID", accessToken), zap.Error(err))
		return nil, err
	}

	//timeT := time.Unix(claim.ExpiresAt, 0)
	return claim, err
}

func (u Usecase) GetUserProfileByWalletAddress(userAddr string) (*entity.Users, error) {

	logger.AtLog.Info("GetUserProfileByWalletAddress", zap.String("userAddr", userAddr))
	user, err := u.Repo.FindUserByWalletAddress(userAddr)
	if err != nil {
		logger.AtLog.Error("GetUserProfileByBtcAddressTaproot", zap.String("userAddr", userAddr), zap.Error(err))
		return nil, err
	}
	logger.AtLog.Info("GetUserProfileByBtcAddressTaproot", zap.String("userAddr", userAddr), zap.Any("user", user))
	return user, nil
}

func (u Usecase) CreateUserHistory(ctx context.Context, data *structure.CreateHistoryMessage) (*entity.UserHistories, error) {

	logger.AtLog.Info("CreateUserHistory", zap.String("userAddr", data.WalletAddress))
	input := &entity.UserHistories{}
	input.WalletAddress = data.WalletAddress
	input.TxHash = data.TxHash
	input.DappTypeTxHash = data.DappTypeTxHash
	input.FromAddress = data.FromAddress
	input.ToAddress = data.ToAddress
	input.Time = data.Time
	input.Value = data.Value
	input.Decimal = data.Decimal
	input.Status = entity.HISTORY_PENDING
	input.BTCTxHash = data.BTCTxHash

	_, err := u.Repo.InsertOne(input)
	if err != nil {
		logger.AtLog.Error("GetUserProfileByBtcAddressTaproot", zap.String("userAddr", data.WalletAddress), zap.Error(err))
		return nil, err
	}

	logger.AtLog.Info("GetUserProfileByBtcAddressTaproot", zap.String("userAddr", data.WalletAddress), zap.Any("history", input))
	return input, nil
}

func (u Usecase) GetUserHistories(ctx context.Context, filter request.HistoriesFilter) ([]entity.UserHistories, error) {
	res := []entity.UserHistories{}
	f := bson.D{}

	if filter.TxHash != nil && *filter.TxHash != "" {
		f = append(f, bson.E{"tx_hash", primitive.Regex{Pattern: *filter.TxHash, Options: "i"}})
	}

	if filter.WalletAdress != nil && *filter.WalletAdress != "" {
		f = append(f, bson.E{"wallet_address", primitive.Regex{Pattern: *filter.WalletAdress, Options: "i"}})
	}

	sort := 1
	if filter.Sort != nil {
		sort = *filter.Sort
	}

	sortBy := "created_at"
	if filter.SortBy != nil {
		sort = *filter.Sort
	}

	s := bson.D{{"status", -1}, {sortBy, sort}}
	err := u.Repo.Find(utils.COLLECTION_USER_HISTORIES, f, int64(*filter.Limit), int64(*filter.Offset), &res, s)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u Usecase) GetUserProfileByBtcAddressTaproot(userAddr string) (*entity.Users, error) {

	user, err := u.Repo.FindUserByBTCTaprootWalletAddress(userAddr)
	if err != nil {
		logger.AtLog.Error("GetUserProfileByBtcAddressTaproot", zap.String("userAddr", userAddr), zap.Error(err))
		return nil, err
	}
	logger.AtLog.Info("GetUserProfileByBtcAddressTaproot", zap.String("userAddr", userAddr), zap.Any("user", user))
	return user, nil
}

func (u Usecase) ConfirmUserHistory(ctx context.Context, userAddr string, txHashData *request.ConfirmHistoriesReq) ([]entity.UserHistories, error) {

	resp := []entity.UserHistories{}

	for _, el := range txHashData.Data {
		if el.BTCHash == "" {
			// skip empty btc has
			continue
		}
		if el.Status != entity.HISTORY_CONFIRMED && el.Status != entity.HISTORY_PENDING {
			// skip invalid status
			continue
		}
		for _, txHash := range el.TxHash {

			f := bson.D{
				{"wallet_address", userAddr},
				{"tx_hash", txHash},
				{"status", entity.HISTORY_PENDING},
			}

			data, err := u.Repo.FindOne(utils.COLLECTION_USER_HISTORIES, f)
			if err != nil {
				logger.AtLog.Error("ConfirmUserHistory", zap.Any("txHashData", txHashData), zap.String("userAddr", userAddr), zap.Error(err))
				return nil, fmt.Errorf("Cannot find transaction: %s - %v", txHash, err.Error())
			}

			h := &entity.UserHistories{}
			err = data.Decode(h)
			if err != nil {
				logger.AtLog.Error("ConfirmUserHistory", zap.Any("txHashData", txHashData), zap.String("userAddr", userAddr), zap.Error(err))
				return nil, fmt.Errorf("Cannot find transaction: %s - %v", txHash, err.Error())
			}

			h.Status = el.Status
			h.BTCTxHash = el.BTCHash

			_, err = u.Repo.ReplaceOne(f, h)
			if err != nil {
				logger.AtLog.Error("ConfirmUserHistory", zap.Any("txHashData", txHashData), zap.String("userAddr", userAddr), zap.Error(err))
				return nil, fmt.Errorf("Cannot update transaction: %s - %v", txHash, err.Error())
			}
			resp = append(resp, *h)
		}
	}

	logger.AtLog.Info("ConfirmUserHistory", zap.Any("txHashData", txHashData), zap.String("userAddr", userAddr), zap.Any("histories", len(resp)))
	return resp, nil
}
