package nika

type Module interface {
    Controllers() []interface{}
    Providers() []interface{}
    Imports() []Module
}