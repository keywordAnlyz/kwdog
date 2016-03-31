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
 
fn :=func(w *Word) (string,string){
		return "<B>","</B>"
	}
	for k, v := range testCases {
		got := Highlight([]byte(k),fn, v...)
		if got != wants[k] {
			t.Fatal("高亮处理失败，不符合预期:" + got)
		}
	}

}

func TestHighlight2(t *testing.T) {

	
	bytes:=[]byte("我是一名中国共产党员人，我爱中国，我爱中国共产党，中国万岁万岁万万岁！") 
	Config.MinFre = 3
	words, err := SegmentByte(bytes)
	if err != nil {
		t.Error(err)
	}
 
	w := []*Word{} 
	for _, v := range words { 
		w = append(w, v)
	}
	fn :=func(w *Word) (string,string){
		return "<Strong>","</Strong>"
	}
	got :=  Highlight(bytes,fn, w...)
	if got != "<Strong>我</Strong>是一名<Strong>中国</Strong>共产党员人，<Strong>我</Strong>爱<Strong>中国</Strong>，<Strong>我</Strong>爱中国共产党，<Strong>中国</Strong>万岁万岁万万岁！" {
		t.Fatal("高亮处理失败，不符合预期:" + got)
	}
}

func TestHighlight3(t *testing.T) {


	testCaes :=map[string]string{
		"吃葡萄不吐葡萄皮不吃葡萄到吐葡萄皮,吃葡萄皮葡萄比葡萄皮好吃":"吃葡萄不吐<B>葡萄皮</B>不吃葡萄到吐<B>葡萄皮</B>,吃<B>葡萄皮</B>葡萄比<B>葡萄皮</B>好吃",
		"我是我你是你我就是我":"<B>我</B>是<B>我</B>你是你<B>我</B>就是<B>我</B>",
	}

	Config.MinFre = 4
	fn :=func(w *Word) (string,string){
			return "<B>","</B>"
		}

	for k,want :=range testCaes{
		words, err := SegmentText(k)
		if err != nil {
			t.Error(err)
		}
		w := []*Word{} 
		for _, v := range words { 
			w = append(w, v)
		}
		
		got :=  Highlight([]byte(k),fn, w...)
	    if got!=want{
	    	t.Fatal("高亮处理失败，不符合预期:"+got)
	    }
	}
}
