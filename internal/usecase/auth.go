package usecase

import (
	"context"
	"crypto/rand"
	"dapp-moderator/internal/entity"
	"dapp-moderator/internal/usecase/structure"
	"dapp-moderator/utils"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"go.mongodb.org/mongo-driver/bson"
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

	// validate data
	if data.ETHSignature == "" || data.Signature == "" ||
		data.Address == "" || data.AddressBTC == nil || *data.AddressBTC == "" || data.AddressBTCSegwit == nil || *data.AddressBTCSegwit == "" ||
		data.MessagePrefix == nil || *data.MessagePrefix == "" {
		return nil, errors.New("invalid params")
	}

	addrr := strings.ToLower(data.Address)
	user, err := u.Repo.FindUserByWalletAddress(addrr)
	if err != nil {
		logger.AtLog.Error("VerifyMessage", zap.Any("walletAddress", data.Address), zap.Error(err))
		return nil, err
	}
	userID := user.ID.Hex()
	
	isVeried, err := u.verifyBTCSegwit(user.Message, *data)
	if err != nil {
		logger.AtLog.Error("VerifyMessage", zap.Any("walletAddress", data.Address), zap.Error(err))
		return nil, err
	}
	
	if !isVeried {
		err := errors.New("Cannot verify wallet address")
		logger.AtLog.Error("VerifyMessage", zap.Any("walletAddress", data.Address), zap.Error(err))
		return nil, err
	}

	if *data.AddressBTC != "" {
		user2, _ := u.Repo.FindUserByBTCWalletAddress(*data.AddressBTC)
		if user2 != nil {
			if user2.WalletAddressBTCTaproot == *data.AddressBTC {
				if data.Address != user2.WalletAddress {
					err := errors.New("invalid wallet address")
					logger.AtLog.Error("VerifyMessage", zap.Any("walletAddress", data.Address), zap.Error(err))
					return nil, err
				}
			}
		}
	}

	token, refreshToken, err := u.Auth2.GenerateAllTokens(user.WalletAddress, "", "", "", userID)
	if err != nil {
		logger.AtLog.Error("VerifyMessage", zap.Any("walletAddress", data.Address), zap.Error(err))
		return nil, err
	}

	logger.AtLog.Info("token", token)
	tokenMd5 := helpers.GenerateMd5String(token)
	logger.AtLog.Info("tokenMd5", tokenMd5)
	
	err = u.Cache.SetDataWithExpireTime(tokenMd5, userID, int(utils.TOKEN_CACHE_EXPIRED_TIME))
	if err != nil {
		logger.AtLog.Error("VerifyMessage", zap.Any("walletAddress", data.Address), zap.Error(err))
		return nil, err
	}

	if data.AddressBTC != nil && *data.AddressBTC != "" {
		if user.WalletAddressBTCTaproot == "" {
			user.WalletAddressBTCTaproot = *data.AddressBTC
			logger.AtLog.Info("user.WalletAddressBTCTaproot.Updated", true)
		}
		if user.WalletAddressBTC == "" {
			user.WalletAddressBTC = *data.AddressBTC
			logger.AtLog.Info("user.WalletAddressBTC.Updated", true)
		}
	}

	if user.WalletAddressPayment == "" {
		if data.AddressPayment == "" {
			if user.WalletType != entity.WalletType_BTC_PRVKEY {
				user.WalletAddressPayment = user.WalletAddress
				logger.AtLog.Info("user.WalletAddressPayment.Updated", true)
			}
		} else {
			user.WalletAddressPayment = data.AddressPayment
			logger.AtLog.Info("user.WalletAddressPayment.Updated", true)
		}
	}

	updated, err := u.Repo.ReplaceOne(bson.D{
		{utils.KEY_WALLET_ADDRESS, user.WalletAddress},
	}, user)

	if err != nil {
		logger.AtLog.Error("VerifyMessage", zap.Any("walletAddress", data.Address), zap.Error(err))
		return nil, err
	}

	_ = updated
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

	logger.AtLog.Info("verify",  zap.Bool("isVerified", isVerified), zap.String("signerHex", signerHex), zap.String("signatureHex", signatureHex), zap.String("signer", signer), zap.String("msgStr", msgStr),  zap.Any("recoveredAddr", recoveredAddr))
	return isVerified, nil
}
