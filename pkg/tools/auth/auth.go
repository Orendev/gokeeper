package auth

const ServicePath = "/keeper.keeperService/"

func AccessibleRoles() map[string][]string {

	return map[string][]string{
		ServicePath + "CreateAccount": {"admin", "user"},
		ServicePath + "UpdateAccount": {"user"},
		ServicePath + "ListAccount":   {"user"},
	}
}
