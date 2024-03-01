package accountlist

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/pkg/errors"
)

// GetSTSEnabledOrgClient accepts an AWS Config object, assumes the OrgRole, and
// returns an initialized client for AWS Organizations. See
// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/organizations#Client
// for more information.
func GetSTSEnabledOrgClient(config *aws.Config, orgRole string) *organizations.Client {
	// Connect STS as a credential provider.
	stsClient := sts.NewFromConfig(*config)
	config.Credentials = aws.NewCredentialsCache(
		stscreds.NewAssumeRoleProvider(stsClient, orgRole),
	)

	// Create service client value configured for credentials from assumed role.
	orgClient := organizations.NewFromConfig(*config)

	return orgClient
}

// GetOrgListAccountsPaginator accepts an organizations.Client object and
// returns a paginator object for the ListAccounts operation.
func GetOrgListAccountsPaginator(
	orgClient *organizations.Client,
	orgResultCount int32,
) *organizations.ListAccountsPaginator {
	paginator := organizations.NewListAccountsPaginator(
		orgClient,
		&organizations.ListAccountsInput{
			MaxResults: aws.Int32(orgResultCount),
		},
		func(o *organizations.ListAccountsPaginatorOptions) {
			o.Limit = orgResultCount
		},
	)

	return paginator
}

// CollectAccountTags iterates over the paginator to produce a result set of
// AccountTag objects. Optionally, you can pass a callback function which
// receives a single AccountTag object in a streaming manner and can perform
// an action.
func CollectAccountTags(
	ctx context.Context,
	orgClient *organizations.Client,
	paginator *organizations.ListAccountsPaginator,
	optCallback ...func(*AccountTag),
) ([]AccountTag, error) {
	accountTags := make([]AccountTag, 0)

	// Execute the AWS requests in a paginated effort.
	for paginator.HasMorePages() {
		results, err := paginator.NextPage(ctx)
		if err != nil {
			return accountTags, errors.Wrap(err, "failed to next page of results")
		}

		for i := range results.Accounts {
			account := &results.Accounts[i]

			orgTagResponse, err := orgClient.ListTagsForResource(ctx, &organizations.ListTagsForResourceInput{
				ResourceId: account.Id,
			})
			if err != nil {
				return accountTags, errors.Wrapf(err, "failed to list tags for account %s", *account.Id)
			}

			accountTag := AccountTag{
				ID:    *account.Id,
				ARN:   *account.Arn,
				Name:  *account.Name,
				Email: *account.Email,
			}

			for i := range orgTagResponse.Tags {
				tag := &orgTagResponse.Tags[i]

				accountTag.Tags = append(accountTag.Tags, TagValues{
					Key:   *tag.Key,
					Value: *tag.Value,
				})
			}

			accountTag.OUs, err = getOU(ctx, orgClient, *account.Id)
			if err != nil {
				return accountTags, err
			}

			// Run the callback.
			if len(optCallback) >= 1 {
				fn := optCallback[0]
				fn(&accountTag)
			}

			accountTags = append(accountTags, accountTag)
		}
	}

	return accountTags, nil
}

func getOU(
	ctx context.Context,
	orgClient *organizations.Client,
	accountID string,
) ([]OUType, error) {
	var results []OUType

	parents, err := orgClient.ListParents(ctx, &organizations.ListParentsInput{
		ChildId: aws.String(accountID),
	})
	if err != nil {
		return results, errors.Wrapf(err, "failed to list parents for account %s", accountID)
	}

	for i := range parents.Parents {
		parent := parents.Parents[i]

		if parent.Type == "ORGANIZATIONAL_UNIT" {
			ou, err := orgClient.DescribeOrganizationalUnit(ctx, &organizations.DescribeOrganizationalUnitInput{
				OrganizationalUnitId: parent.Id,
			})
			if err != nil {
				return results, errors.Wrapf(err, "failed to describe OU %s", *parent.Id)
			}

			results = append(results, OUType{
				accountID: accountID,
				ID:        *ou.OrganizationalUnit.Id,
				ARN:       *ou.OrganizationalUnit.Arn,
				Name:      *ou.OrganizationalUnit.Name,
			})
		}
	}

	return results, nil
}
