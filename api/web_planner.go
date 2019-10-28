package api

import (
	"github.com/morezig/go-odoo/types"
)

type WebPlannerService struct {
	client *Client
}

func NewWebPlannerService(c *Client) *WebPlannerService {
	return &WebPlannerService{client: c}
}

func (svc *WebPlannerService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.WebPlannerModel, name)
}

func (svc *WebPlannerService) GetByIds(ids []int64) (*types.WebPlanners, error) {
	w := &types.WebPlanners{}
	return w, svc.client.getByIds(types.WebPlannerModel, ids, w)
}

func (svc *WebPlannerService) GetByName(name string) (*types.WebPlanners, error) {
	w := &types.WebPlanners{}
	return w, svc.client.getByName(types.WebPlannerModel, name, w)
}

func (svc *WebPlannerService) GetByField(field string, value string) (*types.WebPlanners, error) {
	w := &types.WebPlanners{}
	return w, svc.client.getByField(types.WebPlannerModel, field, value, w)
}

func (svc *WebPlannerService) GetAll() (*types.WebPlanners, error) {
	w := &types.WebPlanners{}
	return w, svc.client.getAll(types.WebPlannerModel, w)
}

func (svc *WebPlannerService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.WebPlannerModel, fields, relations)
}

func (svc *WebPlannerService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.WebPlannerModel, ids, fields, relations)
}

func (svc *WebPlannerService) Delete(ids []int64) error {
	return svc.client.delete(types.WebPlannerModel, ids)
}
