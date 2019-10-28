package api

import (
	"github.com/morezig/go-odoo/types"
)

type ResConfigService struct {
	client *Client
}

func NewResConfigService(c *Client) *ResConfigService {
	return &ResConfigService{client: c}
}

func (svc *ResConfigService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ResConfigModel, name)
}

func (svc *ResConfigService) GetByIds(ids []int64) (*types.ResConfigs, error) {
	r := &types.ResConfigs{}
	return r, svc.client.getByIds(types.ResConfigModel, ids, r)
}

func (svc *ResConfigService) GetByName(name string) (*types.ResConfigs, error) {
	r := &types.ResConfigs{}
	return r, svc.client.getByName(types.ResConfigModel, name, r)
}

func (svc *ResConfigService) GetByField(field string, value string) (*types.ResConfigs, error) {
	r := &types.ResConfigs{}
	return r, svc.client.getByField(types.ResConfigModel, field, value, r)
}

func (svc *ResConfigService) GetAll() (*types.ResConfigs, error) {
	r := &types.ResConfigs{}
	return r, svc.client.getAll(types.ResConfigModel, r)
}

func (svc *ResConfigService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ResConfigModel, fields, relations)
}

func (svc *ResConfigService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResConfigModel, ids, fields, relations)
}

func (svc *ResConfigService) Delete(ids []int64) error {
	return svc.client.delete(types.ResConfigModel, ids)
}
