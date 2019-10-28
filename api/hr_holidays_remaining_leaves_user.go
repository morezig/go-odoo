package api

import (
	"github.com/morezig/go-odoo/types"
)

type HrHolidaysRemainingLeavesUserService struct {
	client *Client
}

func NewHrHolidaysRemainingLeavesUserService(c *Client) *HrHolidaysRemainingLeavesUserService {
	return &HrHolidaysRemainingLeavesUserService{client: c}
}

func (svc *HrHolidaysRemainingLeavesUserService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.HrHolidaysRemainingLeavesUserModel, name)
}

func (svc *HrHolidaysRemainingLeavesUserService) GetByIds(ids []int64) (*types.HrHolidaysRemainingLeavesUsers, error) {
	h := &types.HrHolidaysRemainingLeavesUsers{}
	return h, svc.client.getByIds(types.HrHolidaysRemainingLeavesUserModel, ids, h)
}

func (svc *HrHolidaysRemainingLeavesUserService) GetByName(name string) (*types.HrHolidaysRemainingLeavesUsers, error) {
	h := &types.HrHolidaysRemainingLeavesUsers{}
	return h, svc.client.getByName(types.HrHolidaysRemainingLeavesUserModel, name, h)
}

func (svc *HrHolidaysRemainingLeavesUserService) GetByField(field string, value string) (*types.HrHolidaysRemainingLeavesUsers, error) {
	h := &types.HrHolidaysRemainingLeavesUsers{}
	return h, svc.client.getByField(types.HrHolidaysRemainingLeavesUserModel, field, value, h)
}

func (svc *HrHolidaysRemainingLeavesUserService) GetAll() (*types.HrHolidaysRemainingLeavesUsers, error) {
	h := &types.HrHolidaysRemainingLeavesUsers{}
	return h, svc.client.getAll(types.HrHolidaysRemainingLeavesUserModel, h)
}

func (svc *HrHolidaysRemainingLeavesUserService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.HrHolidaysRemainingLeavesUserModel, fields, relations)
}

func (svc *HrHolidaysRemainingLeavesUserService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.HrHolidaysRemainingLeavesUserModel, ids, fields, relations)
}

func (svc *HrHolidaysRemainingLeavesUserService) Delete(ids []int64) error {
	return svc.client.delete(types.HrHolidaysRemainingLeavesUserModel, ids)
}
