# auto-apidoc
## 介绍
auto-apidoc是一款自动生成接口文档的工具，只有一个可执行文件，java开发者基本不需要额外配置直接可以使用，非常便利；文档输出格式为markdown；支持单体工程和聚合工程。

设计之初想到的就是swagger，但是swagger需要复杂的配置以及接口上需要添加一大堆swagger的注解，所以排除；经过调研发现开源软件JApiDocs，JApiDocs是基于java doc注解的api自动生成工具。但是JApiDocs输出内容为静态页面，同时api的参数解析有一些不符合预期，同时项目只能支持单项目模式（即所有内容都在一个项目中）。

所以有了auto-apidoc；
auto-apidoc受到JApiDocs启发，使用go包装java解析程序，最终将接口文档生成为markdown格式，可以二次编辑；如果有感兴趣的朋友可以自行开发其他语言的解析器。markdown文档为json结构化数据渲染得来的，所以也可以对接口的格式化数据做更多的处理，比如同步文档服务平台系统等等。

## 使用说明
自行编译或者使用main 路径下的autojapi.exe，以下示例内容内容均使用autojapi.exe
1. 使用命令行在项目根目录下执行 `autojapi.exe`
  1. 初次使用
    需要设置API导出路径（之后生成好的文档原始数据均保存该路径下）
    若没有设置maven的环境变量，需要设置mavenHome和Setting file的路径(直接copy idea中的配置即可)
  2. 非初次使用
    直接执行即可，没有额外操作
2. 代码中需要改动
需要在方法备注用使用 `@param `来保证参数会解析到文档中，没有使用注解标注的变量则会被忽略。
返回值推荐使用范类，不然内部属性不会被解析
接受参数推荐使用实体，这样可以逻辑清晰，工具会使用字段注解作为描述使用；
如果参数较少，可以用作为 `net.sf.json.JSONObject` 参数接收，不过这样描述就会是变量名，后续自行修改
> 注意
> request body 必须使用实体接收参数
> return 必须使用范类
```java
// 示例代码
/**
 * 新增接口
 *
 * @param convert
 */
@RequestMapping(value = "/add", produces = "application/json;charset=utf-8", method = RequestMethod.POST)
public MsgVo<Object> add(@RequestBody AdAccountConvert convert) {
    // business do
    return MsgVo.success(object);
}
 
 /**
 * 根据id获取详情
 *
 * @param reqJsonObj
 * @return
 */
@RequestMapping(value = "/get", produces = "application/json;charset=utf-8", method = RequestMethod.POST)
public MsgVo<AdAccountConvert> get(@RequestBody JSONObject reqJsonObj) {
    Integer id = reqJsonObj.containsKey("convertId") ? reqJsonObj.getInt("convertId") : null;
    // business do
    return MsgVo.success(convert);
}
```
### 额外功能
#### 1. 帮助
  autojapi.exe -h
#### 2. 以需求维度生成接口文档
  autojapi.exe -r
  1. 输入需求名称
  2. 选择需要生成的controller
  3. 选择需要生成的接口方法(如果都是你要的直接输入a就行了)
  4. 得到你想要的接口文档了
#### 3. 对比历史，只输出接口发生变化（任何变化）的接口内容
  1. autojapi.exe -c
  2. sohujapi.exe -c52 （与52版本对比）
    1. 使用与-c一致
#### 4. 需求维度和对比历史可以同时使用
autojapi.exe -r -c52
