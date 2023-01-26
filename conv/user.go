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

func ToPasswordDtoFromProto(input *pb.Password) (dto.Password, error) {
	if input == nil {
		return dto.Password{}, nil
	}

	createdAt := utils.TimestampPbAsTimeNow(input.GetCreatedAt()).Local()
	updatedAt := utils.TimestampPbAsTimeNow(input.GetUpdatedAt()).Local()
	output := dto.Password{
		ID:        input.GetId(),
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		UserID:    input.GetUserID(),
		Value:     input.GetValue(),
	}

	return output, nil
}

func ToUserDtoFromProto(input *pb.User) (dto.User, error) {
	if input == nil {
		return dto.User{}, nil
	}

	createdAt := utils.TimestampPbAsTimeNow(input.GetCreatedAt()).Local()
	updatedAt := utils.TimestampPbAsTimeNow(input.GetUpdatedAt()).Local()
	password, err := ToPasswordDtoFromProto(input.GetPassword())
	if err != nil {
		return dto.User{}, err
	}
	personal, err := ToPersonalDtoFromProto(input.GetPersonal())
	if err != nil {
		return dto.User{}, err
	}
	emails, err := ToEmailListDtoFromProto(input.GetEmails())
	if err != nil {
		return dto.User{}, err
	}

	output := dto.User{
		ID:          input.GetId(),
		CreatedAt:   &createdAt,
		UpdatedAt:   &updatedAt,
		LoginID:     input.GetLogindID(),
		LoginType:   input.GetLoginType(),
		LoginSource: input.GetLoginSource(),

		Password: password,
		Personal: personal,
		Emails:   emails,
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

func ToPersonalProtoFromDto(input dto.Personal) (*pb.Personal, error) {

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

func ToPasswordProtoFromDto(input dto.Password) (*pb.Password, error) {

	output := pb.Password{
		Id:        input.ID,
		CreatedAt: utils.NewTimestampPB(input.CreatedAt),
		UpdatedAt: utils.NewTimestampPB(input.UpdatedAt),
		UserID:    input.UserID,
		Value:     input.Value,
	}

	return &output, nil
}

func ToUserProtoFromDto(input dto.User) (*pb.User, error) {
	password, _ := ToPasswordProtoFromDto(input.Password)
	personal, _ := ToPersonalProtoFromDto(input.Personal)
	emails, _ := ToEmailListProtoFromdto(input.Emails)

	output := pb.User{
		Id:          input.ID,
		CreatedAt:   utils.NewTimestampPB(input.CreatedAt),
		UpdatedAt:   utils.NewTimestampPB(input.UpdatedAt),
		LogindID:    input.LoginID,
		LoginType:   input.LoginType,
		LoginSource: input.LoginSource,
		Password:    password,
		Personal:    personal,
		Emails:      emails,
	}

	return &output, nil
}
