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

package commands

import (
	"github.com/mailchain/mailchain/cmd/mailchain/config"
	"github.com/mailchain/mailchain/cmd/mailchain/config/names"
	"github.com/mailchain/mailchain/internal/mailchain/commands/setup"
	"github.com/mailchain/mailchain/internal/pkg/encoding"
	"github.com/spf13/cobra"
)

func cfgChainEthereum() *cobra.Command {
	return &cobra.Command{
		Use:      "ethereum",
		Short:    "setup ethereum",
		PostRunE: config.WriteConfig,
		RunE: func(cmd *cobra.Command, args []string) error {
			chain := encoding.Ethereum
			network, err := setup.Network(cmd, args, chain, names.Empty)
			if err != nil {
				return err
			}

			cmd.Printf("%s configured\n", network)
			return nil
		},
	}
}
