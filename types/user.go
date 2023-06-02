package types

type UserDoc struct {
	Email     string `firestore:"email"`
	FirstName string `firestore:"firstName"`
	LastName  string `firestore:"lastName"`
}
