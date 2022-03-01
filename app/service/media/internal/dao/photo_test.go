// Copyright 2022 Teamgram Authors
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
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"testing"
)

//func init2() {
//	// rand.Seed(time.Now().UnixNano())
//	mysqlConfig1 := mysql_client.MySQLConfig{
//		Name:   "immaster",
//		DSN:    "root:@/chatengine?charset=utf8",
//		Active: 5,
//		Idle:   2,
//	}
//	mysqlConfig2 := mysql_client.MySQLConfig{
//		Name:   "imslave",
//		DSN:    "root:@/chatengine?charset=utf8",
//		Active: 5,
//		Idle:   2,
//	}
//	mysql_client.InstallMysqlClientManager([]mysql_client.MySQLConfig{mysqlConfig1, mysqlConfig2})
//	// dao.InstallMysqlDAOManager(mysql_client.GetMysqlClientManager())
//}

// go test -v -run=TestUploadPhotoFile
func TestUploadPhotoFile(t *testing.T) {
	var (
	// fileId     = int64(986511829842923520)
	// accessHash = int64(2540815227215546042)
	)

	// fileData, err := file.MakeFileDataByLoad(fileId, accessHash)
	// if err != nil {
	//	fmt.Errorf("not found <%d, %d>", fileId, accessHash)
	//	return
	// }
	//
	// _ = fileData
	// UploadPhotoFile(fileData.FileId, 1, fileData.FilePath, fileData.Ext, false)
	// UploadPhotoFile(fileData.FileId, 2, fileData.FilePath, fileData.Ext, true)

	//func MakeFileDataByLoad(fileId, accessHash int64) (*fileData, error) {
	//
	//}
}
