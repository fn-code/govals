package govals

type noValRule struct{}

func (d *noValRule) validate() (bool, error) {
	return true, nil
}
