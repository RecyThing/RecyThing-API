package response

import "recything/features/voucher/entity"

func CoreVoucherToResponVoucher(data entity.VoucherCore) VoucherResponse {
	return VoucherResponse{
		Id:         data.Id,
		RewardName: data.RewardName,
		Point:      data.Point,
		StartDate:  data.StartDate,
		EndDate:    data.EndDate,
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
