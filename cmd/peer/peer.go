// Copyright 2020 Authors of Hubble
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

package peer

import (
	"github.com/cilium/hubble/cmd/common/config"
	"github.com/cilium/hubble/cmd/common/template"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// New creates a new hidden peer command.
func New(vp *viper.Viper) *cobra.Command {
	peerCmd := &cobra.Command{
		Use:     "peers",
		Aliases: []string{"peer"},
		Short:   "Get information about Hubble peers",
		Long:    `Get information about Hubble peers.`,
		Hidden:  true, // this command is only useful for development/debugging purposes
	}

	// add config.ServerFlags to the help template as these flags are used by
	// this command
	peerCmd.SetUsageTemplate(template.Usage(config.ServerFlags))

	peerCmd.AddCommand(
		newWatchCommand(vp),
	)
	return peerCmd
}
