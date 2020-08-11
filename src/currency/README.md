# China currency workflow for Alfred4 
# 简介

不论是出国换汇，还是信用卡购汇还款，在不同时间操作价格差异会比较大，所以能够实时了解到外汇买卖价格是比较有用的。

对于一般消费者，外汇的市场价格（在岸、离岸、中间价）没有什么意义，真正有用的是中国主流银行的外汇牌价（现钞/现汇买入卖出价格）。

因为时常需要购汇还款，所以做了这个能够实时查看银行购买和出售外汇价格的 Alfred 插件。

点这里下载打包好的插件，[插件下载](https://github.com/marsmay/alfred/blob/master/src/currency/Currency.alfredworkflow?raw=true)。



# 功能

- 查询购汇、结汇、购钞、结钞指定金额外币的各主流银行价格；

- 购买外币使用银行卖出价格，分为购汇和购钞，按价格升序排列；

- 出售外币使用银行买入价格，分为结汇和结钞，按价格降序排列；

- 银行间转汇成本很高，价格仅供参考，可帮助确定还款或换汇最佳时间点；

  

# 数据

数据来源于 [新浪金融](http://vip.stock.finance.sina.com.cn)  接口，均为实时查询，不对准确性做任何保证；



# 使用

## 支持币种

- **USD**：美元;
- **EUR**：欧元;
- **HKD**：港币;
- **JPY**：日元;
- **GBP**：英镑;
- **AUD**：澳大利亚元;
- **CAD**：加拿大元;
- **THB**：泰国铢;
- **SGD**：新加坡元;
- **SEK**：瑞典克朗;
- **CHF**：瑞士法郎;
- **DKK**：丹麦克朗;
- **NOK**：挪威克朗;
- **RUB**：俄罗斯卢布;
- **MOP**：澳门元;
- **PHP**：菲律宾比索;
- **NZD**：新西兰元;
- **KRW**：韩元;
- **MYR**：马来西亚林吉特;
- **ZAR**：南非兰特;
- **TWD**：新台币;
- **BRL**：巴西雷亚尔;



## 支持指令

- `cbc`：**cny buy cash**，购买外币现钞的人民币价格；;
- `cbe`：**cny buy exchange**，购买外币现汇的人民币价格；;
- `csc`：**cny sell cash**，出售外币现钞的人民币价格；;
- `cse`：**cny sell exchange**，出售外币现汇的人民币价格；;



## 询价

- `opt amount currency`，例如：`cbc 200 usd`；

  ![cbc](https://github.com/marsmay/alfred/blob/master/src/currency/images/cbc.jpg?raw=true)



# 自定义

需要进行自定义调整或添加功能，可以直接修改 `src/currency` 中的 **golang** 源代码，并使用 `make currency` 指令编译可执行文件，并替换安装后的 `workflow` 目录中的可执行文件 `currency`；



# 已知问题

- 提示未知开发者或无法执行可执行文件 currency

  > 进入 alfred 设置 Workflows 中，找到 China Currency，右键选择 Open in Finder，找到可执行文件 currency，右键选择打开；
  >
  > 执行一次后，不会再用类似提示，可正常使用；
  >
  > 或者自己使用源代码编译 currency，替换  Workflow 文件夹中的可执行文件；

  