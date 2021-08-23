# genFigSpec

This package contains a Cobra Command "GenFigSpec", which introspects a Cobra CLI and produces a partial skeleton Fig autocomplete spec.

This package is still in development. The end goal is for it to be used in a GitHub Action that can be easily attached to any Cobra CLI.

## Usage

The method of usage will evolve quickly as this package is developed. This is how you should use it in its current state:

1. In your Cobra CLI's source code, import this package in the root file of the CLI
    - Add "github.com/EthanOrlander/genFigSpec" to imports
2. Attach the `GenFigSpec` command to the root command. This will look like `<rootCmd>.AddCommand(genFigSpec.GenFigSpec)`
3. Once you rebuild the CLI with this update, simply use `<cliname> genFigSpec > <cliname>.ts` in your command line.
4. This creates a TypeScript file **\<cliname\>.ts** containing the skeletal Fig autocomplete spec for your CLI. It can now be copied to [fig/autocomplete](https://github.com/withfig/autocomplete) and completed.

## Example

As an example, I have performed the above steps in [this fork of pulumi](https://github.com/EthanOrlander/pulumi/tree/genFigSpec).
You'll see that [**pkg/cmd/pulumi/pulumi.go**](https://github.com/EthanOrlander/pulumi/blob/genFigSpec/pkg/cmd/pulumi/pulumi.go#L39) now imports this package and attaches the command on [line 218](https://github.com/EthanOrlander/pulumi/blob/genFigSpec/pkg/cmd/pulumi/pulumi.go#L218) using `cmd.AddCommand(genFigSpec.GenFigSpec)`.
You'll also see a file [**fig/pulumi.ts**](https://github.com/EthanOrlander/pulumi/blob/genFigSpec/fig/pulumi.ts) that contains the generated skeleton Fig autocomplete spec for Pulumi. (I ran it through Fig's linter manually if you're wondering why it looks formatted)

## What's next

1. Complete spec generator
    - Decide how to handle command arguments
      - It appears that in a Cobra CLI you only specify whether or not a command takes arguments, and the number or number range of arguments. The arguments are not defined with any other information (such as names)
    - Decide whether or not to include hidden commands in generated spec. Perhaps we add a Cobra Flag for this to the command, which defaults to false
    - It looks like when the package is imported to a Cobra CLI, it is updating some dependencies.
        - Appears this can be solved by using this package as a plugin
        - Alternatively, I can adjust this package so it's dependency versions align with those of Cobra
2. Create GitHub Action
    - Action could work as follows:
      1. Clone repo
      2. Add "github.com/EthanOrlander/genFigSpec" import to Cobra CLI root
         - Root file path provided as environment variable
      3. Add `GenFigSpec` command to root command
         - Root command reference (within root file) provided as environment variable
      4. Build
      5. Run `cli genFigSpec`
      6. Publish fig spec *somewhere?*
         - Maybe run a diff on the previous fig spec, and make a PR to [withfig/autocomplete](https://github.com/withfig/autocomplete) if anything's changed
         - This gets more complicated since this utility only generates a *partial skeleton*, and will be modified after
