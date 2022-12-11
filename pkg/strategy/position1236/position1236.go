package position1236

import (
	"github/com/cbl315/crypto-center/pkg/consts"
)

type PositionAndAmount struct {
	Price  float64
	Amount float64
}

type PositionRes1236 struct {
	LosePrice float64
	Direction consts.ContractDirection
	Positions []PositionAndAmount
}

type PercentRate struct {
	Percent float64
	Rate    float64
}

var (
	firstPositionPercentage  = 0.01
	secondPositionPercentage = 0.002
	losePositionPercentage   = 0.012
)

func getDirectionIndex(direction consts.ContractDirection) float64 {
	switch direction {
	case consts.Short:
		return -1
	case consts.Long:
		return 1
	default:
		return 1
	}
}

func GetPositionAndAmount(keyPrice, amount float64, direction consts.ContractDirection, lever uint) (res PositionRes1236) {
	res = PositionRes1236{
		Direction: direction,
		Positions: []PositionAndAmount{},
	}
	// get index according to direction
	var percentIndex float64 = getDirectionIndex(direction)
	var TotalAmount = amount * float64(lever)
	res.LosePrice = keyPrice * (1 - losePositionPercentage*percentIndex)
	pricePercentageArr := []PercentRate{
		{
			Percent: 1 + firstPositionPercentage*percentIndex,
			Rate:    1 / float64(1+2+3+6),
		},
		{
			Percent: 1 + secondPositionPercentage*percentIndex,
			Rate:    2 / float64(1+2+3+6),
		},
		{
			Percent: 1 - secondPositionPercentage*percentIndex,
			Rate:    3 / float64(1+2+3+6),
		},
		{
			Percent: 1 - firstPositionPercentage*percentIndex,
			Rate:    6 / float64(1+2+3+6),
		},
	}
	for _, pricePercent := range pricePercentageArr {
		onePrice := keyPrice * pricePercent.Percent
		oneAmount := TotalAmount * pricePercent.Rate
		res.Positions = append(res.Positions, PositionAndAmount{
			Price:  onePrice,
			Amount: oneAmount,
		})
	}
	return res
}
