# Fishing Diary Roadmap

This project is a backend system for logging fishing catches with enrichment data (location, weather, media).

---

## 🟢 Phase 1 — Core API (DONE)
- [x] POST /catches
- [x] GET /catches
- [x] In-memory storage
- [x] Basic tests
- [x] GitHub repo + CI setup

---

## 🟡 Phase 2 — Architecture (DONE)
- [x] Service layer introduced
- [x] Store abstraction
- [x] Dependency injection
- [x] Test isolation improved

---

## 🔵 Phase 3 — Data Model Expansion (NEXT)
- [ ] Add location (lat/lng) to catches
- [ ] Validate input fields
- [ ] Improve timestamps handling
- [ ] Define OpenAPI spec

---

## 🌦 Phase 4 — External Data Integration
- [ ] Weather API integration
- [ ] Attach weather snapshot to catch
- [ ] Time-based weather lookup fallback

---

## 🗺 Phase 5 — Spatial Features
- [ ] Map visualization support
- [ ] Cluster catches by location
- [ ] Filter by region

---

## 📦 Phase 6 — Persistence
- [ ] Replace in-memory store with MySQL
- [ ] Migrations
- [ ] Repository layer

---

## 📊 Phase 7 — Observability
- [ ] Structured logging
- [ ] Metrics (Prometheus-style thinking)
- [ ] Basic request tracing

---

## 🚀 Phase 8 — Polish (Portfolio Level)
- [ ] OpenAPI documentation
- [ ] Docker setup
- [ ] GitHub Actions CI
- [ ] Clean README with diagrams