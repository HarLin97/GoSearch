# GoSearch

全库全文检索

# 环境

# 运行

``
chmod 777 search ./search
``

同级目录创建 config.yml

    # 扫描的文件夹
    rootdirs:
    - /root
    # 正则规则
    regulars:
    - description: email
      expression: ^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$
    - description: idCard
      expression: ^\d{15}|\d{18}$
    - description: ip
      expression: ^(.*)\d+\.\d+\.\d+\.\d+(.*)\n
    # 递归扫描
    search: all
    # 文件后缀
    suffix:
    - txt
