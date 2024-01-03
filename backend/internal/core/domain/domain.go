package domain

type Config struct {
	AppPort         string `env:"APP_PORT"`
	AppName         string `env:"APP_NAME"`
	AppHost         string `env:"APP_HOST"`
	CognitoClientId string `env:"COGNITO_CLIENT_ID"`
}

type EmptyResponseDTO struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message,omitempty" example:"Operação realizada com sucesso"`
	Code    int    `json:"code,omitempty" example:"200"`
} //	@name	EmptyResponseDTO

type ResponseDTO[T any] struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message,omitempty" example:"Operação realizada com sucesso"`
	Code    int    `json:"code,omitempty" example:"200"`
	Data    *T     `json:"data,omitempty"`
} //	@name	ResponseDTO

type UserToken struct {
	RefreshToken  *string `json:"RefreshToken,omitempty"`
	AccessToken   *string `json:"AccessToken,omitempty"`
	AuthChallenge *string `json:"AuthChallenge,omitempty"`
	AuthSession   *string `json:"AuthSession,omitempty"`
} // @name	UserToken

type UserLoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
} // @name	UserLoginDTO

type UserLoginChangeDTO struct {
	Email       string `json:"email" validate:"required,email"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
	Session     string `json:"session" validate:"required"`
} // @name	UserLoginChangeDTO
