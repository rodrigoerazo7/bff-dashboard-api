package domain

type User struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
}

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

// Status returns the user's classification based on age.
// Veteran: age > 50. Rookie: otherwise.
func (u User) Status() string {
	if u.Age > 50 {
		return "Veteran"
	}
	return "Rookie"
}
