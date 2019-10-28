package api

import (
	"github.com/morezig/go-odoo/types"
)

type ReportProductReportPricelistService struct {
	client *Client
}

func NewReportProductReportPricelistService(c *Client) *ReportProductReportPricelistService {
	return &ReportProductReportPricelistService{client: c}
}

func (svc *ReportProductReportPricelistService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ReportProductReportPricelistModel, name)
}

func (svc *ReportProductReportPricelistService) GetByIds(ids []int64) (*types.ReportProductReportPricelists, error) {
	r := &types.ReportProductReportPricelists{}
	return r, svc.client.getByIds(types.ReportProductReportPricelistModel, ids, r)
}

func (svc *ReportProductReportPricelistService) GetByName(name string) (*types.ReportProductReportPricelists, error) {
	r := &types.ReportProductReportPricelists{}
	return r, svc.client.getByName(types.ReportProductReportPricelistModel, name, r)
}

func (svc *ReportProductReportPricelistService) GetByField(field string, value string) (*types.ReportProductReportPricelists, error) {
	r := &types.ReportProductReportPricelists{}
	return r, svc.client.getByField(types.ReportProductReportPricelistModel, field, value, r)
}

func (svc *ReportProductReportPricelistService) GetAll() (*types.ReportProductReportPricelists, error) {
	r := &types.ReportProductReportPricelists{}
	return r, svc.client.getAll(types.ReportProductReportPricelistModel, r)
}

func (svc *ReportProductReportPricelistService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ReportProductReportPricelistModel, fields, relations)
}

func (svc *ReportProductReportPricelistService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ReportProductReportPricelistModel, ids, fields, relations)
}

func (svc *ReportProductReportPricelistService) Delete(ids []int64) error {
	return svc.client.delete(types.ReportProductReportPricelistModel, ids)
}
