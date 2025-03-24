package register

import (
	"crypto/rand"
	"errors"
	"log"
	"matcha/go/components"
	"matcha/go/database"
	"matcha/go/database/enums"
	"matcha/go/database/sqlc"
	"matcha/go/globals"
	"matcha/go/pages"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"vendor/golang.org/x/crypto/sha3"

	"github.com/a-h/templ"
	"github.com/pquerna/otp/totp"
	"goji.io/v3"
	"goji.io/v3/pat"
)

type registerForm struct {
	firstName   string
	lastName    string
	mail        string
	password    []byte
	birthDate   int64
	sex         enums.Sex
	orientation enums.Orientation
	country     int64
	city        int64
	bio         string
}

func (this *registerForm) parse(r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return errors.New("Invalid form data")
	}

	var input string
	namePat := regexp.MustCompile(`^[a-zA-Z\-]{2,16}`)

	if input = r.Form.Get("first-name"); input == "" {
		return errors.New("Missing first-name in form data")
	} else if !namePat.Match([]byte(input)) {
		return errors.New("First name contains illegal characters")
	} else {
		this.firstName = input
	}

	if input = r.Form.Get("last-name"); input == "" {
		return errors.New("Missing last-name in form data")
	} else if !namePat.Match([]byte(input)) {
		return errors.New("Last name contains illegal characters")
	} else {
		this.lastName = input
	}

	if input = r.Form.Get("mail"); input == "" {
		return errors.New("Missing mail in form data")
	} else if len(input) > 64 {
		return errors.New("Mail is too big")
	} else {
		this.mail = input
	}

	if input = r.Form.Get("password"); input == "" {
		return errors.New("Missing password in form data")
	} else if !isValidPassword(input) {
		return errors.New("Password does not conform to requirements")
	} else {
		hasher := sha3.New512()
		hasher.Write([]byte(input))
		this.password = hasher.Sum(nil)
	}

	if input = r.Form.Get("birth-date"); input == "" {
		return errors.New("Missing birth-date in form data")
	} else if birthDate, err := time.Parse("2006-01-02", input); err != nil {
		return errors.New("Invalid format for birth date")
	} else {
		if time.Now().Before(birthDate.AddDate(18, 0, 0)) {
			return errors.New("User is younger than 18 years old")
		}
		this.birthDate = birthDate.Unix()
	}

	if input = r.Form.Get("sex"); input == "" {
		return errors.New("Missing sex in form data")
	} else if sex, err := strconv.ParseInt(input, 10, 0); err != nil {
		return errors.New("Invalid value for sex (not a number)")
	} else {
		this.sex = enums.Sex(sex)
		if !enums.Sexes.Contains(this.sex) {
			return errors.New("Invalid value for sex (out of range)")
		}
	}

	if input = r.Form.Get("orientation"); input == "" {
		return errors.New("Missing orientation in form data")
	} else if orientation, err := strconv.ParseInt(input, 10, 0); err != nil {
		return errors.New("Invalid value for orientation (not a number)")
	} else {
		this.orientation = enums.Orientation(orientation)
		if !enums.Orientations.Contains(this.orientation) {
			return errors.New("Invalid value for orientation (out of range)")
		}
	}

	if input = r.Form.Get("country"); input == "" {
		return errors.New("Missing country in form data")
	} else if country, err := strconv.ParseInt(input, 10, 0); err != nil {
		return errors.New("Invalid value for country (not a number)")
	} else if country < 0 || country >= int64(len(globals.Countries)) {
		return errors.New("Invalid value for country (out of range)")
	} else {
		this.country = country
	}

	if input = r.Form.Get("city"); input == "" {
		return errors.New("Missing city in form data")
	} else if city, err := strconv.ParseInt(input, 10, 0); err != nil {
		return errors.New("Invalid value for city (not a number)")
	} else if city < 0 || city >= int64(len(globals.Cities)) {
		return errors.New("Invalid value for city (out of range)")
	} else if globals.Cities[city].Country != this.country {
		return errors.New("City doesn't match country")
	}

	if input = r.Form.Get("bio"); len(input) > 512 {
		return errors.New("Bio is too big")
	} else {
		this.bio = input
	}

	return nil
}

func isValidPassword(password string) bool {
	passwordBin := []byte(password)
	containsLower := regexp.MustCompile(`[a-z]`)
	containsUpper := regexp.MustCompile(`[A-Z]`)
	containsDigit := regexp.MustCompile(`[0-9]`)
	containsSpec := regexp.MustCompile(`[^a-zA-Z0-9]`)

	return len(passwordBin) >= 8 && len(passwordBin) <= 32 &&
		containsLower.Match(passwordBin) && containsUpper.Match(passwordBin) &&
		containsDigit.Match(passwordBin) && containsSpec.Match(passwordBin)
}

var Page = pages.NewPage(&pages.PageOptions{
	Path: "register",
	Setup: func(mux *goji.Mux) http.Handler {
		mux.HandleFunc(pat.Post("submit"), func(w http.ResponseWriter, r *http.Request) {
			var form registerForm
			if err := form.parse(r); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if id, err := database.Handle.CreateUser(r.Context(), sqlc.CreateUserParams{
				// TODO: fill with form
			}); err != nil {
				log.Fatalf("Failed to create user in database: %s\n", err.Error())
			} else if key, err := totp.Generate(totp.GenerateOpts{
				Issuer:      "matcha",
				AccountName: strconv.FormatInt(id, 10),
				Period:      300,
				Digits:      6,
				Rand:        rand.Reader,
			}); err != nil {
				log.Fatalf("Failed to generate one-time password for user %s: %s\n", id, err.Error())
			} else {
				// TODO: send mail to user and return the confirm component
			}
		})
		mux.Handle(pat.Post("country"), countryHandler)
		return templ.Handler(components.Page("Register", body()))
	},
})
