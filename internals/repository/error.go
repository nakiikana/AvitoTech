package repository

import "errors"

var (
	ErrorNoRowsFound             = errors.New("no rows found")
	ErrorFindingBanner           = errors.New("an error occured when looking for a banner")
	ErrorMarshalingBannerContent = errors.New("couldn't marshal banner's content")
)
