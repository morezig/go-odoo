package api

import (
	"github.com/morezig/go-odoo/types"
)

type CrmLead2opportunityPartnerService struct {
	client *Client
}

func NewCrmLead2opportunityPartnerService(c *Client) *CrmLead2opportunityPartnerService {
	return &CrmLead2opportunityPartnerService{client: c}
}

func (svc *CrmLead2opportunityPartnerService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.CrmLead2opportunityPartnerModel, name)
}

func (svc *CrmLead2opportunityPartnerService) GetByIds(ids []int64) (*types.CrmLead2opportunityPartners, error) {
	c := &types.CrmLead2opportunityPartners{}
	return c, svc.client.getByIds(types.CrmLead2opportunityPartnerModel, ids, c)
}

func (svc *CrmLead2opportunityPartnerService) GetByName(name string) (*types.CrmLead2opportunityPartners, error) {
	c := &types.CrmLead2opportunityPartners{}
	return c, svc.client.getByName(types.CrmLead2opportunityPartnerModel, name, c)
}

func (svc *CrmLead2opportunityPartnerService) GetByField(field string, value string) (*types.CrmLead2opportunityPartners, error) {
	c := &types.CrmLead2opportunityPartners{}
	return c, svc.client.getByField(types.CrmLead2opportunityPartnerModel, field, value, c)
}

func (svc *CrmLead2opportunityPartnerService) GetAll() (*types.CrmLead2opportunityPartners, error) {
	c := &types.CrmLead2opportunityPartners{}
	return c, svc.client.getAll(types.CrmLead2opportunityPartnerModel, c)
}

func (svc *CrmLead2opportunityPartnerService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.CrmLead2opportunityPartnerModel, fields, relations)
}

func (svc *CrmLead2opportunityPartnerService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CrmLead2opportunityPartnerModel, ids, fields, relations)
}

func (svc *CrmLead2opportunityPartnerService) Delete(ids []int64) error {
	return svc.client.delete(types.CrmLead2opportunityPartnerModel, ids)
}
