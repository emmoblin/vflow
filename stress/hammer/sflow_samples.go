// Package hammer generates ipfix packets
//: ----------------------------------------------------------------------------
//: Copyright (C) 2017 Verizon.  All Rights Reserved.
//: All Rights Reserved
//:
//: file:    sflow_samples.go
//: details: TODO
//: author:  Mehrdad Arshad Rad
//: date:    02/01/2017
//:
//: Licensed under the Apache License, Version 2.0 (the "License");
//: you may not use this file except in compliance with the License.
//: You may obtain a copy of the License at
//:
//:     http://www.apache.org/licenses/LICENSE-2.0
//:
//: Unless required by applicable law or agreed to in writing, software
//: distributed under the License is distributed on an "AS IS" BASIS,
//: WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//: See the License for the specific language governing permissions and
//: limitations under the License.
//: ----------------------------------------------------------------------------
package hammer

var sFlowDataSamples = [][]byte{
	{0x0, 0x0, 0x0, 0x5, 0x0, 0x0, 0x0, 0x1, 0xc0, 0xe5, 0xd6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5f, 0xa9, 0xb2, 0xa3, 0x8a, 0xc7, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0xa4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x16, 0x0, 0x0, 0x7, 0xd0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x16, 0x0, 0x0, 0x2, 0x1a, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x64, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x56, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x52, 0xd4, 0x4, 0xff, 0x1, 0x18, 0x1e, 0xde, 0xad, 0x7a, 0x48, 0xcc, 0x37, 0x81, 0x0, 0x0, 0x7, 0x8, 0x0, 0x45, 0x0, 0x0, 0x40, 0xe3, 0xa7, 0x0, 0x0, 0x4, 0x1, 0x3e, 0xc3, 0xc0, 0xe5, 0xd6, 0x17, 0xc0, 0x10, 0x3d, 0x45, 0x8, 0x0, 0x63, 0x3a, 0x8f, 0x44, 0x5, 0x81, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xe9, 0x0, 0x0, 0x0, 0x10, 0x0, 0x0, 0x0, 0x7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
	{0x0, 0x0, 0x0, 0x5, 0x0, 0x0, 0x0, 0x1, 0xc0, 0xe5, 0xd6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x47, 0x7c, 0xb2, 0xa4, 0x3b, 0x1d, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0xa0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x31, 0x0, 0x0, 0x1f, 0x40, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x31, 0x0, 0x0, 0x2, 0x1a, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x60, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x52, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x4e, 0xd4, 0x4, 0xff, 0x1, 0x1f, 0x5e, 0xf4, 0xcc, 0x55, 0xde, 0x1a, 0x92, 0x8, 0x0, 0x45, 0x0, 0x0, 0x40, 0xe6, 0x47, 0x0, 0x0, 0x5, 0x1, 0x40, 0x5, 0xc0, 0xe5, 0xd8, 0x35, 0xc0, 0x10, 0x36, 0x45, 0x8, 0x0, 0xa6, 0x2, 0x4b, 0x79, 0x6, 0x84, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xe9, 0x0, 0x0, 0x0, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
	{0x0, 0x0, 0x0, 0x5, 0x0, 0x0, 0x0, 0x1, 0xc0, 0xe5, 0xd6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x47, 0x79, 0xb2, 0xa4, 0x2a, 0x83, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0xb8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x1a, 0x0, 0x0, 0x7, 0xd0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xfe, 0x0, 0x0, 0x2, 0x1a, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x78, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x6b, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x67, 0x84, 0x18, 0x88, 0xf2, 0xe9, 0x98, 0xd4, 0x4, 0xff, 0x1, 0x1d, 0x91, 0x8, 0x0, 0x45, 0x0, 0x0, 0x59, 0x12, 0x74, 0x0, 0x0, 0x3e, 0x11, 0x1a, 0x4b, 0xc0, 0xe5, 0xd8, 0x8d, 0x41, 0x37, 0x75, 0x2b, 0x6d, 0xa6, 0x0, 0x35, 0x0, 0x45, 0xcc, 0xd9, 0xd0, 0xc4, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xa, 0x63, 0x6f, 0x6c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4, 0x62, 0x6c, 0x6f, 0x62, 0x4, 0x63, 0x6f, 0x72, 0x65, 0x7, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x3, 0x6e, 0x65, 0x74, 0x0, 0x0, 0x1, 0x0, 0x1, 0x0, 0x0, 0x29, 0x10, 0x0, 0x0, 0x0, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xe9, 0x0, 0x0, 0x0, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
	{0x0, 0x0, 0x0, 0x5, 0x0, 0x0, 0x0, 0x1, 0xc0, 0xe5, 0xd6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5f, 0xa8, 0xb2, 0xa3, 0x7b, 0xf8, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0xa4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x16, 0x0, 0x0, 0x7, 0xd0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x16, 0x0, 0x0, 0x2, 0x28, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x64, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x56, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x52, 0xd4, 0x4, 0xff, 0x1, 0x18, 0x1e, 0xde, 0xad, 0x7a, 0x48, 0xcc, 0x37, 0x81, 0x0, 0x0, 0x7, 0x8, 0x0, 0x45, 0x0, 0x0, 0x40, 0x86, 0x41, 0x0, 0x0, 0xa, 0x1, 0x1c, 0x65, 0xc0, 0xe5, 0xd6, 0x17, 0x2e, 0x16, 0x49, 0x4, 0x8, 0x0, 0x70, 0x3b, 0x38, 0x40, 0x4f, 0x84, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xe9, 0x0, 0x0, 0x0, 0x10, 0x0, 0x0, 0x0, 0x7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
}
