package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type LambdaInfraStackProps struct {
	awscdk.StackProps
}

func NewLambdaInfraStack(scope constructs.Construct, id string, props *LambdaInfraStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here
	lambdaFn := awslambda.NewFunction(stack, jsii.String("LambdaFunction"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Code:    awslambda.Code_FromAsset(jsii.String("../bin"), &awss3assets.AssetOptions{}),
		Handler: jsii.String("lambdaHandler"),
	})

	fnUrl := lambdaFn.AddFunctionUrl(&awslambda.FunctionUrlOptions{
		AuthType: awslambda.FunctionUrlAuthType_NONE,
	})

	awscdk.NewCfnOutput(stack, jsii.String("TheUrl"), &awscdk.CfnOutputProps{
		Value: fnUrl.Url(),
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewLambdaInfraStack(app, "LambdaInfraStack", &LambdaInfraStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
