package interfaces

import . "go-rti-testing/src/domain"

var product = Product{
	Name: "Игровой",
	Components: []Component{
		{
			IsMain: true,
			Name:   "Интернет",
			Prices: []Price{
				{
					Cost:      100,
					PriceType: PriceTypeCost,
					RuleApplicabilities: []RuleApplicability{
						{
							CodeName: "technology",
							Operator: OperatorEqual,
							Value:    "adsl",
						},
						{
							CodeName: "internetSpeed",
							Operator: OperatorEqual,
							Value:    "10",
						},
					},
				},
				{
					Cost:      150,
					PriceType: PriceTypeCost,
					RuleApplicabilities: []RuleApplicability{
						{
							CodeName: "technology",
							Operator: OperatorEqual,
							Value:    "adsl",
						},
						{
							CodeName: "internetSpeed",
							Operator: OperatorEqual,
							Value:    "15",
						},
					},
				},
				{
					Cost:      500,
					PriceType: PriceTypeCost,
					RuleApplicabilities: []RuleApplicability{
						{
							CodeName: "technology",
							Operator: OperatorEqual,
							Value:    "xpon",
						},
						{
							CodeName: "internetSpeed",
							Operator: OperatorEqual,
							Value:    "100",
						},
					},
				},
				{
					Cost:      900,
					PriceType: PriceTypeCost,
					RuleApplicabilities: []RuleApplicability{
						{
							CodeName: "technology",
							Operator: OperatorEqual,
							Value:    "xpon",
						},
						{
							CodeName: "internetSpeed",
							Operator: OperatorEqual,
							Value:    "200",
						},
					},
				},
				{
					Cost:      200,
					PriceType: PriceTypeCost,
					RuleApplicabilities: []RuleApplicability{
						{
							CodeName: "technology",
							Operator: OperatorEqual,
							Value:    "fttb",
						},
						{
							CodeName: "internetSpeed",
							Operator: OperatorEqual,
							Value:    "30",
						},
					},
				},
				{
					Cost:      400,
					PriceType: PriceTypeCost,
					RuleApplicabilities: []RuleApplicability{
						{
							CodeName: "technology",
							Operator: OperatorEqual,
							Value:    "fttb",
						},
						{
							CodeName: "internetSpeed",
							Operator: OperatorEqual,
							Value:    "50",
						},
					},
				},
				{
					Cost:      600,
					PriceType: PriceTypeCost,
					RuleApplicabilities: []RuleApplicability{
						{
							CodeName: "technology",
							Operator: OperatorEqual,
							Value:    "fttb",
						},
						{
							CodeName: "internetSpeed",
							Operator: OperatorEqual,
							Value:    "200",
						},
					},
				},
				{
					Cost:      10,
					PriceType: PriceTypeDiscount,
					RuleApplicabilities: []RuleApplicability{
						{
							CodeName: "internetSpeed",
							Operator: OperatorGreaterThanOrEqual,
							Value:    "50",
						},
					},
				},
				{
					Cost:      15,
					PriceType: PriceTypeDiscount,
					RuleApplicabilities: []RuleApplicability{
						{
							CodeName: "internetSpeed",
							Operator: OperatorGreaterThanOrEqual,
							Value:    "100",
						},
					},
				},
			},
		},
		{
			Name: "ADSL Модем",
			Prices: []Price{
				{
					Cost:      300,
					PriceType: PriceTypeCost,
					RuleApplicabilities: []RuleApplicability{
						{
							CodeName: "technology",
							Operator: OperatorEqual,
							Value:    "adsl",
						},
					},
				},
			},
		},
	},
}

func getProduct() *Product {
	return &product
}
