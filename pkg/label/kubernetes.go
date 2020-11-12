package label

// App label is used to identify Kubernetes resources. It is considered
// deprecated by upstream and is replaced by `app.kubernetes.io/name`.
const App = "app"

// AppKubernetesInstance is a unique name identifying an instance of an app.
const AppKubernetesInstance = "app.kubernetes.io/instance"

// AppKubernetesName label is used to identify Kubernetes resources.
const AppKubernetesName = "app.kubernetes.io/name"

// AppKubernetesVersion label is used to identify the version of Kubernetes
// resources.
const AppKubernetesVersion = "app.kubernetes.io/version"
