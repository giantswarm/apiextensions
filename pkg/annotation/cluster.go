package annotation

// ClusterDescription is the cluster annotation used for storing
// a customer's cluster description.
const ClusterDescription = "cluster.giantswarm.io/description"

// support:
//   - crd: clusters.cluster.x-k8s.io
//     apiversion: v1alpha3
// documentation:
//   This annotation is used to define the desired target release for a scheduled upgrade of the cluster.
//   The upgrade to the specified version will be applied if the "update-schedule-target-time" annotation has been set
//   and the time defined there has been reached. The value has to be only the desired release version, e.g "15.2.1".
const UpdateScheduleTargetRelease = "alpha.giantswarm.io/update-schedule-target-release"

// support:
//   - crd: clusters.cluster.x-k8s.io
//     apiversion: v1alpha3
// documentation:
//   This annotation is used to define the desired target time for a scheduled upgrade of the cluster.
//   The upgrade will be applied at the specified time if the "update-schedule-target-release" annotation has been set
//   to the target release version. The value has to be in RFC822 Format and UTC time zone. e.g. "30 Jan 21 15:04 UTC"
const UpdateScheduleTargetTime = "alpha.giantswarm.io/update-schedule-target-time"
