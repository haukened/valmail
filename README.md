# valmail
[![codecov](https://codecov.io/gh/haukened/valmail/branch/main/graph/badge.svg?token=ZBN2N8O8XV)](https://codecov.io/gh/haukened/valmail)

a GoLang email validator that allows projects to validate emails for free - instead of using pay-to-play SAAS options.

## Project Goals and Status

| Finished? | Feature | Description | Field / Function Name | Field Type |
|---|---|---|---|---|
| 🛠️ | Email Reachability | How confident are we in sending an email to this address? | IsReachable | Boolean |
| ✅ | Syntax Validation | Is the address syntactically valid? | ValidSyntax | Boolean |
| ✅ | MX Record Validation | Does the domain of the email have valid MX records in DNS? | MXValid | Boolean |
| ✅ | Disposable Email Address (DEA) Validation | Is the address provided by a known disposable email address provider? | IsDisposable | Boolean |
| 🛠️ | SMTP Server Validation | Is an email sent to this address deliverable? | SMTPValid | Boolean |
| 🛠️ | Mailbox Disabled | Has this address been disabled by the email provider? | IsDisabled | Boolean |
| 🛠️ | Full Inbox | Is the inbox of this mailbox full? | HasFullInbox | Boolean |
| 🛠️ | Catch-All Address | Is this email a catch all address? | IsCatchAll | Boolean |
| 🛠️ | Gravatar Validation | Does this email have a gravatar? | HasGravatar | Boolean |
| 🛠️ | Gravatar URL | The URL of the gravatar image for this email (if present) | GravatarURL | String |
| 🛠️ | Has this been Pwned? | Has this email been comprimised in a data breach? | Pwned | Boolean |
