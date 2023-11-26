package request

import "recything/features/voucher/entity"

func RequestVoucherToCoreVoucher(data VoucherRequest) entity.VoucherCore {
	return entity.VoucherCore{
		Image:       data.Image,
		RewardName:  data.Reward_Name,
		Point:       data.Point,
		Description: data.Description,
		StartDate:   data.Start_Date,
		EndDate:     data.End_Date,
	}
}

// func ListRequestVoucherToCoreVoucher(data []VoucherRequest) []entity.VoucherCore {
// 	list := []entity.VoucherCore{}
// 	for _, v := range data {
// 		result := RequestVoucherToCoreVoucher(v)
// 		list = append(list, result)
// 	}
// 	return list
// }