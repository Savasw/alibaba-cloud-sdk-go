package cloudphoto

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

type Event struct {
	WatermarkPhotoId string `json:"WatermarkPhotoId" xml:"WatermarkPhotoId"`
	WeixinTitle      string `json:"WeixinTitle" xml:"WeixinTitle"`
	SplashPhotoId    string `json:"SplashPhotoId" xml:"SplashPhotoId"`
	EndAt            int    `json:"EndAt" xml:"EndAt"`
	State            string `json:"State" xml:"State"`
	LibraryId        string `json:"LibraryId" xml:"LibraryId"`
	ViewsCount       int    `json:"ViewsCount" xml:"ViewsCount"`
	Title            string `json:"Title" xml:"Title"`
	Mtime            int    `json:"Mtime" xml:"Mtime"`
	StartAt          int    `json:"StartAt" xml:"StartAt"`
	Identity         string `json:"Identity" xml:"Identity"`
	Ctime            int    `json:"Ctime" xml:"Ctime"`
	IdStr            string `json:"IdStr" xml:"IdStr"`
	Id               int    `json:"Id" xml:"Id"`
	BannerPhotoId    string `json:"BannerPhotoId" xml:"BannerPhotoId"`
}