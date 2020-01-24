package main

// Metadata about a file
type Meta struct {
	Directory	bool	// Are we a directory?
	Shared		bool	// Can other users read the file?
	User		string	// Who owns the file
	Modified	int64	// Unix timestamp of time last modified
}

// Represents a File
type File struct {
	Meta					// Metadata for the file
	Data	[maxData]byte	// Stored information in file (if not a directory)
	Kids	[maxKids]*File	// Child files (if a directory)
}


// Number of bytes in the file
func (f *File) Size() int {
	return len(f.Data)
}
