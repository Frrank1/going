package wait_test

import (
	"fmt"
	"testing"

	. "github.com/markbates/going/wait"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Name string
}

type Users []User

func Test_Array(t *testing.T) {
	a := assert.New(t)

	users := []User{User{"A"}, User{"B"}}
	Wait(len(users), func(index int) {
		user := users[index]
		user.Name = fmt.Sprintf("User: %d", index)
		users[index] = user
	})

	a.Equal(users[0].Name, "User: 0")
	a.Equal(users[1].Name, "User: 1")
}

func Test_Collection(t *testing.T) {
	a := assert.New(t)

	users := Users{User{"A"}, User{"B"}}
	Wait(len(users), func(index int) {
		user := users[index]
		user.Name = fmt.Sprintf("User: %d", index)
		users[index] = user
	})

	a.Equal(users[0].Name, "User: 0")
	a.Equal(users[1].Name, "User: 1")
}

// var w sync.WaitGroup
// w.Add(len(users))
// for i := 0; i < len(users); i++ {
// 	user := users[i]
// 	go func(w *sync.WaitGroup, user *models.User, i int) {
// 		// models.DecorateUser(tx, user)
// 		user.AfterFind(tx)
// 		models.DecorateWithPermissions(context.CurrentUser, user)
// 		users[i] = *user
// 		w.Done()
// 	}(&w, &user, i)
// }
// w.Wait()
