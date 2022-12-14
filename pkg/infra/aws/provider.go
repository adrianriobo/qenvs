package aws

import (
	"context"
	"os"

	infraUtil "github.com/adrianriobo/qenvs/pkg/infra/util"
	"github.com/adrianriobo/qenvs/pkg/util/logging"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

const (
	CONFIG_AWS_REGION     string = "aws:region"
	CONFIG_AWS_ACCESS_KEY string = "aws:accessKey"
	CONFIG_AWS_SECRET_KEY string = "aws:secretKey"
)

// pulumi config key : aws env credential
var credentials = map[string]string{
	CONFIG_AWS_REGION:     "AWS_DEFAULT_REGION",
	CONFIG_AWS_ACCESS_KEY: "AWS_ACCESS_KEY_ID",
	CONFIG_AWS_SECRET_KEY: "AWS_SECRET_ACCESS_KEY"}

func SetAWSCredentials(ctx context.Context, stack auto.Stack, fixedCredentials map[string]string) error {
	for configKey, envKey := range credentials {
		if value, ok := fixedCredentials[configKey]; ok {
			if err := stack.SetConfig(ctx, configKey,
				auto.ConfigValue{Value: value}); err != nil {
				logging.Errorf("Failed setting credential: %v", err)
				return err
			}
		} else {
			if err := stack.SetConfig(ctx, configKey,
				auto.ConfigValue{Value: os.Getenv(envKey)}); err != nil {
				logging.Errorf("Failed setting credential: %v", err)
				return err
			}
		}
	}
	return nil
}

func GetPluginAWS(fixedCredentials map[string]string) infraUtil.PluginInfo {
	return infraUtil.PluginInfo{
		Name:              "aws",
		Version:           "v4.0.0",
		SetCredentialFunc: SetAWSCredentials,
		FixedCredentials:  fixedCredentials}
}

var PluginAWSDefault = GetPluginAWS(nil)
