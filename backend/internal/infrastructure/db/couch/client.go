package couch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/dKariakin/purser/internal/domain/model"
	"github.com/dKariakin/purser/internal/infrastructure/db/dto"
)

type Config struct {
	DbHost    string
	TableName string
}

type CouchClient struct {
	httpClient *http.Client
	conf       Config
}

func New(hc *http.Client, c Config) *CouchClient {
	return &CouchClient{
		httpClient: hc,
		conf:       c,
	}
}

func (c *CouchClient) CreateExpense(expense model.Expense) (model.Expense, error) {
	req := c.createItemRequest(dto.FromDomain(expense).Reader())
	res, err := c.queryDb(req)
	if err != nil {
		return expense, err
	}

	if !res.Ok {
		return expense, errors.New("Couldn't create a new expense record in DB")
	}

	expense.Id = res.Id

	return expense, nil
}

func (c *CouchClient) createItemRequest(body io.Reader) *http.Request {
	path := fmt.Sprintf("%s/%s", c.conf.DbHost, c.conf.TableName)
	req, err := http.NewRequest(http.MethodPut, path, body)
	if err != nil {
		return nil
	}
	req.Header.Add("Content-Type", "application/json")

	return req
}

func (c *CouchClient) queryDb(req *http.Request) (dto.DbResponse, error) {
	var body []byte
	resp := dto.DbResponse{}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return resp, err
	}
	res.Body.Read(body)

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
