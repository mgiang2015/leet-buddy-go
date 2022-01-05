package cmd

type Command interface {
	Execute() (string, error)
}
