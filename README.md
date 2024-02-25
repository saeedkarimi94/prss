# Rss API Integration

A simple Go program using the Gin framework to create an API endpoint that fetches an RSS feed from [Zoomit](https://www.zoomit.ir/feed/) and returns it as JSON.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Error Handling](#error-handling)
- [Contributing](#contributing)
- [License](#license)

## Installation

To run the Zoomit API Integration locally, make sure you have Go installed. Clone this repository and execute the following command:

```bash
go run main.go
```

## Usage
- Once the application is running, access the API endpoint at http://localhost:8080/api/zoomit to retrieve the Zoomit RSS feed in JSON format.
- Once the application is running, access the API endpoint at http://localhost:8080/api/digiato to retrieve the Digiato RSS feed in JSON format.

## API Endpoints
GET /api/zoomit
Fetches the Zoomit RSS feed and returns it as JSON.

GET /api/digiato
Fetches the Digiato RSS feed and returns it as JSON.

## Error Handling
The API provides structured error responses for various scenarios:
- Failed RSS Feed Retrieval
```json
{
  "error": "Failed to fetch the RSS feed"
}
```
- Unexpected Status Code 
```json
{
  "error": "Unexpected status code: 502"
}
```
- Invalid XML Format
```json
{
  "error": "Invalid XML format"
}
```
- Failed JSON Conversion
```json
{
  "error": "Failed to convert data to JSON"
}
```
## Contributing
If you'd like to contribute to this project, feel free to open an issue or submit a pull request. Please follow the Contribution Guidelines.

## License
This Zoomit API Integration is licensed under the MIT License. See the LICENSE file for details.
