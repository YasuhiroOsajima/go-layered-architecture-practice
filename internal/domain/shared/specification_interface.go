package shared

// Specification objects can be taken as function args by repository objects
// like complex read data SQL to avoid 1+N performance problem.
// In these cases, repository functions should take this interface object
// to avoid dependency to upper side layer.
