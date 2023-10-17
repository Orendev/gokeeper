package account

import (
	"os"
	"testing"

	"github.com/Orendev/gokeeper/internal/pkg/domain/account"
	"github.com/Orendev/gokeeper/pkg/tools/encryption"
	"github.com/Orendev/gokeeper/pkg/type/title"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	ac                *account.Account
	acEnc             *account.Account
	enc               *encryption.Enc
	key               = "supersecret"
	updateAccountArgs = &UpdateAccountArgs{}
	createAccountArgs = &CreateAccountArgs{}
	titleError        string
)

func TestMain(m *testing.M) {
	titleObj, _ := title.New("Test title")
	userID := uuid.New()
	login := []byte("test login")
	password := []byte("test pasword")
	url := []byte("ya.ru")
	comment := []byte("test comment")

	ac, _ = account.New(
		userID,
		*titleObj,
		login,
		password,
		url,
		comment,
	)

	updateAccountArgs.ID = ac.ID().String()
	updateAccountArgs.Login = string(ac.Login())
	updateAccountArgs.Password = string(ac.Password())
	updateAccountArgs.Title = ac.Title().String()
	updateAccountArgs.Comment = string(ac.Comment())
	updateAccountArgs.URL = string(ac.URL())
	updateAccountArgs.UserID = ac.UserID().String()

	createAccountArgs.Login = string(ac.Login())
	createAccountArgs.Password = string(ac.Password())
	createAccountArgs.Title = ac.Title().String()
	createAccountArgs.Comment = string(ac.Comment())
	createAccountArgs.URL = string(ac.URL())
	createAccountArgs.UserID = ac.UserID().String()

	titleError = "Банальные, но неопровержимые выводы, а также ключевые особенности структуры проекта " +
		"формируют глобальную экономическую сеть и при этом — описаны максимально подробно. Как уже неоднократно упомянуто, " +
		"некоторые особенности внутренней политики, превозмогая сложившуюся непростую экономическую ситуацию, объявлены нарушающими " +
		"общечеловеческие нормы этики и морали."

	enc = encryption.New(key)

	os.Exit(m.Run())
}

func TestToDecAccount(t *testing.T) {

	assertion := assert.New(t)

	t.Run("positive test To Enc Update Account", func(t *testing.T) {
		result, err := ToEncUpdateAccount(enc, updateAccountArgs)
		assertion.NoError(err)
		acEnc = result
	})

	t.Run("positive test To Dec Account", func(t *testing.T) {

		result, err := ToDecAccount(enc, acEnc)
		assertion.NoError(err)

		assertion.Equal(result.ID(), ac.ID())
		assertion.Equal(result.URL(), ac.URL())
		assertion.Equal(result.Login(), ac.Login())
		assertion.Equal(result.Password(), ac.Password())
		assertion.Equal(result.Comment(), ac.Comment())

	})

	t.Run("positive test To Enc Create Account", func(t *testing.T) {
		result, err := ToEncCreateAccount(enc, createAccountArgs)
		assertion.NoError(err)
		acEnc = result
	})

	t.Run("positive test To Dec Account", func(t *testing.T) {

		result, err := ToDecAccount(enc, acEnc)
		assertion.NoError(err)

		assertion.Equal(result.URL(), ac.URL())
		assertion.Equal(result.Login(), ac.Login())
		assertion.Equal(result.Password(), ac.Password())
		assertion.Equal(result.Comment(), ac.Comment())

	})

	t.Run("negative test To Enc Update Account", func(t *testing.T) {
		updateAccountArgs.Title = titleError

		result, err := ToEncUpdateAccount(enc, updateAccountArgs)
		assertion.Errorf(err, title.ErrWrongLength.Error())
		assertion.Nil(result)
	})

	t.Run("negative test To Enc Create Account", func(t *testing.T) {
		createAccountArgs.Title = titleError

		result, err := ToEncCreateAccount(enc, createAccountArgs)
		assertion.Errorf(err, title.ErrWrongLength.Error())
		assertion.Nil(result)
	})
}
