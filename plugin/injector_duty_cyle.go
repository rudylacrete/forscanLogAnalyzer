package plugin

import (
	"fmt"

	"github.com/rudylacrete/forscanLogAnalyzer/models"
	"github.com/rudylacrete/forscanLogAnalyzer/utils"
)

const injectorField = "fuelpw"
const rpmField = "rpm"

type InjectorDutyCycle struct{}

func NewInjectorDutyCyclePlugin() *InjectorDutyCycle {
	return &InjectorDutyCycle{}
}

func (p *InjectorDutyCycle) Info() string {
	return "Injector plugin calculates injector duty cycle in percent"
}

func (p *InjectorDutyCycle) Transform(logs *models.ForscanLogs) error {
	injectorIndex := utils.ArrayIndex(logs.Fields, injectorField)
	rpmIndex := utils.ArrayIndex(logs.Fields, rpmField)
	if injectorIndex == -1 || rpmIndex == -1 {
		// required field are not here
		return fmt.Errorf("injector pulse width or rpm field missing")
	}
	logs.Fields = append(logs.Fields, "injectorDutyCycle")
	for i, e := range logs.Values {
		if e[injectorIndex] != 0 && e[rpmIndex] != 0 {
			idc := e[injectorIndex] / ((60 / e[rpmIndex]) * 2 * 1000) * 100
			logs.Values[i] = append(e, idc)
		} else {
			// add 0 to keep the same array length
			logs.Values[i] = append(e, 0)
		}
	}
	return nil
}
