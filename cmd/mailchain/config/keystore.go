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
	"fmt"

	"github.com/mailchain/mailchain/cmd/mailchain/config/defaults"
	"github.com/mailchain/mailchain/cmd/mailchain/config/names"
	"github.com/mailchain/mailchain/internal/mailchain/commands/prompts"
	"github.com/mailchain/mailchain/internal/pkg/keystore/nacl"
	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper" // nolint: depguard
	"github.com/ttacon/chalk"
)

// GetKeystore create new keystore from config
func GetKeystore() (*nacl.FileStore, error) {
	if viper.GetString("storage.keys") == names.KeystoreNACLFilestore {
		fs := nacl.NewFileStore(viper.GetString(fmt.Sprintf("stores.%s.path", names.KeystoreNACLFilestore)))
		return &fs, nil
	}

	return nil, errors.Errorf("unknown keystore type")
}

func SetKeystore(cmd *cobra.Command, keystoreType string) error {
	viper.Set("storage.keys", keystoreType)
	switch keystoreType {
	case names.KeystoreNACLFilestore:
		// NACL only needs to set the path
		return setKeystorePath(cmd, keystoreType)
	default:
		return errors.Errorf("unsupported key store type")
	}
}

func setKeystorePath(cmd *cobra.Command, keystoreType string) error {
	if keystorePath, _ := cmd.Flags().GetString("keystore-path"); keystorePath != "" {
		viper.Set(fmt.Sprintf("stores.%s.path", keystoreType), keystorePath)
	}
	keystorePath, err := prompts.RequiredInputWithDefault("path", defaults.KeystorePath)
	if err != nil {
		return err
	}
	viper.Set(fmt.Sprintf("stores.%s.path", keystoreType), keystorePath)
	return nil
}

// Passphrase is extracted from the command
func Passphrase(cmd *cobra.Command) (string, error) {
	passphrase, err := cmd.Flags().GetString("passphrase")
	if err != nil {
		return "", errors.WithMessage(err, "could not get `passphrase`")
	}
	if passphrase != "" {
		return passphrase, nil
	}
	emptyPassphrase, err := cmd.Flags().GetBool("empty-passphrase")
	if err != nil {
		return "", errors.WithMessage(err, "could not get `empty-passphrase`")
	}
	if emptyPassphrase {
		return "", nil
	}
	fmt.Println(chalk.Yellow, "Note: To derive a storage key passphrase is required. The passphrase must be secure and not guessable.")
	return passphraseFromPrompt()
}

func passphraseFromPrompt() (string, error) {
	prompt := promptui.Prompt{
		Label: "Passphrase",
		Mask:  '*',
	}
	password, err := prompt.Run()
	if err != nil {
		return "", errors.Errorf("failed read passphrase")
	}

	confirmPrompt := promptui.Prompt{
		Label: "Repeat passphrase: ",
		Mask:  '*',
	}
	confirm, err := confirmPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", errors.Errorf("failed read passphrase confirmation")
	}
	if password != confirm {
		return "", errors.Errorf("Passphrases do not match")
	}

	return password, nil
}
