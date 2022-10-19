package service

type AuthService struct {
	AllowedTokens    []string
	RegisteredTokens []string
}

type AuthServiceRequest struct {
	Token string
}

type AuthServiceResponse struct {
	Status  string
	Message string
}

func (t *AuthService) LogIn(request AuthServiceRequest, response *AuthServiceResponse) error {
	response.Status = "fail"
	if request.Token == "" {
		response.Message = "Missing token"
	}
	// Check if logged
	for _, registeredToken := range t.RegisteredTokens {
		if registeredToken == request.Token {
			response.Status = "fail"
			response.Message = "Token registered"
			return nil
		}
	}
	// Check if allowed
	for _, allowedToken := range t.AllowedTokens {
		if allowedToken == request.Token {
			response.Status = "success"
			t.RegisteredTokens = append(t.RegisteredTokens, request.Token)
			return nil
		}
	}
	response.Message = "Unknown token"
	return nil
}

func (t *AuthService) LogOut(request AuthServiceRequest, response *AuthServiceResponse) error {
	response.Status = "fail"
	for idx, registeredToken := range t.RegisteredTokens {
		if registeredToken == request.Token {
			response.Status = "success"
			t.RegisteredTokens = append(t.RegisteredTokens[:idx], t.RegisteredTokens[idx+1:]...)
			return nil
		}
	}
	return nil
}
