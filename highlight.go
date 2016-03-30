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
	"bytes"
	"image/color"
	"log"
	"sort"
)

type PositionsSlice []Position

func (p PositionsSlice) Len() int {
	return len(p)
}
func (p PositionsSlice) Less(i, j int) bool {
	return p[i].Start < p[j].Start
}
func (p PositionsSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type Highlight struct {
	Color color.Color
}

//将数据中的词汇高亮标注
func (h *Highlight) Highlight(data []byte, words ...*Word) string {

	if len(words) == 0 {
		return string(data)
	}

	// html := bytes.NewBufferString("")

	var positions PositionsSlice

	//收集
	for _, w := range words {
		positions = append(positions, w.Positions...)
	}
	//排序
	sort.Sort(positions)

	log.Print(positions)

	//按顺序修改
	// 记录修改增长
	var (
		tagBegin = []byte("<B>")
		tagEnd   = []byte("</B>")
		addLen   = 0
		signLen  = len(tagBegin) + len(tagEnd)
	)

	var preEnd, newStart, newEnd int
	for _, p := range positions {

		newStart = p.Start + addLen
		newEnd = p.End + addLen

		if preEnd > newEnd {
			newEnd -= len(tagEnd)
			newStart -= len(tagEnd)
		}

		log.Println(string(data), newStart, newEnd, p)
		//插入 tag
		newData := make([]byte, 0, len(data)+signLen)

		if newStart > len(tagBegin) {

			log.Println(bytes.NewBuffer(data[newStart-len(tagBegin):newStart]).String(), "===", len(tagBegin))
		}

		if newStart > len(tagBegin) && string(data[newStart-len(tagBegin):newStart]) == string(tagBegin) {

			//newStart = newStart - len(tagBegin) - 1
		}

		newData = append(newData, data[:newStart]...)
		newData = append(newData, tagBegin...)
		newData = append(newData, data[newStart:newEnd]...)
		newData = append(newData, tagEnd...)
		newData = append(newData, data[newEnd:]...)

		data = newData

		addLen += signLen
		preEnd = newEnd
	}

	return string(data)
}
