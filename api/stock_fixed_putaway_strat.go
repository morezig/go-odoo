package api

import (
	"github.com/morezig/go-odoo/types"
)

type StockFixedPutawayStratService struct {
	client *Client
}

func NewStockFixedPutawayStratService(c *Client) *StockFixedPutawayStratService {
	return &StockFixedPutawayStratService{client: c}
}

func (svc *StockFixedPutawayStratService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.StockFixedPutawayStratModel, name)
}

func (svc *StockFixedPutawayStratService) GetByIds(ids []int64) (*types.StockFixedPutawayStrats, error) {
	s := &types.StockFixedPutawayStrats{}
	return s, svc.client.getByIds(types.StockFixedPutawayStratModel, ids, s)
}

func (svc *StockFixedPutawayStratService) GetByName(name string) (*types.StockFixedPutawayStrats, error) {
	s := &types.StockFixedPutawayStrats{}
	return s, svc.client.getByName(types.StockFixedPutawayStratModel, name, s)
}

func (svc *StockFixedPutawayStratService) GetByField(field string, value string) (*types.StockFixedPutawayStrats, error) {
	s := &types.StockFixedPutawayStrats{}
	return s, svc.client.getByField(types.StockFixedPutawayStratModel, field, value, s)
}

func (svc *StockFixedPutawayStratService) GetAll() (*types.StockFixedPutawayStrats, error) {
	s := &types.StockFixedPutawayStrats{}
	return s, svc.client.getAll(types.StockFixedPutawayStratModel, s)
}

func (svc *StockFixedPutawayStratService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.StockFixedPutawayStratModel, fields, relations)
}

func (svc *StockFixedPutawayStratService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockFixedPutawayStratModel, ids, fields, relations)
}

func (svc *StockFixedPutawayStratService) Delete(ids []int64) error {
	return svc.client.delete(types.StockFixedPutawayStratModel, ids)
}
