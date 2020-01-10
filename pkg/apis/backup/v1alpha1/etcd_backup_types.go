package v1alpha1

import (
	"gopkg.in/yaml.v2"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const etcdBackupCRDYAML = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: etcdbackups.backup.giantswarm.io
spec:
  group: backup.giantswarm.io
  scope: Cluster
  version: v1alpha1
  names:
    kind: ETCDBackup
    plural: etcdbackups
    singular: etcdbackup
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          type: object
          properties:
            guestBackup:
              type: boolean
            prefix:
              type: string
              pattern: "^[a-z0-9A-Z]+$"
            provider:
              enum:
              - aws
              - azure
              - kvm
            etcdV2:
              type: object
              properties:
                dataDir:
                  type: string
                  pattern: "^/"
              required:
              - dataDir
            etcdV3:
              type: object
              properties:
                endpoints:
                  type: string
                cacert:
                  type: string
                cert:
                  type: string
                key:
                  type: string
              required:
              - endpoints
              - cacert
              - cert
              - key
            storage:
              type: object
              properties:
                type:
                  enum:
                  - S3
                s3:
                  type: object
                  properties:
                    bucket:
                      type: string
                    region:
                      type: string
              required:
              - type
          required:
          - guestBackup
          - prefix
          - provider
          - etcdV2
          - etcdV3
          - storage
        status:
          type: object
          properties:
            status:
              enum:
              - Pending
              - Running
              - Completed
              - Failed
            startedTimestamp:
              type: string
              format: date-time
            finishedTimestamp:
              type: string
              format: date-time
            instances:
              type: object
              properties:
                attempts:
                  type: integer
                status:
                  enum:
                  - Pending
                  - Running
                  - Completed
                  - Failed
                startedTimestamp:
                  type: string
                  format: date-time
                finishedTimestamp:
                  type: string
                  format: date-time
              required:
          	  - status
              - attempts
          required:
          - status
          - instances
  additionalPrinterColumns:
  - name: guestBackup
    type: boolean
    description: Wether guest backups are backed up or not.
    JSONPath: .spec.guestBackup
  - name: Storage
    type: string
    description: The destination storage
    JSONPath: .spec.storage.type
  - name: Status
    type: string
    description: The status this backup is in
    JSONPath: .status.status
  - name: Started
    type: date
    description: The date the backup has been first attempted
    JSONPath: .status.startedTimestamp
  - name: Finished
    type: date
    description: The date the backup has been last attempted
    JSONPath: .status.finishedTimestamp
  - name: Attempts
    type: date
    description: The number of backups attempted
    JSONPath: .status.attempts
`

var etcdBackupCRD *apiextensionsv1beta1.CustomResourceDefinition

func init() {
	err := yaml.Unmarshal([]byte(etcdBackupCRDYAML), &etcdBackupCRD)
	if err != nil {
		panic(err)
	}
}

func NewEtcdBackupCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return etcdBackupCRD.DeepCopy()
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type EtcdBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              EtcdBackupSpec   `json:"spec"`
	Status            EtcdBackupStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

type EtcdBackupSpec struct {
	// GuestBackup is a boolean indicating if the tenant clusters have to be backupped
	GuestBackup bool `json:"guestBackup" yaml:"guestBackup"`
	// [mandatory] Prefix to use in etcd filenames
	Prefix string `json:"prefix" yaml:"prefix"`
	// [mandatory] Provider (aws, azure or kvm)
	Provider string `json:"provider" yaml:"provider"`
	// ETCDv2 backup settings
	ETCDv2 ETCDv2Settings `json:"etcdV2" yaml:"etcdV2"`
	// ETCDv3 backup settings
	ETCDv3 ETCDv3Settings `json:"etcdV3" yaml:"etcdV3"`
	// Settings for the backup storage
	Storage StorageSettings `json:"storage" yaml:"storage"`
}

type ETCDv2Settings struct {
	DataDir string `json:"dataDir" yaml:"dataDir"`
}

type ETCDv3Settings struct {
	Endpoints string `json:"endpoints" yaml:"endpoints"`
	CaCert    string `json:"cacert" yaml:"cacert"`
	Cert      string `json:"cert" yaml:"cert"`
	Key       string `json:"key" yaml:"key"`
}

type StorageSettings struct {
	// Storage type: only allowed value is "S3"
	Type string `json:"type" yaml:"type"`
	// Configuration for storage type: "S3"
	S3 S3Settings `json:"s3,omitempty" yaml:"s3,omitempty"`
}

type S3Settings struct {
	Bucket string `json:"bucket" yaml:"bucket"`
	Region string `json:"region" yaml:"region"`
}

type EtcdBackupStatus struct {
	// array for the state of the backup for all instances
	Instances []EtcdInstanceBackupStatus `json:"instances" yaml:"instances"`
	// Status of the whole backup job (can be 'Pending', 'Running'. 'Completed', 'Failed')
	Status string `json:"status" yaml:"status"`
	// Timestamp when the first attempt was made
	StartedTimestamp DeepCopyTime `json:"startedTimestamp,omitempty" yaml:"startedTimestamp"`
	// Timestamp when the last (final) attempt was made (when the Phase became either 'Completed' or 'Failed'
	FinishedTimestamp DeepCopyTime `json:"finishedTimestamp,omitempty" yaml:"finishedTimestamp"`
}

type EtcdInstanceBackupStatus struct {
	// Name of the tenant cluster or 'Control Plane'
	Name string `json:"name" yaml:"name"`
	// Status of this isntance's backup job (can be 'Pending', 'Running'. 'Completed', 'Failed')
	Status string `json:"status" yaml:"status"`
	// Attempts number of backup attempts made
	Attempts int8 `json:"attempts" yaml:"attempts"`
	// Timestamp when the first attempt was made
	StartedTimestamp DeepCopyTime `json:"startedTimestamp,omitempty" yaml:"startedTimestamp"`
	// Timestamp when the last (final) attempt was made (when the Phase became either 'Completed' or 'Failed'
	FinishedTimestamp DeepCopyTime `json:"finishedTimestamp,omitempty" yaml:"finishedTimestamp"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type EtcdBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []EtcdBackup `json:"items"`
}
