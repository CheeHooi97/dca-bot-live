package errcode

type ErrorCode struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var (
	InternalServerError = ErrorCode{
		Code:    "INTERNAL_ERROR",
		Message: "Internal Server Error",
	}
	RegisteredEmail = ErrorCode{
		Code:    "REGISTERED_EMAIL",
		Message: "Email has been registered",
	}
	RegisteredPhoneNumber = ErrorCode{
		Code:    "REGISTERED_PHONE_NUMBER",
		Message: "Phone number has been registered",
	}
	InvalidRequest = ErrorCode{
		Code:    "INVALID_REQUEST",
		Message: "Invalid Request",
	}
	InvalidOtp = ErrorCode{
		Code:    "INVALID_OTP",
		Message: "Invalid otp",
	}
	InvalidTAC = ErrorCode{
		Code:    "INVALID_TAC",
		Message: "Invalid TAC",
	}
	InvalidDate = ErrorCode{
		Code:    "INVALID_DATE",
		Message: "Invalid date",
	}
	ValidationError = ErrorCode{
		Code:    "VALIDATION_ERROR",
		Message: "Validation error",
	}
	EncryptionError = ErrorCode{
		Code:    "ENCRYPTION_ERROR",
		Message: "Encryption error",
	}
	FileError = ErrorCode{
		Code:    "FILE_ERROR",
		Message: "File error",
	}
	RedisError = ErrorCode{
		Code:    "REDIS_ERROR",
		Message: "Redis error",
	}
	SamePinError = ErrorCode{
		Code:    "SAME_PIN_ERROR",
		Message: "New pin cannot same as current pin",
	}
	TransactionFailed = ErrorCode{
		Code:    "TRANSACTION_FAILED",
		Message: "Transaction failed",
	}
	RegisteredSocialUId = ErrorCode{
		Code:    "REGISTERED_SOCIAL_UID",
		Message: "Social UId has been registered",
	}
	AuthenticationFailed = ErrorCode{
		Code:    "AUTHENTICATION_FAILED",
		Message: "Failed to authenticate",
	}
	InvalidEncryptedText = ErrorCode{
		Code:    "INVALID_ENCRYPTED_TEXT",
		Message: "Invalid encrypted text",
	}
	InvalidRegistration = ErrorCode{
		Code:    "INVALID_REGISTRATION",
		Message: "Invalid registration process",
	}
	InvalidToken = ErrorCode{
		Code:    "INVALID_TOKEN",
		Message: "Token is invalid",
	}
	InvalidRole = ErrorCode{
		Code:    "INVALID_ROLE",
		Message: "Invalid role",
	}
	InvalidLocation = ErrorCode{
		Code:    "INVALID_LOCATION",
		Message: "Current location is not nearby the destination area",
	}
	InvalidOwner = ErrorCode{
		Code:    "INVALID_OWNER",
		Message: "Invalid Owner",
	}
	InvalidTenant = ErrorCode{
		Code:    "INVALID_TENANT",
		Message: "Invalid tenant",
	}
	InvalidAuthorizePin = ErrorCode{
		Code:    "INVALID_AUTHORIZE_PIN",
		Message: "Invalid authorize pin",
	}
	InvalidSetAuthorizePin = ErrorCode{
		Code:    "INVALID_SET_AUTHORIZE_PIN",
		Message: "Authorize pin has been set",
	}
	InvalidUrl = ErrorCode{
		Code:    "INVALID_URL",
		Message: "Invalid url",
	}
	InvalidStage = ErrorCode{
		Code:    "INVALID_STAGE",
		Message: "Invalid stage",
	}
	TokenNotFound = ErrorCode{
		Code:    "TOKEN_NOT_FOUND",
		Message: "Token not found",
	}
	UserNotFound = ErrorCode{
		Code:    "USER_NOT_FOUND",
		Message: "User not found",
	}
	ActorNotFound = ErrorCode{
		Code:    "ACTOR_NOT_FOUND",
		Message: "Actor not found",
	}
	EkycNotFound = ErrorCode{
		Code:    "EKYC_NOT_FOUND",
		Message: "Ekyc not found",
	}
	BankNotFound = ErrorCode{
		Code:    "BANK_NOT_FOUND",
		Message: "Bank not found",
	}
	DepositNotFound = ErrorCode{
		Code:    "DEPOSIT_NOT_FOUND",
		Message: "Deposit not found",
	}
	IdNotFound = ErrorCode{
		Code:    "ID_NOT_FOUND",
		Message: "Id not found",
	}
	InfoNotFound = ErrorCode{
		Code:    "INFO_NOT_FOUND",
		Message: "Info not found",
	}
	AppointmentNotFound = ErrorCode{
		Code:    "APPOINTMENT_NOT_FOUND",
		Message: "Appointment not found",
	}
	HomeNotFound = ErrorCode{
		Code:    "HOME_NOT_FOUND",
		Message: "Home not found",
	}
	HomeInfoNotFound = ErrorCode{
		Code:    "HOME_INFO_NOT_FOUND",
		Message: "Home info not found",
	}
	SettingNotFound = ErrorCode{
		Code:    "SETTING_NOT_FOUND",
		Message: "Setting not found",
	}
	ContractNotFound = ErrorCode{
		Code:    "CONTRACT_NOT_FOUND",
		Message: "Contract not found",
	}
	StageNotFound = ErrorCode{
		Code:    "STAGE_NOT_FOUND",
		Message: "Stage not found",
	}
	TransactionNotFound = ErrorCode{
		Code:    "TRANSACTION_NOT_FOUND",
		Message: "Transaction not found",
	}
	PaymentRequestNotFound = ErrorCode{
		Code:    "PAYMENT_REQUEST_NOT_FOUND",
		Message: "Payment Request not found",
	}
	FailedGetUser = ErrorCode{
		Code:    "FAILED_GET_USER",
		Message: "Failed to get user",
	}
	RegisteredUsername = ErrorCode{
		Code:    "REGISTERED_USERNAME",
		Message: "Username has been used",
	}
	AppointmentSameUserExisted = ErrorCode{
		Code:    "APPOINTMENT_SAME_USER_EXISTED",
		Message: "Same user cannot create appointment for his/her own property",
	}
	AppointmentHomeExisted = ErrorCode{
		Code:    "APPOINTMENT_EXISTED_HOME",
		Message: "Failed to make appointment as there is existed appointment for same home/property",
	}
	HomeEditError = ErrorCode{
		Code:    "HOME_EDIT_ERROR",
		Message: "Cannot edit home details unless there is no offer",
	}
	HomePublishError = ErrorCode{
		Code:    "HOME_PUBLISH_ERROR",
		Message: "Home is not completed, not approved or rented",
	}
	HomePhotoError = ErrorCode{
		Code:    "HOME_PHOTO_ERROR",
		Message: "At least five photos are required",
	}
	HomeInfoError = ErrorCode{
		Code:    "HOME_INFO_ERROR",
		Message: "Home info is incorrect",
	}
	HomeError = ErrorCode{
		Code:    "HOME_ERROR",
		Message: "Home is not owned by this user",
	}
	ContractError = ErrorCode{
		Code:    "CONTRACT_ERROR",
		Message: "Contract is not existed",
	}
	ActiveRoleError = ErrorCode{
		Code:    "ACTIVE_ROLE_ERROR",
		Message: "Active role error",
	}
	AddressError = ErrorCode{
		Code:    "ADDRESS_ERROR",
		Message: "Address field is required",
	}
	PropertyTypeError = ErrorCode{
		Code:    "PROPERTY_TYPE_ERROR",
		Message: "Property Type field is required",
	}
	TitleError = ErrorCode{
		Code:    "TITLE_ERROR",
		Message: "Title field is required",
	}
	RoomError = ErrorCode{
		Code:    "ROOM_ERROR",
		Message: "Room field is required",
	}
	SquareFeetError = ErrorCode{
		Code:    "SQUARE_FEET_ERROR",
		Message: "Square Feet field is required",
	}
	PriceError = ErrorCode{
		Code:    "PRICE_ERROR",
		Message: "Price must be greater than 0",
	}
	TenurePeriodError = ErrorCode{
		Code:    "TENURE_PERIOD_ERROR",
		Message: "Tenure Period must choose at least one",
	}
	OfferSameHomeExisted = ErrorCode{
		Code:    "OFFER_SAME_HOME_EXISTED",
		Message: "Offer for the same home/property existed",
	}
	InvalidHome = ErrorCode{
		Code:    "INVALID_HOME",
		Message: "Home is unpublished or rented",
	}
	InvalidHomeSameUser = ErrorCode{
		Code:    "INVALID_HOME_SAME_USER",
		Message: "The user is same as the home owner",
	}
	InvalidTenure = ErrorCode{
		Code:    "INVALID_TENURE",
		Message: "Invalid tenure period",
	}
	InvalidSignature = ErrorCode{
		Code:    "INVALID_SIGNATURE",
		Message: "Invalid signature",
	}
	PhotoUrlFieldRequired = ErrorCode{
		Code:    "PHOTO_URL_FIELD_REQUIRED",
		Message: "PhotoUrl Field is required",
	}
	TACAttemptReached = ErrorCode{
		Code:    "TAC_ATTEMPT_REACHED",
		Message: "Your account has reached the maximum attempt limit within this time frame, please try again later.",
	}
	AmountNotTally = ErrorCode{
		Code:    "AMOUNT_NOT_TALLY",
		Message: "Amount is not tally",
	}
)
