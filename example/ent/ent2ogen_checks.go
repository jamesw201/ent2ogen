// Code generated by ent, DO NOT EDIT.

package ent

import (
	openapi "github.com/jamesw201/go-starter/example/api"
	"github.com/jamesw201/go-starter/example/ent/keycapmodel"
	"github.com/jamesw201/go-starter/example/ent/switchmodel"
)

func _() {
	_ = struct {
		ID       int64
		Name     string
		Switches openapi.Switches
		Keycaps  openapi.Keycaps
		Price    int64
		Discount openapi.NilInt64
	}(openapi.Keyboard{})
}

func _() {
	_ = struct {
		ID       int64
		Name     string
		Profile  string
		Material openapi.KeycapsMaterial
	}(openapi.Keycaps{})
	_ = map[bool]struct{}{
		string(openapi.KeycapsMaterialABS) == string(keycapmodel.MaterialABS): {},
		false: {},
	}
	_ = map[bool]struct{}{
		string(openapi.KeycapsMaterialPBT) == string(keycapmodel.MaterialPBT): {},
		false: {},
	}
}

func _() {
	_ = struct {
		ID         int64
		Name       string
		SwitchType openapi.SwitchesSwitchType
	}(openapi.Switches{})
	_ = map[bool]struct{}{
		string(openapi.SwitchesSwitchTypeMechanical) == string(switchmodel.SwitchTypeMechanical): {},
		false: {},
	}
	_ = map[bool]struct{}{
		string(openapi.SwitchesSwitchTypeOptical) == string(switchmodel.SwitchTypeOptical): {},
		false: {},
	}
	_ = map[bool]struct{}{
		string(openapi.SwitchesSwitchTypeElectrocapacitive) == string(switchmodel.SwitchTypeElectrocapacitive): {},
		false: {},
	}
}
