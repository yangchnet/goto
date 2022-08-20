package store

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var filepath string = "./goto.json"

var s Store

func TestMain(m *testing.M) {
	s = NewJsonStore(filepath)

	fd, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}

	jsonBytes, _ := json.Marshal(entrys)
	if _, err := fd.Write(jsonBytes); err != nil {
		log.Fatal(err)
	}

	exitCode := m.Run()

	_ = os.Remove(filepath)

	os.Exit(exitCode)
}

func Test_AddGetDelete(t *testing.T) {
	require.NoError(t, s.Add(solo))

	e, err := s.Get(func(e *Entry) bool {
		return e.Name == "solo_server"
	})
	require.NoError(t, err)

	require.NotNil(t, e)
}

func Test_Delete(t *testing.T) {
	require.NoError(t, s.Delete(func(e *Entry) bool {
		return e.Name == "first_server"
	}))

	e, _ := s.Get(func(e *Entry) bool {
		return e.Name == "first_server"
	})
	require.Nil(t, e)
}

func Test_Update(t *testing.T) {
	err := s.Update(func(e *Entry) bool {
		return e.Name == "first_server"
	}, solo)

	require.NoError(t, err)

	e, err := s.Get(func(e *Entry) bool {
		return e.Name == "solo_server"
	})
	require.NoError(t, err)

	require.Equal(t, "57.175.18.11", e.Ip)

}
