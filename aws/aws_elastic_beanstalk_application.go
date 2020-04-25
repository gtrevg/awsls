// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
)

func ListElasticBeanstalkApplication(client *Client) ([]Resource, error) {
	req := client.elasticbeanstalkconn.DescribeApplicationsRequest(&elasticbeanstalk.DescribeApplicationsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Applications) > 0 {
		for _, r := range resp.Applications {

			result = append(result, Resource{
				Type:   "aws_elastic_beanstalk_application",
				ID:     *r.ApplicationName,
				Region: client.elasticbeanstalkconn.Config.Region,
			})
		}
	}

	return result, nil
}
