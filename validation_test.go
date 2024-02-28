package golang_validation

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestValidation(t *testing.T) {
	validate := validator.New()
	if validate == nil {
		t.Error("validate is nill")
	}
}

func TestValidationVariable(t *testing.T) {
	validation := validator.New()

	user := ""
	err := validation.Var(user, "required")
	if err != nil {
		fmt.Print(err.Error())
	}
}

func TestValidateTwoVariable(t *testing.T) {
	validate := validator.New()

	password := "rahasia"
	confirmPassword := "mautahuaja"
	err := validate.VarWithValue(password, confirmPassword, "eqfield")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMultipleTag(t *testing.T) {
	validation := validator.New()

	user := "roni1234"
	err := validation.Var(user, "required,number")
	if err != nil {
		fmt.Print(err.Error())
	}
}

func TestTagParameter(t *testing.T) {
	validation := validator.New()

	user := "99"
	err := validation.Var(user, "required,numeric,min=5,max=10")
	if err != nil {
		fmt.Print(err.Error())
	}
}

func TestStruct(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,min=5,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()
	login := LoginRequest{
		Password: "rahasia",
		Username: "ahmadroni@gmail.com",
	}

	err := validate.Struct(login)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationError(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,min=5,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()
	login := LoginRequest{
		Password: "roni",
		Username: "roni",
	}

	err := validate.Struct(login)
	if err != nil {
		var validationError validator.ValidationErrors
		errors.As(err, &validationError)
		for _, field := range validationError {
			fmt.Println("error ", field.Field(), " on tag ", field.Tag(), " with error ", field.Error())
		}
	}
}

func TestStructCrossField(t *testing.T) {
	type LoginRequest struct {
		Username        string `validate:"required,min=5,email"`
		Password        string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
	}

	validate := validator.New()
	login := LoginRequest{
		Password:        "rahasia",
		ConfirmPassword: "123",
		Username:        "ahmadroni@gmail.com",
	}

	err := validate.Struct(login)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestNestedStruct(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      uint64  `validate:"required"`
		Name    string  `validate:"required"`
		Address Address `validate:"required"`
	}

	user := User{
		Id:   0,
		Name: "",
		Address: Address{
			City:    "",
			Country: "",
		},
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      uint64    `validate:"required"`
		Name    string    `validate:"required"`
		Address []Address `validate:"required,dive"`
	}

	user := User{
		Id:   0,
		Name: "",
		Address: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestBasicCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      uint64    `validate:"required"`
		Name    string    `validate:"required"`
		Address []Address `validate:"required,dive"`
		Hobbies []string  `validate:"required,dive,required,min=1"`
	}

	user := User{
		Id:   0,
		Name: "",
		Address: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
		Hobbies: []string{"Coding"},
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMap(t *testing.T) {
	type School struct {
		Name string `validate:"required"`
	}

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      uint64            `validate:"required"`
		Name    string            `validate:"required"`
		Address []Address         `validate:"required,dive"`
		Hobbies []string          `validate:"required,dive,required,min=1"`
		Schools map[string]School `validate:"required,dive,keys,required,min=2,endkeys"`
	}

	user := User{
		Id:   0,
		Name: "",
		Address: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
		Hobbies: []string{"Coding"},
		Schools: map[string]School{
			"SD": {
				Name: "SD 1 Indonesia",
			},
			"SMP": {
				Name: "",
			},
			"": {
				Name: "",
			},
		},
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestBasicMap(t *testing.T) {
	type School struct {
		Name string `validate:"required"`
	}

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      uint64            `validate:"required"`
		Name    string            `validate:"required"`
		Address []Address         `validate:"required,dive"`
		Hobbies []string          `validate:"required,dive,required,min=1"`
		Schools map[string]School `validate:"required,dive,keys,required,min=2,endkeys"`
		Wallet  map[string]int    `validate:"required,dive,keys,required,endkeys,required,gt=0"`
	}

	user := User{
		Id:   0,
		Name: "",
		Address: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
		Hobbies: []string{"Coding"},
		Schools: map[string]School{
			"SD": {
				Name: "SD 1 Indonesia",
			},
			"SMP": {
				Name: "",
			},
			"": {
				Name: "",
			},
		},
		Wallet: map[string]int{
			"BNI": 1,
			"BCA": 3,
		},
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestAlias(t *testing.T) {
	validate := validator.New()
	validate.RegisterAlias("varchar", "required,max=255")

	type Seller struct {
		Id     string `validate:"varchar"`
		Name   string `validate:"varchar"`
		Owner  string `validate:"varchar"`
		Slogan string `validate:"varchar"`
	}

	seller := Seller{
		Id:     "",
		Name:   "",
		Owner:  "",
		Slogan: "",
	}

	err := validate.Struct(seller)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func MustValidUsername(level validator.FieldLevel) bool {
	value, ok := level.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value) {
			return false
		}

		if len(value) < 5 {
			return false
		}
	}
	return true
}

func TestCustomerValidation(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("username", MustValidUsername)
	if err != nil {
		return
	}

	type LoginRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required"`
	}

	login := LoginRequest{
		Username: "roni",
		Password: "roni",
	}

	err = validate.Struct(login)
	if err != nil {
		fmt.Println(err)
	}
}

var regexNumber = regexp.MustCompile("^[0-9]+$")

func MustValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}

	value := field.Field().String()
	if !regexNumber.MatchString(value) {
		return false
	}

	return len(value) == length
}

func TestCustomeValidationParameter(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", MustValidUsername)
	validate.RegisterValidation("pin", MustValidPin)

	type LoginRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required,pin=5"`
	}

	login := LoginRequest{
		Username: "RONIPUR",
		Password: "1234",
	}

	err := validate.Struct(login)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestOrRule(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email|numeric"`
		Password string `validate:"required"`
	}

	validate := validator.New()
	login := LoginRequest{
		Username: "ahmadroni@gmail.com",
		Password: "",
	}

	err := validate.Struct(login)
	if err != nil {
		fmt.Println(err)
	}
}

func MustEqualIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		panic("field not ok")
	}

	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestCrossValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("field_equal_ignore_case", MustEqualIgnoreCase)

	type User struct {
		Username string `validate:"required,field_equal_ignore_case=Email|field_equal_ignore_case=Phone"`
		Email    string `validate:"required,email"`
		Phone    string `validate:"required,numeric"`
		Name     string `validate:"required"`
	}

	user := User{
		Username: "roni@gmail.com",
		Email:    "roni@gmail.com",
		Phone:    "081321321",
		Name:     "roni",
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err)
	}

	user = User{
		Username: "081321321",
		Email:    "roni@gmail.com",
		Phone:    "081321321",
		Name:     "roni",
	}

	err = validate.Struct(user)
	if err != nil {
		fmt.Println(err)
	}
}

type ReqisterRequest struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,numeric"`
	Password string `validate:"required"`
}

func MustValidRegisterSuccess(level validator.StructLevel) {
	registerRequest := level.Current().Interface().(ReqisterRequest)
	if registerRequest.Email == registerRequest.Username || registerRequest.Phone == registerRequest.Username {

	} else {
		// gagal
		level.ReportError(registerRequest.Username, "Username", "Username", "username", "")
	}
}
