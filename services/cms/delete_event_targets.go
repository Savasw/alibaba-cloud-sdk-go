package cms

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

// DeleteEventTargets invokes the cms.DeleteEventTargets API synchronously
// api document: https://help.aliyun.com/api/cms/deleteeventtargets.html
func (client *Client) DeleteEventTargets(request *DeleteEventTargetsRequest) (response *DeleteEventTargetsResponse, err error) {
	response = CreateDeleteEventTargetsResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteEventTargetsWithChan invokes the cms.DeleteEventTargets API asynchronously
// api document: https://help.aliyun.com/api/cms/deleteeventtargets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteEventTargetsWithChan(request *DeleteEventTargetsRequest) (<-chan *DeleteEventTargetsResponse, <-chan error) {
	responseChan := make(chan *DeleteEventTargetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteEventTargets(request)
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

// DeleteEventTargetsWithCallback invokes the cms.DeleteEventTargets API asynchronously
// api document: https://help.aliyun.com/api/cms/deleteeventtargets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteEventTargetsWithCallback(request *DeleteEventTargetsRequest, callback func(response *DeleteEventTargetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteEventTargetsResponse
		var err error
		defer close(result)
		response, err = client.DeleteEventTargets(request)
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

// DeleteEventTargetsRequest is the request struct for api DeleteEventTargets
type DeleteEventTargetsRequest struct {
	*requests.RpcRequest
	Ids      *[]string `position:"Query" name:"Ids"  type:"Repeated"`
	RuleName string    `position:"Query" name:"RuleName"`
}

// DeleteEventTargetsResponse is the response struct for api DeleteEventTargets
type DeleteEventTargetsResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteEventTargetsRequest creates a request to invoke DeleteEventTargets API
func CreateDeleteEventTargetsRequest() (request *DeleteEventTargetsRequest) {
	request = &DeleteEventTargetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "DeleteEventTargets", "cms", "openAPI")
	return
}

// CreateDeleteEventTargetsResponse creates a response to parse from DeleteEventTargets response
func CreateDeleteEventTargetsResponse() (response *DeleteEventTargetsResponse) {
	response = &DeleteEventTargetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
