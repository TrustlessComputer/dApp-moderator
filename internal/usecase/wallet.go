package usecase

import (
	"context"
	"dapp-moderator/external/quicknode"
	"dapp-moderator/internal/usecase/structure"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func (u Usecase) getUTXOFromBlockStream(address string) ([]quicknode.WalletAddressBalanceResp, error) {
	url := u.Config.BlockStream + "/api/address/" + address + "/utxo"

	fmt.Printf("URL blockstream: %v\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Printf("resp getUTXOFromBlockStream: %+v\n", resp)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyStr := string(body)
	if strings.Contains(bodyStr, "RPC error") {
		return nil, errors.New(bodyStr)
	}

	respUTXOs := []structure.UTXOFromBlockStream{}
	err = json.Unmarshal(body, &respUTXOs)
	if err != nil {
		return nil, err
	}

	result := []quicknode.WalletAddressBalanceResp{}
	for _, utxo := range respUTXOs {
		if utxo.Status.Confirmed {
			result = append(result, quicknode.WalletAddressBalanceResp{
				Height:  int64(utxo.Status.BlockHeight),
				Address: address,
				Hash:    utxo.Txid,
				Index:   utxo.Vout,
				Value:   uint64(utxo.Value),
			})
		}
	}

	fmt.Printf("result: %+v\n", result)
	return result, nil
}

func (u *Usecase) GetBTCWalletInfo(ctx context.Context, address string) (*structure.WalletInfo, error) {
	var result structure.WalletInfo

	t := time.Now()

	abs := []quicknode.WalletAddressBalanceResp{}
	var err error

	switch u.Config.ENV {
	case "develop":
		{
			abs, err = u.getUTXOFromBlockStream(address)
			if err != nil {
				return nil, err
			}
		}
	case "production":
		{
			abs, err = u.QuickNode.AddressBalance(address)
			if err != nil {
				return nil, err
			}
		}
	default:
		{
			return nil, fmt.Errorf("Invalid network env: %v", u.Config.ENV)
		}
	}

	trackT1 := time.Since(t)

	inscriptionByOutput := make(map[string][]structure.WalletInscriptionByOutput)

	txRefs := []*structure.TxRef{}
	txRefsChan := make(chan *structure.TxRef, len(abs))

	balance := 0
	for _, ab := range abs {
		go func(ab quicknode.WalletAddressBalanceResp, txRefsChan chan *structure.TxRef) {
			tmp := &structure.TxRef{}

			defer func() {
				txRefsChan <- tmp
			}()

			out := fmt.Sprintf("%s:%d", ab.Hash, ab.Index)

			data := &structure.InscriptionByOutput{}
			if u.Config.ENV == "production" {
				data, err = u.GetInscriptionByOutput(out)
				if err != nil {
					return
				}
			}

			tmp.TxHash = ab.Hash
			tmp.BlockHeight = int(ab.Height)
			tmp.TxInputN = 0
			tmp.TxOutputN = ab.Index
			tmp.Value = int(ab.Value)

			if data != nil && len(data.Inscriptions) > 0 {
				tmp.IsOrdinal = true
			}

			balance += tmp.Value
		}(ab, txRefsChan)
	}

	for _, _ = range abs {
		dataFromChan := <-txRefsChan
		txRefs = append(txRefs, dataFromChan)
	}

	trackT2 := time.Since(t)
	trackT3 := time.Since(t)

	result.Address = address
	result.TotalReceived = 0
	result.TotalSent = 0
	result.Balance = balance
	result.UnconfirmedBalance = 0
	result.FinalBalance = balance
	result.UnconfirmedNTx = 0
	result.FinalNTx = 0
	result.Txrefs = txRefs
	result.TxURL = ""
	result.Inscriptions = []structure.WalletInscriptionInfo{}
	result.InscriptionsByOutputs = inscriptionByOutput
	result.Loadtime = make(map[string]string)
	result.Loadtime["trackT1"] = trackT1.String()
	result.Loadtime["trackT2"] = trackT2.String()
	result.Loadtime["trackT3"] = trackT3.String()

	return &result, nil
}

func (u *Usecase) GetBTCWalletTXS(ctx context.Context, address string) (interface{}, error) {
	return u.BlockStream.Txs(address)
}
