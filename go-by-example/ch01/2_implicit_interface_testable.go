// implicit_interface_testable.go
//
// This example demonstrates how Go's implicit interfaces
// make code easily testable - even when the external API
// (like a payment processor) does not define its own interfaces.
//
// We define a small interface (Charger) that represents the
// behavior we need, and then write code that depends on that interface.
// In production, we pass the real implementation.
// In tests, we pass a fake one.

package main

import (
	"fmt"
)

// Charger code only cares that something can "Charge" a user.
// It doesn’t matter whether it’s a real API client or a fake one.
type Charger interface {
	Charge(userID string, amount int) error
}

// MithrilClient Production implementation serving a real API
type MithrilClient struct{}

func (m *MithrilClient) Charge(userID string, amount int) error {
	fmt.Printf("[Mithril] Charging user %s %d\n", userID, amount)
	return nil
}

func ProcessPayment(c Charger, userId string, amount int) error {
	fmt.Println("Processing payment...")
	return c.Charge(userId, amount)
}

// FakeCharger satisfies the Charger interface implicitly
// because it has a Charge method with the same signature.
type FakeCharger struct {
	Calls []string
}

func (f *FakeCharger) Charge(userId string, amount int) error {
	fmt.Printf("[Fake] Pretending to charge user %s %d\n", userId, amount)
	f.Calls = append(f.Calls, fmt.Sprintf("%s:%d", userId, amount))
	return nil
}

func testProcessPayment() {
	fake := &FakeCharger{}
	_ = ProcessPayment(fake, "johnDoe", 100)
}

func main() {
	// Production: use the real API clien
	client := &MithrilClient{}
	_ = ProcessPayment(client, "user12", 50)

	fmt.Println()

	// Testing: use a fake that doesn’t hit real servers
	testProcessPayment()
}
