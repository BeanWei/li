## li-engine
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