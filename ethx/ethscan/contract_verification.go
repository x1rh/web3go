package ethscan

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"net/url"
)

// type VerifyContractReq struct {
// 	Address        string            `json:"address"`
// 	Chain          string            `json:"chain"`
// 	Files          map[string]string `json:"files"`
// 	CreatorTxHash  string            `json:"creatorTxHash"`
// 	ChosenContract string            `json:"chosenContract"`
// }


type VerifyContractReq struct {
	URL, apiKey, sourceCode, contractAddress, contractName, metadataJson string
}

type VerifyContractResp struct {
	Result []struct {
		Address    string `json:"address,omitempty"`
		ChainId    string `json:"chainId,omitempty"`
		Status     string `json:"status,omitempty"`
		LibraryMap struct {
			Lib1 string `json:"lib1,omitempty"`
			Lib2 string `json:"lib2,omitempty"`
		} `json:"libraryMap"`
	} `json:"result,omitempty"`
	Error string `json:"error,omitempty"`
}

type VerifyResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}


func VerifyContract(URL, apiKey, sourceCode, contractAddress, contractName, metadataJson string) (string, error) {
	//logx.Debugf("_url: |%s|\n", _url)
	//logx.Debugf("apiKey: |%s|\n", apiKey)
	//logx.Debugf("sourceCode: |%s|\n", sourceCode)
	//logx.Debugf("contractAddress: |%s|\n", contractAddress)
	//logx.Debugf("contractName: |%s|\n", contractName)

	compilerVersion := "v0.8.22+commit.4fc1097e"
	runs := "200"
	optimizationUsed := "1"

	data := url.Values{}
	data.Set("apiKey", apiKey)
	data.Set("module", "contract")
	data.Set("action", "verifysourcecode")
	data.Set("sourceCode", sourceCode)
	data.Set("contractaddress", contractAddress)
	data.Set("codeformat", "solidity-single-file")
	data.Set("contractname", contractName)
	data.Set("compilerversion", compilerVersion)
	data.Set("optimizationUsed", optimizationUsed)
	data.Set("runs", runs)

	// only for payTokenType = 0
	encodedArgs := "0000000000000000000000000000000000000000000000000000000000000000"
	data.Set("constructorArguements", encodedArgs)

	//logx.Debug("args: ", encodedArgs)

	//constructorArguements:
	//data.Set("evmversion", "paris")
	//data.Set("licenseType", "1")

	payload := bytes.NewBufferString(data.Encode())

	resp, err := http.Post(URL, "application/x-www-form-urlencoded", payload)
	if err != nil {
		logx.Error(err)
		return "", err
	}
	defer resp.Body.Close()

	body := new(bytes.Buffer)
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		logx.Error("err:", err)
		return "", err
	}

	// todo: check response status
	fmt.Println("response status:", resp.Status)
	fmt.Println("response body:", body.String())

	var res VerifyResp
	err = json.Unmarshal(body.Bytes(), &res)
	if err != nil {
		logx.Error(err)
		return "", err
	}

	//{"status":"0","message":"NOTOK","result":"Invalid constructor arguments provided. Please verify that they are in ABI-encoded format"}
	if res.Status == "0" && res.Message == "NOTOK" {
		return "", errors.New(res.Result)
	}

	return res.Result, nil
}

func QueryResult() error {
	_url := "https://api-goerli.etherscan.io/api"

	data := url.Values{}
	data.Set("module", "contract")
	data.Set("action", "checkverifystatus")
	data.Set("apikey", "QEAE2M96IB94MVPUN7ESQEBNI416F1EWRR")
	data.Set("guid", "ciqqtpiziyenbwmrlizzlu1wbxdgcxsg1lsygqdtrs6ikbbvta")

	payload := bytes.NewBufferString(data.Encode())
	resp, err := http.Post(_url, "application/x-www-form-urlencoded", payload)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body := new(bytes.Buffer)
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("response status:", resp.Status)
	fmt.Println("response body:", body.String())
	return nil
}
