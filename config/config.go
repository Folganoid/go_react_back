package config

func Setup() map[string]string {

	setup := make(map[string]string)

	// mysql
	//setup["db_name"] = "go"
	//setup["db_user"] = "fg"
	//setup["db_pass"] = "56195619"
	//setup["db_prefix"] = ""

	//postgres
	setup["db_host"] = "127.0.0.1"
	setup["db_port"] = "5432"
	setup["db_pass"] = "5619"
	setup["db_user"] = "fg"
	setup["db_pass"] = "56195619"
	setup["db_name"] = "postgres"
	setup["db_prefix"] = "test."

	setup["frontPath"] = "http://127.0.0.1:3000"

	return setup
}
