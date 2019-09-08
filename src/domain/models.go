package domain

import (
	"errors"
	"strconv"
)

const (
	PriceTypeCost              = "COST"
	PriceTypeDiscount          = "Discount"
	OperatorEqual              = "EQ"
	OperatorGreaterThanOrEqual = "GTE"
	OperatorLessThanOrEqual    = "LTE"
)

type RuleApplicability struct {
	CodeName string `json:"codeName"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type Price struct {
	Cost                float64             `json:"cost"`
	PriceType           string              `json:"priceType,omitempty"`
	RuleApplicabilities []RuleApplicability `json:"ruleApplicabilities,omitempty"`
}

type Component struct {
	Name   string  `json:"name"`
	IsMain bool    `json:"isMain,omitempty"`
	Prices []Price `json:"prices"`
}

type Product struct {
	Name       string      `json:"name"`
	Components []Component `json:"components"`
}

type Condition struct {
	RuleName string `json:"ruleName"`
	Value    string `json:"value"`
}

type Offer struct {
	Product
	TotalCost Price `json:"totalCost"`
}

func Calculate(product *Product, conditions []Condition) (offer *Offer, err error) {
	if product == nil || conditions == nil {
		return nil, nil
	}
	productForOffer := Product{}
	offer = &Offer{}
	components := []Component{}
	matchComponent := Component{}
	for _, component := range product.Components {
		isMatchComponent := false
		matchPrices := []Price{}
		for _, price := range component.Prices {
			if price.PriceType == PriceTypeDiscount {
				matchPrices = append(matchPrices, price)
				matchComponent = Component{
					Name:   component.Name,
					IsMain: component.IsMain,
					Prices: matchPrices,
				}
				continue
			}
			isMatch := []bool{}
			for _, rule := range price.RuleApplicabilities {
				for _, condition := range conditions {
					if condition.RuleName == rule.CodeName {
						isMatchPrice, err := isMatchRule(condition, rule)
						isMatch = append(isMatch, isMatchPrice)
						if err != nil {
							return nil, err
						}
					}
				}
			}
			isMatchPrice := true
			for _, match := range isMatch {
				if !match {
					isMatchPrice = false
				}
			}
			if isMatchPrice {
				isMatchComponent = true
				matchPrices = append(matchPrices, price)
				matchComponent = Component{
					Name:   component.Name,
					IsMain: component.IsMain,
					Prices: matchPrices,
				}
			}
		}
		if isMatchComponent {
			components = append(components, matchComponent)
		}
	}
	productForOffer = Product{
		Name:       product.Name,
		Components: components,
	}
	offer, err = makeOffer(&productForOffer)
	if err != nil {
		return nil, err
	}
	return offer, nil
}

func makeOffer(product *Product) (*Offer, error) {
	isHasMain := false
	offer := Offer{Product: Product{Name: product.Name}}
	for _, component := range product.Components {
		matchDiscounts := []Price{}
		totalDiscount := 0.0
		if component.IsMain {
			isHasMain = true
		}
		discounts := []Price{}
		prices := []Price{}
		for _, price := range component.Prices {
			if price.PriceType == PriceTypeDiscount {
				discounts = append(discounts, price)
			} else {
				prices = append(prices, price)
			}
		}
		if len(prices) > 1 || len(prices) == 0 {
			return nil, nil
		} else {
			component.Prices = prices
		}
		if len(discounts) == 0 {
			offer.TotalCost.Cost += component.Prices[0].Cost
			offer.Components = append(offer.Components, component)
			continue
		}
		for _, discount := range discounts {
			for _, discountRule := range discount.RuleApplicabilities {
				for _, priceRule := range component.Prices[0].RuleApplicabilities {
					if discountRule.CodeName == priceRule.CodeName {
						condition := Condition{
							RuleName: priceRule.CodeName,
							Value:    priceRule.Value,
						}
						isMatch, err := isMatchRule(condition, discountRule)
						if err != nil {
							return nil, err
						}
						if isMatch {
							matchDiscounts = append(matchDiscounts, discount)
						}
					}
				}
			}
		}
		for _, discount := range matchDiscounts {
			if totalDiscount < discount.Cost {
				totalDiscount = discount.Cost
			}
		}
		component.Prices[0].Cost *= (100 - totalDiscount) / 100
		component.Prices[0] = Price{Cost: component.Prices[0].Cost}
		offer.TotalCost.Cost += component.Prices[0].Cost
		offer.Components = append(offer.Components, component)
	}
	if !isHasMain {
		return &Offer{}, nil
	}
	return &offer, nil
}

func isMatchRule(condition Condition, applicability RuleApplicability) (bool, error) {
	isMatch := false
	switch applicability.Operator {
	case OperatorEqual:
		if condition.Value == applicability.Value {
			isMatch = true
		}
	case OperatorGreaterThanOrEqual:
		condValue, applicabilityValue, err := convertRuleValues(condition.Value, applicability.Value)
		if err != nil {
			isMatch = false
			return isMatch, err
		}
		if condValue != nil && applicabilityValue != nil && *condValue >= *applicabilityValue {
			isMatch = true
		}
	case OperatorLessThanOrEqual:
		condValue, applicabilityValue, err := convertRuleValues(condition.Value, applicability.Value)
		if err != nil {
			isMatch = false
			return isMatch, err
		}
		if condValue != nil && applicabilityValue != nil && *condValue <= *applicabilityValue {
			isMatch = true
		}
	}
	return isMatch, nil
}

func convertRuleValues(condValue string, applicabilityValue string) (*int64, *int64, error) {
	condValueOut, err := strconv.ParseInt(condValue, 10, 64)
	if err != nil {
		return nil, nil, err
	}
	applicabilityValueOut, err := strconv.ParseInt(applicabilityValue, 10, 64)
	if err != nil {
		err = errors.New("invalid applicability")
		return nil, nil, err
	}
	return &condValueOut, &applicabilityValueOut, nil
}
