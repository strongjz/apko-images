// Copyright 2022 Chainguard, Inc.
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

package cli

import (
	cranecmd "github.com/google/go-containerregistry/cmd/crane/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"log"
	"sigs.k8s.io/release-utils/version"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "apko",
		DisableAutoGenTag: true,
		SilenceUsage:      true,
	}

	cmd.AddCommand(cranecmd.NewCmdAuthLogin("apko")) // apko login
	cmd.AddCommand(buildCmd())
	cmd.AddCommand(buildMinirootFS())
	cmd.AddCommand(showConfig())
	cmd.AddCommand(publish())
	cmd.AddCommand(version.Version())
	err := doc.GenMarkdownTree(cmd, "./docs/")
	if err != nil {
		log.Fatal(err)
	}
	return cmd
}
