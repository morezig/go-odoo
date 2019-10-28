package api

import (
	"github.com/morezig/go-odoo/types"
)

type ResPartnerTitleService struct {
	client *Client
}

func NewResPartnerTitleService(c *Client) *ResPartnerTitleService {
	return &ResPartnerTitleService{client: c}
}

func (svc *ResPartnerTitleService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ResPartnerTitleModel, name)
}

func (svc *ResPartnerTitleService) GetByIds(ids []int64) (*types.ResPartnerTitles, error) {
	r := &types.ResPartnerTitles{}
	return r, svc.client.getByIds(types.ResPartnerTitleModel, ids, r)
}

func (svc *ResPartnerTitleService) GetByName(name string) (*types.ResPartnerTitles, error) {
	r := &types.ResPartnerTitles{}
	return r, svc.client.getByName(types.ResPartnerTitleModel, name, r)
}

func (svc *ResPartnerTitleService) GetByField(field string, value string) (*types.ResPartnerTitles, error) {
	r := &types.ResPartnerTitles{}
	return r, svc.client.getByField(types.ResPartnerTitleModel, field, value, r)
}

func (svc *ResPartnerTitleService) GetAll() (*types.ResPartnerTitles, error) {
	r := &types.ResPartnerTitles{}
	return r, svc.client.getAll(types.ResPartnerTitleModel, r)
}

func (svc *ResPartnerTitleService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ResPartnerTitleModel, fields, relations)
}

func (svc *ResPartnerTitleService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResPartnerTitleModel, ids, fields, relations)
}

func (svc *ResPartnerTitleService) Delete(ids []int64) error {
	return svc.client.delete(types.ResPartnerTitleModel, ids)
}
