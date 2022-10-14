package regions

import (
	awsCommon "github.com/adrianriobo/qenvs/pkg/infra/aws"
	utilInfra "github.com/adrianriobo/qenvs/pkg/infra/util"
	"github.com/adrianriobo/qenvs/pkg/util"
)

func GetRegions(projectName, backedURL string) ([]string, error) {
	stack := utilInfra.Stack{
		StackName:   StackGetRegionsName,
		ProjectName: projectName,
		BackedURL:   backedURL,
		Plugin:      awsCommon.PluginAWSDefault,
		DeployFunc:  getDefaultRegionsRequest().GetRegions,
	}
	// Exec stack
	stackResult, err := utilInfra.UpStack(stack)
	if err != nil {
		return nil, err
	}
	regions, ok := stackResult.Outputs[StackGetRegionsOutputAWSRegions].Value.([]interface{})
	if !ok {
		return nil, err
	}
	return util.ArrayConvert[string](regions), nil
}
