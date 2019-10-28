package api

import (
	"github.com/morezig/go-odoo/types"
)

type IrQwebFieldMonetaryService struct {
	client *Client
}

func NewIrQwebFieldMonetaryService(c *Client) *IrQwebFieldMonetaryService {
	return &IrQwebFieldMonetaryService{client: c}
}

func (svc *IrQwebFieldMonetaryService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrQwebFieldMonetaryModel, name)
}

func (svc *IrQwebFieldMonetaryService) GetByIds(ids []int64) (*types.IrQwebFieldMonetarys, error) {
	i := &types.IrQwebFieldMonetarys{}
	return i, svc.client.getByIds(types.IrQwebFieldMonetaryModel, ids, i)
}

func (svc *IrQwebFieldMonetaryService) GetByName(name string) (*types.IrQwebFieldMonetarys, error) {
	i := &types.IrQwebFieldMonetarys{}
	return i, svc.client.getByName(types.IrQwebFieldMonetaryModel, name, i)
}

func (svc *IrQwebFieldMonetaryService) GetByField(field string, value string) (*types.IrQwebFieldMonetarys, error) {
	i := &types.IrQwebFieldMonetarys{}
	return i, svc.client.getByField(types.IrQwebFieldMonetaryModel, field, value, i)
}

func (svc *IrQwebFieldMonetaryService) GetAll() (*types.IrQwebFieldMonetarys, error) {
	i := &types.IrQwebFieldMonetarys{}
	return i, svc.client.getAll(types.IrQwebFieldMonetaryModel, i)
}

func (svc *IrQwebFieldMonetaryService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrQwebFieldMonetaryModel, fields, relations)
}

func (svc *IrQwebFieldMonetaryService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrQwebFieldMonetaryModel, ids, fields, relations)
}

func (svc *IrQwebFieldMonetaryService) Delete(ids []int64) error {
	return svc.client.delete(types.IrQwebFieldMonetaryModel, ids)
}
