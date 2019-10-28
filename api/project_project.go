package api

import (
	"github.com/morezig/go-odoo/types"
)

type ProjectProjectService struct {
	client *Client
}

func NewProjectProjectService(c *Client) *ProjectProjectService {
	return &ProjectProjectService{client: c}
}

func (svc *ProjectProjectService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ProjectProjectModel, name)
}

func (svc *ProjectProjectService) GetByIds(ids []int64) (*types.ProjectProjects, error) {
	p := &types.ProjectProjects{}
	return p, svc.client.getByIds(types.ProjectProjectModel, ids, p)
}

func (svc *ProjectProjectService) GetByName(name string) (*types.ProjectProjects, error) {
	p := &types.ProjectProjects{}
	return p, svc.client.getByName(types.ProjectProjectModel, name, p)
}

func (svc *ProjectProjectService) GetByField(field string, value string) (*types.ProjectProjects, error) {
	p := &types.ProjectProjects{}
	return p, svc.client.getByField(types.ProjectProjectModel, field, value, p)
}

func (svc *ProjectProjectService) GetAll() (*types.ProjectProjects, error) {
	p := &types.ProjectProjects{}
	return p, svc.client.getAll(types.ProjectProjectModel, p)
}

func (svc *ProjectProjectService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ProjectProjectModel, fields, relations)
}

func (svc *ProjectProjectService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProjectProjectModel, ids, fields, relations)
}

func (svc *ProjectProjectService) Delete(ids []int64) error {
	return svc.client.delete(types.ProjectProjectModel, ids)
}
