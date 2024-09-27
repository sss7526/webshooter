package processor

type targetOptions struct {
	verbose 	bool
	saveToImage	bool
	saveToPDF	bool
	translate	bool
	useTorProxy	bool
}

type option func(*targetOptions)

func WithVerbose(verbose bool) option {
	return func(opts *targetOptions) {
		opts.verbose = verbose
	}
}

func WithSaveToImage(save bool) option {
	return func(opts *targetOptions) {
		opts.saveToImage = save
	}
}

func WithSaveToPDF(save bool) option {
	return func(opts *targetOptions) {
		opts.saveToPDF = save
	}
}

func WithTranslate(translate bool) option {
	return func(opts *targetOptions) {
		opts.translate = translate
	}
}

func WithTorProxy(useTor bool) option {
	return func(opts *targetOptions) {
		opts.useTorProxy = useTor
	}
}

func NewTargetOptions(options ...option) *targetOptions {
	opts := &targetOptions{}
	for _, opt := range options {
		opt(opts)
	}
	return opts
}
