package service

import (
	"context"
	"kalkulator_pinjaman/model"
	"math"
)

func CalculateLoan(ctx context.Context, payload model.CalculatorLoanRequest) (result []model.LoanTable, err error) {
	var (
		remainingInstallments = payload.LimitLoan
		monthLength           = int64(12)
		interestRate          = payload.InterestRate / 100
	)

	if payload.LoanDuration < monthLength {
		monthLength = payload.LoanDuration
	}

	for i := 0; i < int(payload.LoanDuration); i++ {
		var (
			interest             = interestRate / float64(monthLength)
			anuitasPercentage    = math.Pow(1+interest, float64(payload.LoanDuration))
			anuitas              = (interest * payload.LimitLoan) * anuitasPercentage / (anuitasPercentage - 1)
			interestInstallments = (interestRate / float64(30*monthLength)) * 30 * remainingInstallments
			date                 = payload.StartDate
		)

		if i > 0 {
			date = date.AddDate(0, 1, 0)
		}

		tempResult := model.LoanTable{
			LoanNumber:            int64(i + 1),
			Date:                  date,
			TotalInstallments:     roundFloat(anuitas, 2),
			PrincipalInstallments: roundFloat(anuitas-interestInstallments, 2),
			InterestInstallments:  roundFloat(interestInstallments, 2),
		}

		tempResult.RemainingInstallments = roundFloat(remainingInstallments-tempResult.PrincipalInstallments, 2)

		result = append(result, tempResult)
		remainingInstallments = tempResult.RemainingInstallments
	}
	return
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
