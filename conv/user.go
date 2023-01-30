package conv

import (
	"github.com/w-woong/common/dto"
	pb "github.com/w-woong/common/dto/protos/user/v1"
	"github.com/w-woong/common/utils"
)

func ToEmailDtoFromProto(input *pb.Email) (dto.Email, error) {
	if input == nil {
		return dto.Email{}, nil
	}

	createdAt := utils.TimestampPbAsTimeNow(input.GetCreatedAt()).Local()
	updatedAt := utils.TimestampPbAsTimeNow(input.GetUpdatedAt()).Local()
	output := dto.Email{
		ID:        input.GetId(),
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		UserID:    input.GetUserID(),
		Email:     input.GetEmail(),
		Priority:  uint8(input.GetPriority()),
	}

	return output, nil
}

func ToEmailListDtoFromProto(input []*pb.Email) (dto.Emails, error) {
	if input == nil {
		return nil, nil
	}

	var err error
	output := make(dto.Emails, len(input))
	for i, p := range input {
		output[i], err = ToEmailDtoFromProto(p)
		if err != nil {
			return nil, err
		}
	}

	return output, nil
}

func ToPersonalDtoFromProto(input *pb.Personal) (dto.Personal, error) {
	if input == nil {
		return dto.Personal{}, nil
	}

	createdAt := utils.TimestampPbAsTimeNow(input.GetCreatedAt()).Local()
	updatedAt := utils.TimestampPbAsTimeNow(input.GetUpdatedAt()).Local()
	birthDate := input.GetBirthDate().AsTime().Local()
	output := dto.Personal{
		ID:          input.GetId(),
		CreatedAt:   &createdAt,
		UpdatedAt:   &updatedAt,
		UserID:      input.GetUserID(),
		FirstName:   input.GetFirstName(),
		MiddleName:  input.GetMiddleName(),
		LastName:    input.GetLastName(),
		BirthYear:   int(input.GetBirthYear()),
		BirthMonth:  int(input.GetBirthMonth()),
		BirthDay:    int(input.GetBirthDay()),
		BirthDate:   &birthDate,
		Gender:      input.GetGender(),
		Nationality: input.GetNationality(),
	}

	return output, nil
}

func ToCredentialPasswordDtoFromProto(input *pb.CredentialPassword) (dto.CredentialPassword, error) {
	if input == nil {
		return dto.CredentialPassword{}, nil
	}

	createdAt := utils.TimestampPbAsTimeNow(input.GetCreatedAt()).Local()
	updatedAt := utils.TimestampPbAsTimeNow(input.GetUpdatedAt()).Local()
	output := dto.CredentialPassword{
		ID:        input.GetId(),
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		UserID:    input.GetUserID(),
		Value:     input.GetValue(),
	}

	return output, nil
}

func ToCredentialTokenDtoFromProto(input *pb.CredentialToken) (dto.CredentialToken, error) {
	if input == nil {
		return dto.CredentialToken{}, nil
	}

	createdAt := utils.TimestampPbAsTimeNow(input.GetCreatedAt()).Local()
	updatedAt := utils.TimestampPbAsTimeNow(input.GetUpdatedAt()).Local()
	output := dto.CredentialToken{
		ID:        input.GetId(),
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		UserID:    input.GetUserID(),
		Value:     input.GetValue(),
	}

	return output, nil
}

func ToDeliveryRequestTypeDtoFromProto(input *pb.DeliveryRequestType) (dto.DeliveryRequestType, error) {
	if input == nil {
		return dto.DeliveryRequestType{}, nil
	}

	createdAt := utils.TimestampPbAsTimeNow(input.GetCreatedAt()).Local()
	updatedAt := utils.TimestampPbAsTimeNow(input.GetUpdatedAt()).Local()
	output := dto.DeliveryRequestType{
		ID:        input.GetId(),
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		Name:      input.GetName(),
	}

	return output, nil
}

func ToDeliveryRequestDtoFromProto(input *pb.DeliveryRequest) (dto.DeliveryRequest, error) {
	if input == nil {
		return dto.DeliveryRequest{}, nil
	}

	createdAt := utils.TimestampPbAsTimeNow(input.GetCreatedAt()).Local()
	updatedAt := utils.TimestampPbAsTimeNow(input.GetUpdatedAt()).Local()
	deliveryRequestType, err := ToDeliveryRequestTypeDtoFromProto(input.GetDeliveryRequestType())
	if err != nil {
		return dto.NilDeliveryRequest, err
	}
	output := dto.DeliveryRequest{
		ID:                    input.GetId(),
		CreatedAt:             &createdAt,
		UpdatedAt:             &updatedAt,
		DeliveryAddressID:     input.GetDeliveryAddressID(),
		DeliveryRequestTypeID: input.GetDeliveryRequestTypeID(),
		DeliveryRequestType:   deliveryRequestType,
		RequestMessage:        input.GetRequestMessage(),
	}

	return output, nil
}

func ToDeliveryAddressDtoFromProto(input *pb.DeliveryAddress) (dto.DeliveryAddress, error) {
	if input == nil {
		return dto.DeliveryAddress{}, nil
	}

	createdAt := utils.TimestampPbAsTimeNow(input.GetCreatedAt()).Local()
	updatedAt := utils.TimestampPbAsTimeNow(input.GetUpdatedAt()).Local()
	deliveryRequest, err := ToDeliveryRequestDtoFromProto(input.GetDeliveryRequest())
	if err != nil {
		return dto.NilDeliveryAddress, err
	}
	output := dto.DeliveryAddress{
		ID:              input.GetId(),
		CreatedAt:       &createdAt,
		UpdatedAt:       &updatedAt,
		UserID:          input.GetUserID(),
		IsDefault:       input.GetIsDefault(),
		ReceiverName:    input.GetReceiverName(),
		ReceiverContact: input.ReceiverContact,
		PostCode:        input.PostCode,
		Address:         input.Address,
		AddressDetail:   input.AddressDetail,
		DeliveryRequest: deliveryRequest,
	}

	return output, nil
}

func ToPaymentTypeDtoFromProto(input *pb.PaymentType) (dto.PaymentType, error) {
	if input == nil {
		return dto.PaymentType{}, nil
	}

	createdAt := utils.TimestampPbAsTimeNow(input.GetCreatedAt()).Local()
	updatedAt := utils.TimestampPbAsTimeNow(input.GetUpdatedAt()).Local()
	output := dto.PaymentType{
		ID:        input.GetId(),
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		Name:      input.GetName(),
	}

	return output, nil
}

func ToPaymentMethodDtoFromProto(input *pb.PaymentMethod) (dto.PaymentMethod, error) {
	if input == nil {
		return dto.PaymentMethod{}, nil
	}

	createdAt := utils.TimestampPbAsTimeNow(input.GetCreatedAt()).Local()
	updatedAt := utils.TimestampPbAsTimeNow(input.GetUpdatedAt()).Local()
	paymentType, err := ToPaymentTypeDtoFromProto(input.GetPaymentType())
	if err != nil {
		return dto.NilPaymentMethod, err
	}
	output := dto.PaymentMethod{
		ID:            input.GetId(),
		CreatedAt:     &createdAt,
		UpdatedAt:     &updatedAt,
		UserID:        input.GetUserID(),
		PaymentTypeID: input.GetPaymentTypeID(),
		PaymentType:   paymentType,
		Identity:      input.GetIdentity(),
		Option:        input.GetOption(),
	}

	return output, nil
}

func ToUserDtoFromProto(input *pb.User) (dto.User, error) {
	if input == nil {
		return dto.User{}, nil
	}

	createdAt := utils.TimestampPbAsTimeNow(input.GetCreatedAt()).Local()
	updatedAt := utils.TimestampPbAsTimeNow(input.GetUpdatedAt()).Local()
	credentialPassword, err := ToCredentialPasswordDtoFromProto(input.GetCredentialPassword())
	if err != nil {
		return dto.NilUser, err
	}
	credentialToken, err := ToCredentialTokenDtoFromProto(input.GetCredentialToken())
	if err != nil {
		return dto.NilUser, err
	}
	personal, err := ToPersonalDtoFromProto(input.GetPersonal())
	if err != nil {
		return dto.NilUser, err
	}
	emails, err := ToEmailListDtoFromProto(input.GetEmails())
	if err != nil {
		return dto.NilUser, err
	}

	deliveryAddress, err := ToDeliveryAddressDtoFromProto(input.GetDeliveryAddress())
	if err != nil {
		return dto.NilUser, err
	}
	paymentMethod, err := ToPaymentMethodDtoFromProto(input.GetPaymentMethod())
	if err != nil {
		return dto.NilUser, err
	}

	output := dto.User{
		ID:          input.GetId(),
		CreatedAt:   &createdAt,
		UpdatedAt:   &updatedAt,
		LoginID:     input.GetLogindID(),
		LoginType:   input.GetLoginType(),
		LoginSource: input.GetLoginSource(),

		CredentialPassword: &credentialPassword,
		CredentialToken:    &credentialToken,
		Personal:           &personal,
		Emails:             emails,
		DeliveryAddress:    &deliveryAddress,
		PaymentMethod:      &paymentMethod,
	}

	return output, nil
}

func ToEmailProtoFromdto(input dto.Email) (*pb.Email, error) {

	output := pb.Email{
		Id:        input.ID,
		CreatedAt: utils.NewTimestampPB(input.CreatedAt),
		UpdatedAt: utils.NewTimestampPB(input.UpdatedAt),
		UserID:    input.UserID,
		Email:     input.Email,
		Priority:  uint32(input.Priority),
	}

	return &output, nil
}

func ToEmailListProtoFromdto(input dto.Emails) ([]*pb.Email, error) {
	if len(input) == 0 {
		return nil, nil
	}

	var err error
	output := make([]*pb.Email, len(input))
	for i, v := range input {
		output[i], err = ToEmailProtoFromdto(v)
		if err != nil {
			return nil, err
		}
	}

	return output, nil
}

func ToPersonalProtoFromDto(input *dto.Personal) (*pb.Personal, error) {
	if input == nil {
		return nil, nil
	}
	output := pb.Personal{
		Id:          input.ID,
		CreatedAt:   utils.NewTimestampPB(input.CreatedAt),
		UpdatedAt:   utils.NewTimestampPB(input.UpdatedAt),
		UserID:      input.UserID,
		FirstName:   input.FirstName,
		MiddleName:  input.MiddleName,
		LastName:    input.LastName,
		BirthYear:   int64(input.BirthYear),
		BirthMonth:  int64(input.BirthMonth),
		BirthDay:    int64(input.BirthDay),
		BirthDate:   utils.NewTimestampPB(input.BirthDate),
		Gender:      input.Gender,
		Nationality: input.Nationality,
	}

	return &output, nil
}

func ToCredentialPasswordProtoFromDto(input *dto.CredentialPassword) (*pb.CredentialPassword, error) {
	if input == nil {
		return nil, nil
	}
	output := pb.CredentialPassword{
		Id:        input.ID,
		CreatedAt: utils.NewTimestampPB(input.CreatedAt),
		UpdatedAt: utils.NewTimestampPB(input.UpdatedAt),
		UserID:    input.UserID,
		Value:     input.Value,
	}

	return &output, nil
}
func ToCredentialTokenProtoFromDto(input *dto.CredentialToken) (*pb.CredentialToken, error) {
	if input == nil {
		return nil, nil
	}
	output := pb.CredentialToken{
		Id:        input.ID,
		CreatedAt: utils.NewTimestampPB(input.CreatedAt),
		UpdatedAt: utils.NewTimestampPB(input.UpdatedAt),
		UserID:    input.UserID,
		Value:     input.Value,
	}

	return &output, nil
}

func ToDeliveryRequestTypeProtoFromDto(input dto.DeliveryRequestType) (*pb.DeliveryRequestType, error) {

	output := pb.DeliveryRequestType{
		Id:        input.ID,
		CreatedAt: utils.NewTimestampPB(input.CreatedAt),
		UpdatedAt: utils.NewTimestampPB(input.UpdatedAt),
		Name:      input.Name,
	}

	return &output, nil
}

func ToDeliveryRequestProtoFromDto(input dto.DeliveryRequest) (*pb.DeliveryRequest, error) {

	deliveryRequestType, err := ToDeliveryRequestTypeProtoFromDto(input.DeliveryRequestType)
	if err != nil {
		return nil, err
	}
	output := pb.DeliveryRequest{
		Id:                    input.ID,
		CreatedAt:             utils.NewTimestampPB(input.CreatedAt),
		UpdatedAt:             utils.NewTimestampPB(input.UpdatedAt),
		DeliveryAddressID:     input.DeliveryAddressID,
		DeliveryRequestTypeID: input.DeliveryRequestTypeID,
		DeliveryRequestType:   deliveryRequestType,
		RequestMessage:        input.RequestMessage,
	}

	return &output, nil
}

func ToDeliveryAddressProtoFromDto(input *dto.DeliveryAddress) (*pb.DeliveryAddress, error) {
	if input == nil {
		return nil, nil
	}
	deliveryRequest, err := ToDeliveryRequestProtoFromDto(input.DeliveryRequest)
	if err != nil {
		return nil, err
	}
	output := pb.DeliveryAddress{
		Id:              input.ID,
		CreatedAt:       utils.NewTimestampPB(input.CreatedAt),
		UpdatedAt:       utils.NewTimestampPB(input.UpdatedAt),
		UserID:          input.UserID,
		IsDefault:       input.IsDefault,
		ReceiverName:    input.ReceiverName,
		ReceiverContact: input.ReceiverContact,
		PostCode:        input.PostCode,
		Address:         input.Address,
		AddressDetail:   input.AddressDetail,
		DeliveryRequest: deliveryRequest,
	}

	return &output, nil
}

func ToPaymentTypeProtoFromDto(input dto.PaymentType) (*pb.PaymentType, error) {

	output := pb.PaymentType{
		Id:        input.ID,
		CreatedAt: utils.NewTimestampPB(input.CreatedAt),
		UpdatedAt: utils.NewTimestampPB(input.UpdatedAt),
		Name:      input.Name,
	}

	return &output, nil
}

func ToPaymentMethodProtoFromDto(input *dto.PaymentMethod) (*pb.PaymentMethod, error) {
	if input == nil {
		return nil, nil
	}

	paymentType, err := ToPaymentTypeProtoFromDto(input.PaymentType)
	if err != nil {
		return nil, err
	}
	output := pb.PaymentMethod{
		Id:            input.ID,
		CreatedAt:     utils.NewTimestampPB(input.CreatedAt),
		UpdatedAt:     utils.NewTimestampPB(input.UpdatedAt),
		UserID:        input.UserID,
		PaymentTypeID: input.PaymentTypeID,
		PaymentType:   paymentType,
		Identity:      input.Identity,
		Option:        input.Option,
	}

	return &output, nil
}

func ToUserProtoFromDto(input dto.User) (*pb.User, error) {
	password, err := ToCredentialPasswordProtoFromDto(input.CredentialPassword)
	if err != nil {
		return nil, err
	}
	token, err := ToCredentialTokenProtoFromDto(input.CredentialToken)
	if err != nil {
		return nil, err
	}
	personal, err := ToPersonalProtoFromDto(input.Personal)
	if err != nil {
		return nil, err
	}
	emails, err := ToEmailListProtoFromdto(input.Emails)
	if err != nil {
		return nil, err
	}
	deliveryAddress, err := ToDeliveryAddressProtoFromDto(input.DeliveryAddress)
	if err != nil {
		return nil, err
	}
	paymentMethod, err := ToPaymentMethodProtoFromDto(input.PaymentMethod)
	if err != nil {
		return nil, err
	}
	output := pb.User{
		Id:                 input.ID,
		CreatedAt:          utils.NewTimestampPB(input.CreatedAt),
		UpdatedAt:          utils.NewTimestampPB(input.UpdatedAt),
		LogindID:           input.LoginID,
		LoginType:          input.LoginType,
		LoginSource:        input.LoginSource,
		CredentialPassword: password,
		CredentialToken:    token,
		Personal:           personal,
		Emails:             emails,
		DeliveryAddress:    deliveryAddress,
		PaymentMethod:      paymentMethod,
	}

	return &output, nil
}
