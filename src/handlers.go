package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *Server) handlePostToAssetRegister() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//get JSON payload.
		body, err := ioutil.ReadAll(r.Body)
		funclocList := FunclocList{}
		err = json.Unmarshal(body, &funclocList)

		//response variable for posted assets
		assetresponse := ResponseList{}
		assetresponse.ResponseList = []ARPostResult{}

		//response variable for crud errors recieved
		assetresponse.ErrorList = []ErrorResult{}

		//For loop which deals with
		for _, element := range funclocList.Flist {
			//	for i := 0; i < len(funclocList.Flist); i++ {
			//fmt.Println("JSON:", funclocList.Flist[i])
			js, jserr1 := json.Marshal(element)
			if jserr1 != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to create JSON result object from DB result to post Funcloc flex value.")
				return
			}
			//post to crud service
			req, respErr := http.Post("http://"+config.CRUDHost+":"+config.CRUDPort+"/toAssetRegister", "application/json", bytes.NewBuffer(js))
			//check for response error of 500
			if respErr != nil {
				w.WriteHeader(500)
				fmt.Fprint(w, respErr.Error())
				fmt.Println("Error in communication with CRUD service endpoint for request to post assets")
				return
			}

			if req.StatusCode == 500 {
				w.WriteHeader(500)
				bodyBytes, err := ioutil.ReadAll(req.Body)
				if err != nil {
					log.Fatal(err)
				}
				bodyString := string(bodyBytes)
				fmt.Fprintf(w, "Request to DB can't be completed...  "+bodyString)
				fmt.Println("Request to DB can't be completed...  " + bodyString)
				return
			}
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprint(w, err.Error())
				fmt.Println("Posting is not able to be completed by internal error")
				return
			}
			//close the request
			defer req.Body.Close()
			//create new response struct
			var postAssetsResponse ARPostResult

			//decode request into decoder which converts to the struct.
			decoder := json.NewDecoder(req.Body)
			err = decoder.Decode(&postAssetsResponse)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprint(w, err.Error())
				fmt.Println("Error occured in decoding Posting response")
				return
			}

			assetresponse.ResponseList = append(assetresponse.ResponseList, ARPostResult{postAssetsResponse.Error, postAssetsResponse.FunclocSuccess, postAssetsResponse.FunclocMessage, postAssetsResponse.FunclocID, postAssetsResponse.FunclocflexvalSuccess, postAssetsResponse.FunclocflexvalMessage, postAssetsResponse.FunclocnodeSuccess, postAssetsResponse.FunclocnodeMessage, postAssetsResponse.FunclocnodeID, postAssetsResponse.FuncloclinkSuccess, postAssetsResponse.FuncloclinkMessage, postAssetsResponse.FunclocnodeflexvalSuccess, postAssetsResponse.FunclocnodeflexvalMessage, postAssetsResponse.AssetSuccess, postAssetsResponse.AssetMessage, postAssetsResponse.PostedAssetList, postAssetsResponse.AssetflexvalSuccess, postAssetsResponse.AssetflexvalMessage, postAssetsResponse.ObservationflexvalSuccess, postAssetsResponse.ObservationflexvalMessage})

		}

		//convert struct back to JSON
		js, jserr4 := json.Marshal(assetresponse)

		if jserr4 != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON result object from DB result to update Asset.")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
