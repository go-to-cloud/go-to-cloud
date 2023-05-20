import{_ as s,c as a,o as l,N as n}from"./chunks/framework.cd9250a1.js";const h=JSON.parse('{"title":"二进制部署","description":"","frontmatter":{},"headers":[],"relativePath":"installation/binary.md","lastUpdated":1679836149000}'),o={name:"installation/binary.md"},e=n(`<h1 id="二进制部署" tabindex="-1">二进制部署 <a class="header-anchor" href="#二进制部署" aria-label="Permalink to &quot;二进制部署&quot;">​</a></h1><h2 id="前置准备" tabindex="-1">前置准备 <a class="header-anchor" href="#前置准备" aria-label="Permalink to &quot;前置准备&quot;">​</a></h2><ul><li>Mysql <code>8.0+</code></li><li>K8S <code>1.18+</code></li><li>Registry <code>V2</code></li></ul><div class="tip custom-block"><p class="custom-block-title">Registry V2</p><p>正式环境可以使用sonatype提供的<a href="https://github.com/sonatype/docker-nexus3" target="_blank" rel="noreferrer">nexus</a></p><p>体验环境可以使用Docker提供的镜像<a href="https://hub.docker.com/_/registry/tags" target="_blank" rel="noreferrer">registry:2</a>，如下：</p><div class="language-shell"><button title="Copy Code" class="copy"></button><span class="lang">shell</span><pre class="shiki material-theme-palenight"><code><span class="line"><span style="color:#FFCB6B;">docker</span><span style="color:#A6ACCD;"> </span><span style="color:#C3E88D;">run</span><span style="color:#A6ACCD;"> </span><span style="color:#C3E88D;">-d</span><span style="color:#A6ACCD;"> </span><span style="color:#C3E88D;">-p</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">5000</span><span style="color:#C3E88D;">:</span><span style="color:#F78C6C;">5000</span><span style="color:#A6ACCD;"> </span><span style="color:#C3E88D;">--restart</span><span style="color:#A6ACCD;"> </span><span style="color:#C3E88D;">always</span><span style="color:#A6ACCD;"> </span><span style="color:#C3E88D;">--name</span><span style="color:#A6ACCD;"> </span><span style="color:#C3E88D;">registry</span><span style="color:#A6ACCD;"> </span><span style="color:#C3E88D;">registry:</span><span style="color:#F78C6C;">2</span></span>
<span class="line"></span></code></pre></div></div><h2 id="下载地址" tabindex="-1">下载地址 <a class="header-anchor" href="#下载地址" aria-label="Permalink to &quot;下载地址&quot;">​</a></h2><div class="vp-code-group"><div class="tabs"><input type="radio" name="group-seMmk" id="tab-SRv24ML" checked="checked"><label for="tab-SRv24ML">macos arm64</label><input type="radio" name="group-seMmk" id="tab-9kI04UF"><label for="tab-9kI04UF">macos x64</label><input type="radio" name="group-seMmk" id="tab-aA1LQ0c"><label for="tab-aA1LQ0c">linux x64</label><input type="radio" name="group-seMmk" id="tab-FSvz9Om"><label for="tab-FSvz9Om">windows</label></div><div class="blocks"><div class="language-shell active"><button title="Copy Code" class="copy"></button><span class="lang">shell</span><pre class="shiki material-theme-palenight"><code><span class="line"><span style="color:#FFCB6B;">wget</span><span style="color:#A6ACCD;"> </span><span style="color:#C3E88D;">https://github.com/go-to-cloud/go-to-cloud/releases/download/1.0.0-beta/gotocloud-macos_arm64.tar.gz</span></span>
<span class="line"></span></code></pre></div><div class="language-shell"><button title="Copy Code" class="copy"></button><span class="lang">shell</span><pre class="shiki material-theme-palenight"><code><span class="line"><span style="color:#FFCB6B;">wget</span><span style="color:#A6ACCD;"> </span><span style="color:#C3E88D;">https://github.com/go-to-cloud/go-to-cloud/releases/download/1.0.0-beta/gotocloud-macos_x64.tar.gz</span></span>
<span class="line"></span></code></pre></div><div class="language-shell"><button title="Copy Code" class="copy"></button><span class="lang">shell</span><pre class="shiki material-theme-palenight"><code><span class="line"><span style="color:#FFCB6B;">wget</span><span style="color:#A6ACCD;"> </span><span style="color:#C3E88D;">https://github.com/go-to-cloud/go-to-cloud/releases/download/1.0.0-beta/gotocloud-linux_x64.tar.gz</span></span>
<span class="line"></span></code></pre></div><div class="language-shell"><button title="Copy Code" class="copy"></button><span class="lang">shell</span><pre class="shiki material-theme-palenight"><code><span class="line"><span style="color:#676E95;font-style:italic;"># 下载地址</span></span>
<span class="line"><span style="color:#FFCB6B;">wget</span><span style="color:#A6ACCD;"> </span><span style="color:#C3E88D;">https://github.com/go-to-cloud/go-to-cloud/releases/download/1.0.0-beta/gotocloud-windows_x64.tar.gz</span></span>
<span class="line"></span></code></pre></div></div></div><h3 id="解压文件" tabindex="-1">解压文件 <a class="header-anchor" href="#解压文件" aria-label="Permalink to &quot;解压文件&quot;">​</a></h3><div class="language-shell"><button title="Copy Code" class="copy"></button><span class="lang">shell</span><pre class="shiki material-theme-palenight"><code><span class="line"><span style="color:#676E95;font-style:italic;">#!bin/bash</span></span>
<span class="line"><span style="color:#FFCB6B;">tar</span><span style="color:#A6ACCD;"> </span><span style="color:#C3E88D;">-zxvf</span><span style="color:#A6ACCD;"> </span><span style="color:#C3E88D;">./gotocloud-</span><span style="color:#89DDFF;">&lt;</span><span style="color:#C3E88D;">目标平</span><span style="color:#A6ACCD;">台</span><span style="color:#89DDFF;">&gt;</span><span style="color:#C3E88D;">.tar.gz</span></span>
<span class="line"></span></code></pre></div><h3 id="修改配置" tabindex="-1">修改配置 <a class="header-anchor" href="#修改配置" aria-label="Permalink to &quot;修改配置&quot;">​</a></h3><div class="language-yaml"><button title="Copy Code" class="copy"></button><span class="lang">yaml</span><pre class="shiki material-theme-palenight"><code><span class="line"><span style="color:#F07178;">db</span><span style="color:#89DDFF;">:</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;"># 数据库配置</span></span>
<span class="line"><span style="color:#A6ACCD;">  </span><span style="color:#F07178;">user</span><span style="color:#89DDFF;">:</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;"># 数据库用户名</span></span>
<span class="line"><span style="color:#A6ACCD;">  </span><span style="color:#F07178;">password</span><span style="color:#89DDFF;">:</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;"># 数据库密码</span></span>
<span class="line"><span style="color:#A6ACCD;">  </span><span style="color:#F07178;">host</span><span style="color:#89DDFF;">:</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;"># 数据库地址</span></span>
<span class="line"><span style="color:#A6ACCD;">  </span><span style="color:#F07178;">schema</span><span style="color:#89DDFF;">:</span><span style="color:#A6ACCD;"> </span><span style="color:#C3E88D;">gotocloud</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;"># 数据库名，默认不用修改</span></span>
<span class="line"><span style="color:#F07178;">jwt</span><span style="color:#89DDFF;">:</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;"># jwt配置，建议修改security，其他保持不变</span></span>
<span class="line"><span style="color:#A6ACCD;">  </span><span style="color:#F07178;">security</span><span style="color:#89DDFF;">:</span><span style="color:#A6ACCD;"> </span><span style="color:#C3E88D;">thisisunsafeuntilyouchangit</span></span>
<span class="line"><span style="color:#A6ACCD;">  </span><span style="color:#F07178;">realm</span><span style="color:#89DDFF;">:</span><span style="color:#A6ACCD;"> </span><span style="color:#C3E88D;">GOTOCLOUD</span></span>
<span class="line"><span style="color:#A6ACCD;">  </span><span style="color:#F07178;">idkey</span><span style="color:#89DDFF;">:</span><span style="color:#A6ACCD;"> </span><span style="color:#C3E88D;">id</span></span>
<span class="line"><span style="color:#F07178;">builder</span><span style="color:#89DDFF;">:</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;"># 打包镜像的服务，一般不用修改</span></span>
<span class="line"><span style="color:#A6ACCD;">  </span><span style="color:#F07178;">kaniko</span><span style="color:#89DDFF;">:</span><span style="color:#A6ACCD;"> </span><span style="color:#C3E88D;">go-to-cloud-docker.pkg.coding.net/devops/kaniko/executor:v1.9.1-debug</span></span>
<span class="line"></span></code></pre></div><div class="info custom-block"><p class="custom-block-title">INFO</p><p>请参照注释做相应的配置</p></div><h2 id="运行程序" tabindex="-1">运行程序 <a class="header-anchor" href="#运行程序" aria-label="Permalink to &quot;运行程序&quot;">​</a></h2><div class="vp-code-group"><div class="tabs"><input type="radio" name="group-GCR9A" id="tab-Y_U3lZ6" checked="checked"><label for="tab-Y_U3lZ6">默认方式</label><input type="radio" name="group-GCR9A" id="tab-NDWM2-p"><label for="tab-NDWM2-p">指定端口</label></div><div class="blocks"><div class="language-shell active"><button title="Copy Code" class="copy"></button><span class="lang">shell</span><pre class="shiki material-theme-palenight"><code><span class="line"><span style="color:#89DDFF;">&lt;</span><span style="color:#A6ACCD;">二进制文件目录</span><span style="color:#89DDFF;">&gt;</span><span style="color:#A6ACCD;">/gotocloud</span></span>
<span class="line"></span></code></pre></div><div class="language-shell"><button title="Copy Code" class="copy"></button><span class="lang">shell</span><pre class="shiki material-theme-palenight"><code><span class="line"><span style="color:#89DDFF;">&lt;</span><span style="color:#A6ACCD;">二进制文件目录</span><span style="color:#89DDFF;">&gt;</span><span style="color:#A6ACCD;">/gotocloud -port </span><span style="color:#F78C6C;">8080</span></span>
<span class="line"></span></code></pre></div></div></div>`,13),p=[e];function t(c,r,i,y,d,C){return l(),a("div",null,p)}const u=s(o,[["render",t]]);export{h as __pageData,u as default};
