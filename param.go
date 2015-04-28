// Copyright 2015 The Tango Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tango

type (
	param struct {
		Name  string
		Value string
	}
	Params []param
)

func (p *Params) Get(key string) string {
	for _, v := range *p {
		if v.Name == key {
			return v.Value
		}
	}
	return ""
}

func (p *Params) Set(key, value string) {
	for i, v := range *p {
		if v.Name == key {
			(*p)[i].Value = value
			return
		}
	}

	*p = append(*p, param{key, value})
}

type Paramer interface {
	SetParams([]param)
}

func (p *Params) SetParams(params []param) {
	*p = params
}

func Param() HandlerFunc {
	return func(ctx *Context) {
		if action := ctx.Action(); action != nil {
			if p, ok := action.(Paramer); ok {
				p.SetParams(*ctx.Params())
			}
		}
		ctx.Next()
	}
}
