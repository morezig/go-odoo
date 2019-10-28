package api

import (
	"github.com/morezig/go-odoo/types"
)

type BaseImportTestsModelsCharStillreadonlyService struct {
	client *Client
}

func NewBaseImportTestsModelsCharStillreadonlyService(c *Client) *BaseImportTestsModelsCharStillreadonlyService {
	return &BaseImportTestsModelsCharStillreadonlyService{client: c}
}

func (svc *BaseImportTestsModelsCharStillreadonlyService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.BaseImportTestsModelsCharStillreadonlyModel, name)
}

func (svc *BaseImportTestsModelsCharStillreadonlyService) GetByIds(ids []int64) (*types.BaseImportTestsModelsCharStillreadonlys, error) {
	b := &types.BaseImportTestsModelsCharStillreadonlys{}
	return b, svc.client.getByIds(types.BaseImportTestsModelsCharStillreadonlyModel, ids, b)
}

func (svc *BaseImportTestsModelsCharStillreadonlyService) GetByName(name string) (*types.BaseImportTestsModelsCharStillreadonlys, error) {
	b := &types.BaseImportTestsModelsCharStillreadonlys{}
	return b, svc.client.getByName(types.BaseImportTestsModelsCharStillreadonlyModel, name, b)
}

func (svc *BaseImportTestsModelsCharStillreadonlyService) GetByField(field string, value string) (*types.BaseImportTestsModelsCharStillreadonlys, error) {
	b := &types.BaseImportTestsModelsCharStillreadonlys{}
	return b, svc.client.getByField(types.BaseImportTestsModelsCharStillreadonlyModel, field, value, b)
}

func (svc *BaseImportTestsModelsCharStillreadonlyService) GetAll() (*types.BaseImportTestsModelsCharStillreadonlys, error) {
	b := &types.BaseImportTestsModelsCharStillreadonlys{}
	return b, svc.client.getAll(types.BaseImportTestsModelsCharStillreadonlyModel, b)
}

func (svc *BaseImportTestsModelsCharStillreadonlyService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.BaseImportTestsModelsCharStillreadonlyModel, fields, relations)
}

func (svc *BaseImportTestsModelsCharStillreadonlyService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseImportTestsModelsCharStillreadonlyModel, ids, fields, relations)
}

func (svc *BaseImportTestsModelsCharStillreadonlyService) Delete(ids []int64) error {
	return svc.client.delete(types.BaseImportTestsModelsCharStillreadonlyModel, ids)
}
