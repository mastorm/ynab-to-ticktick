package internal

import (
	ynab "github.com/brunomvsouza/ynab.go"
	"github.com/mastorm/ynab-to-ticktick/pkg/ticktick"
)

type Application struct {
	Ticktick *ticktick.Client
	Ynab     ynab.ClientServicer
}

type ConfigProvider interface {
	Get(key string) (string, error)
}

const (
	ticktickClientId    = "TICKTICK_CLIENT_ID"
	ticktickAccessToken = "TICKTICK_ACCESS_TOKEN"
	ynabAccessToken     = "YNAB_ACCESS_TOKEN"
)

func buildTicktickClient(configProvider ConfigProvider) (*ticktick.Client, error) {
	clientId, err := configProvider.Get(ticktickClientId)
	if err != nil {
		return nil, err
	}

	accessToken, err := configProvider.Get(ticktickAccessToken)
	if err != nil {
		return nil, err
	}

	args := ticktick.ClientArgs{
		ClientId:    clientId,
		AccessToken: accessToken,
		Scopes:      []ticktick.Scope{ticktick.TasksRead, ticktick.TasksWrite},
	}

	return ticktick.NewClient(args), nil
}

func buildYnabClient(configProvider ConfigProvider) (ynab.ClientServicer, error) {
	accessToken, err := configProvider.Get(ynabAccessToken)
	if err != nil {
		return nil, err
	}

	return ynab.NewClient(accessToken), nil
}

func NewApplication(configProvider ConfigProvider) (*Application, error) {
	ticktick, err := buildTicktickClient(configProvider)
	if err != nil {
		return nil, err
	}

	ynab, err := buildYnabClient(configProvider)
	if err != nil {
		return nil, err
	}

	app := &Application{
		Ticktick: ticktick,
		Ynab:     ynab,
	}

	return app, err
}

func (a *Application) SyncTransactions() {

}
