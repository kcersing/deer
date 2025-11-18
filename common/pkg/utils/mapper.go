package utils

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
)

type CopierMapper[DTO any, ENTITY any] struct{}

func NewCopierMapper[DTO any, ENTITY any]() *CopierMapper[DTO, ENTITY] {
	return &CopierMapper[DTO, ENTITY]{}
}
func (m *CopierMapper[DTO, ENTITY]) ToEntity(dto *DTO) *ENTITY {
	if dto == nil {
		return nil
	}

	var entity ENTITY
	if err := copier.CopyWithOption(&entity, dto, copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		klog.Errorf("failed to copy DTO to entity: %w", err)
		return nil
	}

	return &entity
}

func (m *CopierMapper[DTO, ENTITY]) ToDTO(entity *ENTITY) *DTO {
	if entity == nil {
		return nil
	}

	var dto DTO
	if err := copier.CopyWithOption(&dto, entity, copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		klog.Errorf("failed to copy entity to DTO: %w", err)
		return nil
	}

	return &dto
}
func (m *CopierMapper[DTO, ENTITY]) ToDTOyList(entits []*ENTITY) []*DTO {
	if entits == nil {
		return nil
	}

	dtos := make([]*DTO, len(entits))
	for i, entity := range entits {
		dtos[i] = m.ToDTO(entity)
	}
	return dtos
}
