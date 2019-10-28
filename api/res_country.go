package api

import (
	"github.com/morezig/go-odoo/types"
)

type ResCountryService struct {
	client *Client
}

func NewResCountryService(c *Client) *ResCountryService {
	return &ResCountryService{client: c}
}

func (svc *ResCountryService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ResCountryModel, name)
}

func (svc *ResCountryService) GetByIds(ids []int64) (*types.ResCountrys, error) {
	r := &types.ResCountrys{}
	return r, svc.client.getByIds(types.ResCountryModel, ids, r)
}

func (svc *ResCountryService) GetByName(name string) (*types.ResCountrys, error) {
	r := &types.ResCountrys{}
	return r, svc.client.getByName(types.ResCountryModel, name, r)
}

func (svc *ResCountryService) GetByField(field string, value string) (*types.ResCountrys, error) {
	r := &types.ResCountrys{}
	return r, svc.client.getByField(types.ResCountryModel, field, value, r)
}

func (svc *ResCountryService) GetAll() (*types.ResCountrys, error) {
	r := &types.ResCountrys{}
	return r, svc.client.getAll(types.ResCountryModel, r)
}

func (svc *ResCountryService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ResCountryModel, fields, relations)
}

func (svc *ResCountryService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResCountryModel, ids, fields, relations)
}

func (svc *ResCountryService) Delete(ids []int64) error {
	return svc.client.delete(types.ResCountryModel, ids)
}
