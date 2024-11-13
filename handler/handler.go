// Package handler contains the business logic to initiate and manage the scanning process.
// It orchestrates the flow of execution by following four essential steps:
// 1. Install the required binary for scanning.
// 2. Start the scan using the installed binary.
// 3. Format the generated scan reports.
// 4. Upload the formatted report to SSD.
//
// This package acts as the core of the scanning process, coordinating various components
// to perform end-to-end operations efficiently.
package handler

type ScanHanlder interface {
	Run() error
}
