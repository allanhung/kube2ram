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

// PutMetricAlarm invokes the cms.PutMetricAlarm API synchronously
// api document: https://help.aliyun.com/api/cms/putmetricalarm.html
func (client *Client) PutMetricAlarm(request *PutMetricAlarmRequest) (response *PutMetricAlarmResponse, err error) {
	response = CreatePutMetricAlarmResponse()
	err = client.DoAction(request, response)
	return
}

// PutMetricAlarmWithChan invokes the cms.PutMetricAlarm API asynchronously
// api document: https://help.aliyun.com/api/cms/putmetricalarm.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PutMetricAlarmWithChan(request *PutMetricAlarmRequest) (<-chan *PutMetricAlarmResponse, <-chan error) {
	responseChan := make(chan *PutMetricAlarmResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PutMetricAlarm(request)
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

// PutMetricAlarmWithCallback invokes the cms.PutMetricAlarm API asynchronously
// api document: https://help.aliyun.com/api/cms/putmetricalarm.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PutMetricAlarmWithCallback(request *PutMetricAlarmRequest, callback func(response *PutMetricAlarmResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PutMetricAlarmResponse
		var err error
		defer close(result)
		response, err = client.PutMetricAlarm(request)
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

// PutMetricAlarmRequest is the request struct for api PutMetricAlarm
type PutMetricAlarmRequest struct {
	*requests.RpcRequest
	Period              string           `position:"Query" name:"Period"`
	Webhook             string           `position:"Query" name:"Webhook"`
	ContactGroups       string           `position:"Query" name:"ContactGroups"`
	Level               requests.Integer `position:"Query" name:"Level"`
	Subject             string           `position:"Query" name:"Subject"`
	AlertName           string           `position:"Query" name:"AlertName"`
	GroupId             string           `position:"Query" name:"GroupId"`
	Description         string           `position:"Query" name:"Description"`
	Resources           string           `position:"Query" name:"Resources"`
	Threshold           string           `position:"Query" name:"Threshold"`
	EffectiveInterval   string           `position:"Query" name:"EffectiveInterval"`
	GroupName           string           `position:"Query" name:"GroupName"`
	Filter              string           `position:"Query" name:"Filter"`
	NoEffectiveInterval string           `position:"Query" name:"NoEffectiveInterval"`
	DisplayName         string           `position:"Query" name:"DisplayName"`
	Namespace           string           `position:"Query" name:"Namespace"`
	EvaluationCount     requests.Integer `position:"Query" name:"EvaluationCount"`
	SilenceTime         requests.Integer `position:"Query" name:"SilenceTime"`
	Interval            string           `position:"Query" name:"Interval"`
	MetricName          string           `position:"Query" name:"MetricName"`
	DeepDive            string           `position:"Query" name:"DeepDive"`
	ComparisonOperator  string           `position:"Query" name:"ComparisonOperator"`
	Dimensions          string           `position:"Query" name:"Dimensions"`
	Statistics          string           `position:"Query" name:"Statistics"`
}

// PutMetricAlarmResponse is the response struct for api PutMetricAlarm
type PutMetricAlarmResponse struct {
	*responses.BaseResponse
	Success    bool   `json:"Success" xml:"Success"`
	Code       string `json:"Code" xml:"Code"`
	Message    string `json:"Message" xml:"Message"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Datapoints string `json:"Datapoints" xml:"Datapoints"`
}

// CreatePutMetricAlarmRequest creates a request to invoke PutMetricAlarm API
func CreatePutMetricAlarmRequest() (request *PutMetricAlarmRequest) {
	request = &PutMetricAlarmRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "PutMetricAlarm", "cms", "openAPI")
	return
}

// CreatePutMetricAlarmResponse creates a response to parse from PutMetricAlarm response
func CreatePutMetricAlarmResponse() (response *PutMetricAlarmResponse) {
	response = &PutMetricAlarmResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
