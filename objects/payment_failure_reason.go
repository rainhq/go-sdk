// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects

import (
	"encoding/json"
)

type PaymentFailureReason int

const (
	PaymentFailureReasonUndefined PaymentFailureReason = iota

	PaymentFailureReasonNone

	PaymentFailureReasonTimeout

	PaymentFailureReasonNoRoute

	PaymentFailureReasonError

	PaymentFailureReasonIncorrectPaymentDetails

	PaymentFailureReasonInsufficientBalance

	PaymentFailureReasonInvoiceAlreadyPaid

	PaymentFailureReasonSelfPayment

	PaymentFailureReasonInvoiceExpired
)

func (a *PaymentFailureReason) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		*a = PaymentFailureReasonUndefined
	case "NONE":
		*a = PaymentFailureReasonNone
	case "TIMEOUT":
		*a = PaymentFailureReasonTimeout
	case "NO_ROUTE":
		*a = PaymentFailureReasonNoRoute
	case "ERROR":
		*a = PaymentFailureReasonError
	case "INCORRECT_PAYMENT_DETAILS":
		*a = PaymentFailureReasonIncorrectPaymentDetails
	case "INSUFFICIENT_BALANCE":
		*a = PaymentFailureReasonInsufficientBalance
	case "INVOICE_ALREADY_PAID":
		*a = PaymentFailureReasonInvoiceAlreadyPaid
	case "SELF_PAYMENT":
		*a = PaymentFailureReasonSelfPayment
	case "INVOICE_EXPIRED":
		*a = PaymentFailureReasonInvoiceExpired

	}
	return nil
}

func (a PaymentFailureReason) StringValue() string {
	var s string
	switch a {
	default:
		s = "undefined"
	case PaymentFailureReasonNone:
		s = "NONE"
	case PaymentFailureReasonTimeout:
		s = "TIMEOUT"
	case PaymentFailureReasonNoRoute:
		s = "NO_ROUTE"
	case PaymentFailureReasonError:
		s = "ERROR"
	case PaymentFailureReasonIncorrectPaymentDetails:
		s = "INCORRECT_PAYMENT_DETAILS"
	case PaymentFailureReasonInsufficientBalance:
		s = "INSUFFICIENT_BALANCE"
	case PaymentFailureReasonInvoiceAlreadyPaid:
		s = "INVOICE_ALREADY_PAID"
	case PaymentFailureReasonSelfPayment:
		s = "SELF_PAYMENT"
	case PaymentFailureReasonInvoiceExpired:
		s = "INVOICE_EXPIRED"

	}
	return s
}

func (a PaymentFailureReason) MarshalJSON() ([]byte, error) {
	s := a.StringValue()
	return json.Marshal(s)
}
