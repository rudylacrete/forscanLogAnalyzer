// this plugin convert raw MAP values to converted turbo pressure
package plugin

import (
	"fmt"
	"math"

	"github.com/rudylacrete/forscanLogAnalyzer/models"
	"github.com/rudylacrete/forscanLogAnalyzer/utils"
)

type TurboUnit string

const Bar = TurboUnit("bar")
const DefaultUnit = Bar

type TurboPlugin struct {
	unit TurboUnit
}

const pluginField = "turbo"

func NewTurboPlugin(unit TurboUnit) *TurboPlugin {
	switch unit {
	case Bar:
		return &TurboPlugin{
			unit: unit,
		}
	default:
		fmt.Printf("This unit is not supported: %s. Use %s (default)", unit, DefaultUnit)
		return &TurboPlugin{
			unit: DefaultUnit,
		}
	}
}

func (p *TurboPlugin) Info() string {
	return "Turbo plugin is used to convert MAP value to human readable turbo pressure"
}

func (p *TurboPlugin) Transform(logs *models.ForscanLogs) error {
	fi := utils.ArrayIndex(logs.Fields, pluginField)
	if fi == -1 {
		// if turbo field is not here, nothing to do
		return nil
	}
	//TODO manage other units
	for _, e := range logs.Values {
		if e[fi] != 0 {
			e[fi] = math.Max((e[fi]-100)/100, 0)
		}
	}
	return nil
}
