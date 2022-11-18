package interfaces


type UserValidationInterface interface{
	ValidateUserCreation(data User)error
}