package v1alpha2

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1alpha2 "sigs.k8s.io/cluster-api/api/v1alpha2"

	"github.com/giantswarm/apiextensions/pkg/id"
	"github.com/giantswarm/apiextensions/pkg/label"
)

const (
	defaultMasterInstanceType = "m5.xlarge"
)

type ClusterConfig struct {
	ClusterID         string
	ControlPlaneID    string
	Credential        string
	Domain            string
	ExternalSNAT      bool
	MasterAZ          []string
	Name              string
	PodsCIDR          string
	Owner             string
	Region            string
	ReleaseComponents map[string]string
	ReleaseVersion    string
	Labels            map[string]string
}

type ClusterCRs struct {
	Cluster         *apiv1alpha2.Cluster
	AWSCluster      *AWSCluster
	G8sControlPlane *G8sControlPlane
	AWSControlPlane *AWSControlPlane
}

func NewClusterCRs(config ClusterConfig) (ClusterCRs, error) {
	// Default some essentials in case certain information are not given. E.g.
	// the Tenant Cluster ID may be provided by the user.
	{
		if config.ClusterID == "" {
			config.ClusterID = id.Generate()
		}
		if config.ControlPlaneID == "" {
			config.ControlPlaneID = id.Generate()
		}
	}

	awsClusterCR := newAWSClusterCR(config)
	clusterCR := newClusterCR(awsClusterCR, config)
	awsControlPlaneCR := newAWSControlPlaneCR(config)
	g8sControlPlaneCR := newG8sControlPlaneCR(awsControlPlaneCR, config)

	crs := ClusterCRs{
		Cluster:         clusterCR,
		AWSCluster:      awsClusterCR,
		G8sControlPlane: g8sControlPlaneCR,
		AWSControlPlane: awsControlPlaneCR,
	}

	return crs, nil
}

func newAWSClusterCR(c ClusterConfig) *AWSCluster {
	awsClusterCR := &AWSCluster{
		TypeMeta: metav1.TypeMeta{
			Kind:       kindAWSCluster,
			APIVersion: SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      c.ClusterID,
			Namespace: metav1.NamespaceDefault,
			Labels: map[string]string{
				label.AWSOperatorVersion: c.ReleaseComponents["aws-operator"],
				label.Cluster:            c.ClusterID,
				label.Organization:       c.Owner,
				label.ReleaseVersion:     c.ReleaseVersion,
			},
		},
		Spec: AWSClusterSpec{
			Cluster: AWSClusterSpecCluster{
				Description: c.Name,
				DNS: AWSClusterSpecClusterDNS{
					Domain: c.Domain,
				},
				OIDC: AWSClusterSpecClusterOIDC{},
			},
			Provider: AWSClusterSpecProvider{
				CredentialSecret: AWSClusterSpecProviderCredentialSecret{
					Name:      c.Credential,
					Namespace: "giantswarm",
				},
				Pods: AWSClusterSpecProviderPods{
					CIDRBlock:    c.PodsCIDR,
					ExternalSNAT: &c.ExternalSNAT,
				},
				Region: c.Region,
			},
		},
	}

	// Single master node
	if len(c.MasterAZ) == 1 {
		awsClusterCR.Spec.Provider.Master = AWSClusterSpecProviderMaster{
			AvailabilityZone: c.MasterAZ[0],
			InstanceType:     defaultMasterInstanceType,
		}
	}

	return awsClusterCR
}

func newAWSControlPlaneCR(c ClusterConfig) *AWSControlPlane {
	return &AWSControlPlane{
		TypeMeta: metav1.TypeMeta{
			Kind:       kindAWSControlPlane,
			APIVersion: SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      c.ControlPlaneID,
			Namespace: metav1.NamespaceDefault,
			Labels: map[string]string{
				label.AWSOperatorVersion: c.ReleaseComponents["aws-operator"],
				label.Cluster:            c.ClusterID,
				label.ControlPlane:       c.ControlPlaneID,
				label.Organization:       c.Owner,
				label.ReleaseVersion:     c.ReleaseVersion,
			},
		},
		Spec: AWSControlPlaneSpec{
			AvailabilityZones: c.MasterAZ,
			InstanceType:      defaultMasterInstanceType,
		},
	}
}

func newClusterCR(obj *AWSCluster, c ClusterConfig) *apiv1alpha2.Cluster {
	clusterLabels := map[string]string{}
	{
		for key, value := range c.Labels {
			clusterLabels[key] = value
		}

		gsLabels := map[string]string{
			label.ClusterOperatorVersion: c.ReleaseComponents["cluster-operator"],
			label.Cluster:                c.ClusterID,
			label.Organization:           c.Owner,
			label.ReleaseVersion:         c.ReleaseVersion,
		}

		for key, value := range gsLabels {
			clusterLabels[key] = value
		}
	}

	clusterCR := &apiv1alpha2.Cluster{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Cluster",
			APIVersion: "cluster.x-k8s.io/v1alpha2",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      c.ClusterID,
			Namespace: metav1.NamespaceDefault,
			Labels:    clusterLabels,
		},
		Spec: apiv1alpha2.ClusterSpec{
			InfrastructureRef: &corev1.ObjectReference{
				APIVersion: obj.TypeMeta.APIVersion,
				Kind:       obj.TypeMeta.Kind,
				Name:       obj.GetName(),
				Namespace:  obj.GetNamespace(),
			},
		},
	}

	return clusterCR
}

func newG8sControlPlaneCR(obj *AWSControlPlane, c ClusterConfig) *G8sControlPlane {
	return &G8sControlPlane{
		TypeMeta: metav1.TypeMeta{
			Kind:       kindG8sControlPlane,
			APIVersion: SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      c.ControlPlaneID,
			Namespace: metav1.NamespaceDefault,
			Labels: map[string]string{
				label.ClusterOperatorVersion: c.ReleaseComponents["cluster-operator"],
				label.Cluster:                c.ClusterID,
				label.ControlPlane:           c.ControlPlaneID,
				label.Organization:           c.Owner,
				label.ReleaseVersion:         c.ReleaseVersion,
			},
		},
		Spec: G8sControlPlaneSpec{
			Replicas: len(c.MasterAZ),
			InfrastructureRef: corev1.ObjectReference{
				APIVersion: obj.TypeMeta.APIVersion,
				Kind:       obj.TypeMeta.Kind,
				Name:       obj.GetName(),
				Namespace:  obj.GetNamespace(),
			},
		},
	}
}
