```bash
➜  ~ k describe cm configmap-demo
Name:         configmap-demo
Namespace:    default
Labels:       <none>
Annotations:  <none>

Data
====
file.txt:
----
name=John
age=30

key:
----
value

key2:
----
value2

mysql.conf:
----
color.good=blue
color.good=red
color.bad=ok


BinaryData
====

Events:  <none>
```