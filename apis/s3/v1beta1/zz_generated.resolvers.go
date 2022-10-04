/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta13 "github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1"
	v1beta12 "github.com/upbound/official-providers/provider-aws/apis/kms/v1beta1"
	v1beta11 "github.com/upbound/official-providers/provider-aws/apis/sns/v1beta1"
	v1beta1 "github.com/upbound/official-providers/provider-aws/apis/sqs/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this BucketACL.
func (mg *BucketACL) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BucketAccelerateConfiguration.
func (mg *BucketAccelerateConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BucketAnalyticsConfiguration.
func (mg *BucketAnalyticsConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageClassAnalysis); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.StorageClassAnalysis[i3].DataExport); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.StorageClassAnalysis[i3].DataExport[i4].Destination); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.StorageClassAnalysis[i3].DataExport[i4].Destination[i5].S3BucketDestination); i6++ {
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageClassAnalysis[i3].DataExport[i4].Destination[i5].S3BucketDestination[i6].BucketArn),
						Extract:      resource.ExtractParamPath("arn", true),
						Reference:    mg.Spec.ForProvider.StorageClassAnalysis[i3].DataExport[i4].Destination[i5].S3BucketDestination[i6].BucketArnRef,
						Selector:     mg.Spec.ForProvider.StorageClassAnalysis[i3].DataExport[i4].Destination[i5].S3BucketDestination[i6].BucketArnSelector,
						To: reference.To{
							List:    &BucketList{},
							Managed: &Bucket{},
						},
					})
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.StorageClassAnalysis[i3].DataExport[i4].Destination[i5].S3BucketDestination[i6].BucketArn")
					}
					mg.Spec.ForProvider.StorageClassAnalysis[i3].DataExport[i4].Destination[i5].S3BucketDestination[i6].BucketArn = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.StorageClassAnalysis[i3].DataExport[i4].Destination[i5].S3BucketDestination[i6].BucketArnRef = rsp.ResolvedReference

				}
			}
		}
	}

	return nil
}

// ResolveReferences of this BucketCorsConfiguration.
func (mg *BucketCorsConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BucketIntelligentTieringConfiguration.
func (mg *BucketIntelligentTieringConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BucketInventory.
func (mg *BucketInventory) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Destination); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Destination[i3].Bucket); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Destination[i3].Bucket[i4].BucketArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.Destination[i3].Bucket[i4].BucketArnRef,
				Selector:     mg.Spec.ForProvider.Destination[i3].Bucket[i4].BucketArnSelector,
				To: reference.To{
					List:    &BucketList{},
					Managed: &Bucket{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Destination[i3].Bucket[i4].BucketArn")
			}
			mg.Spec.ForProvider.Destination[i3].Bucket[i4].BucketArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Destination[i3].Bucket[i4].BucketArnRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this BucketLifecycleConfiguration.
func (mg *BucketLifecycleConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BucketLogging.
func (mg *BucketLogging) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TargetBucket),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.TargetBucketRef,
		Selector:     mg.Spec.ForProvider.TargetBucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TargetBucket")
	}
	mg.Spec.ForProvider.TargetBucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TargetBucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BucketMetric.
func (mg *BucketMetric) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BucketNotification.
func (mg *BucketNotification) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Queue); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Queue[i3].QueueArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.Queue[i3].QueueArnRef,
			Selector:     mg.Spec.ForProvider.Queue[i3].QueueArnSelector,
			To: reference.To{
				List:    &v1beta1.QueueList{},
				Managed: &v1beta1.Queue{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Queue[i3].QueueArn")
		}
		mg.Spec.ForProvider.Queue[i3].QueueArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Queue[i3].QueueArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Topic); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Topic[i3].TopicArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.Topic[i3].TopicArnRef,
			Selector:     mg.Spec.ForProvider.Topic[i3].TopicArnSelector,
			To: reference.To{
				List:    &v1beta11.TopicList{},
				Managed: &v1beta11.Topic{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Topic[i3].TopicArn")
		}
		mg.Spec.ForProvider.Topic[i3].TopicArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Topic[i3].TopicArnRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this BucketObject.
func (mg *BucketObject) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1beta12.KeyList{},
			Managed: &v1beta12.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyID")
	}
	mg.Spec.ForProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BucketObjectLockConfiguration.
func (mg *BucketObjectLockConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BucketOwnershipControls.
func (mg *BucketOwnershipControls) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BucketPolicy.
func (mg *BucketPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BucketPublicAccessBlock.
func (mg *BucketPublicAccessBlock) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BucketReplicationConfiguration.
func (mg *BucketReplicationConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Role),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.RoleRef,
		Selector:     mg.Spec.ForProvider.RoleSelector,
		To: reference.To{
			List:    &v1beta13.RoleList{},
			Managed: &v1beta13.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Role")
	}
	mg.Spec.ForProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Rule); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Rule[i3].Destination); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Rule[i3].Destination[i4].Bucket),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.Rule[i3].Destination[i4].BucketRef,
				Selector:     mg.Spec.ForProvider.Rule[i3].Destination[i4].BucketSelector,
				To: reference.To{
					List:    &BucketList{},
					Managed: &Bucket{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Rule[i3].Destination[i4].Bucket")
			}
			mg.Spec.ForProvider.Rule[i3].Destination[i4].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Rule[i3].Destination[i4].BucketRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this BucketRequestPaymentConfiguration.
func (mg *BucketRequestPaymentConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BucketServerSideEncryptionConfiguration.
func (mg *BucketServerSideEncryptionConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Rule); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Rule[i3].ApplyServerSideEncryptionByDefault); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Rule[i3].ApplyServerSideEncryptionByDefault[i4].KMSMasterKeyID),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.Rule[i3].ApplyServerSideEncryptionByDefault[i4].KMSMasterKeyIDRef,
				Selector:     mg.Spec.ForProvider.Rule[i3].ApplyServerSideEncryptionByDefault[i4].KMSMasterKeyIDSelector,
				To: reference.To{
					List:    &v1beta12.KeyList{},
					Managed: &v1beta12.Key{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Rule[i3].ApplyServerSideEncryptionByDefault[i4].KMSMasterKeyID")
			}
			mg.Spec.ForProvider.Rule[i3].ApplyServerSideEncryptionByDefault[i4].KMSMasterKeyID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Rule[i3].ApplyServerSideEncryptionByDefault[i4].KMSMasterKeyIDRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this BucketVersioning.
func (mg *BucketVersioning) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BucketWebsiteConfiguration.
func (mg *BucketWebsiteConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Object.
func (mg *Object) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1beta12.KeyList{},
			Managed: &v1beta12.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyID")
	}
	mg.Spec.ForProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyIDRef = rsp.ResolvedReference

	return nil
}