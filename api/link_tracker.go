package api

import (
	"github.com/morezig/go-odoo/types"
)

type LinkTrackerService struct {
	client *Client
}

func NewLinkTrackerService(c *Client) *LinkTrackerService {
	return &LinkTrackerService{client: c}
}

func (svc *LinkTrackerService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.LinkTrackerModel, name)
}

func (svc *LinkTrackerService) GetByIds(ids []int64) (*types.LinkTrackers, error) {
	l := &types.LinkTrackers{}
	return l, svc.client.getByIds(types.LinkTrackerModel, ids, l)
}

func (svc *LinkTrackerService) GetByName(name string) (*types.LinkTrackers, error) {
	l := &types.LinkTrackers{}
	return l, svc.client.getByName(types.LinkTrackerModel, name, l)
}

func (svc *LinkTrackerService) GetByField(field string, value string) (*types.LinkTrackers, error) {
	l := &types.LinkTrackers{}
	return l, svc.client.getByField(types.LinkTrackerModel, field, value, l)
}

func (svc *LinkTrackerService) GetAll() (*types.LinkTrackers, error) {
	l := &types.LinkTrackers{}
	return l, svc.client.getAll(types.LinkTrackerModel, l)
}

func (svc *LinkTrackerService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.LinkTrackerModel, fields, relations)
}

func (svc *LinkTrackerService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.LinkTrackerModel, ids, fields, relations)
}

func (svc *LinkTrackerService) Delete(ids []int64) error {
	return svc.client.delete(types.LinkTrackerModel, ids)
}
