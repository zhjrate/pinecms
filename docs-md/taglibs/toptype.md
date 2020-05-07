# 顶级栏目标签
|功能| 描述|
| :------------- |:-------------|
| 作用      | 获取当前栏目的顶级栏目对象 |
| 备注 | 只可使用在`列表页` 或 `详情页`, 如果当前栏目为一级栏目, 则返回栏目本身 |   

> 重复调用会本次页面缓存

# 标签暴露变量
|变量| 描述|
| :------------- |:-------------|
| field | 一条栏目信息(不包括content) |

# 实例说明
### 获取当前页面的顶级栏目 
```html
{{yield toptype() content}}
    {{field.Catname}} 分类
{{end}}
```