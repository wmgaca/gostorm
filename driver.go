package main

// Driver describes the interface for Gostorm's datasource driver
type Driver interface {

	// Get value from datastore
	Get(string, chan string, chan error)
}
