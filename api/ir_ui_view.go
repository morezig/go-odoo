package api

import (
	"github.com/morezig/go-odoo/types"
)

type IrUiViewService struct {
	client *Client
}

func NewIrUiViewService(c *Client) *IrUiViewService {
	return &IrUiViewService{client: c}
}

func (svc *IrUiViewService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrUiViewModel, name)
}

func (svc *IrUiViewService) GetByIds(ids []int64) (*types.IrUiViews, error) {
	i := &types.IrUiViews{}
	return i, svc.client.getByIds(types.IrUiViewModel, ids, i)
}

func (svc *IrUiViewService) GetByName(name string) (*types.IrUiViews, error) {
	i := &types.IrUiViews{}
	return i, svc.client.getByName(types.IrUiViewModel, name, i)
}

func (svc *IrUiViewService) GetByField(field string, value string) (*types.IrUiViews, error) {
	i := &types.IrUiViews{}
	return i, svc.client.getByField(types.IrUiViewModel, field, value, i)
}

func (svc *IrUiViewService) GetAll() (*types.IrUiViews, error) {
	i := &types.IrUiViews{}
	return i, svc.client.getAll(types.IrUiViewModel, i)
}

func (svc *IrUiViewService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrUiViewModel, fields, relations)
}

func (svc *IrUiViewService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrUiViewModel, ids, fields, relations)
}

func (svc *IrUiViewService) Delete(ids []int64) error {
	return svc.client.delete(types.IrUiViewModel, ids)
}
