# genFigSpec

This package contains a Cobra Command factory "NewCmdGenFigSpec", which introspects a Cobra CLI and produces a partial skeleton Fig autocomplete spec.

## Usage

1. Add this package to your project using `go get github.com/EthanOrlander/genFigSpec`
2. In your Cobra CLI's source code, import this package in the file containing the root command
    - Add "github.com/EthanOrlander/genFigSpec" to imports
3. Attach the fig spec generator command to a command. This will look like `<cmd>.AddCommand(genFigSpec.NewCmdGenFigSpec())`
   - Often the command will be attached to the root command or to a `generate` or `gen` subcommand.
   - NewCmdGenFigSpec accepts an optional argument of type object `genFigSpec.Opts` that can be used to customize the command. [See how to use it here](#Opts).
4. Update go.sum using `go mod tidy`
5. If the project uses go vendors, run `go mod vendor`
6. Once you rebuild the CLI with this update, run the command. If you are using the default command, use `<cliname> genFigSpec > <cliname>.ts` in your command line.
7. This creates a TypeScript file **\<cliname\>.ts** containing the skeletal Fig autocomplete spec for your CLI. It can now be copied to [fig/autocomplete](https://github.com/withfig/autocomplete) and completed.

## Opts (customizing the command)

A `genFigSpec.Opts` object can be passed as an argument to the NewCmdGenFigSpec() factory function to customize the command:

```go
type Opts struct {
 Use                 string                      # Override provided cobra.Command.Use
 Short               string                      # Override provided cobra.Command.Use
 Visible             bool                        # Override provided cobra.Command.Hidden
 Long                string                      # Override provided cobra.Command.Long
 commandArgGenerator func(*cobra.Command) Args   # You can also provide a function to generate command arguments
}
```

### Example

```go
generatorCmd.AddCommand(genFigSpec.NewCmdGenFigSpec(&> genFigSpec.Opts{
   Use: "fig",
   Short: "Generate a skeleton fig spec for mycli",
   Visible: true,
   commandArgGenerator: func(cmd *cobra.Command){
      var args []Arg
      # ...generate args from cmd
      return args
   }
}))
```

## Using your own command

If you require even more customization, you can create your own command and simply use the `genFigSpec.MakeFigSpec(rootCmd)` function. This function returns a `Spec` object, from which the TypeScript Fig Spec can be produced using the `toTypescript()`  receiver on `Spec`. The `MakeFigSpec` function accepts a command from which to generate the Spec. In most cases, this will be the root command.


## Example

As an example, I have performed the above steps in [this fork of pulumi](https://github.com/EthanOrlander/pulumi/tree/genFigSpec).
You'll see that [**pkg/cmd/pulumi/pulumi.go**](https://github.com/EthanOrlander/pulumi/blob/genFigSpec/pkg/cmd/pulumi/pulumi.go#L39) now imports this package and attaches the command on [line 218](https://github.com/EthanOrlander/pulumi/blob/genFigSpec/pkg/cmd/pulumi/pulumi.go#L218) using `cmd.AddCommand(genFigSpec.NewCmdGenFigSpec(&genFigSpec.Opts{...}))`.
You'll also see a file [**fig/pulumi.ts**](https://github.com/EthanOrlander/pulumi/blob/genFigSpec/fig/pulumi.ts) that contains the generated skeleton Fig autocomplete spec for Pulumi. (I ran it through Fig's linter manually if you're wondering why it looks formatted).

I've also added the command in [this fork of OpenShift](https://github.com/EthanOrlander/oc/tree/genFigSpec).
