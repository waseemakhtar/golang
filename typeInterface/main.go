package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type CreateVnfRequest struct {
	VnfdID                 string   `json:"vnfdId"`
	VnfInstanceName        string   `json:"vnfInstanceName"`
	VnfInstanceDescription string   `json:"vnfInstanceDescription,omitempty"`
}

// InstantiateVnfRequest struct
type InstantiateVnfRequest struct {
	VnfInstanceID        string                      `json:"vnfInstanceId,omitempty"`
	FlavourID            string                      `json:"flavourId"`
	InstantiationLevelID string                      `json:"instantiationLevelId,omitempty"`
	VnfProfileResourceID string                      `json:"vnfProfileResourceId,omitempty"`
	LocalizationLanguage string                      `json:"localizationLanguage"`
	AdditionalParams     json.RawMessage             `json:"additionalParams,omitempty"`
	AdditionalParamsV2   json.RawMessage             `json:"additionalParams_v2,omitempty"`
	Extensions           json.RawMessage             `json:"extensions_v2,omitempty"`
}

type VnfDB struct {
	instantiateVnfRequest InstantiateVnfRequest
	createVnfRequest CreateVnfRequest
}

func (d *VnfDB) SetVnfRequest(req interface{}) {
	v := reflect.ValueOf(req)

	switch v.Type().String() {
	case "main.CreateVnfRequest":
		d.createVnfRequest =  v.Interface().(CreateVnfRequest)
	case "main.InstantiateVnfRequest":
		d.instantiateVnfRequest = v.Interface().(InstantiateVnfRequest)
	}

}


func main() {
	vnfDB := &VnfDB{}
	fmt.Println("%#v", vnfDB)

	createVnfReq := CreateVnfRequest{VnfdID:"WaseemVnf"}
	instantiateVnfReq := InstantiateVnfRequest{FlavourID:"WaseemFlavor"}
	vnfDB.SetVnfRequest(createVnfReq)
	fmt.Printf("After-1: %#v\n", vnfDB)
	vnfDB.SetVnfRequest(instantiateVnfReq)
	fmt.Printf("After-2 %#v\n", vnfDB)
}
