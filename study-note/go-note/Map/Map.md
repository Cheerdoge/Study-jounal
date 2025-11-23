怀疑自己水平ing，还是有点记不清楚的和没咋练得记下来
1. `mapAssigned = mapLit` 表示`mapAssigned`引用了`mapLit` 对前者修改会影响到后者的值
2. map增长到容量上限时，大小自动加1
3. 一key对多值，考虑把值定义为切片类型
4. key值不存在时，调用会自动得到一个值类型空值。
	如想知道是否存在，可以`_,ok=map[key]` ，`ok` 得到一个bool值
5. `delete(map, key)` 即使key本来就不存在也不会报错
6. 使用range遍历map时获得得值是一个拷贝，修改不会对原本的值产生影响