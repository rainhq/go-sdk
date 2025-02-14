// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import "lightspark/types"

type SendPaymentOutput struct {

	// The payment that has been sent.
	Payment types.EntityWrapper `json:"send_payment_output_payment"`
}

const (
	SendPaymentOutputFragment = `
fragment SendPaymentOutputFragment on SendPaymentOutput {
    __typename
    send_payment_output_payment: payment {
        id
    }
}
`
)
