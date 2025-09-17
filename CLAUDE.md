# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

### Update tools index
```bash
cd scripts && go run update-tools.go
```
This command scans all directories for tool HTML files, extracts metadata (title, description), and regenerates the `app.js` file that displays tools on the main index page.

### Local testing
Open `index.html` in a browser to test the main tools listing page locally.

## Architecture

This is a collection of standalone web tools hosted at tools.spencerlyon.com. Each tool:
- Lives in its own directory at the root level
- Contains a self-contained `index.html` with inline CSS/JS
- Uses Tailwind CSS via CDN for styling
- Is automatically discovered and indexed by the Go script

Key files:
- `scripts/update-tools.go`: Scans directories, extracts tool metadata from HTML files, generates the tools index
- `app.js`: Auto-generated file containing the tools array and rendering logic
- `index.html`: Main landing page that displays all available tools
- `.github/workflows/deploy.yml`: GitHub Actions workflow that runs the Go script and deploys to GitHub Pages on push to main

## Development Workflow

When adding a new tool:
1. Create a new directory with the tool name
2. Add an `index.html` file with proper `<title>` and `<meta name="description">` tags
3. Push to main branch - GitHub Actions will automatically update the index and deploy

The Go script extracts metadata by looking for:
- Tool name from the `<title>` tag
- Description from `<meta name="description" content="...">` or first `<h2>` or `<p>` tag (truncated to 100 chars)