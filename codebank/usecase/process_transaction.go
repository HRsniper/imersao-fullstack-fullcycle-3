package usecase

import (
	"time"

	"github.com/hrsniper/imersao-fullstack-fullcycle-3/domain"
	"github.com/hrsniper/imersao-fullstack-fullcycle-3/dto"
)

type UseCaseTransaction struct {
	TransactionRepository domain.TransactionRepository
}

func NewUseCaseTransaction(transactionRepository domain.TransactionRepository) UseCaseTransaction {
	return UseCaseTransaction{TransactionRepository: transactionRepository}
}

// method of UseCaseTransaction struct
func (useCaseTransaction UseCaseTransaction) ProcessTransaction(transactionDto dto.Transaction) (domain.Transaction, error) {
	creditCard := useCaseTransaction.hydrateCreditCard(transactionDto)

	creditCardBalanceAndLimit, err := useCaseTransaction.TransactionRepository.GetCreditCard(*creditCard)

	if err != nil {
		return domain.Transaction{}, err
	}

	creditCard.ID = creditCardBalanceAndLimit.ID
	creditCard.Limit = creditCardBalanceAndLimit.Limit
	creditCard.Balance = creditCardBalanceAndLimit.Balance

	transaction := useCaseTransaction.newTransaction(transactionDto, creditCardBalanceAndLimit)

	transaction.ProcessAndValidate(creditCard)

	err = useCaseTransaction.TransactionRepository.SaveTransaction(*transaction, *creditCard)

	if err != nil {
		return domain.Transaction{}, err
	}

	return *transaction, nil
}

// method of UseCaseTransaction struct
func (useCaseTransaction UseCaseTransaction) hydrateCreditCard(transactionDto dto.Transaction) *domain.CreditCard {
	creditCard := domain.NewCreditCard()
	creditCard.Name = transactionDto.Name
	creditCard.Number = transactionDto.Number
	creditCard.ExpirationMonth = transactionDto.ExpirationMonth
	creditCard.ExpirationYear = transactionDto.ExpirationYear
	creditCard.CVV = transactionDto.CVV

	return creditCard
}

// method of UseCaseTransaction struct
func (useCaseTransaction UseCaseTransaction) newTransaction(transactionDto dto.Transaction, creditCard domain.CreditCard) *domain.Transaction {
	transaction := domain.NewTransaction()
	transaction.CreditCardId = creditCard.ID
	transaction.Amount = transactionDto.Amount
	transaction.Store = transactionDto.Store
	transaction.Description = transactionDto.Description
	transaction.CreatedAt = time.Now()

	return transaction
}