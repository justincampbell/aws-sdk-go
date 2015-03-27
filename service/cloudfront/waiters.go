// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudfront

import (
	"github.com/awslabs/aws-sdk-go/internal/waiter"
)
var waiterDistributionDeployed *waiter.Config

func (c *CloudFront) WaitUntilDistributionDeployed(input *GetDistributionInput) error {
	if waiterDistributionDeployed == nil {
		waiterDistributionDeployed = &waiter.Config{
			Operation:   "GetDistribution",
			Delay:       60,
			MaxAttempts: 25,
			Acceptors: []waiter.WaitAcceptor{
				waiter.WaitAcceptor{
					State:    "success",
					Matcher:  "path",
					Argument: "Distribution.Status",
					Expected: Deployed,
				},
				
			},
		}
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterDistributionDeployed,
	}
	return w.Wait()
}

var waiterInvalidationCompleted *waiter.Config

func (c *CloudFront) WaitUntilInvalidationCompleted(input *GetInvalidationInput) error {
	if waiterInvalidationCompleted == nil {
		waiterInvalidationCompleted = &waiter.Config{
			Operation:   "GetInvalidation",
			Delay:       20,
			MaxAttempts: 60,
			Acceptors: []waiter.WaitAcceptor{
				waiter.WaitAcceptor{
					State:    "success",
					Matcher:  "path",
					Argument: "Invalidation.Status",
					Expected: Completed,
				},
				
			},
		}
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterInvalidationCompleted,
	}
	return w.Wait()
}

var waiterStreamingDistributionDeployed *waiter.Config

func (c *CloudFront) WaitUntilStreamingDistributionDeployed(input *GetStreamingDistributionInput) error {
	if waiterStreamingDistributionDeployed == nil {
		waiterStreamingDistributionDeployed = &waiter.Config{
			Operation:   "GetStreamingDistribution",
			Delay:       60,
			MaxAttempts: 25,
			Acceptors: []waiter.WaitAcceptor{
				waiter.WaitAcceptor{
					State:    "success",
					Matcher:  "path",
					Argument: "StreamingDistribution.Status",
					Expected: Deployed,
				},
				
			},
		}
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterStreamingDistributionDeployed,
	}
	return w.Wait()
}