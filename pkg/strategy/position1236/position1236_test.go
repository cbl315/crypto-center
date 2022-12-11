package position1236

import (
	"github/com/cbl315/crypto-center/pkg/consts"
	"reflect"
	"testing"
)

func TestGetPositionAndAmount(t *testing.T) {
	type Input struct {
		KeyPrice  float64
		amount    float64
		direction consts.ContractDirection
		lever     uint
	}
	type TestCase struct {
		Description string
		input       Input
		Expected    PositionRes1236
	}
	var testCases = []TestCase{
		{
			Description: "Long",
			input: Input{
				KeyPrice:  20000,
				amount:    1200,
				direction: consts.Long,
				lever:     10,
			},
			Expected: PositionRes1236{
				LosePrice: 20000 * 0.988,
				Direction: consts.Long,
				Positions: []PositionAndAmount{
					{
						Price:  20000 * 1.01,
						Amount: 12000 * 1 / (1 + 2 + 3 + 6),
					},
					{
						Price:  20000 * 1.002,
						Amount: 12000 * 2 / (1 + 2 + 3 + 6),
					},
					{
						Price:  20000 * 0.998,
						Amount: 12000 * 3 / (1 + 2 + 3 + 6),
					},
					{
						Price:  20000 * 0.99,
						Amount: 12000 * 6 / (1 + 2 + 3 + 6),
					},
				},
			},
		},
		{
			Description: "Short",
			input: Input{
				KeyPrice:  20000,
				amount:    1200,
				direction: consts.Short,
				lever:     10,
			},
			Expected: PositionRes1236{
				LosePrice: 20000 * 1.012,
				Direction: consts.Short,
				Positions: []PositionAndAmount{
					{
						Price:  20000 * 0.99,
						Amount: 12000 * 1 / (1 + 2 + 3 + 6),
					},
					{
						Price:  20000 * 0.998,
						Amount: 12000 * 2 / (1 + 2 + 3 + 6),
					},
					{
						Price:  20000 * 1.002,
						Amount: 12000 * 3 / (1 + 2 + 3 + 6),
					},
					{
						Price:  20000 * 1.01,
						Amount: 12000 * 6 / (1 + 2 + 3 + 6),
					},
				},
			},
		},
	}
	for _, testCase := range testCases {
		got := GetPositionAndAmount(testCase.input.KeyPrice, testCase.input.amount, testCase.input.direction, testCase.input.lever)
		if !reflect.DeepEqual(got, testCase.Expected) {
			t.Errorf("case %s failed, got: %v, expected: %v", testCase.Description, got, testCase.Expected)
		}
	}
}
