package bfs_service

import (
	"dapp-moderator/utils/config"
	"dapp-moderator/utils/helpers"
	"dapp-moderator/utils/redis"
	"fmt"
)

type BfsService struct {
	conf *config.Config
	serverURL string
	cache     redis.IRedisCache
}

func NewBfsService(conf *config.Config, cache redis.IRedisCache) *BfsService {
	return &BfsService{
		conf:      conf,
		serverURL: conf.BFSService,
		cache:     cache,
	}
}

func (q BfsService) Files(walletAddress string) ([]string, error) {
	headers := make(map[string]string)	
	data, _, _, err := helpers.JsonRequest(fmt.Sprintf("%s/files/%s",q.serverURL, walletAddress,), "GET", headers, nil)
	if err != nil {
		return nil, err
	}

	
	resp, err := q.ParseData(data)
	if err != nil {
		return nil, err
	}
	
	
	return resp.ToFiles(), nil
}

func (q BfsService) BrowseFiles(walletAddress string, path string) (*BrowsedFileResp, error) {
	headers := make(map[string]string)	
	data, _, _, err := helpers.JsonRequest(fmt.Sprintf("%s/browse/%s?path=%s",q.serverURL, walletAddress,path), "GET", headers, nil)
	if err != nil {
		return nil, err
	}


	resp, err := q.ParseData(data)
	if err != nil {
		return nil, err
	}
	
	
	return resp.ToBrowedFiles(), nil
}

func (q BfsService) FileInfo(walletAddress string, path string) (*FileInfoResp, error) {
	headers := make(map[string]string)	
	data, _, _, err := helpers.JsonRequest(fmt.Sprintf("%s/info?path=%s/%s",q.serverURL, walletAddress,path), "GET", headers, nil)
	if err != nil {
		return nil, err
	}


	resp, err := q.ParseData(data)
	if err != nil {
		return nil, err
	}
	
	
	return resp.ToFileInfo(), nil
}

func (q BfsService) ParseData(data []byte) (*ServiceResp, error) {
	resp := &ServiceResp{}
	err := helpers.ParseData(data, resp)
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, resp.Error
	}

	return resp, nil
}