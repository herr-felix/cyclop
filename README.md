# Cyclop

Cyclop is a tree-based entity store that verifies the complete integrity of an entity on save.

## Entity

- An entity has an `ID`, that is unique between all entities.

- An entity has an `EntityType`, to know wht set of properties to expect from this entity.

- An entity has a `Data` field, that contains all the properties of an entity encoded with MsgPack.

## Warning

Each entity types has a method `Inspect` for verifying the integrity of a given entity with a custom business logic.

If `Inspect` returns a list _warnings_, which is empty if the entity is in perfect condition.

All the warning are put into a centralized store that can be looked up with the entity's ID.



## Storing the Data

All the entities (and their relations to each others) and warnings are stored in a relational database (SQL Server, PostgreSQL, MySQL, etc...).

### `Types`
- `ID`: ID of the type 

### `Entities`
- `ID`: ID of an entity
- `Type`: Foreign key to `Types.ID`
- `Data`: MsgPack blob

### `Links`
- `ParentID`: FK to `Entities.ID`
- `ChildID`: FK to `Entities.ID`

Every entities and warnings are also cached and can be accessible directly by their ID.

As for searching, entities can be indexed in a full-text-search database, in which they are categorized by their entity types.
