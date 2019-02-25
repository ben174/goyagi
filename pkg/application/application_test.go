package application

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
    app, err := New()
    assert.Nil(t, err)
    assert.NotNil(t, app)
    assert.NotNil(t, app.Config)
}
