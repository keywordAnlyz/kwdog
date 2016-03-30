// Copyright 2016 Author ysqi. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// @Author: ysqi
// @Email: devysq@gmail.com or 460857340@qq.com

package worddog

import "testing"

func TestHighlight(t *testing.T) {

	testCases := map[string][]*Word{
		// "我是一名中国人，我爱中国，中国万岁万岁万万岁！": []*Word{
		// 	{Text: "我", Positions: []Position{{0, 3}, {24, 27}}},
		// 	{Text: "中国", Positions: []Position{{12, 18}, {30, 36}, {39, 45}}},
		// 	{Text: "中国人", Positions: []Position{{12, 20}}},
		// },
		"abcdabcefgb": []*Word{
			{Text: "bc", Positions: []Position{{1, 3}, {5, 7}}},
			{Text: "b", Positions: []Position{{1, 2}, {5, 6}, {10, 11}}},
		},
	}

	h := Highlight{}

	for k, v := range testCases {
		got := h.Highlight([]byte(k), v...)
		t.Log(got)
	}

}
