package response

import "recything/features/community/entity"

func CoreCommunityToResponCommunity(data entity.CommunityCore) CommunityResponse {
	return CommunityResponse{
		Name:      data.Name,
		Location:  data.Location,
		CreatedAt: data.CreatedAt,
	}
}

func CoreCommunityToResponCommunityForDetails(data entity.CommunityCore) CommunityResponseForDetails {
	return CommunityResponseForDetails{
		Name:        data.Name,
		Description: data.Description,
		Location:    data.Location,
		MaxMembers:  data.MaxMembers,
		Image:       data.Image,
		CreatedAt:   data.CreatedAt,
	}
}

func ListCoreCommunityToResponseCommunity(data []entity.CommunityCore) []CommunityResponse {
	list := []CommunityResponse{}
	for _, v := range data {
		result := CoreCommunityToResponCommunity(v)
		list = append(list, result)
	}
	return list
}
