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
	"context"
	"encoding/json"
	"fmt"

	"chainguard.dev/melange/pkg/build"
	"github.com/invopop/jsonschema"
	"github.com/spf13/cobra"
)

func Schema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schema",
		Short: "Generate the melange config json schema.",
		Example: `
	  melange schema > melange.schema.json
	  `,
		Args: cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return SchemaCmd(cmd.Context())
		},
	}

	return cmd
}

func SchemaCmd(ctx context.Context) error {
	r := new(jsonschema.Reflector)
	r.AddGoComments("chainguard.dev/melange", "./")

	s := r.Reflect(&build.Configuration{})
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}
