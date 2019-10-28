package api

import (
	"github.com/morezig/go-odoo/types"
)

type ResourceCalendarLeavesService struct {
	client *Client
}

func NewResourceCalendarLeavesService(c *Client) *ResourceCalendarLeavesService {
	return &ResourceCalendarLeavesService{client: c}
}

func (svc *ResourceCalendarLeavesService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ResourceCalendarLeavesModel, name)
}

func (svc *ResourceCalendarLeavesService) GetByIds(ids []int64) (*types.ResourceCalendarLeavess, error) {
	r := &types.ResourceCalendarLeavess{}
	return r, svc.client.getByIds(types.ResourceCalendarLeavesModel, ids, r)
}

func (svc *ResourceCalendarLeavesService) GetByName(name string) (*types.ResourceCalendarLeavess, error) {
	r := &types.ResourceCalendarLeavess{}
	return r, svc.client.getByName(types.ResourceCalendarLeavesModel, name, r)
}

func (svc *ResourceCalendarLeavesService) GetByField(field string, value string) (*types.ResourceCalendarLeavess, error) {
	r := &types.ResourceCalendarLeavess{}
	return r, svc.client.getByField(types.ResourceCalendarLeavesModel, field, value, r)
}

func (svc *ResourceCalendarLeavesService) GetAll() (*types.ResourceCalendarLeavess, error) {
	r := &types.ResourceCalendarLeavess{}
	return r, svc.client.getAll(types.ResourceCalendarLeavesModel, r)
}

func (svc *ResourceCalendarLeavesService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ResourceCalendarLeavesModel, fields, relations)
}

func (svc *ResourceCalendarLeavesService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResourceCalendarLeavesModel, ids, fields, relations)
}

func (svc *ResourceCalendarLeavesService) Delete(ids []int64) error {
	return svc.client.delete(types.ResourceCalendarLeavesModel, ids)
}
