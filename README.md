# Server Core 🎯

> A production-ready Go foundation for building hypermedia-driven web applications. Because sometimes the frontend framework treadmill gets exhausting.

A comprehensive Go server application template with type-safe HTML templating, modern hypermedia patterns, and i18n support. This is what happens when you realize that shipping HTML from the server doesn't have to mean trading away developer experience.

## Overview

Server Core is a foundational Go template designed to be extended for various web projects. It implements modern hypermedia-driven architecture using templ for type-safe HTML generation, Echo for high-performance routing, and HTMX-ready patterns for progressive enhancement.

## Skills Demonstrated

This project showcases professional Go backend development competencies:

- **Type-Safe HTML Templating with templ**: Uses the a-h/templ library for compile-time verified HTML generation with full Go type safety. Similar in spirit to the TSX controllers approach, but leveraging Go's compile-time guarantees instead of TypeScript. Components are written as Go functions that return templ.Component interfaces, enabling composable, testable UI patterns without runtime template parsing overhead.

- **Modern Go Web Architecture**: Echo v4 framework with middleware patterns, graceful shutdown, structured logging, and production-ready HTTP server configuration

- **Hypermedia-Driven Design**: HTMX-compatible endpoints serving HTML fragments, progressive enhancement patterns, and server-side state management without JavaScript framework overhead

- **Component Architecture**: Reusable UI components (header, footer, user avatar, logo, main content) built with templ's composable functions, demonstrating modular frontend patterns in a server-rendered context

- **Internationalization (i18n)**: Built-in translation support using golang.org/x/text with message catalogs, locale detection, and runtime language switching

- **CLI Patterns**: Cobra-based command structure for server, admin, and extensible subcommands with flags and configuration management

- **Configuration Management**: Viper integration supporting environment variables, config files, and sensible defaults for 12-factor app deployment

- **Session Management**: Gorilla sessions integration for stateful authentication patterns with secure cookie handling

- **Input Validation**: go-playground/validator for struct-level validation with custom rules and localized error messages

- **Hot Reload Development**: Air integration for rapid development cycles with automatic recompilation on file changes

- **Modern CSS Tooling**: TailwindCSS utility-first approach with DaisyUI component library for rapid UI development

## Tech Stack

### Core Framework
- **Go** 1.21+ - Systems language with modern concurrency primitives
- **Echo** v4.12.0 - High-performance, extensible web framework
- **templ** v0.2.771 - Type-safe HTML templating

### Frontend Integration
- **HTMX-ready** - Hypermedia-driven architecture with progressive enhancement
- **TailwindCSS** - Utility-first CSS framework with JIT compilation
- **DaisyUI** - Tailwind-based component library
- **Flowbite** (optional) - Additional component patterns

### Configuration & CLI
- **Cobra** v1.8.1 - Modern CLI framework for Go applications
- **Viper** v1.19.0 - Complete configuration solution

### Internationalization
- **golang.org/x/text** - Unicode and i18n support including message catalogs

### Session & Validation
- **Gorilla Sessions** v1.3.0 - Cookie and filesystem session management
- **validator/v10** v10.22.0 - Struct and field validation

### Development Tools
- **Air** - Hot reload for Go applications
- **fatih/color** - Colorized terminal output for better logging

## Project Structure

```
server-core/
├── cmd/
│   └── main.go              # Application entry point with CLI
├── internal/
│   ├── assets/              # Static asset handling
│   ├── handlers/            # HTTP request handlers
│   ├── middleware/          # Echo middleware configuration
│   ├── router/              # Route registration
│   ├── server/              # Server initialization
│   └── translations/        # i18n message catalogs
├── pkg/
│   ├── components/          # Reusable templ components
│   ├── config/              # Configuration parsing
│   ├── htmx/                # HTMX utilities
│   ├── layouts/             # Page layout templates
│   ├── logger/              # Structured logging
│   ├── pages/               # Full page templates
│   ├── templ/               # templ utilities
│   └── translation/         # i18n utilities
├── public/                  # Static assets (CSS, JS, images)
├── .air.toml                # Hot reload configuration
└── go.mod                   # Go module definition
```

## Getting Started

### Prerequisites
- Go >= 1.21
- Air (for hot reload): `go install github.com/air-verse/air@latest`
- templ CLI: `go install github.com/a-h/templ/cmd/templ@latest`

### Installation

```bash
# Clone the repository
git clone https://github.com/ashmortar/server-core.git
cd server-core

# Install dependencies
go mod download

# Generate templ templates (if not already generated)
templ generate

# Build CSS (if using TailwindCSS)
npm install
npm run build:css  # or equivalent tailwind build command
```

### Development

```bash
# Start development server with hot reload
air

# Or run directly
go run cmd/main.go

# Run admin interface
go run cmd/main.go admin
```

The server will start on the configured port (default: typically 3000 or 8080, check your config).

### Building for Production

```bash
# Generate templ templates
templ generate

# Build the binary
go build -o server-core cmd/main.go

# Run production server
./server-core
```

### Configuration

Server Core uses Viper for flexible configuration. Create a config file or use environment variables:

```bash
# Example environment variables
export PORT=3000
export LOG_LEVEL=debug
export SESSION_SECRET=your-secret-key-here

# Or create a config file (config.yaml, config.json, etc.)
# See pkg/config/config.go for available options
```

## Architecture Highlights

### templ Components

Components are written as Go functions returning `templ.Component`:

```go
// pkg/components/user-avatar.templ
package components

templ UserAvatar(username string, avatarURL string) {
  <div class="avatar">
    <img src={avatarURL} alt={username}/>
  </div>
}
```

After running `templ generate`, these become type-safe Go functions:

```go
func HandleUserProfile(c echo.Context) error {
  user := getUserFromContext(c)
  component := components.UserAvatar(user.Name, user.Avatar)
  return templ.Render(c.Response(), component)
}
```

### HTMX Integration

The `pkg/htmx` package provides utilities for working with HTMX requests:

```go
if htmx.IsHTMXRequest(c) {
  // Return HTML fragment
  return templ.Render(c.Response(), components.SomeFragment())
}
// Return full page
return templ.Render(c.Response(), layouts.FullPage(content))
```

### Internationalization

Messages are defined in message catalogs and accessed via the translation utilities:

```go
import "core/pkg/translation"

func HandlePage(c echo.Context) error {
  tr := translation.FromContext(c)
  message := tr.Get("welcome.message")
  // Use message in templates
}
```

## Extending This Template

This template is designed to be forked and extended:

1. **Fork or clone** the repository
2. **Customize** the components and layouts for your UI needs
3. **Add** your domain-specific handlers and routes
4. **Extend** the CLI with additional commands as needed
5. **Deploy** with your preferred hosting (containerized, systemd service, etc.)

## Contributing

This is a personal template project, but feedback and suggestions are welcome via issues.

## License

UNLICENSED - This is a personal template project.

---

**Author:** [Aaron Ross](https://github.com/ashmortar)

*Part of a curated collection of production-ready templates and architectural experiments. This project explores the intersection of type-safe server-side rendering and modern backend frameworks - specifically investigating how Go's compile-time guarantees and templ's component model can deliver the developer experience of modern frontend frameworks while maintaining the simplicity and performance of hypermedia-driven architectures. The question: can we have our type safety and eat our bandwidth savings too?*
