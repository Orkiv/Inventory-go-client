/* 
 * InventoryAPI
 *
 * Orkiv Inventory API client 
 *
 * OpenAPI spec version: 1.0.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package InventoryClient

type Order struct {

	// Order ID
	OrderId string `json:"order_id,omitempty"`

	// Customer email
	InfoEmail string `json:"info_email,omitempty"`

	// Customer first name
	InfoFirst string `json:"info_first,omitempty"`

	// Customer last name
	InfoLast string `json:"info_last,omitempty"`

	// Customer phone number
	Phone string `json:"phone,omitempty"`

	// Customer billing address matches shipping address
	Shipset bool `json:"shipset,omitempty"`

	// Customer billing address line '1'
	InfoAdr1 string `json:"info_adr1,omitempty"`

	// Customer billing address line '2'
	InfoAdr2 string `json:"info_adr2,omitempty"`

	// Customer billing city
	InfoCty string `json:"info_cty,omitempty"`

	// Customer billing zip code
	InfoZip string `json:"info_zip,omitempty"`

	// Customer billing state
	State string `json:"state,omitempty"`

	// Customer shipping address line '1'
	InfoSadr1 string `json:"info_sadr1,omitempty"`

	// Customer shipping address line '2'
	InfoSadr2 string `json:"info_sadr2,omitempty"`

	// Customer shipping city
	InfoScty string `json:"info_scty,omitempty"`

	// Customer shipping zip code
	InfoSzip string `json:"info_szip,omitempty"`

	// Customer shipping state
	Sstate string `json:"sstate,omitempty"`

	// Tax amount in hundreds. Must divide by '100' for USD value
	TaxAmount float32 `json:"tax_amount,omitempty"`

	// Shipping total in USD
	ShippingAmount float32 `json:"shipping_amount,omitempty"`

	// Checkout total in USD
	AmountTotal float32 `json:"amount_total,omitempty"`

	// Array of items purchased at checkout
	ItemIDs []string `json:"itemIDs,omitempty"`
}
