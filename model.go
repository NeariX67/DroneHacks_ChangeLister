package main

import "fmt"

var (
	ErrInvalidFilterName  = fmt.Errorf("argument filterName is not a valid regular expression")
	ErrInvalidFilterIndex = fmt.Errorf("argument filterIndex is not a valid regular expression")
	ErrMissingFileName    = fmt.Errorf("missing fileName parameter")
)

type flags struct {
	fileName    string
	filterName  string
	filterIndex string
	difference  float64
}

type parameterConfigs []struct {
	parameterConfig
}

type parameterConfig struct {
	TableNo    int     `json:"table_no"`
	ParamIndex int     `json:"param_index"`
	Name       string  `json:"name"`
	Min        float64 `json:"min"`
	Max        float64 `json:"max"`
	Default    float64 `json:"default"`
	Value      string  `json:"value"`
	Changed    bool    `json:"changed"`
	TypeID     int     `json:"type_id"`
}

func (param *parameterConfig) toString() string {
	return fmt.Sprintf("Value: %s, Default: %f, Name: %s, Min: %f, Max: %f, Changed: %t, TableNo: %d, ParamIndex: %d, TypeID: %d", param.Value, param.Default, param.Name, param.Min, param.Max, param.Changed, param.TableNo, param.ParamIndex, param.TypeID)
}
