// Package egress provides implementations for clients that handle outgoing network calls.
// It is responsible for interacting with external services or systems, such as uploading data or installing binaries.
// This package facilitates the export of processed data, typically after it has been handled by the business logic layers,
// ensuring that the necessary information is sent to the appropriate external endpoints (e.g., APIs, servers, or storage locations).
// It abstracts the communication details for tasks like data transfer and installation, ensuring a clean separation
// between business logic and external interactions.
package egress

type Installer interface {
	InstallScanner(scannerName string, version string) error
}

type Uploader interface {
	UploadReport(reportPath string) error
}
