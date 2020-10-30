package label

// CatalogName label is used to identify resources belonging to a Giant Swarm
// app catalog.
const CatalogName = "application.giantswarm.io/catalog"

// CatalogType label is used to identify the type of Giant Swarm app catalog
// e.g. stable or test.
const CatalogType = "application.giantswarm.io/catalog-type"

// CatalogVisibility label is used to determine how Giant Swarm app catalogs
// are displayed in the UX. e.g. public or internal.
const CatalogVisibility = "application.giantswarm.io/catalog-visibility"
