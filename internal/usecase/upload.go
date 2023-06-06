package usecase

import (
	"context"
	"dapp-moderator/internal/entity"
	"dapp-moderator/internal/usecase/structure"
	"dapp-moderator/utils/googlecloud"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/logger"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type File interface {
	io.ReadSeeker
}

type Chunk struct {
	Bufsize int
	Offset  int64
	Index   int
	Data    []byte
	Err     error
}

type CalculatedChunk struct {
	FileName   string
	FileType   string
	Chunks     int
	BufferSize int //kb
	FileSize   int //kb
}

func (u *Usecase) UploadFile(fileHeader *multipart.FileHeader) (string, error) {
	gf := googlecloud.GcsFile{
		FileHeader: fileHeader,
	}
	uploaded, err := u.Storage.FileUploadToBucket(gf)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", os.Getenv("GCS_DOMAIN"), uploaded.Name), nil
}

func (u *Usecase) UploadFileMultipart(fileHeader *multipart.FileHeader) (*entity.UploadedFile, error) {
	path := "artifact"
	gf := googlecloud.GcsFile{
		FileHeader: fileHeader,
		Path:       &path,
	}
	uploaded, err := u.Storage.FileUploadToBucket(gf)
	if err != nil {
		return nil, err
	}

	uploadedFIle := &entity.UploadedFile{
		Name:     uploaded.Name,
		Size:     int(uploaded.Size),
		Path:     uploaded.Path,
		FileType: uploaded.Minetype,
		FullPath: fmt.Sprintf("%s/%s", os.Getenv("GCS_DOMAIN"), uploaded.Name),
	}

	err = u.Repo.InsertUploadedFile(uploadedFIle)
	if err != nil {
		return nil, err
	}

	return uploadedFIle, nil
}

func (u *Usecase) MakeChunkWorker(file *os.File, chunksizes []Chunk, i int, dataChan chan Chunk) {
	chunk := chunksizes[i]
	buffer := make([]byte, chunk.Bufsize)
	bytesread, err := file.ReadAt(buffer, chunk.Offset)

	if err != nil && err != io.EOF {
		chunk.Err = err
		dataChan <- chunk
		return
	}

	chunk.Data = buffer[:bytesread]
	dataChan <- chunk
}

func (u *Usecase) CalculateChunks(fileName string) (*CalculatedChunk, error) {
	res := CalculatedChunk{}

	// Read the file in 4-byte chunks
	bufferSizeStr := os.Getenv("FILE_CHUNK_SIZE")
	if bufferSizeStr == "" {
		bufferSizeStr = "350"
	}
	bufferSize, err := strconv.Atoi(bufferSizeStr)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("MakeChunks  strconv.Atoi - %s", fileName), zap.Int("bufferSize", bufferSize), zap.String("fileName", fileName), zap.Error(err))
		return nil, err
	}

	//TODO - temporary remove for testing
	//bufferSize = 1000 * bufferSize
	f, err := u.Storage.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	filesize := len(f)
	concurrency := filesize / bufferSize
	if remainder := filesize % bufferSize; remainder != 0 {
		concurrency++
	}

	res.Chunks = concurrency
	res.BufferSize = bufferSize
	res.FileSize = filesize
	res.FileName = fileName

	ext1 := filepath.Ext(fileName)
	res.FileType = ext1

	return &res, nil
}

// Refs: https://askgolang.com/how-to-read-a-file-in-chunks-in-golang/
func (u *Usecase) MakeChunks(fileName string, uploadedIndex string) ([]Chunk, *CalculatedChunk, error) {

	calculated, err := u.CalculateChunks(fileName)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("MakeChunks -  CalculateChunks.Err - %s", fileName), zap.String("fileName", fileName), zap.Error(err))
		return nil, nil, err
	}

	concurrency := calculated.Chunks
	bufferSize := calculated.BufferSize
	filesize := calculated.FileSize

	bytes, err := u.Storage.ReadFile(calculated.FileName)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("MakeChunks -  u.Storage.ReadFile - %s", fileName), zap.String("fileName", fileName), zap.Error(err))
		return nil, nil, err
	}

	tmp := strings.ReplaceAll(calculated.FileName, `/`, "-")
	tmp = fmt.Sprintf("%s-%s", uploadedIndex, tmp)

	fc, err := os.Create(tmp)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("MakeChunks -  os.Create - %s", fileName), zap.String("fileName", fileName), zap.Error(err))
		return nil, nil, err
	}
	defer func() {
		os.Remove(tmp)
	}()

	defer fc.Close()
	_, err = fc.Write(bytes)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("MakeChunks -  fc.Write - %s", fileName), zap.String("fileName", fileName), zap.Error(err))
		return nil, nil, err
	}

	//file, err := os.Open(calculated.FileName)
	//if err != nil {
	//	logger.AtLog.Logger.Error(fmt.Sprintf("MakeChunks -  Open.Err - %s", fileName), zap.String("fileName", fileName), zap.Error(err))
	//	return nil, nil, err
	//}

	chunks := []Chunk{}
	chunksizes := make([]Chunk, concurrency)

	// calculate each chunk size
	for i := 0; i < concurrency; i++ {
		chunksizes[i].Offset = int64(i * bufferSize)
		chunksizes[i].Bufsize = bufferSize
		chunksizes[i].Index = i
		if i == concurrency-1 {
			chunksizes[i].Bufsize = filesize % bufferSize
		}
	}

	dataChan := make(chan Chunk)
	for i := 0; i < concurrency; i++ {
		go u.MakeChunkWorker(fc, chunksizes, i, dataChan)
	}

	for i := 0; i < concurrency; i++ {
		dataFromChan := <-dataChan
		if dataFromChan.Err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("MakeChunks -  dataFromChan.Err - %s", fileName), zap.String("fileName", fileName), zap.Error(dataFromChan.Err))
			return nil, nil, dataFromChan.Err
		}
		chunks = append(chunks, dataFromChan)
	}

	return chunks, calculated, nil
}

func (u *Usecase) MergeChunks(fileName string, chunks []Chunk) (*string, error) {
	f, err := os.Create(fileName)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("MergeChunks - %s", fileName), zap.String("fileName", fileName), zap.Error(err))
		os.Remove(fileName)
		return nil, err
	}
	defer f.Close()

	//sor by index
	sort.SliceStable(chunks, func(i, j int) bool {
		if chunks[i].Index < chunks[j].Index {
			return true
		}
		return false
	})

	for key, chunk := range chunks {
		_ = key
		_, err = f.Write(chunk.Data)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("MergeChunks - %s", fileName), zap.String("fileName", fileName), zap.Error(err))
			os.Remove(fileName)
			return nil, err
		}

	}

	return nil, nil
}

func (u *Usecase) FilterChunks(filter *entity.FilterChunks) ([]entity.UploadedFileChunk, error) {
	return u.Repo.ListChunks(filter)
}

func (u *Usecase) GetUploadedFiles(f *entity.FilterUploadedFile) ([]entity.QueriedUploadedFile, error) {
	return u.Repo.ListUploadedFiles(f)
}

func (u *Usecase) UpdateTxHashForUploadedFile(data *structure.UpdateUploadedFileTxHash) (*entity.UploadedFile, error) {
	uploadedFile, err := u.Repo.GetUploadedFile(data.FileID)
	if err != nil {
		logger.AtLog.Logger.Error("UpdateTxHashForUploadedFile GetUploadedFile", zap.String("fileID", data.FileID), zap.String("txHash", data.TxHash), zap.Error(err))
		return nil, err
	}

	if uploadedFile.TxHash != "" && uploadedFile.Chunks != 0 {
		return nil, errors.New(fmt.Sprintf("chunks of file %s have been created", data.FileID))
	}

	chunks, calculated, err := u.MakeChunks(uploadedFile.Name, uploadedFile.ID.Hex())
	if err != nil {
		logger.AtLog.Logger.Error("UpdateTxHashForUploadedFile MakeChunks", zap.String("fileID", data.FileID), zap.String("txHash", data.TxHash), zap.Error(err))

		return nil, err
	}

	fileChunks := []entity.IEntity{}
	for _, chunk := range chunks {
		fileChunk := &entity.UploadedFileChunk{
			FileID:     uploadedFile.ID,
			ChunkIndex: chunk.Index,
			ChunkData:  chunk.Data,
			Status:     entity.ChunkNew,
		}
		fileChunks = append(fileChunks, fileChunk)
	}

	err = u.Repo.UpdateChunksTxHashForUploadedFile(data.FileID, calculated.FileSize, data.TxHash, data.WalletAddress, calculated.Chunks, calculated.BufferSize)
	if err != nil {
		logger.AtLog.Logger.Error("UpdateTxHashForUploadedFile update info", zap.String("fileID", data.FileID), zap.String("txHash", data.TxHash), zap.Error(err))

		return nil, err
	}

	_, err = u.Repo.InsertMany(fileChunks)
	if err != nil {
		logger.AtLog.Logger.Error("UpdateTxHashForUploadedFile InsertMany chunks", zap.String("fileID", data.FileID), zap.String("txHash", data.TxHash), zap.Error(err))
		return nil, err
	}

	uploadedFile.Size = calculated.FileSize
	uploadedFile.TxHash = data.TxHash
	uploadedFile.WalletAddress = data.WalletAddress
	uploadedFile.Chunks = calculated.Chunks
	uploadedFile.ChunkSize = calculated.BufferSize

	return uploadedFile, nil
}

func (u *Usecase) UpdateTxHashForAChunk(fileID string, chunkID string, txHash string) (*entity.UploadedFileChunk, error) {
	//verify chunk
	verifiedDataFromChain := &structure.TxHashInfo{}
	verifiedData, _, statusCode, err := helpers.TxHashInfo(txHash)
	if err == nil && statusCode == 200 {
		err = json.Unmarshal(verifiedData, verifiedDataFromChain)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("UpdateTxHashForAChunk - verify hash - %s", fileID), zap.String("fileID", fileID), zap.String("chunkID", chunkID), zap.String("txHash", txHash), zap.Error(err))
			return nil, err
		}
	}

	c, err := u.Repo.FindChunk(fileID, chunkID)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("UpdateTxHashForAChunk - FindChunk - %s", fileID), zap.String("fileID", fileID), zap.String("chunkID", chunkID), zap.String("txHash", txHash), zap.Error(err))
		return nil, err
	}

	uploadedFile, err := u.Repo.FindUploadedFileByID(c.FileID.Hex())
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("UpdateTxHashForAChunk - FindUploadedFileByID - %s", fileID), zap.String("fileID", fileID), zap.String("chunkID", chunkID), zap.String("txHash", txHash), zap.Error(err))
		return nil, err
	}

	if strings.ToLower(uploadedFile.WalletAddress) != strings.ToLower(verifiedDataFromChain.Result.From) {
		err := errors.New(fmt.Sprintf("Chunk is not valid"))
		logger.AtLog.Logger.Error(fmt.Sprintf("UpdateTxHashForAChunk - FindUploadedFileByID - %s", fileID), zap.String("fileID", fileID), zap.String("chunkID", chunkID), zap.String("txHash", txHash), zap.Error(err))
		return nil, err
	}

	if strings.ToLower(uploadedFile.TxHash) == strings.ToLower(txHash) {
		err := errors.New(fmt.Sprintf("Cannot use txHash of uploaded file for a chunk"))
		logger.AtLog.Logger.Error(fmt.Sprintf("UpdateTxHashForAChunk - FindUploadedFileByID - %s", fileID), zap.String("fileID", fileID), zap.String("chunkID", chunkID), zap.String("txHash", txHash), zap.Error(err))
		return nil, err
	}

	err = u.Repo.UpdateChunkTxHash(fileID, chunkID, txHash)
	if err != nil {
		logger.AtLog.Logger.Error(fmt.Sprintf("UpdateTxHashForAChunk - verify hash - %s", fileID), zap.String("fileID", fileID), zap.String("chunkID", chunkID), zap.String("txHash", txHash), zap.Error(err))
		return nil, err
	}
	return c, nil
}

func (u *Usecase) GetChunkByID(fileID string, chunkID string) (*entity.UploadedFileChunk, error) {

	c, err := u.Repo.FindChunk(fileID, chunkID)
	if err != nil {
		return nil, err
	}

	return c, nil
}

//s3 multipart upload

func (u *Usecase) CreateMultipartUpload(ctx context.Context, group string, fileName string) (*string, error) {

	group = helpers.GenerateSlug(group)
	group = fmt.Sprintf("%s-%d", group, time.Now().UTC().Nanosecond())

	fileName = helpers.GenerateSlug(fileName)
	uploaded, err := u.S3Adapter.CreateMultiplePartsUpload(ctx, "artifact/"+group, fileName)

	return uploaded.UploadId, err
}

func (u *Usecase) UploadPart(ctx context.Context, uploadID string, file File, fileSize int64, partNumber int) error {

	if err := u.S3Adapter.UploadPart(uploadID, file, fileSize, partNumber); err != nil {
		return err
	}
	return nil
}

func (u *Usecase) CompleteMultipartUpload(ctx context.Context, uploadID string, walletAddress string) (*entity.UploadedFile, error) {
	uploaded, err := u.S3Adapter.CompleteMultipartUpload(ctx, uploadID)
	if err != nil {
		return nil, err
	}

	name := *uploaded.Key
	//bytes, err := u.Storage.ReadFile(*uploaded.Key)
	//if err != nil {
	//	return nil, err
	//}

	nameArray := strings.Split(name, ".")
	fType := ""
	if len(nameArray) > 1 {
		fType = nameArray[len(nameArray)-1]
	}

	//TODO - insert uploaded file here
	uploadedFIle := &entity.UploadedFile{
		Name: name,
		//Size:     len(bytes),
		Path:          name,
		FileType:      fType,
		WalletAddress: walletAddress,
		FullPath:      fmt.Sprintf("%s/%s", os.Getenv("GCS_DOMAIN"), name),
	}

	err = u.Repo.InsertUploadedFile(uploadedFIle)
	if err != nil {
		return nil, err
	}

	return uploadedFIle, nil
}

type chunkDataChan struct {
	Data      entity.UploadedFileChunk
	Err       error
	IsPending bool
	Hash      *types.Transaction
}

func (u *Usecase) ListenedChunks() error {
	chunks, err := u.Repo.GetUploadingChunks()
	if err != nil {
		logger.AtLog.Logger.Error("ListenedChunks", zap.Error(err))
		return err
	}

	inputChan := make(chan entity.UploadedFileChunk, len(chunks))
	resultChan := make(chan chunkDataChan, len(chunks))

	for i := 0; i < len(chunks); i++ {
		go u.ListenedChunkWorker(inputChan, resultChan)
	}

	for i, chunk := range chunks {
		inputChan <- chunk
		if i > 0 && i%100 == 0 {
			time.Sleep(time.Millisecond * 500)
		}
	}

	for i := 0; i < len(chunks); i++ {
		dataFromChan := <-resultChan
		if dataFromChan.Err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("ListenedChunks"), zap.Error(dataFromChan.Err))

			return dataFromChan.Err
		}

		chunk := dataFromChan.Data
		isPending := dataFromChan.IsPending

		logger.AtLog.Logger.Info(fmt.Sprintf("ListenedChunks - %s", chunk.TxHash), zap.String("hash", chunk.TxHash), zap.Bool("isPending", isPending))

		if !isPending {
			err := u.Repo.UpdateChunkTxHashStatus(chunk.ID.Hex(), chunk.TxHash, entity.ChunkUploaded) // uploaded to blockchain
			if err != nil {
				logger.AtLog.Logger.Error(fmt.Sprintf("ListenedChunks - %s", chunk.ID.Hex()), zap.Error(err), zap.String("txHash", chunk.TxHash))
				return err
			}
		}

	}

	return nil
}

func (u *Usecase) ListenedChunkWorker(input chan entity.UploadedFileChunk, output chan chunkDataChan) {
	inData := <-input

	hash, isPending, err := u.TCPublicNode.TransactionByHash(common.HexToHash(inData.TxHash))
	output <- chunkDataChan{
		Data:      inData,
		Err:       err,
		IsPending: isPending,
		Hash:      hash,
	}
}

type ArtifactChan struct {
	Nfts []entity.Nfts
	Err  error
	Done bool
}

type NftSizeChan struct {
	Nft entity.Nfts
	Err error
}

func (u *Usecase) CrontabUpdateImageSize(ctx context.Context) error {

	nftsChan := make(chan ArtifactChan)
	stopSig := make(chan bool)
	nfts := []entity.Nfts{}

	go func(stopSig chan bool) {
		contract := "0x16EfDc6D3F977E39DAc0Eb0E123FefFeD4320Bc0"
		page := 1
		limit := 50

		for {

			go u.GetNftArtifacts(ctx, contract, page, limit, nftsChan)
			page++

			stop := <-stopSig
			if stop {
				return
			}

		}

	}(stopSig)

	for {
		dataFromChan := <-nftsChan
		stopSig <- dataFromChan.Done

		nftsFromChan := dataFromChan.Nfts
		nfts = append(nfts, nftsFromChan...)

		if dataFromChan.Done {
			break
		}
	}

	workerInputChan := make(chan entity.Nfts, len(nfts))
	outputChan := make(chan NftSizeChan, len(nfts))

	for i := 0; i < len(nfts); i++ {
		go u.UpdateImageSize(workerInputChan, outputChan)
	}

	for i := 0; i < len(nfts); i++ {
		workerInputChan <- nfts[i]
		if i > 0 && i%5 == 0 {
			time.Sleep(500 * time.Millisecond)
		}
	}

	//create backup here
	if os.Getenv("ENV") != "production" {
		err := helpers.CreateFile("backup-nfts.json", nfts)
		if err != nil {
			logger.AtLog.Logger.Error("CrontabUpdateImageSize", zap.Error(err), zap.Int("nfts", len(nfts)))
			return err
		}
	}

	data := []entity.Nfts{}
	for i := 0; i < len(nfts); i++ {
		dataFromChan := <-outputChan
		if dataFromChan.Err != nil {
			logger.AtLog.Logger.Error("CrontabUpdateImageSize", zap.Error(dataFromChan.Err))
			continue
		}

		nft := dataFromChan.Nft
		//update DB
		updated, err := u.Repo.UpdateNftSize(nft.ContractAddress, nft.TokenID, nft.Size)
		if err != nil {
			logger.AtLog.Logger.Error(fmt.Sprintf("CrontabUpdateImageSize - %s - %s", nft.ContractAddress, nft.TokenID), zap.Error(err), zap.Int64("size", nft.Size))
			return err
		}

		logger.AtLog.Logger.Info(fmt.Sprintf("CrontabUpdateImageSize - %s - %s", nft.ContractAddress, nft.TokenID), zap.Any("updated", updated), zap.Int64("size", nft.Size))

		//only used for testing
		data = append(data, nft)
	}

	err := helpers.CreateFile("nfts.json", data)
	return err
}

func (u *Usecase) UpdateImageSize(input chan entity.Nfts, output chan NftSizeChan) {
	var err error
	inputData := <-input
	image := inputData.Image

	tokenURI := fmt.Sprintf("https://dapp.trustless.computer/dapp/api/nft-explorer/collections/%s/nfts/%s", inputData.ContractAddress, inputData.TokenID)
	inputData.TokenURI = tokenURI

	image = fmt.Sprintf("https://dapp.trustless.computer/dapp/api/nft-explorer/collections/%s/nfts/%s/content", inputData.ContractAddress, inputData.TokenID)
	inputData.Image = image

	res, headers, status, err := helpers.HttpRequest(image, "GET", make(map[string]string), nil)
	if err != nil {
		logger.AtLog.Logger.Info(fmt.Sprintf("UpdateImageSize - %s", inputData.TokenID), zap.Any("headers", headers), zap.Any("status", status))
	} else {
		logger.AtLog.Logger.Error(fmt.Sprintf("UpdateImageSize - %s", inputData.TokenID), zap.Any("headers", headers), zap.Any("status", status), zap.Error(err))

	}

	inputData.Size = int64(len(res))
	defer func() {
		output <- NftSizeChan{
			Nft: inputData,
			Err: err,
		}
	}()

}

func (u *Usecase) GetNftArtifacts(ctx context.Context, contract string, page, limit int, nftsChan chan ArtifactChan) {
	offset := (page - 1) * limit
	var err error
	nfts := []entity.Nfts{}

	defer func() {
		done := false
		if len(nfts) == 0 || err != nil {
			done = true
		}

		nftsChan <- ArtifactChan{
			Nfts: nfts,
			Err:  err,
			Done: done,
		}
	}()

	nfts, err = u.Repo.GetNfts(contract, offset, limit)
	if err != nil {
		logger.AtLog.Logger.Error("GetNftArtifacts", zap.Error(err))
	}
}
