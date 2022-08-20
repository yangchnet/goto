package store

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type JsonStore struct {
	FilePath string
}

var _ Store = (*JsonStore)(nil)

func NewJsonStore(filepath string) Store {
	return &JsonStore{
		FilePath: filepath,
	}
}

func (s *JsonStore) load() ([]*Entry, error) {
	storeByte, err := ioutil.ReadFile(s.FilePath)
	if err != nil {
		return nil, err
	}

	entrys := make([]*Entry, 0)
	if err := json.Unmarshal(storeByte, &entrys); err != nil {
		return nil, err
	}

	return entrys, nil
}

func (s *JsonStore) save(entrys []*Entry) error {
	ebs, err := json.Marshal(entrys)
	if err != nil {
		return nil
	}

	if err := ioutil.WriteFile(s.FilePath, ebs, 755); err != nil {
		return err
	}

	return nil
}

func (s *JsonStore) Add(e *Entry) error {
	entrys, err := s.load()
	if err != nil {
		return err
	}

	entrys = append(entrys, e)

	return s.save(entrys)
}

func (s *JsonStore) Get(filter func(*Entry) bool) (*Entry, error) {
	entrys, err := s.load()
	if err != nil {
		return nil, err
	}

	for _, e := range entrys {
		if filter(e) {
			return e, nil
		}
	}

	return nil, errors.New("not found")
}

func (s *JsonStore) List() ([]*Entry, error) {
	return s.load()
}

func (s *JsonStore) Delete(filter func(*Entry) bool) error {
	entrys, err := s.load()
	if err != nil {
		return err
	}

	for i, e := range entrys {
		if filter(e) {
			entrys = append(entrys[:i], entrys[i+1:]...)
		}
	}

	return s.save(entrys)
}

func (s *JsonStore) Update(filter func(*Entry) bool, entry *Entry) error {
	entrys, err := s.load()
	if err != nil {
		return err
	}

	for _, e := range entrys {
		if filter(e) {
			e.Name = entry.Name
			e.Passwd = entry.Passwd
			e.Ip = entry.Ip
			e.PrivateKey = entry.PrivateKey
			e.PublicKey = entry.PublicKey
			e.Short = entry.Short
			e.User = entry.User
			break
		}
	}

	return s.save(entrys)
}
