# GoWork

## 1. interface test
** tips ** [jordanorelli](https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go)
>* create abstractions by considering the functionality that is common between datatypes, instead of the fields that are common between datatypes
* an interface{} value is not of any type; it is of interface{} type
interfaces are two words wide; schematically they look like (type, value)
it is better to accept an interface{} value than it is to return an interface{} value
* a pointer type may call the methods of its associated value type, but not vice versa
* everything is pass by value, even the receiver of a method
* an interface value isn’t strictly a pointer or not a pointer, it’s just an interface
* if you need to completely overwrite a value inside of a method, use the * operator to manually dereference a pointer