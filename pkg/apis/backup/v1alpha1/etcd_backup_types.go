package v1alpha1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/apis/backup"
	"github.com/giantswarm/apiextensions/pkg/crd"
	"github.com/giantswarm/apiextensions/pkg/key"
)

// NewETCDBackupCRD returns a CRD defining an ETCDBackup.
func NewETCDBackupCRD() *v1.CustomResourceDefinition {
	return crd.LoadV1(backup.Group, backup.KindETCDBackup)
}

// NewETCDBackupCR returns an EtcdBackup custom resource.
func NewETCDBackupCR(name string) *ETCDBackup {
	cr := ETCDBackup{}
	cr.TypeMeta, cr.ObjectMeta = key.NewMeta(SchemeGroupVersion, backup.KindETCDBackup, name, "")
	return &cr
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=common;giantswarm,scope=Cluster

type ETCDBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ETCDBackupSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	Status ETCDBackupStatus `json:"status,omitempty"`
}

type ETCDBackupSpec struct {
	// GuestBackup is a boolean indicating if the tenant clusters have to be backupped
	GuestBackup bool `json:"guestBackup"`
}

type ETCDBackupStatus struct {
	// map containing the state of the backup for all instances
	Instances map[string]ETCDInstanceBackupStatusIndex `json:"instances,omitempty"`
	// Status of the whole backup job (can be 'Pending', 'Running'. 'Completed', 'Failed')
	Status string `json:"status"`
	// Timestamp when the first attempt was made
	StartedTimestamp metav1.Time `json:"startedTimestamp,omitempty"`
	// Timestamp when the last (final) attempt was made (when the Phase became either 'Completed' or 'Failed'
	FinishedTimestamp metav1.Time `json:"finishedTimestamp,omitempty"`
}

type ETCDInstanceBackupStatusIndex struct {
	// Name of the tenant cluster or 'Control Plane'
	Name string `json:"name"`
	// Status of the V2 backup for this instance
	V2 ETCDInstanceBackupStatus `json:"v2"`
	// Status of the V3 backup for this instance
	V3 ETCDInstanceBackupStatus `json:"v3"`
}

type ETCDInstanceBackupStatus struct {
	// Status of this isntance's backup job (can be 'Pending', 'Running'. 'Completed', 'Failed')
	Status string `json:"status"`
	// Timestamp when the first attempt was made
	StartedTimestamp metav1.Time `json:"startedTimestamp,omitempty"`
	// Timestamp when the last (final) attempt was made (when the Phase became either 'Completed' or 'Failed'
	FinishedTimestamp metav1.Time `json:"finishedTimestamp,omitempty"`
	// Latest backup error message
	LatestError string `json:"latestError,omitempty"`
	// Time took by the backup creation process
	CreationTime int64 `json:"creationTime,omitempty"`
	// Time took by the backup encryption process
	EncryptionTime int64 `json:"encryptionTime,omitempty"`
	// Time took by the backup upload process
	UploadTime int64 `json:"uploadTime,omitempty"`
	// Size of the backup file
	BackupFileSize int64 `json:"backupFileSize,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ETCDBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ETCDBackup `json:"items"`
}
