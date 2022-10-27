package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/adrianriobo/qenvs/pkg/infra/aws/modules/environment"
	"github.com/adrianriobo/qenvs/pkg/util/logging"
)

const (
	hostCmdCreateDescription string = "create a supported host"
)

func init() {
	hostCmd.AddCommand(hostCreateCmd)
	flagSet := pflag.NewFlagSet(createCmdName, pflag.ExitOnError)
	flagSet.StringP(projectName, "", "", projectNameDesc)
	flagSet.StringP(backedURL, "", "", backedURLDesc)
	flagSet.StringP(supportedHostID, "", "", supportedHostIDDesc)
	hostCreateCmd.Flags().AddFlagSet(flagSet)
}

var hostCreateCmd = &cobra.Command{
	Use:   createCmdName,
	Short: hostCmdCreateDescription,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}
		if err := environment.Create(
			viper.GetString(projectName),
			viper.GetString(backedURL),
			// fixed as public to ensure sync, when PR is merged we can offer as param
			// https://github.com/pulumi/pulumi-command/pull/132
			true,
			viper.GetString(supportedHostID)); err != nil {
			logging.Error(err)
		}
		return nil
	},
}
