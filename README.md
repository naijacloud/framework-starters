# Framework starters (QA)

Minimal starters NaijaCloud uses to check that each Python, PHP and Go framework deploys as it comes (TGL-11/12/13).

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
