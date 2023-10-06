package auth

const ServicePath = "/keeper.keeperService/"

func AccessibleRoles() map[string][]string {

	return map[string][]string{
		ServicePath + "CreateAccount": {"admin", "user"},
		ServicePath + "UpdateAccount": {"user"},
		ServicePath + "DeleteAccount": {"user"},
		ServicePath + "ListAccount":   {"user"},
		ServicePath + "CreateText":    {"user"},
		ServicePath + "UpdateText":    {"user"},
		ServicePath + "DeleteText":    {"user"},
		ServicePath + "ListText":      {"user"},
		ServicePath + "CreateBinary":  {"user"},
		ServicePath + "UpdateBinary":  {"user"},
		ServicePath + "DeleteBinary":  {"user"},
		ServicePath + "ListBinary":    {"user"},
		ServicePath + "CreateCard":    {"user"},
		ServicePath + "UpdateCard":    {"user"},
		ServicePath + "DeleteCard":    {"user"},
		ServicePath + "ListCard":      {"user"},
	}
}
