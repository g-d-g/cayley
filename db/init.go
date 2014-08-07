// Copyright 2014 The Cayley Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

import (
	"errors"
	"fmt"

	"github.com/google/cayley/config"
	"github.com/google/cayley/graph"
)

var ErrNotPersistent = errors.New("database type is not persistent")

func Init(cfg *config.Config, triplePath string) error {
	if !graph.IsPersistent(cfg.DatabaseType) {
		return fmt.Errorf("ignoring unproductive database initialization request: %v", ErrNotPersistent)
	}

	err := graph.InitTripleStore(cfg.DatabaseType, cfg.DatabasePath, cfg.DatabaseOptions)
	if err != nil {
		return err
	}
	if triplePath != "" {
		ts, err := Open(cfg)
		if err != nil {
			return err
		}
		err = Load(ts, cfg, triplePath)
		if err != nil {
			return err
		}
		ts.Close()
	}
	return err
}
