package excp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/fury_discount-api/src/api/infrastructure/i18n"
	"github.com/mercadolibre/fury_discount-api/src/api/infrastructure/logger"
)

type Exception interface {
	GetCode() int
	GetType() Type
	GetCause() error
	GetFullMessage() string
}

type exceptionImpl struct {
	Key     string `json:"key"`
	Message string `json:"message"`
	model   Type
	err     error
}

func New(ctx context.Context, model TypeStruct, params ...interface{}) exceptionImpl {
	message := i18n.Translate(ctx, model.key, params...)
	return exceptionImpl{
		Key:     model.key,
		Message: message,
		model:   model,
	}
}

func Propagated(code int, jsonBytes []byte) (exceptionImpl, bool) {
	var propagated exceptionImpl
	if err := json.Unmarshal(jsonBytes, &propagated); err == nil && propagated.Key != "" {
		propagated.model = TypeStruct{
			code: code,
			key:  propagated.Key,
		}
		return propagated, true
	}
	return propagated, false
}

func FromError(ctx context.Context, model TypeStruct, err error, log *logger.Logger, params ...interface{}) exceptionImpl {
	log.WithError(err).Error(ctx)

	exception := New(ctx, model, params...)
	exception.err = err
	return exception
}

func (ex exceptionImpl) GetCode() int {
	return ex.model.GetCode()
}

func (ex exceptionImpl) GetType() Type {
	return ex.model
}

func (ex exceptionImpl) GetCause() error {
	return ex.err
}

func (ex exceptionImpl) GetFullMessage() string {
	fullTraceMessage := ex.Message
	if ex.err != nil {
		trace := fmt.Sprintf("%+v\n", ex.err)
		fullTraceMessage = fullTraceMessage + " - Cause: \n" + trace
	}
	return fullTraceMessage
}
