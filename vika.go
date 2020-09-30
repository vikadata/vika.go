/*
 *
 * MIT License
 *
 * Copyright (c) 2020 VIKADATA
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package vika

import (
	"reflect"
	"time"
)

const (
	FieldKeyName          string        = "name"
	FieldKeyId            string        = "id"
	OrderDesc             string        = "desc"
	OrderAsc              string        = "asc"
	DefaultHost           string        = "https://api.vika.cn"
	DefaultPageSize       int           = 100
	MaxPageSize           int           = 1000
	DefaultPageNum        int           = 1
	DefaultRequestTimeout time.Duration = 60000000000
)

type IVika interface {
	Datasheet(datasheetId string) Datasheet
}

type Vika struct {
	VikaConfig
}

type VikaConfig struct {
	/*
	  (必填) string 你的 API Token，用于鉴权
	*/
	Token string `json:"token"`
	/*
	  (选填）全局指定 field 的查询和返回的 key。默认使用列名  'name' 。指定为 'id' 时将以 fieldId 作为查询和返回方式（使用 id 可以避免列名的修改导致代码失效问题）
	*/
	FieldKey string `json:"fieldKey"`
	/*
	 （选填）请求失效时间
	*/
	RequestTimeout time.Duration `json:"requestTimeout"`
	/*
	 (选填）目标服务器地址 默认值: https://api.vika.cn
	*/
	Host string `json:"host"`
}

// 纳秒
var defaultVikaConfig = VikaConfig{RequestTimeout: DefaultRequestTimeout, Host: DefaultHost, FieldKey: FieldKeyName}

func New(config VikaConfig) *Vika {
	vika := Vika{defaultVikaConfig}
	vika.Token = config.Token
	if !reflect.ValueOf(config.Host).IsZero() {
		vika.Host = config.Host
	}
	if !reflect.ValueOf(config.FieldKey).IsZero() {
		vika.FieldKey = config.FieldKey
	}
	if !reflect.ValueOf(config.RequestTimeout).IsZero() {
		vika.RequestTimeout = config.RequestTimeout
	}
	return &vika
}

func (vika Vika) Datasheet(datasheetId string) *Datasheet {
	return newDatasheet(datasheetId, vika)
}
