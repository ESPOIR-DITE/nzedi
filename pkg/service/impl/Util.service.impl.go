package impl

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/service"
	"github.com/google/uuid"
)

type UtilServiceImpl struct {
}

func (u UtilServiceImpl) GetID(typeName string) string {
	return typeName + "-" + uuid.New().String()
}

func NewUtilServiceImpl() *UtilServiceImpl {
	return &UtilServiceImpl{}
}

var _ service.UtilService = &UtilServiceImpl{}
