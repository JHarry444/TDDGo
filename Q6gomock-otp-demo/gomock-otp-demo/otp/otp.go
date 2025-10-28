package otp

// OTPService defines external dependency for sending and verifying OTPs.
type OTPService interface {
	SendOTP(email string) (string, error)
	VerifyOTP(email, otp string) bool
}

// AuthManager is the CUT (Code Under Test).
type AuthManager struct {
	service OTPService
	store   map[string]string // in-memory: email -> last sent OTP
}

func NewAuthManager(svc OTPService) *AuthManager {
	return &AuthManager{
		service: svc,
		store:   make(map[string]string),
	}
}

// SendOTP asks the service to send an OTP and stores the value per email.
func (a *AuthManager) SendOTP(email string) error {
	otp, err := a.service.SendOTP(email)
	if err != nil {
		return err
	}
	a.store[email] = otp
	return nil
}

// VerifyOTP returns true only if service verifies AND stored value matches.
func (a *AuthManager) VerifyOTP(email, otp string) bool {
	stored, ok := a.store[email]
	if !ok {
		return false
	}
	if stored != otp {
		return false
	}
	return a.service.VerifyOTP(email, otp)
}
