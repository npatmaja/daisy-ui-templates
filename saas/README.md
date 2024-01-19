# A Simple and Responsive SaaS template

This is an HTML template for a SaaS product using [daisyui](https://daisyui.com/) and bit of [alpinejs](https://alpinejs.dev/).

!["The screenshot of the SaaS dashboard template"](daisy-ui-saas.png)

## Features
- Responsive
- Dark mode
- Side bar
- Example forms
- Date picker

## Usage

### Static file server

You may want to use a tool to serve a static file directory over HTTP, e.g., [go-live](https://github.com/antsankov/go-live),
[live-server](https://github.com/tapio/live-server). In my case, I'm using `go-live`. Before running the server, build the `CSS`:

```text
npx tailwindcss -i ./src/input.css -o ./dist/output.css --watch
```

Then serve the `index.html` via `go-live`:

```
go-live -q
```

### Go Templ
You need to install [air](https://github.com/cosmtrek/air) to auto reload the [Go Templ](https://templ.guide/). To
run the project, first make sure that you are in the `saas` directory then run `air`, then open port 8081 to see the result.
```
air
```

## Cooking
- [ ] Range date picker
- [x] Go templ

