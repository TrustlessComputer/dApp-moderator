package usecase

import (
	"context"
	"dapp-moderator/external/quicknode"
	"dapp-moderator/internal/usecase/structure"
	"fmt"
	"time"
)

func (u Usecase) GetBTCWalletInfo(ctx context.Context, address string) (*structure.WalletInfo, error) {
	var result structure.WalletInfo

	t := time.Now()
	abs, err := u.QuickNode.AddressBalance(address)
	if err != nil {
		return nil, err
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
			data, err := u.GetInscriptionByOutput(out)
			if err != nil {
				return
			}

			tmp.TxHash = ab.Hash
			tmp.BlockHeight = int(ab.Height)
			tmp.TxInputN = 0
			tmp.TxOutputN = ab.Index
			tmp.Value = int(ab.Value)

			if len(data.Inscriptions) > 0 {
				tmp.IsOrdinal = true
			}

			balance += tmp.Value
		}(ab, txRefsChan)
	}

	for _,_ = range abs {
		dataFromChan := <- txRefsChan
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
