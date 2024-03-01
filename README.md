# AWS Account List

Generates a list of all AWS accounts registered in an AWS Organizations account.

## Install as a CLI tool

1. You must have the Golang toolchain installed first.

    ```bash
    brew install go
    ```

1. Add `$GOPATH/bin` to your `$PATH` environment variable. By default (i.e., without configuration), `$GOPATH` is defined as `$HOME/go`.

    ```bash
    export PATH="$PATH:$GOPATH/bin"
    ```

1. Once you've done everything above, you can use `go get`.

    ```bash
    go get -u github.com/northwood-labs/aws-account-list
    ```

## Usage as CLI Tool

Examples assume the use of [AWS Vault] and [AWS Identity Center].

Gets a list of AWS accounts that are part of the AWS Organization as JSON.

```bash
aws-account-list --help
```

Read directly from the AWS Organizations management account.

```bash
aws-vault exec management-account -- aws-account-list
```

Assume the `AWS_ORG_ROLE` IAM role first, then read the AWS Organizations management account using that IAM role.

```bash
AWS_ORG_ROLE="arn:aws:iam::0123456789012:role/OrganizationReadOnlyAccess"
aws-vault exec management-account -- aws-account-list
```

## Usage as Library

This can also be used as a library in your own applications for generating a list in-memory. The library should fetch data for accounts asynchronously for better performance, but does not yet. This has been tested on AWS Organizations up to ~200 accounts.

```go
import "github.com/northwood-labs/aws-account-list/accountlist"
```

See `main.go`, which implements this library to produce this very same CLI tool.

[AWS Identity Center]: https://aws.amazon.com/iam/identity-center/
[AWS Vault]: https://github.com/99designs/aws-vault
