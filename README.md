# Weather API Wrapper Service

A Go-based weather API that fetches and returns weather data from a 3rd party provider with Redis caching.

**Project URL:** https://roadmap.sh/projects/weather-api-wrapper-service

## Overview

This project is a weather API wrapper service that demonstrates:
- Working with 3rd party APIs (Visual Crossing Weather API)
- Redis caching with automatic expiration (12-hour TTL)
- Environment variable configuration
- Rate limiting to prevent API abuse
- Layered architecture with proper separation of concerns

## Architecture

```
cmd/
└── weather-api/
    └── main.go           # Application entry point & dependency wiring

internal/
├── api/
│   ├── handlers/         # HTTP request handlers
│   └── routes/           # Route definitions
├── cache/                # Redis cache implementation
├── config/               # Configuration structs
├── models/               # Data structures
└── services/             # Business logic
```

## Features

- **Weather Data Retrieval**: Get weather forecasts by country and city
- **Redis Caching**: Responses are cached for 12 hours to reduce API calls
- **Rate Limiting**: 100 requests per 10 seconds per IP/endpoint
- **Configurable Days**: Request 1-15 days of weather data via query parameter

## Prerequisites

- Go 1.24+
- Redis server running locally
- Visual Crossing API key ([Get free API key](https://www.visualcrossing.com/weather-api))

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Yer01/weather-app.git
cd weather-app
```

2. Create a `.env` file from the example:
```bash
cp .env.example .env
```

3. Add your Visual Crossing API key to `.env`:
```
API_KEY=your_api_key_here
```

4. Make sure Redis is running on default port (6379)

5. Run the application:
```bash
go run ./cmd/weather-api
```

## API Usage

### Get Weather Report

```
GET /report/{country}/{city}?days={number}
```

**Parameters:**
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| country | path | Yes | Country name |
| city | path | Yes | City name |
| days | query | No | Number of days (1-14), default: 7 |

**Example Request:**
```bash
curl "http://localhost:8081/report/Kazakhstan/Almaty?days=3"
```

**Example Response:**
```json
{
  "address": "Almaty, Kazakhstan",
  "timezone": "Asia/Almaty",
  "tzoffset": 5,
  "days": [
    {
      "datetime": "2026-01-15",
      "temp": -5.2,
      "humidity": 78.5,
      "windspeed": 12.3,
      "pressure": 1015.2,
      "cloudcover": 45.0,
      "sunrise": "08:23:00",
      "sunset": "18:05:00"
    }
  ],
  "stations": {}
}
```

## Tech Stack

- **Language**: Go 1.24
- **Router**: [chi](https://github.com/go-chi/chi)
- **Rate Limiting**: [httprate](https://github.com/go-chi/httprate)
- **Cache**: [Redis](https://redis.io/) via [go-redis](https://github.com/redis/go-redis)
- **Weather API**: [Visual Crossing](https://www.visualcrossing.com/weather-api)

## License

MIT
