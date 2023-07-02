package valmail

import (
	"net"
)

func (e *EmailAddress) populateMX() error {
	// don't re-populate if we already have records
	if e.mxRecords != nil || len(e.mxRecords) > 0 {
		return nil
	}
	// otherwise lets do an MX lookup for the domain
	mxRecords, err := net.LookupMX(e.Domain)
	if err != nil {
		return err
	}
	// check if they are nil or zero length
	if len(mxRecords) == 0 {
		return err
	} else {
		// if they exist store them in case we need them later
		e.mxRecords = mxRecords
	}
	return nil
}

func (e *EmailAddress) MXValid() bool {
	// can't validate if this isn't a real email
	if !e.SyntaxValid {
		return false
	}
	err := e.populateMX()
	if err != nil {
		return false
	}
	if len(e.mxRecords) > 0 {
		return true
	}
	return false
}
