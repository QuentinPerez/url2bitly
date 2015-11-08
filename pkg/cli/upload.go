package cli

var cmdUpload = &Command{
	Exec:        runUpload,
	UsageLine:   "upload URL",
	Description: "",
	Help:        "",
}

func runUpload(cmd *Command, rawArgs []string) error {
	if len(rawArgs) == 0 {
		return cmd.PrintUsage()
	}
	return nil
}
