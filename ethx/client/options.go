package client

type Option func(*options)

type options struct {
	origin string
}

func applyOptions(opts []Option) options {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	return o
}

func WithOrigin(origin string) Option {
	return func(o *options) {
		o.origin = origin
	}
}
