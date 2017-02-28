// Package hammer generates ipfix packets
//: ----------------------------------------------------------------------------
//: Copyright (C) 2017 Verizon.  All Rights Reserved.
//: All Rights Reserved
//:
//: file:    ipfix_samples.go
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

var ipfixDataSamples = [][]byte{
	{0x0, 0xa, 0x1, 0xa4, 0x58, 0x6c, 0x2d, 0xe2, 0x55, 0x81, 0x3f, 0x2d, 0x0, 0x0, 0x80, 0x3, 0x1, 0x0, 0x1, 0x94, 0xc0, 0x10, 0x30, 0xd9, 0xba, 0x13, 0xd5, 0xf0, 0x2, 0x6, 0x1, 0xbb, 0xfb, 0x7d, 0x0, 0x0, 0x0, 0x0, 0x3, 0x2b, 0x0, 0x0, 0x18, 0x14, 0xfa, 0x56, 0xea, 0x0, 0x0, 0x0, 0x6c, 0x63, 0xbd, 0x7d, 0x9, 0xd5, 0x10, 0x0, 0x0, 0x4, 0x59, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x57, 0xe4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xf, 0x3f, 0x3f, 0x0, 0x0, 0x1, 0x59, 0x66, 0x92, 0x4f, 0x8a, 0x0, 0x0, 0x1, 0x59, 0x66, 0x92, 0xee, 0x30, 0x2, 0x0, 0x0, 0x0, 0x0, 0xb1, 0xcf, 0x4c, 0xf3, 0xc0, 0x10, 0x3a, 0x8, 0x0, 0x6, 0xe8, 0xb7, 0x0, 0x50, 0x0, 0x0, 0x0, 0x0, 0x4, 0x5, 0x0, 0x0, 0x15, 0x18, 0x0, 0x0, 0x49, 0xc1, 0xfa, 0x56, 0xea, 0x0, 0xc0, 0x10, 0x29, 0x36, 0x10, 0x0, 0x0, 0x4, 0x11, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x28, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x7c, 0x7c, 0x0, 0x0, 0x1, 0x59, 0x66, 0x92, 0x4c, 0x28, 0x0, 0x0, 0x1, 0x59, 0x66, 0x92, 0x4c, 0x28, 0x1, 0x0, 0x0, 0x0, 0x0, 0xc8, 0x49, 0x37, 0x22, 0xc0, 0x10, 0x30, 0xd9, 0x0, 0x6, 0x27, 0xb3, 0x1, 0xbb, 0x0, 0x0, 0x0, 0x0, 0x4, 0x59, 0x0, 0x0, 0x12, 0x18, 0x0, 0x0, 0x49, 0x3b, 0xfa, 0x56, 0xea, 0x0, 0xc0, 0x10, 0x29, 0x31, 0x10, 0x0, 0x0, 0x4, 0x11, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x34, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x35, 0x35, 0x0, 0x0, 0x1, 0x59, 0x66, 0x92, 0x41, 0x4f, 0x0, 0x0, 0x1, 0x59, 0x66, 0x92, 0x41, 0x4f, 0x1, 0x0, 0x0, 0x0, 0x0, 0xb3, 0x6c, 0x87, 0x14, 0xc0, 0x10, 0x30, 0xd9, 0x28, 0x6, 0xba, 0x83, 0x1, 0xbb, 0x0, 0x0, 0x0, 0x0, 0x4, 0x5, 0x0, 0x0, 0x18, 0x18, 0x0, 0x4, 0x3, 0x3d, 0xfa, 0x56, 0xea, 0x0, 0xc0, 0x10, 0x29, 0x31, 0x10, 0x0, 0x0, 0x4, 0x11, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x34, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x37, 0x37, 0x0, 0x0, 0x1, 0x59, 0x66, 0x92, 0x50, 0x45, 0x0, 0x0, 0x1, 0x59, 0x66, 0x92, 0x50, 0x45, 0x1, 0x0, 0x0, 0x0, 0x0, 0xbd, 0x51, 0xeb, 0xbf, 0xc0, 0x10, 0x30, 0xc8, 0x0, 0x6, 0xc7, 0x58, 0x1, 0xbb, 0x0, 0x0, 0x0, 0x0, 0x4, 0x59, 0x0, 0x0, 0x13, 0x18, 0x0, 0x0, 0x1e, 0x3a, 0xfa, 0x56, 0xea, 0x0, 0xc0, 0x10, 0x29, 0x31, 0x10, 0x0, 0x0, 0x4, 0x11, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xbc, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0x78, 0x78, 0x0, 0x0, 0x1, 0x59, 0x66, 0x92, 0x3e, 0xfb, 0x0, 0x0, 0x1, 0x59, 0x66, 0x92, 0x75, 0x8e, 0x2, 0x0, 0x0, 0x0, 0x0},
	{0x0, 0xa, 0x0, 0xb4, 0x58, 0x90, 0xd5, 0xf3, 0x25, 0x4c, 0xd3, 0x40, 0x0, 0x0, 0x84, 0x0, 0x1, 0x0, 0x0, 0xa4, 0xc0, 0xe5, 0xd3, 0x28, 0xd8, 0x6d, 0x33, 0x33, 0x0, 0x6, 0x0, 0x50, 0xde, 0xb6, 0x0, 0x0, 0x0, 0x0, 0x2, 0xe1, 0x0, 0x0, 0x18, 0x12, 0xfa, 0x56, 0xea, 0x0, 0x0, 0x0, 0xd, 0x88, 0x3e, 0x73, 0x9, 0x81, 0x10, 0x0, 0x0, 0x8, 0xf5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5, 0x8c, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x3f, 0x3f, 0x0, 0x0, 0x1, 0x59, 0xf5, 0xc2, 0xc4, 0xaf, 0x0, 0x0, 0x1, 0x59, 0xf5, 0xc2, 0xc4, 0xaf, 0x1, 0x0, 0x0, 0x0, 0x0, 0xc0, 0xe5, 0xbb, 0xa3, 0xc6, 0x7, 0x15, 0x6f, 0x0, 0x6, 0xd7, 0xa0, 0x1, 0xbb, 0x0, 0x0, 0x0, 0x0, 0x2, 0xe1, 0x0, 0x0, 0x19, 0x18, 0x0, 0x0, 0x3b, 0x1d, 0x0, 0x0, 0x3b, 0x1d, 0x3e, 0x73, 0x9, 0x81, 0x10, 0x0, 0x0, 0x8, 0xf5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x28, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x3f, 0x3f, 0x0, 0x0, 0x1, 0x59, 0xf5, 0xc2, 0xcf, 0x82, 0x0, 0x0, 0x1, 0x59, 0xf5, 0xc2, 0xcf, 0x82, 0x1, 0x0, 0x0, 0x0, 0x0},
	{0x0, 0xa, 0x1, 0xa4, 0x58, 0x90, 0xd6, 0x4d, 0x2b, 0xea, 0x94, 0xbd, 0x0, 0x0, 0x80, 0x0, 0x1, 0x0, 0x1, 0x94, 0xaa, 0x4a, 0xf8, 0x15, 0x48, 0x15, 0x51, 0xc8, 0x0, 0x6, 0xc8, 0x34, 0x1, 0xbb, 0x0, 0x0, 0x0, 0x0, 0x6, 0x5e, 0x0, 0x0, 0x16, 0x18, 0x0, 0x0, 0x13, 0xb9, 0xfa, 0x56, 0xea, 0x0, 0xc0, 0xe5, 0x9c, 0x2c, 0x10, 0x0, 0x0, 0x3, 0xc0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x28, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x3a, 0x3a, 0x0, 0x0, 0x1, 0x59, 0xf5, 0xc4, 0x2f, 0x4c, 0x0, 0x0, 0x1, 0x59, 0xf5, 0xc4, 0x2f, 0x4c, 0x1, 0x0, 0x0, 0x0, 0x0, 0x48, 0x15, 0x5b, 0x46, 0xc6, 0xf5, 0xc0, 0x88, 0x0, 0x6, 0x1, 0xbb, 0x9a, 0xdc, 0x0, 0x0, 0x0, 0x0, 0x3, 0xab, 0x0, 0x0, 0x18, 0x18, 0xfa, 0x56, 0xea, 0x0, 0x0, 0x0, 0xd, 0x1c, 0x4, 0x44, 0x47, 0xc5, 0x10, 0x0, 0x0, 0x6, 0x5e, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5, 0xdc, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x3f, 0x3f, 0x0, 0x0, 0x1, 0x59, 0xf5, 0xc4, 0x29, 0x31, 0x0, 0x0, 0x1, 0x59, 0xf5, 0xc4, 0x29, 0x31, 0x1, 0x0, 0x0, 0x0, 0x0, 0xc0, 0xe5, 0xd2, 0xb5, 0xa5, 0xe6, 0xe1, 0xc, 0x0, 0x6, 0x1, 0xbb, 0x70, 0x1d, 0x0, 0x0, 0x0, 0x0, 0x3, 0xab, 0x0, 0x0, 0x18, 0x17, 0xfa, 0x56, 0xea, 0x0, 0x0, 0x0, 0x0, 0x2e, 0x3f, 0x41, 0xba, 0x31, 0x10, 0x0, 0x0, 0x6, 0x61, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xb, 0xb8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x3f, 0x3f, 0x0, 0x0, 0x1, 0x59, 0xf5, 0xc4, 0x32, 0xf, 0x0, 0x0, 0x1, 0x59, 0xf5, 0xc4, 0x32, 0x6f, 0x1, 0x0, 0x0, 0x0, 0x0, 0xc0, 0xe5, 0xd2, 0xb5, 0x68, 0xc9, 0xf8, 0x4e, 0x0, 0x6, 0x0, 0x50, 0x86, 0x6c, 0x0, 0x0, 0x0, 0x0, 0x3, 0xab, 0x0, 0x0, 0x18, 0x12, 0xfa, 0x56, 0xea, 0x0, 0x0, 0x0, 0x74, 0x98, 0x4, 0x44, 0x47, 0xc5, 0x10, 0x0, 0x0, 0x6, 0x5e, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5, 0x92, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x3f, 0x3f, 0x0, 0x0, 0x1, 0x59, 0xf5, 0xc4, 0x28, 0xe9, 0x0, 0x0, 0x1, 0x59, 0xf5, 0xc4, 0x28, 0xe9, 0x1, 0x0, 0x0, 0x0, 0x0, 0x48, 0x15, 0x51, 0xfd, 0xa7, 0x15, 0x8e, 0x2a, 0x0, 0x6, 0x0, 0x50, 0x12, 0xce, 0x0, 0x0, 0x0, 0x0, 0x3, 0xab, 0x0, 0x0, 0x18, 0x11, 0xfa, 0x56, 0xea, 0x0, 0x0, 0x0, 0x77, 0xb1, 0x4, 0x44, 0x47, 0xc5, 0x10, 0x0, 0x0, 0x6, 0x5e, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1d, 0x4c, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5, 0x3f, 0x3f, 0x0, 0x0, 0x1, 0x59, 0xf5, 0xbd, 0xc5, 0xd6, 0x0, 0x0, 0x1, 0x59, 0xf5, 0xc4, 0xee, 0x4b, 0x2, 0x0, 0x0, 0x0, 0x0},
	{0x0, 0xa, 0x1, 0xa4, 0x58, 0x6c, 0x2d, 0xe2, 0x43, 0x7b, 0xab, 0xce, 0x0, 0x0, 0x8a, 0x3, 0x1, 0x0, 0x1, 0x94, 0xb8, 0xb3, 0x6e, 0xa2, 0xc0, 0xe5, 0xd3, 0x28, 0x0, 0x6, 0xd7, 0x2d, 0x0, 0x50, 0x0, 0x0, 0x0, 0x0, 0x7, 0x48, 0x0, 0x0, 0x16, 0x18, 0x0, 0x0, 0x58, 0xf5, 0xfa, 0x56, 0xea, 0x0, 0xc0, 0xe5, 0x82, 0x2f, 0x10, 0x0, 0x0, 0x3, 0xed, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x7b, 0x7b, 0x0, 0x0, 0x1, 0x59, 0x66, 0x92, 0x45, 0x95, 0x0, 0x0, 0x1, 0x59, 0x66, 0x92, 0x45, 0x95, 0x1, 0x0, 0x0, 0x0, 0x0, 0x44, 0x60, 0x6b, 0x3, 0xc0, 0xe5, 0xa3, 0xf8, 0x0, 0x6, 0xd9, 0xc1, 0x1, 0xbb, 0x0, 0x0, 0x0, 0x0, 0x7, 0x48, 0x0, 0x0, 0x13, 0x18, 0x0, 0x0, 0x58, 0xf5, 0x0, 0x0, 0xad, 0x3b, 0xc0, 0xe5, 0x82, 0x2e, 0x10, 0x0, 0x0, 0x3, 0xed, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x50, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x7c, 0x7c, 0x0, 0x0, 0x1, 0x59, 0x66, 0x92, 0x43, 0x86, 0x0, 0x0, 0x1, 0x59, 0x66, 0x92, 0x46, 0x18, 0x1, 0x0, 0x0, 0x0, 0x0, 0x48, 0x15, 0x5b, 0x9, 0x46, 0xb4, 0x8d, 0xfa, 0x0, 0x6, 0x0, 0x50, 0xe8, 0x49, 0x0, 0x0, 0x0, 0x0, 0x3, 0xeb, 0x0, 0x0, 0x18, 0x11, 0x0, 0x0, 0xad, 0x3b, 0x0, 0x0, 0x58, 0xf5, 0x46, 0xa7, 0x97, 0x41, 0x10, 0x0, 0x0, 0x6, 0x6e, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5, 0xdc, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x3f, 0x3f, 0x0, 0x0, 0x1, 0x59, 0x66, 0x92, 0x46, 0x7c, 0x0, 0x0, 0x1, 0x59, 0x66, 0x92, 0x46, 0x7c, 0x1, 0x0, 0x0, 0x0, 0x0, 0x48, 0xc1, 0x43, 0x37, 0xc0, 0xe5, 0xd2, 0xa3, 0x0, 0x6, 0xc4, 0x9e, 0x1, 0xbb, 0x0, 0x0, 0x0, 0x0, 0x7, 0x48, 0x0, 0x0, 0x10, 0x18, 0x0, 0x0, 0x58, 0xf5, 0x0, 0x0, 0xad, 0x3b, 0xc0, 0xe5, 0x82, 0x28, 0x10, 0x0, 0x0, 0x3, 0xed, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x28, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x3b, 0x3b, 0x0, 0x0, 0x1, 0x59, 0x66, 0x92, 0x4b, 0x66, 0x0, 0x0, 0x1, 0x59, 0x66, 0x92, 0x4b, 0x66, 0x1, 0x0, 0x0, 0x0, 0x0, 0xc0, 0xe5, 0xad, 0x5f, 0x44, 0x6c, 0x92, 0x27, 0x0, 0x6, 0x1, 0xbb, 0x99, 0xec, 0x0, 0x0, 0x0, 0x0, 0x3, 0xeb, 0x0, 0x0, 0x18, 0x12, 0x0, 0x0, 0xad, 0x3b, 0x0, 0x0, 0x58, 0xf5, 0x48, 0xd7, 0xe0, 0xaa, 0x10, 0x0, 0x0, 0x7, 0x48, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5, 0xdc, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x3f, 0x3f, 0x0, 0x0, 0x1, 0x59, 0x66, 0x92, 0x48, 0x7e, 0x0, 0x0, 0x1, 0x59, 0x66, 0x92, 0x48, 0x7e, 0x1, 0x0, 0x0, 0x0, 0x0},
}

var ipfixTemplates = [][]byte{
	{0x0, 0xa, 0x0, 0x7c, 0x58, 0x90, 0xd5, 0xf4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x8a, 0x0, 0x0, 0x2, 0x0, 0x6c, 0x1, 0x0, 0x0, 0x19, 0x0, 0x8, 0x0, 0x4, 0x0, 0xc, 0x0, 0x4, 0x0, 0x5, 0x0, 0x1, 0x0, 0x4, 0x0, 0x1, 0x0, 0x7, 0x0, 0x2, 0x0, 0xb, 0x0, 0x2, 0x0, 0x20, 0x0, 0x2, 0x0, 0xa, 0x0, 0x4, 0x0, 0x3a, 0x0, 0x2, 0x0, 0x9, 0x0, 0x1, 0x0, 0xd, 0x0, 0x1, 0x0, 0x10, 0x0, 0x4, 0x0, 0x11, 0x0, 0x4, 0x0, 0xf, 0x0, 0x4, 0x0, 0x6, 0x0, 0x1, 0x0, 0xe, 0x0, 0x4, 0x0, 0x1, 0x0, 0x8, 0x0, 0x2, 0x0, 0x8, 0x0, 0x34, 0x0, 0x1, 0x0, 0x35, 0x0, 0x1, 0x0, 0x98, 0x0, 0x8, 0x0, 0x99, 0x0, 0x8, 0x0, 0x88, 0x0, 0x1, 0x0, 0xf3, 0x0, 0x2, 0x0, 0xf5, 0x0, 0x2},
	{0x0, 0xa, 0x0, 0x7c, 0x58, 0x90, 0xd5, 0xf4, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0, 0xaa, 0x0, 0x0, 0x2, 0x0, 0x6c, 0x1, 0x1, 0x0, 0x19, 0x0, 0x1b, 0x0, 0x10, 0x0, 0x1c, 0x0, 0x10, 0x0, 0x5, 0x0, 0x1, 0x0, 0x4, 0x0, 0x1, 0x0, 0x7, 0x0, 0x2, 0x0, 0xb, 0x0, 0x2, 0x0, 0x8b, 0x0, 0x2, 0x0, 0xa, 0x0, 0x4, 0x0, 0x3a, 0x0, 0x2, 0x0, 0x1d, 0x0, 0x1, 0x0, 0x1e, 0x0, 0x1, 0x0, 0x10, 0x0, 0x4, 0x0, 0x11, 0x0, 0x4, 0x0, 0x3e, 0x0, 0x10, 0x0, 0x6, 0x0, 0x1, 0x0, 0xe, 0x0, 0x4, 0x0, 0x1, 0x0, 0x8, 0x0, 0x2, 0x0, 0x8, 0x0, 0x34, 0x0, 0x1, 0x0, 0x35, 0x0, 0x1, 0x0, 0x98, 0x0, 0x8, 0x0, 0x99, 0x0, 0x8, 0x0, 0x88, 0x0, 0x1, 0x0, 0xf3, 0x0, 0x2, 0x0, 0xf5, 0x0, 0x2},
	{0x0, 0xa, 0x0, 0x7c, 0x58, 0x6c, 0x2d, 0xe2, 0x14, 0x81, 0xaf, 0x2e, 0x0, 0x0, 0x87, 0x0, 0x0, 0x2, 0x0, 0x6c, 0x1, 0x0, 0x0, 0x19, 0x0, 0x8, 0x0, 0x4, 0x0, 0xc, 0x0, 0x4, 0x0, 0x5, 0x0, 0x1, 0x0, 0x4, 0x0, 0x1, 0x0, 0x7, 0x0, 0x2, 0x0, 0xb, 0x0, 0x2, 0x0, 0x20, 0x0, 0x2, 0x0, 0xa, 0x0, 0x4, 0x0, 0x3a, 0x0, 0x2, 0x0, 0x9, 0x0, 0x1, 0x0, 0xd, 0x0, 0x1, 0x0, 0x10, 0x0, 0x4, 0x0, 0x11, 0x0, 0x4, 0x0, 0xf, 0x0, 0x4, 0x0, 0x6, 0x0, 0x1, 0x0, 0xe, 0x0, 0x4, 0x0, 0x1, 0x0, 0x8, 0x0, 0x2, 0x0, 0x8, 0x0, 0x34, 0x0, 0x1, 0x0, 0x35, 0x0, 0x1, 0x0, 0x98, 0x0, 0x8, 0x0, 0x99, 0x0, 0x8, 0x0, 0x88, 0x0, 0x1, 0x0, 0xf3, 0x0, 0x2, 0x0, 0xf5, 0x0, 0x2},
	{0x0, 0xa, 0x0, 0x7c, 0x58, 0x6c, 0x2d, 0xe2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x87, 0x1, 0x0, 0x2, 0x0, 0x6c, 0x1, 0x0, 0x0, 0x19, 0x0, 0x8, 0x0, 0x4, 0x0, 0xc, 0x0, 0x4, 0x0, 0x5, 0x0, 0x1, 0x0, 0x4, 0x0, 0x1, 0x0, 0x7, 0x0, 0x2, 0x0, 0xb, 0x0, 0x2, 0x0, 0x20, 0x0, 0x2, 0x0, 0xa, 0x0, 0x4, 0x0, 0x3a, 0x0, 0x2, 0x0, 0x9, 0x0, 0x1, 0x0, 0xd, 0x0, 0x1, 0x0, 0x10, 0x0, 0x4, 0x0, 0x11, 0x0, 0x4, 0x0, 0xf, 0x0, 0x4, 0x0, 0x6, 0x0, 0x1, 0x0, 0xe, 0x0, 0x4, 0x0, 0x1, 0x0, 0x8, 0x0, 0x2, 0x0, 0x8, 0x0, 0x34, 0x0, 0x1, 0x0, 0x35, 0x0, 0x1, 0x0, 0x98, 0x0, 0x8, 0x0, 0x99, 0x0, 0x8, 0x0, 0x88, 0x0, 0x1, 0x0, 0xf3, 0x0, 0x2, 0x0, 0xf5, 0x0, 0x2},
}

var ipfixTemplatesOpt = [][]byte{
	{0x0, 0xa, 0x0, 0x34, 0x58, 0x90, 0xd6, 0x4e, 0x0, 0x1f, 0x4d, 0x23, 0x0, 0x0, 0x8b, 0x0, 0x0, 0x3, 0x0, 0x24, 0x2, 0x0, 0x0, 0x6, 0x0, 0x1, 0x0, 0x90, 0x0, 0x4, 0x0, 0xa0, 0x0, 0x8, 0x0, 0x82, 0x0, 0x4, 0x0, 0x83, 0x0, 0x10, 0x0, 0xd6, 0x0, 0x1, 0x0, 0xd7, 0x0, 0x1, 0x0, 0x0},
	{0x0, 0xa, 0x0, 0x34, 0x58, 0x6c, 0x2d, 0xe2, 0x0, 0xd, 0xcc, 0x64, 0x0, 0x0, 0x87, 0x0, 0x0, 0x3, 0x0, 0x24, 0x2, 0x0, 0x0, 0x6, 0x0, 0x1, 0x0, 0x90, 0x0, 0x4, 0x0, 0xa0, 0x0, 0x8, 0x0, 0x82, 0x0, 0x4, 0x0, 0x83, 0x0, 0x10, 0x0, 0xd6, 0x0, 0x1, 0x0, 0xd7, 0x0, 0x1, 0x0, 0x0},
	{0x0, 0xa, 0x0, 0x34, 0x58, 0x6c, 0x2d, 0xe2, 0x0, 0x23, 0xfb, 0x19, 0x0, 0x0, 0x82, 0x0, 0x0, 0x3, 0x0, 0x24, 0x2, 0x0, 0x0, 0x6, 0x0, 0x1, 0x0, 0x90, 0x0, 0x4, 0x0, 0xa0, 0x0, 0x8, 0x0, 0x82, 0x0, 0x4, 0x0, 0x83, 0x0, 0x10, 0x0, 0xd6, 0x0, 0x1, 0x0, 0xd7, 0x0, 0x1, 0x0, 0x0},
}
