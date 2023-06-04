package types

type UserDoc struct {
	Email     string `firestore:"email"`
	FirstName string `firestore:"firstName"`
	LastName  string `firestore:"lastName"`
}

func NewUserDoc(email string, firstName string, lastName string) *UserDoc {
	return &UserDoc{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}
}
