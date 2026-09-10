package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogic_WrongArgs(t *testing.T) {
	cases := []struct {
		name string
		arg  string
	}{
		{
			name: "wrong argument",
			arg:  "шарик",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := f(tc.arg)
			require.Error(t, err)
			require.Contains(t, err.Error(), "Неверная команда")
		})
	}

}
