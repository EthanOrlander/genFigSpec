# genFigSpec

This package contains a Cobra Command "GenFigSpec", which introspects a Cobra CLI and produces a partial skeleton Fig autocomplete spec.

This package is still in development. The end goal is for it to be used in a GitHub Action that can be easily attached to any Cobra CLI.

## Usage

The method of usage will evolve quickly as this package is developed. This is how you should use it in its current state:

1. In your Cobra CLI's source code, import this package in the root file for the CLI
    - Import using "github.com/EthanOrlander/genFigSpec"
2. Attach the command to the root command. This will look `<rootCmd>.AddCommand(genFigSpec.GenFigSpec)`
3. Once you rebuild the CLI with this update, simply use `<cliname> genFigSpec > <cliname>.ts` in your command line.
4. This creates a TypeScript file contains the skeletal Fig autocomplete spec for your CLI. It can now be copied to [fig/autocomplete](https://github.com/withfig/autocomplete) and completed.

## Example

As an example, I have performed the above steps in [this fork of pulumi](https://github.com/withfig/autocomplete).
You'll see that **pkg/cmd/pulumi/pulumi.go** now imports this package and attaches the command on line 218 using `cmd.AddCommand(genFigSpec.GenFigSpec)`.
You'll also see a file at the **fig/pulumi.ts** that contains the skeleton autocomplete spec for Pulumi.

## What's next

1. Complete spec generator
    - Decide how to handle command arguments
    - Decode whether or not to include hidden commands in generated spec
2. Create GitHub Action
