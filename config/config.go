package config

func Setup() map[string]string {

	setup := make(map[string]string)

	//setup["db_name"] = "go"
	//setup["db_user"] = "t430"
	//setup["db_pass"] = "56195619"

	setup["db_name"] = "go"
	setup["db_user"] = "fg"
	setup["db_pass"] = "56195619"

	setup["frontPath"] = "http://127.0.0.1:3000"

	return setup
}
