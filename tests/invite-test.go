package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"net/http"

)

func TestValidateInviteToken(t *testing.T) { // GET /validate
	assert := assert.New(t)
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/validate", nil)
    router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestRetrieveInviteTokens(t *testing.T) { // GET /invite-token

}

func TestGenerateInviteToken(t *testing.T) { // POST /invite-token
	// Connect to Test Database
	// Add test user and passwor into the test database
	// Check if JWT is created.
	// Clear database entry
}

func TestUpdateInviteToken(t *testing.T) { // PUT /invite-token
	
}

func TestDeleteInviteToken(t *testing.T) { // DELETE /invite-token
	
}