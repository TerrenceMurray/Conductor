# Conductor 

The Distributed Workflow Engine is a self-hosted platform for submitting asynchronous jobs,
assigning them to registered workers, tracking execution, recovering from failures, and
visualizing system state in real time. It begins as a deliberately small Go/PostgreSQL system and
evolves toward a reliable distributed workflow platform with Go and Python workers,
scheduling, retries, observability, and dependency-based workflows.

The project is intentionally designed to exercise core engineering skills rather than maximize
feature count. Important mechanisms such as worker coordination, leases, heartbeats,
cancellation, idempotency, retry behavior, concurrency control, and graceful shutdown should be
implemented and understood before introducing infrastructure that abstracts them away.
