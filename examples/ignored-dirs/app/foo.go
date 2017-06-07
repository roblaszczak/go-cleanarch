package app

// Foo shouldn't be imported from app layer to domain layer, but all dirs with imports Foo are ignored.
type Foo int
