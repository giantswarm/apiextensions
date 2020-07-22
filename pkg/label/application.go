package label

// ApplicationCatalogType describes the type of catalog that the labelled
// AppCatalog CR represents. This label helps our UI determine how to
// render this app catalog. Possible values:
//
// - `"internal"`: Will not show up in our UIs at all.
// - `"managed"`: Will gain a 'managed' banner, helping it stand out from
//   other catalogs.
// - `"incubator"`: Will show some expectation management message before
//   installing an app from this catalog, that it is still a work in progress,
//   but expected to at least install and somewhat work.
// - `"community"`: Will show a more strongly worded expectation management
//   message, indicating that apps from this catalog will most likely not
//   work without some adjustments.
const ApplicationCatalogType = "application.giantswarm.io/catalog-type"
