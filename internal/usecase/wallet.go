package usecase

import (
	"context"
	"dapp-moderator/internal/usecase/structure"
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
	
	txRefs := []structure.TxRef{}
	
	balance := 0
	for _, ab := range abs {
		tmp := structure.TxRef{
			TxHash: ab.Hash,
			BlockHeight: int(ab.Height),
			TxInputN: 0,
			TxOutputN: 0,
			Value: int(ab.Value),
			//Confirmed: ,
		}

		balance += tmp.Value
		txRefs = append(txRefs, tmp)
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