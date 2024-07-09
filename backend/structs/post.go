package structs

type Post struct {
    PostID          int
    UserID          int
    Username        string
    Dislike         int
    Like            int
    PostHeading     string
    Postdescription string
    Categories      [] string
}
