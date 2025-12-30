# Roadmap: Semantic Search Database for SML

A step-by-step plan to implement a semantic search database for the Small Model Language (SML) in pure Go.

## 1. Requirements & Goals
- Define what should be searchable (models, entities, attributes, methods, documentation)
- Specify use cases (e.g., find similar models, search by concept, retrieve related documentation)

## 2. Data Model & Storage
- Design a data model for storing SML entities and metadata
- Choose a storage solution (in-memory, file-based, or embedded DB like BadgerDB)

## 3. Indexing
- Implement an indexing mechanism for fast retrieval
- Store raw text and metadata for each model/entity

## 4. Embedding & Representation
- Select or implement a Go-based text embedding method (e.g., TF-IDF, Bag-of-Words, or use pre-trained embeddings compatible with Go)
- Convert model texts and queries into vector representations

## 5. Similarity Search
- Implement similarity metrics (e.g., cosine similarity, Euclidean distance)
- Build a search function to compare query vectors with indexed vectors

## 6. API Design
- Design a Go API for semantic search queries (search, add, update, delete)
- Provide functions for ranking and filtering results

## 7. CLI & Integration
- Integrate semantic search into the SML CLI tool
- Allow users to perform semantic queries from the command line

## 8. Testing & Evaluation
- Write unit and integration tests for all components
- Evaluate search quality and performance

## 9. Documentation
- Document the architecture, API usage, and CLI commands
- Provide examples and best practices

---
*Last updated: December 30, 2025*
