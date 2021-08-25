/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package genFigSpec

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var includeHidden bool

var GenFigSpec = &cobra.Command{
	Use:    "genFigSpec",
	Short:  "Generate a fig spec",
	Hidden: true,
	Long: `
Fig is a tool for your command line that adds autocomplete.
This command generates a TypeScript file with the skeleton
Fig autocomplete spec for your Cobra CLI.
`,
	Run: func(cmd *cobra.Command, args []string) {
		root := cmd.Root()
		spec := makeFigSpec(root)
		fmt.Println(spec.toTypescript())
	},
}

func makeFigSpec(root *cobra.Command) Spec {
	spec := Spec{
		Subcommand: &Subcommand{
			BaseSuggestion: &BaseSuggestion{
				description: root.Short,
			},
			options:     options(root),
			subcommands: subcommands(root),
			args:        commandArguments(root),
		},
		name: root.Name(),
	}
	return spec
}

func subcommands(cmd *cobra.Command) []Subcommand {
	var subs []Subcommand
	for _, sub := range cmd.Commands() {
		if !includeHidden && (!sub.IsAvailableCommand() || sub.IsAdditionalHelpTopicCommand() || sub.Hidden) {
			continue
		}
		subs = append(subs, Subcommand{
			BaseSuggestion: &BaseSuggestion{
				description: sub.Short,
			},
			name:        append(sub.Aliases, sub.Name()),
			options:     options(sub),
			subcommands: subcommands(sub),
			args:        commandArguments(sub),
		})
	}
	return subs
}

func options(cmd *cobra.Command) []Option {
	var opts []Option
	attachFlags := func(flag *pflag.Flag) {
		option := Option{
			BaseSuggestion: &BaseSuggestion{
				displayName: flag.Name,
				description: flag.Usage,
			},
			name:         []string{fmt.Sprintf("--%v", flag.Name)},
			isRepeatable: strings.Contains(strings.ToLower(flag.Value.Type()), "array"),
		}
		if flag.Shorthand != "" {
			option.name = append(option.name, fmt.Sprintf("-%v", flag.Shorthand))
		}
		option.args = flagArguments(flag)
		opts = append(opts, option)
	}
	/* Help command doesn't appear to be included in any of the FlagSets
	 * Find another way to include the help command if it's available in the CLI
	 */
	cmd.NonInheritedFlags().VisitAll(attachFlags)
	return opts
}

/*
 * In Cobra, you only specify the number of arguments.
 * Not sure how we want to handle this (if at all)
 * https://github.com/spf13/cobra/blob/v1.2.1/user_guide.md#positional-and-custom-arguments
 */
func commandArguments(cmd *cobra.Command) []Arg {
	var args []Arg
	// switch t := cmd.Args.(T) {
	// case cobra.ArbitraryArgs:

	// }
	// fmt.Println(cmd.Args)
	return args
}

func flagArguments(flag *pflag.Flag) []Arg {
	var args []Arg
	if flag.Value.Type() != "bool" {
		arg := Arg{
			name:       flag.Name,
			defaultVal: flag.DefValue,
		}
		args = append(args, arg)
	}
	return args
}

func init() {
	GenFigSpec.Flags().BoolVar(
		&includeHidden, "include-hidden", false,
		"Include hidden commands in generated Fig autocomplete spec")
}
