package api

import (
	"github.com/morezig/go-odoo/types"
)

type AccountJournalService struct {
	client *Client
}

func NewAccountJournalService(c *Client) *AccountJournalService {
	return &AccountJournalService{client: c}
}

func (svc *AccountJournalService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountJournalModel, name)
}

func (svc *AccountJournalService) GetByIds(ids []int64) (*types.AccountJournals, error) {
	a := &types.AccountJournals{}
	return a, svc.client.getByIds(types.AccountJournalModel, ids, a)
}

func (svc *AccountJournalService) GetByName(name string) (*types.AccountJournals, error) {
	a := &types.AccountJournals{}
	return a, svc.client.getByName(types.AccountJournalModel, name, a)
}

func (svc *AccountJournalService) GetByField(field string, value string) (*types.AccountJournals, error) {
	a := &types.AccountJournals{}
	return a, svc.client.getByField(types.AccountJournalModel, field, value, a)
}

func (svc *AccountJournalService) GetAll() (*types.AccountJournals, error) {
	a := &types.AccountJournals{}
	return a, svc.client.getAll(types.AccountJournalModel, a)
}

func (svc *AccountJournalService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountJournalModel, fields, relations)
}

func (svc *AccountJournalService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountJournalModel, ids, fields, relations)
}

func (svc *AccountJournalService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountJournalModel, ids)
}
