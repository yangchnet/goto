package store

var entrys []*Entry = []*Entry{
	{
		Name:        "first_server",
		Short:       "001",
		User:        "root",
		Ip:          "127.0.0.1",
		Passwd:      "123456",
		PrivateKey:  "./key",
		PublicKey:   "./key.pub",
		Description: "first test data",
	},
	{
		Name:        "second_server",
		Short:       "002",
		User:        "root",
		Ip:          "192.168.1.1",
		Passwd:      "123456",
		Description: "second test data",
	},
}

var solo *Entry = &Entry{
	Name:        "solo_server",
	Short:       "solo",
	User:        "root",
	Ip:          "57.175.18.11",
	Passwd:      "123456",
	Description: "solo server",
}
