# generic-go
This project explore the way to implement a generic type in Go

The generic type `T` in Java will the same as `interface{}` - concrete type

The other generic type as `BaseModel` will need a specific method of model for compiler recognize your generic type, example is `GetId` method in my case.

So whenever you need to pass a model to repo or service layer, and the object need to be a model, we will use as e.g
```
  func ReplaceOne(ctx context.Context, model BaseModel) {
      // replace your object to database
  }
```

The difficult step in this part is your models need to implement method `GetID`, we can use tool to generate, in case your models are following encapslution pattern (has getter setter method) this will be best place for this solution.
