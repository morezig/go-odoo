package api

import (
	"github.com/morezig/go-odoo/types"
)

type ImLivechatReportChannelService struct {
	client *Client
}

func NewImLivechatReportChannelService(c *Client) *ImLivechatReportChannelService {
	return &ImLivechatReportChannelService{client: c}
}

func (svc *ImLivechatReportChannelService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ImLivechatReportChannelModel, name)
}

func (svc *ImLivechatReportChannelService) GetByIds(ids []int64) (*types.ImLivechatReportChannels, error) {
	i := &types.ImLivechatReportChannels{}
	return i, svc.client.getByIds(types.ImLivechatReportChannelModel, ids, i)
}

func (svc *ImLivechatReportChannelService) GetByName(name string) (*types.ImLivechatReportChannels, error) {
	i := &types.ImLivechatReportChannels{}
	return i, svc.client.getByName(types.ImLivechatReportChannelModel, name, i)
}

func (svc *ImLivechatReportChannelService) GetByField(field string, value string) (*types.ImLivechatReportChannels, error) {
	i := &types.ImLivechatReportChannels{}
	return i, svc.client.getByField(types.ImLivechatReportChannelModel, field, value, i)
}

func (svc *ImLivechatReportChannelService) GetAll() (*types.ImLivechatReportChannels, error) {
	i := &types.ImLivechatReportChannels{}
	return i, svc.client.getAll(types.ImLivechatReportChannelModel, i)
}

func (svc *ImLivechatReportChannelService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ImLivechatReportChannelModel, fields, relations)
}

func (svc *ImLivechatReportChannelService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ImLivechatReportChannelModel, ids, fields, relations)
}

func (svc *ImLivechatReportChannelService) Delete(ids []int64) error {
	return svc.client.delete(types.ImLivechatReportChannelModel, ids)
}
