package config

import (
    "os"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
    cfg := New()
    assert.NotNil(t, cfg, "returned config shouldn't be nil")
}

func TestEnvironments(t *testing.T) {
    originalEnv := os.Getenv(environmentENV)
    defer func() {
        err := os.Setenv(environmentENV, originalEnv)
        require.NoError(t, err)
    }()

    envs := []string{"development", "test"}

    for _, env := range envs {
        err := os.Setenv(environmentENV, env)
        require.NoError(t, err)

        cfg := New()
        assert.Equal(t, cfg.Environment, env, "incorrect environment")
    }
}
