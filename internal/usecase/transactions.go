package usecase

import (
	"context"
	"dapp-moderator/internal/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *Usecase) ScanTransactions(ctx context.Context) error {

	const layout = "2006-01-02T15:04:05.000000Z"

	isBreak := false
	numDup := 0
	nextParams := &NextParams{}
	nextParams.BlockNumber = "0"
	nextParams.Index = "0"
	nextParams.PagesLimit = "0"
	nextParams.PageNumber = "0"
	nextParams.PageSize = "50"

	for !isBreak {
		rs, err := u.GetTransactionInfo(nextParams)
		if err != nil {
			fmt.Println(nextParams.PageNumber)
			return err
		}
		nextParams = rs.NextParams
		for _, htmlStr := range rs.Items {
			htmlStr = strings.ReplaceAll(htmlStr, `\n`, "")
			htmlStr = strings.ReplaceAll(htmlStr, `\"`, `"`)

			doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlStr))
			if err != nil {
				log.Fatal(err)
			}
			hash := doc.Find("a[data-test='transaction_hash_link']").Text()

			tmp, _ := u.Repo.FindTransactionByHash(ctx, hash)
			if tmp != nil {
				fmt.Println(nextParams.PageNumber)
				fmt.Println(hash)
				numDup += 1
				if numDup >= 10 {
					return nil
				}
				continue
			}

			transactionType := doc.Find("span.tile-label").Text()
			transactionStatus := doc.Find("span.tile-status-label").Text()
			method := doc.Find("div.bs-label.method.ml-1").Text()

			var from, to string

			doc.Find(`a[data-test="address_hash_link"]`).Each(
				func(i int, s *goquery.Selection) {
					if i == 0 {
						from = s.Find("span.d-none.d-md-none.d-xl-inline").Text()
					}
					if i == 1 {
						classContract := doc.Find(`span[class="contract-address"]`)
						if classContract != nil {
							to = classContract.AttrOr("data-address-hash", "")
						} else {
							to = s.Find("span.d-none.d-md-none.d-xl-inline").Text()
						}
					}
				},
			)

			toName := doc.Find("span.d-none.d-md-none.d-lg-inline.d-xl-inline").Text()

			amountQ := doc.Find(`span[class="d-flex flex-md-row flex-column mt-3 mt-md-0"]`)

			amount := strings.TrimSpace(strings.ReplaceAll(amountQ.Find("span.tile-title").Text(), "TC", ""))
			amount = strings.ReplaceAll(amount, ",", "")

			fee := strings.TrimSpace(strings.ReplaceAll(amountQ.Find("span.ml-0.ml-md-1.text-nowrap").Text(), "TX Fee", ""))
			fee = strings.ReplaceAll(fee, ",", "")

			var tokenTransferFrom1, tokenTransferTo1, tokenTransferToken1, tokenTransferAmount1, tokenTransferID1, tokenTransferName1 string
			var tokenTransferFrom2, tokenTransferTo2, tokenTransferToken2, tokenTransferAmount2, tokenTransferID2, tokenTransferName2 string
			doc.Find(`div[data-test="token_transfer"]`).Each(func(i int, s *goquery.Selection) {
				switch i {
				case 0:
					{
						s.Find(`span[class="d-inline-block tile-type-token-transfer-short-name"]`).Each(
							func(i2 int, s2 *goquery.Selection) {
								if i2 == 0 {
									tokenTransferFrom1 = strings.TrimSpace(strings.ReplaceAll(s2.Find(`a[data-test="address_hash_link"]`).AttrOr("href", ""), "/address/", ""))
								}
								if i2 == 1 {
									tokenTransferTo1 = strings.TrimSpace(strings.ReplaceAll(s2.Find(`a[data-test="address_hash_link"]`).AttrOr("href", ""), "/address/", ""))
								}
							},
						)
						tokenTransferToken := strings.TrimSpace(strings.ReplaceAll(s.Find(`span[class="col-xs-12 col-lg-4 ml-3 ml-sm-0"]`).Find(`a[data-test="token_link"]`).AttrOr("href", ""), "/token/", ""))
						tokenTransferToken = strings.Split(tokenTransferToken, "/")[0]
						tokenTransferToken1 = tokenTransferToken
						tokenTransferAmountOrId := strings.TrimSpace(s.Find(`span[class="col-xs-12 col-lg-4 ml-3 ml-sm-0"]`).Text())
						if strings.Contains(tokenTransferAmountOrId, "TokenID") {
							tokenTransferAmountOrId = strings.TrimSpace(strings.Split(tokenTransferAmountOrId, " ")[1])
							tokenTransferAmountOrId = strings.ReplaceAll(tokenTransferAmountOrId, "[", "")
							tokenTransferAmountOrId = strings.ReplaceAll(tokenTransferAmountOrId, "]", "")
							tokenTransferID1 = strings.ReplaceAll(tokenTransferAmountOrId, ",", "")
						} else {
							tokenTransferAmount1 = strings.TrimSpace(strings.Split(tokenTransferAmountOrId, " ")[0])
							tokenTransferAmount1 = strings.ReplaceAll(tokenTransferAmount1, ",", "")
							tokenTransferName1 = strings.TrimSpace(strings.Split(tokenTransferAmountOrId, " ")[1])
						}
					}
				case 1:
					{
						s.Find(`span[class="d-inline-block tile-type-token-transfer-short-name"]`).Each(
							func(i2 int, s2 *goquery.Selection) {
								if i2 == 0 {
									tokenTransferFrom2 = strings.TrimSpace(strings.ReplaceAll(s2.Find(`a[data-test="address_hash_link"]`).AttrOr("href", ""), "/address/", ""))
								}
								if i2 == 1 {
									tokenTransferTo2 = strings.TrimSpace(strings.ReplaceAll(s2.Find(`a[data-test="address_hash_link"]`).AttrOr("href", ""), "/address/", ""))
								}
							},
						)
						tokenTransferToken := strings.TrimSpace(strings.ReplaceAll(s.Find(`span[class="col-xs-12 col-lg-4 ml-3 ml-sm-0"]`).Find(`a[data-test="token_link"]`).AttrOr("href", ""), "/token/", ""))
						tokenTransferToken = strings.Split(tokenTransferToken, "/")[0]
						tokenTransferToken2 = tokenTransferToken
						tokenTransferAmountOrId := strings.TrimSpace(s.Find(`span[class="col-xs-12 col-lg-4 ml-3 ml-sm-0"]`).Text())
						if strings.Contains(tokenTransferAmountOrId, "TokenID") {
							tokenTransferAmountOrId = strings.TrimSpace(strings.Split(tokenTransferAmountOrId, " ")[1])
							tokenTransferAmountOrId = strings.ReplaceAll(tokenTransferAmountOrId, "[", "")
							tokenTransferAmountOrId = strings.ReplaceAll(tokenTransferAmountOrId, "]", "")
							tokenTransferID2 = strings.ReplaceAll(tokenTransferAmountOrId, ",", "")
						} else {
							tokenTransferAmount2 = strings.TrimSpace(strings.Split(tokenTransferAmountOrId, " ")[0])
							tokenTransferAmount2 = strings.ReplaceAll(tokenTransferAmount2, ",", "")
							tokenTransferName2 = strings.TrimSpace(strings.Split(tokenTransferAmountOrId, " ")[1])
						}
					}
				}
			})

			//block
			divBlock := doc.Find("div.col-md-3.col-lg-2.d-flex.flex-row.flex-md-column.flex-nowrap.justify-content-center.text-md-right.mt-3.mt-md-0.tile-bottom span a")
			blockNumber := strings.TrimSpace(strings.ReplaceAll(divBlock.Text(), "Block #", ""))
			// fmt.Println("blockNumber")
			// fmt.Println(blockNumber)

			blockTime := strings.TrimSpace(doc.Find("span[data-from-now]").AttrOr("data-from-now", ""))
			// fmt.Println("blockTime")
			// fmt.Println(blockTime)

			now := time.Now()
			transaction := &entity.Transactions{}
			transaction.CreatedAt = &now
			transaction.UpdatedAt = &now
			transaction.DeletedAt = nil
			transaction.TransactionType = strings.TrimSpace(transactionType)
			transaction.TransactionStatus = strings.TrimSpace(transactionStatus)
			transaction.Method = strings.TrimSpace(method)
			transaction.Hash = strings.TrimSpace(hash)
			transaction.FromAddress = strings.TrimSpace(from)
			transaction.ToAddress = strings.TrimSpace(to)
			transaction.ToName = strings.TrimSpace(toName)
			transaction.Amount, _ = primitive.ParseDecimal128(amount)
			transaction.Fee, _ = primitive.ParseDecimal128(fee)
			transaction.TransferFrom1 = strings.TrimSpace(tokenTransferFrom1)
			transaction.TransferTo1 = strings.TrimSpace(tokenTransferTo1)
			transaction.TransferToken1 = strings.TrimSpace(tokenTransferToken1)
			transaction.TransferAmount1, _ = primitive.ParseDecimal128(tokenTransferAmount1)
			transaction.TransferTokenID1 = strings.TrimSpace(tokenTransferID1)
			transaction.TransferTokenName1 = strings.TrimSpace(tokenTransferName1)
			transaction.TransferFrom2 = strings.TrimSpace(tokenTransferFrom2)
			transaction.TransferTo2 = strings.TrimSpace(tokenTransferTo2)
			transaction.TransferToken2 = strings.TrimSpace(tokenTransferToken2)
			transaction.TransferTokenName2 = strings.TrimSpace(tokenTransferName2)
			transaction.TransferAmount2, _ = primitive.ParseDecimal128(tokenTransferAmount2)
			transaction.TransferTokenID2 = strings.TrimSpace(tokenTransferID2)
			tmpBlock, err := strconv.ParseUint(blockNumber, 10, 64)

			if err != nil {
				return err
			}
			transaction.Block = uint(tmpBlock)
			transaction.BlockTime, err = time.Parse(layout, blockTime)
			if err != nil {
				return err
			}
			_, err = u.Repo.InsertOne(transaction)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Usecase) doWithAuth(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(req)
}

func (c *Usecase) getJSON(url string, headers map[string]string, result interface{}) (int, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := c.doWithAuth(req)
	if err != nil {
		return 0, fmt.Errorf("failed request: %v", err)
	}
	if resp.StatusCode >= 300 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return resp.StatusCode, fmt.Errorf("http response bad status %d %s", resp.StatusCode, err.Error())
		}
		return resp.StatusCode, fmt.Errorf("http response bad status %d %s", resp.StatusCode, string(bodyBytes))
	}
	if result != nil {
		return resp.StatusCode, json.NewDecoder(resp.Body).Decode(result)
	}
	return resp.StatusCode, nil
}

type NextParams struct {
	BlockNumber interface{} `json:"block_number"`
	Index       interface{} `json:"index"`
	PageNumber  interface{} `json:"page_number"`
	PageSize    interface{} `json:"page_size"`
	PagesLimit  interface{} `json:"pages_limit"`
}
type TransactionResponse struct {
	Items      []string    `json:"items"`
	NextParams *NextParams `json:"next_page_params"`
}

func (c *Usecase) GetTransactionInfo(params *NextParams) (*TransactionResponse, error) {
	rs := &TransactionResponse{}
	url := ""
	pageNumber := 0
	switch v := params.PageNumber.(type) {
	case float64:
		pageNumber = int(v) + 1
		blockNumber, _ := params.BlockNumber.(float64)
		index, _ := params.Index.(float64)
		pageSize, _ := params.PageSize.(float64)
		pagesLimit, _ := params.PagesLimit.(float64)

		url = fmt.Sprintf("https://explorer.trustless.computer/txs?type=JSON&page_number=%d&block_number=%d&index=%d&page_size=%d&pages_limit=%d", pageNumber, int(blockNumber),
			int(index), int(pageSize), int(pagesLimit))
		if pageNumber >= int(pagesLimit) && int(pagesLimit) > 0 {
			return nil, fmt.Errorf("pageNumber over limit")
		}
	case string:
		pageNumber, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		tmp, _ := params.PagesLimit.(string)
		pagesLimit, err := strconv.Atoi(tmp)
		if err != nil {
			return nil, err
		}
		url = fmt.Sprintf("https://explorer.trustless.computer/txs?type=JSON&page_number=%d&block_number=%s&index=%s&page_size=%s&pages_limit=%d", pageNumber+1, params.BlockNumber,
			params.Index, params.PageSize, int(pagesLimit))
		if pageNumber >= pagesLimit && int(pagesLimit) > 0 {
			return nil, fmt.Errorf("pageNumber over limit")
		}
	default:
		fmt.Printf("i has an unsupported type: %T\n", v)
	}

	fmt.Println(url)
	_, err := c.getJSON(
		url,
		map[string]string{},
		&rs,
	)
	if err != nil {
		return nil, err
	}
	return rs, nil
}
