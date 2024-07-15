package main

type Options struct {
	assetsPath string
}

// to be extended later for customisation
func defaultOptions() Options {
	return Options{
		assetsPath: "assets",
	}
}
