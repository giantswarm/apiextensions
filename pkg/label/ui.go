package label

// I the ui.giantswarm.io label namespace we define labels that
// exist in order to affect the user experience in user
// interfaces, like the Giant Swarm web UI or CLIs.

// Affects whether or not a resource is intended for display in
// a user interface. For example, it can be used to hide
// irrelevant system roles from users by default. The value can
// either be "true" or "false".
const DisplayInUserInterface = "ui.giantswarm.io/display"
