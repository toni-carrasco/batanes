# Batanes

This was build super quick with [Litepage](https://litepage.dev).

We have just one file `content/index.md`. Write your content in here and it will be transformed to HTML ready to be hosted somewhere.

## Adding images

Put your images in `/public` directory. Make sure they are **optimized**!

A simple tool to do this is https://squoosh.app/.

Upload your image there, convert it to `.webp` format. The default compression settings should be fine. Once you did that, copy the compressed `.webp` file into your public directory and reference that image in your markdown doc.

## Running locally

If you want to run locally, make sure you have Golang setup on your machine.

### Build

```bash
make build
```

This will build the site locally on your machine and places it in `/dist` directory.

### Serve

```bash
make serve
```

This will serve the site locally on http://localhost:3000/batanes

### Dev

```bash
make dev
```

This will serve the site locally on http://localhost:3000/batanes and refresh automatically on changes. **NOTE** You need Air installed to automatically rebuild the site on changes. If not installed, run:

```bash
make install-devtools
```
