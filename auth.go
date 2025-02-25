package fenrir

type LoginResponse struct {
	Result       string              `json:"result"`
	ID           string              `json:"_id"`
	UserID       string              `json:"user_id"`
	Token        string              `json:"token"`
	Name         string              `json:"name"`
	Subscription WebpushSubscription `json:"subscription"`
}

type Sessions struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}
type Account struct {
	ID    string `json:"_id"`
	Email string `json:"email"`
}

type Onboarding struct {
	Onboarding bool `json:"onboarding"`
}

type MFA struct {
	// Unvalidated or authorised MFA ticket; used to resolve the correct account
	MfaTicket string `json:"mfa_ticket"`

	// MFA response
	MfaResponse MFAResponse `json:"mfa_response"`

	// Friendly name used for the session
	FriendlyName string `json:"friendly_name"`
}

type MFAResponse struct {
	Password string `json:"password"`
}

type ChangeEmail struct {
	Ticket Ticket `json:"ticket"` // Why is this nested
}

type Ticket struct {
	ID           string `json:"_id"`
	AccountID    string `json:"account_id"`
	Token        string `json:"token"`
	Validated    bool   `json:"validated"`
	Authorised   bool   `json:"authorised"`
	LastTOTPCode string `json:"last_totp_code"`
}
