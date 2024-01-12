package hosts

import (
	params "github.com/adrianriobo/qenvs/cmd/cmd/constants"
	qenvsContext "github.com/adrianriobo/qenvs/pkg/manager/context"
	"github.com/adrianriobo/qenvs/pkg/provider/aws/action/mac"
	"github.com/adrianriobo/qenvs/pkg/util/logging"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	cmdMac         = "mac"
	cmdMacDesc     = "manage mac instances"
	requestCmd     = "request"
	requestCmdDesc = "request mac machine"
	releaseCmd     = "release"
	releaseCmdDesc = "release mac machine"

	arch              string = "arch"
	archDesc          string = "mac architecture allowed values x86, m1, m2"
	archDefault       string = "m2"
	osVersion         string = "version"
	osVersionDesc     string = "macos operating system vestion 11, 12 on x86 and m1/m2; 13, 14 on all archs"
	osDefault         string = "14"
	fixedLocation     string = "fixed-location"
	fixedLocationDesc string = "if this flag is set the host will be created only on the region set by the AWS Env (AWS_DEFAULT_REGION)"
)

func GetMacCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   cmdMac,
		Short: cmdMacDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
	}
	c.AddCommand(getMacCreate(), getMacDestroy(), getMacRequest(), getMacRelease())
	return c
}

func getMacCreate() *cobra.Command {
	c := &cobra.Command{
		Use:   params.CreateCmdName,
		Short: params.CreateCmdName,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}

			// Initialize context
			qenvsContext.Init(
				viper.GetString(params.ProjectName),
				viper.GetString(params.BackedURL),
				viper.GetString(params.ConnectionDetailsOutput),
				viper.GetStringMapString(params.Tags))

			// Run create
			if err := mac.Create(
				&mac.MacRequest{
					Prefix:        "main",
					Architecture:  viper.GetString(arch),
					Version:       viper.GetString(osVersion),
					FixedLocation: viper.IsSet(fixedLocation),
					Airgap:        viper.IsSet(airgap)}); err != nil {
				logging.Error(err)
			}
			return nil
		},
	}
	flagSet := pflag.NewFlagSet(params.CreateCmdName, pflag.ExitOnError)
	flagSet.StringP(params.ConnectionDetailsOutput, "", "", params.ConnectionDetailsOutputDesc)
	flagSet.StringToStringP(params.Tags, "", nil, params.TagsDesc)
	flagSet.StringP(arch, "", archDefault, archDesc)
	flagSet.StringP(osVersion, "", osDefault, osVersionDesc)
	flagSet.Bool(fixedLocation, false, fixedLocationDesc)
	flagSet.Bool(airgap, false, airgapDesc)
	c.PersistentFlags().AddFlagSet(flagSet)
	return c
}

func getMacRequest() *cobra.Command {
	c := &cobra.Command{
		Use:   requestCmd,
		Short: requestCmd,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}

			// Initialize context
			qenvsContext.Init(
				viper.GetString(params.ProjectName),
				viper.GetString(params.BackedURL),
				viper.GetString(params.ConnectionDetailsOutput),
				viper.GetStringMapString(params.Tags))

			// Run create
			if err := mac.Request(
				&mac.MacRequest{
					Prefix:        "main",
					Architecture:  viper.GetString(arch),
					Version:       viper.GetString(osVersion),
					FixedLocation: viper.IsSet(fixedLocation),
					Airgap:        viper.IsSet(airgap)}); err != nil {
				logging.Error(err)
			}
			return nil
		},
	}
	flagSet := pflag.NewFlagSet(requestCmd, pflag.ExitOnError)
	flagSet.StringP(params.ConnectionDetailsOutput, "", "", params.ConnectionDetailsOutputDesc)
	flagSet.StringToStringP(params.Tags, "", nil, params.TagsDesc)
	flagSet.StringP(arch, "", archDefault, archDesc)
	flagSet.StringP(osVersion, "", osDefault, osVersionDesc)
	flagSet.Bool(fixedLocation, false, fixedLocationDesc)
	flagSet.Bool(airgap, false, airgapDesc)
	c.PersistentFlags().AddFlagSet(flagSet)
	return c
}

func getMacRelease() *cobra.Command {
	c := &cobra.Command{
		Use:   releaseCmd,
		Short: releaseCmd,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}

			// Initialize context
			qenvsContext.Init(
				viper.GetString(params.ProjectName),
				viper.GetString(params.BackedURL),
				viper.GetString(params.ConnectionDetailsOutput),
				viper.GetStringMapString(params.Tags))

			// Run create
			if err := mac.Release(
				&mac.MacRequest{
					Prefix:       "main",
					Architecture: viper.GetString(arch),
					Version:      viper.GetString(osVersion),
				}); err != nil {
				logging.Error(err)
			}
			return nil
		},
	}
	flagSet := pflag.NewFlagSet(releaseCmd, pflag.ExitOnError)
	flagSet.StringToStringP(params.Tags, "", nil, params.TagsDesc)
	flagSet.StringP(arch, "", archDefault, archDesc)
	flagSet.StringP(osVersion, "", osDefault, osVersionDesc)
	c.PersistentFlags().AddFlagSet(flagSet)
	return c
}

func getMacDestroy() *cobra.Command {
	c := &cobra.Command{
		Use:   params.DestroyCmdName,
		Short: params.DestroyCmdName,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}

			qenvsContext.InitBase(
				viper.GetString(params.ProjectName),
				viper.GetString(params.BackedURL))

			if err := mac.Destroy(
				"main",
				viper.GetString(arch)); err != nil {
				logging.Error(err)
			}
			return nil
		},
	}
	flagSet := pflag.NewFlagSet(params.DestroyCmdName, pflag.ExitOnError)
	flagSet.StringP(arch, "", archDefault, archDesc)
	c.PersistentFlags().AddFlagSet(flagSet)
	return c
}
