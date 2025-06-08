// This file is auto-generated, don't edit it. Thanks.
package main

import (
	"log"
	"os"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	credential "github.com/aliyun/credentials-go/credentials"
	"github.com/joho/godotenv"
)

// Description:
//
// # Initialize the Client with the credentials
//
// @return Client
//
// @throws Exception
func CreateClient() (_result *openapi.Client, _err error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found or failed to load")
	}
	// It is recommended to use the default credential. For more credentials, please refer to: https://help.aliyun.com/document_detail/378661.html.
	credential, _err := credential.NewCredential(nil)
	if _err != nil {
		return _result, _err
	}

	config := &openapi.Config{
		Credential: credential,
	}
	// Get MongoDB connection string from environment
	aliCloudID := os.Getenv("ALIBABA_CLOUD_ACCOUNT_ID")
	if aliCloudID == "" {
		log.Fatal("ALIBABA_CLOUD_ACCOUNT_ID not set in environment")
	}
	// See https://api.alibabacloud.com/product/FC.
	config.Endpoint = tea.String(aliCloudID + ".ap-southeast-1.fc.aliyuncs.com")
	_result = &openapi.Client{}
	_result, _err = openapi.NewClient(config)
	return _result, _err
}

// Description:
//
// # API Info
//
// @param path - string Path parameters
//
// @return OpenApi.Params
func CreateApiInfo() (_result *openapi.Params) {
	params := &openapi.Params{
		// API Name
		Action: tea.String("DescribeRegions"),
		// API Version
		Version: tea.String("2023-03-30"),
		// Protocol
		Protocol: tea.String("HTTPS"),
		// HTTP Method
		Method:   tea.String("GET"),
		AuthType: tea.String("AK"),
		Style:    tea.String("FC"),
		// API PATH
		Pathname: tea.String("/2023-03-30/regions"),
		// Request body content format
		ReqBodyType: tea.String("json"),
		// Response body content format
		BodyType: tea.String("json"),
	}
	_result = params
	return _result
}

func _main(args []*string) (_err error) {
	client, _err := CreateClient()
	if _err != nil {
		return _err
	}

	params := CreateApiInfo()
	// runtime options
	runtime := &util.RuntimeOptions{}
	request := &openapi.OpenApiRequest{}
	// Copy the code to run, please print the return value of the API by yourself.
	// The return value is of Map type, and three types of data can be obtained from Map: response body, response headers, HTTP status code.
	resp, _err := client.CallApi(params, request, runtime)
	if _err != nil {
		return _err
	}

	console.Log(util.ToJSONString(resp))
	return _err
}

func main() {
	err := _main(tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
}
