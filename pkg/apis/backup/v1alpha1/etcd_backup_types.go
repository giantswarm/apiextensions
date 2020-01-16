package v1alpha1

import (
	"github.com/ghodss/yaml"
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
            etcdV2:
              type: object
              properties:
                enabled:
                  type: boolean
                dataDir:
                  type: string
            etcdV3:
              type: object
              properties:
                enabled:
                  type: boolean
                endpoints:
                  type: string
                cacert:
                  type: string
                cert:
                  type: string
                key:
                  type: string
          required:
          - guestBackup
          - etcdV2
          - etcdV3
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
              type: array
              items:
                type: object
                properties:
                  name:
                    type: string
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
                  latestError:
                    type: string
                required:
                - name
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
type ETCDBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ETCDBackupSpec   `json:"spec"`
	Status            ETCDBackupStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

type ETCDBackupSpec struct {
	// GuestBackup is a boolean indicating if the tenant clusters have to be backupped
	GuestBackup bool `json:"guestBackup" yaml:"guestBackup"`
	// ETCDv2 backup settings
	ETCDv2 ETCDv2Settings `json:"etcdV2" yaml:"etcdV2"`
	// ETCDv3 backup settings
	ETCDv3 ETCDv3Settings `json:"etcdV3" yaml:"etcdV3"`
}

type ETCDv2Settings struct {
	Enabled bool   `json:"enabled" yaml:"enabled"`
	DataDir string `json:"dataDir" yaml:"dataDir"`
}

type ETCDv3Settings struct {
	Enabled   bool   `json:"enabled" yaml:"enabled"`
	Endpoints string `json:"endpoints" yaml:"endpoints"`
	CaCert    string `json:"cacert" yaml:"cacert"`
	Cert      string `json:"cert" yaml:"cert"`
	Key       string `json:"key" yaml:"key"`
}

type ETCDBackupStatus struct {
	// array for the state of the backup for all instances
	Instances []ETCDInstanceBackupStatus `json:"instances" yaml:"instances"`
	// Status of the whole backup job (can be 'Pending', 'Running'. 'Completed', 'Failed')
	Status string `json:"status" yaml:"status"`
	// Timestamp when the first attempt was made
	StartedTimestamp DeepCopyTime `json:"startedTimestamp,omitempty" yaml:"startedTimestamp"`
	// Timestamp when the last (final) attempt was made (when the Phase became either 'Completed' or 'Failed'
	FinishedTimestamp DeepCopyTime `json:"finishedTimestamp,omitempty" yaml:"finishedTimestamp"`
}

type ETCDInstanceBackupStatus struct {
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
	// Latest backup error message
	LatestError string `json:"latestError,omitempty" yaml:"latestError,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ETCDBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ETCDBackup `json:"items"`
}
