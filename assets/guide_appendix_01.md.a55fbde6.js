import{_ as s,c as n,o as a,N as l}from"./chunks/framework.cd9250a1.js";const F=JSON.parse('{"title":"权限代码","description":"","frontmatter":{},"headers":[],"relativePath":"guide/appendix_01.md","lastUpdated":1679753726000}'),o={name:"guide/appendix_01.md"},p=l(`<h1 id="权限代码" tabindex="-1">权限代码 <a class="header-anchor" href="#权限代码" aria-label="Permalink to &quot;权限代码&quot;">​</a></h1><div class="language-go"><button title="Copy Code" class="copy"></button><span class="lang">go</span><pre class="shiki material-theme-palenight"><code><span class="line"><span style="color:#A6ACCD;">	</span><span style="color:#676E95;font-style:italic;">// MainMenuProject 主菜单-&gt;项目</span></span>
<span class="line"><span style="color:#A6ACCD;">	MainMenuProject            AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">10000</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResProjectNew              AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">10100</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 新建项目</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResProjectDelete           AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">10201</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 删除项目</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResProjectUpdate           AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">10202</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 更新项目信息</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResProjectSourceCode       AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">10310</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 管理模块：代码仓库</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResProjectSourceCodeImport AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">10311</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 管理模块：导入代码仓库</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResProjectSourceDelete     AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">10312</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 管理模块：删除代码仓库</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResProjectCI               AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">10320</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 管理模块：持续集成</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResProjectCINew            AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">10321</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 管理模块：新建构建计划</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResProjectCIStart          AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">10322</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 管理模块：开始构建</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResProjectCIHistory        AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">10323</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 管理模块：构建计划</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResProjectCIDelete         AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">10324</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 管理模块：删除计划</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResProjectCD               AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">10330</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 管理模块：持续部署</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResProjectCDNew            AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">10331</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 管理模块：新建部署方案</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResProjectCDStart          AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">10332</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 管理模块：开始部署/重新部署</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResProjectCDHistory        AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">10333</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 管理模块：部署历史</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResProjectCDRollback       AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">10334</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 管理模块：回滚部署</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResProjectCDDelete         AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">10334</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 管理模块：删除部署方案</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResProjectCDMonitor        AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">10335</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 管理模块：转到应用监控</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResProjectArtifacts        AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">10340</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 管理模块：制品仓库</span></span>
<span class="line"></span>
<span class="line"><span style="color:#A6ACCD;">	</span><span style="color:#676E95;font-style:italic;">// MainMenuConfigure 主菜单-&gt;配置</span></span>
<span class="line"><span style="color:#A6ACCD;">	MainMenuConfigure                 AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">20000</span></span>
<span class="line"><span style="color:#A6ACCD;">	SubMenuConfigureCodeRepo          AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">20010</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 代码仓库</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResConfigureCodeRepoBind          AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">20011</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 绑定代码仓库</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResConfigureCodeRepoUpdate        AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">20012</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 编辑代码仓库</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResConfigureCodeRepoRemove        AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">20013</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 移除代码仓库</span></span>
<span class="line"><span style="color:#A6ACCD;">	SubMenuConfigureBuildNode         AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">20020</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 构建节点</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResConfigureBuildNodeBind         AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">20021</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 安装构建节点</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResConfigureBuildNodeUpdate       AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">20022</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 更新构建节点</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResConfigureBuildNodeRemove       AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">20023</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 卸载构建节点</span></span>
<span class="line"><span style="color:#A6ACCD;">	SubMenuConfigureArtifactRepo      AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">20030</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 制品仓库</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResConfigureArtifactRepoBind      AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">20031</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 绑定制品仓库</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResConfigureArtifactRepoUpdate    AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">20032</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 更新制品仓库</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResConfigureArtifactRepoRemove    AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">20033</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 移除制品仓库</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResConfigureArtifactDeleteHistory AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">20034</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 删除历史制品</span></span>
<span class="line"><span style="color:#A6ACCD;">	SubMenuConfigureDeployRepo        AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">20040</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 部署环境</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResConfigureDeployBind            AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">20041</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 绑定部署环境</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResConfigureDeployUpdate          AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">20042</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 更新部署环境</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResConfigureDeployRemove          AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">20043</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 移除部署环境</span></span>
<span class="line"></span>
<span class="line"><span style="color:#A6ACCD;">	</span><span style="color:#676E95;font-style:italic;">// MainMenuMonitor 主菜单-&gt;应用监控</span></span>
<span class="line"><span style="color:#A6ACCD;">	MainMenuMonitor          AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">30000</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResourceMonitorScale     AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">30010</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 伸缩副本</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResourceMonitorRestart   AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">30011</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 重新应用</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResourceMonitorDelete    AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">30012</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 删除应用</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResourceMonitorShell     AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">30013</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 容器Shell</span></span>
<span class="line"><span style="color:#A6ACCD;">	ResourceMonitorDeletePod AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">30014</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 删除容器</span></span>
<span class="line"></span>
<span class="line"><span style="color:#A6ACCD;">	</span><span style="color:#676E95;font-style:italic;">// MainMenuUsers 主菜单-&gt;用户管理</span></span>
<span class="line"><span style="color:#A6ACCD;">	MainMenuUsers AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">40000</span></span>
<span class="line"><span style="color:#A6ACCD;">	SubMenuUser   AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">40010</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 用户管理</span></span>
<span class="line"><span style="color:#A6ACCD;">	SubMenuOrgs   AuthCode </span><span style="color:#89DDFF;">=</span><span style="color:#A6ACCD;"> </span><span style="color:#F78C6C;">40020</span><span style="color:#A6ACCD;"> </span><span style="color:#676E95;font-style:italic;">// 组织管理</span></span>
<span class="line"></span></code></pre></div>`,2),e=[p];function t(C,c,r,y,A,D){return a(),n("div",null,e)}const u=s(o,[["render",t]]);export{F as __pageData,u as default};