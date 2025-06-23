package stripe

import (
	"context"
	"fmt"
	"os"

	"github.com/stripe/stripe-go/v82"
)

func InitStripe() *stripe.Client {
	stripeKey := os.Getenv("STRIPE_API_KEY")
	if stripeKey == "" {
		panic("STRIPE_API_KEY não definido")
	}

	sc := stripe.NewClient(stripeKey, nil)
	return sc
}

func CreateStripeCustomer(ctx context.Context, sc *stripe.Client, email string) (*stripe.Customer, error) {
	params := &stripe.CustomerCreateParams{
		Description:      stripe.String("Stripe Developer"),
		Email:            stripe.String(email),
		PreferredLocales: stripe.StringSlice([]string{"en", "es"}),
	}

	customer, err := sc.V1Customers.Create(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar cliente Stripe: %w", err)
	}
	return customer, nil
}

func ListCustomerPaymentIntents(ctx context.Context, sc *stripe.Client, customerID string) stripe.Seq2[*stripe.PaymentIntent, error] {
	params := &stripe.PaymentIntentListParams{
		Customer: stripe.String(customerID),
	}
	return sc.V1PaymentIntents.List(ctx, params)
}

func ListStripeEvents(ctx context.Context, sc *stripe.Client) stripe.Seq2[*stripe.Event, error] {
	return sc.V1Events.List(ctx, nil)
}

// func a() {
// 	ctx := context.Background()
// 	sc := stripe.InitStripe()

// 	// 1. Cria cliente
// 	customer, err := stripe.CreateStripeCustomer(ctx, sc, "client@email.com")
// 	if err != nil {
// 		log.Fatal("Erro ao criar cliente:", err)
// 	}

// 	// 2. Cria assinatura
// 	subscription, err := stripe.CreateSubscription(ctx, sc, customer.ID, "price_1PEgTgF2...")
// 	if err != nil {
// 		log.Fatal("Erro ao criar assinatura:", err)
// 	}

// 	fmt.Println("Assinatura criada:", subscription.ID)
// }
