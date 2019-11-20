# c++读书笔记
## 类型与申明
- __auto类型修饰符__

- [x] 自动推断类型
```cpp
int a1  = 123;
auto a2 = 123;  //auto 自动将a2的类型推断为int
```

- [x] 声明语句有=号时，auto应该避免使用{}
```cpp
template <class T, class U>
auto operator + (const Matrix<T>& a, const Matrix<U>& b) -> Matrix<decltype(T{}+U{})>   // ->为后置返回类型语法 
             //使用decltype()来隐式声明返回值类型
{
    Matrix<ecltype(T{}+U{})> res;
    for (int i = 0; i != a.rows(); ++i)
    {
        for (int j = 0; j != a.cols(); ++j)
        {
            res(i,j) = a(i,j) + b(i,j);
        }
        return res;
    }
}

