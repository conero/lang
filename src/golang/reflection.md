# reflection 反射机制



## 反射规则

1. Reflection goes from interface value to reflection object.
2. Reflection goes from reflection object to interface value.
3. To modify a reflection object, the value must be settable.





## 注意项

- *reflect* 用于调用struct的方法时属于外部调用，因此命名上一定符合此规则，即首字母大写(公共)
  - 私有方法可能无法调用



## 参考资料

- [laws-of-reflection](https://blog.golang.org/laws-of-reflection)

