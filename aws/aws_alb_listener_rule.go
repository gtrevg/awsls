// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

func ListAlbListenerRule(client *Client) ([]Resource, error) {
	req := client.elasticloadbalancingv2conn.DescribeRulesRequest(&elasticloadbalancingv2.DescribeRulesInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Rules) > 0 {
		for _, r := range resp.Rules {

			result = append(result, Resource{
				Type:   "aws_alb_listener_rule",
				ID:     *r.RuleArn,
				Region: client.elasticloadbalancingv2conn.Config.Region,
			})
		}
	}

	return result, nil
}
