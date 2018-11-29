// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Benqi (wubenqi@gmail.com)

package webpage

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/mtproto"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
)

// webPage#5f07b4bc flags:# id:long url:string display_url:string hash:int type:flags.0?string site_name:flags.1?string title:flags.2?string description:flags.3?string photo:flags.4?Photo embed_url:flags.5?string embed_type:flags.5?string embed_width:flags.6?int embed_height:flags.6?int duration:flags.7?int author:flags.8?string document:flags.9?Document cached_page:flags.10?Page = WebPage;
func GetWebPagePreview(rawurl string) *mtproto.WebPage {
	u, err := url.Parse(rawurl)
	if err != nil {
		glog.Warning("getWebPagePreview error - ", err)
		return mtproto.NewTLWebPageEmpty().To_WebPage()
	}

	glog.Info(u)
	// ogParams := []string{"image", "site_name", "title", "description", "url"}
	ogContents := GetWebpageOgList(u.String(), []string{"image", "site_name", "title", "description", "url"})
	glog.Info("ogContents - ", ogContents)

	if len(ogContents) == 0 {
		return mtproto.NewTLWebPageEmpty().To_WebPage()
	} else {
		var webPage = mtproto.NewTLWebPage()

		// TODO(@benqi): save to db
		webPage.SetId(rand.Int63())
		webPage.SetUrl(rawurl)
		webPage.SetDisplayUrl(u.String()[len(u.Scheme)+3:])
		webPage.SetType("article")
		webPage.Data2.Title, _ = ogContents["title"]
		webPage.Data2.SiteName, _ = ogContents["site_name"]
		webPage.Data2.Description, _ = ogContents["description"]

		var imageBody = []byte{}

		rawImageUrl, _ := ogContents["image"]
		if rawImageUrl != "" {
			// var  *http.Response
			resp, err := http.Get(rawImageUrl)
			if err != nil {
				glog.Warning("get image body error - ", err)
			} else {
				imageBody, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					glog.Warning("read image body error - ", imageBody)
				}
			}
		} else {
			glog.Warning("image empty")
		}

		if len(imageBody) > 0 {
			// TODO(@benqi): gen photo
			// 1. upload photo
			// 2. getPhoto
			// webPage.SetPhoto(mtproto.NewTLPhotoEmpty().To_Photo())
		}

		// webPage.SetPhoto(mtproto.NewTLPhotoEmpty().To_Photo())

		// TODO(@benqi): image ==> photoSize
		return webPage.To_WebPage()
	}
}

// author @Will
func GetWebpageOgList(url string, ogParams []string) map[string]string {
	resp, err := http.Get(url)
	if err != nil {
		glog.Info("get url error - ", err)
		return nil
	}
	contentBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Info("contentBytes error - ", err)
		return nil
	}
	return GetWebpageOgListFromContent(string(contentBytes), ogParams)
}

func GetWebpageOgListFromContent(content string, ogParams []string) map[string]string {
	pattern := regexp.MustCompile(`<meta\s+property\s*=\s*"og:([0-9a-zA-Z-]+)"\s+content\s*=\s*"([^"]*?)"\s*/>`)
	allMatches := pattern.FindAllStringSubmatch(content, -1)
	allParams := make(map[string]string)
	for _, val := range allMatches {
		glog.Infof("og val: %v = %v\n", val[1], val[2])
		k := val[1]
		v := val[2]
		allParams[k] = v
	}

	return allParams
	//if ogParams == nil {
	//	return allParams
	//}
	//res := make(map[string]string)
	//for _, k := range ogParams {
	//	if v, ok := allParams[k]; ok {
	//		res[k] = v
	//	} else {
	//		res[k] = ""
	//	}
	//}
	//return res
}
