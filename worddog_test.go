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

func TestFile(t *testing.T) {

	Config.MinFre=0
	words, err := SegmentFile("testdata/text1.txt")
	if err != nil {
		t.Error(err)
	}

	//我是一名中国人，我爱中国，中国万岁万岁万万岁！
	want := []struct{
			Text string 
			Fre int  
		}{
		 	{ Text:"我",Fre:2 },
		 	{ Text:"是",Fre:1 }, 
		 	{ Text:"一名",Fre:1 },
		 	{ Text:"中国",Fre:3 }, 
		 	{ Text:"爱",Fre:1 },
		 	{ Text:"人",Fre:1 },
		 	{ Text:"万岁",Fre:2 },
		 	{ Text:"万万岁",Fre:1 },
		 	{ Text:"，",Fre:2 },
		 	{ Text:"！",Fre:1 },
		}
		
	if len(words)!=len(want){
		t.Error("期望的词汇量和结果不一致")
	}

	for _, w :=range want{
		if v,ok:=words[w.Text];!ok{
			t.Errorf("期望词汇%q不存在\r\n", w.Text)
		}else if v.Frequency()!=w.Fre {
			t.Errorf("词汇%q频次期望是 %d,实际得到 %d\r\n",w.Text, w.Fre, v.Frequency() )
		}
	}
 
}

// 仅仅是测试显示信息
func TestFile2(t *testing.T) {
	Config.MinFre=2
	words, err := SegmentFile("testdata/焦裕禄.txt")
	if err != nil {
		t.Error(err)
	}
	for _,v:=range words{
		t.Logf("text=%q,pos=%q,fre=%d,positions=%v",v.Text,v.Pos,v.Frequency(),v.Positions )
	}
}
