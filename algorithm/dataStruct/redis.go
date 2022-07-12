package dataStruct

/*
redis中各种数据类型的实现
redis只有键=>值，键只能是字符串，值的类型有字符串、列表、字典、集合、有序集合。
为了提高效率，不同的值类型，redis设计了一些数据结构来实现

list: 压缩列表（ziplist) + 双向循环链表
*/
