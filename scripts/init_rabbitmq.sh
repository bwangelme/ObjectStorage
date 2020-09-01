
/usr/local/sbin/rabbitmqctl add_vhost /objs
# 如果在生产环境使用，建议使用强度更高的密码
/usr/local/sbin/rabbitmqctl add_user objs objs
# 生产环境的话，不需要添加 administrator 标签
/usr/local/sbin/rabbitmqctl set_user_tags objs administrator
# 注意权限必须用单引号，否则会导致 bash 自动补全成当前文件
/usr/local/sbin/rabbitmqctl set_permissions -p /objs objs '.*' '.*' '.*'
/usr/local/sbin/rabbitmqadmin -V /objs -u objs -p objs declare exchange name=api_servers type=direct
/usr/local/sbin/rabbitmqadmin -V /objs -u objs -p objs declare exchange name=data_servers type=direct


