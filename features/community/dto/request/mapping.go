package request

import "recything/features/community/entity"

func RequestCommunityToCoreCommunity(data CommunityRequest) entity.CommunityCore {
	return entity.CommunityCore{
		Name:        data.Name,
		Description: data.Description,
		Location:    data.Location,
		MaxMembers:  data.Max_Members,
		Image:       data.Image,
	}
}
