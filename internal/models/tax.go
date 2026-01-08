package models

// Deductions represents any allowable tax deductions.
type Deductions struct {
	name string `name:"name"`
	Amount float64 `amount:"amount"`
}

// TaxRequest represents user input for tax calculation. 
type TaxRequest struct {
	MonthlySalary float64 `json:"monthly_salary"`
	AnnualREnt float64 `json:"annual_rent"`
	Pension float64 `json:"pension"`
	Deductions []Deductions `json:"deductions"`
}

// TaxResponse represents the system output after tax calculation.
type TaxResponse struct {
	AnnualIncome float64 `json:"annual_income"`
	RentRelief float64 `json:"rent_relief"`
	TotalDeductions float64 `json:"total_deductions"`
	TaxableIncome float64 `json:"taxable_income"`
	AnnualTax float64 `json:"annual_tax"`
	MonthlyTax float64 `json:"monthly_tax"`
}