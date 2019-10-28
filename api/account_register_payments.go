package api

import (
	"github.com/morezig/go-odoo/types"
)

type AccountRegisterPaymentsService struct {
	client *Client
}

func NewAccountRegisterPaymentsService(c *Client) *AccountRegisterPaymentsService {
	return &AccountRegisterPaymentsService{client: c}
}

func (svc *AccountRegisterPaymentsService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountRegisterPaymentsModel, name)
}

func (svc *AccountRegisterPaymentsService) GetByIds(ids []int64) (*types.AccountRegisterPaymentss, error) {
	a := &types.AccountRegisterPaymentss{}
	return a, svc.client.getByIds(types.AccountRegisterPaymentsModel, ids, a)
}

func (svc *AccountRegisterPaymentsService) GetByName(name string) (*types.AccountRegisterPaymentss, error) {
	a := &types.AccountRegisterPaymentss{}
	return a, svc.client.getByName(types.AccountRegisterPaymentsModel, name, a)
}

func (svc *AccountRegisterPaymentsService) GetByField(field string, value string) (*types.AccountRegisterPaymentss, error) {
	a := &types.AccountRegisterPaymentss{}
	return a, svc.client.getByField(types.AccountRegisterPaymentsModel, field, value, a)
}

func (svc *AccountRegisterPaymentsService) GetAll() (*types.AccountRegisterPaymentss, error) {
	a := &types.AccountRegisterPaymentss{}
	return a, svc.client.getAll(types.AccountRegisterPaymentsModel, a)
}

func (svc *AccountRegisterPaymentsService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountRegisterPaymentsModel, fields, relations)
}

func (svc *AccountRegisterPaymentsService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountRegisterPaymentsModel, ids, fields, relations)
}

func (svc *AccountRegisterPaymentsService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountRegisterPaymentsModel, ids)
}
