package api

import (
	"github.com/morezig/go-odoo/types"
)

type MailShortcodeService struct {
	client *Client
}

func NewMailShortcodeService(c *Client) *MailShortcodeService {
	return &MailShortcodeService{client: c}
}

func (svc *MailShortcodeService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.MailShortcodeModel, name)
}

func (svc *MailShortcodeService) GetByIds(ids []int64) (*types.MailShortcodes, error) {
	m := &types.MailShortcodes{}
	return m, svc.client.getByIds(types.MailShortcodeModel, ids, m)
}

func (svc *MailShortcodeService) GetByName(name string) (*types.MailShortcodes, error) {
	m := &types.MailShortcodes{}
	return m, svc.client.getByName(types.MailShortcodeModel, name, m)
}

func (svc *MailShortcodeService) GetByField(field string, value string) (*types.MailShortcodes, error) {
	m := &types.MailShortcodes{}
	return m, svc.client.getByField(types.MailShortcodeModel, field, value, m)
}

func (svc *MailShortcodeService) GetAll() (*types.MailShortcodes, error) {
	m := &types.MailShortcodes{}
	return m, svc.client.getAll(types.MailShortcodeModel, m)
}

func (svc *MailShortcodeService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.MailShortcodeModel, fields, relations)
}

func (svc *MailShortcodeService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailShortcodeModel, ids, fields, relations)
}

func (svc *MailShortcodeService) Delete(ids []int64) error {
	return svc.client.delete(types.MailShortcodeModel, ids)
}
