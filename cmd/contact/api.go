package contact

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateContactApi(e *echo.Echo) {
	contactPage := NewPage()

	e.GET("/contacts", func(c echo.Context) error {
		return c.Render(http.StatusOK, "contact_page", contactPage)
	})

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")

		if contactPage.List.hasEmail(email) {
			formData := NewFormData()
			formData.Values["name"] = name
			formData.Values["email"] = email
			formData.Errors["email"] = "Email already exists"

			return c.Render(http.StatusUnprocessableEntity, "form", formData)
		}

		contact := NewContact(name, email)
		contactPage.List.Contacts = append(contactPage.List.Contacts, contact)

		c.Render(http.StatusOK, "form", NewFormData())
		return c.Render(http.StatusOK, "oob-contact", contact)
	})

	e.DELETE("/contacts/:id", func(c echo.Context) error {
		time.Sleep(3 * time.Second)
		idStr := c.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid id")
		}

		index := contactPage.List.indexOf(id)
		if index == -1 {
			return c.String(http.StatusNotFound, "Contact not found")
		}

		contactPage.List.Contacts = append(contactPage.List.Contacts[:index], contactPage.List.Contacts[index+1:]...)

		return c.NoContent(http.StatusOK)
	})
}
