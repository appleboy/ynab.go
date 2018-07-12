package transaction_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"

	"bmvs.io/ynab"
	"bmvs.io/ynab/api"
	"bmvs.io/ynab/api/transaction"
)

func TestService_GetTransactions(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/transactions"
	httpmock.RegisterResponder("GET", url,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, `{
  "data": {
    "transactions": [
      {
        "id": "e6ad88f5-6f16-4480-9515-5377012750dd",
        "date": "2018-03-10",
        "amount": -43950,
        "memo": "nice memo",
        "cleared": "reconciled",
        "approved": true,
        "flag_color": null,
        "account_id": "09eaca5e-6f16-4480-9515-828fb90638f2",
        "account_name": "Bank Name",
        "payee_id": "6216ab4b-6f16-4480-9515-be2dee26ab0d",
        "payee_name": "Supermarket",
        "category_id": "e9517027-6f16-4480-9515-5981bed2e9e1",
        "category_name": "Split (Multiple Categories)...",
        "transfer_account_id": null,
        "import_id": null,
        "deleted": false,
        "subtransactions": [
          {
            "id": "9453526b-2f58-4c02-9683-a30c2a1192d7",
            "transaction_id": "e6ad88f5-6f16-4480-9515-5377012750dd",
            "amount": -33970,
            "memo": "Debit Card Payment",
            "payee_id": "6216ab4b-bb05-4574-b4b5-be2dee26ab0d",
            "category_id": "080985e4-4175-43e4-96bb-d207a9d2c8ce",
            "transfer_account_id": null,
            "deleted": false
          }
				]
			}
		]
	}
}
		`), nil
		},
	)

	client := ynab.NewClient("")
	transactions, err := client.Transaction().GetTransactions("aa248caa-eed7-4575-a990-717386438d2c", nil)
	assert.NoError(t, err)

	expectedDate, err := api.NewDateFromString("2018-03-10")
	assert.NoError(t, err)

	expectedMemo := "nice memo"
	expectedPayeeID := "6216ab4b-6f16-4480-9515-be2dee26ab0d"
	expectedPayeeName := "Supermarket"
	expectedCategoryID := "e9517027-6f16-4480-9515-5981bed2e9e1"
	expectedCategoryName := "Split (Multiple Categories)..."

	expectedSubTransactionMemo := "Debit Card Payment"
	expectedSubTransactionPayeeID := "6216ab4b-bb05-4574-b4b5-be2dee26ab0d"
	expectedSubTransactionCategoryID := "080985e4-4175-43e4-96bb-d207a9d2c8ce"

	expected := &transaction.Transaction{
		ID:           "e6ad88f5-6f16-4480-9515-5377012750dd",
		Date:         expectedDate,
		Amount:       int64(-43950),
		Memo:         &expectedMemo,
		Cleared:      transaction.ClearingStatusReconciled,
		Approved:     true,
		AccountID:    "09eaca5e-6f16-4480-9515-828fb90638f2",
		AccountName:  "Bank Name",
		PayeeID:      &expectedPayeeID,
		PayeeName:    &expectedPayeeName,
		CategoryID:   &expectedCategoryID,
		CategoryName: &expectedCategoryName,
		Deleted:      false,
		SubTransactions: []*transaction.SubTransaction{
			{
				ID:            "9453526b-2f58-4c02-9683-a30c2a1192d7",
				TransactionID: "e6ad88f5-6f16-4480-9515-5377012750dd",
				Amount:        int64(-33970),
				Memo:          &expectedSubTransactionMemo,
				PayeeID:       &expectedSubTransactionPayeeID,
				CategoryID:    &expectedSubTransactionCategoryID,
				Deleted:       false,
			},
		},
	}
	assert.Equal(t, expected, transactions[0])
}

func TestService_GetTransactionsByAccount(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/accounts/09eaca5e-6f16-4480-9515-828fb90638f2/transactions"
	httpmock.RegisterResponder("GET", url,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, `{
  "data": {
    "transactions": [
      {
        "id": "e6ad88f5-6f16-4480-9515-5377012750dd",
        "date": "2018-03-10",
        "amount": -43950,
        "memo": "nice memo",
        "cleared": "reconciled",
        "approved": true,
        "flag_color": null,
        "account_id": "09eaca5e-6f16-4480-9515-828fb90638f2",
        "account_name": "Bank Name",
        "payee_id": "6216ab4b-6f16-4480-9515-be2dee26ab0d",
        "payee_name": "Supermarket",
        "category_id": "e9517027-6f16-4480-9515-5981bed2e9e1",
        "category_name": "Split (Multiple Categories)...",
        "transfer_account_id": null,
        "import_id": null,
        "deleted": false,
        "subtransactions": [
          {
            "id": "9453526b-2f58-4c02-9683-a30c2a1192d7",
            "transaction_id": "e6ad88f5-6f16-4480-9515-5377012750dd",
            "amount": -33970,
            "memo": "Debit Card Payment",
            "payee_id": "6216ab4b-bb05-4574-b4b5-be2dee26ab0d",
            "category_id": "080985e4-4175-43e4-96bb-d207a9d2c8ce",
            "transfer_account_id": null,
            "deleted": false
          }
				]
			}
		]
	}
}
		`), nil
		},
	)

	client := ynab.NewClient("")
	transactions, err := client.Transaction().GetTransactionsByAccount(
		"aa248caa-eed7-4575-a990-717386438d2c",
		"09eaca5e-6f16-4480-9515-828fb90638f2",
		nil,
	)
	assert.NoError(t, err)

	expectedDate, err := api.NewDateFromString("2018-03-10")
	assert.NoError(t, err)

	expectedMemo := "nice memo"
	expectedPayeeID := "6216ab4b-6f16-4480-9515-be2dee26ab0d"
	expectedPayeeName := "Supermarket"
	expectedCategoryID := "e9517027-6f16-4480-9515-5981bed2e9e1"
	expectedCategoryName := "Split (Multiple Categories)..."

	expectedSubTransactionMemo := "Debit Card Payment"
	expectedSubTransactionPayeeID := "6216ab4b-bb05-4574-b4b5-be2dee26ab0d"
	expectedSubTransactionCategoryID := "080985e4-4175-43e4-96bb-d207a9d2c8ce"

	expected := &transaction.Transaction{
		ID:           "e6ad88f5-6f16-4480-9515-5377012750dd",
		Date:         expectedDate,
		Amount:       int64(-43950),
		Memo:         &expectedMemo,
		Cleared:      transaction.ClearingStatusReconciled,
		Approved:     true,
		AccountID:    "09eaca5e-6f16-4480-9515-828fb90638f2",
		AccountName:  "Bank Name",
		PayeeID:      &expectedPayeeID,
		PayeeName:    &expectedPayeeName,
		CategoryID:   &expectedCategoryID,
		CategoryName: &expectedCategoryName,
		Deleted:      false,
		SubTransactions: []*transaction.SubTransaction{
			{
				ID:            "9453526b-2f58-4c02-9683-a30c2a1192d7",
				TransactionID: "e6ad88f5-6f16-4480-9515-5377012750dd",
				Amount:        int64(-33970),
				Memo:          &expectedSubTransactionMemo,
				PayeeID:       &expectedSubTransactionPayeeID,
				CategoryID:    &expectedSubTransactionCategoryID,
				Deleted:       false,
			},
		},
	}
	assert.Equal(t, expected, transactions[0])
}

func TestService_GetTransactionsByCategory(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/categories/a33c906e-444c-469c-be27-04c8e0c9959f/transactions"
	httpmock.RegisterResponder("GET", url,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, `{
  "data": {
    "transactions": [
      {
        "type": "transaction",
        "id": "c132c55c-1200-4606-a321-99f4ec24b4df",
        "parent_transaction_id": null,
        "date": "2018-01-10",
        "amount": -42000,
        "memo": "",
        "cleared": "reconciled",
        "approved": true,
        "flag_color": null,
        "account_id": "134d159-444c-469c-be27-44094e388fa0",
        "account_name": "Cash",
        "payee_id": "b391144e-444c-469c-be27-fed6aa352a7a",
        "payee_name": "Landlord",
        "category_id": "a33c906e-444c-469c-be27-04c8e0c9959f",
        "category_name": "Rent",
        "transfer_account_id": null,
        "import_id": null,
        "deleted": false
      }
		]
	}
}
		`), nil
		},
	)

	client := ynab.NewClient("")
	transactions, err := client.Transaction().GetTransactionsByCategory(
		"aa248caa-eed7-4575-a990-717386438d2c",
		"a33c906e-444c-469c-be27-04c8e0c9959f",
		nil,
	)
	assert.NoError(t, err)

	expectedDate, err := api.NewDateFromString("2018-01-10")
	assert.NoError(t, err)

	expectedMemo := ""
	expectedPayeeID := "b391144e-444c-469c-be27-fed6aa352a7a"
	expectedPayeeName := "Landlord"
	expectedCategoryID := "a33c906e-444c-469c-be27-04c8e0c9959f"
	expectedCategoryName := "Rent"

	expected := &transaction.Hybrid{
		Type:         transaction.TypeTransaction,
		ID:           "c132c55c-1200-4606-a321-99f4ec24b4df",
		Date:         expectedDate,
		Amount:       int64(-42000),
		Memo:         &expectedMemo,
		Cleared:      transaction.ClearingStatusReconciled,
		Approved:     true,
		AccountID:    "134d159-444c-469c-be27-44094e388fa0",
		AccountName:  "Cash",
		PayeeID:      &expectedPayeeID,
		PayeeName:    &expectedPayeeName,
		CategoryID:   &expectedCategoryID,
		CategoryName: &expectedCategoryName,
		Deleted:      false,
	}
	assert.Equal(t, expected, transactions[0])
}

func TestFilter_ToQuery(t *testing.T) {
	sinceDate, err := api.NewDateFromString("2020-02-02")
	assert.NoError(t, err)

	var zeroDate api.Date

	uncategorizedTransaction := transaction.StatusUncategorized
	unapprovedTransaction := transaction.StatusUnapproved

	table := []struct {
		Input  transaction.Filter
		Output string
	}{
		{
			Input:  transaction.Filter{SinceDate: &sinceDate, Type: &unapprovedTransaction},
			Output: "since_date=2020-02-02&type=unapproved",
		},
		{
			Input:  transaction.Filter{SinceDate: &sinceDate, Type: &uncategorizedTransaction},
			Output: "since_date=2020-02-02&type=uncategorized",
		},
		{
			Input:  transaction.Filter{SinceDate: &sinceDate},
			Output: "since_date=2020-02-02",
		},
		{
			Input:  transaction.Filter{Type: &uncategorizedTransaction},
			Output: "type=uncategorized",
		},
		{
			Input:  transaction.Filter{SinceDate: &zeroDate, Type: &uncategorizedTransaction},
			Output: "type=uncategorized",
		},
		{
			Input:  transaction.Filter{},
			Output: "",
		},
	}

	for _, test := range table {
		assert.Equal(t, test.Output, test.Input.ToQuery())
	}
}