# Trivy Demo App

Simple Gin + MongoDB app built for Trivy demo. Includes a small regex-based scanner in `exploit/` and a custom Trivy secret rule for hardcoded Mongo URIs.

## Structure

- `app/`: Gin app + Dockerfile + compose
- `exploit/`: Regex scanner
- `trivy-secret.yaml`: Custom Trivy secret rule

## Run the app

From `app/`:

```bash
docker compose up --build
```

## API

- `POST /users` with JSON body:

```json
{
	"username": "alice",
	"email": "alice@example.com"
}
```

- `GET /users` returns an array of users.

## Trivy secret scan

```bash
./trivy-secret-scan.sh
```

## Trivy image scan

Build the image first (from `app/`):

```bash
docker compose up --build
```

Then scan the image:

```bash
trivy image app-app
```
