## li-engine

## Entity
```go
type Post struct {
    entity.Schema
}

func (Post) Mixin() []entity.Mixin {
    return []entity.Mixin {

    }
}

func (Post) Annotations() []entity.Annotation {
    return []entity.Annotation{

    }
}

func (Post) Fields() []entity.Field {
    return []entity.Field{
        entity.FieldStr("title"),
    }
}

func (Post) Indexes() []entity.Index {
    return []entity.Index{
        entity.IndexFields("title"),
    }
}
```


## Page
```go
type PostListPage struct {
    page.Schema
}

func (PostListPage) Annotations() []page.Annotation {
    return []page.Annotation{

    }
}

func (PostListPage) Mixin() []page.Mixin {
    return []page.Mixin {

    }
}

func (PostListPage) Nodes() []page.Node {
    return []page.Node{
        page.NodeTable("node_posts").
            RowKey("id").
            RowSelection("checkbox").
            Column(&page.Node{
                page.NodeInput("title"),
            }).
            Default("").
            Query("")
    }
}
```