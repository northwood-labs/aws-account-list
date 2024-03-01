package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/northwood-labs/aws-account-list/accountlist"
	"github.com/northwood-labs/awsutils"
	"github.com/northwood-labs/golang-utils/exiterrorf"
)

const (
	// The number of results to request per page.
	orgResultCount = 20
)

var (
	// Global flags.
	retries *int
	verbose *bool

	// Context.
	ctx = context.Background()

	orgRole = os.Getenv("AWS_ORG_ROLE")
)

func main() {
	// Flags
	retries = flag.Int("retries", 5, "The number of times to retry failed requests to AWS.")
	verbose = flag.Bool("verbose", false, "Output internal data to stdout.")
	flag.Parse()

	config, err := awsutils.GetAWSConfig(ctx, "", "", *retries, *verbose)
	if err != nil {
		exiterrorf.ExitErrorf(err)
	}

	var orgClient *organizations.Client

	if orgRole == "" {
		orgClient = organizations.NewFromConfig(config)
	} else {
		orgClient = accountlist.GetSTSEnabledOrgClient(&config, orgRole)
	}
	paginator := accountlist.GetOrgListAccountsPaginator(orgClient, orgResultCount)

	accountTags, err := accountlist.CollectAccountTags(ctx, orgClient, paginator)
	if err != nil {
		exiterrorf.ExitErrorf(err)
	}

	b, err := json.Marshal(accountTags)
	if err != nil {
		exiterrorf.ExitErrorf(err)
	}

	fmt.Println(string(b))
}
