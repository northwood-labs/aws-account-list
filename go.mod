module github.com/northwood-labs/aws-account-list

go 1.22

toolchain go1.22.0

require (
	github.com/aws/aws-sdk-go-v2 v1.27.2
	github.com/aws/aws-sdk-go-v2/credentials v1.17.4
	github.com/aws/aws-sdk-go-v2/service/organizations v1.25.1
	github.com/aws/aws-sdk-go-v2/service/sts v1.28.12
	github.com/northwood-labs/awsutils v0.0.0-20220620172853-924504e83dfb
	github.com/northwood-labs/golang-utils/exiterrorf v0.0.0-20240301191325-850f76df0fb0
	github.com/pkg/errors v0.9.1
)

require (
	github.com/aws/aws-sdk-go-v2/config v1.27.4 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.15.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.9 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.9 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.11.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.11.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.20.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.23.1 // indirect
	github.com/aws/smithy-go v1.20.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
)
