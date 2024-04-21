# IoT-Project

Contains three packages with the following dependency chain:
- mqtt
- broker (uses mqtt)
- publisher (uses mqtt)

To handle subscribers, use either the broker to create individual handlers, or a two-tier broker to separate publisher handling from subscriber handling.
For now, using the former approach.

Example use case of subscriber: database...
