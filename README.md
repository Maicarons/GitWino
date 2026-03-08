# GitWino

![GitWino Screenshot](website.png)

**Languages:** [English](README.md) | [中文](README_zh.md)

GitWino is a powerful tool for querying and analyzing Git repository history. It provides detailed timeline information about code commits, repository creation, version tags, and releases across multiple Git platforms.

![License](https://img.shields.io/badge/license-Apache_2.0-blue.svg)
![Go Version](https://img.shields.io/badge/go-1.25.0-brightgreen.svg)

## Features

- **Multi-Platform Support**: Supports GitHub, Gitee, and GitLab
- **Comprehensive Timeline**: Query the following key timestamps:
  - First code commit time
  - Repository creation time
  - First version tag time
  - First release time
- **Time Zone Display**: Shows both UTC and local time for all timestamps
- **Time Difference Analysis**: Automatically calculates time differences between events
- **Export Options**: Export results as JSON or image
- **Modern UI**: Beautiful gradient interface with responsive design
- **Multi-Language**: Supports Chinese, English, Korean, and French

## Technology Stack

### Backend
- **Language**: Go 1.25.0
- **Web Framework**: Gin
- **Git Library**: go-git/v5
- **Deployment**: Vercel serverless deployment support

### Frontend
- **UI Framework**: Layui
- **Libraries**: html2canvas (for image export)
- **CDN**: BootCDN for fast asset delivery

## Installation

### Local Development

1. **Clone the repository**
```bash
git clone https://github.com/Maicarons/GitWino.git
cd gitwino
```

2. **Install dependencies**
```bash
go mod download
```

3. **Run the application**
```bash
go run main.go
```

4. **Access the application**
Open your browser and navigate to `http://localhost:8080`

### Deploy to Vercel

GitWino supports serverless deployment on Vercel:

1. Install Vercel CLI:
```bash
npm install -g vercel
```

2. Deploy:
```bash
vercel
```

The `vercel.json` configuration file is already included in the project.

## Usage

1. **Enter Repository URL**: Input a Git repository URL (supports GitHub, Gitee, or GitLab)
   - Example: `https://github.com/owner/repo`
   - Example: `https://gitee.com/owner/repo`
   - Example: `https://gitlab.com/owner/repo`

2. **Click Search**: The system will automatically:
   - Clone the repository temporarily
   - Analyze commit history
   - Query platform APIs for creation and release information
   - Calculate time differences between events

3. **View Results**: Results include:
   - Repository URL
   - First commit time (with UTC and local time)
   - Repository creation time (with time difference from first commit)
   - First tag time (with time difference from first commit)
   - First release time (with time difference from first commit)

4. **Export Data**: 
   - Click "Export JSON" to download data in JSON format
   - Click "Export Image" to save results as a PNG image

## API Reference

### Query Repository History

**Endpoint**: `GET /api`

**Parameters**:
- `repo` (required): Git repository URL

**Response Example**:
```json
{
  "earliest_commit_time": "2020-01-15T08:30:00Z",
  "repo_creation_time": "2020-01-10T12:00:00Z",
  "earliest_tag_time": "2020-02-01T10:00:00Z",
  "earliest_release_time": "2020-02-05T14:30:00Z",
  "repo_url": "https://github.com/owner/repo.git"
}
```

**Error Response**:
```json
{
  "error": "Error message description"
}
```

## Project Structure

```
gitwino/
├── api/                    # API handlers
│   └── index.go           # HTTP request handlers
├── frontend/              # Frontend static files
│   ├── index.html        # Main HTML page
│   ├── app.js            # Application logic
│   ├── style.css         # Styles
│   └── i18n.js           # Internationalization
├── internal/git/          # Git service implementation
│   ├── providers/        # Platform-specific providers
│   │   ├── github.go     # GitHub API integration
│   │   ├── gitee.go      # Gitee API integration
│   │   ├── gitlab.go     # GitLab API integration
│   │   └── provider_manager.go  # Provider management
│   ├── git.go           # Main Git logic
│   ├── repository.go    # Repository operations
│   └── types.go         # Data structures
├── main.go              # Application entry point
├── go.mod               # Go module definition
├── go.sum               # Dependency checksums
└── vercel.json          # Vercel deployment config
```

## Supported Platforms

- **GitHub**: Full support for commit history, repository creation, tags, and releases
- **Gitee**: Support for commit history, repository creation, and releases
- **GitLab**: Full support for commit history, repository creation, tags, and releases
- **Other Platforms**: Limited support for commit history
## How It Works

1. **URL Parsing**: The system identifies the Git platform from the repository URL
2. **Repository Cloning**: Temporarily clones the repository using go-git
3. **Commit Analysis**: Scans all branches to find the earliest commit
4. **Tag Analysis**: Finds the earliest tag by commit time
5. **API Integration**: Queries platform APIs for repository creation and release times
6. **Data Processing**: Calculates time differences and formats timestamps
7. **Result Display**: Presents data with both UTC and local time zones

## Development

### Requirements
- Go 1.25.0 or higher
- Node.js (optional, for Vercel CLI)

### Build
```bash
go build -o gitwino main.go
```

### Run Tests
```bash
go test ./...
```

## Limitations

- Private repositories require API token configuration (not yet implemented)
- Very large repositories may take longer to analyze due to cloning time
- Some platforms may have API rate limits

## Future Enhancements

- Support for private repositories with authentication
- Additional Git platform support (Bitbucket, etc.)
- Historical statistics and charts
- Batch repository analysis
- API token configuration interface

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the Apache 2.0 License.

## Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [go-git](https://github.com/go-git/go-git)
- [Layui](https://layui.dev/)
- [html2canvas](https://html2canvas.hertzen.com/)

## Contact

- **Author**: Maicarons
- **Project Link**: [https://github.com/Maicarons/GitWino](https://github.com/Maicarons/GitWino)

---

Made with ❤️ by Maicarons
