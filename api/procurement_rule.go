package api

import (
	"github.com/morezig/go-odoo/types"
)

type ProcurementRuleService struct {
	client *Client
}

func NewProcurementRuleService(c *Client) *ProcurementRuleService {
	return &ProcurementRuleService{client: c}
}

func (svc *ProcurementRuleService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ProcurementRuleModel, name)
}

func (svc *ProcurementRuleService) GetByIds(ids []int64) (*types.ProcurementRules, error) {
	p := &types.ProcurementRules{}
	return p, svc.client.getByIds(types.ProcurementRuleModel, ids, p)
}

func (svc *ProcurementRuleService) GetByName(name string) (*types.ProcurementRules, error) {
	p := &types.ProcurementRules{}
	return p, svc.client.getByName(types.ProcurementRuleModel, name, p)
}

func (svc *ProcurementRuleService) GetByField(field string, value string) (*types.ProcurementRules, error) {
	p := &types.ProcurementRules{}
	return p, svc.client.getByField(types.ProcurementRuleModel, field, value, p)
}

func (svc *ProcurementRuleService) GetAll() (*types.ProcurementRules, error) {
	p := &types.ProcurementRules{}
	return p, svc.client.getAll(types.ProcurementRuleModel, p)
}

func (svc *ProcurementRuleService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ProcurementRuleModel, fields, relations)
}

func (svc *ProcurementRuleService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProcurementRuleModel, ids, fields, relations)
}

func (svc *ProcurementRuleService) Delete(ids []int64) error {
	return svc.client.delete(types.ProcurementRuleModel, ids)
}
