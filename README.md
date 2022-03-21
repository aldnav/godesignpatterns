# Go Design Patterns

Following the course of Joe Marini on Go Design Patterns

## Creational Patterns

- Builder
- Factory
- Singleton

### Builder pattern

Purpose:

Encapsulates an object's construction process along with specifying the various parts of a complex API

Enables flexible creation of an object that can have many different representations

Increases code readability for complex types

Scenarios:

Objects that have complex APIs, multiple constructor op`tions, and several different possible representations

```txt
          new()           create()
Director ------> Builder ----------> Complex Object
newBuilder()    setPropA()
                setPropB()
                setPropC()
                create()
```

Advantages:

1. Encapsulation of the construction of the internal representation of the complex object
2. Builder provides detailed control over how object is constructed

Disadvantages:

1. Creating a builder for each type of complex object
2. Cause some difficulty to other software engineering techniques like dependency injection
3. Each object must be a mutable object which might not always be desirable

### Factory pattern

Purpose:

Allows for the construction of objects when the types of those objects is not predetermined at runtime

Scenarios:

Produces code that is more readable when there are multiple ways of creating a particular object

Situations where object creation needs to be flexible and cannot be known beforehand

```txt
                      Product
Creator interface    --------->  Product
 + factoryMethod()

       ^
       |
       |              Create
 Concrete Creator    --------->  Concrete Product
 + factoryMethod()
```

Go has no inheritance. Go has the notion of interfaces and structs so we can use these to design a factory pattern.

