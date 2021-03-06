// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
)

func ListSfnStateMachine(client *Client) ([]Resource, error) {
	req := client.sfnconn.ListStateMachinesRequest(&sfn.ListStateMachinesInput{})

	var result []Resource

	p := sfn.NewListStateMachinesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.StateMachines {

			t := *r.CreationDate
			result = append(result, Resource{
				Type:   "aws_sfn_state_machine",
				ID:     *r.StateMachineArn,
				Region: client.sfnconn.Config.Region,

				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
