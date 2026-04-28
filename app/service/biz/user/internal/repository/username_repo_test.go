package repository

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/user/internal/repository/model"
)

type usernameModelNotFound struct {
	model.UsernameModel
}

func (usernameModelNotFound) SelectByUsername(context.Context, string) (*model.Username, error) {
	return nil, &model.NotFoundError{Resource: "username", Key: "username=teamgram"}
}

func TestIsValidUsername(t *testing.T) {
	valid := []string{"teamgram", "team_gram", "abc12", "a1_b2"}
	for _, username := range valid {
		if !isValidUsername(username) {
			t.Fatalf("expected valid username %q", username)
		}
	}

	invalid := []string{"", "ab", "12345", "_abcde", "abcde_", "a__bcde", "UPPER", "with-dash", "with.dot"}
	for _, username := range invalid {
		if isValidUsername(username) {
			t.Fatalf("expected invalid username %q", username)
		}
	}
}

func TestCheckUsernameTreatsModelNotFoundAsAvailable(t *testing.T) {
	r := &Repository{model: &model.Models{UsernameModel: usernameModelNotFound{}}}
	got, err := r.CheckUsername(context.Background(), "teamgram")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := got.ToUsernameNotExisted(); !ok {
		t.Fatalf("expected username not existed, got %#v", got)
	}
}
