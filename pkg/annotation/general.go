// Package annotation defines annotation keys used by Giant Swarm within
// Kubernetes resources.
package annotation

// Docs is the docs annotation put into all CRs to link to its CR specific
// documentation. This aims to help understanding all the moving parts within
// the system and how they relate to each other.
const Docs = "giantswarm.io/docs"

// Notes is for informational messages for resources managed by operators. Such
// as whether the resource may or may not be edited.
const Notes = "giantswarm.io/notes"

// ReleaseNotesURL defines where to find release notes about the CR at hand.
// The value is expected to be a URI, e. g.
// "https://github.com/giantswarm/releases/tree/master/aws/v11.5.0".
const ReleaseNotesURL = "giantswarm.io/release-notes"
