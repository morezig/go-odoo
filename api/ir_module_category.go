package api

import (
	"github.com/morezig/go-odoo/types"
)

type IrModuleCategoryService struct {
	client *Client
}

func NewIrModuleCategoryService(c *Client) *IrModuleCategoryService {
	return &IrModuleCategoryService{client: c}
}

func (svc *IrModuleCategoryService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrModuleCategoryModel, name)
}

func (svc *IrModuleCategoryService) GetByIds(ids []int64) (*types.IrModuleCategorys, error) {
	i := &types.IrModuleCategorys{}
	return i, svc.client.getByIds(types.IrModuleCategoryModel, ids, i)
}

func (svc *IrModuleCategoryService) GetByName(name string) (*types.IrModuleCategorys, error) {
	i := &types.IrModuleCategorys{}
	return i, svc.client.getByName(types.IrModuleCategoryModel, name, i)
}

func (svc *IrModuleCategoryService) GetByField(field string, value string) (*types.IrModuleCategorys, error) {
	i := &types.IrModuleCategorys{}
	return i, svc.client.getByField(types.IrModuleCategoryModel, field, value, i)
}

func (svc *IrModuleCategoryService) GetAll() (*types.IrModuleCategorys, error) {
	i := &types.IrModuleCategorys{}
	return i, svc.client.getAll(types.IrModuleCategoryModel, i)
}

func (svc *IrModuleCategoryService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrModuleCategoryModel, fields, relations)
}

func (svc *IrModuleCategoryService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrModuleCategoryModel, ids, fields, relations)
}

func (svc *IrModuleCategoryService) Delete(ids []int64) error {
	return svc.client.delete(types.IrModuleCategoryModel, ids)
}
