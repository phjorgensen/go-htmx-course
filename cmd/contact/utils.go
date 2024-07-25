package contact

var lastContactId = 0

func NewContact(name string, email string) Contact {
	lastContactId++
	return Contact{
		Id:    lastContactId,
		Name:  name,
		Email: email,
	}
}

func NewFormData() ContactFormData {
	return ContactFormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

func NewListData() ContactListData {
	return ContactListData{
		Contacts: []Contact{
			NewContact("John", "jd@gmail.com"),
			NewContact("Clara", "cd@gmail.com"),
		},
	}
}

func NewPage() ContactPage {
	return ContactPage{
		List: NewListData(),
		Form: NewFormData(),
	}
}
