package api

import (
	"github.com/morezig/go-odoo/types"
)

type ResLangService struct {
	client *Client
}

func NewResLangService(c *Client) *ResLangService {
	return &ResLangService{client: c}
}

func (svc *ResLangService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ResLangModel, name)
}

func (svc *ResLangService) GetByIds(ids []int64) (*types.ResLangs, error) {
	r := &types.ResLangs{}
	return r, svc.client.getByIds(types.ResLangModel, ids, r)
}

func (svc *ResLangService) GetByName(name string) (*types.ResLangs, error) {
	r := &types.ResLangs{}
	return r, svc.client.getByName(types.ResLangModel, name, r)
}

func (svc *ResLangService) GetByField(field string, value string) (*types.ResLangs, error) {
	r := &types.ResLangs{}
	return r, svc.client.getByField(types.ResLangModel, field, value, r)
}

func (svc *ResLangService) GetAll() (*types.ResLangs, error) {
	r := &types.ResLangs{}
	return r, svc.client.getAll(types.ResLangModel, r)
}

func (svc *ResLangService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ResLangModel, fields, relations)
}

func (svc *ResLangService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResLangModel, ids, fields, relations)
}

func (svc *ResLangService) Delete(ids []int64) error {
	return svc.client.delete(types.ResLangModel, ids)
}
