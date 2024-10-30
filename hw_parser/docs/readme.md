# The solution includes two primary modules in Go:

-Configuration Parsing Module: Parses, maps, and structures hardware data from Excel sheets.
E-xcel Formatting Module: Generates a formatted Excel document from structured data.

## Solution Outline
Define Configuration Structures

### Map Python dictionaries to Go structs.

Use structs like ModuleType for each device type and IOReplace for module replacement mappings.

### Parse Input Excel Files

Load and parse Excel sheets with excelize, grouping data as required.

### Generate Formatted Output

Using the data parsed, write to a new Excel file with formatting, following the specifications in IO_DOCUMENT.

### Main Application Structure

Load mappings from JSON/YAML files.
-Parse raw Excel input.
-Structure and format output.

## Project Hierachy

hw_parser/
│
├── config/
│   ├── config_loader.go   # This file contains LoadConfig and Configuration struct
│   ├── config_test.go      # Unit tests for the config_loader
├── logger/
│   ├── logger.go
├── parser/
│   ├── parser.go
├── formatter/
│   ├── formatter.go
|
├── main.go                 # Main application file
├── test_config.yaml        # Test configuration file
├── go.mod                  # Go module file
├── go.sum                  # Go module dependencies file
└── run.sh                  # Bash script to build and run the application
