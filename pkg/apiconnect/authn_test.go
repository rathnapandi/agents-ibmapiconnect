package apiconnect

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	tests := []struct {
		name  string
		token string
		err   error
	}{
		{
			name:  "should return no token and an error",
			token: "",
			err:   fmt.Errorf("no token here"),
		},
	}
	for i := range tests {
		tc := tests[i]
		t.Run(tc.name, func(t *testing.T) {
			c := &authClient{}
			user := &Credential{}

			c.On("GetAccessToken").Return(tc.token, user, time.Duration(10), tc.err)

			auth, err := NewAuth(c)
			assert.Equal(t, tc.err, err)
			if tc.err == nil {
				assert.NotNil(t, auth)
				token := auth.GetToken()
				assert.Equal(t, tc.token, token)
				auth.Stop()

			} else {
				assert.Nil(t, auth)
			}

		})
	}

}

func Test_startRefreshToken(t *testing.T) {

	client := &authClientRefreshErr{
		stop: make(chan bool),
	}
	a := &auth{
		client: client,
	}
	a.startRefreshToken(1000)
	done := <-client.stop
	assert.True(t, done)
}

type authClientRefreshErr struct {
	stop chan bool
}

func (a authClientRefreshErr) GetAccessToken() (string, *Credential, time.Duration, error) {
	a.stop <- true
	return "", &Credential{}, 0, fmt.Errorf("auth error")
}

type authClient struct {
	mock.Mock
}

func (a *authClient) GetAccessToken() (string, *Credential, time.Duration, error) {
	args := a.Called()
	token := args.String(0)
	user := args.Get(1)
	duration := args.Get(2)
	return token, user.(*Credential), duration.(time.Duration), args.Error(3)
}
