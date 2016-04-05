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
	"errors"
	"io/ioutil"
	"path"
	"strings" 

	"github.com/keywordAnlyz/sego"
)

var (
	Segmenter = sego.Segmenter{}
)

func init() {
	initConfig()

	Segmenter.LoadDictionary(strings.Join(Config.DictionaryFiles, ","))
}

//词汇在文本中的位置
type Position struct {
	//开始位置
	Start int
	//结束位置，不包含该位置
	End int
}

// 词汇信息
type Word struct {
	//词汇文本信息
	Text string
	//词汇属性
	Pos       string
	//在字典中登记的频次
	DictFrequency int 
	Positions []Position
}

//词汇频次
func (w *Word) Frequency() int {
	return len(w.Positions)
}

// 解析本地文件(文件格式必须是UTF-B格式)
func SegmentFile(filename string) (map[string]*Word, error) {
	if path.Ext(filename) != ".txt" {
		return nil, errors.New("仅支持 .txt 纯文本文件解析")
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return SegmentByte(bytes)
}

// 解析文本
func SegmentText(text string) (map[string]*Word, error) {
	return SegmentByte([]byte(text))
}

// 解析 Bytes 数据，Bytes数据必须是UTF-B格式
func SegmentByte(bytes []byte) (map[string]*Word, error) {
	if len(bytes) == 0 {
		return map[string]*Word{}, nil
	}

	segments := Segmenter.Segment(bytes)

	return machin(segments), nil
}

//解析词汇
func machin(segments []sego.Segment) map[string]*Word {
	words := make(map[string]*Word, len(segments)/2)

	for _, s := range segments {
		text := s.Token().Text()
		if item, ok := words[text]; ok {
			item.Positions = append(item.Positions,
				Position{
					Start: s.Start(),
					End:   s.End(),
				})
		} else {
			words[text] = &Word{
				Text: text,
				Pos:  s.Token().Pos(),
				DictFrequency:s.Token().Frequency(),
				Positions: []Position{
					{
						Start: s.Start(),
						End:   s.End(),
					},
				},
			}
		}
	}

	//移除低于频次下限的词汇
	if Config.MinFre > 0 {
		for k, v := range words {
			if check(v) == false {
				delete(words, k)
			}
		}
	}
	return words
}

// 检查词汇是否符合配置要求
// 检查项：
//   1.符合最低频次要求
//   2.不属于黑名单词汇
func check(w *Word) bool {

	//删除标点符号
	if w.Pos == "x" {
		return false
	}
	if w.Frequency() < Config.MinFre {
		return false
	}
	if Config.BlackWords[w.Text] {
		return false
	}
	return true
}
