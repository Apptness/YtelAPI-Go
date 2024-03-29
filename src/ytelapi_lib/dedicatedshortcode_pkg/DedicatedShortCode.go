/*
 * ytelapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package dedicatedshortcode_pkg

import "ytelapi_lib/configuration_pkg"

/*
 * Interface for the DEDICATEDSHORTCODE_IMPL
 */
type DEDICATEDSHORTCODE interface {
    UpdateShortcode (string, *string, *string, *string, *string, *string) (string, error)

    CreateListShortcodes (*string, *string, *string) (string, error)

    CreateListInboundSMS (*int64, *int64, *string, *string, *string) (string, error)

    CreateViewSMS (string) (string, error)

    CreateListSMS (*string, *string, *string, *int64, *int64) (string, error)

    CreateSendSMS (int64, float64, string, *string, *string) (string, error)

    CreateViewSMS (string) (string, error)
}

/*
 * Factory for the DEDICATEDSHORTCODE interaface returning DEDICATEDSHORTCODE_IMPL
 */
func NewDEDICATEDSHORTCODE(config configuration_pkg.CONFIGURATION) *DEDICATEDSHORTCODE_IMPL {
    client := new(DEDICATEDSHORTCODE_IMPL)
    client.config = config
    return client
}