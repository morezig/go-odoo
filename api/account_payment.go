package api

import (
	"github.com/morezig/go-odoo/types"
)

type AccountPaymentService struct {
	client *Client
}

func NewAccountPaymentService(c *Client) *AccountPaymentService {
	return &AccountPaymentService{client: c}
}

func (svc *AccountPaymentService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountPaymentModel, name)
}

func (svc *AccountPaymentService) GetByIds(ids []int64) (*types.AccountPayments, error) {
	a := &types.AccountPayments{}
	return a, svc.client.getByIds(types.AccountPaymentModel, ids, a)
}

func (svc *AccountPaymentService) GetByName(name string) (*types.AccountPayments, error) {
	a := &types.AccountPayments{}
	return a, svc.client.getByName(types.AccountPaymentModel, name, a)
}

func (svc *AccountPaymentService) GetByField(field string, value string) (*types.AccountPayments, error) {
	a := &types.AccountPayments{}
	return a, svc.client.getByField(types.AccountPaymentModel, field, value, a)
}

func (svc *AccountPaymentService) GetAll() (*types.AccountPayments, error) {
	a := &types.AccountPayments{}
	return a, svc.client.getAll(types.AccountPaymentModel, a)
}

func (svc *AccountPaymentService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountPaymentModel, fields, relations)
}

func (svc *AccountPaymentService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountPaymentModel, ids, fields, relations)
}

func (svc *AccountPaymentService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountPaymentModel, ids)
}
