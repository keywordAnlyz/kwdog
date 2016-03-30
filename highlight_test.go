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

import (
	"image/color"
	"io/ioutil"
	"testing"
)

func TestHighlight(t *testing.T) {

	testCases := map[string][]*Word{
		"我是一名中国人，我爱中国，中国万岁万岁万万岁！": []*Word{
			{Text: "我", Positions: []Position{{0, 3}, {24, 27}}},
			{Text: "中国", Positions: []Position{{12, 18}, {30, 36}, {39, 45}}},
			{Text: "中国人", Positions: []Position{{12, 21}}},
		},
	}
	wants := map[string]string{
		"我是一名中国人，我爱中国，中国万岁万岁万万岁！": "<B>我</B>是一名<B>中国人</B>，<B>我</B>爱<B>中国</B>，<B>中国</B>万岁万岁万万岁！",
	}

	h := Highlight{Color: color.Black}

	for k, v := range testCases {
		got := h.Highlight([]byte(k), v...)
		if got != wants[k] {
			t.Fatal("高亮处理失败，不符合预期:" + got)
		}
	}

}

func TestHeightlith2(t *testing.T) {

	file := "testdata/text1.txt"

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		t.Error(err)
	}
	Config.MinFre = 3
	words, err := SegmentByte(bytes)
	if err != nil {
		t.Error(err)
	}

	h := Highlight{Color: color.Black}

	w := []*Word{}

	for _, v := range words {
		t.Log(v)
		w = append(w, v)
	}
	got := h.Highlight(bytes, w...)
	if got != "<B>我</B>是一名<B>中国</B>共产党员人，<B>我</B>爱<B>中国</B>，<B>我</B>爱中国共产党，<B>中国</B>万岁万岁万万岁！" {
		t.Fatal("高亮处理失败，不符合预期:" + got)
	}
}
