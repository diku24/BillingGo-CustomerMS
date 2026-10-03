package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindConfigFileWalksParentDirectories(t *testing.T) {
	root := t.TempDir()
	configDir := filepath.Join(root, "config")
	if err := os.Mkdir(configDir, 0o700); err != nil {
		t.Fatal(err)
	}
	configPath := filepath.Join(configDir, "configuration.env")
	if err := os.WriteFile(configPath, []byte("SERVERPORT=:8382\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	start := filepath.Join(root, "handler")
	if err := os.Mkdir(start, 0o700); err != nil {
		t.Fatal(err)
	}

	got, err := findConfigFile(start)
	if err != nil {
		t.Fatal(err)
	}
	if got != configPath {
		t.Fatalf("findConfigFile() = %q, want %q", got, configPath)
	}
}

func TestEnvVarReadUsesFileValueWithoutMutatingEnvironment(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "configuration.env")
	if err := os.WriteFile(path, []byte("SERVERPORT=:9999\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	t.Setenv("BILLINGGO_CONFIG_FILE", path)
	t.Setenv("SERVERPORT", "stale-value")

	got := EnvVarRead("SERVERPORT")
	if got != ":9999" {
		t.Fatalf("EnvVarRead() = %q, want %q", got, ":9999")
	}
	if os.Getenv("SERVERPORT") != "stale-value" {
		t.Fatal("EnvVarRead mutated the process environment")
	}
}
