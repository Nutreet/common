package common

import (
	"context"
	"testing"
)

func TestGetApp(t *testing.T) {
	app := GetApp()
	if app == nil {
		t.Errorf("app is nil\n")
	}
}

func TestFirestore(t *testing.T) {
	Firebase := GetApp()

	if Firebase == nil {
		t.Errorf("app is nil\n")
	}

	firestore, err := Firebase.Firestore(context.Background())
	if err != nil {
		t.Errorf("expected to retrieve firestore with no error")
	}

	if firestore == nil {
		t.Errorf("expected to retrive firestore client, got nil")
	}

	firestore.Close()
}

func TestStorage(t *testing.T) {
	Firebase := GetApp()

	if Firebase == nil {
		t.Errorf("app is nil\n")
	}

	storage, err := Firebase.Storage(context.Background())
	if err != nil {
		t.Errorf("expected to retrieve storage with no error")
	}

	if storage == nil {
		t.Errorf("expected to retrive storage client, got nil")
	}
}

func TestAuth(t *testing.T) {
	Firebase := GetApp()

	if Firebase == nil {
		t.Errorf("app is nil\n")
	}

	auth, err := Firebase.Auth(context.Background())
	if err != nil {
		t.Errorf("expected to retrieve auth with no error")
	}

	if auth == nil {
		t.Errorf("expected to retrive auth client, got nil")
	}
}
