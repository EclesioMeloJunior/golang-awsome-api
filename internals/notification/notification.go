package notification

type Notifier interface {
	Notify() error
}
