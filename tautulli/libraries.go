package tautulli

import (
	"net/http"

	"github.com/dghubble/sling"
)

type LibrariesService struct {
	sling *sling.Sling
}

func newLibrariesService(sling *sling.Sling) *LibrariesService {
	return &LibrariesService{
		sling: sling,
	}
}

type Libraries struct {
	Response struct {
		Message string
		Result  string
		Data    []Library
	} `json:"response"`
}

type Library struct {
	Agent       string `json:"agent"`
	Art         string `json:"art"`
	Count       int    `json:"count"`
	SectionID   int    `json:"section_id"`
	SectionName string `json:"section_name"`
	SectionType string `json:"section_type"`
	Thumb       string `json:"thumb"`
}

type LibraryTable struct {
}

type ExtendedLibrary struct {
	Agent       string `json:"agent"`
	Art         string `json:"art"`
	Count       int    `json:"count"`
	SectionID   int    `json:"section_id"`
	SectionName string `json:"section_name"`
	SectionType string `json:"section_type"`
	Thumb       string `json:"thumb"`
}

func (s *LibrariesService) GetLibraries() ([]Library, *http.Response, error) {

	queryStruct := struct {
		Cmd string
	}{
		Cmd: "get_libraries",
	}

	libraries := new(Libraries)
	resp, err := s.sling.New().Get("").QueryStruct(queryStruct).Receive(libraries, nil)
	return libraries.Response.Data, resp, err
}
