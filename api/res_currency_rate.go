package api

import (
	"github.com/morezig/go-odoo/types"
)

type ResCurrencyRateService struct {
	client *Client
}

func NewResCurrencyRateService(c *Client) *ResCurrencyRateService {
	return &ResCurrencyRateService{client: c}
}

func (svc *ResCurrencyRateService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ResCurrencyRateModel, name)
}

func (svc *ResCurrencyRateService) GetByIds(ids []int64) (*types.ResCurrencyRates, error) {
	r := &types.ResCurrencyRates{}
	return r, svc.client.getByIds(types.ResCurrencyRateModel, ids, r)
}

func (svc *ResCurrencyRateService) GetByName(name string) (*types.ResCurrencyRates, error) {
	r := &types.ResCurrencyRates{}
	return r, svc.client.getByName(types.ResCurrencyRateModel, name, r)
}

func (svc *ResCurrencyRateService) GetByField(field string, value string) (*types.ResCurrencyRates, error) {
	r := &types.ResCurrencyRates{}
	return r, svc.client.getByField(types.ResCurrencyRateModel, field, value, r)
}

func (svc *ResCurrencyRateService) GetAll() (*types.ResCurrencyRates, error) {
	r := &types.ResCurrencyRates{}
	return r, svc.client.getAll(types.ResCurrencyRateModel, r)
}

func (svc *ResCurrencyRateService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ResCurrencyRateModel, fields, relations)
}

func (svc *ResCurrencyRateService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResCurrencyRateModel, ids, fields, relations)
}

func (svc *ResCurrencyRateService) Delete(ids []int64) error {
	return svc.client.delete(types.ResCurrencyRateModel, ids)
}
