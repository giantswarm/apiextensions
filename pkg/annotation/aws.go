package annotation

// AWSUpdateMaxBatchSize is the aws update annotation used for configuring
// maximum batch size for instances during ASG update.
// The value can be either a whole number specifying the number of instances
// or a percentage of total instances as decimal number ie: `0.3` for 30%.
// https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-rollingupdate-maxbatchsize
const AWSUpdateMaxBatchSize = "alpha.aws.giantswarm.io/update-max-batch-size"

// AWSUpdatePauseTime is the aws update annotation used for configuring
// time pause between rolling a single batch during ASG update.
// The value must be in ISO 8601 duration format, e. g. "PT5M" for five minutes or "PT10S" for 10 seconds.
// https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-rollingupdate-pausetime
const AWSUpdatePauseTime = "alpha.aws.giantswarm.io/update-pause-time"

// AWSMetadataV2 is the aws update annotation used for configuring
// the state of token usage for your instance metadata requests.
// The value can be either "optional" or "required". Default is "optional".
// https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-metadataoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-metadataoptions-httptokens
const AWSMetadataV2 = "alpha.aws.giantswarm.io/metadata-v2"

// AWSSubnetSize is the aws update annotation used for configuring
// the subnet size of AWSCluster or AWSMachineDeployment.
// The value is a number that will represent the subnet mask used when creating the subnet. This value must be smaller than 28 due to AWS restrictions.
const AWSSubnetSize = "alpha.aws.giantswarm.io/aws-subnet-size"
