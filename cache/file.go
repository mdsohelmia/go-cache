package cache

import "time"

type File struct{}

func (f File) get(key string) (any, error) {
	return nil, nil
}
func (f File) has(key string) bool {
	return false
}
func (f File) remember(key string, ttl time.Time, callback func()) any {
	return nil
}
func (f File) rememberForever(key string, callback func()) any {
	return nil
}
func (f File) pull(key string) any {
	return nil
}
func (f File) put(key string, value any, ttl ...time.Time) {

}
func (f File) add(key string, value any, ttl ...time.Time) {

}
func (f File) forever(key string, value any) {

}
func (f File) forget(key string) error {
	return nil
}
func (f File) flush() string {
	return "flush"
}
