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

package crypto

import (
	"encoding/hex"
	"fmt"
	"testing"
)

/*
  telegram客户端使用的是pkcs1格式公钥证书，但GO语言没找到支持PKCS8格式的操作，
  参考
     http://blog.qiujinwu.com/2017/07/14/rsa/
     https://medium.com/@oyrxx/rsa%E7%A7%98%E9%92%A5%E4%BB%8B%E7%BB%8D%E5%8F%8Aopenssl%E7%94%9F%E6%88%90%E5%91%BD%E4%BB%A4-d3fcc689513f

  使用如下方法生成:
    openssl genrsa -out server.key 2048

    转成pcks8:
    openssl pkcs8 -topk8 -inform PEM -in server.key -outform pem -nocrypt -out server8.key
    openssl rsa -in server.key -pubout > public_pkcs8.pub

    转成pcks1:
    openssl rsa -in server.key -outform PEM -RSAPublicKey_out -out public_pkcs1.key
*/

var pkcs1PemPrivateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAvKLEOWTzt9Hn3/9Kdp/RdHcEhzmd8xXeLSpHIIzaXTLJDw8B
hJy1jR/iqeG8Je5yrtVabqMSkA6ltIpgylH///FojMsX1BHu4EPYOXQgB0qOi6kr
08iXZIH9/iOPQOWDsL+Lt8gDG0xBy+sPe/2ZHdzKMjX6O9B4sOsxjFrk5qDoWDri
oJorAJ7eFAfPpOBf2w73ohXudSrJE0lbQ8pCWNpMY8cB9i8r+WBitcvouLDAvmtn
TX7akhoDzmKgpJBYliAY4qA73v7u5UIepE8QgV0jCOhxJCPubP8dg+/PlLLVKyxU
5CdiQtZj2EMy4s9xlNKzX8XezE0MHEa6bQpnFwIDAQABAoIBACd+SGjfyursZoiO
MW/ejAK/PFJ3bKtNI8P++v9Enh8vF8swUBgMmzIdv93jZfnnD1mtT46kU6mXd3fy
FMunGVrjlwkLKET9MC8B5U46Es6T/H4fAA8KCzA+ywefOEnVA5pIsB7dIFFhyNDB
uO8zrBAFfsu+Y1KMlggsZaZGDXB/WVyUJDbEOMZstVx4uNhpcEgKYp28YQMP/yvv
dp4UgnTxXXXpDghzO5iqi5tUWY0p1lH2ii2OZBxEdqdDl7TirorhUDYIivyoe3B5
H30RNBRok/6w7W0WPyY2lSIcjd3cLPte6vx0QfBXVo2A6N9LTKAtAw3iWBp0x9NZ
N5p8OeECgYEA8QywXlM8nH5M7Sg2sMUYBOHA22O26ZPio7rJzcb8dlkV5gVHm+Kl
aDP61Uy8KoYABQ5kFdem/IQAUPepLxmJmiqfbwOIjfajOD3uVAQunFnDCHBWm4Uk
onbpdA5NlT/OUoSjIBemiBR/4CpDK1cEby/sg+EvQaGxqtedEe4xFmcCgYEAyFXe
MyAAOLpzmnCs9NYTTvMPofW8y+kLDodfbskl7M8q6l20VMo/E+g1gQ+65Aah901Z
/LKGi6HpzmHi5q9O2OJtqyI6FVwjXa07M5ueDbHcVKJw4hC9W0oHpMg8hqumPAWF
+MoN/Toy77p5LzoR30WUdhPvOAJPEL1p2a6r29ECgYEAiXfCEVkI5PqGZm2bmv4b
75TLhpJ8WwMSqms48Vi828V8Xpy+NOFxkVargv9rBBk9Y6TMYUSGH9Yr1AEZhBnd
RoVuPUJXmxaACPAQvetQpavvNR3T1od82AZWpvANQMONp7Oqz/+M4mhGcRHJEqti
hQJgsOk4KQbMqvChy/r6FZsCgYEAwyaqgkD9FkXC0UJLqWFUg8bQhqPcGwLUC34h
n8kAUbPpiU5omWQ+mATPAf8xvmkbo81NCJVb7W93U90U7ET/2NSRonCABkiwBtP2
ZKqGB68oA6YNspo960ytL38DPui80aFLxXQGtpPYBKEw5al6uXWNTozSrkvJe3QY
Rb4amdECgYBpGk7zPcK1TbJ++W5fkiory4qOdf0L1Zf0NbML4fY6dIww+dwMVUpq
FbsgCLqimqOFaaECU+LQEFUHHM7zrk7NBf7GzBvQ+qJx8zhJ66sFVox+IirBUyR9
Vh0+z5tIbFbKmYkO06NbeMlq87JexSlocPZtA3HMhEga5/0fHNHsNw==
-----END RSA PRIVATE KEY-----
`)

var pkcs8PemPrivateKey = []byte(`
-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC8osQ5ZPO30eff
/0p2n9F0dwSHOZ3zFd4tKkcgjNpdMskPDwGEnLWNH+Kp4bwl7nKu1VpuoxKQDqW0
imDKUf//8WiMyxfUEe7gQ9g5dCAHSo6LqSvTyJdkgf3+I49A5YOwv4u3yAMbTEHL
6w97/Zkd3MoyNfo70Hiw6zGMWuTmoOhYOuKgmisAnt4UB8+k4F/bDveiFe51KskT
SVtDykJY2kxjxwH2Lyv5YGK1y+i4sMC+a2dNftqSGgPOYqCkkFiWIBjioDve/u7l
Qh6kTxCBXSMI6HEkI+5s/x2D78+UstUrLFTkJ2JC1mPYQzLiz3GU0rNfxd7MTQwc
RrptCmcXAgMBAAECggEAJ35IaN/K6uxmiI4xb96MAr88Undsq00jw/76/0SeHy8X
yzBQGAybMh2/3eNl+ecPWa1PjqRTqZd3d/IUy6cZWuOXCQsoRP0wLwHlTjoSzpP8
fh8ADwoLMD7LB584SdUDmkiwHt0gUWHI0MG47zOsEAV+y75jUoyWCCxlpkYNcH9Z
XJQkNsQ4xmy1XHi42GlwSApinbxhAw//K+92nhSCdPFddekOCHM7mKqLm1RZjSnW
UfaKLY5kHER2p0OXtOKuiuFQNgiK/Kh7cHkffRE0FGiT/rDtbRY/JjaVIhyN3dws
+17q/HRB8FdWjYDo30tMoC0DDeJYGnTH01k3mnw54QKBgQDxDLBeUzycfkztKDaw
xRgE4cDbY7bpk+KjusnNxvx2WRXmBUeb4qVoM/rVTLwqhgAFDmQV16b8hABQ96kv
GYmaKp9vA4iN9qM4Pe5UBC6cWcMIcFabhSSidul0Dk2VP85ShKMgF6aIFH/gKkMr
VwRvL+yD4S9BobGq150R7jEWZwKBgQDIVd4zIAA4unOacKz01hNO8w+h9bzL6QsO
h19uySXszyrqXbRUyj8T6DWBD7rkBqH3TVn8soaLoenOYeLmr07Y4m2rIjoVXCNd
rTszm54NsdxUonDiEL1bSgekyDyGq6Y8BYX4yg39OjLvunkvOhHfRZR2E+84Ak8Q
vWnZrqvb0QKBgQCJd8IRWQjk+oZmbZua/hvvlMuGknxbAxKqazjxWLzbxXxenL40
4XGRVquC/2sEGT1jpMxhRIYf1ivUARmEGd1GhW49QlebFoAI8BC961Clq+81HdPW
h3zYBlam8A1Aw42ns6rP/4ziaEZxEckSq2KFAmCw6TgpBsyq8KHL+voVmwKBgQDD
JqqCQP0WRcLRQkupYVSDxtCGo9wbAtQLfiGfyQBRs+mJTmiZZD6YBM8B/zG+aRuj
zU0IlVvtb3dT3RTsRP/Y1JGicIAGSLAG0/ZkqoYHrygDpg2ymj3rTK0vfwM+6LzR
oUvFdAa2k9gEoTDlqXq5dY1OjNKuS8l7dBhFvhqZ0QKBgGkaTvM9wrVNsn75bl+S
KivLio51/QvVl/Q1swvh9jp0jDD53AxVSmoVuyAIuqKao4VpoQJT4tAQVQcczvOu
Ts0F/sbMG9D6onHzOEnrqwVWjH4iKsFTJH1WHT7Pm0hsVsqZiQ7To1t4yWrzsl7F
KWhw9m0DccyESBrn/R8c0ew3
-----END PRIVATE KEY-----
`)

var pkcs1PemPublicKey = []byte(`
-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEAvKLEOWTzt9Hn3/9Kdp/RdHcEhzmd8xXeLSpHIIzaXTLJDw8BhJy1
jR/iqeG8Je5yrtVabqMSkA6ltIpgylH///FojMsX1BHu4EPYOXQgB0qOi6kr08iX
ZIH9/iOPQOWDsL+Lt8gDG0xBy+sPe/2ZHdzKMjX6O9B4sOsxjFrk5qDoWDrioJor
AJ7eFAfPpOBf2w73ohXudSrJE0lbQ8pCWNpMY8cB9i8r+WBitcvouLDAvmtnTX7a
khoDzmKgpJBYliAY4qA73v7u5UIepE8QgV0jCOhxJCPubP8dg+/PlLLVKyxU5Cdi
QtZj2EMy4s9xlNKzX8XezE0MHEa6bQpnFwIDAQAB
-----END RSA PUBLIC KEY-----
`)

var pkcs8PemPublicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvKLEOWTzt9Hn3/9Kdp/R
dHcEhzmd8xXeLSpHIIzaXTLJDw8BhJy1jR/iqeG8Je5yrtVabqMSkA6ltIpgylH/
//FojMsX1BHu4EPYOXQgB0qOi6kr08iXZIH9/iOPQOWDsL+Lt8gDG0xBy+sPe/2Z
HdzKMjX6O9B4sOsxjFrk5qDoWDrioJorAJ7eFAfPpOBf2w73ohXudSrJE0lbQ8pC
WNpMY8cB9i8r+WBitcvouLDAvmtnTX7akhoDzmKgpJBYliAY4qA73v7u5UIepE8Q
gV0jCOhxJCPubP8dg+/PlLLVKyxU5CdiQtZj2EMy4s9xlNKzX8XezE0MHEa6bQpn
FwIDAQAB
-----END PUBLIC KEY-----
`)

// bca2c43964f3b7d1e7dfff4a769fd174770487399df315de2d2a47208cda5d32c90f0f01849cb58d1fe2a9e1bc25ee72aed55a6ea312900ea5b48a60ca51fffff1688ccb17d411eee043d8397420074a8e8ba92bd3c8976481fdfe238f40e583b0bf8bb7c8031b4c41cbeb0f7bfd991ddcca3235fa3bd078b0eb318c5ae4e6a0e8583ae2a09a2b009ede1407cfa4e05fdb0ef7a215ee752ac913495b43ca4258da4c63c701f62f2bf96062b5cbe8b8b0c0be6b674d7eda921a03ce62a0a49058962018e2a03bdefeeee5421ea44f10815d2308e8712423ee6cff1d83efcf94b2d52b2c54e4276242d663d84332e2cf7194d2b35fc5decc4d0c1c46ba6d0a6717

func TestNewRSACryptor(t *testing.T) {

	rsa := NewRSACryptor("test_pkcs1.key")

	testData := []byte("rsa 2048 key!!!!")
	fmt.Println(string(testData))

	encData := rsa.Encrypt(testData)
	fmt.Println(hex.EncodeToString(encData))

	decData := rsa.Decrypt(encData)
	fmt.Println("len = ", len(decData), ", data: ", string(decData))
}
