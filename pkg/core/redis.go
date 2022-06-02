// Copyright 2020 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import (
	"encoding/json"

	"github.com/pingcap/errors"
)

const (
	RedisSentinelRestartAction  = "restart"
	RedisSentinelStopAction     = "stop"
	RedisCachePenetrationAction = "penetration"
)

var _ AttackConfig = &RedisCommand{}

type RedisCommand struct {
	CommonAttackConfig

	Addr        string `json:"addr,omitempty"`
	Password    string `json:"password,omitempty"`
	Conf        string `json:"conf,omitempty"`
	FlushConfig bool   `json:"flushConfig,omitempty"`
	RedisPath   string `json:"redisPath,omitempty"`
	RequestNum  int    `json:"requestNum,omitempty"`
}

func (r *RedisCommand) Validate() error {
	if err := r.CommonAttackConfig.Validate(); err != nil {
		return err
	}
	if len(r.Addr) == 0 {
		return errors.New("addr of redis server is required")
	}
	switch r.Action {
	case RedisCachePenetrationAction:
		if r.RequestNum == 0 {
			return errors.New("request-num is required")
		}
	}
	return nil
}

func (r RedisCommand) RecoverData() string {
	data, _ := json.Marshal(r)

	return string(data)
}

func NewRedisCommand() *RedisCommand {
	return &RedisCommand{
		CommonAttackConfig: CommonAttackConfig{
			Kind: RedisAttack,
		},
	}
}