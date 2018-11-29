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

package media

/*
	inputMediaEmpty#9664f57f = InputMedia;
	inputMediaUploadedPhoto#2f37e231 flags:# file:InputFile caption:string stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
	inputMediaPhoto#81fa373a flags:# id:InputPhoto caption:string ttl_seconds:flags.0?int = InputMedia;
	inputMediaGeoPoint#f9c44144 geo_point:InputGeoPoint = InputMedia;
	inputMediaContact#a6e45987 phone_number:string first_name:string last_name:string = InputMedia;
	inputMediaUploadedDocument#e39621fd flags:# file:InputFile thumb:flags.2?InputFile mime_type:string attributes:Vector<DocumentAttribute> caption:string stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
	inputMediaDocument#5acb668e flags:# id:InputDocument caption:string ttl_seconds:flags.0?int = InputMedia;
	inputMediaVenue#2827a81a geo_point:InputGeoPoint title:string address:string provider:string venue_id:string = InputMedia;
	inputMediaGifExternal#4843b0fd url:string q:string = InputMedia;
	inputMediaPhotoExternal#922aec1 flags:# url:string caption:string ttl_seconds:flags.0?int = InputMedia;
	inputMediaDocumentExternal#b6f74335 flags:# url:string caption:string ttl_seconds:flags.0?int = InputMedia;
	inputMediaGame#d33f43f3 id:InputGame = InputMedia;
	inputMediaInvoice#92153685 flags:# title:string description:string photo:flags.0?InputWebDocument invoice:Invoice payload:bytes provider:string start_param:string = InputMedia;


	messageMediaEmpty#3ded6320 = MessageMedia;
	messageMediaPhoto#b5223b0f flags:# photo:flags.0?Photo caption:flags.1?string ttl_seconds:flags.2?int = MessageMedia;
	messageMediaGeo#56e0d474 geo:GeoPoint = MessageMedia;
	messageMediaContact#5e7d2f39 phone_number:string first_name:string last_name:string user_id:int = MessageMedia;
	messageMediaUnsupported#9f84f49e = MessageMedia;
	messageMediaDocument#7c4414d3 flags:# document:flags.0?Document caption:flags.1?string ttl_seconds:flags.2?int = MessageMedia;
	messageMediaWebPage#a32dd600 webpage:WebPage = MessageMedia;
	messageMediaVenue#7912b71f geo:GeoPoint title:string address:string provider:string venue_id:string = MessageMedia;
	messageMediaGame#fdb19008 game:Game = MessageMedia;
	messageMediaInvoice#84551347 flags:# shipping_address_requested:flags.1?true test:flags.3?true title:string description:string photo:flags.0?WebDocument receipt_msg_id:flags.2?int currency:string total_amount:long start_param:string = MessageMedia;
*/
