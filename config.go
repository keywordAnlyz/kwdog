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
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/utils"
)

var (
	//Config配置信息
	Config *WordConfig
	//配置管理器
	Configer *configer
	//运行包位置
	APPPath string
	//配置文件存储路径
	configPath string
)

// KW 配置信息
type WordConfig struct {
	//运行环境，默认为 dev
	RunMode string
	//词汇字典文件路径
	DictionaryFiles []string
	//词汇黑名单
	BlackWords map[string]bool
	//词汇最小频次，如果为0则解析所有
	MinFre int
}

func initConfig() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	APPPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))

	defaultDict := "data/dictionary.txt"
	Config = &WordConfig{
		RunMode:         "dev",
		DictionaryFiles: []string{defaultDict},
		MinFre:          2,
		BlackWords:      map[string]bool{},
	}

	configPath = filepath.Join(APPPath, "conf", "kwdog.conf")
	if !utils.FileExists(configPath) {
		Configer = &configer{innerConfig: config.NewFakeConfig()}
		return
	}

	if err := parseConfig(configPath); err != nil {
		panic(err)
	}

}

// 配置解析，配置文件仅支持ini格式
func parseConfig(configPath string) (err error) {
	Configer, err = newAppConfig("ini", configPath)
	if err != nil {
		return err
	}

	if runMode := Configer.String("RunMode"); runMode != "" {
		Config.RunMode = runMode
	}

	if files := Configer.DefaultStrings("DictionaryFiles", Config.DictionaryFiles); len(files) > 0 {
		Config.DictionaryFiles = []string{}
		for _, f := range files {
			if !utils.FileExists(f) {
				log.Fatalf("worddog:字典文件%q不存在\r\n", f)
			} else {
				Config.DictionaryFiles = append(Config.DictionaryFiles, f)
			}
		}
	}
	if len(Config.DictionaryFiles) == 0 {
		return errors.New("没有配置字典")
	}

	Config.MinFre = Configer.DefaultInt("minfre", Config.MinFre)

	//初始化黑名单
	if blackWords := Configer.Strings("BlackWords"); len(blackWords) > 0 {
		for _, v := range blackWords {
			if Config.BlackWords[v] == false {
				Config.BlackWords[v] = true
			}
		}
	}

	return nil

}

//配置文件解析器
type configer struct {
	innerConfig config.Configer
}

func newAppConfig(appConfigProvider, appConfigPath string) (*configer, error) {
	ac, err := config.NewConfig(appConfigProvider, appConfigPath)
	if err != nil {
		return nil, err
	}
	return &configer{ac}, nil
}

func (b *configer) Set(key, val string) error {
	if err := b.innerConfig.Set(key, val); err != nil {
		return err
	}
	return b.innerConfig.Set(key, val)
}

func (b *configer) String(key string) string {
	if v := b.innerConfig.String(key); v != "" {
		return v
	}
	return b.innerConfig.String(key)
}

func (b *configer) Strings(key string) []string {
	if v := b.innerConfig.Strings(key); v[0] != "" {
		return v
	}
	return b.innerConfig.Strings(key)
}

func (b *configer) Int(key string) (int, error) {
	if v, err := b.innerConfig.Int(key); err == nil {
		return v, nil
	}
	return b.innerConfig.Int(key)
}

func (b *configer) Int64(key string) (int64, error) {
	if v, err := b.innerConfig.Int64(key); err == nil {
		return v, nil
	}
	return b.innerConfig.Int64(key)
}

func (b *configer) Bool(key string) (bool, error) {
	if v, err := b.innerConfig.Bool(key); err == nil {
		return v, nil
	}
	return b.innerConfig.Bool(key)
}

func (b *configer) Float(key string) (float64, error) {
	if v, err := b.innerConfig.Float(key); err == nil {
		return v, nil
	}
	return b.innerConfig.Float(key)
}

func (b *configer) DefaultString(key string, defaultVal string) string {
	if v := b.String(key); v != "" {
		return v
	}
	return defaultVal
}

func (b *configer) DefaultStrings(key string, defaultVal []string) []string {
	if v := b.Strings(key); len(v) != 0 {
		return v
	}
	return defaultVal
}

func (b *configer) DefaultInt(key string, defaultVal int) int {
	if v, err := b.Int(key); err == nil {
		return v
	}
	return defaultVal
}

func (b *configer) DefaultInt64(key string, defaultVal int64) int64 {
	if v, err := b.Int64(key); err == nil {
		return v
	}
	return defaultVal
}

func (b *configer) DefaultBool(key string, defaultVal bool) bool {
	if v, err := b.Bool(key); err == nil {
		return v
	}
	return defaultVal
}

func (b *configer) DefaultFloat(key string, defaultVal float64) float64 {
	if v, err := b.Float(key); err == nil {
		return v
	}
	return defaultVal
}

func (b *configer) DIY(key string) (interface{}, error) {
	return b.innerConfig.DIY(key)
}

func (b *configer) GetSection(section string) (map[string]string, error) {
	return b.innerConfig.GetSection(section)
}

func (b *configer) SaveConfigFile(filename string) error {
	return b.innerConfig.SaveConfigFile(filename)
}
