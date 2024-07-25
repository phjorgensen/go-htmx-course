package contact

type Contact struct {
	Id    int
	Name  string
	Email string
}

type Contacts = []Contact

type ContactListData struct {
	Contacts Contacts
}

func (listData ContactListData) hasEmail(email string) bool {
	for _, contact := range listData.Contacts {
		if contact.Email == email {
			return true
		}
	}

	return false
}

func (listData ContactListData) indexOf(id int) int {
	for i, contact := range listData.Contacts {
		if contact.Id == id {
			return i
		}
	}

	return -1
}

type ContactFormData struct {
	Values map[string]string
	Errors map[string]string
}

type ContactPage struct {
	List ContactListData
	Form ContactFormData
}
