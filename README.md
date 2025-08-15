# Performance Testing: DB & Redis

## Overview

This project demonstrates performance testing of a system using MariaDB and Redis. It uses **k6** for load testing, **InfluxDB** for metrics, and **Grafana** for visualization. Backend is built with **GoFiber**, **GORM**, and **Hexagonal Architecture**. All services run in **Docker**.

## Features

* GoFiber + GORM (Hexagonal Architecture)
* MariaDB + Redis
* Load testing with k6
* Metrics in InfluxDB + Grafana
* Docker Compose setup

## Run

```bash
docker-compose up -d
```

Run k6 test:

```bash
docker compose run --rm k6 run /scripts/test.js
```

## Notes

* Demonstrates Cache Aside pattern with Redis + DB
* Learning project for Hexagonal Architecture, performance testing, and caching.
