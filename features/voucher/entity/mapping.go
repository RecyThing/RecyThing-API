package entity

import "recything/features/voucher/model"

func CoreVoucherToModelVoucher(data VoucherCore) model.Voucher {
	return model.Voucher{
		Image:       data.Image,
		RewardName:  data.RewardName,
		Point:       data.Point,
		Description: data.Description,
		StartDate:   data.StartDate,
		EndDate:     data.EndDate,
	}
}

func ListCoreVoucherToModelVoucher(data []VoucherCore) []model.Voucher {
	list := []model.Voucher{}
	for _, v := range data {
		result := CoreVoucherToModelVoucher(v)
		list = append(list, result)
	}
	return list
}

func ModelVoucherToCoreVoucher(data model.Voucher) VoucherCore {
	return VoucherCore{
		Id:          data.Id,
		Image:       data.Image,
		RewardName:  data.RewardName,
		Point:       data.Point,
		Description: data.Description,
		StartDate:   data.StartDate,
		EndDate:     data.EndDate,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func ListModelVoucherToCoreVoucher(data []model.Voucher) []VoucherCore {
	list := []VoucherCore{}
	for _, v := range data {
		result := ModelVoucherToCoreVoucher(v)
		list = append(list, result)
	}
	return list
}
