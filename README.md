# AWS-LAMBDA-GOLANG

Example repository of how to deploy an AWS Lambda Function written in Go using CDK Go SDK.

## Prerequisites


Check if your environment variables are set

```bash

export CDK_DEFAULT_ACCOUNT=
export CDK_DEFAULT_REGION=

```


## Steps


1. Run go build

```bash
make
```

2. Run cdk bootstrap

```bash
make bootstrap
```

3. Run cdk deploy 

```bash
make deploy
```

4. At the end run cdk deploy, when are you done

```bash
make destroy
```
