package whatsmeow_service

import (
	"testing"

	"go.mau.fi/whatsmeow/store"
)

func TestInitWhatsAppVersion(t *testing.T) {
	v := InitWhatsAppVersion(nil, nil)
	t.Logf("Initialized version: %+v", v)

	currentWAVersion := store.GetWAVersion()
	t.Logf("store.GetWAVersion(): %s", currentWAVersion.String())

	if currentWAVersion.String() == "2.3000.1042386815" {
		t.Fatalf("Expected store WAVersion to be updated, but it is still the outdated default 2.3000.1042386815")
	}
}
