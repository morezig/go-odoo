package api

import (
	"github.com/morezig/go-odoo/types"
)

type ResourceMixinService struct {
	client *Client
}

func NewResourceMixinService(c *Client) *ResourceMixinService {
	return &ResourceMixinService{client: c}
}

func (svc *ResourceMixinService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ResourceMixinModel, name)
}

func (svc *ResourceMixinService) GetByIds(ids []int64) (*types.ResourceMixins, error) {
	r := &types.ResourceMixins{}
	return r, svc.client.getByIds(types.ResourceMixinModel, ids, r)
}

func (svc *ResourceMixinService) GetByName(name string) (*types.ResourceMixins, error) {
	r := &types.ResourceMixins{}
	return r, svc.client.getByName(types.ResourceMixinModel, name, r)
}

func (svc *ResourceMixinService) GetByField(field string, value string) (*types.ResourceMixins, error) {
	r := &types.ResourceMixins{}
	return r, svc.client.getByField(types.ResourceMixinModel, field, value, r)
}

func (svc *ResourceMixinService) GetAll() (*types.ResourceMixins, error) {
	r := &types.ResourceMixins{}
	return r, svc.client.getAll(types.ResourceMixinModel, r)
}

func (svc *ResourceMixinService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ResourceMixinModel, fields, relations)
}

func (svc *ResourceMixinService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResourceMixinModel, ids, fields, relations)
}

func (svc *ResourceMixinService) Delete(ids []int64) error {
	return svc.client.delete(types.ResourceMixinModel, ids)
}
