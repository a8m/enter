## enter

A CLI for generating [ER (entity-relationship) diagrams](https://en.wikipedia.org/wiki/Entity%E2%80%93relationship_model)
for [Ent schema](https://entgo.io/docs/schema-def) using [mermaid.js](https://mermaid-js.github.io).

### Install

```
go get -u github.com/a8m/enter
```

### Examples:

```
enter 
```

Pass custom package schema path (defaults to `ent/schema`):

```
enter ./pkg/ent/schema
```
