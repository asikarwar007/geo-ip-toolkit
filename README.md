# geo-ip-toolkit

geo-ip-toolkit is an open-source platform offering unified access to extensive IP geolocation data, integrating information from leading data providers like ipinfo.io and ip-api.com. Designed for developers, analysts, and tech enthusiasts, it simplifies the retrieval of IP insights, enhancing applications in cybersecurity, analytics, and personalized content delivery.

## Features

- **Unified API**: Aggregate geolocation data from multiple sources through a single, simplified API.
- **Extensive IP Insights**: Access geographical location, ISP details, organization info, network ASN information, and security flags for any IP address.
- **Flexible and Extensible**: Easily integrate with existing systems and expand to include additional data sources as needed.
- **Open Source**: Contribute to or modify the project to suit your specific requirements.

## Getting Started

### Prerequisites

Ensure you have Go installed on your system (version 1.16 or later).

### Installation

1.  **Clone the repository:**
```bash
git clone https://github.com/asikarwar007/geo-ip-toolkit
```

2. **Navigate to the project directory:**
```bash
cd geo-ip-toolkit
```

3. **Load environment variables (ensure you have a .env file with the necessary API tokens):**
```bash
source .env
```

4. **Run the project:**
```bash
go run .
```

## Authentication

The IPinfo Go library can be authenticated with your IPinfo API access token, which is stored in .env file. Your IPInfo access token can be found in the account section of IPinfo's website after you have signed in: https://ipinfo.io/account/token

```bash
API_TOKEN=your_api_token_here
```


## Usage
geo-ip-toolkit simplifies accessing extensive IP geolocation data through a unified API. After setting up and running the service, you can query IP information by making HTTP requests to the local server.

### Fetching IP Information
To retrieve details for a specific IP address, use the following HTTP GET request format:

```bash
http://localhost:8080/info?ip={IP}&provider={IP_Provider}
```
Replace {IP} with the IP address you want to query (e.g., 8.8.8.8) and {IP_Provider} with the identifier for the IP information provider you wish to use (e.g., ipinfo for ipinfo.io or ipapi for ip-api.com).

### Example Request
For example, to fetch information about the IP address 8.8.8.8 using ipinfo.io as the provider, you would access:

```bash
http://localhost:8080/info?ip=8.8.8.8&provider=ipinfo
```

### Supported Providers
Currently, geo-ip-toolkit supports the following providers:

1. `ipinfo`: Utilizes ipinfo.io for fetching IP data.
2. `ipapi`: Utilizes ip-api.com for fetching IP data.

You can extend geo-ip-toolkit to support additional providers by implementing the IPInfoProvider interface for each new service.




# Contributing
We welcome contributions to geo-ip-toolkit! If you have suggestions for improvement or want to contribute code, please follow these steps:

## Fork the repository.
1. Create a new branch for your feature (git checkout -b feature/amazing-feature).
2. Commit your changes (git commit -am 'Add some amazing feature').
3. Push to the branch (git push origin feature/amazing-feature).
4. Open a Pull Request.

# License
Distributed under the MIT License. See LICENSE for more information.

# Acknowledgments
Thanks to ipinfo.io and ip-api.com for providing IP data services.
Special thanks to all contributors and users of the geo-ip-toolkit project.
