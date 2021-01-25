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

type ARPostResult struct {
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
