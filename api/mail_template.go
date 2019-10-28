package api

import (
	"github.com/morezig/go-odoo/types"
)

type MailTemplateService struct {
	client *Client
}

func NewMailTemplateService(c *Client) *MailTemplateService {
	return &MailTemplateService{client: c}
}

func (svc *MailTemplateService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.MailTemplateModel, name)
}

func (svc *MailTemplateService) GetByIds(ids []int64) (*types.MailTemplates, error) {
	m := &types.MailTemplates{}
	return m, svc.client.getByIds(types.MailTemplateModel, ids, m)
}

func (svc *MailTemplateService) GetByName(name string) (*types.MailTemplates, error) {
	m := &types.MailTemplates{}
	return m, svc.client.getByName(types.MailTemplateModel, name, m)
}

func (svc *MailTemplateService) GetByField(field string, value string) (*types.MailTemplates, error) {
	m := &types.MailTemplates{}
	return m, svc.client.getByField(types.MailTemplateModel, field, value, m)
}

func (svc *MailTemplateService) GetAll() (*types.MailTemplates, error) {
	m := &types.MailTemplates{}
	return m, svc.client.getAll(types.MailTemplateModel, m)
}

func (svc *MailTemplateService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.MailTemplateModel, fields, relations)
}

func (svc *MailTemplateService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailTemplateModel, ids, fields, relations)
}

func (svc *MailTemplateService) Delete(ids []int64) error {
	return svc.client.delete(types.MailTemplateModel, ids)
}
