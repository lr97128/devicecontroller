package main

type Device struct {
	Host      string `json:"host"`
	Username  string `json:"username"`
	CiphPass  string `json:"ciphpass"`
	PlainPass string
	Secret    string   `json:"secret"`
	Key       string   `json:"key"`
	UseKey    bool     `json:"usekey"`
	Command   []string `json:"command"`
}

func (d *Device) Exec() error {
	return nil
}
