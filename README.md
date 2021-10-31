## enter

A CLI for generating [ER (entity-relationship) diagrams](https://en.wikipedia.org/wiki/Entity%E2%80%93relationship_model)
for [Ent schema](https://entgo.io/docs/schema-def) using [mermaid.js](https://mermaid-js.github.io).

### Install

```
go get -u github.com/a8m/enter
```

### Example

```
enter 
```

Pass a custom schema path (defaults to `ent/schema`):

```
enter ./pkg/ent/schema
```

Open `er.html` in web browser to view the diagram:

![ERD](https://user-images.githubusercontent.com/7413593/113307613-00f7d800-930e-11eb-8d22-0627b5dfd41d.png)
