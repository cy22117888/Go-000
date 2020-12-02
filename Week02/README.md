学习笔记
### 2020.12.02 
* 详细的学习笔记后面补，现在Go的语法还没学全，讲的课只是能大概听懂的程度。还要再观看复习两遍，才能来做总结。

#### 下面是 Week02 第一遍学习的体会：
* 感觉 Go 的 Error 机制与 Rust 相似，分为可恢复错误和不可恢复错误。
* 不可恢复错误 panic 会直接终止程序的运行。
* 可恢复错误更像是 Java 异常机制的替代品，有点类似于 Java 的 Option<T> 类，不过 Option<T> 貌似只能判空。最像的还是 Rust 的 Result<T, E> 枚举。
```Rust
Rust 的 Result<T, E> 定义如下：

    enum Result<T, E> {
        Ok(T),
        Err(E),
    }
```
* Go 中的方法需要返回 error 时，需要使用多返回值的形式，将 error 作为其中的一个返回值。
* 当被上层方法调用时，先使用 if 结构处理 可能产生的 error：
    * 当产生了 error 时，开辟分支处理 error，便于方法的提前返回；
    * 当未产生 error 时，方法主线继续执行业务逻辑。
* Go 中对 error 的处理类似于 Rust 中的 match 分支处理。