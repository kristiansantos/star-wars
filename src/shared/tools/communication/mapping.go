package communication

type mapItem struct {
	Message string
	Code    int
}

type ResponseMapping struct {
	Mapping map[string]mapItem `json:"mapping"`
}

var singletonResponseMapping *ResponseMapping = nil

//New ...
func New() ResponseMapping {
	if singletonResponseMapping == nil {
		mapping := ResponseMapping{}
		mapping.populate()
		singletonResponseMapping = &mapping
	}
	return *singletonResponseMapping
}

//Response responsible to make success response
func (response *ResponseMapping) Response(status int, identifier string, data interface{}) Response {
	return Response{
		Status:  status,
		Code:    response.Mapping[identifier].Code,
		Message: response.Mapping[identifier].Message,
		Data:    data,
	}
}

//Response responsible to make error response
func (response *ResponseMapping) ResponseError(status int, identifier string, err error) Response {
	return Response{
		Status:  status,
		Code:    response.Mapping[identifier].Code,
		Message: response.Mapping[identifier].Message,
		Error:   err,
	}
}

func (response *ResponseMapping) populate() {
	data := make(map[string]mapItem)

	data["already_exists"] = mapItem{Message: "Already exists", Code: 100000}
	data["validate_required"] = mapItem{Message: "Required", Code: 100001}
	data["validate_invalid"] = mapItem{Message: "Invalid", Code: 100002}
	data["validate_email"] = mapItem{Message: "Invalid e-mail", Code: 100003}
	data["validate_date"] = mapItem{Message: "Invalid date", Code: 100004}
	data["validate_password_length"] = mapItem{Message: "Password must be between 6 and 40 characters.", Code: 100005}
	data["success"] = mapItem{Message: "Success", Code: 100006}
	data["not_found"] = mapItem{Message: "Not found", Code: 100007}
	data["error"] = mapItem{Message: "Error", Code: 100008}
	data["success_create"] = mapItem{Message: "Record successfully created", Code: 100009}
	data["success_update"] = mapItem{Message: "Registro atualizado com sucesso", Code: 100010}
	data["success_delete"] = mapItem{Message: "Record successfully deleted", Code: 100011}
	data["error_create"] = mapItem{Message: "Unable to create record", Code: 100012}
	data["error_update"] = mapItem{Message: "Unable to update record", Code: 100013}
	data["error_delete"] = mapItem{Message: "Unable to delete record", Code: 100014}
	data["error_list"] = mapItem{Message: "Unable to list record", Code: 100015}
	data["authenticate_failed"] = mapItem{Message: "Authenticate failed", Code: 100016}
	data["authenticate_success"] = mapItem{Message: "Authenticate success", Code: 100017}
	data["validate_failed"] = mapItem{Message: "Validation failed", Code: 100018}
	data["endpoint_not_found"] = mapItem{Message: "Endpoint not found", Code: 100019}
	data["unexpected"] = mapItem{Message: "Unexpected", Code: 100020}
	data["error_extract_token"] = mapItem{Message: "Token not found", Code: 100021}
	data["error_check_token"] = mapItem{Message: "Invalid token", Code: 100022}
	data["user_not_found"] = mapItem{Message: "User not found", Code: 100023}
	data["email_send_successfully"] = mapItem{Message: "E-mail enviado com sucesso", Code: 100023}
	data["validate_new_password_to_old"] = mapItem{Message: "A nova senha não pode ser igual a senha atual", Code: 100024}
	data["validate_new_password_compare"] = mapItem{Message: "As senhas não coincidem", Code: 100025}
	data["invalid_integer"] = mapItem{Message: "Invalid integer", Code: 100026}
	data["greater_or_equal_than"] = mapItem{Message: "Deve ser maior ou igual à", Code: 100027}
	data["less_or_equal_than"] = mapItem{Message: "Deve ser menor ou igual à", Code: 100028}
	data["error_create_schedule"] = mapItem{Message: "Erro ao tentar criar o job no cloud schedule", Code: 100029}
	data["error_update_schedule"] = mapItem{Message: "Erro ao tentar atualizar o job no cloud schedule", Code: 100030}
	data["error_delete_schedule"] = mapItem{Message: "Erro ao tentar deletar o job no cloud schedule", Code: 100031}
	data["greater_than"] = mapItem{Message: "Deve ser maior à", Code: 100032}
	data["less_than"] = mapItem{Message: "Deve ser menor", Code: 100033}

	response.Mapping = data
}

func (item *mapItem) CustomMessage(message string) {
	item.Message = message
}
