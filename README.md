# MotoEx — Vehicle Import Catalog

A full-stack web application for a small car import business. Customers browse imported vehicles; the owner manages inventory through an admin panel.

**Live project built independently from scratch.**

---

## What I Built

- Public catalog with filtering, photo carousel, and lightbox
- Vehicle detail pages with full specs and gallery
- Admin panel: create / edit / delete vehicles, upload multiple photos
- Multilingual UI — Polish, English, Russian (i18next)
- REST API with simple token-based authentication

---

## Stack

**Backend** — Go, Gin, PostgreSQL (pgx v5)  
**Frontend** — React 19, TypeScript, Vite, Tailwind CSS  
**Other** — React Router v7, Embla Carousel, i18next

---

## Architecture

```
Go backend (port 8080)
  ├── REST API — /cars, /admin
  ├── serves built React app from static/dist/
  └── stores uploaded photos on filesystem

React frontend (Vite dev server → proxies to backend)
  ├── pages: Home, Catalog, Vehicle Details
  └── custom hooks for data fetching
```

Photo uploads are renamed to UUIDs on the server and stored under `static/uploads/`. Paths are saved in the database.

Auth is a single admin key passed via `Authorization` header — intentionally simple for a solo-owner use case.

---

## Technical Decisions

**Graceful shutdown** — the server listens for `SIGINT`/`SIGTERM` and gives in-flight requests 30 seconds to finish before exiting. HTTP timeouts are also set explicitly (read 10s, write 30s, idle 120s) so no connection can hang the server indefinitely.

**Connection pool** — database access uses `pgxpool` instead of a single connection. Under concurrent requests the pool manages multiple connections automatically without extra setup.

**Constant-time auth comparison** — the admin key is compared with `crypto/subtle.ConstantTimeCompare` to prevent timing attacks. A naive `==` comparison leaks information about how many characters matched; constant-time comparison doesn't.

**Parameterized queries** — all SQL uses `$1, $2, ...` placeholders via pgx. No string concatenation, no SQL injection surface.

**Context propagation** — every database call receives `c.Request.Context()`. If the client disconnects mid-request, the query is cancelled rather than running to completion for nobody.

**RowsAffected check on mutations** — after UPDATE and DELETE, the code checks whether any row was actually affected and returns 404 if not, rather than silently returning 200 for a missing record.

**UUID file naming** — uploaded photos are saved as `<uuid>.<ext>`. This prevents filename collisions and makes paths unpredictable so they can't be enumerated.

**SPA fallback routing** — unknown paths that aren't API routes serve `index.html`, so React Router handles client-side navigation without 404s on hard refresh.

**Structured logging** — `log/slog` with a JSON handler. Every error log includes the error value as a structured field, making logs parseable by any log aggregator.

---

## What This Project Demonstrates

- Building a production-grade REST API in Go: middleware, CORS, timeouts, graceful shutdown
- Connecting React to a real backend (not mock data) with custom hooks
- File upload handling and static asset serving from Go
- Shipping a monorepo: `make build` compiles the frontend into the Go binary's static folder
