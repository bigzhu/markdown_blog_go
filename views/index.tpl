<!doctype html>

<html lang="cn">
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1" />

        <title>{{.title}}</title>
        <meta name="description" content="The BigZhu's Blog">
        <meta name="author" content="{{.author}}">

        <link rel="stylesheet" href="/static/css/github-markdown-css/github-markdown.css">
        <link rel="stylesheet" href="/static/css/style.css">
        <style>
            .markdown-body {
                box-sizing: border-box;
                margin: 0 auto;
                padding: 8px;
                position: relative;
            }
            @media screen and (max-width: 600px) {
                .mobile-hidden {
                    display: none !important;
                }
                .copy-right.mobile {
                    right: 13px;
                    bottom: 27px;
                    font-size: 12px;
                }

            }
            @media screen and (min-width: 992px) {
                .markdown-body {
                    box-sizing: border-box;
                    min-width: 200px;
                    max-width: 980px;
                    margin: 0 auto;
                    padding-bottom: 180px;
                }
                #search {
                    float: right;
                }
            }
            /*@media screen and (min-width: 767px) {
                .copy-right {
                    right: 13px;
                    bottom: 27px;
                    font-size: 12px;
                }
            }*/
            html {
                height: 100%;
                min-height: 100%;
            }
            body {
                min-height: 100%;
                position: relative;
                margin: 0;
                padding-bottom: 8rem;
            }
            .pre {
                float: left;
            }
            .old {
                float: right;
            }

            .footer {
                position: absolute;
                bottom: 0;
                width: 100%;
                margin-top: 0px;
                border-top: 5px solid rgba(0, 0, 0, .05);
                background-color: #fff;
            }
            .container {
                padding: 0rem auto;
                display: block;
                max-width: 980px;
                margin: 0 auto;
            }
            .logo {
                margin-top: 0px;
                margin-bottom: 4px;
                padding-bottom: 0px;
                border-style: none;
                -webkit-transition: all 500ms ease;
                transition: all 500ms ease;
                font-family: 'Ubuntu', Helvetica, sans-serif;
                color: rgba(51, 51, 51, .81);
                font-size: 26px;
                font-weight: bolder;
                text-decoration: none;
                cursor: pointer;
            }
            .logo:hover {
                color: #333;
            }
            .web-address {
                display: inline-block;
                margin-top: 0px;
                margin-bottom: 0px;
                border: 0px none #000;
                font-family: Georgia, Times, 'Times New Roman', serif;
                color: #999;
                font-size: 0.9rem;
                text-decoration: none;
                cursor: pointer;
            }
            .logo-web {
                display: inline-block;
                padding: 20px;
            }
            .copy-right {
                position: absolute;
                right: 103px;
                bottom: 32px;
                display: inline-block;
                float: right;
                font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;
                color: rgba(51, 51, 51, .75);
                font-size: 13px;
            }
            .copy-right a {
                color: #407CC4;
                text-decoration: none;
            }
        </style>
        <!--[if lt IE 9]>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/html5shiv/3.7.3/html5shiv.js"></script>
        <![endif]-->
    </head>

    <body>
        <article class="markdown-body">
        <div class="mobile-hidden" id="search">
            <form action="https://www.google.com" onsubmit="return dispatch()" target="_blank">
                <input type="text" maxlength="44" name="q" id="q" value="">
            </form>
        </div>

        <header>
        <h2>{{.title}}</h2>

        <p class="meta">
        <a href="{{.author_link}}" target="blank">{{.author}}</a> <time datetime="{{ .modify_time }}">{{ .modify_time }}</time>
        </p>
        </header>
        {{.toc}}
        {{str2html .content}}
        <hr>
        <b>bigzhu:「我就是在瞎说, 别让我举证」</b>
        <h2>苹果用户专用赞赏二维码</h2>
        <img width="200" src="https://ws3.sinaimg.cn/large/006tNc79gy1fhp9os94xyj30jp0jkjsu.jpg">
        <p>
        {{ if .pre }}
        <a class="pre" href="/{{.quote_pre}}" >←  {{.pre}} </a>
        {{ end }}

        {{ if .old }}
        <a class="old" href="/{{.quote_old}}" >{{.old}} → </a>
        {{ end }}
        </p>
        </article>
        <div class="footer">
            <div class="container">
                <div class="logo-web">
                    <h3 class="logo"><a href="http://www.lorstone.com" class="logo">Lorstone.</a></h3><a class="web-address" href="http://www.lorstone.com">www.lorstone.com</a>
                </div>
                <div class="copy-right mobile">© Copyright 2017</br>Powered by <a href="https://github.com/bigzhu/markdown_blog">markdown_blog</a></div>
            </div>
        </div>
        <script>
            function dispatch() {
                var q = document.getElementById("q");
                if (q.value != "") {
                    var url = 'https://www.google.com/search?q=site:bigzhu.lorstone.com %20' + q.value;
                    if (navigator.userAgent.indexOf('iPad') > -1 || navigator.userAgent.indexOf('iPod') > -1 || navigator.userAgent.indexOf('iPhone') > -1) {
                        location.href = url;
                        } else {
                        window.open(url, "_blank");
                    }
                    return false;
                    } else {
                    return false;
                }
            }
        </script>
        <script src="http://hammerjs.github.io/dist/hammer.min.js"></script>
        <!--
        <script>
            var body = document.getElementsByTagName("body")[0];
            var mc = new Hammer(body);
            mc.on('panleft panright', function(ev) {
                console.log()
                if (ev.velocityX >=0.5) {
                    {% if pre is not None%}
                    window.location = '/'+'{{.quote_pre}}';
                    {% end %}
                }
                if (ev.velocityX <=-0.5) {
                    {% if old is not None%}
                    window.location = '/'+'{{.quote_old}}';
                    {% end %}
                }
            });
        </script>
        -->
    </body>
</html>
