package api

import (
	"github.com/morezig/go-odoo/types"
)

type AccountFiscalPositionService struct {
	client *Client
}

func NewAccountFiscalPositionService(c *Client) *AccountFiscalPositionService {
	return &AccountFiscalPositionService{client: c}
}

func (svc *AccountFiscalPositionService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountFiscalPositionModel, name)
}

func (svc *AccountFiscalPositionService) GetByIds(ids []int64) (*types.AccountFiscalPositions, error) {
	a := &types.AccountFiscalPositions{}
	return a, svc.client.getByIds(types.AccountFiscalPositionModel, ids, a)
}

func (svc *AccountFiscalPositionService) GetByName(name string) (*types.AccountFiscalPositions, error) {
	a := &types.AccountFiscalPositions{}
	return a, svc.client.getByName(types.AccountFiscalPositionModel, name, a)
}

func (svc *AccountFiscalPositionService) GetByField(field string, value string) (*types.AccountFiscalPositions, error) {
	a := &types.AccountFiscalPositions{}
	return a, svc.client.getByField(types.AccountFiscalPositionModel, field, value, a)
}

func (svc *AccountFiscalPositionService) GetAll() (*types.AccountFiscalPositions, error) {
	a := &types.AccountFiscalPositions{}
	return a, svc.client.getAll(types.AccountFiscalPositionModel, a)
}

func (svc *AccountFiscalPositionService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountFiscalPositionModel, fields, relations)
}

func (svc *AccountFiscalPositionService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountFiscalPositionModel, ids, fields, relations)
}

func (svc *AccountFiscalPositionService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountFiscalPositionModel, ids)
}
