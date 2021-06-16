package excp

import "net/http"

var (
	InternErrorEx     = NewType("ex.generic.intern-error", http.StatusInternalServerError)
	EntityNotFoundEx  = NewType("ex.generic.entity-not-found", http.StatusNoContent)
	NotResultsFoundEx = NewType("ex.generic.no-results-found", http.StatusNoContent)
	KvsOverQuotaEx    = NewType("ex.generic.kvs-over-quota", http.StatusInternalServerError)
)

type Type interface {
	Catch(exception Exception) bool
	GetCode() int
}

type TypeStruct struct {
	key  string
	code int
}

func NewType(key string, code int) TypeStruct {
	return TypeStruct{
		key:  key,
		code: code,
	}
}

func (model TypeStruct) Catch(exception Exception) bool {
	if exception == nil {
		return false
	}
	return model == exception.GetType()
}

func (model TypeStruct) GetCode() int {
	return model.code
}
