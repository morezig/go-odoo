package api

import (
	"github.com/morezig/go-odoo/types"
)

type HrEmployeeCategoryService struct {
	client *Client
}

func NewHrEmployeeCategoryService(c *Client) *HrEmployeeCategoryService {
	return &HrEmployeeCategoryService{client: c}
}

func (svc *HrEmployeeCategoryService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.HrEmployeeCategoryModel, name)
}

func (svc *HrEmployeeCategoryService) GetByIds(ids []int64) (*types.HrEmployeeCategorys, error) {
	h := &types.HrEmployeeCategorys{}
	return h, svc.client.getByIds(types.HrEmployeeCategoryModel, ids, h)
}

func (svc *HrEmployeeCategoryService) GetByName(name string) (*types.HrEmployeeCategorys, error) {
	h := &types.HrEmployeeCategorys{}
	return h, svc.client.getByName(types.HrEmployeeCategoryModel, name, h)
}

func (svc *HrEmployeeCategoryService) GetByField(field string, value string) (*types.HrEmployeeCategorys, error) {
	h := &types.HrEmployeeCategorys{}
	return h, svc.client.getByField(types.HrEmployeeCategoryModel, field, value, h)
}

func (svc *HrEmployeeCategoryService) GetAll() (*types.HrEmployeeCategorys, error) {
	h := &types.HrEmployeeCategorys{}
	return h, svc.client.getAll(types.HrEmployeeCategoryModel, h)
}

func (svc *HrEmployeeCategoryService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.HrEmployeeCategoryModel, fields, relations)
}

func (svc *HrEmployeeCategoryService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.HrEmployeeCategoryModel, ids, fields, relations)
}

func (svc *HrEmployeeCategoryService) Delete(ids []int64) error {
	return svc.client.delete(types.HrEmployeeCategoryModel, ids)
}
