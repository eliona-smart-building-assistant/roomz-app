package eliona

import (
	"fmt"
	assetmodel "roomz/model/asset"
	confmodel "roomz/model/conf"

	"github.com/eliona-smart-building-assistant/go-eliona/asset"
	"github.com/eliona-smart-building-assistant/go-utils/log"
)

const ClientReference string = "roomz"

func UpsertAssetData(config confmodel.Configuration, a assetmodel.Sensor) error {
	for _, projectId := range config.ProjectIDs {
		log.Debug("Eliona", "upserting data for asset: config %d and asset '%v'", config.Id, a.GetGAI())
		assetId, err := a.GetAssetID(projectId)
		if err != nil {
			return err
		}
		if assetId == nil {
			// This might happen in case of filtered or newly added devices.
			log.Debug("conf", "unable to find asset ID for %v", a.GetGAI())
			continue
		}

		data := asset.Data{
			AssetId:         *assetId,
			Data:            a,
			ClientReference: ClientReference,
		}
		if err := asset.UpsertAssetDataIfAssetExists(data); err != nil {
			return fmt.Errorf("upserting data: %v", err)
		}
	}
	return nil
}
