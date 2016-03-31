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

import "sort"

type PositionsSlice []postionwithword

type postionwithword struct {
	Position
	Word *Word
}

func (p PositionsSlice) Len() int {
	return len(p)
}
func (p PositionsSlice) Less(i, j int) bool {
	return p[i].Start <= p[j].Start && p[i].End <= p[j].End
}
func (p PositionsSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// 删除重叠部分
func (p PositionsSlice) remoteOverlapZone() PositionsSlice {
	//排序
	sort.Sort(p)

	//默认已按 Start 和 End 排序，遍历删除重叠区间
	// 前面元素必然比后面元素范围更小，只需要检查当前元素是否包含在下一个元素中即可。
	newS := p[:0]
	for i := 0; i < len(p); i++ {
		if i == len(p)-1 {
			newS = append(newS, p[i])
		} else {
			p1, p2 := p[i], p[i+1]
			// p1 在 p2 范围内
			if !(p1.Start >= p2.Start && p1.End <= p2.End) {
				newS = append(newS, p1)
			}
		}
	}
	return newS
}

//默认高亮处理，将使用黄色背景高亮词汇。
func HighlightDefault(data []byte, words ...*Word) string {
	return Highlight(data,
		func(w *Word) (string, string) {
			return `<font style="background-color:#FFFF00">`, "</font>"
		},
		words...)
}

//将数据中的词汇高亮标注
// tagFn 用于处理高亮代码格式，返回对应的tag前缀和tag后缀，如: <font color="red">  </font>
func Highlight(data []byte, tagFn func(*Word) (string, string), words ...*Word) string {

	if len(words) == 0 {
		return string(data)
	}

	var positions PositionsSlice

	//收集
	for _, w := range words {
		for _, p := range w.Positions {
			positions = append(positions, postionwithword{Position: p, Word: w})
		}
	}
	//排序去重
	positions = positions.remoteOverlapZone()

	//按顺序修改
	// 记录修改增长

	var (
		newStart, newEnd, addLen int
		tagBegin, tagEnd         []byte
	)
	for _, p := range positions {

		newStart = p.Start + addLen
		newEnd = p.End + addLen

		tagBeginStr, tabEndStr := tagFn(p.Word)
		tagBegin, tagEnd = []byte(tagBeginStr), []byte(tabEndStr)

		//插入 tag
		newData := make([]byte, 0, len(data)+len(tagBegin)+len(tagEnd))

		newData = append(newData, data[:newStart]...)
		newData = append(newData, tagBegin...)
		newData = append(newData, data[newStart:newEnd]...)
		newData = append(newData, tagEnd...)
		newData = append(newData, data[newEnd:]...)

		data = newData
		addLen += len(tagBegin) + len(tagEnd)

	}

	return string(data)
}
