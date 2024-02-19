package services

import (
	"encoding/json"
	"fmt"
	"strconv"
	"watchAlert/controllers/repo"
	"watchAlert/globals"
	"watchAlert/models"
	"watchAlert/utils/cmd"
	"watchAlert/utils/http"
)

type AlertDataSourceService struct {
}

type InterAlertDataSourceService interface {
	Create(dataSource models.AlertDataSource) error
	Update(dataSource models.AlertDataSource) error
	Delete(id string) error
	List() ([]models.AlertDataSource, error)
	Get(id, dsType string) []models.AlertDataSource
}

func NewInterAlertDataSourceService() InterAlertDataSourceService {
	return &AlertDataSourceService{}
}

func (adss *AlertDataSourceService) Create(dataSource models.AlertDataSource) error {

	err := adss.Check(dataSource)
	if err != nil {
		return err
	}

	id := "ds-" + cmd.RandId()

	httpStr, _ := json.Marshal(dataSource.HTTPJson)

	data := models.AlertDataSource{
		Id:          id,
		Name:        dataSource.Name,
		Type:        dataSource.Type,
		HTTP:        string(httpStr),
		Enabled:     strconv.FormatBool(dataSource.EnabledBool),
		Description: dataSource.Description,
	}

	err = repo.DBCli.Create(models.AlertDataSource{}, &data)
	if err != nil {
		return err
	}

	return nil

}

func (adss *AlertDataSourceService) Update(dataSource models.AlertDataSource) error {

	httpStr, _ := json.Marshal(dataSource.HTTPJson)

	data := repo.Updates{
		Table: models.AlertDataSource{},
		Where: []string{"id = ?", dataSource.Id},
		Updates: models.AlertDataSource{
			Id:          dataSource.Id,
			Name:        dataSource.Name,
			Type:        dataSource.Type,
			HTTP:        string(httpStr),
			Enabled:     strconv.FormatBool(dataSource.EnabledBool),
			Description: dataSource.Description,
		},
	}

	err := repo.DBCli.Updates(data)
	if err != nil {
		return err
	}

	return nil

}

func (adss *AlertDataSourceService) Delete(id string) error {

	data := repo.Delete{
		Table: models.AlertDataSource{},
		Where: []string{"id = ?", id},
	}

	err := repo.DBCli.Delete(data)
	if err != nil {
		return err
	}

	return nil

}

func (adss *AlertDataSourceService) List() ([]models.AlertDataSource, error) {

	var (
		data []models.AlertDataSource
	)

	globals.DBCli.Find(&data)

	for k, v := range data {
		var httpJson models.HTTP
		_ = json.Unmarshal([]byte(v.HTTP), &httpJson)
		data[k].HTTPJson = httpJson
		data[k].EnabledBool, _ = strconv.ParseBool(v.Enabled)
	}

	return data, nil

}

func (adss *AlertDataSourceService) Get(id, dsType string) []models.AlertDataSource {

	query := "type = ?"
	args := []interface{}{dsType}

	if id != "" {
		query += " AND id = ?"
		args = append(args, id)
	}

	var data []models.AlertDataSource
	err := globals.DBCli.Where(query, args...).Find(&data).Error
	if err != nil {
		return []models.AlertDataSource{}
	}

	for k := range data {
		var httpJson models.HTTP
		if err = json.Unmarshal([]byte(data[k].HTTP), &httpJson); err == nil {
			data[k].HTTPJson = httpJson
		}
		data[k].EnabledBool, _ = strconv.ParseBool(data[k].Enabled)
	}

	return data

}

func (adss *AlertDataSourceService) Check(dataSource models.AlertDataSource) error {

	switch dataSource.Type {
	case "Prometheus":
		path := "/api/v1/format_query?query=foo/bar"
		fullPath := dataSource.HTTPJson.URL + path
		res, err := http.Get(fullPath)
		if err != nil {
			return err
		}
		if res.StatusCode != 200 {
			return fmt.Errorf("StatusCode 非预期 -> %d", res.StatusCode)
		}
	}

	return nil

}