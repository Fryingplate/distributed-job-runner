# Distributed Job Runner

A distributed background job processing system built with Go, PostgreSQL, Redis, and Docker.

The application allows clients to submit jobs through a REST API. Jobs are stored in PostgreSQL and published to a Redis queue, where worker processes consume and execute them asynchronously.

### Features

* REST API for job submission
* PostgreSQL job persistence
* Redis-based job queue
* Multiple worker support
* Asynchronous job processing
* Dockerized local setup

### Technologies

* Go
* PostgreSQL
* Redis
* Docker

### Purpose

This project was built to learn backend development, distributed systems concepts, asynchronous processing, queue-based architectures, and scalable worker patterns using Go.

