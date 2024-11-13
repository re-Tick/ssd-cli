# Overview

This directory contains various submodules, each dedicated to a specific scanner. The goal of these modules is to embed the corresponding scanner binary directly into the Go program, allowing for efficient and self-contained scanning operations.

## Structure and Purpose

Each subdirectory within this folder is designed to:
- **Embed**: Include a specific scanner binary as part of the Go program, enabling seamless integration.
- **Build**: Generate a compiled Go binary that packages the embedded scanner, making it easy to distribute and execute.
- **Run Scans**: Execute the embedded scanner binary to perform security or vulnerability scans on the target code or artifacts.
- **Format Reports**: Process and format the scan results to ensure consistency, making them suitable for further analysis or uploading.

## Workflow

The Go program generated from these modules performs the following steps:
1. **Create a Self-Contained Binary**: Combines the Go code and the embedded scanner binary into a single executable file.
2. **Execute Scans**: Uses the embedded scanner to initiate and complete the scanning process.
3. **Report Formatting**: Formats the raw output from the scan into a structured and standardized report format.
4. **Ready for Upload**: The formatted report can then be easily exported or uploaded for integration into larger systems, such as SSD platforms.

By embedding the binaries and automating the scanning workflow, this setup simplifies deployment and enhances the reliability of security checks within CI/CD pipelines or standalone environments.

It also reduces the size of the CLI binary by decoupling the scanners into individual modules, allowing users to include only the necessary scanners in their deployments.