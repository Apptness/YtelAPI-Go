/*
 * ytelapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package email_pkg

import "ytelapi_lib/models_pkg"
import "ytelapi_lib/configuration_pkg"

/*
 * Interface for the EMAIL_IMPL
 */
type EMAIL interface {
    CreateBlockedEmails (*string, *string) (string, error)

    CreateRemoveInvalidEmail (string) (string, error)

    CreateInvalidEmails (*string, *string) (string, error)

    CreateRemoveBouncedEmail (string) (string, error)

    CreateBouncedEmails (*string, *string) (string, error)

    CreateSpamEmails (*string, *string) (string, error)

    CreateSendEmail (string, models_pkg.TypeEnum, string, string, *string, *string, *string, *string) (string, error)

    CreateRemoveBlockedAddress (string) (string, error)

    AddEmailUnsubscribe (string) (string, error)

    CreateRemoveUnsubscribedEmail (string) (string, error)

    CreateListUnsubscribedEmails (*string, *string) (string, error)

    CreateRemoveSpamAddress (string) (string, error)
}

/*
 * Factory for the EMAIL interaface returning EMAIL_IMPL
 */
func NewEMAIL(config configuration_pkg.CONFIGURATION) *EMAIL_IMPL {
    client := new(EMAIL_IMPL)
    client.config = config
    return client
}