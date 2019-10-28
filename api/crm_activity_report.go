package api

import (
	"github.com/morezig/go-odoo/types"
)

type CrmActivityReportService struct {
	client *Client
}

func NewCrmActivityReportService(c *Client) *CrmActivityReportService {
	return &CrmActivityReportService{client: c}
}

func (svc *CrmActivityReportService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.CrmActivityReportModel, name)
}

func (svc *CrmActivityReportService) GetByIds(ids []int64) (*types.CrmActivityReports, error) {
	c := &types.CrmActivityReports{}
	return c, svc.client.getByIds(types.CrmActivityReportModel, ids, c)
}

func (svc *CrmActivityReportService) GetByName(name string) (*types.CrmActivityReports, error) {
	c := &types.CrmActivityReports{}
	return c, svc.client.getByName(types.CrmActivityReportModel, name, c)
}

func (svc *CrmActivityReportService) GetByField(field string, value string) (*types.CrmActivityReports, error) {
	c := &types.CrmActivityReports{}
	return c, svc.client.getByField(types.CrmActivityReportModel, field, value, c)
}

func (svc *CrmActivityReportService) GetAll() (*types.CrmActivityReports, error) {
	c := &types.CrmActivityReports{}
	return c, svc.client.getAll(types.CrmActivityReportModel, c)
}

func (svc *CrmActivityReportService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.CrmActivityReportModel, fields, relations)
}

func (svc *CrmActivityReportService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CrmActivityReportModel, ids, fields, relations)
}

func (svc *CrmActivityReportService) Delete(ids []int64) error {
	return svc.client.delete(types.CrmActivityReportModel, ids)
}
