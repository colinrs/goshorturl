package sdk

type Option func(*Options)

type Options struct {
	host       string
	bizTagName string
}

func WithHost(host string) Option {
	return func(options *Options) {
		options.host = host
	}
}

func WithBizTagName(bizTagName string) Option {
	return func(options *Options) {
		options.bizTagName = bizTagName
	}
}
