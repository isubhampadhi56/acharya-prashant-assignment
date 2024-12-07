package tokencache

type BlackListedToken interface {
	Set(string, int64)
	Remove(string)
	IsPresent(string) bool
	GetExpTime(string) (int64, error)
}
