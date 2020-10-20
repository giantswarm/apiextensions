package annotation

// ConditionsToSet is used to add conditions array (in JSON format) that will
// be added to CR status in the first next reconciliation loop. The handler
// that reads this annotation and sets the conditions also MUST update the
// annotation by removing the conditions that are already set in the CR status.
const ConditionsToSet = "giantswarm.io/conditions-to-set"
