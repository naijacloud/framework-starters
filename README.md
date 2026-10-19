# Framework starters (QA)

Minimal starters NaijaCloud uses to check that each framework deploys as it comes (TGL-10, TGL-11/12/13).

Each folder is one app. Deploy it as a web service from this public repo with the folder as its root directory.

| Folder | Stack |
|---|---|
| `go-stdlib` | net/http with a subpackage |
| `go-gin`, `go-fiber`, `go-echo` | the three popular Go web frameworks |
| `go-cmd` | `cmd/<app>` layout |
| `go-work` | a go.work workspace |
| `py-django-gunicorn`, `py-django-plain` | Django from `django-admin startproject`, with and without gunicorn |
| `py-flask`, `py-fastapi` | Flask, and FastAPI in `app/main.py` |
| `py-poetry`, `py-uv`, `py-pyproject` | Poetry, uv, and a plain PEP 621 pyproject |
| `php-plain`, `php-laravel`, `php-symfony` | plain PHP, `laravel new`, and the Symfony skeleton |
| `js-vite-express` | Vite + React with an Express server serving the build and an API |
| `js-sveltekit` | the SvelteKit starter as it comes (adapter-auto) |
| `js-astro-ssr` | Astro with `@astrojs/node` |
| `js-angular-ssr` | Angular with server-side rendering |
| `js-yarn-berry` | a Yarn 4 (Plug'n'Play) Node server |
