package annotation

// support:
//   - crd: awsclusters.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: "Since 14.0.0"
// documentation:
//   AWSCNIMinimumIPTarget is an annotation to configure the value for MINIMUM_IP_TARGET for AWS CNI.
//   See [CNI Configuration Variables](https://github.com/aws/amazon-vpc-cni-k8s#cni-configuration-variables)
//   and [ENI and IP Target](https://github.com/aws/amazon-vpc-cni-k8s/blob/master/docs/eni-and-ip-target.md)
const AWSCNIMinimumIPTarget = "alpha.cni.aws.giantswarm.io/minimum-ip-target"

// support:
//   - crd: awsclusters.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: "Since 14.0.0"
// documentation:
//   AWSCNIWarmIPTarget is an annotation to configure the value for WARM_IP_TARGET for AWS CNI.
//   See [CNI Configuration Variables](https://github.com/aws/amazon-vpc-cni-k8s#cni-configuration-variables)
//   and [ENI and IP Target](https://github.com/aws/amazon-vpc-cni-k8s/blob/master/docs/eni-and-ip-target.md)
const AWSCNIWarmIPTarget = "alpha.cni.aws.giantswarm.io/warm-ip-target"

// support:
//   - crd: awsclusters.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: "Since 12.7.0"
//   - crd: awsmachinedeployments.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: "Since 12.7.0"
// documentation:
//   AWSUpdateMaxBatchSize is the aws update annotation used for configuring
//   maximum batch size for instances during ASG update.
//   The value can be either a whole number specifying the number of instances
//   or a percentage of total instances as decimal number ie `0.3` for 30%.
//   See [Fine Tuning Upgrades](https://docs.giantswarm.io/guides/fine-tuning-upgrade-disruption-on-aws/)
//   and [AWS Documentation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-rollingupdate-maxbatchsize) for additional information.
const AWSUpdateMaxBatchSize = "alpha.aws.giantswarm.io/update-max-batch-size"

// support:
//   - crd: awsclusters.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: "Since 12.7.0"
//   - crd: awsmachinedeployments.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: "Since 12.7.0"
// documentation:
//   AWSUpdatePauseTime is the aws update annotation used for configuring time pause between rolling a single batch during ASG update.
//   The value must be in ISO 8601 duration format, e. g. "PT5M" for five minutes or "PT10S" for 10 seconds.
//   See [Fine Tuning Upgrades](https://docs.giantswarm.io/guides/fine-tuning-upgrade-disruption-on-aws/)
//   and [AWS Documentation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-rollingupdate-maxbatchsize) for additional information.
const AWSUpdatePauseTime = "alpha.aws.giantswarm.io/update-pause-time"

// support:
//   - crd: awsclusters.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: "Since 12.7.0"
//   - crd: awsmachinedeployments.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: "Since 12.7.0"
// documentation:
//   AWSMetadataV2 configures token usage for your AWS EC2 instance metadata requests.
//   If the value is 'optional', you can choose to retrieve instance metadata with or without a signed token
//   header on your request. If you retrieve the IAM role credentials without a token, the version 1.0 role
//   credentials are returned. If you retrieve the IAM role credentials using a valid signed token, the version
//   2.0 role credentials are returned.
//   If the state is 'required', you must send a signed token header with any instance metadata retrieval
//   requests. In this state, retrieving the IAM role credentials always returns the version 2.0 credentials; the
//   version 1.0 credentials are not available.
//   Default value is 'optional'
//
//   [AWS Documentation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-metadataoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-metadataoptions-httptokens)
const AWSMetadataV2 = "alpha.aws.giantswarm.io/metadata-v2"

// support:
//   - crd: awsclusters.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: "Since 12.7.0"
//   - crd: awsmachinedeployments.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: "Since 12.7.0"
// documentation:
//   AWSSubnetSize is the aws update annotation used for configuring the subnet size of AWSCluster or AWSMachineDeployment.
//   The value is a number that will represent the subnet mask used when creating the subnet. This value must be smaller than 28 due to AWS restrictions.
const AWSSubnetSize = "alpha.aws.giantswarm.io/aws-subnet-size"
