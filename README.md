# SGL Tools

A collection of simple, useful web tools hosted at [tools.spencerlyon.com](https://tools.spencerlyon.com)

## Adding a New Tool

1. Create a new folder in the root directory with a descriptive name (e.g., `json-formatter`, `base64-encoder`)

2. Inside the folder, create at minimum an `index.html` file. You can also add:
   - CSS files (or use Tailwind CSS via CDN)
   - JavaScript files
   - Any other assets needed

3. Structure your `index.html`:

   ```html
   <!DOCTYPE html>
   <html lang="en">
   <head>
       <meta charset="UTF-8">
       <meta name="viewport" content="width=device-width, initial-scale=1.0">
       <meta name="description" content="Brief description of your tool">
       <title>Your Tool Name</title>
       <script src="https://cdn.tailwindcss.com"></script>
   </head>
   <body>
       <!-- Your tool content -->
   </body>
   </html>
   ```

4. Push to the `main` branch - GitHub Actions will automatically:
   - Run the Go script to update the tools index
   - Deploy to GitHub Pages

## Local Development

To test the tools index generation locally:

```bash
cd scripts
go run update-tools.go
```

Then open `index.html` in your browser.

## How It Works

- The Go script (`scripts/update-tools.go`) scans all directories for `index.html` files
- It extracts the title and description from each tool's HTML
- It generates the `app.js` file with the tools list
- GitHub Actions runs this script on every push and deploys to GitHub Pages
