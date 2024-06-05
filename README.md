# GoPortfolioSite

## Introduction

GoPortfolioSite is a personal portfolio website designed to present my projects, skills, and professional experiences. Built using Go, this website aims to provide a fast, responsive, and modern interface to highlight my work.

## Technologies

- **Backend**: Go (Golang)
- **Frontend**: HTML, CSS
- **Version Control**: Git

# Go Server with Nodemon Setup

This project sets up a simple Go server to serve static HTML files and uses `nodemon` to automatically restart the server when changes are detected.

## Prerequisites

- Go (https://golang.org/dl/)
- Node.js and npm (https://nodejs.org/)
- Nodemen (https://nodemon.io/)

## Project Structure

```
GoDevPortfolio/
├── main.go
├── static/
│   ├── home.html
│   ├── cv.html
│   └── blog.html
├── nodemon.json
└── README.md
```

## Setup Instructions

1. **Clone the repository:**

   ```sh
   git clone https://github.com/ArjinAlbay/GoDevPortfolio.git
   cd GoDevPortfolio
   ```

2. **Install `nodemon`:**

   Ensure you have Node.js installed, then install `nodemon` globally:

   ```sh
   npm install -g nodemon
   ```

3. **Create `nodemon.json` file:**

   In your project root directory, create a file named `nodemon.json` with the following content:

   ```json
   {
     "watch": ["*.go", "static/*"],
     "ext": "go html css",
     "exec": "go run main.go"
   }
   ```

4. **Run the server with `nodemon`:**

   ```sh
   nodemon
   ```

   This command will watch for changes in `.go`, `.html`, and `.css` files, and restart the server automatically.

## Usage

- Open your browser and navigate to `http://127.0.0.1:3000/home` to see the home page.
- Navigate to `http://127.0.0.1:3000/cv` to see the CV page.
- Navigate to `http://127.0.0.1:3000/blog` to see the blog page.


## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements, bug fixes, or new features.

1. Fork the repository
2. Create a new branch (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -m 'Add some feature'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Open a pull request
