package api

import (
	"github.com/morezig/go-odoo/types"
)

type IrActionsActUrlService struct {
	client *Client
}

func NewIrActionsActUrlService(c *Client) *IrActionsActUrlService {
	return &IrActionsActUrlService{client: c}
}

func (svc *IrActionsActUrlService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrActionsActUrlModel, name)
}

func (svc *IrActionsActUrlService) GetByIds(ids []int64) (*types.IrActionsActUrls, error) {
	i := &types.IrActionsActUrls{}
	return i, svc.client.getByIds(types.IrActionsActUrlModel, ids, i)
}

func (svc *IrActionsActUrlService) GetByName(name string) (*types.IrActionsActUrls, error) {
	i := &types.IrActionsActUrls{}
	return i, svc.client.getByName(types.IrActionsActUrlModel, name, i)
}

func (svc *IrActionsActUrlService) GetByField(field string, value string) (*types.IrActionsActUrls, error) {
	i := &types.IrActionsActUrls{}
	return i, svc.client.getByField(types.IrActionsActUrlModel, field, value, i)
}

func (svc *IrActionsActUrlService) GetAll() (*types.IrActionsActUrls, error) {
	i := &types.IrActionsActUrls{}
	return i, svc.client.getAll(types.IrActionsActUrlModel, i)
}

func (svc *IrActionsActUrlService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrActionsActUrlModel, fields, relations)
}

func (svc *IrActionsActUrlService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrActionsActUrlModel, ids, fields, relations)
}

func (svc *IrActionsActUrlService) Delete(ids []int64) error {
	return svc.client.delete(types.IrActionsActUrlModel, ids)
}
