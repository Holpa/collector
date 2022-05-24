package graph

import (
	"math/big"
	"time"

	"github.com/steschwa/hopper-analytics-collector/utils"
)

type (
	DayDataInterface interface {
		DayId() uint
		AmountFloat() float64
	}
	DayData struct {
		Id             uint
		StartTimestamp time.Time
	}

	MintDayData struct {
		DayData
		Minted *big.Int
	}
	BurnDayData struct {
		DayData
		Burned *big.Int
	}
	CirculatingDayData struct {
		DayData
		Circulating *big.Int
	}
	StakeDepositDayData struct {
		DayData
		Deposited *big.Int
	}
	StakeWithdrawDayData struct {
		DayData
		Withdrawn *big.Int
	}
)

// ----------------------------------------
// Interface implementations
// ----------------------------------------

func (dayData DayData) DayId() uint {
	return dayData.Id
}

func (dayData MintDayData) AmountFloat() float64 {
	f, _ := utils.ToDecimal(dayData.Minted, 18).Float64()
	return f
}

func (dayData BurnDayData) AmountFloat() float64 {
	f, _ := utils.ToDecimal(dayData.Burned, 18).Float64()
	return f
}

func (dayData CirculatingDayData) AmountFloat() float64 {
	f, _ := utils.ToDecimal(dayData.Circulating, 18).Float64()
	return f
}

func (dayData StakeDepositDayData) AmountFloat() float64 {
	f, _ := utils.ToDecimal(dayData.Deposited, 18).Float64()
	return f
}

func (dayData StakeWithdrawDayData) AmountFloat() float64 {
	f, _ := utils.ToDecimal(dayData.Withdrawn, 18).Float64()
	return f
}
