package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"nigeria-tax-api/internal/models"
	"nigeria-tax-api/internal/services"
)

func CalculateTax(c *gin.Context){
	var req models.TaxRequest

	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})

		return
	}

	annualIncome := req.MonthlySalary * 12

	// Extract deduction amounts 
	var deductionAmounts []float64

	for _, d := range req.Deductions{
		if d.Amount > 0 {
			deductionAmounts = append(deductionAmounts, d.Amount)
		}
	}

	service := services.TaxService{}
	rentRelief, totalDeductions, taxableIncome, annualTax :=
		service.CalculateTax(
			annualIncome, 
			req.AnnualREnt,
			req.Pension,
			deductionAmounts,
		)
	
	response := models.TaxResponse{
		AnnualIncome: annualIncome,
		RentRelief: rentRelief,
		TotalDeductions: totalDeductions,
		TaxableIncome: taxableIncome,
		AnnualTax: annualTax,
		MonthlyTax: annualTax / 12,
	}

	c.JSON(http.StatusOK, response)
}

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}