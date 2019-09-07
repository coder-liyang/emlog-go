{{range .blogs}}
    <article class="post post-1">
        <header class="entry-header">
            <h1 class="entry-title">
                <a href="/{{.Gid}}.html">{{.Title}}</a>
            </h1>
            <div class="entry-meta">
                <span class="post-category"><a href="#">Django 博客教程</a></span>
                <span class="post-date"><a href="#"><time class="entry-date" datetime="{{ .Date | convertTime }}">{{ .Date | convertTime }}</time></a></span>
                <span class="post-author"><a href="#">{{.User.Nickname}}</a></span><!-- 要显示作者姓名 -->
                <span class="comments-link"><a href="#">4 评论</a></span>
                <span class="views-count"><a href="#">{{.Views}} 阅读</a></span>
            </div>
        </header>
        <div class="entry-content clearfix">
            {{or .Excerpt .Content | str2html}}
            <div class="read-more cl-effect-14">
                <a href="/{{.Gid}}.html" class="more-link">继续阅读 <span class="meta-nav">→</span></a>
            </div>
        </div>
    </article>
{{end}}
<!-- 简单分页效果
                <div class="pagination-simple">
                    <a href="#">上一页</a>
                    <span class="current">第 6 页 / 共 11 页</span>
                    <a href="#">下一页</a>
                </div>
                -->
<!-- <div class="pagination">
    <ul>
        <li><a href="">1</a></li>
        <li><a href="">...</a></li>
        <li><a href="">4</a></li>
        <li><a href="">5</a></li>
        <li class="current"><a href="">6</a></li>
        <li><a href="">7</a></li>
        <li><a href="">8</a></li>
        <li><a href="">...</a></li>
        <li><a href="">11</a></li>
    </ul>
</div>
-->
{{if gt .paginator.PageNums 1}}
    <ul class="pagination pagination-sm">
        {{if .paginator.HasPrev}}
            <li><a href="{{.paginator.PageLinkFirst}}">首页</a></li>
            <li><a href="{{.paginator.PageLinkPrev}}">&lt;</a></li>
        {{else}}
            <li class="disabled"><a>首页</a></li>
            <li class="disabled"><a>&lt;</a></li>
        {{end}}
        {{range $index, $page := .paginator.Pages}}
            <li{{if $.paginator.IsActive .}} class="active"{{end}}>
                <a href="{{$.paginator.PageLink $page}}">{{$page}}</a>
            </li>
        {{end}}
        {{if .paginator.HasNext}}
            <li><a href="{{.paginator.PageLinkNext}}">&gt;</a></li>
            <li><a href="{{.paginator.PageLinkLast}}">尾页</a></li>
        {{else}}
            <li class="disabled"><a>&gt;</a></li>
            <li class="disabled"><a>尾页</a></li>
        {{end}}
    </ul>
{{end}}