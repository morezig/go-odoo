package api

import (
	"github.com/morezig/go-odoo/types"
)

type IrMailServerService struct {
	client *Client
}

func NewIrMailServerService(c *Client) *IrMailServerService {
	return &IrMailServerService{client: c}
}

func (svc *IrMailServerService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrMailServerModel, name)
}

func (svc *IrMailServerService) GetByIds(ids []int64) (*types.IrMailServers, error) {
	i := &types.IrMailServers{}
	return i, svc.client.getByIds(types.IrMailServerModel, ids, i)
}

func (svc *IrMailServerService) GetByName(name string) (*types.IrMailServers, error) {
	i := &types.IrMailServers{}
	return i, svc.client.getByName(types.IrMailServerModel, name, i)
}

func (svc *IrMailServerService) GetByField(field string, value string) (*types.IrMailServers, error) {
	i := &types.IrMailServers{}
	return i, svc.client.getByField(types.IrMailServerModel, field, value, i)
}

func (svc *IrMailServerService) GetAll() (*types.IrMailServers, error) {
	i := &types.IrMailServers{}
	return i, svc.client.getAll(types.IrMailServerModel, i)
}

func (svc *IrMailServerService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrMailServerModel, fields, relations)
}

func (svc *IrMailServerService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrMailServerModel, ids, fields, relations)
}

func (svc *IrMailServerService) Delete(ids []int64) error {
	return svc.client.delete(types.IrMailServerModel, ids)
}
