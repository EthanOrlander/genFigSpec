# genFigSpec

This package contains a Cobra Command "GenFigSpec", which introspects a Cobra CLI and produces a partial skeleton Fig autocomplete spec.

This package is still in development. The end goal is for it to be used in a GitHub Action that can be easily attached to any Cobra CLI.

## Usage

The method of usage will evolve quickly as this package is developed. This is how you should use it in its current state:

1. Add dependency to project using `go get github.com/EthanOrlander/genFigSpec`
2. In your Cobra CLI's source code, import this package in the file containing the root command
    - Add "github.com/EthanOrlander/genFigSpec" to imports
3. Attach the `GenFigSpec` command to the root command. This will look like `<rootCmd>.AddCommand(genFigSpec.GenFigSpec)`
4. Update go.sum using `go mod tidy`
5. Once you rebuild the CLI with this update, simply use `<cliname> genFigSpec > <cliname>.ts` in your command line.
6. This creates a TypeScript file **\<cliname\>.ts** containing the skeletal Fig autocomplete spec for your CLI. It can now be copied to [fig/autocomplete](https://github.com/withfig/autocomplete) and completed.

## Example

As an example, I have performed the above steps in [this fork of pulumi](https://github.com/EthanOrlander/pulumi/tree/genFigSpec).
You'll see that [**pkg/cmd/pulumi/pulumi.go**](https://github.com/EthanOrlander/pulumi/blob/genFigSpec/pkg/cmd/pulumi/pulumi.go#L39) now imports this package and attaches the command on [line 218](https://github.com/EthanOrlander/pulumi/blob/genFigSpec/pkg/cmd/pulumi/pulumi.go#L218) using `cmd.AddCommand(genFigSpec.GenFigSpec)`.
You'll also see a file [**fig/pulumi.ts**](https://github.com/EthanOrlander/pulumi/blob/genFigSpec/fig/pulumi.ts) that contains the generated skeleton Fig autocomplete spec for Pulumi. (I ran it through Fig's linter manually if you're wondering why it looks formatted)

## What's next

1. Complete spec generator - See [issues](https://github.com/EthanOrlander/genFigSpec/issues)
    - Decide how to handle command arguments
      - Shell completions look like a promising reference
      - It appears that in a Cobra CLI you only specify whether or not a command takes arguments, and the number or number range of arguments. The arguments are not defined with any other information (such as names)
    - It looks like when the package is imported to a Cobra CLI, it is updating some dependencies.
        - Appears this can be solved by using this package as a plugin
        - Alternatively, I can adjust this package so it's dependency versions align with those of Cobra
2. Create GitHub Action
    - Action could work as follows:
      1. Clone repo
      2. Add dependency to project using `go get github.com/EthanOrlander/genFigSpec`
      3. Add "github.com/EthanOrlander/genFigSpec" import to Cobra CLI root file
         - Root file path provided as environment variable
      4. Add `GenFigSpec` command to root command
         - Root command reference (within root file) provided as environment variable
      5. Build
         - Action will also require direction for how to build
      6. Run `cli genFigSpec`
      7. Publish fig spec *somewhere?*
         - Maybe run a diff on the previous fig spec, and make a PR to [withfig/autocomplete](https://github.com/withfig/autocomplete) if anything's changed
         - This gets more complicated since this utility only generates a *partial skeleton*, and will be modified after
