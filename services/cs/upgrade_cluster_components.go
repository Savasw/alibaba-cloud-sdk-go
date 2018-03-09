package cs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// invoke UpgradeClusterComponents api with *UpgradeClusterComponentsRequest synchronously
// api document: https://help.aliyun.com/api/cs/upgradeclustercomponents.html
func (client *Client) UpgradeClusterComponents(request *UpgradeClusterComponentsRequest) (response *UpgradeClusterComponentsResponse, err error) {
	response = CreateUpgradeClusterComponentsResponse()
	err = client.DoAction(request, response)
	return
}

// invoke UpgradeClusterComponents api with *UpgradeClusterComponentsRequest asynchronously
// api document: https://help.aliyun.com/api/cs/upgradeclustercomponents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpgradeClusterComponentsWithChan(request *UpgradeClusterComponentsRequest) (<-chan *UpgradeClusterComponentsResponse, <-chan error) {
	responseChan := make(chan *UpgradeClusterComponentsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpgradeClusterComponents(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// invoke UpgradeClusterComponents api with *UpgradeClusterComponentsRequest asynchronously
// api document: https://help.aliyun.com/api/cs/upgradeclustercomponents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpgradeClusterComponentsWithCallback(request *UpgradeClusterComponentsRequest, callback func(response *UpgradeClusterComponentsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpgradeClusterComponentsResponse
		var err error
		defer close(result)
		response, err = client.UpgradeClusterComponents(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type UpgradeClusterComponentsRequest struct {
	*requests.RoaRequest
	ClusterId   string `position:"Path" name:"ClusterId"`
	ComponentId string `position:"Path" name:"ComponentId"`
}

type UpgradeClusterComponentsResponse struct {
	*responses.BaseResponse
}

// create a request to invoke UpgradeClusterComponents API
func CreateUpgradeClusterComponentsRequest() (request *UpgradeClusterComponentsRequest) {
	request = &UpgradeClusterComponentsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "UpgradeClusterComponents", "/clusters/[ClusterId]/components/[ComponentId]/upgrade", "", "")
	request.Method = requests.POST
	return
}

// create a response to parse from UpgradeClusterComponents response
func CreateUpgradeClusterComponentsResponse() (response *UpgradeClusterComponentsResponse) {
	response = &UpgradeClusterComponentsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}