package error_test

import (
	error2 "github.com/ruancaetano/go-gin-fruits/internal/presentation/error"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHttpError_Error(t *testing.T) {
	err := error2.HttpError{
		Message: "invalid params",
		Status:  400,
	}

	assert.Equal(t, err.Error(), "invalid params")
}
