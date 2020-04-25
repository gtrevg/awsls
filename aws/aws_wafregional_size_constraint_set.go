// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func ListWafregionalSizeConstraintSet(client *Client) ([]Resource, error) {
	req := client.wafregionalconn.ListSizeConstraintSetsRequest(&wafregional.ListSizeConstraintSetsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.SizeConstraintSets) > 0 {
		for _, r := range resp.SizeConstraintSets {

			result = append(result, Resource{
				Type:   "aws_wafregional_size_constraint_set",
				ID:     *r.SizeConstraintSetId,
				Region: client.wafregionalconn.Config.Region,
			})
		}
	}

	return result, nil
}
