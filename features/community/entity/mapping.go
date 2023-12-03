package entity

import "recything/features/community/model"

func CoreCommunityToModelCommunity(data CommunityCore) model.Community {
	return model.Community{
		Name:        data.Name,
		Description: data.Description,
		Location:    data.Location,
		Members:     data.Members,
		MaxMembers:  data.MaxMembers,
		Image:       data.Image,
	}
}

func ListCoreCommunityToModelCommunity(data []CommunityCore) []model.Community {
	list := []model.Community{}
	for _, v := range data {
		result := CoreCommunityToModelCommunity(v)
		list = append(list, result)
	}
	return list
}

func ModelCommunityToCoreCommunity(data model.Community) CommunityCore {
	return CommunityCore{
		Id:          data.Id,
		Name:        data.Name,
		Description: data.Description,
		Location:    data.Location,
		Members:     data.Members,
		MaxMembers:  data.MaxMembers,
		Image:       data.Image,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func ListModelCommunityToCoreCommunity(data []model.Community) []CommunityCore {
	list := []CommunityCore{}
	for _, v := range data {
		result := ModelCommunityToCoreCommunity(v)
		list = append(list, result)
	}
	return list
}
