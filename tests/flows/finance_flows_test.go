package flows

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

type apiClient struct {
	baseURL string
	client  *http.Client
}

type apiResponse struct {
	StatusCode int
	Body       map[string]any
	Raw        string
}

func newAPIClient(t *testing.T) apiClient {
	t.Helper()

	baseURL := strings.TrimRight(os.Getenv("FINANCE_API_BASE_URL"), "/")
	if baseURL == "" {
		t.Skip("set FINANCE_API_BASE_URL, for example http://localhost:8000/api, to run flow tests")
	}

	return apiClient{
		baseURL: baseURL,
		client:  &http.Client{Timeout: 10 * time.Second},
	}
}

func uniqueName(prefix string) string {
	return fmt.Sprintf("%s-%d", prefix, time.Now().UnixNano())
}

func ptr[T any](v T) *T {
	return &v
}

func (c apiClient) do(t *testing.T, method, path string, payload any) apiResponse {
	t.Helper()

	var body io.Reader
	if payload != nil {
		rawPayload, err := json.Marshal(payload)
		if err != nil {
			t.Fatalf("marshal payload: %v", err)
		}
		body = bytes.NewReader(rawPayload)
	}

	req, err := http.NewRequest(method, c.baseURL+path, body)
	if err != nil {
		t.Fatalf("create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		t.Fatalf("%s %s failed: %v", method, path, err)
	}
	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read response body: %v", err)
	}

	decoded := map[string]any{}
	if len(bytes.TrimSpace(rawBody)) > 0 {
		if err := json.Unmarshal(rawBody, &decoded); err != nil {
			t.Fatalf("decode response body %q: %v", string(rawBody), err)
		}
	}

	return apiResponse{
		StatusCode: resp.StatusCode,
		Body:       decoded,
		Raw:        string(rawBody),
	}
}

func requireStatus(t *testing.T, resp apiResponse, want int) {
	t.Helper()

	if resp.StatusCode != want {
		t.Fatalf("status = %d, want %d, body = %s", resp.StatusCode, want, resp.Raw)
	}
}

func dataMap(t *testing.T, resp apiResponse) map[string]any {
	t.Helper()

	data, ok := resp.Body["data"].(map[string]any)
	if !ok {
		t.Fatalf("response data is missing or invalid: %s", resp.Raw)
	}
	return data
}

func nestedMap(t *testing.T, resp apiResponse, key string) map[string]any {
	t.Helper()

	value, ok := dataMap(t, resp)[key].(map[string]any)
	if !ok {
		t.Fatalf("response data.%s is missing or invalid: %s", key, resp.Raw)
	}
	return value
}

func nestedSlice(t *testing.T, resp apiResponse, key string) []any {
	t.Helper()

	value, ok := dataMap(t, resp)[key].([]any)
	if !ok {
		t.Fatalf("response data.%s is missing or invalid: %s", key, resp.Raw)
	}
	return value
}

func idFrom(t *testing.T, body map[string]any) int {
	t.Helper()

	id, ok := body["id"].(float64)
	if !ok || id <= 0 {
		t.Fatalf("invalid id in response: %#v", body)
	}
	return int(id)
}

func createFinancialAccount(t *testing.T, c apiClient, name string, initialBalance float64) int {
	t.Helper()

	resp := c.do(t, http.MethodPost, "/financial-account", map[string]any{
		"name":            name,
		"type":            "checking",
		"initial_balance": initialBalance,
		"opened_at":       "2026-07-12",
	})
	requireStatus(t, resp, http.StatusCreated)
	return idFrom(t, nestedMap(t, resp, "financial_account"))
}

func createCategory(t *testing.T, c apiClient, name, categoryType string, parentID *int) int {
	t.Helper()

	payload := map[string]any{
		"name": name,
		"type": categoryType,
	}
	if parentID != nil {
		payload["parent_id"] = *parentID
	}

	resp := c.do(t, http.MethodPost, "/categories", payload)
	requireStatus(t, resp, http.StatusCreated)
	return idFrom(t, nestedMap(t, resp, "category"))
}

func currentBalance(t *testing.T, c apiClient, accountID int) float64 {
	t.Helper()

	resp := c.do(t, http.MethodGet, fmt.Sprintf("/financial-account/current-balance/%d", accountID), nil)
	requireStatus(t, resp, http.StatusOK)
	balance, ok := nestedMap(t, resp, "current_balance")["balance"].(float64)
	if !ok {
		t.Fatalf("invalid current balance body: %s", resp.Raw)
	}
	return balance
}

func TestFinancialAccountsFlow(t *testing.T) {
	c := newAPIClient(t)
	accountName := uniqueName("backlog-account")

	accountID := createFinancialAccount(t, c, accountName, 100)

	resp := c.do(t, http.MethodPost, "/financial-account", map[string]any{
		"name":            "",
		"type":            "checking",
		"initial_balance": 0,
		"opened_at":       "2026-07-12",
	})
	requireStatus(t, resp, http.StatusUnprocessableEntity)

	resp = c.do(t, http.MethodPost, "/financial-account", map[string]any{
		"name":            uniqueName("invalid-account-type"),
		"type":            "invalid",
		"initial_balance": 0,
		"opened_at":       "2026-07-12",
	})
	requireStatus(t, resp, http.StatusUnprocessableEntity)

	resp = c.do(t, http.MethodPost, "/financial-account", map[string]any{
		"name":            uniqueName("negative-account-balance"),
		"type":            "checking",
		"initial_balance": -1,
		"opened_at":       "2026-07-12",
	})
	requireStatus(t, resp, http.StatusUnprocessableEntity)

	resp = c.do(t, http.MethodGet, "/financial-account", nil)
	requireStatus(t, resp, http.StatusOK)
	requireListContainsID(t, nestedSlice(t, resp, "financial_accounts"), accountID)

	resp = c.do(t, http.MethodGet, fmt.Sprintf("/financial-account/%d", accountID), nil)
	requireStatus(t, resp, http.StatusOK)

	resp = c.do(t, http.MethodGet, "/financial-account/not-a-number", nil)
	requireStatus(t, resp, http.StatusBadRequest)

	resp = c.do(t, http.MethodGet, "/financial-account/999999999", nil)
	requireStatus(t, resp, http.StatusNotFound)

	resp = c.do(t, http.MethodPut, fmt.Sprintf("/financial-account/%d", accountID), map[string]any{
		"name":            accountName + "-updated",
		"type":            "savings",
		"initial_balance": 150,
		"opened_at":       "2026-07-12",
	})
	requireStatus(t, resp, http.StatusOK)

	resp = c.do(t, http.MethodDelete, fmt.Sprintf("/financial-account/delete/%d", accountID), nil)
	requireStatus(t, resp, http.StatusOK)
}

func TestCategoriesFlow(t *testing.T) {
	c := newAPIClient(t)

	parentID := createCategory(t, c, uniqueName("backlog-category-parent"), "both", nil)
	childID := createCategory(t, c, uniqueName("backlog-category-child"), "expense", &parentID)

	resp := c.do(t, http.MethodPost, "/categories", map[string]any{
		"name": "",
		"type": "expense",
	})
	requireStatus(t, resp, http.StatusUnprocessableEntity)

	resp = c.do(t, http.MethodPost, "/categories", map[string]any{
		"name":      uniqueName("missing-parent-category"),
		"type":      "expense",
		"parent_id": 999999999,
	})
	requireStatus(t, resp, http.StatusUnprocessableEntity)

	resp = c.do(t, http.MethodGet, "/categories", nil)
	requireStatus(t, resp, http.StatusOK)
	requireListContainsID(t, nestedSlice(t, resp, "categories"), parentID)
	requireListContainsID(t, nestedSlice(t, resp, "categories"), childID)

	resp = c.do(t, http.MethodPut, fmt.Sprintf("/categories/%d", childID), map[string]any{
		"name":      uniqueName("self-parent-category"),
		"type":      "expense",
		"parent_id": childID,
	})
	requireStatus(t, resp, http.StatusUnprocessableEntity)

	resp = c.do(t, http.MethodPut, fmt.Sprintf("/categories/%d", childID), map[string]any{
		"name":      uniqueName("updated-child-category"),
		"type":      "expense",
		"parent_id": parentID,
	})
	requireStatus(t, resp, http.StatusOK)

	resp = c.do(t, http.MethodDelete, fmt.Sprintf("/categories/delete/%d", parentID), nil)
	requireStatus(t, resp, http.StatusOK)
}

func TestFinancialTransactionsFlow(t *testing.T) {
	c := newAPIClient(t)

	accountID := createFinancialAccount(t, c, uniqueName("movement-account"), 100)
	incomeCategoryID := createCategory(t, c, uniqueName("movement-income-category"), "income", nil)
	expenseCategoryID := createCategory(t, c, uniqueName("movement-expense-category"), "expense", nil)

	resp := c.do(t, http.MethodPost, "/financial-transaction", map[string]any{
		"financial_account_id": accountID,
		"category_id":          incomeCategoryID,
		"description":          "Backlog income flow",
		"amount":               50,
		"movement_date":        "2026-07-12",
		"reference_date":       "2026-07-12",
		"movement_type":        "entry",
		"operation_type":       "original",
	})
	requireStatus(t, resp, http.StatusCreated)
	entryID := idFrom(t, nestedMap(t, resp, "financial_transaction"))

	if got := currentBalance(t, c, accountID); got != 150 {
		t.Fatalf("balance after entry = %.2f, want 150.00", got)
	}

	resp = c.do(t, http.MethodPost, "/financial-transaction", map[string]any{
		"financial_account_id": accountID,
		"category_id":          expenseCategoryID,
		"description":          "Backlog expense flow",
		"amount":               20,
		"movement_date":        "2026-07-12",
		"reference_date":       "2026-07-12",
		"movement_type":        "exit",
		"operation_type":       "original",
	})
	requireStatus(t, resp, http.StatusCreated)

	if got := currentBalance(t, c, accountID); got != 130 {
		t.Fatalf("balance after exit = %.2f, want 130.00", got)
	}

	resp = c.do(t, http.MethodPost, "/financial-transaction", map[string]any{
		"financial_account_id": accountID,
		"category_id":          expenseCategoryID,
		"description":          "Invalid income category",
		"amount":               10,
		"movement_date":        "2026-07-12",
		"reference_date":       "2026-07-12",
		"movement_type":        "entry",
		"operation_type":       "original",
	})
	requireStatus(t, resp, http.StatusUnprocessableEntity)

	resp = c.do(t, http.MethodGet, "/financial-transaction", nil)
	requireStatus(t, resp, http.StatusOK)
	requireListContainsID(t, nestedSlice(t, resp, "financial_transactions"), entryID)
}

func TestFinancialObligationsAndSettlementsFlow(t *testing.T) {
	c := newAPIClient(t)

	accountID := createFinancialAccount(t, c, uniqueName("obligation-account"), 500)
	expenseCategoryID := createCategory(t, c, uniqueName("obligation-expense-category"), "expense", nil)
	incomeCategoryID := createCategory(t, c, uniqueName("obligation-income-category"), "income", nil)

	payableID := createObligation(t, c, expenseCategoryID, "payable", 200)
	receivableID := createObligation(t, c, incomeCategoryID, "receivable", 300)

	if got := currentBalance(t, c, accountID); got != 500 {
		t.Fatalf("balance after obligations = %.2f, want 500.00", got)
	}

	resp := c.do(t, http.MethodGet, "/financial-obligation?status=pending", nil)
	requireStatus(t, resp, http.StatusOK)
	requireListContainsID(t, nestedSlice(t, resp, "financial_obligations"), payableID)
	requireListContainsID(t, nestedSlice(t, resp, "financial_obligations"), receivableID)

	resp = c.do(t, http.MethodPost, "/financial-obligation/pay", map[string]any{
		"financial_obligation_id": payableID,
		"financial_account_id":    accountID,
		"amount_paid":             80,
		"payment_date":            "2026-07-12",
	})
	requireStatus(t, resp, http.StatusOK)

	if got := currentBalance(t, c, accountID); got != 420 {
		t.Fatalf("balance after partial payment = %.2f, want 420.00", got)
	}

	resp = c.do(t, http.MethodPost, "/financial-obligation/pay", map[string]any{
		"financial_obligation_id": payableID,
		"financial_account_id":    accountID,
		"amount_paid":             120,
		"payment_date":            "2026-07-12",
	})
	requireStatus(t, resp, http.StatusOK)

	resp = c.do(t, http.MethodPost, "/financial-obligation/pay", map[string]any{
		"financial_obligation_id": payableID,
		"financial_account_id":    accountID,
		"amount_paid":             1,
		"payment_date":            "2026-07-12",
	})
	requireStatus(t, resp, http.StatusUnprocessableEntity)

	resp = c.do(t, http.MethodPost, "/financial-obligation/pay", map[string]any{
		"financial_obligation_id": receivableID,
		"financial_account_id":    accountID,
		"amount_paid":             100,
		"payment_date":            "2026-07-12",
	})
	requireStatus(t, resp, http.StatusUnprocessableEntity)
}

func TestPendingBacklogFlowsWithoutRegisteredRoutes(t *testing.T) {
	t.Skip("backlog transferencias, recebimento de obrigacao receivable, consulta de liquidacoes e estornos/cancelamento de movimentacao ainda nao possuem rotas publicas registradas")
}

func createObligation(t *testing.T, c apiClient, categoryID int, obligationType string, amount float64) int {
	t.Helper()

	resp := c.do(t, http.MethodPost, "/financial-obligation", map[string]any{
		"category_id":     categoryID,
		"description":     uniqueName("backlog-obligation"),
		"type":            obligationType,
		"original_amount": amount,
		"due_date":        "2026-07-30",
		"competence_date": ptr("2026-07-01"),
		"notes":           ptr("created by flow tests"),
	})
	requireStatus(t, resp, http.StatusCreated)
	return idFrom(t, nestedMap(t, resp, "financial_obligation"))
}

func requireListContainsID(t *testing.T, list []any, id int) {
	t.Helper()

	for _, item := range list {
		body, ok := item.(map[string]any)
		if !ok {
			continue
		}

		itemID, ok := body["id"].(float64)
		if ok && int(itemID) == id {
			return
		}
	}

	t.Fatalf("list does not contain id %d: %#v", id, list)
}
