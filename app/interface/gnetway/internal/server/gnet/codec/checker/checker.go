// Copyright © 2024 Teamgram Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/interface/gnetway/internal/server/gnet"
)

var (
	enData = []string{
		"f40e658272167595eb783273874a5aadee1aa55f2a64bd6618cb2f26738904afd716e26f31ad298ce9d6d624d30166488386cfb9f0d9a1fbcd7cb4720c8d31253f69cd0a912c5632af120601ed4948113817c41d8f26b72c",
		"f40e65827216759595e4228f3c76f43637d8003d3d00631cea739821b3e82f5b657153c00977686bee8a0b689f26400dc76caf99cc42b262d440f8ed23f9a561438be617fce97d9781ed5e95fe842d6f8983dba97c257eedfb7ad1354a1dbd002dc4e4c09d072afcf24d00eccb37166d5d0dab6fa7a1ba5c20cb2abb98b0b7070448873905fb4daac64aa2d654373ded8a2075074bab4509c385123dade9a620c51fe73d7089e3ff6a1725e8d160f10977b8e8d0910ea26d7733fa8f0c362a0b10934be78718b0f99d04ce7cc234066f92a95077a74e7c8c945010261edeac377002c1222a67dd5958a97ed9d7bdb8ee0b48e80d283f0d31423c4bfc000026dfd3c32950971cd83f",
		//"7e6fe5cb218cce94dbfd101e5c5c7d7411267c1a7e92bdabebe2c9ebc858f3f5d47e1a23d9aa56902150295ecb0cae3aa420fd9f7c6872c464f3f86f87442af7e59332c8d70fbd66b31faf41e2faed1dfbd3b3e7138170444ba2edad5f1e5c167db4017f782f9debcccca589644f7b7285c4d22a928ed96335720b3a18bd7d2b3dd0a517a26b7cff6c9eb57f38811261393bad92e8d16b9b",
		//"f76cdb13419f239d66bc514e6ca5eadbb4b1ef652ffe3f0febe271cd6cc51d8e302d1f9490c966dcd107331e61e09cf6caa0ff91caaa0c25e7ae3d868fa5f923df9092c718198b175aca966cdcb8e1150a5699267a7ba2f55bc692c51226f4210bcaf4bfb3969ef5cc516958dc9747d98ecc290d85231edc6595e2971f603d8d6e8ab110b66927b8fd8450f673a970201e1f657ce5b0ea9a",
		//"00000000000000009459a3662b8d686644000000f18e7ebef2fe9b45bce716f237f56d2810f8832c76b7f72f7feebe7b07a884b443f994dcddbfea7fdb7e36105c6f4bfad85527d62803cd6611924b5997534d416ee4880f536997865edb97c103d440b0f09f64ca2e2dfc718a348cd74477547c48a8db53cad456efc543a240cf2b11331f2880ad20b15af6b7c09ac996c2420ac22ed6f17e774001a941f6c71227a159c09fb0c3d06fe814e8e2bd41c21f54c0cdccf4651fc30eb51605d7d3effcd620c8295a82",
		//"f70b0a589e3e72f49a648443d1a28843c82b395f20fb532336c0090245d16a7f2692aca0a0654cc515baa0f457374993feb701b05b067e1e0759e2aa301b8c5bbb214a3495ab8b4d22fb93a8520de8bbdcc3363bc59b5f1032a8541a8a66dff6e56aebb08bfcc4344ace38fd3be8c16f18f333d223c69c47bea944b9f1f8dab4772d04cab36272bf858baf761870dc0dab643f971b7290493c417b0e1db9f5584e50a5c3f5b9362ee3072f1d7d7f748a26044a813f23d330cbd7b10a8aadabdc6ab2dff629d9d6a36c5331ff8ae64695b57e80ad0f94556bfcde71e019649944cc90405b3fec9f78",
		//"f70b0a589e3e72f49a648443d1a28843c82b395f20fb532336c0090245d16a7f2692aca0a0654cc515baa0f457374993feb701b05b067e1e0759e2aa301b8c5bbb214a3495ab8b4d22fb93a8520de8bbdcc3363bc59b5f10",
		//"fac0b625882f0d48b2ea5a89cfd594a0ad8663f2be27ff206b798f68c32b2e7a38d8aa0903bf405da7ba99b1b9308c1ea21d94ae643ac6efefefef0100447a43000000000000000058e2595aea8c6866f80000007897466094ca2e1ad7386edda6b4d4e677d598156537f67a43f99301277328629ba15b1eeb272fa5285d15e6b9bd3591e37f4bf81f90e3edd768583283325e49400daf4cd736b024890d2ed4c70bb3bf4232d35192ab7154cd1084c206535396ddf5bb2c34f9cccf357cc3c893b9d5734d403c2bf16ad00c9211446310ffc506d678d4609b38c7c8381f56d9a81f014a65301e583b006b53896ea51025b46a81408d40fb03a1342786444ca5cbcedb190cf03c7ba2c6d2d522f76b132b890ee0f32c76c7374e2de752f094bf065dadfdd1900ccd2a8e801c9a45102d8862e071b7bc8a66",
		//"543f53e433ae3f8e51d83ea1759a9837fa75ab7a3f0ad33114be6be0d963bdd835d4a8429f7a82303cd405d3f95bb4f2aec7363527100761a86c8ce6d7829e196faf97e82887e877d292f990b05b9d254a177af3c9809c8b57d9318ed70edd1b2bf39859e432598e4ed0e0b5ffa08ee75e525518c3faed77b82135cf18434a6c2b2ba326689f3b37de02469c0572ad935d5deb15a5288f4a05924ac74e0cf434cfc12decf27e0799270c61f106d47254e26f00ff71a21833a5a15240e569b98b2cf47d49f874fe5a",
	}

	kDataList = map[int64]string{
		4444856049120896726:  "GRwMkGqX4h4l+MWffPovraIzOfFwxaxI3SD3GVga6tjydNGvadHBtP6wDSsCO69m8lwzqX0reXPuyPO44prQzTuZVe2b+LUVk5w31c1ToMTQ1WRe7DZEZ4Uq/kIm47OaVPsmIFpuRqwyg912eryyJoqfUovjYoPWE3M+VWdpEvJ3FQbnmnyYm4NLSnmxpf6uO1ZdXtIYhyWkv/Bl7UwISqsCFyColWgu35mPhjFtBPPhaVZeL+8Ba8RrUJLh79Th+VG9IOxlnujUmX9Mg/D1EJcqhSuhQdPyBVpOR9drGXuXNVUhvf4iSS75YmpXfDto4jML6jRIebWImGhvTVcwSw",
		-7677205308713398540: "sY1pGU6lWLpnpvqj1U9H5WvqLPF2Xx3hdzNQ/QCkI6Bo1e8XAic+g5i/CmjfEH8zfCaMHTR1o2+mGcUpRd3Ny2Eqki8et58THlREb4zDEeZCoOIdI7voV3lBvKEemgYDnAiyfeJlzTsITAL3KLczPWjPyaHKgkDY8de7UZSrZuZoAgnJg/5tbstcJE0PqCkYmRbUKaZtSxC1YVwzrgprdN4TF0JnqEJGpZ4tDqEazzG0YIwczmV7Y70f4hgvTrJaQh8XFbCnTe5qLMMyG4F70SOjGuMUFlEVY4szq9lY0zmRrNDUSIGN/hys60n/NzSWFuZJ9l6va2ILKKticcSLNQ",
		4875127898191746537:  "zVyeyLkD/PxZnp8v9zIwdFMjhz1tN6zjJyUuTcehMHonmT9puHEO+INbX8hZN3Z/pFUyByTMEjJAJZKvOHomY7SAuOdNsz4o2YS5hlpErcpOFKbdxAop6idYeelXP0AblBDkae00mfgfI1A2B9LX5al5AaUikmjEm/tzQpFG1lpm/ubbxtZ8Zf8AZDIDZ+AHm9ACrl0CRc6aWbihz/E8S8iV1fSRc7j3c8xiL1z9Y8EuSXSmN9hXZaO9dpd3TfRhSVMpEhxqhy6gommdPfyIR+Yh77ubrJVJ5xKWUrIBa8VTLCdUQysyA0OWFQtTwiZAqfWSjE6yE9ZBPXb8a0ZV2Q",
		-832534131307312137:  "i1tPjkrDOD2QmqPv4oQhd7nRgMbw5pxMMKIE1eoiBHdKsxtQe0KIt5GU0FITmHeKdQOlinl+T76ONcipkOs9neTkq4Nse2y98r7iZc+bbYGzNL4ZwhwGhkjJDFbbfPTY2GcWy5d+t4HGLqVrkempElwbRe8EuQODy07l9o2vy6QcX8DZjTUNpJpjdV57JRvpOeF2GTmPZZuz4Pc6MI55+sBc8MHCw61JPK2S5tpzbd6YgCCJI6WoYX/O+c3cuA7f+nvL1ojUah6I1KgJvvEI5C0jSQvLCk8Vy+sVr0t+20N3PQPbkHQaggTZDv9SkcsKKDAwyMlWYFLOzYDHercJCg",
		-8196641258893787308: "cuRwp3mlGhQoS/4dNtfEvC1P6v8glPhfM4jbRBMYQ+JlSad6hMf0VoO+iabKFqIup65Xlha+wSC612b8PKz3+Ddzj7X0wqsCmEXNFUZwS7Y4v6ChKfx3pFYee264HuEgidvi2dpIfODH81gSDiqGws/AbwxVfuPfuhfh2demf7vO+82/ENTVkjnRryga0CGmAAmhB/furC6/0y7kvMfFuSUXxiTLjVYjlL4+/IzfkCS1Po9w+R/IVfcjLfN8ZFt9UBCgJ3RSWnKR6TcjxIpDHZt+5sIVd0vef/iAXdcBbghonq9iibRtNO6k1xXbGP23jKbvrt0/JGvZQk+lhI1y6w",
	}
)

func parseFromIncomingMessage(b []byte) (msgId int64, obj mtproto.TLObject, err error) {
	dBuf := mtproto.NewDecodeBuf(b)

	msgId = dBuf.Long()
	_ = dBuf.Int()
	obj = dBuf.Object()
	err = dBuf.GetError()

	return
}

func main() {
	for i, v := range enData {
		data, _ := hex.DecodeString(v)
		fmt.Println("data: ", i, "len:", len(data))
		id := int64(binary.LittleEndian.Uint64(data[:8]))
		fmt.Println(id)
		if id == 0 {
			msgId, obj, _ := parseFromIncomingMessage(data[8:])
			fmt.Println("msgId:", msgId)
			fmt.Println("obj:", obj)
		} else {
			k, ok := kDataList[id]
			if ok {
				kData, _ := hex.DecodeString(k)
				kInfo := mtproto.NewAuthKeyInfo(id, kData, mtproto.AuthKeyTypeTemp)
				authKey := gnet.NewAuthKeyUtil(kInfo)
				rawData, err := authKey.AesIgeDecrypt(data[8:24], data[24:])
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("rawData:", rawData)
			}
		}
	}
}
