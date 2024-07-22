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

package assetmodel

import (
	"context"
	"fmt"
	"roomz/conf"
	confmodel "roomz/model/conf"

	"github.com/eliona-smart-building-assistant/go-eliona/asset"
)

type Sensor struct {
	ID       string `eliona:"id" subtype:"info"`
	Presence int8   `eliona:"presence" subtype:"input"`

	Config *confmodel.Configuration
}

func (d *Sensor) GetName() string {
	return d.ID
}

func (d *Sensor) GetDescription() string {
	return "ROOMZ presence sensor"
}

func (d *Sensor) GetAssetType() string {
	return "roomz_presence_sensor"
}

func (d *Sensor) GetGAI() string {
	return d.GetAssetType() + "_" + d.ID
}

func (d *Sensor) GetAssetID(projectID string) (*int32, error) {
	return conf.GetAssetId(context.Background(), *d.Config, projectID, d.GetGAI())
}

func (d *Sensor) SetAssetID(assetID int32, projectID string) error {
	if err := conf.InsertAsset(context.Background(), *d.Config, projectID, d.GetGAI(), assetID, d.ID); err != nil {
		return fmt.Errorf("inserting asset to config db: %v", err)
	}
	return nil
}

func (d *Sensor) GetLocationalChildren() []asset.LocationalNode {
	return []asset.LocationalNode{}
}

func (d *Sensor) GetFunctionalChildren() []asset.FunctionalNode {
	return []asset.FunctionalNode{}
}

type Root struct {
	Sensors []Sensor

	Config *confmodel.Configuration
}

func (r *Root) GetName() string {
	return "ROOMZ"
}

func (r *Root) GetDescription() string {
	return "Root asset for ROOMZ devices"
}

func (r *Root) GetAssetType() string {
	return "roomz_root"
}

func (r *Root) GetGAI() string {
	return r.GetAssetType()
}

func (r *Root) GetAssetID(projectID string) (*int32, error) {
	return conf.GetAssetId(context.Background(), *r.Config, projectID, r.GetGAI())
}

func (r *Root) SetAssetID(assetID int32, projectID string) error {
	if err := conf.InsertAsset(context.Background(), *r.Config, projectID, r.GetGAI(), assetID, ""); err != nil {
		return fmt.Errorf("inserting asset to config db: %v", err)
	}
	return nil
}

func (r *Root) GetLocationalChildren() []asset.LocationalNode {
	locationalChildren := make([]asset.LocationalNode, 0, len(r.Sensors))
	for i := range r.Sensors {
		locationalChildren = append(locationalChildren, &r.Sensors[i])
	}
	return locationalChildren
}

func (r *Root) GetFunctionalChildren() []asset.FunctionalNode {
	functionalChildren := make([]asset.FunctionalNode, 0, len(r.Sensors))
	for i := range r.Sensors {
		functionalChildren = append(functionalChildren, &r.Sensors[i])
	}
	return functionalChildren
}
