package args

var Debug bool

var Dataset struct {
	Dir string
}

var Database struct {
	Type   string
	Host   string
	User   string
	Passwd string
	Name   string
	Pragma string
}

var Logger struct {
	Dir    string
	Level  string
	Stdout bool
	ToFile bool
}

var Server struct {
	DSN      string
	Listen   string
	Register bool
}

var Worker struct {
	Remote string
}