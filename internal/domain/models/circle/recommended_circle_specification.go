package circle

// Specification object expresses a domain rule for domain objects.
// Specification object takes domain objects in itself's arguments.
//
// And if domain object needs repository, the process should be write in specification method in each 'specification' file.
// Entities and value objects should not contain processes which uses repository.

func IsRecommended(circle *Circle) bool {
	if circle.CountMembers() < 10 {
		return false
	}

	return true
}
