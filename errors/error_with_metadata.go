package errors

import (
	"fmt"

	"github.com/PlanckProject/go-commons/logger"
)

type ErrorWithMetadata interface {
	error
	fmt.Stringer

	ErrorValue() string
	Metadata() interface{}

	SetError(string) ErrorWithMetadata
	SetMetadata(interface{}) ErrorWithMetadata
}

func NewErrorWithMetadata() ErrorWithMetadata {
	return &errWithMetadataImpl{"", ""}
}

type errWithMetadataImpl struct {
	errorVal    string
	metadataVal interface{}
}

func (e *errWithMetadataImpl) Error() string {
	return e.String()
}

func (e *errWithMetadataImpl) String() string {
	metadata, _ := e.metadataVal.(string)

	return fmt.Sprintf("%s:%s", e.errorVal, metadata)
}

func (e *errWithMetadataImpl) ErrorValue() string {
	return e.errorVal
}

func (e *errWithMetadataImpl) Metadata() interface{} {
	return e.metadataVal
}

func (e *errWithMetadataImpl) SetError(err string) ErrorWithMetadata {
	e.errorVal = err
	return e
}

func (e *errWithMetadataImpl) SetMetadata(metadata interface{}) ErrorWithMetadata {
	_, ok := metadata.(fmt.Stringer) // Currently only supports the types that are fmt.Stringer

	if !ok {
		logger.WithField("metadata", metadata).Error("Currently metadata needs to be fmt.Stringer")
		metadata = ""
	}

	e.metadataVal = metadata
	return e
}
