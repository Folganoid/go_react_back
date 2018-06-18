package config

func Setup() map[string]string {

	setup := make(map[string]string)

	setup["db_name"] = "go"
	setup["db_user"] = "fg"
	setup["db_pass"] = "5619"


	return setup
}