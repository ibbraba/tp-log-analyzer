package analyzer

import "fmt"

type FileNotFoundError struct {
	URL string
	Err error
}

type ParsingError struct {
	URL string
	Err error
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("Log introuvable : %s (%v)", e.URL, e.Err)
}

func (e *FileNotFoundError) Unwrap() error {
	return e.Err
}

func (e *ParsingError) Error() string {
	return fmt.Sprintf("Erreur d'analyse du log : %s (%v)", e.URL, e.Err)
}

func (e *ParsingError) Unwrap() error {
	return e.Err
}
