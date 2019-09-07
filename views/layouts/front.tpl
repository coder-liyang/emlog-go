<!DOCTYPE html>
<html>
<head>
    <title>{{.site_title}}</title>

    <!-- meta -->
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="keywords" content="{{.site_key}}" />
    <meta name="description" content="{{.site_description}}" />

    <!-- css -->
    <link rel="stylesheet" href="/static/css/bootstrap.min.css">
    <link rel="stylesheet" href="/static/css/ionicons.min.css">
    <link rel="stylesheet" href="/static/css/pace.css">
    <link rel="stylesheet" href="/static/css/custom.css">

    <!-- js -->
    <script src="/static/js/jquery-2.1.3.min.js"></script>
    <script src="/static/js/bootstrap.min.js"></script>
    <script src="/static/js/pace.min.js"></script>
    <script src="/static/js/modernizr.custom.js"></script>
</head>

<body>
<div class="container">
    <header id="site-header">
        <div class="row">
            <div class="col-md-4 col-sm-5 col-xs-8">
                <div class="logo">
                    <h1><a href="/">{{.blogname}}</a></h1>
                </div>
                <sub>{{.bloginfo}}</sub>
            </div><!-- col-md-4 -->
            <div class="col-md-8 col-sm-7 col-xs-4">
                <nav class="main-nav" role="navigation">
                    <div class="navbar-header">
                        <button type="button" id="trigger-overlay" class="navbar-toggle">
                            <span class="ion-navicon"></span>
                        </button>
                    </div>

                    <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
                        <ul class="nav navbar-nav navbar-right">
                            <li class="cl-effect-11"><a href="index.html" data-hover="首页">首页</a></li>
                            <li class="cl-effect-11"><a href="full-width.html" data-hover="博客">博客</a></li>
                            <li class="cl-effect-11"><a href="/about" data-hover="关于">关于</a></li>
                            <li class="cl-effect-11"><a href="contact.html" data-hover="联系">联系</a></li>
                        </ul>
                    </div><!-- /.navbar-collapse -->
                </nav>
                <div id="header-search-box">
                    <a id="search-menu" href="#"><span id="search-icon" class="ion-ios-search-strong"></span></a>
                    <div id="search-form" class="search-form">
                        <form role="search" method="get" id="searchform" action="#">
                            <input type="search" placeholder="搜索" required>
                            <button type="submit"><span class="ion-ios-search-strong"></span></button>
                        </form>
                    </div>
                </div>
            </div><!-- col-md-8 -->
        </div>
    </header>
</div>
<div class="copyrights">Collect from <a href="http://www.cssmoban.com/">网页模板</a></div>
<div class="copyrights">Modified by <a href="http://zmrenwu.com/">追梦人物的博客</a></div>

<div class="content-body">
    <div class="container">
        <div class="row">
            <main class="col-md-8">
                {{.LayoutContent}}
            </main>
            <aside class="col-md-4">
                <div class="widget widget-recent-posts">
                    <h3 class="widget-title">最新文章</h3>
                    <ul>
                        <li>
                            <a href="#">Django 博客开发入门教程：前言</a>
                        </li>
                        <li>
                            <a href="#">Django 博客使用 Markdown 自动生成文章目录</a>
                        </li>
                        <li>
                            <a href="#">部署 Django 博客</a>
                        </li>
                    </ul>
                </div>
                <div class="widget widget-archives">
                    <h3 class="widget-title">归档</h3>
                    <ul>
                        <li>
                            <a href="#">2017 年 5 月</a>
                        </li>
                        <li>
                            <a href="#">2017 年 4 月</a>
                        </li>
                        <li>
                            <a href="#">2017 年 3 月</a>
                        </li>
                    </ul>
                </div>

                <div class="widget widget-category">
                    <h3 class="widget-title">分类</h3>
                    <ul>
                        <li>
                            <a href="#">Django 博客教程 <span class="post-count">(13)</span></a>
                        </li>
                        <li>
                            <a href="#">Python 教程 <span class="post-count">(11)</span></a>
                        </li>
                        <li>
                            <a href="#">Django 用户认证 <span class="post-count">(8)</span></a>
                        </li>
                    </ul>
                </div>

                <div class="widget widget-tag-cloud">
                    <h3 class="widget-title">标签云</h3>
                    <ul>
                        <li>
                            <a href="#">Django</a>
                        </li>
                        <li>
                            <a href="#">Python</a>
                        </li>
                        <li>
                            <a href="#">Java</a>
                        </li>
                        <li>
                            <a href="#">笔记</a>
                        </li>
                        <li>
                            <a href="#">文档</a>
                        </li>
                        <li>
                            <a href="#">AngularJS</a>
                        </li>
                        <li>
                            <a href="#">CSS</a>
                        </li>
                        <li>
                            <a href="#">JavaScript</a>
                        </li>
                        <li>
                            <a href="#">Snippet</a>
                        </li>
                        <li>
                            <a href="#">jQuery</a>
                        </li>
                    </ul>
                </div>
                <div class="rss">
                    <a href=""><span class="ion-social-rss-outline"></span> RSS 订阅</a>
                </div>
            </aside>
        </div>
    </div>
</div>
<footer id="site-footer">
    <div class="container">
        <div class="row">
            <div class="col-md-12">
                <p class="copyright">
                    &copy 2019 - <a href="http://www.beian.miit.gov.cn/" target="_blank" title="">{{.icp}}</a>
                </p>
            </div>
        </div>
    </div>
</footer>

<!-- Mobile Menu -->
<div class="overlay overlay-hugeinc">
    <button type="button" class="overlay-close"><span class="ion-ios-close-empty"></span></button>
    <nav>
        <ul>
            <li><a href="index.html">首页</a></li>
            <li><a href="full-width.html">博客</a></li>
            <li><a href="about.html">关于</a></li>
            <li><a href="contact.html">联系</a></li>
        </ul>
    </nav>
</div>

<script src="/static/js/script.js"></script>

</body>
</html>