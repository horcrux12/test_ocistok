package model

import "time"

type CalculatorLoanRequest struct {
	StartDateStr string    `json:"start_date"`
	StartDate    time.Time `json:"-"`
	LimitLoan    float64   `json:"limit_loan"`
	InterestRate float64   `json:"interest_rate"`
	LoanDuration int64     `json:"loan_duration"`
}

type LoanTable struct {
	LoanNumber            int64     `json:"loan_number"`
	Date                  time.Time `json:"date"`
	TotalInstallments     float64   `json:"total_installments"`
	PrincipalInstallments float64   `json:"principal_installments"`
	InterestInstallments  float64   `json:"interest_installments"`
	RemainingInstallments float64   `json:"remaining_installments"`
}
