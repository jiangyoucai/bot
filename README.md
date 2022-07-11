# funbot

数据爬取集合

1. 「民政部」全国行政区划信息

    包含省、市、县行政区划、面积、人口、邮编、区号等

    - 数据来源
    
        http://xzqh.mca.gov.cn/map

    - 采集源码
    
        https://github.com/fundipper/funbot/tree/main/mca

    - 在线访问
    
        https://www.ipanpan.com/query/place

    - 数据下载（免费）
    
        稍候提供
    
    - 接口调用（免费）

        ```
        Method: POST
        
        URL: https://www.ipanpan.com/v7/query/place

        Body: {
            keyword: 城市名称
        }
        ```
2. 「国家统计局」统计用区划和城乡划分代码

    包含省、市、县/区、乡/镇/街道、村/居委会区划代码，名称

    - 数据来源
    
        http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2021/index.html

    - 采集源码
    
        https://github.com/fundipper/funbot/tree/main/stats

    - 在线访问
    
        https://www.ipanpan.com/place/110000

    - 数据下载（免费）
    
        稍候提供
    
    - 接口调用（免费）

        ```
        Method: GET
        
        URL: https://www.ipanpan.com/v7/place/{城市编码}

        ```
        


