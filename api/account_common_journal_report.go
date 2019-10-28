package api

import (
	"github.com/morezig/go-odoo/types"
)

type AccountCommonJournalReportService struct {
	client *Client
}

func NewAccountCommonJournalReportService(c *Client) *AccountCommonJournalReportService {
	return &AccountCommonJournalReportService{client: c}
}

func (svc *AccountCommonJournalReportService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountCommonJournalReportModel, name)
}

func (svc *AccountCommonJournalReportService) GetByIds(ids []int64) (*types.AccountCommonJournalReports, error) {
	a := &types.AccountCommonJournalReports{}
	return a, svc.client.getByIds(types.AccountCommonJournalReportModel, ids, a)
}

func (svc *AccountCommonJournalReportService) GetByName(name string) (*types.AccountCommonJournalReports, error) {
	a := &types.AccountCommonJournalReports{}
	return a, svc.client.getByName(types.AccountCommonJournalReportModel, name, a)
}

func (svc *AccountCommonJournalReportService) GetByField(field string, value string) (*types.AccountCommonJournalReports, error) {
	a := &types.AccountCommonJournalReports{}
	return a, svc.client.getByField(types.AccountCommonJournalReportModel, field, value, a)
}

func (svc *AccountCommonJournalReportService) GetAll() (*types.AccountCommonJournalReports, error) {
	a := &types.AccountCommonJournalReports{}
	return a, svc.client.getAll(types.AccountCommonJournalReportModel, a)
}

func (svc *AccountCommonJournalReportService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountCommonJournalReportModel, fields, relations)
}

func (svc *AccountCommonJournalReportService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountCommonJournalReportModel, ids, fields, relations)
}

func (svc *AccountCommonJournalReportService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountCommonJournalReportModel, ids)
}
