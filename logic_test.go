package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogic_WrongArg(t *testing.T) {
	cmd := "шарик"

	err := f(cmd)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Неверная команда")
}
