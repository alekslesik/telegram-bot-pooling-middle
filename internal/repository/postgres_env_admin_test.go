package repository

import "testing"

func TestIsEnvAdminAdminTelegramIDs(t *testing.T) {
	t.Setenv("ADMIN_TELEGRAM_IDS", "12345, 777 ; 42")
	t.Setenv("ADMIN_IDS", "")
	if !isEnvAdmin(777) {
		t.Fatal("expected user 777 to be admin from ADMIN_TELEGRAM_IDS")
	}
	if isEnvAdmin(999) {
		t.Fatal("did not expect user 999 to be admin")
	}
}

func TestIsEnvAdminFallbackAdminIDs(t *testing.T) {
	t.Setenv("ADMIN_TELEGRAM_IDS", "")
	t.Setenv("ADMIN_IDS", "1001 1002")
	if !isEnvAdmin(1002) {
		t.Fatal("expected user 1002 to be admin from ADMIN_IDS")
	}
}

