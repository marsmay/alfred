# China finance workflow for Alfred4 
# 简介

因为 macos 自带的股票软件不实时，延误了 15 分钟，而且看不到基金，于是自己做了一个能够实时查看股票基金行情的 Alfred 插件。

点这里下载打包好的插件，[插件下载](https://github.com/marsmay/alfred/blob/master/src/finance/Finance.alfredworkflow?raw=true)。



# 功能

- 查询中国市场股票、基金实时行情；

- 维护自选列表，显示自选列表所有股票基金实时行情；

- 设置自选列表排序，设置 Quicklook 图表类型；

  

# 数据

数据来源于 [新浪财经](https://hq.sinajs.cn)  接口，均为实时查询；

行情数据特别是基金实时估值，均来源于新浪财经，不对准确性做任何保证；



# 指令

## 股票基金代码

- **沪市股票**：`sh` 前缀加 **6** 位数字代码，如 **上证指数**  `sh000001`;
- **深市股票**：`sz` 前缀加 **6** 位数字代码，如 **深圳成指**  `sz399001`;
- **开放式基金**：`fu` 前缀加 **6** 位数字代码，如  **易方达沪深300ETF联接C**  `fu007339`;



## 自选列表维护指令

- 添加股票、基金到自选列表： `fin add code`，例如：`fin add sh000001`；

- 删除股票、基金到自选列表： `fin del code`，例如：`fin del sh000001`；
- 清空自选列表： `fin clean`



## 自选列表排序维护指令

- 设置自选列表行情显示排序： `fin sort sort_by`，例如：`fin sort va`；
  - `va`：**value asc**，按当前净值（价格）从低到高;
  - `vd`：**value desc**，按当前净值（价格）从高到低;
  - `ra`：**gain ratio asc**，按当前涨跌幅从低到高;
  - `rd`：**gain ratio desc**，按当前涨跌幅从高到低;
  - `nn`：**none none**，不实用自定义排序，默认按列表添加顺序，新添加的在最前;



## Quicklook 图表类型维护指令

> 在选中某一条行情信息时，单击 **shift** 可激活 **Quicklook** 功能，显示当前项目的行情走势图；

- 设置 **Quicklook** 显示的图表类型： `fin chart chart_type`，例如：`fin chart daily`；

  - `daily`：近期行情走势图，以天为单位;
![chart_daily](https://github.com/marsmay/alfred/blob/master/src/finance/images/chart_daily.jpg?raw=true)

  - `today`：当天行情走势图，以分钟为单位;
![chart_today](https://github.com/marsmay/alfred/blob/master/src/finance/images/chart_today.jpg?raw=true)
    
    
    

## 查询指令

- 查询指定股票、基金行情：`fin code`，例如：`fin sh000001`；

![search](https://github.com/marsmay/alfred/blob/master/src/finance/images/search.jpg?raw=true)

- 查询自选列表行情：`fin`

![list](https://github.com/marsmay/alfred/blob/master/src/finance/images/list.jpg?raw=true)



# 自定义

需要进行自定义调整或添加功能，可以直接修改 `src/main.go`，并使用 `GO111MODULE="off" go build -o fin main.go` 指令编译可执行文件，并替换 `workflow` 中的可执行文件 `fin`；



# 已知问题

- 提示未知开发者或无法执行可执行文件 fin

  > 进入 alfred 设置 Workflows 中，找到 China Finance，右键选择 Open in Finder，找到可执行文件 fin，右键选择打开；
  >
  > 执行一次后，不会再用类似提示，可正常使用；
  >
  > 或者自己使用源代码编译 fin，替换  Workflow 文件夹中的可执行文件；

  