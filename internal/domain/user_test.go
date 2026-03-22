package domain

import "testing"

func TestUser_FullName(t *testing.T) {
	u := User{FirstName: "Terry", LastName: "Medhurst"}
	if got := u.FullName(); got != "Terry Medhurst" {
		t.Errorf("FullName(): want 'Terry Medhurst', got %q", got)
	}
}

func TestUser_Status_Veteran(t *testing.T) {
	u := User{Age: 51}
	if got := u.Status(); got != "Veteran" {
		t.Errorf("Status(): want Veteran, got %q", got)
	}
}

func TestUser_Status_Boundary(t *testing.T) {
	// age == 50 is NOT veteran
	u := User{Age: 50}
	if got := u.Status(); got != "Rookie" {
		t.Errorf("Status(): age 50 should be Rookie, got %q", got)
	}
}

func TestUser_Status_Rookie(t *testing.T) {
	u := User{Age: 25}
	if got := u.Status(); got != "Rookie" {
		t.Errorf("Status(): want Rookie, got %q", got)
	}
}
