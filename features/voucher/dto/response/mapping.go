package response

import "recything/features/voucher/entity"

func CoreVoucherToResponVoucher(data entity.VoucherCore) VoucherResponse {
	return VoucherResponse{
		Id:         data.Id,
		Image:      data.Image,
		RewardName: data.RewardName,
		Point:      data.Point,
		StartDate:  data.StartDate,
		EndDate:    data.EndDate,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ListCoreVoucherToCoreVoucher(data []entity.VoucherCore) []VoucherResponse {
	list := []VoucherResponse{}
	for _, v := range data {
		result := CoreVoucherToResponVoucher(v)
		list = append(list, result)
	}
	return list
}
