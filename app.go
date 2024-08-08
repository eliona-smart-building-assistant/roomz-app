//  This file is part of the Eliona project.
//  Copyright © 2024 IoTEC AG. All Rights Reserved.
//  ______ _ _
// |  ____| (_)
// | |__  | |_  ___  _ __   __ _
// |  __| | | |/ _ \| '_ \ / _` |
// | |____| | | (_) | | | | (_| |
// |______|_|_|\___/|_| |_|\__,_|
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING
//  BUT NOT LIMITED  TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
//  NON INFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
//  DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"context"
	"net/http"
	"roomz/apiserver"
	"roomz/apiservices"
	"roomz/conf"
	"roomz/eliona"
	assetmodel "roomz/model/asset"
	confmodel "roomz/model/conf"
	"roomz/roomz"
	"sync"

	"github.com/eliona-smart-building-assistant/go-eliona/app"
	"github.com/eliona-smart-building-assistant/go-eliona/asset"
	"github.com/eliona-smart-building-assistant/go-eliona/dashboard"
	"github.com/eliona-smart-building-assistant/go-eliona/frontend"
	"github.com/eliona-smart-building-assistant/go-utils/common"
	"github.com/eliona-smart-building-assistant/go-utils/db"
	utilshttp "github.com/eliona-smart-building-assistant/go-utils/http"
	"github.com/eliona-smart-building-assistant/go-utils/log"
)

func initialization() {
	ctx := context.Background()

	// Necessary to close used init resources
	conn := db.NewInitConnectionWithContextAndApplicationName(ctx, app.AppName())
	defer conn.Close(ctx)

	// Init the app before the first run.
	app.Init(conn, app.AppName(),
		app.ExecSqlFile("conf/init.sql"),
		asset.InitAssetTypeFiles("resources/asset-types/*.json"),
		dashboard.InitWidgetTypeFiles("resources/widget-types/*.json"),
	)
}

var once sync.Once

func collectData() {
	configs, err := conf.GetConfigs(context.Background())
	if err != nil {
		log.Fatal("conf", "Couldn't read configs from DB: %v", err)
		return
	}
	if len(configs) == 0 {
		once.Do(func() {
			log.Info("conf", "No configs in DB. Please configure the app in Eliona.")
		})
		return
	}

	for _, config := range configs {
		if !config.Enable {
			if config.Active {
				conf.SetConfigActiveState(context.Background(), config, false)
			}
			continue
		}

		if !config.Active {
			conf.SetConfigActiveState(context.Background(), config, true)
			log.Info("conf", "Collecting initialized with Configuration %d:\n"+
				"Enable: %t\n"+
				"Project IDs: %v\n",
				config.Id,
				config.Enable,
				config.ProjectIDs)
		}

		common.RunOnceWithParam(startWebhookListener, config, 1) // We allow only one listener to work (we have only one port available).
	}
}

func startWebhookListener(config confmodel.Configuration) {
	handlePresenceChange := func(workspaceId string, presenceStatus roomz.PresenceStatus) error {
		var presence int8
		switch presenceStatus {
		case roomz.Free:
			presence = 0
		case roomz.Busy:
			presence = 1
		default:
			log.Warn("roomz", "received unsupported state: %v", presenceStatus)
		}
		sensor := assetmodel.Sensor{
			ID:       workspaceId,
			Presence: presence,
			Config:   &config,
		}
		root := assetmodel.Root{Config: &config, Sensors: []assetmodel.Sensor{sensor}}
		if err := eliona.CreateAssets(config, &root); err != nil {
			log.Error("eliona", "creating assets for root %+v: %v", root, err)
			return err
		}

		if err := eliona.UpsertAssetData(config, sensor); err != nil {
			log.Error("eliona", "upserting data for sensor %+v: %v", sensor, err)
			return err
		}

		return nil
	}

	roomz.StartWebhookListener(config.Secret, handlePresenceChange)
}

// listenApi starts the API server and listen for requests
func listenApi() {
	err := http.ListenAndServe(":"+common.Getenv("API_SERVER_PORT", "3000"),
		frontend.NewEnvironmentHandler(
			utilshttp.NewCORSEnabledHandler(
				apiserver.NewRouter(
					apiserver.NewConfigurationAPIController(apiservices.NewConfigurationAPIService()),
					apiserver.NewVersionAPIController(apiservices.NewVersionAPIService()),
					apiserver.NewCustomizationAPIController(apiservices.NewCustomizationAPIService()),
				))))
	log.Fatal("main", "API server: %v", err)
}
