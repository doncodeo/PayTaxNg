package services 

import "math"

const (
	TaxFreeThreshold = 800000
	RentReliefRate  = 0.20
	RentReliefCap   = 500000
)

type TaxService struct{}

func (s *TaxService) CalculateTax(
	annualIncome float64,
	annualRent float64,
	pension float64,
	deductions[]float64,
)(RentRelief float64, totalDeductions float64, taxableIncome float64, annualTax float64){

	// Calculate Rent Relief
	RentRelief = math.Min(annualRent*RentReliefRate, RentReliefCap)

	// Sum dynamic deductions
	var otherDeductions float64
	for _, amount := range deductions {
		otherDeductions += amount
	}

	//Total Deductions 
	totalDeductions = RentRelief + pension + otherDeductions

	// Chargeable Income
	taxableIncome = annualIncome - totalDeductions

	if taxableIncome <= TaxFreeThreshold {
		return RentRelief, totalDeductions, taxableIncome, 0
	}

	annualTax = calculateProgressiveTax(taxableIncome)

	return 
}	

// Progressive Tax Calculation
func calculateProgressiveTax(taxableIncome float64) float64 {
	var tax float64

	brackets := []struct {
		limit  float64
		rate   float64
	}{
		{3000000, 0.15},
		{12000000, 0.18},
		{25000000, 0.21},
		{50000000, 0.23},
		{math.Inf(1), 0.24},
	}

	remaining := taxableIncome - TaxFreeThreshold
	prevLimit := float64(TaxFreeThreshold)

	for _, bracket := range brackets {
		if remaining <= 0 {
			break
		}

		taxable := math.Min(bracket.limit - prevLimit, remaining)
		tax += taxable * bracket.rate 
		remaining -= taxable 
		prevLimit = bracket.limit
	}

	return tax
}