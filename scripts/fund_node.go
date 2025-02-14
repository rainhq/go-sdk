// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package scripts

import "lightspark/objects"

const FUND_NODE_MUTATION = `
mutation FundNode(
    $node_id: ID!,
    $amount_sats: Long
) {
    fund_node(input: { node_id: $node_id, amount_sats: $amount_sats }) {
        amount {
            ...CurrencyAmountFragment
        }
    }
}

` + objects.CurrencyAmountFragment
