// Copyright 2019 Finobo
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

package config

import (
	"github.com/mailchain/mailchain/internal/pkg/stores"
	"github.com/mailchain/mailchain/internal/pkg/stores/leveldb"
	"github.com/pkg/errors"
	"github.com/spf13/viper" // nolint: depguard
)

// GetStateStore create all the clients based on configuration
func GetStateStore(vpr *viper.Viper) (stores.State, error) {
	if vpr.GetString("storage.state") == "leveldb" {
		return leveldb.New(viper.GetString("stores.leveldb.path"), 256, 0)
	}
	return nil, errors.New("unsupported mailbox state type")
}
