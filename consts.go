package wal

const (
	// pageSize defines the size of a single page in the wal
	pageSize = 4096

	// maxPayloadSize defines the max size a payload can have to still fit in a single page
	maxPayloadSize = pageSize - 64
)

const (
	pageStatusInvalid = iota
	pageStatusOther
	pageStatusWritten
	pageStatusComitted
	pageStatusApplied
)

// Metadata contains the header and version of the data being stored.
type Metadata struct {
	Header, Version string
}

var (
	metadata = Metadata{
		Header:  "WAL",
		Version: "1.0",
	}
)
