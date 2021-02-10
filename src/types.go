package main

import (
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}
type Config struct {
	CRUDHost        string
	CRUDPort        string
	ListenServePort string
}

type toAssetRegisterResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Asset struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}

type AssetFlexVal struct {
	ID        string `json:"id,omitempty"`
	AssetID   string `json:"assetid,omitempty"`
	Flexfldid string `json:"flexfldid,omitempty"`
	Value     string `json:"value,omitempty"`
}

type ObservationFlexVal struct {
	ID        string `json:"id,omitempty"`
	AssetID   string `json:"assetid,omitempty"`
	Flexfldid string `json:"flexfldid,omitempty"`
	Value     string `json:"value,omitempty"`
}

type toAssetRegister struct {
	ID                          string               `json:"id"`
	Name                        string               `json:"name"`
	AssetType                   string               `json:"assettype"`
	CompatibleUnitID            string               `json:"compatibleunitid"`
	DerecognitionDate           string               `json:"derecognitiondate,omitempty"`
	DerecognitionValue          string               `json:"derecognitionvalue,omitempty"`
	Description                 string               `json:"description,omitempty"`
	Dimension1Value             string               `json:"dimension1value,omitempty"`
	Dimension2Value             string               `json:"dimension2value,omitempty"`
	Dimension3Value             string               `json:"dimension3value,omitempty"`
	Dimension4Value             string               `json:"dimension4value,omitempty"`
	Dimension5Value             string               `json:"dimension5value,omitempty"`
	Extent                      string               `json:"extent,omitempty"`
	ExtentConfidence            string               `json:"extentconfidence,omitempty"`
	ManufactureDate             string               `json:"manufacturedate,omitempty"`
	ManufactureDateConfidence   string               `json:"manufacturedateconfidence,omitempty"`
	TakeOnDate                  string               `json:"takeondate,omitempty"`
	SerialNo                    string               `json:"serialno,omitempty"`
	Latitude                    string               `json:"lat"`
	Longitude                   string               `json:"lon"`
	Geom                        string               `json:"geom"`
	FunclocID                   string               `json:"funclocid"`
	InstallDate                 string               `json:"installdate"`
	Status                      string               `json:"status,omitempty"`
	Age                         string               `json:"age,omitempty"`
	CarryingValueClosingBalance string               `json:"carryingvalueclosingbalance,omitempty"`
	CarryingValueOpeningBalance string               `json:"carryingvalueopeningbalance,omitempty"`
	CostClosingBalance          string               `json:"costclosingbalance,omitempty"`
	CostOpeningBalance          string               `json:"costopeningbalance,omitempty"`
	Crc                         string               `json:"crc,omitempty"`
	DepreciationClosingBalance  string               `json:"depreciationclosingbalance,omitempty"`
	DepreciationOpeningBalance  string               `json:"depreciationopeningbalance,omitempty"`
	ImpairmentClosingBalance    string               `json:"impairmentclosingbalance,omitempty"`
	ImpairmentOpeningBalance    string               `json:"impairmentopeningbalance,omitempty"`
	ResidualValue               string               `json:"residualvalue,omitempty"`
	RulYears                    string               `json:"rulyears,omitempty"`
	Drc                         string               `json:"drc,omitempty"`
	Fy                          string               `json:"fy,omitempty"`
	AssetValID                  string               `json:"assetvalid"`
	FlvList                     []AssetFlexVal       `json:"assetflexvals"`
	OFlvList                    []ObservationFlexVal `json:"observationflexvals"`
}

type toAssetRegsiterList struct {
	Alist []toAssetRegister `json:"assets"`
}

type FunclocFlexVal struct {
	ID        string `json:"id,omitempty"`
	FunclocID string `json:"funclocid,omitempty"`
	Flexfldid string `json:"flexfldid,omitempty"`
	Value     string `json:"value,omitempty"`
}

type Funcloc struct {
	FunclocID   string            `json:"funclocid,omitempty"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Latitude    string            `json:"latitude,omitempty"`
	Longitude   string            `json:"longitude,omitempty"`
	Geom        string            `json:"geom,omitempty"`
	FLNlist     []FunclocNode     `json:"funclocnodes,omitempty"`
	FLFVlist    []FunclocFlexVal  `json:"funclocflexvals,omitempty"`
	Alist       []toAssetRegister `json:"assets,omitempty"`
}

type FunclocList struct {
	Flist []Funcloc `json:"funcloc"`
}

type FunclocNodeFlexVal struct {
	ID            string `json:"id,omitempty"`
	FunclocNodeID string `json:"funclocnodeid,omitempty"`
	Flexfldid     string `json:"flexfldid,omitempty"`
	Value         string `json:"value,omitempty"`
}
type FunclocNode struct {
	ID         string               `json:"funclocnodeid,omitempty"`
	Name       string               `json:"name,omitempty"`
	AliasName  string               `json:"aliasname,omitempty"`
	Latitude   string               `json:"latitude,omitempty"`
	Longitude  string               `json:"longitude,omitempty"`
	Geom       string               `json:"geom,omitempty"`
	NodeTypeID string               `json:"nodetypeid,omitempty"`
	ParentID   string               `json:"parentid,omitempty"`
	FLNFVlist  []FunclocNodeFlexVal `json:"funclocnodeflexvals,omitempty"`
}

type FunclocNodeList struct {
	Fnodelist []FunclocNode `json:"funclocnode"`
}

type ARPostResult struct {
	Error                     string  `json:"error,omitempty"`
	FunclocSuccess            bool    `json:"funclocsuccess,omitempty"`
	FunclocMessage            string  `json:"funclocmessage,omitempty"`
	FunclocID                 string  `json:"funclocid,omitempty"`
	FunclocflexvalSuccess     bool    `json:"funclocflexvalsuccess,omitempty"`
	FunclocflexvalMessage     string  `json:"funclocflexvalmessage,omitempty"`
	FunclocnodeSuccess        bool    `json:"funclocnodesuccess,omitempty"`
	FunclocnodeMessage        string  `json:"funclocnodemessage,omitempty"`
	FunclocnodeID             string  `json:"funclocnodeid,omitempty"`
	FuncloclinkSuccess        bool    `json:"funcloclinksuccess,omitempty"`
	FuncloclinkMessage        string  `json:"funcloclinkmessage,omitempty"`
	FunclocnodeflexvalSuccess bool    `json:"funclocnodeflexvalsuccess,omitempty"`
	FunclocnodeflexvalMessage string  `json:"funclocnodeflexvalmessage,omitempty"`
	AssetSuccess              bool    `json:"assetsuccess,omitempty"`
	AssetMessage              string  `json:"assetmessage,omitempty"`
	PostedAssetList           []Asset `json:"postedassets,omitempty"`
	AssetflexvalSuccess       bool    `json:"assetflexvalsuccess,omitempty"`
	AssetflexvalMessage       string  `json:"assetflexvalmessage,omitempty"`
	ObservationflexvalSuccess bool    `json:"observationflexvalsuccess,omitempty"`
	ObservationflexvalMessage string  `json:"observationflexvalmessage,omitempty"`
}

type ErrorResult struct {
	Message string `json:"message,omitempty"`
}

type ResponseList struct {
	ErrorList    []ErrorResult  `json:"errorresponse,omitempty"`
	ResponseList []ARPostResult `json:"result,omitempty"`
}
