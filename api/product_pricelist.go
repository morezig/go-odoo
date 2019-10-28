package api

import (
	"github.com/morezig/go-odoo/types"
)

type ProductPricelistService struct {
	client *Client
}

func NewProductPricelistService(c *Client) *ProductPricelistService {
	return &ProductPricelistService{client: c}
}

func (svc *ProductPricelistService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ProductPricelistModel, name)
}

func (svc *ProductPricelistService) GetByIds(ids []int64) (*types.ProductPricelists, error) {
	p := &types.ProductPricelists{}
	return p, svc.client.getByIds(types.ProductPricelistModel, ids, p)
}

func (svc *ProductPricelistService) GetByName(name string) (*types.ProductPricelists, error) {
	p := &types.ProductPricelists{}
	return p, svc.client.getByName(types.ProductPricelistModel, name, p)
}

func (svc *ProductPricelistService) GetByField(field string, value string) (*types.ProductPricelists, error) {
	p := &types.ProductPricelists{}
	return p, svc.client.getByField(types.ProductPricelistModel, field, value, p)
}

func (svc *ProductPricelistService) GetAll() (*types.ProductPricelists, error) {
	p := &types.ProductPricelists{}
	return p, svc.client.getAll(types.ProductPricelistModel, p)
}

func (svc *ProductPricelistService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ProductPricelistModel, fields, relations)
}

func (svc *ProductPricelistService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProductPricelistModel, ids, fields, relations)
}

func (svc *ProductPricelistService) Delete(ids []int64) error {
	return svc.client.delete(types.ProductPricelistModel, ids)
}
