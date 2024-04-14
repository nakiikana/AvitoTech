package repository

import "errors"

var (
	ErrorNoRowsFound             = errors.New("no rows found")
	ErrorFindingBanner           = errors.New("an error occured when looking for a banner")
	ErrorMarshalingBannerContent = errors.New("couldn't marshal banner's content")
	ErrorStartingTransaction     = errors.New("error starting a new transaction")
	ErrorCreatingBanner          = errors.New("error creating a banner")
	ErrorAddingTF                = errors.New("error adding tags and features")
	ErrorCommittigTransaction    = errors.New("error committing transaction")
	ErrorDelete                  = errors.New("an error occured when deleting")
	ErrorUpdate                  = errors.New("an error occured when updating")
	ErrorGetAdminBanner          = errors.New("error whem completeing queryx")
	ErrorScan                    = errors.New("an error occured during transaction")
	ErrorFilter                  = errors.New("filter query error")
)
