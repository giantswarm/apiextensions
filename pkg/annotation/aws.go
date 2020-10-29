package annotation

// AWSUpdateMaxBatchSize is the aws update annotation used for configuring
// maximum batch size for instances during ASG update.
// The value can be either a whole number specifying the number of instances
// or a percentage of total instances as decimal number ie: `0.3` for 30%.
// https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-rollingupdate-maxbatchsize
const AWSUpdateMaxBatchSize = "aws.giantswarm.io/update-max-batch-size"

// AWSUpdatePauseTime is the aws update annotation used for configuring
// time pause between rolling a single batch during ASG update.
// The value must be in ISO 8601 duration format, e. g. "PT5M" for five minutes or "PT10S" for 10 seconds.
// https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-rollingupdate-pausetime
const AWSUpdatePauseTime = "aws.giantswarm.io/update-pause-time"
