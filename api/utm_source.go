package api

import (
	"github.com/morezig/go-odoo/types"
)

type UtmSourceService struct {
	client *Client
}

func NewUtmSourceService(c *Client) *UtmSourceService {
	return &UtmSourceService{client: c}
}

func (svc *UtmSourceService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.UtmSourceModel, name)
}

func (svc *UtmSourceService) GetByIds(ids []int64) (*types.UtmSources, error) {
	u := &types.UtmSources{}
	return u, svc.client.getByIds(types.UtmSourceModel, ids, u)
}

func (svc *UtmSourceService) GetByName(name string) (*types.UtmSources, error) {
	u := &types.UtmSources{}
	return u, svc.client.getByName(types.UtmSourceModel, name, u)
}

func (svc *UtmSourceService) GetByField(field string, value string) (*types.UtmSources, error) {
	u := &types.UtmSources{}
	return u, svc.client.getByField(types.UtmSourceModel, field, value, u)
}

func (svc *UtmSourceService) GetAll() (*types.UtmSources, error) {
	u := &types.UtmSources{}
	return u, svc.client.getAll(types.UtmSourceModel, u)
}

func (svc *UtmSourceService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.UtmSourceModel, fields, relations)
}

func (svc *UtmSourceService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.UtmSourceModel, ids, fields, relations)
}

func (svc *UtmSourceService) Delete(ids []int64) error {
	return svc.client.delete(types.UtmSourceModel, ids)
}
