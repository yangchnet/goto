package store

type Entry struct {
	Name        string `json:"name"`
	Short       string `json:"short"`
	User        string `json:"user"`
	Ip          string `json:"ip"`
	Passwd      string `json:"passwd"`
	PrivateKey  string `json:"private_key"`
	PublicKey   string `json:"public_key"`
	Description string `json:"description"`
}

type Store interface {
	Add(*Entry) error
	Get(filter func(*Entry) bool) (*Entry, error)
	List() ([]*Entry, error)
	Delete(filter func(*Entry) bool) error
	Update(filter func(*Entry) bool, e *Entry) error
}
