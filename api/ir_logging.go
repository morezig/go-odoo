package api

import (
	"github.com/morezig/go-odoo/types"
)

type IrLoggingService struct {
	client *Client
}

func NewIrLoggingService(c *Client) *IrLoggingService {
	return &IrLoggingService{client: c}
}

func (svc *IrLoggingService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrLoggingModel, name)
}

func (svc *IrLoggingService) GetByIds(ids []int64) (*types.IrLoggings, error) {
	i := &types.IrLoggings{}
	return i, svc.client.getByIds(types.IrLoggingModel, ids, i)
}

func (svc *IrLoggingService) GetByName(name string) (*types.IrLoggings, error) {
	i := &types.IrLoggings{}
	return i, svc.client.getByName(types.IrLoggingModel, name, i)
}

func (svc *IrLoggingService) GetByField(field string, value string) (*types.IrLoggings, error) {
	i := &types.IrLoggings{}
	return i, svc.client.getByField(types.IrLoggingModel, field, value, i)
}

func (svc *IrLoggingService) GetAll() (*types.IrLoggings, error) {
	i := &types.IrLoggings{}
	return i, svc.client.getAll(types.IrLoggingModel, i)
}

func (svc *IrLoggingService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrLoggingModel, fields, relations)
}

func (svc *IrLoggingService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrLoggingModel, ids, fields, relations)
}

func (svc *IrLoggingService) Delete(ids []int64) error {
	return svc.client.delete(types.IrLoggingModel, ids)
}
