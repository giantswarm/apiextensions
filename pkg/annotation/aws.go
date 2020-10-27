package annotation

// AWSUpdateMaxBatchSize is the aws update annotation used for configuring
// maximum batch size for instances during ASG update.
const AWSUpdateMaxBatchSize = "aws.giantswarm.io/update-max-batch-size"

// AWSUpdatePauseTime is the aws update annotation used for configuring
// time pause between rolling a single batch during ASG update.
const AWSUpdatePauseTime = "aws.giantswarm.io/update-pause-time"
